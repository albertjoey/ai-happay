<template>
  <div class="recommend-list">
    <a-card title="推荐位管理">
      <template #extra>
        <a-space>
          <a-select v-model:value="selectedChannel" style="width: 200px" @change="loadRecommends">
            <a-select-option v-for="ch in channels" :key="ch.id" :value="ch.id">
              {{ ch.name }}
            </a-select-option>
          </a-select>
          <a-button type="primary" @click="handleAdd">
            <template #icon><PlusOutlined /></template>
            添加推荐位
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
        <vxe-column field="title" title="标题" width="150"></vxe-column>
        <vxe-column field="display_type" title="展示类型" width="120">
          <template #default="{ row }">
            <a-tag v-if="row.display_type === 'single'" color="blue">单图</a-tag>
            <a-tag v-else-if="row.display_type === 'double'" color="green">双图</a-tag>
            <a-tag v-else-if="row.display_type === 'triple'" color="orange">三图</a-tag>
            <a-tag v-else-if="row.display_type === 'list'" color="purple">列表</a-tag>
            <a-tag v-else>未知</a-tag>
          </template>
        </vxe-column>
        <vxe-column field="source_type" title="内容来源" width="120">
          <template #default="{ row }">
            <a-tag v-if="row.source_type === 'algorithm'" color="cyan">算法推荐</a-tag>
            <a-tag v-else-if="row.source_type === 'manual'" color="gold">人工推荐</a-tag>
            <a-tag v-else-if="row.source_type === 'filter'" color="lime">条件筛选</a-tag>
            <a-tag v-else>未知</a-tag>
          </template>
        </vxe-column>
        <vxe-column field="sort" title="排序" width="80"></vxe-column>
        <vxe-column field="status" title="状态" width="100">
          <template #default="{ row }">
            <a-tag v-if="row.status === 1" color="success">启用</a-tag>
            <a-tag v-else color="error">禁用</a-tag>
          </template>
        </vxe-column>
        <vxe-column field="description" title="描述" min-width="150"></vxe-column>
        <vxe-column title="操作" width="200" fixed="right">
          <template #default="{ row }">
            <a-space>
              <a-button type="link" size="small" @click="handleEdit(row)">编辑</a-button>
              <a-button type="link" size="small" @click="handlePreview(row)">预览</a-button>
              <a-popconfirm
                title="确定要删除此推荐位吗?"
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
      width="800px"
    >
      <a-form
        ref="formRef"
        :model="formState"
        :rules="formRules"
        :label-col="{ span: 5 }"
        :wrapper-col="{ span: 18 }"
      >
        <a-form-item label="标题" name="title">
          <a-input v-model:value="formState.title" placeholder="请输入标题" />
        </a-form-item>
        <a-form-item label="展示类型" name="display_type">
          <a-select v-model:value="formState.display_type">
            <a-select-option value="single">单图</a-select-option>
            <a-select-option value="double">双图</a-select-option>
            <a-select-option value="triple">三图</a-select-option>
            <a-select-option value="list">列表</a-select-option>
          </a-select>
        </a-form-item>
        <a-form-item label="内容来源" name="source_type">
          <a-radio-group v-model:value="formState.source_type">
            <a-radio value="algorithm">算法推荐</a-radio>
            <a-radio value="manual">人工推荐</a-radio>
            <a-radio value="filter">条件筛选</a-radio>
          </a-radio-group>
        </a-form-item>

        <!-- 算法推荐配置 -->
        <template v-if="formState.source_type === 'algorithm'">
          <a-form-item label="算法类型" name="filter_rule.algorithm_type">
            <a-select v-model:value="formState.filter_rule.algorithm_type">
              <a-select-option value="hot">热度排序</a-select-option>
              <a-select-option value="time">时间排序</a-select-option>
              <a-select-option value="personal">个性化推荐</a-select-option>
            </a-select>
          </a-form-item>
          <a-form-item label="内容数量" name="filter_rule.limit">
            <a-input-number v-model:value="formState.filter_rule.limit" :min="1" :max="50" />
          </a-form-item>
        </template>

        <!-- 人工推荐配置 -->
        <template v-if="formState.source_type === 'manual'">
          <a-form-item label="内容ID列表" name="content_ids">
            <a-select
              v-model:value="formState.content_ids"
              mode="tags"
              placeholder="输入内容ID后回车添加"
              :tokenSeparators="[',']"
            >
            </a-select>
          </a-form-item>
        </template>

        <!-- 条件筛选配置 -->
        <template v-if="formState.source_type === 'filter'">
          <a-form-item label="内容类型" name="filter_rule.content_type">
            <a-select v-model:value="formState.filter_rule.content_type" mode="multiple" placeholder="选择内容类型">
              <a-select-option :value="1">长视频</a-select-option>
              <a-select-option :value="2">短视频</a-select-option>
              <a-select-option :value="3">短剧</a-select-option>
              <a-select-option :value="4">漫剧</a-select-option>
              <a-select-option :value="5">小说</a-select-option>
              <a-select-option :value="6">图文</a-select-option>
            </a-select>
          </a-form-item>
          <a-form-item label="排序方式" name="filter_rule.sort_by">
            <a-select v-model:value="formState.filter_rule.sort_by">
              <a-select-option value="hot_score">热度排序</a-select-option>
              <a-select-option value="created_at">时间排序</a-select-option>
              <a-select-option value="view_count">浏览量排序</a-select-option>
            </a-select>
          </a-form-item>
          <a-form-item label="内容数量" name="filter_rule.limit">
            <a-input-number v-model:value="formState.filter_rule.limit" :min="1" :max="50" />
          </a-form-item>
        </template>

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

    <!-- 预览对话框 -->
    <a-modal
      v-model:open="previewVisible"
      title="内容预览"
      :footer="null"
      width="600px"
    >
      <a-spin :spinning="previewLoading">
        <div v-if="previewData.length > 0" class="preview-list">
          <div v-for="item in previewData" :key="item.id" class="preview-item">
            <img v-if="item.cover" :src="item.cover" class="preview-cover" />
            <div class="preview-info">
              <div class="preview-title">{{ item.title }}</div>
              <div class="preview-meta">
                <span>热度: {{ item.hot_score?.toFixed(2) || 0 }}</span>
                <span>浏览: {{ item.view_count }}</span>
              </div>
            </div>
          </div>
        </div>
        <a-empty v-else description="暂无内容" />
      </a-spin>
    </a-modal>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted } from 'vue';
