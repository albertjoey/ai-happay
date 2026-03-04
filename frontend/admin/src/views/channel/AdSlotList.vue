<template>
  <div class="ad-slot-list">
    <a-card title="广告位管理">
      <template #extra>
        <a-space>
          <a-select v-model:value="selectedChannel" style="width: 200px" @change="loadAdSlots">
            <a-select-option v-for="ch in channels" :key="ch.id" :value="ch.id">
              {{ ch.name }}
            </a-select-option>
          </a-select>
          <a-button type="primary" @click="handleAdd">
            <template #icon><PlusOutlined /></template>
            添加广告位
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
        <vxe-column field="name" title="名称" width="150"></vxe-column>
        <vxe-column field="insert_type" title="插入方式" width="120">
          <template #default="{ row }">
            <a-tag v-if="row.insert_type === 'fixed'" color="blue">固定位置</a-tag>
            <a-tag v-else-if="row.insert_type === 'interval'" color="green">间隔插入</a-tag>
            <a-tag v-else color="orange">随机插入</a-tag>
          </template>
        </vxe-column>
        <vxe-column field="ad_type" title="广告类型" width="100">
          <template #default="{ row }">
            <a-tag v-if="row.ad_type === 'image'" color="purple">图片</a-tag>
            <a-tag v-else color="cyan">视频</a-tag>
          </template>
        </vxe-column>
        <vxe-column field="insert_rule" title="插入规则" min-width="200">
          <template #default="{ row }">
            <div v-if="row.insert_type === 'fixed'">
              第{{ row.insert_rule.fixed_position }}条后插入,最多{{ row.insert_rule.max_count }}个
            </div>
            <div v-else-if="row.insert_type === 'interval'">
              每{{ row.insert_rule.interval }}条插入,最多{{ row.insert_rule.max_count }}个
            </div>
            <div v-else>
              随机插入,最多{{ row.insert_rule.max_count }}个
            </div>
          </template>
        </vxe-column>
        <vxe-column field="status" title="状态" width="100">
          <template #default="{ row }">
            <a-tag v-if="row.status === 1" color="success">启用</a-tag>
            <a-tag v-else color="error">禁用</a-tag>
          </template>
        </vxe-column>
        <vxe-column field="description" title="描述" min-width="150"></vxe-column>
        <vxe-column title="操作" width="150" fixed="right">
          <template #default="{ row }">
            <a-space>
              <a-button type="link" size="small" @click="handleEdit(row)">编辑</a-button>
              <a-popconfirm
                title="确定要删除此广告位吗?"
                @confirm="handleDelete(row)"
              >
                <a-button type="link" size="small" danger>删除</a-button>
              </a-popconfirm>
            </a-space>
          </template>
        </vxe-column>
      </vxe-table>
    </a-card>

    <!-- 添加/编辑对话框 -->
    <a-modal
      v-model:open="modalVisible"
      :title="modalTitle"
      @ok="handleModalOk"
      @cancel="handleModalCancel"
      width="700px"
    >
      <a-form
        ref="formRef"
        :model="formState"
        :rules="formRules"
        :label-col="{ span: 6 }"
        :wrapper-col="{ span: 16 }"
      >
        <a-form-item label="名称" name="name">
          <a-input v-model:value="formState.name" placeholder="请输入广告位名称" />
        </a-form-item>
        <a-form-item label="插入方式" name="insert_type">
          <a-select v-model:value="formState.insert_type" @change="handleInsertTypeChange">
            <a-select-option value="fixed">固定位置</a-select-option>
            <a-select-option value="interval">间隔插入</a-select-option>
            <a-select-option value="random">随机插入</a-select-option>
          </a-select>
        </a-form-item>
        
        <!-- 固定位置配置 -->
        <a-form-item v-if="formState.insert_type === 'fixed'" label="固定位置">
          <a-input-number v-model:value="formState.insert_rule.fixed_position" :min="1" placeholder="第N条后插入" />
        </a-form-item>
        
        <!-- 间隔插入配置 -->
        <a-form-item v-if="formState.insert_type === 'interval'" label="间隔数量">
          <a-input-number v-model:value="formState.insert_rule.interval" :min="1" placeholder="每N条插入" />
        </a-form-item>
        
        <a-form-item label="最大插入数量">
          <a-input-number v-model:value="formState.insert_rule.max_count" :min="1" :max="10" />
        </a-form-item>
        
        <a-form-item label="广告类型" name="ad_type">
          <a-select v-model:value="formState.ad_type">
            <a-select-option value="image">图片广告</a-select-option>
            <a-select-option value="video">视频广告</a-select-option>
          </a-select>
        </a-form-item>
        
        <a-form-item v-if="formState.ad_type === 'image'" label="图片地址">
          <a-input v-model:value="formState.ad_content.image_url" placeholder="请输入图片URL" />
        </a-form-item>
        
        <a-form-item v-if="formState.ad_type === 'video'" label="视频地址">
          <a-input v-model:value="formState.ad_content.video_url" placeholder="请输入视频URL" />
        </a-form-item>
        
        <a-form-item v-if="formState.ad_type === 'video'" label="视频时长">
          <a-input-number v-model:value="formState.ad_content.duration" :min="1" placeholder="秒" />
        </a-form-item>
        
        <a-form-item label="广告标题">
          <a-input v-model:value="formState.ad_content.title" placeholder="请输入广告标题" />
        </a-form-item>
        
        <a-form-item label="跳转链接">
          <a-input v-model:value="formState.link_url" placeholder="请输入跳转链接" />
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
  getAdSlotList,
  createAdSlot,
  updateAdSlot,
  deleteAdSlot,
  type AdSlot,
} from '@/api/adSlot';
import { getChannelList, type Channel } from '@/api/channel';

