<template>
  <div class="discover-config">
    <a-card title="发现页模块配置" :bordered="false">
      <a-table
        :columns="columns"
        :data-source="configList"
        :loading="loading"
        row-key="id"
      >
        <template #bodyCell="{ column, record }">
          <template v-if="column.key === 'enabled'">
            <a-switch
              v-model:checked="record.enabled"
              @change="handleEnabledChange(record)"
            />
          </template>
          <template v-if="column.key === 'action'">
            <a-button type="link" @click="handleEdit(record)">编辑</a-button>
            <a-button type="link" @click="handleManageItems(record)">管理内容</a-button>
          </template>
        </template>
      </a-table>
    </a-card>

    <!-- 编辑配置弹窗 -->
    <a-modal
      v-model:open="editVisible"
      title="编辑模块配置"
      @ok="handleEditOk"
    >
      <a-form :model="editForm" :label-col="{ span: 6 }">
        <a-form-item label="模块名称">
          <a-input v-model:value="editForm.module" disabled />
        </a-form-item>
        <a-form-item label="模块标题">
          <a-input v-model:value="editForm.title" />
        </a-form-item>
        <a-form-item label="排序">
          <a-input-number v-model:value="editForm.sort_order" :min="1" />
        </a-form-item>
        <a-form-item label="是否启用">
          <a-switch v-model:checked="editForm.enabled" />
        </a-form-item>
      </a-form>
    </a-modal>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue';
import { message } from 'ant-design-vue';
import { getDiscoverConfigList, updateDiscoverConfig, type DiscoverConfig } from '@/api/discover';

const loading = ref(false);
const configList = ref<DiscoverConfig[]>([]);
const editVisible = ref(false);
const editForm = ref<Partial<DiscoverConfig>>({});

const columns = [
  {
    title: 'ID',
    dataIndex: 'id',
    key: 'id',
    width: 80,
  },
  {
    title: '模块名称',
    dataIndex: 'module',
    key: 'module',
  },
  {
    title: '模块标题',
    dataIndex: 'title',
    key: 'title',
  },
  {
    title: '排序',
    dataIndex: 'sort_order',
    key: 'sort_order',
    width: 100,
  },
  {
    title: '是否启用',
    dataIndex: 'enabled',
    key: 'enabled',
    width: 120,
  },
  {
    title: '操作',
    key: 'action',
    width: 200,
  },
];

// 加载配置列表
const loadConfigList = async () => {
  loading.value = true;
  try {
    const res = await getDiscoverConfigList();
    configList.value = res.data.list || [];
  } catch (error) {
    message.error('加载配置失败');
  } finally {
    loading.value = false;
  }
};

// 启用/禁用状态改变
const handleEnabledChange = async (record: DiscoverConfig) => {
  try {
    await updateDiscoverConfig({
      id: record.id,
      enabled: record.enabled,
    });
    message.success('更新成功');
  } catch (error) {
    message.error('更新失败');
    record.enabled = !record.enabled;
  }
};

// 编辑配置
const handleEdit = (record: DiscoverConfig) => {
  editForm.value = { ...record };
  editVisible.value = true;
};

// 编辑确认
const handleEditOk = async () => {
  try {
    await updateDiscoverConfig(editForm.value);
    message.success('更新成功');
    editVisible.value = false;
    loadConfigList();
  } catch (error) {
    message.error('更新失败');
  }
};

// 管理内容
const handleManageItems = (record: DiscoverConfig) => {
  // 跳转到内容管理页面
  window.location.href = `/discover/items?module=${record.module}`;
};

onMounted(() => {
  loadConfigList();
});
</script>

<style scoped>
.discover-config {
  padding: 24px;
}
</style>
