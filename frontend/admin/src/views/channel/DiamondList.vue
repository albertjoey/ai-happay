<template>
  <div class="diamond-list">
    <a-card title="金刚位管理">
      <template #extra>
        <a-space>
          <a-select v-model:value="selectedChannel" style="width: 200px" @change="loadDiamonds">
            <a-select-option v-for="ch in channels" :key="ch.id" :value="ch.id">
              {{ ch.name }}
            </a-select-option>
          </a-select>
          <a-button type="primary" @click="handleAdd">
            <template #icon><PlusOutlined /></template>
            添加金刚位
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
          <template v-if="column.key === 'group_id'">
            <a-tag color="blue">第{{ record.group_id }}组</a-tag>
          </template>
          <template v-else-if="column.key === 'icon'">
            <span class="icon-preview">{{ record.icon }}</span>
          </template>
          <template v-else-if="column.key === 'link_type'">
            <a-tag v-if="record.link_type === 'channel'" color="green">频道</a-tag>
            <a-tag v-else-if="record.link_type === 'topic'" color="blue">话题</a-tag>
            <a-tag v-else-if="record.link_type === 'content'" color="orange">内容</a-tag>
            <a-tag v-else color="purple">外链</a-tag>
          </template>
          <template v-else-if="column.key === 'status'">
            <a-tag v-if="record.status === 1" color="success">启用</a-tag>
            <a-tag v-else color="error">禁用</a-tag>
          </template>
          <template v-else-if="column.key === 'action'">
            <a-space>
              <a-button type="link" size="small" @click="handleEdit(record)">编辑</a-button>
              <a-popconfirm
                title="确定要删除此金刚位吗?"
                @confirm="handleDelete(record)"
              >
                <a-button type="link" size="small" danger>删除</a-button>
              </a-popconfirm>
            </a-space>
          </template>
        </template>
      </a-table>
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
        <a-form-item label="分组" name="group_id">
          <a-input-number v-model:value="formState.group_id" :min="1" :max="5" />
        </a-form-item>
        <a-form-item label="排序" name="sort">
          <a-input-number v-model:value="formState.sort" :min="0" :max="999" />
        </a-form-item>
        <a-form-item label="图标" name="icon">
          <a-input v-model:value="formState.icon" placeholder="请输入图标(emoji)" />
        </a-form-item>
        <a-form-item label="标题" name="title">
          <a-input v-model:value="formState.title" placeholder="请输入标题" />
        </a-form-item>
        <a-form-item label="链接类型" name="link_type">
          <a-select v-model:value="formState.link_type">
            <a-select-option value="channel">频道</a-select-option>
            <a-select-option value="topic">话题</a-select-option>
            <a-select-option value="content">内容</a-select-option>
            <a-select-option value="external">外链</a-select-option>
          </a-select>
        </a-form-item>
        <a-form-item label="链接值" name="link_value">
          <a-input v-model:value="formState.link_value" placeholder="请输入链接值" />
        </a-form-item>
        <a-form-item label="状态" name="status">
          <a-radio-group v-model:value="formState.status">
            <a-radio :value="1">启用</a-radio>
            <a-radio :value="0">禁用</a-radio>
          </a-radio-group>
        </a-form-item>
        <a-form-item label="描述" name="description">
          <a-textarea v-model:value="formState.description" placeholder="请输入描述" :rows="3" />
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
  getDiamondList,
  createDiamond,
  updateDiamond,
  deleteDiamond,
  type Diamond,
} from '@/api/diamond';
import { getChannelList, type Channel } from '@/api/channel';

const loading = ref(false);
const tableData = ref<Diamond[]>([]);
const channels = ref<Channel[]>([]);
const selectedChannel = ref<number>(0);
const modalVisible = ref(false);
const modalTitle = ref('添加金刚位');
const currentId = ref<number>(0);
const formRef = ref<FormInstance>();

const formState = reactive({
  group_id: 1,
  sort: 0,
  icon: '',
  title: '',
  link_type: 'channel',
  link_value: '',
  status: 1,
  description: '',
});

// 表格列定义
const columns = [
  { title: 'ID', dataIndex: 'id', key: 'id', width: 80 },
  { title: '分组', dataIndex: 'group_id', key: 'group_id', width: 80 },
  { title: '排序', dataIndex: 'sort', key: 'sort', width: 80 },
  { title: '图标', dataIndex: 'icon', key: 'icon', width: 80 },
  { title: '标题', dataIndex: 'title', key: 'title', width: 150 },
  { title: '链接类型', dataIndex: 'link_type', key: 'link_type', width: 120 },
  { title: '链接值', dataIndex: 'link_value', key: 'link_value', width: 150 },
  { title: '状态', dataIndex: 'status', key: 'status', width: 100 },
  { title: '操作', key: 'action', width: 150, fixed: 'right' as const },
];

const formRules = {
  title: [{ required: true, message: '请输入标题', trigger: 'blur' }],
  link_type: [{ required: true, message: '请选择链接类型', trigger: 'change' }],
  link_value: [{ required: true, message: '请输入链接值', trigger: 'blur' }],
};

// 加载频道列表
const loadChannels = async () => {
  try {
    const res = await getChannelList({ page: 1, page_size: 100 });
    channels.value = res.list || [];
    // 自动选择第一个频道
    if (channels.value.length > 0 && selectedChannel.value === 0) {
      selectedChannel.value = channels.value[0].id;
      loadDiamonds();
    }
  } catch (error) {
    message.error('加载频道列表失败');
  }
};

// 加载金刚位列表
const loadDiamonds = async () => {
  loading.value = true;
  try {
    const res = await getDiamondList({ channel_id: selectedChannel.value });
    tableData.value = res.list || [];
  } catch (error) {
    message.error('加载金刚位列表失败');
  } finally {
    loading.value = false;
  }
};

// 添加
const handleAdd = () => {
  modalTitle.value = '添加金刚位';
  currentId.value = 0;
  Object.assign(formState, {
    group_id: 1,
    sort: 0,
    icon: '',
    title: '',
    link_type: 'channel',
    link_value: '',
    status: 1,
    description: '',
  });
  modalVisible.value = true;
};

// 编辑
const handleEdit = (row: Diamond) => {
  modalTitle.value = '编辑金刚位';
  currentId.value = row.id;
  Object.assign(formState, {
    group_id: row.group_id,
    sort: row.sort,
    icon: row.icon,
    title: row.title,
    link_type: row.link_type,
    link_value: row.link_value,
    status: row.status,
    description: row.description,
  });
  modalVisible.value = true;
};

// 删除
const handleDelete = async (row: Diamond) => {
  try {
    await deleteDiamond(row.id);
    message.success('删除成功');
    await loadDiamonds();
  } catch (error) {
    message.error('删除失败');
  }
};

// 提交
const handleModalOk = async () => {
  try {
    await formRef.value?.validate();
    const data = {
      ...formState,
      channel_id: selectedChannel.value,
    };

    if (currentId.value) {
      await updateDiamond(currentId.value, data);
      message.success('更新成功');
    } else {
      await createDiamond(data);
      message.success('创建成功');
    }
    modalVisible.value = false;
    await loadDiamonds();
  } catch (error) {
    message.error('操作失败');
  }
};

// 取消
const handleModalCancel = () => {
  modalVisible.value = false;
  formRef.value?.resetFields();
};

onMounted(() => {
  loadChannels();
  loadDiamonds();
});
</script>

<style scoped>
.diamond-list {
  padding: 0;
}

.icon-preview {
  font-size: 20px;
}
</style>
