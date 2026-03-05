<template>
  <div class="channel-list">
    <a-card title="频道管理">
      <template #extra>
        <a-button type="primary" @click="handleAdd">
          <template #icon><PlusOutlined /></template>
          添加频道
        </a-button>
      </template>

      <vxe-table
        border
        stripe
        :data="tableData"
        :loading="loading"
        row-config="{keyField: 'id'}"
      >
        <vxe-column type="seq" width="60" title="序号"></vxe-column>
        <vxe-column field="id" title="ID" width="80"></vxe-column>
        <vxe-column field="name" title="频道名称" width="150"></vxe-column>
        <vxe-column field="icon" title="图标" width="100">
          <template #default="{ row }">
            <span class="icon-preview">{{ row.icon }}</span>
          </template>
        </vxe-column>
        <vxe-column field="description" title="描述" min-width="200"></vxe-column>
        <vxe-column field="sort" title="排序" width="80">
          <template #default="{ row }">
            <a-tag color="blue">{{ row.sort }}</a-tag>
          </template>
        </vxe-column>
        <vxe-column field="status" title="状态" width="100">
          <template #default="{ row }">
            <a-tag v-if="row.status === 1" color="success">启用</a-tag>
            <a-tag v-else color="error">禁用</a-tag>
          </template>
        </vxe-column>
        <vxe-column field="created_at" title="创建时间" width="180"></vxe-column>
        <vxe-column title="操作" width="280" fixed="right">
          <template #default="{ row }">
            <a-space>
              <a-button type="link" size="small" @click="handleEdit(row)">编辑</a-button>
              <a-button type="link" size="small" @click="handleConfig(row)">配置</a-button>
              <a-button type="link" size="small" @click="handleMoveUp(row)">上移</a-button>
              <a-button type="link" size="small" @click="handleMoveDown(row)">下移</a-button>
              <a-popconfirm
                title="确定要删除此频道吗?"
                @confirm="handleDelete(row)"
              >
                <a-button type="link" size="small" danger>删除</a-button>
              </a-popconfirm>
            </a-space>
          </template>
        </vxe-column>
      </vxe-table>

      <div class="pagination">
        <a-pagination
          v-model:current="pagination.current"
          v-model:pageSize="pagination.pageSize"
          :total="pagination.total"
          show-size-changer
          @change="handlePageChange"
        />
      </div>
    </a-card>

    <!-- 添加/编辑频道对话框 -->
    <a-modal
      v-model:open="modalVisible"
      :title="modalTitle"
      @ok="handleModalOk"
      @cancel="handleModalCancel"
      width="600px"
    >
      <a-form
        ref="formRef"
        :model="formState"
        :rules="formRules"
        :label-col="{ span: 6 }"
        :wrapper-col="{ span: 16 }"
      >
        <a-form-item label="频道名称" name="name">
          <a-input v-model:value="formState.name" placeholder="请输入频道名称" />
        </a-form-item>
        <a-form-item label="图标" name="icon">
          <a-input v-model:value="formState.icon" placeholder="请输入图标(emoji或图标名)" />
        </a-form-item>
        <a-form-item label="描述" name="description">
          <a-textarea v-model:value="formState.description" placeholder="请输入频道描述" :rows="3" />
        </a-form-item>
        <a-form-item label="排序" name="sort">
          <a-input-number v-model:value="formState.sort" :min="0" :max="999" />
        </a-form-item>
        <a-form-item label="状态" name="status">
          <a-radio-group v-model:value="formState.status">
            <a-radio :value="1">启用</a-radio>
            <a-radio :value="0">禁用</a-radio>
          </a-radio-group>
        </a-form-item>
      </a-form>
    </a-modal>

    <!-- 频道配置对话框 -->
    <a-modal
      v-model:open="configModalVisible"
      title="频道配置"
      @ok="handleConfigModalOk"
      @cancel="handleConfigModalCancel"
      width="700px"
    >
      <a-form
        :model="configFormState"
        :label-col="{ span: 6 }"
        :wrapper-col="{ span: 16 }"
      >
        <a-form-item label="内容类型">
          <a-checkbox-group v-model:value="contentTypes">
            <a-checkbox value="video">视频</a-checkbox>
            <a-checkbox value="article">文章</a-checkbox>
            <a-checkbox value="image">图片</a-checkbox>
          </a-checkbox-group>
        </a-form-item>
        <a-form-item label="展示模式">
          <a-select v-model:value="configFormState.display_mode" style="width: 200px">
            <a-select-option value="default">默认模式</a-select-option>
            <a-select-option value="waterfall">瀑布流</a-select-option>
            <a-select-option value="list">列表模式</a-select-option>
            <a-select-option value="grid">网格模式</a-select-option>
          </a-select>
        </a-form-item>
        <a-form-item label="显示作者">
          <a-switch v-model:checked="configFormState.custom_data.show_author" />
        </a-form-item>
        <a-form-item label="显示统计">
          <a-switch v-model:checked="configFormState.custom_data.show_stats" />
        </a-form-item>
        <a-form-item label="自动播放">
          <a-switch v-model:checked="configFormState.custom_data.auto_play" />
        </a-form-item>
      </a-form>
    </a-modal>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted, computed } from 'vue';
