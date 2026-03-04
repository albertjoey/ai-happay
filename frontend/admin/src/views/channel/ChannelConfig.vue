<template>
  <div class="channel-config">
    <a-card title="频道配置">
      <template #extra>
        <a-space>
          <a-select v-model:value="selectedChannel" style="width: 200px" @change="loadConfig">
            <a-select-option v-for="ch in channels" :key="ch.id" :value="ch.id">
              {{ ch.name }}
            </a-select-option>
          </a-select>
          <a-button type="primary" @click="saveConfig" :loading="saving">
            <template #icon><SaveOutlined /></template>
            保存配置
          </a-button>
        </a-space>
      </template>

      <a-spin :spinning="loading">
        <a-tabs v-model:activeKey="activeTab">
          <!-- Banner配置 -->
          <a-tab-pane key="banner" tab="Banner配置">
            <a-form :label-col="{ span: 4 }" :wrapper-col="{ span: 18 }">
              <a-form-item label="启用Banner">
                <a-switch v-model:checked="config.page_config.banner.enabled" />
              </a-form-item>
              <a-form-item label="选择Banner" v-if="config.page_config.banner.enabled">
                <a-select 
                  v-model:value="config.page_config.banner.banner_ids" 
                  mode="multiple"
                  placeholder="选择要显示的Banner"
                  style="width: 100%"
                >
                  <a-select-option v-for="banner in banners" :key="banner.id" :value="banner.id">
                    {{ banner.title }}
                  </a-select-option>
                </a-select>
              </a-form-item>
            </a-form>
          </a-tab-pane>

          <!-- 金刚位配置 -->
          <a-tab-pane key="diamond" tab="金刚位配置">
            <a-form :label-col="{ span: 4 }" :wrapper-col="{ span: 18 }">
              <a-form-item label="启用金刚位">
                <a-switch v-model:checked="config.page_config.diamond.enabled" />
              </a-form-item>
              <a-form-item label="显示分组" v-if="config.page_config.diamond.enabled">
                <a-checkbox-group v-model:value="config.page_config.diamond.group_ids">
                  <a-checkbox :value="1">第1组</a-checkbox>
                  <a-checkbox :value="2">第2组</a-checkbox>
                  <a-checkbox :value="3">第3组</a-checkbox>
                  <a-checkbox :value="4">第4组</a-checkbox>
                  <a-checkbox :value="5">第5组</a-checkbox>
                </a-checkbox-group>
              </a-form-item>
            </a-form>
          </a-tab-pane>

          <!-- 推荐位配置 -->
          <a-tab-pane key="recommend" tab="推荐位配置">
            <div class="config-section">
              <a-button type="dashed" block @click="addRecommend">
                <template #icon><PlusOutlined /></template>
                添加推荐位
              </a-button>
              
              <a-table 
                :dataSource="config.page_config.recommends" 
                :columns="recommendColumns"
                :pagination="false"
                rowKey="id"
                class="mt-16"
              >
                <template #bodyCell="{ column, record, index }">
                  <template v-if="column.key === 'id'">
                    <a-select v-model:value="record.id" style="width: 200px" placeholder="选择推荐位">
                      <a-select-option v-for="rec in recommendList" :key="rec.id" :value="rec.id">
                        {{ rec.title }}
                      </a-select-option>
                    </a-select>
                  </template>
                  <template v-if="column.key === 'title'">
                    <a-input v-model:value="record.title" placeholder="自定义标题（可选）" />
                  </template>
                  <template v-if="column.key === 'sort'">
                    <a-input-number v-model:value="record.sort" :min="0" :max="999" />
                  </template>
                  <template v-if="column.key === 'action'">
                    <a-button type="link" danger size="small" @click="removeRecommend(index)">
                      删除
                    </a-button>
                  </template>
                </template>
              </a-table>
            </div>
          </a-tab-pane>

          <!-- Feed流配置 -->
          <a-tab-pane key="feed" tab="Feed流配置">
            <a-form :label-col="{ span: 4 }" :wrapper-col="{ span: 18 }">
              <a-form-item label="启用Feed流">
                <a-switch v-model:checked="config.page_config.feed.enabled" />
              </a-form-item>
              <template v-if="config.page_config.feed.enabled">
                <a-form-item label="选择Feed配置">
                  <a-select v-model:value="config.page_config.feed.feed_id" style="width: 300px" placeholder="选择Feed配置">
                    <a-select-option v-for="feed in feedList" :key="feed.id" :value="feed.id">
                      {{ feed.title }}
                    </a-select-option>
                  </a-select>
                </a-form-item>
                <a-form-item label="自动加载">
                  <a-switch v-model:checked="config.page_config.feed.auto_load" />
                  <span class="ml-8 text-gray">开启后滚动到底部自动加载更多</span>
                </a-form-item>
                <a-form-item label="显示标题">
                  <a-switch v-model:checked="config.page_config.feed.show_title" />
                  <span class="ml-8 text-gray">是否显示Feed流标题</span>
                </a-form-item>
              </template>
            </a-form>
          </a-tab-pane>
        </a-tabs>
      </a-spin>
    </a-card>

    <!-- 预览 -->
    <a-card title="配置预览" class="mt-16">
      <a-button @click="previewConfig" :loading="previewing">
        <template #icon><EyeOutlined /></template>
        预览H5效果
      </a-button>
      <div class="preview-info mt-16">
        <a-descriptions :column="2" bordered size="small">
          <a-descriptions-item label="Banner">
            {{ config.page_config.banner.enabled ? '已启用' : '未启用' }}
            <span v-if="config.page_config.banner.enabled && config.page_config.banner.banner_ids?.length">
              ({{ config.page_config.banner.banner_ids.length }}个)
            </span>
          </a-descriptions-item>
          <a-descriptions-item label="金刚位">
            {{ config.page_config.diamond.enabled ? '已启用' : '未启用' }}
            <span v-if="config.page_config.diamond.enabled && config.page_config.diamond.group_ids?.length">
              ({{ config.page_config.diamond.group_ids.length }}组)
            </span>
          </a-descriptions-item>
          <a-descriptions-item label="推荐位">
            {{ config.page_config.recommends?.length || 0 }}个
          </a-descriptions-item>
          <a-descriptions-item label="Feed流">
            {{ config.page_config.feed.enabled ? '已启用' : '未启用' }}
          </a-descriptions-item>
        </a-descriptions>
      </div>
    </a-card>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted } from 'vue';