import { message } from 'ant-design-vue';
import { PlusOutlined } from '@ant-design/icons-vue';
import type { FormInstance } from 'ant-design-vue';
import {
  getRecommendList,
  createRecommend,
  updateRecommend,
  deleteRecommend,
  type Recommend,
} from '@/api/recommend';
import { getChannelList, type Channel } from '@/api/channel';
import axios from 'axios';

const loading = ref(false);
const tableData = ref<Recommend[]>([]);
const channels = ref<Channel[]>([]);
const selectedChannel = ref<number>(7);
const modalVisible = ref(false);
const modalTitle = ref('添加推荐位');
const currentId = ref<number>(0);
const formRef = ref<FormInstance>();

// 预览相关
const previewVisible = ref(false);
const previewLoading = ref(false);
const previewData = ref<any[]>([]);

const getDefaultFilterRule = () => ({
  algorithm_type: 'hot',
  content_type: [],
  tag_ids: [],
  topic_ids: [],
  sort_by: 'hot_score',
  sort_order: 'desc',
  limit: 10,
});

const formState = reactive({
  title: '',
  display_type: 'single',
  source_type: 'algorithm',
  content_ids: [] as number[],
  filter_rule: getDefaultFilterRule(),
  sort: 0,
  status: 1,
  description: '',
});

const formRules = {
  title: [{ required: true, message: '请输入标题', trigger: 'blur' }],
  source_type: [{ required: true, message: '请选择内容来源', trigger: 'change' }],
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

// 加载推荐位列表
const loadRecommends = async () => {
  loading.value = true;
  try {
    const res = await getRecommendList({ channel_id: selectedChannel.value });
    tableData.value = res.list || [];
  } catch (error) {
    message.error('加载推荐位列表失败');
  } finally {
    loading.value = false;
  }
};

// 添加
const handleAdd = () => {
  modalTitle.value = '添加推荐位';
  currentId.value = 0;
  Object.assign(formState, {
    title: '',
    display_type: 'single',
    source_type: 'algorithm',
    content_ids: [],
    filter_rule: getDefaultFilterRule(),
    sort: 0,
    status: 1,
    description: '',
  });
  modalVisible.value = true;
};

// 编辑
const handleEdit = (row: Recommend) => {
  modalTitle.value = '编辑推荐位';
  currentId.value = row.id;
  Object.assign(formState, {
    title: row.title,
    display_type: row.display_type,
    source_type: row.source_type,
    content_ids: row.content_ids || [],
    filter_rule: row.filter_rule || getDefaultFilterRule(),
    sort: row.sort,
    status: row.status,
    description: row.description,
  });
  modalVisible.value = true;
};

// 预览
const handlePreview = async (row: Recommend) => {
  previewVisible.value = true;
  previewLoading.value = true;
  previewData.value = [];

  try {
    // 通过后端代理调用推荐服务
    const requestBody: any = {
      strategy: row.source_type === 'algorithm' ? 'algorithm' : 
                row.source_type === 'manual' ? 'manual' : 'filter',
      limit: row.filter_rule?.limit || 10,
    };

    if (row.source_type === 'algorithm') {
      requestBody.algorithm_type = row.filter_rule?.algorithm_type || 'hot';
      requestBody.content_type = row.filter_rule?.content_type || [];
    } else if (row.source_type === 'manual') {
      requestBody.content_ids = row.content_ids || [];
    } else {
      requestBody.content_type = row.filter_rule?.content_type || [];
      requestBody.sort_by = row.filter_rule?.sort_by || 'hot_score';
    }

    const response = await axios.post('http://localhost:4004/api/v1/recommend/preview', requestBody);
    previewData.value = response.data?.content || [];
  } catch (error: any) {
    message.error(error.response?.data?.message || '预览失败，请确保推荐服务已启动');
  } finally {
    previewLoading.value = false;
  }
};

// 删除
const handleDelete = async (row: Recommend) => {
  try {
    await deleteRecommend(row.id);
    message.success('删除成功');
    await loadRecommends();
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
      await updateRecommend(currentId.value, data);
      message.success('更新成功');
    } else {
      await createRecommend(data);
      message.success('创建成功');
    }
    modalVisible.value = false;
    await loadRecommends();
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
  loadRecommends();
});
</script>

<style scoped>
.recommend-list {
  padding: 0;
}

.preview-list {
  max-height: 400px;
  overflow-y: auto;
}

.preview-item {
  display: flex;
  padding: 12px;
  border-bottom: 1px solid #f0f0f0;
}

.preview-cover {
  width: 80px;
  height: 60px;
  object-fit: cover;
  border-radius: 4px;
  margin-right: 12px;
}

.preview-info {
  flex: 1;
}

.preview-title {
  font-weight: 500;
  margin-bottom: 8px;
}

.preview-meta {
  color: #999;
  font-size: 12px;
}

.preview-meta span {
  margin-right: 16px;
}
</style>