import { useRouter } from 'vue-router';
import { message } from 'ant-design-vue';
import { PlusOutlined } from '@ant-design/icons-vue';
import type { FormInstance } from 'ant-design-vue';
import {
  getChannelList,
  createChannel,
  updateChannel,
  deleteChannel,
  sortChannels,
  getChannelConfig,
  updateChannelConfig,
  type Channel,
  type ChannelConfig,
} from '@/api/channel';

const loading = ref(false);
const tableData = ref<Channel[]>([]);
const modalVisible = ref(false);
const configModalVisible = ref(false);
const modalTitle = ref('添加频道');
const currentChannelId = ref<number>(0);
const formRef = ref<FormInstance>();
const router = useRouter();

const pagination = reactive({
  current: 1,
  pageSize: 10,
  total: 0,
});

const formState = reactive({
  name: '',
  icon: '',
  description: '',
  sort: 0,
  status: 1,
});

const formRules = {
  name: [{ required: true, message: '请输入频道名称', trigger: 'blur' }],
  icon: [{ required: true, message: '请输入图标', trigger: 'blur' }],
};

const configFormState = reactive<ChannelConfig>({
  channel_id: 0,
  content_type: {
    video: true,
    article: true,
    image: true,
  },
  display_mode: 'default',
  custom_data: {
    show_author: true,
    show_stats: true,
    auto_play: false,
  },
});

const contentTypes = computed({
  get: () => {
    const types: string[] = [];
    if (configFormState.content_type.video) types.push('video');
    if (configFormState.content_type.article) types.push('article');
    if (configFormState.content_type.image) types.push('image');
    return types;
  },
  set: (values: string[]) => {
    configFormState.content_type = {
      video: values.includes('video'),
      article: values.includes('article'),
      image: values.includes('image'),
    };
  },
});

// 加载频道列表
const loadChannels = async () => {
  loading.value = true;
  try {
    const res = await getChannelList({
      page: pagination.current,
      page_size: pagination.pageSize,
    });
    tableData.value = res.list || [];
    pagination.total = res.total || 0;
  } catch (error) {
    message.error('加载频道列表失败');
  } finally {
    loading.value = false;
  }
};

// 添加频道
const handleAdd = () => {
  modalTitle.value = '添加频道';
  currentChannelId.value = 0;
  Object.assign(formState, {
    name: '',
    icon: '',
    description: '',
    sort: 0,
    status: 1,
  });
  modalVisible.value = true;
};

// 编辑频道
const handleEdit = (row: Channel) => {
  modalTitle.value = '编辑频道';
  currentChannelId.value = row.id;
  Object.assign(formState, {
    name: row.name,
    icon: row.icon,
    description: row.description,
    sort: row.sort,
    status: row.status,
  });
  modalVisible.value = true;
};

// 配置频道 - 跳转到频道配置页面
const handleConfig = (row: Channel) => {
  router.push({ path: '/channel/config', query: { channel_id: row.id } });
};

// 上移频道
const handleMoveUp = async (row: Channel) => {
  const index = tableData.value.findIndex(item => item.id === row.id);
  if (index > 0) {
    const items = [
      { id: row.id, sort: tableData.value[index - 1].sort },
      { id: tableData.value[index - 1].id, sort: row.sort },
    ];
    try {
      await sortChannels(items);
      message.success('排序成功');
      await loadChannels();
    } catch (error) {
      message.error('排序失败');
    }
  }
};

// 下移频道
const handleMoveDown = async (row: Channel) => {
  const index = tableData.value.findIndex(item => item.id === row.id);
  if (index < tableData.value.length - 1) {
    const items = [
      { id: row.id, sort: tableData.value[index + 1].sort },
      { id: tableData.value[index + 1].id, sort: row.sort },
    ];
    try {
      await sortChannels(items);
      message.success('排序成功');
      await loadChannels();
    } catch (error) {
      message.error('排序失败');
    }
  }
};

// 删除频道
const handleDelete = async (row: Channel) => {
  try {
    await deleteChannel(row.id);
    message.success('删除成功');
    await loadChannels();
  } catch (error) {
    message.error('删除失败');
  }
};

// 提交表单
const handleModalOk = async () => {
  try {
    await formRef.value?.validate();
    if (currentChannelId.value) {
      await updateChannel(currentChannelId.value, formState);
      message.success('更新成功');
    } else {
      await createChannel(formState);
      message.success('创建成功');
    }
    modalVisible.value = false;
    await loadChannels();
  } catch (error) {
    message.error('操作失败');
  }
};

// 取消表单
const handleModalCancel = () => {
  modalVisible.value = false;
  formRef.value?.resetFields();
};

// 提交配置
const handleConfigModalOk = async () => {
  try {
    await updateChannelConfig(currentChannelId.value, configFormState);
    message.success('配置保存成功');
    configModalVisible.value = false;
  } catch (error) {
    message.error('配置保存失败');
  }
};

// 取消配置
const handleConfigModalCancel = () => {
  configModalVisible.value = false;
};

// 分页改变
const handlePageChange = (page: number, pageSize: number) => {
  pagination.current = page;
  pagination.pageSize = pageSize;
  loadChannels();
};

onMounted(() => {
  loadChannels();
});
</script>

<style scoped>
.channel-list {
  padding: 0;
}

.pagination {
  margin-top: 16px;
  text-align: right;
}

.icon-preview {
  font-size: 20px;
}
</style>
