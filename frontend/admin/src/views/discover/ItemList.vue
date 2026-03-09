<template>
  <div class="discover-items">
    <a-card :title="`${moduleTitle}内容管理`" :bordered="false">
      <template #extra>
        <a-button type="primary" @click="handleAdd">添加内容</a-button>
      </template>

      <a-table
        :columns="columns"
        :data-source="itemList"
        :loading="loading"
        row-key="id"
        :pagination="pagination"
        @change="handleTableChange"
      >
        <template #bodyCell="{ column, record }">
          <template v-if="column.key === 'cover_url'">
            <a-image
              v-if="record.cover_url"
              :src="record.cover_url"
              :width="80"
              :height="80"
              style="object-fit: cover; border-radius: 4px;"
            />
            <div
              v-else
              :style="{
                width: '80px',
                height: '80px',
                background: record.color || '#FF6B6B',
                borderRadius: '4px',
                display: 'flex',
                alignItems: 'center',
                justifyContent: 'center',
                color: 'white',
                fontSize: '24px',
              }"
            >
              {{ record.title.slice(0, 2) }}
            </div>
          </template>
          <template v-if="column.key === 'is_enabled'">
            <a-tag :color="record.is_enabled ? 'green' : 'red'">
              {{ record.is_enabled ? '启用' : '禁用' }}
            </a-tag>
          </template>
          <template v-if="column.key === 'action'">
            <a-button type="link" @click="handleEdit(record)">编辑</a-button>
            <a-popconfirm
              title="确定删除吗?"
              @confirm="handleDelete(record.id)"
            >
              <a-button type="link" danger>删除</a-button>
            </a-popconfirm>
          </template>
        </template>
      </a-table>
    </a-card>

    <!-- 添加/编辑弹窗 -->
    <a-modal
      v-model:open="editVisible"
      :title="editForm.id ? '编辑内容' : '添加内容'"
      @ok="handleEditOk"
      width="600px"
    >
      <a-form :model="editForm" :label-col="{ span: 6 }">
        <a-form-item label="标题" required>
          <a-input v-model:value="editForm.title" />
        </a-form-item>
        <a-form-item label="封面图片">
          <a-input v-model:value="editForm.cover_url" placeholder="图片URL" />
        </a-form-item>
        <a-form-item label="背景颜色">
          <a-input v-model:value="editForm.color" placeholder="#FF6B6B" />
        </a-form-item>
        <a-form-item label="作者">
          <a-input v-model:value="editForm.author" />
        </a-form-item>
        <a-form-item label="描述">
          <a-textarea v-model:value="editForm.description" :rows="3" />
        </a-form-item>
        <a-form-item label="播放量">
          <a-input-number v-model:value="editForm.views" :min="0" />
        </a-form-item>
        <a-form-item label="点赞数">
          <a-input-number v-model:value="editForm.likes" :min="0" />
        </a-form-item>
        <a-form-item label="评论数">
          <a-input-number v-model:value="editForm.comments" :min="0" />
        </a-form-item>
        <a-form-item label="粉丝数">
          <a-input-number v-model:value="editForm.fans" :min="0" />
        </a-form-item>
        <a-form-item label="参与数">
          <a-input-number v-model:value="editForm.count" :min="0" />
        </a-form-item>
        <a-form-item label="排序">
          <a-input-number v-model:value="editForm.sort_order" :min="1" />
        </a-form-item>
        <a-form-item label="是否启用">
          <a-switch v-model:checked="editForm.is_enabled" />
        </a-form-item>
      </a-form>
    </a-modal>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, computed } from 'vue';
import { useRoute } from 'vue-router';
import { message } from 'ant-design-vue';
import {
  getDiscoverItemList,
  createDiscoverItem,
  updateDiscoverItem,
  deleteDiscoverItem,
  type DiscoverItem,
} from '@/api/discover';

const route = useRoute();
const module = ref(route.query.module as string || 'hot_topics');
const loading = ref(false);
const itemList = ref<DiscoverItem[]>([]);
const editVisible = ref(false);
const editForm = ref<Partial<DiscoverItem>>({
  module: module.value,
  is_enabled: true,
  sort_order: 1,
});

const pagination = ref({
  current: 1,
  pageSize: 20,
  total: 0,
});

const moduleTitle = computed(() => {
  const titles: Record<string, string> = {
    hot_topics: '热门话题',
    hot_rank: '热门榜单',
    recommend_creators: '推荐创作者',
    guess_you_like: '猜你喜欢',
  };
  return titles[module.value] || module.value;
});

const columns = [
  {
    title: 'ID',
    dataIndex: 'id',
    key: 'id',
    width: 80,
  },
  {
    title: '封面',
    dataIndex: 'cover_url',
    key: 'cover_url',
    width: 100,
  },
  {
    title: '标题',
    dataIndex: 'title',
    key: 'title',
  },
  {
    title: '作者',
    dataIndex: 'author',
    key: 'author',
    width: 120,
  },
  {
    title: '播放量',
    dataIndex: 'views',
    key: 'views',
    width: 100,
  },
  {
    title: '排序',
    dataIndex: 'sort_order',
    key: 'sort_order',
    width: 80,
  },
  {
    title: '状态',
    dataIndex: 'is_enabled',
    key: 'is_enabled',
    width: 100,
  },
  {
    title: '操作',
    key: 'action',
    width: 150,
  },
];

// 加载内容列表
const loadItemList = async () => {
  loading.value = true;
  try {
    const res = await getDiscoverItemList({
      module: module.value,
      page: pagination.value.current,
      page_size: pagination.value.pageSize,
    });
    itemList.value = res.data.list || [];
    pagination.value.total = res.data.total || 0;
  } catch (error) {
    message.error('加载内容失败');
  } finally {
    loading.value = false;
  }
};

// 表格分页改变
const handleTableChange = (pag: any) => {
  pagination.value.current = pag.current;
  loadItemList();
};

// 添加内容
const handleAdd = () => {
  editForm.value = {
    module: module.value,
    is_enabled: true,
    sort_order: 1,
  };
  editVisible.value = true;
};

// 编辑内容
const handleEdit = (record: DiscoverItem) => {
  editForm.value = { ...record };
  editVisible.value = true;
};

// 编辑确认
const handleEditOk = async () => {
  try {
    if (editForm.value.id) {
      await updateDiscoverItem(editForm.value);
      message.success('更新成功');
    } else {
      await createDiscoverItem(editForm.value);
      message.success('添加成功');
    }
    editVisible.value = false;
    loadItemList();
  } catch (error) {
    message.error('操作失败');
  }
};

// 删除内容
const handleDelete = async (id: number) => {
  try {
    await deleteDiscoverItem(id);
    message.success('删除成功');
    loadItemList();
  } catch (error) {
    message.error('删除失败');
  }
};

onMounted(() => {
  loadItemList();
});
</script>

<style scoped>
.discover-items {
  padding: 24px;
}
</style>
