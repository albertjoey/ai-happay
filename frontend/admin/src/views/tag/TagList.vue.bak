<template>
  <div class="tag-list">
    <a-card title="标签管理">
      <template #extra>
        <a-space>
          <a-input-search
            v-model:value="searchName"
            placeholder="搜索标签名称"
            style="width: 200px"
            @search="loadTags"
          />
          <a-select v-model:value="searchType" style="width: 120px" placeholder="类型" allowClear @change="loadTags">
            <a-select-option :value="0">通用</a-select-option>
            <a-select-option :value="1">长视频</a-select-option>
            <a-select-option :value="2">短视频</a-select-option>
            <a-select-option :value="3">短剧</a-select-option>
            <a-select-option :value="4">漫剧</a-select-option>
            <a-select-option :value="5">小说</a-select-option>
            <a-select-option :value="6">图文</a-select-option>
          </a-select>
          <a-button type="primary" @click="handleAdd">
            <template #icon><PlusOutlined /></template>
            添加标签
          </a-button>
        </a-space>
      </template>

      <vxe-table
        border
        stripe
        :data="tableData"
        :loading="loading"
      >
        <vxe-column type="seq" width="60" title="序号"></vxe-column>
        <vxe-column field="id" title="ID" width="80"></vxe-column>
        <vxe-column field="name" title="标签名称" width="150"></vxe-column>
        <vxe-column field="type" title="类型" width="120">
          <template #default="{ row }">
            <a-tag :color="getTypeColor(row.type)">{{ getTypeName(row.type) }}</a-tag>
          </template>
        </vxe-column>
        <vxe-column field="status" title="状态" width="100">
          <template #default="{ row }">
            <a-tag v-if="row.status === 1" color="success">启用</a-tag>
            <a-tag v-else color="error">禁用</a-tag>
          </template>
        </vxe-column>
        <vxe-column field="created_at" title="创建时间" width="180"></vxe-column>
        <vxe-column title="操作" width="150" fixed="right">
          <template #default="{ row }">
            <a-space>
              <a-button type="link" size="small" @click="handleEdit(row)">编辑</a-button>
              <a-popconfirm
                title="确定要删除此标签吗?"
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

    <!-- 添加/编辑对话框 -->
    <a-modal
      v-model:open="modalVisible"
      :title="modalTitle"
      @ok="handleModalOk"
      @cancel="handleModalCancel"
      width="500px"
    >
      <a-form
        ref="formRef"
        :model="formState"
        :rules="formRules"
        :label-col="{ span: 6 }"
        :wrapper-col="{ span: 16 }"
      >
        <a-form-item label="标签名称" name="name">
          <a-input v-model:value="formState.name" placeholder="请输入标签名称" />
        </a-form-item>
        <a-form-item label="标签类型" name="type">
          <a-select v-model:value="formState.type">
            <a-select-option :value="0">通用</a-select-option>
            <a-select-option :value="1">长视频</a-select-option>
            <a-select-option :value="2">短视频</a-select-option>
            <a-select-option :value="3">短剧</a-select-option>
            <a-select-option :value="4">漫剧</a-select-option>
            <a-select-option :value="5">小说</a-select-option>
            <a-select-option :value="6">图文</a-select-option>
          </a-select>
        </a-form-item>
        <a-form-item label="状态" name="status">
          <a-radio-group v-model:value="formState.status">
            <a-radio :value="1">启用</a-radio>
            <a-radio :value="0">禁用</a-radio>
          </a-radio-group>
        </a-form-item>
      </a-form>
    </a-modal>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted } from 'vue';
import { message } from 'ant-design-vue';
import { PlusOutlined } from '@ant-design/icons-vue';
import type { FormInstance } from 'ant-design-vue';
import {
  getTagList,
  createTag,
  updateTag,
  deleteTag,
  type Tag,
} from '@/api/tag';

const loading = ref(false);
const tableData = ref<Tag[]>([]);
const searchName = ref('');
const searchType = ref<number | undefined>(undefined);
const modalVisible = ref(false);
const modalTitle = ref('添加标签');
const currentId = ref<number>(0);
const formRef = ref<FormInstance>();

const pagination = reactive({
  current: 1,
  pageSize: 10,
  total: 0,
});

const formState = reactive({
  name: '',
  type: 0,
  status: 1,
});

const formRules = {
  name: [{ required: true, message: '请输入标签名称', trigger: 'blur' }],
};

// 获取类型名称
const getTypeName = (type: number) => {
  const types: Record<number, string> = {
    0: '通用',
    1: '长视频',
    2: '短视频',
    3: '短剧',
    4: '漫剧',
    5: '小说',
    6: '图文',
  };
  return types[type] || '未知';
};

// 获取类型颜色
const getTypeColor = (type: number) => {
  const colors: Record<number, string> = {
    0: 'default',
    1: 'blue',
    2: 'cyan',
    3: 'purple',
    4: 'magenta',
    5: 'orange',
    6: 'green',
  };
  return colors[type] || 'default';
};

// 加载标签列表
const loadTags = async () => {
  loading.value = true;
  try {
    const res = await getTagList({
      page: pagination.current,
      page_size: pagination.pageSize,
      name: searchName.value,
      type: searchType.value,
    });
    tableData.value = res.list || [];
    pagination.total = res.total || 0;
  } catch (error) {
    message.error('加载标签列表失败');
  } finally {
    loading.value = false;
  }
};

// 添加标签
const handleAdd = () => {
  modalTitle.value = '添加标签';
  currentId.value = 0;
  Object.assign(formState, {
    name: '',
    type: 0,
    status: 1,
  });
  modalVisible.value = true;
};

// 编辑标签
const handleEdit = (row: Tag) => {
  modalTitle.value = '编辑标签';
  currentId.value = row.id;
  Object.assign(formState, {
    name: row.name,
    type: row.type,
    status: row.status,
  });
  modalVisible.value = true;
};

// 删除标签
const handleDelete = async (row: Tag) => {
  try {
    await deleteTag(row.id);
    message.success('删除成功');
    await loadTags();
  } catch (error) {
    message.error('删除失败');
  }
};

// 提交表单
const handleModalOk = async () => {
  try {
    await formRef.value?.validate();
    if (currentId.value) {
      await updateTag(currentId.value, formState);
      message.success('更新成功');
    } else {
      await createTag(formState);
      message.success('创建成功');
    }
    modalVisible.value = false;
    await loadTags();
  } catch (error) {
    message.error('操作失败');
  }
};

// 取消表单
const handleModalCancel = () => {
  modalVisible.value = false;
  formRef.value?.resetFields();
};

// 分页改变
const handlePageChange = (page: number, pageSize: number) => {
  pagination.current = page;
  pagination.pageSize = pageSize;
  loadTags();
};

onMounted(() => {
  loadTags();
});
</script>

<style scoped>
.tag-list {
  padding: 0;
}

.pagination {
  margin-top: 16px;
  text-align: right;
}
</style>
