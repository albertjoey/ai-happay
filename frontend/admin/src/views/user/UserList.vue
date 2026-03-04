<template>
  <div class="user-list">
    <a-card title="用户管理">
      <template #extra>
        <a-button type="primary" @click="handleAdd">
          <template #icon><PlusOutlined /></template>
          添加用户
        </a-button>
      </template>

      <vxe-table
        border
        stripe
        :data="tableData"
        :loading="loading"
      >
        <vxe-column type="seq" width="60" title="序号"></vxe-column>
        <vxe-column field="id" title="ID" width="80"></vxe-column>
        <vxe-column field="username" title="用户名" width="120"></vxe-column>
        <vxe-column field="nickname" title="昵称" width="120"></vxe-column>
        <vxe-column field="email" title="邮箱" width="200"></vxe-column>
        <vxe-column field="phone" title="手机号" width="150"></vxe-column>
        <vxe-column field="role" title="角色" width="100">
          <template #default="{ row }">
            <a-tag v-if="row.role === 0" color="blue">普通用户</a-tag>
            <a-tag v-else-if="row.role === 1" color="green">博主</a-tag>
            <a-tag v-else-if="row.role === 2" color="red">管理员</a-tag>
          </template>
        </vxe-column>
        <vxe-column field="status" title="状态" width="100">
          <template #default="{ row }">
            <a-tag v-if="row.status === 1" color="success">启用</a-tag>
            <a-tag v-else color="error">禁用</a-tag>
          </template>
        </vxe-column>
        <vxe-column field="created_at" title="创建时间" width="180"></vxe-column>
        <vxe-column title="操作" width="200" fixed="right">
          <template #default="{ row }">
            <a-space>
              <a-button type="link" size="small" @click="handleEdit(row)">编辑</a-button>
              <a-button type="link" size="small" @click="handleView(row)">查看</a-button>
              <a-popconfirm
                title="确定要删除此用户吗？"
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

    <!-- 添加/编辑用户对话框 -->
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
        <a-form-item label="用户名" name="username">
          <a-input v-model:value="formState.username" placeholder="请输入用户名" :disabled="isView" />
        </a-form-item>
        <a-form-item label="昵称" name="nickname">
          <a-input v-model:value="formState.nickname" placeholder="请输入昵称" :disabled="isView" />
        </a-form-item>
        <a-form-item label="邮箱" name="email">
          <a-input v-model:value="formState.email" placeholder="请输入邮箱" :disabled="isView" />
        </a-form-item>
        <a-form-item label="手机号" name="phone">
          <a-input v-model:value="formState.phone" placeholder="请输入手机号" :disabled="isView" />
        </a-form-item>
        <a-form-item v-if="!isEdit" label="密码" name="password">
          <a-input-password v-model:value="formState.password" placeholder="请输入密码" />
        </a-form-item>
        <a-form-item label="角色" name="role">
          <a-select v-model:value="formState.role" :disabled="isView">
            <a-select-option :value="0">普通用户</a-select-option>
            <a-select-option :value="1">博主</a-select-option>
            <a-select-option :value="2">管理员</a-select-option>
          </a-select>
        </a-form-item>
        <a-form-item label="状态" name="status">
          <a-radio-group v-model:value="formState.status" :disabled="isView">
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

const loading = ref(false);
const tableData = ref([
  {
    id: 1,
    username: 'user001',
    nickname: '用户001',
    email: 'user001@example.com',
    phone: '13800138001',
    role: 0,
    status: 1,
    created_at: '2024-01-01 10:00:00',
  },
  {
    id: 2,
    username: 'blogger001',
    nickname: '博主001',
    email: 'blogger001@example.com',
    phone: '13800138002',
    role: 1,
    status: 1,
    created_at: '2024-01-02 11:00:00',
  },
  {
    id: 3,
    username: 'admin',
    nickname: '管理员',
    email: 'admin@example.com',
    phone: '13800138000',
    role: 2,
    status: 1,
    created_at: '2024-01-01 00:00:00',
  },
]);

const pagination = reactive({
  current: 1,
  pageSize: 10,
  total: 100,
});

const modalVisible = ref(false);
const modalTitle = ref('添加用户');
const isEdit = ref(false);
const isView = ref(false);
const currentId = ref<number>(0);
const formRef = ref<FormInstance>();

const formState = reactive({
  username: '',
  nickname: '',
  email: '',
  phone: '',
  password: '',
  role: 0,
  status: 1,
});

const formRules = {
  username: [{ required: true, message: '请输入用户名', trigger: 'blur' }],
  nickname: [{ required: true, message: '请输入昵称', trigger: 'blur' }],
  email: [
    { required: true, message: '请输入邮箱', trigger: 'blur' },
    { type: 'email', message: '请输入正确的邮箱格式', trigger: 'blur' },
  ],
  password: [{ required: true, message: '请输入密码', trigger: 'blur' }],
};

const handleAdd = () => {
  modalTitle.value = '添加用户';
  isEdit.value = false;
  isView.value = false;
  currentId.value = 0;
  Object.assign(formState, {
    username: '',
    nickname: '',
    email: '',
    phone: '',
    password: '',
    role: 0,
    status: 1,
  });
  modalVisible.value = true;
};

const handleEdit = (row: any) => {
  modalTitle.value = '编辑用户';
  isEdit.value = true;
  isView.value = false;
  currentId.value = row.id;
  Object.assign(formState, {
    username: row.username,
    nickname: row.nickname,
    email: row.email,
    phone: row.phone,
    role: row.role,
    status: row.status,
  });
  modalVisible.value = true;
};

const handleView = (row: any) => {
  modalTitle.value = '查看用户';
  isEdit.value = false;
  isView.value = true;
  currentId.value = row.id;
  Object.assign(formState, {
    username: row.username,
    nickname: row.nickname,
    email: row.email,
    phone: row.phone,
    role: row.role,
    status: row.status,
  });
  modalVisible.value = true;
};

const handleDelete = (row: any) => {
  message.success(`删除用户: ${row.username}`);
  // 这里应该调用API删除用户
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
  // 重新加载数据
};

onMounted(() => {
  // 加载数据
});
</script>

<style scoped>
.user-list {
  padding: 0;
}

.pagination {
  margin-top: 16px;
  text-align: right;
}
</style>
