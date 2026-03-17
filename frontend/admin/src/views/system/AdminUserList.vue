<template>
  <div class="admin-user-list">
    <a-card title="管理员管理">
      <template #extra>
        <a-space>
          <a-input-search
            v-model:value="searchKeyword"
            placeholder="搜索用户名/姓名"
            style="width: 200px"
            @search="handleSearch"
          />
          <a-button type="primary" @click="handleAdd">
            <template #icon><PlusOutlined /></template>
            添加管理员
          </a-button>
        </a-space>
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

    <!-- 添加/编辑管理员对话框 -->
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
          <a-input v-model:value="formState.username" placeholder="请输入用户名" :disabled="isEdit" />
        </a-form-item>
        <a-form-item v-if="!isEdit" label="密码" name="password">
          <a-input-password v-model:value="formState.password" placeholder="请输入密码" />
        </a-form-item>
        <a-form-item label="姓名" name="realname">
          <a-input v-model:value="formState.realname" placeholder="请输入姓名" />
        </a-form-item>
        <a-form-item label="邮箱" name="email">
          <a-input v-model:value="formState.email" placeholder="请输入邮箱" />
        </a-form-item>
        <a-form-item label="手机号" name="phone">
          <a-input v-model:value="formState.phone" placeholder="请输入手机号" />
        </a-form-item>
        <a-form-item label="状态" name="status">
          <a-radio-group v-model:value="formState.status">
            <a-radio :value="1">启用</a-radio>
            <a-radio :value="0">禁用</a-radio>
          </a-radio-group>
        </a-form-item>
      </a-form>
    </a-modal>

    <!-- 角色配置对话框 -->
    <a-modal
      v-model:open="roleModalVisible"
      title="角色配置"
      @ok="handleRoleOk"
      width="500px"
    >
      <a-checkbox-group v-model:value="checkedRoleIds" style="width: 100%">
        <div v-for="role in roleList" :key="role.id" style="margin-bottom: 12px">
          <a-checkbox :value="role.id">{{ role.name }} ({{ role.code }})</a-checkbox>
        </div>
      </a-checkbox-group>
    </a-modal>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted } from 'vue';
import { message } from 'ant-design-vue';
import { PlusOutlined } from '@ant-design/icons-vue';
import type { FormInstance } from 'ant-design-vue';
import {
  getAdminUserList,
  createAdminUser,
  updateAdminUser,
  deleteAdminUser,
  getRoleList,
  assignAdminRoles,
  getAdminRoles,
  type AdminUser,
  type Role,
} from '@/api/rbac';

const loading = ref(false);
const tableData = ref<AdminUser[]>([]);
const searchKeyword = ref('');

const pagination = reactive({
  current: 1,
  pageSize: 20,
  total: 0,
});

const modalVisible = ref(false);
const modalTitle = ref('添加管理员');
const isEdit = ref(false);
const currentId = ref<number>(0);
const formRef = ref<FormInstance>();

// 表格列定义
const columns = [
  { title: 'ID', dataIndex: 'id', key: 'id', width: 80 },
  { title: '用户名', dataIndex: 'username', key: 'username', width: 120 },
  { title: '姓名', dataIndex: 'realname', key: 'realname', width: 120 },
  { title: '邮箱', dataIndex: 'email', key: 'email', width: 180 },
  { title: '手机', dataIndex: 'phone', key: 'phone', width: 120 },
  { title: '状态', dataIndex: 'status', key: 'status', width: 100 },
  { title: '创建时间', dataIndex: 'created_at', key: 'created_at', width: 180 },
  { title: '操作', key: 'action', width: 200, fixed: 'right' as const },
];

const formState = reactive({
  username: '',
  password: '',
  realname: '',
  email: '',
  phone: '',
  status: 1,
});

const formRules = {
  username: [{ required: true, message: '请输入用户名', trigger: 'blur' }],
  password: [{ required: true, message: '请输入密码', trigger: 'blur' }],
  realname: [{ required: true, message: '请输入姓名', trigger: 'blur' }],
};

// 角色配置
const roleModalVisible = ref(false);
const roleList = ref<Role[]>([]);
const checkedRoleIds = ref<number[]>([]);
const currentAdminId = ref<number>(0);

const loadData = async () => {
  loading.value = true;
  try {
    const res = await getAdminUserList({
      page: pagination.current,
      page_size: pagination.pageSize,
      keyword: searchKeyword.value,
    });
    tableData.value = res.list || [];
    pagination.total = res.total;
  } catch (error) {
    message.error('加载数据失败');
  } finally {
    loading.value = false;
  }
};

const handleSearch = () => {
  pagination.current = 1;
  loadData();
};

const handlePageChange = (page: number, pageSize: number) => {
  pagination.current = page;
  pagination.pageSize = pageSize;
  loadData();
};

const handleAdd = () => {
  modalTitle.value = '添加管理员';
  isEdit.value = false;
  currentId.value = 0;
  Object.assign(formState, {
    username: '',
    password: '',
    realname: '',
    email: '',
    phone: '',
    status: 1,
  });
  modalVisible.value = true;
};

const handleEdit = (row: AdminUser) => {
  modalTitle.value = '编辑管理员';
  isEdit.value = true;
  currentId.value = row.id;
  Object.assign(formState, {
    username: row.username,
    realname: row.realname,
    email: row.email,
    phone: row.phone,
    status: row.status,
  });
  modalVisible.value = true;
};

const handleDelete = async (row: AdminUser) => {
  try {
    await deleteAdminUser(row.id);
    message.success('删除成功');
    loadData();
  } catch (error) {
    message.error('删除失败');
  }
};

const handleModalOk = async () => {
  try {
    await formRef.value?.validate();
    if (isEdit.value) {
      await updateAdminUser(currentId.value, formState);
      message.success('更新成功');
    } else {
      await createAdminUser(formState as any);
      message.success('创建成功');
    }
    modalVisible.value = false;
    loadData();
  } catch (error) {
    message.error('操作失败');
  }
};

const handleModalCancel = () => {
  modalVisible.value = false;
  formRef.value?.resetFields();
};

// 角色配置
const handleRole = async (row: AdminUser) => {
  currentAdminId.value = row.id;
  try {
    // 加载角色列表
    const res = await getRoleList();
    roleList.value = res.list || [];
    
    // 加载管理员已有角色
    const roleIds = await getAdminRoles(row.id);
    checkedRoleIds.value = roleIds;
    
    roleModalVisible.value = true;
  } catch (error) {
    message.error('加载角色失败');
  }
};

const handleRoleOk = async () => {
  try {
    await assignAdminRoles(currentAdminId.value, checkedRoleIds.value);
    message.success('角色配置成功');
    roleModalVisible.value = false;
  } catch (error) {
    message.error('角色配置失败');
  }
};

onMounted(() => {
  loadData();
});
</script>

<style scoped>
.admin-user-list {
  padding: 0;
}

.pagination {
  margin-top: 16px;
  text-align: right;
}
</style>
