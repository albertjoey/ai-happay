<template>
  <div class="topic-list">
    <a-card title="话题管理">
      <template #extra>
        <a-space>
          <a-input-search
            v-model:value="searchName"
            placeholder="搜索话题名称"
            style="width: 200px"
            @search="loadTopics"
          />
          <a-button type="primary" @click="handleAdd">
            <template #icon><PlusOutlined /></template>
            添加话题
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
        <vxe-column field="cover" title="封面" width="100">
          <template #default="{ row }">
            <img v-if="row.cover" :src="row.cover" class="cover-img" />
            <span v-else class="text-gray">无封面</span>
          </template>
        </vxe-column>
        <vxe-column field="name" title="话题名称" width="150"></vxe-column>
        <vxe-column field="description" title="描述" min-width="200"></vxe-column>
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
                title="确定要删除此话题吗?"
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
      width="600px"
    >
      <a-form
        ref="formRef"
        :model="formState"
        :rules="formRules"
        :label-col="{ span: 6 }"
        :wrapper-col="{ span: 16 }"
      >
        <a-form-item label="话题名称" name="name">
          <a-input v-model:value="formState.name" placeholder="请输入话题名称" />
        </a-form-item>
        <a-form-item label="话题描述" name="description">
          <a-textarea v-model:value="formState.description" placeholder="请输入话题描述" :rows="3" />
        </a-form-item>
        <a-form-item label="封面图片" name="cover">
          <a-input v-model:value="formState.cover" placeholder="请输入封面图片URL" />
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
  getTopicList,
  createTopic,
  updateTopic,
  deleteTopic,
  type Topic,
} from '@/api/topic';

const loading = ref(false);
const tableData = ref<Topic[]>([]);
const searchName = ref('');
const modalVisible = ref(false);
const modalTitle = ref('添加话题');
const currentId = ref<number>(0);
const formRef = ref<FormInstance>();

const pagination = reactive({
  current: 1,
  pageSize: 10,
  total: 0,
});

const formState = reactive({
  name: '',
  description: '',
  cover: '',
  status: 1,
});

const formRules = {
  name: [{ required: true, message: '请输入话题名称', trigger: 'blur' }],
};

// 加载话题列表
const loadTopics = async () => {
  loading.value = true;
  try {
    const res = await getTopicList({
      page: pagination.current,
      page_size: pagination.pageSize,
      name: searchName.value,
    });
    tableData.value = res.list || [];
    pagination.total = res.total || 0;
  } catch (error) {
    message.error('加载话题列表失败');
  } finally {
    loading.value = false;
  }
};

// 添加话题
const handleAdd = () => {
  modalTitle.value = '添加话题';
  currentId.value = 0;
  Object.assign(formState, {
    name: '',
    description: '',
    cover: '',
    status: 1,
  });
  modalVisible.value = true;
};

// 编辑话题
const handleEdit = (row: Topic) => {
  modalTitle.value = '编辑话题';
  currentId.value = row.id;
  Object.assign(formState, {
    name: row.name,
    description: row.description,
    cover: row.cover,
    status: row.status,
  });
  modalVisible.value = true;
};

// 删除话题
const handleDelete = async (row: Topic) => {
  try {
    await deleteTopic(row.id);
    message.success('删除成功');
    await loadTopics();
  } catch (error) {
    message.error('删除失败');
  }
};

// 提交表单
const handleModalOk = async () => {
  try {
    await formRef.value?.validate();
    if (currentId.value) {
      await updateTopic(currentId.value, formState);
      message.success('更新成功');
    } else {
      await createTopic(formState);
      message.success('创建成功');
    }
    modalVisible.value = false;
    await loadTopics();
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
  loadTopics();
};

onMounted(() => {
  loadTopics();
});
</script>

<style scoped>
.topic-list {
  padding: 0;
}

.pagination {
  margin-top: 16px;
  text-align: right;
}

.cover-img {
  width: 60px;
  height: 60px;
  object-fit: cover;
  border-radius: 4px;
}

.text-gray {
  color: #999;
}
</style>
