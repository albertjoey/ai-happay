<template>
  <div class="role-list">
    <a-card title="角色管理">
      <template #extra>
        <a-button type="primary" @click="handleAdd">
          <template #icon><PlusOutlined /></template>
          添加角色
        </a-button>
      </template>

      <vxe-table border stripe :data="tableData" :loading="loading">
        <vxe-column type="seq" width="60" title="序号"></vxe-column>
        <vxe-column field="id" title="ID" width="80"></vxe-column>
        <vxe-column field="name" title="角色名称" width="150"></vxe-column>
        <vxe-column field="code" title="角色编码" width="150"></vxe-column>
        <vxe-column field="description" title="描述" min-width="200"></vxe-column>
        <vxe-column field="status" title="状态" width="100">
          <template #default="{ row }">
            <a-tag v-if="row.status === 1" color="success">启用</a-tag>
            <a-tag v-else color="error">禁用</a-tag>
          </template>
        </vxe-column>
        <vxe-column field="created_at" title="创建时间" width="180"></vxe-column>
        <vxe-column title="操作" width="250" fixed="right">
          <template #default="{ row }">
            <a-space>
              <a-button type="link" size="small" @click="handleEdit(row)">编辑</a-button>
              <a-button type="link" size="small" @click="handlePermission(row)">权限配置</a-button>
              <a-popconfirm title="确定要删除此角色吗？" @confirm="handleDelete(row)">
                <a-button type="link" size="small" danger>删除</a-button>
              </a-popconfirm>
            </a-space>
          </template>
        </vxe-column>
      </vxe-table>
    </a-card>

    <!-- 添加/编辑角色对话框 -->
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
        <a-form-item label="角色名称" name="name">
          <a-input v-model:value="formState.name" placeholder="请输入角色名称" />
        </a-form-item>
        <a-form-item label="角色编码" name="code">
          <a-input v-model:value="formState.code" placeholder="请输入角色编码" :disabled="isEdit" />
        </a-form-item>
        <a-form-item label="描述" name="description">
          <a-textarea v-model:value="formState.description" placeholder="请输入描述" :rows="3" />
        </a-form-item>
        <a-form-item label="状态" name="status">
          <a-radio-group v-model:value="formState.status">
            <a-radio :value="1">启用</a-radio>
            <a-radio :value="0">禁用</a-radio>
          </a-radio-group>
        </a-form-item>
      </a-form>
    </a-modal>

    <!-- 权限配置对话框 -->
    <a-modal
      v-model:open="permissionModalVisible"
      title="权限配置"
      @ok="handlePermissionOk"
      width="600px"
    >
      <a-tree
        v-model:checkedKeys="checkedPermissionIds"
        :tree-data="permissionTree"
        checkable
        :field-names="{ title: 'name', key: 'id', children: 'children' }"
        :defaultExpandAll="true"
      />
    </a-modal>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted } from 'vue';
import { message } from 'ant-design-vue';
import { PlusOutlined } from '@ant-design/icons-vue';
import type { FormInstance } from 'ant-design-vue';
import {
  getRoleList,
  createRole,
  updateRole,
  deleteRole,
  getPermissionTree,
  assignRolePermissions,
  getRolePermissions,
  type Role,
  type Permission,
} from '@/api/rbac';

const loading = ref(false);
const tableData = ref<Role[]>([]);
const modalVisible = ref(false);
const modalTitle = ref('添加角色');
const isEdit = ref(false);
const currentId = ref<number>(0);
const formRef = ref<FormInstance>();

const formState = reactive({
  name: '',
  code: '',
  description: '',
  status: 1,
});

const formRules = {
  name: [{ required: true, message: '请输入角色名称', trigger: 'blur' }],
  code: [{ required: true, message: '请输入角色编码', trigger: 'blur' }],
};

// 权限配置
const permissionModalVisible = ref(false);
const permissionTree = ref<Permission[]>([]);
const checkedPermissionIds = ref<number[]>([]);
const currentRoleId = ref<number>(0);

const loadData = async () => {
  loading.value = true;
  try {
    const res = await getRoleList();
    tableData.value = res.list || [];
  } catch (error) {
    message.error('加载数据失败');
  } finally {
    loading.value = false;
  }
};

const handleAdd = () => {
  modalTitle.value = '添加角色';
  isEdit.value = false;
  currentId.value = 0;
  Object.assign(formState, {
    name: '',
    code: '',
    description: '',
    status: 1,
  });
  modalVisible.value = true;
};

const handleEdit = (row: Role) => {
  modalTitle.value = '编辑角色';
  isEdit.value = true;
  currentId.value = row.id;
  Object.assign(formState, {
    name: row.name,
    code: row.code,
    description: row.description,
    status: row.status,
  });
  modalVisible.value = true;
};

const handleDelete = async (row: Role) => {
  try {
    await deleteRole(row.id);
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
      await updateRole(currentId.value, formState);
      message.success('更新成功');
    } else {
      await createRole(formState);
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

// 权限配置
const handlePermission = async (row: Role) => {
  currentRoleId.value = row.id;
  try {
    // 加载权限树
    const tree = await getPermissionTree();
    permissionTree.value = tree;
    
    // 加载角色已有权限
    const permissionIds = await getRolePermissions(row.id);
    checkedPermissionIds.value = permissionIds;
    
    permissionModalVisible.value = true;
  } catch (error) {
    message.error('加载权限失败');
  }
};

const handlePermissionOk = async () => {
  try {
    await assignRolePermissions(currentRoleId.value, checkedPermissionIds.value as number[]);
    message.success('权限配置成功');
    permissionModalVisible.value = false;
  } catch (error) {
    message.error('权限配置失败');
  }
};

onMounted(() => {
  loadData();
});
</script>

<style scoped>
.role-list {
  padding: 0;
}
</style>
