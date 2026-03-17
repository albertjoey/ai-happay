<template>
  <div class="content-list">
    <a-card title="内容管理">
      <template #extra>
        <a-button type="primary" @click="handleAdd">
          <template #icon><PlusOutlined /></template>
          添加内容
        </a-button>
      </template>

      <a-table
        :columns="columns"
        :data-source="tableData"
        :loading="loading"
        :pagination="{ pageSize: 10 }"
        row-key="id"
      >
        <template #bodyCell="{ column, record }">
        
        
        
        
        
        
        
        
        
        
      </template>
      </a-table>

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

    <!-- 添加/编辑内容对话框 -->
    <a-modal
      v-model:open="modalVisible"
      :title="modalTitle"
      @ok="handleModalOk"
      @cancel="handleModalCancel"
      width="800px"
    >
      <a-form
        ref="formRef"
        :model="formState"
        :rules="formRules"
        :label-col="{ span: 4 }"
        :wrapper-col="{ span: 18 }"
      >
        <a-form-item label="标题" name="title">
          <a-input v-model:value="formState.title" placeholder="请输入标题" :disabled="isView" />
        </a-form-item>
        <a-form-item label="类型" name="type">
          <a-select v-model:value="formState.type" :disabled="isView">
            <a-select-option :value="1">长视频</a-select-option>
            <a-select-option :value="2">短视频</a-select-option>
            <a-select-option :value="3">短剧</a-select-option>
            <a-select-option :value="4">漫剧</a-select-option>
            <a-select-option :value="5">小说</a-select-option>
            <a-select-option :value="6">图文</a-select-option>
          </a-select>
        </a-form-item>
        <a-form-item label="作者" name="author">
          <a-input v-model:value="formState.author" placeholder="请输入作者" :disabled="isView" />
        </a-form-item>
        <a-form-item label="封面" name="cover">
          <a-input v-model:value="formState.cover" placeholder="请输入封面URL" :disabled="isView" />
        </a-form-item>
        <a-form-item label="内容" name="content">
          <a-textarea v-model:value="formState.content" placeholder="请输入内容" :rows="5" :disabled="isView" />
        </a-form-item>
        <a-form-item label="状态" name="status">
          <a-radio-group v-model:value="formState.status" :disabled="isView">
            <a-radio :value="0">草稿</a-radio>
            <a-radio :value="1">已发布</a-radio>
            <a-radio :value="2">已下架</a-radio>
            <a-radio :value="3">审核中</a-radio>
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

const loading = ref(false);
const tableData = ref([
  {
    id: 1,
    title: '精彩短视频分享',
    type: 2,
    author: '创作者A',
    view_count: 1234,
    like_count: 567,
    status: 1,
    created_at: '2024-01-01 10:00:00',
  },
  {
    id: 2,
    title: '热门短剧推荐',
    type: 3,
    author: '创作者B',
    view_count: 2345,
    like_count: 678,
    status: 1,
    created_at: '2024-01-02 11:00:00',
  },
  {
    id: 3,
    title: '漫剧精彩片段',
    type: 4,
    author: '创作者C',
    view_count: 3456,
    like_count: 789,
    status: 3,
    created_at: '2024-01-03 12:00:00',
  },
]);

const pagination = reactive({
  current: 1,
  pageSize: 10,
  total: 100,
});

const modalVisible = ref(false);
const modalTitle = ref('添加内容');
const isEdit = ref(false);
const isView = ref(false);
const currentId = ref<number>(0);
const formRef = ref<FormInstance>();

// 表格列定义
const columns = [
  { title: 'ID', dataIndex: 'id', key: 'id', width: 80 },
  { title: '标题', dataIndex: 'title', key: 'title', width: 200 },
  { title: '类型', dataIndex: 'type', key: 'type', width: 100 },
  { title: '状态', dataIndex: 'status', key: 'status', width: 100 },
  { title: '创建时间', dataIndex: 'created_at', key: 'created_at', width: 180 },
  { title: '操作', key: 'action', width: 150, fixed: 'right' as const },
];

const formState = reactive({
  title: '',
  type: 2,
  author: '',
  cover: '',
  content: '',
  status: 0,
});

const formRules = {
  title: [{ required: true, message: '请输入标题', trigger: 'blur' }],
  type: [{ required: true, message: '请选择类型', trigger: 'change' }],
  author: [{ required: true, message: '请输入作者', trigger: 'blur' }],
};

const handleAdd = () => {
  modalTitle.value = '添加内容';
  isEdit.value = false;
  isView.value = false;
  currentId.value = 0;
  Object.assign(formState, {
    title: '',
    type: 2,
    author: '',
    cover: '',
    content: '',
    status: 0,
  });
  modalVisible.value = true;
};

const handleEdit = (row: any) => {
  modalTitle.value = '编辑内容';
  isEdit.value = true;
  isView.value = false;
  currentId.value = row.id;
  Object.assign(formState, {
    title: row.title,
    type: row.type,
    author: row.author,
    status: row.status,
  });
  modalVisible.value = true;
};

const handleView = (row: any) => {
  modalTitle.value = '查看内容';
  isEdit.value = false;
  isView.value = true;
  currentId.value = row.id;
  Object.assign(formState, {
    title: row.title,
    type: row.type,
    author: row.author,
    status: row.status,
  });
  modalVisible.value = true;
};

const handleAudit = (row: any) => {
  message.info(`审核内容: ${row.title}`);
};

const handleDelete = (row: any) => {
  message.success(`删除内容: ${row.title}`);
};

const handleModalOk = async () => {
  if (isView.value) {
    modalVisible.value = false;
    return;
  }

  try {
    await formRef.value?.validate();
    if (isEdit.value) {
      message.success('更新成功');
    } else {
      message.success('创建成功');
    }
    modalVisible.value = false;
  } catch (error) {
    message.error('操作失败');
  }
};

const handleModalCancel = () => {
  modalVisible.value = false;
  formRef.value?.resetFields();
};

const handlePageChange = (page: number, pageSize: number) => {
  pagination.current = page;
  pagination.pageSize = pageSize;
};

onMounted(() => {
  // 加载数据
});
</script>

<style scoped>
.content-list {
  padding: 0;
}

.pagination {
  margin-top: 16px;
  text-align: right;
}
</style>