import { message } from 'ant-design-vue';
import { SaveOutlined, PlusOutlined, EyeOutlined } from '@ant-design/icons-vue';
import { getChannelList, type Channel } from '@/api/channel';
import request from '@/api/request';

interface Banner {
  id: number;
  title: string;
}

interface RecommendItem {
  id: number;
  title: string;
}

interface FeedItem {
  id: number;
  title: string;
}

const loading = ref(false);
const saving = ref(false);
const previewing = ref(false);
const channels = ref<Channel[]>([]);
const selectedChannel = ref<number>(7);
const activeTab = ref('banner');

const banners = ref<Banner[]>([]);
const recommendList = ref<RecommendItem[]>([]);
const feedList = ref<FeedItem[]>([]);

const getDefaultPageConfig = () => ({
  banner: {
    enabled: true,
    banner_ids: [] as number[],
  },
  diamond: {
    enabled: true,
    group_ids: [1] as number[],
  },
  recommends: [] as Array<{ id: number; title: string; sort: number }>,
  feed: {
    enabled: true,
    feed_id: undefined as number | undefined,
    auto_load: true,
    show_title: true,
  },
});

const config = reactive({
  channel_id: 1,
  content_type: {} as Record<string, boolean>,
  display_mode: 'default',
  custom_data: {} as Record<string, any>,
  page_config: getDefaultPageConfig(),
});

