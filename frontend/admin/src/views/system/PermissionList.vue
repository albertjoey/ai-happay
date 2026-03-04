<template>
  <div class="permission-list">
    <a-card title="权限管理">
      <template #extra>
        <a-button type="primary" @click="handleAdd">
          <template #icon><PlusOutlined /></template>
          添加权限
        </a-button>
      </template>

      <a-table
        :columns="columns"
        :data-source="tableData"
        :loading="loading"
        :pagination="false"
        row-key="id"
        :default-expand-all-rows="true"
      >
        <template #bodyCell="{ column, record }">
          <template v-if="column.key === 'type'">
            <a-tag v-if="record.type === 'menu'" color="blue">菜单</a-tag>
            <a-tag v-else-if="record.type === 'button'" color="green">按钮</a-tag>
            <a-tag v-else-if="record.type === 'api'" color="orange">接口</a-tag>
          </template>
          <template v-else-if="column.key === 'status'">
            <a-tag v-if="record.status === 1" color="success">启用</a-tag>
            <a-tag v-else color="error">禁用</a-tag>
          </template>
          <template v-else-if="column.key === 'action'">
            <a-space>
              <a-button type="link" size="small" @click="handleEdit(record)">编辑</a-button>
              <a-popconfirm title="确定要删除此权限吗？" @confirm="handleDelete(record)">
                <a-button type="link" size="small" danger>删除</a-button>
              </a-popconfirm>
            </a-space>
          </template>
        </template>
      </a-table>
    </a-card>

    <!-- 添加/编辑权限对话框 -->
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
        <a-form-item label="权限名称" name="name">
          <a-input v-model:value="formState.name" placeholder="请输入权限名称" />
        </a-form-item>
        <a-form-item label="权限编码" name="code">
          <a-input v-model:value="formState.code" placeholder="请输入权限编码" :disabled="isEdit" />
        </a-form-item>
        <a-form-item label="权限类型" name="type">
          <a-select v-model:value="formState.type">
            <a-select-option value="menu">菜单</a-select-option>
            <a-select-option value="button">按钮</a-select-option>
            <a-select-option value="api">接口</a-select-option>
          </a-select>
        </a-form-item>
        <a-form-item label="父级权限" name="parent_id">
          <a-tree-select
            v-model:value="formState.parent_id"
            :tree-data="permissionTree"
            :field-names="{ label: 'name', value: 'id', children: 'children' }"
            placeholder="请选择父级权限"
            allow-clear
            tree-default-expand-all
          />
        </a-form-item>
        <a-form-item v-if="formState.type === 'menu'" label="路由路径" name="path">
          <a-input v-model:value="formState.path" placeholder="请输入路由路径" />
        </a-form-item>
        <a-form-item v-if="formState.type === 'menu'" label="图标" name="icon">
          <a-input v-model:value="formState.icon" placeholder="请输入图标名称" />
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
  getPermissionTree,
  createPermission,
  updatePermission,
  deletePermission,
  type Permission,
} from '@/api/rbac';

const loading = ref(false);
const tableData = ref<Permission[]>([]);
const permissionTree = ref<Permission[]>([]);
const modalVisible = ref(false);
const modalTitle = ref('添加权限');
const isEdit = ref(false);
const currentId = ref<number>(0);
const formRef = ref<FormInstance>();

const formState = reactive({
  name: '',
  code: '',
  type: 'menu',
  parent_id: 0,
  path: '',
  icon: '',
  status: 1,
});

const formRules = {
  name: [{ required: true, message: '请输入权限名称', trigger: 'blur' }],
  code: [{ required: true, message: '请输入权限编码', trigger: 'blur' }],
  type: [{ required: true, message: '请选择权限类型', trigger: 'change' }],
};

const columns = [
  { title: '权限名称', dataIndex: 'name', key: 'name', width: 200 },
  { title: '权限编码', dataIndex: 'code', key: 'code', width: 200 },
  { title: '类型', dataIndex: 'type', key: 'type', width: 100 },
  { title: '路由路径', dataIndex: 'path', key: 'path', width: 200 },
  { title: '图标', dataIndex: 'icon', key: 'icon', width: 150 },
  { title: '状态', dataIndex: 'status', key: 'status', width: 100 },
  { title: '操作', key: 'action', width: 150, fixed: 'right' },
];

const loadData = async () => {
  loading.value = true;
  try {
    const tree = await getPermissionTree();
    tableData.value = tree;
    permissionTree.value = [{ id: 0, name: '顶级权限', children: tree } as any];
  } catch (error) {
    message.error('加载数据失败');
  } finally {
    loading.value = false;
  }
};

const handleAdd = () => {
  modalTitle.value = '添加权限';
  isEdit.value = false;
  currentId.value = 0;
  Object.assign(formState, {
    name: '',
    code: '',
    type: 'menu',
    parent_id: 0,
    path: '',
    icon: '',
    status: 1,
  });
  modalVisible.value = true;
};

const handleEdit = (row: Permission) => {
  modalTitle.value = '编辑权限';
  isEdit.value = true;
  currentId.value = row.id;
  Object.assign(formState, {
    name: row.name,
    code: row.code,
    type: row.type,
    parent_id: row.parent_id,
    path: row.path,
    icon: row.icon,
    status: row.status,
  });
  modalVisible.value = true;
};

const handleDelete = async (row: Permission) => {
  try {
    await deletePermission(row.id);
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
      await updatePermission(currentId.value, formState);
      message.success('更新成功');
    } else {
      await createPermission(formState);
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

onMounted(() => {
  loadData();
});
</script>

<style scoped>
.permission-list {
  padding: 0;
}
</style>