const loading = ref(false);
const tableData = ref<AdSlot[]>([]);
const channels = ref<Channel[]>([]);
const selectedChannel = ref<number>(7);
const modalVisible = ref(false);
const modalTitle = ref('添加广告位');
const currentId = ref<number>(0);
const formRef = ref<FormInstance>();

const formState = reactive({
  name: '',
  insert_type: 'fixed',
  insert_rule: {
    fixed_position: 3,
    interval: 5,
    max_count: 1,
  },
  ad_type: 'image',
  ad_content: {
    image_url: '',
    video_url: '',
    title: '',
    duration: 15,
  },
  link_url: '',
  sort: 0,
  status: 1,
  description: '',
});

const formRules = {
  name: [{ required: true, message: '请输入名称', trigger: 'blur' }],
  insert_type: [{ required: true, message: '请选择插入方式', trigger: 'change' }],
  ad_type: [{ required: true, message: '请选择广告类型', trigger: 'change' }],
};

// 加载频道列表
const loadChannels = async () => {
  try {
    const res = await getChannelList({ page: 1, page_size: 100 });
    channels.value = res.list || [];
  } catch (error) {
    message.error('加载频道列表失败');
  }
};

// 加载广告位列表
const loadAdSlots = async () => {
  loading.value = true;
  try {
    const res = await getAdSlotList({ channel_id: selectedChannel.value });
    tableData.value = res.list || [];
  } catch (error) {
    message.error('加载广告位列表失败');
  } finally {
    loading.value = false;
  }
};

// 插入方式改变
const handleInsertTypeChange = () => {
  // 重置插入规则
  formState.insert_rule = {
    fixed_position: 3,
    interval: 5,
    max_count: 1,
  };
};

// 添加
const handleAdd = () => {
  modalTitle.value = '添加广告位';
  currentId.value = 0;
  Object.assign(formState, {
    name: '',
    insert_type: 'fixed',
    insert_rule: {
      fixed_position: 3,
      interval: 5,
      max_count: 1,
    },
    ad_type: 'image',
    ad_content: {
      image_url: '',
      video_url: '',
      title: '',
      duration: 15,
    },
    link_url: '',
    sort: 0,
    status: 1,
    description: '',
  });
  modalVisible.value = true;
};

// 编辑
const handleEdit = (row: AdSlot) => {
  modalTitle.value = '编辑广告位';
  currentId.value = row.id;
  Object.assign(formState, {
    name: row.name,
    insert_type: row.insert_type,
    insert_rule: row.insert_rule || {},
    ad_type: row.ad_type,
    ad_content: row.ad_content || {},
    link_url: row.link_url,
    sort: row.sort,
    status: row.status,
    description: row.description,
  });
  modalVisible.value = true;
};

// 删除
const handleDelete = async (row: AdSlot) => {
  try {
    await deleteAdSlot(row.id);
    message.success('删除成功');
    await loadAdSlots();
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
      await updateAdSlot(currentId.value, data);
      message.success('更新成功');
    } else {
      await createAdSlot(data);
      message.success('创建成功');
    }
    modalVisible.value = false;
    await loadAdSlots();
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
  loadAdSlots();
});
</script>

<style scoped>
.ad-slot-list {
  padding: 0;
}
</style>