const recommendColumns = [
  { title: '推荐位', key: 'id', width: 220 },
  { title: '自定义标题', key: 'title' },
  { title: '排序', key: 'sort', width: 100 },
  { title: '操作', key: 'action', width: 80 },
];

// 加载频道列表
const loadChannels = async () => {
  try {
    const res = await getChannelList({ page: 1, page_size: 100 });
    channels.value = res.list || [];
  } catch (error) {
    message.error('加载频道列表失败');
  }
};

// 加载配置
const loadConfig = async () => {
  loading.value = true;
  try {
    const res = await request.get(`/channel/config/${selectedChannel.value}`);
    Object.assign(config, {
      channel_id: res.channel_id,
      content_type: res.content_type || {},
      display_mode: res.display_mode || 'default',
      custom_data: res.custom_data || {},
      page_config: res.page_config || getDefaultPageConfig(),
    });
    
    // 确保page_config有默认值
    if (!config.page_config) {
      config.page_config = getDefaultPageConfig();
    }
    if (!config.page_config.banner) {
      config.page_config.banner = { enabled: true, banner_ids: [] };
    }
    if (!config.page_config.diamond) {
      config.page_config.diamond = { enabled: true, group_ids: [1] };
    }
    if (!config.page_config.recommends) {
      config.page_config.recommends = [];
    }
    if (!config.page_config.feed) {
      config.page_config.feed = { enabled: true, auto_load: true, show_title: true };
    }
    
    // 加载关联数据
    await loadRelatedData();
  } catch (error) {
    message.error('加载配置失败');
  } finally {
    loading.value = false;
  }
};

// 加载关联数据
const loadRelatedData = async () => {
  try {
    // 加载Banner列表
    const bannerRes = await request.get('/banner/list', { params: { channel_id: selectedChannel.value } });
    banners.value = bannerRes.list || [];
    
    // 加载推荐位列表
    const recRes = await request.get('/recommend/list', { params: { channel_id: selectedChannel.value } });
    recommendList.value = recRes.list || [];
    
    // 加载Feed配置列表
    const feedRes = await request.get('/feed-config/list', { params: { channel_id: selectedChannel.value } });
    feedList.value = feedRes.list || [];
    
    // 加载金刚位列表（用于显示分组）
    // const diamondRes = await request.get('/diamond/list', { params: { channel_id: selectedChannel.value } });
    // diamonds.value = diamondRes.list || [];
  } catch (error) {
    console.error('加载关联数据失败', error);
  }
};

// 添加推荐位
const addRecommend = () => {
  config.page_config.recommends.push({
    id: 0,
    title: '',
    sort: config.page_config.recommends.length,
  });
};

// 删除推荐位
const removeRecommend = (index: number) => {
  config.page_config.recommends.splice(index, 1);
};

// 保存配置
const saveConfig = async () => {
  saving.value = true;
  try {
    // 构建保存数据
    const saveData = {
      channel_id: selectedChannel.value,
      content_type: config.content_type,
      display_mode: config.display_mode,
      custom_data: {
        ...config.custom_data,
        page_config: config.page_config,
      },
    };
    
    await request.put(`/channel/config/${selectedChannel.value}`, saveData);
    message.success('保存成功');
  } catch (error) {
    message.error('保存失败');
  } finally {
    saving.value = false;
  }
};

// 预览配置
const previewConfig = () => {
  const h5Url = `http://localhost:3000/channel/${selectedChannel.value}`;
  window.open(h5Url, '_blank');
};

onMounted(() => {
  loadChannels();
  loadConfig();
});
</script>

<style scoped>
.channel-config {
  padding: 0;
}

.mt-16 {
  margin-top: 16px;
}

.ml-8 {
  margin-left: 8px;
}

.text-gray {
  color: #999;
}

.config-section {
  padding: 16px 0;
}
</style>
