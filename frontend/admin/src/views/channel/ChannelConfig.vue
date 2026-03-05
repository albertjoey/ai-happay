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
                    <div class="banner-option">
                      <img :src="banner.image" class="banner-thumb" />
                      <span>{{ banner.title }}</span>
                    </div>
                  </a-select-option>
                </a-select>
                <!-- 显示已选择的Banner预览 -->
                <div v-if="config.page_config.banner.banner_ids?.length" class="selected-preview">
                  <div class="preview-title">已选择 {{ config.page_config.banner.banner_ids.length }} 个Banner:</div>
                  <div class="preview-items">
                    <div v-for="id in config.page_config.banner.banner_ids" :key="id" class="preview-item">
                      <img :src="getBannerById(id)?.image" class="preview-thumb" />
                      <span>{{ getBannerById(id)?.title }}</span>
                    </div>
                  </div>
                </div>
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
                  <a-checkbox v-for="group in diamondGroups" :key="group.id" :value="group.id">
                    第{{ group.id }}组 ({{ group.count }}个)
                  </a-checkbox>
                </a-checkbox-group>
                <!-- 显示已选择分组的金刚位预览 -->
                <div v-if="config.page_config.diamond.group_ids?.length" class="selected-preview">
                  <div class="preview-title">已选择 {{ config.page_config.diamond.group_ids.length }} 个分组:</div>
                  <div class="preview-items diamond-preview">
                    <div v-for="group in getSelectedDiamondGroups()" :key="group.id" class="diamond-group">
                      <div class="group-title">第{{ group.id }}组 ({{ group.items.length }}个)</div>
                      <div class="group-items">
                        <div v-for="item in group.items" :key="item.id" class="diamond-item">
                          <span class="diamond-icon">{{ item.icon }}</span>
                          <span class="diamond-title">{{ item.title }}</span>
                        </div>
                      </div>
                    </div>
                  </div>
                </div>
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
                        {{ rec.title }} ({{ rec.materials?.length || 0 }}个物料)
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

              <!-- 显示已选择推荐位的物料预览 -->
              <div v-if="config.page_config.recommends?.length" class="selected-preview mt-16">
                <div class="preview-title">已配置推荐位预览:</div>
                <div v-for="rec in config.page_config.recommends" :key="rec.id" class="recommend-preview">
                  <div class="recommend-title">{{ rec.title || getRecommendById(rec.id)?.title }}</div>
                  <div class="recommend-materials">
                    <div v-for="m in (getRecommendById(rec.id)?.materials || []).slice(0, 5)" :key="m.id" class="material-card">
                      <img :src="m.cover_url" class="material-cover" />
                      <div class="material-info">
                        <div class="material-name">{{ m.title }}</div>
                        <div class="material-stats">{{ formatNumber(m.view_count) }}次播放</div>
                      </div>
                    </div>
                    <div v-if="(getRecommendById(rec.id)?.materials?.length || 0) > 5" class="more-hint">
                      还有 {{ (getRecommendById(rec.id)?.materials?.length || 0) - 5 }} 个...
                    </div>
                  </div>
                </div>
              </div>
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
                      {{ feed.title }} ({{ feed.materials?.length || 0 }}个物料)
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
                <!-- 显示已选择Feed的物料预览 -->
                <div v-if="config.page_config.feed.feed_id" class="selected-preview">
                  <div class="preview-title">Feed流物料预览:</div>
                  <div class="feed-materials">
                    <div v-for="m in (getFeedById(config.page_config.feed.feed_id)?.materials || []).slice(0, 6)" :key="m.id" class="material-card">
                      <img :src="m.cover_url" class="material-cover" />
                      <div class="material-info">
                        <div class="material-name">{{ m.title }}</div>
                        <div class="material-stats">{{ formatNumber(m.view_count) }}次播放</div>
                      </div>
                    </div>
                    <div v-if="(getFeedById(config.page_config.feed.feed_id)?.materials?.length || 0) > 6" class="more-hint">
                      还有 {{ (getFeedById(config.page_config.feed.feed_id)?.materials?.length || 0) - 6 }} 个...
                    </div>
                  </div>
                </div>
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
              ({{ config.page_config.diamond.group_ids.length }}组, 共{{ getSelectedDiamondCount() }}个)
            </span>
          </a-descriptions-item>
          <a-descriptions-item label="推荐位">
            {{ config.page_config.recommends?.length || 0 }}个
          </a-descriptions-item>
          <a-descriptions-item label="Feed流">
            {{ config.page_config.feed.enabled ? '已启用' : '未启用' }}
            <span v-if="config.page_config.feed.enabled && config.page_config.feed.feed_id">
              ({{ getFeedById(config.page_config.feed.feed_id)?.title }})
            </span>
          </a-descriptions-item>
        </a-descriptions>
      </div>
    </a-card>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted, computed } from 'vue';
import { useRoute } from 'vue-router';
import { message } from 'ant-design-vue';
import { SaveOutlined, PlusOutlined, EyeOutlined } from '@ant-design/icons-vue';
import { getChannelList, type Channel } from '@/api/channel';
import request from '@/api/request';

interface Banner {
  id: number;
  title: string;
  image: string;
}

interface Diamond {
  id: number;
  group_id: number;
  title: string;
  icon: string;
  link_type: string;
  link_value: string;
}

interface RecommendItem {
  id: number;
  title: string;
  materials?: any[];
}

interface FeedItem {
  id: number;
  title: string;
  materials?: any[];
}

const route = useRoute();
const loading = ref(false);
const saving = ref(false);
const previewing = ref(false);
const channels = ref<Channel[]>([]);
const selectedChannel = ref<number>(7);
const activeTab = ref('banner');

const banners = ref<Banner[]>([]);
const diamonds = ref<Diamond[]>([]);
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

// 按分组整理金刚位
const diamondGroups = computed(() => {
  const groups: Record<number, Diamond[]> = {};
  diamonds.value.forEach(d => {
    if (!groups[d.group_id]) {
      groups[d.group_id] = [];
    }
    groups[d.group_id].push(d);
  });
  return Object.entries(groups).map(([id, items]) => ({
    id: parseInt(id),
    count: items.length,
    items,
  }));
});

// 获取Banner
const getBannerById = (id: number) => banners.value.find(b => b.id === id);

// 获取推荐位
const getRecommendById = (id: number) => recommendList.value.find(r => r.id === id);

// 获取Feed配置
const getFeedById = (id: number) => feedList.value.find(f => f.id === id);

// 获取已选择的金刚位分组
const getSelectedDiamondGroups = () => {
  return diamondGroups.value.filter(g => config.page_config.diamond.group_ids?.includes(g.id));
};

// 获取已选择的金刚位总数
const getSelectedDiamondCount = () => {
  return getSelectedDiamondGroups().reduce((sum, g) => sum + g.items.length, 0);
};

// 格式化数字
const formatNumber = (num: number) => {
  if (num >= 10000) return (num / 10000).toFixed(1) + 'w';
  if (num >= 1000) return (num / 1000).toFixed(1) + 'k';
  return num.toString();
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
    
    // 加载金刚位列表
    const diamondRes = await request.get('/diamond/list', { params: { channel_id: selectedChannel.value } });
    diamonds.value = diamondRes.list || [];
    
    // 加载推荐位列表
    const recRes = await request.get('/recommend/list', { params: { channel_id: selectedChannel.value } });
    recommendList.value = recRes.list || [];
    
    // 加载Feed配置列表
    const feedRes = await request.get('/feed-config/list', { params: { channel_id: selectedChannel.value } });
    feedList.value = feedRes.list || [];
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
  const h5Url = `http://localhost:4000`;
  window.open(h5Url, '_blank');
};

onMounted(async () => {
  await loadChannels();
  
  // 从URL参数获取频道ID
  const channelId = route.query.channel_id;
  if (channelId) {
    selectedChannel.value = parseInt(channelId as string);
  }
  
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

/* Banner选项样式 */
.banner-option {
  display: flex;
  align-items: center;
  gap: 8px;
}

.banner-thumb {
  width: 60px;
  height: 30px;
  object-fit: cover;
  border-radius: 4px;
}

/* 已选择预览样式 */
.selected-preview {
  margin-top: 16px;
  padding: 12px;
  background: #f5f5f5;
  border-radius: 8px;
}

.preview-title {
  font-weight: 500;
  margin-bottom: 12px;
  color: #333;
}

.preview-items {
  display: flex;
  flex-wrap: wrap;
  gap: 12px;
}

.preview-item {
  display: flex;
  align-items: center;
  gap: 8px;
  padding: 8px 12px;
  background: white;
  border-radius: 6px;
  box-shadow: 0 1px 3px rgba(0,0,0,0.1);
}

.preview-thumb {
  width: 80px;
  height: 40px;
  object-fit: cover;
  border-radius: 4px;
}

/* 金刚位预览样式 */
.diamond-preview {
  flex-direction: column;
}

.diamond-group {
  margin-bottom: 12px;
}

.group-title {
  font-weight: 500;
  margin-bottom: 8px;
  color: #1890ff;
}

.group-items {
  display: flex;
  flex-wrap: wrap;
  gap: 8px;
}

.diamond-item {
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: 8px 12px;
  background: white;
  border-radius: 6px;
  box-shadow: 0 1px 3px rgba(0,0,0,0.1);
  min-width: 60px;
}

.diamond-icon {
  font-size: 24px;
  margin-bottom: 4px;
}

.diamond-title {
  font-size: 12px;
  color: #666;
}

/* 推荐位预览样式 */
.recommend-preview {
  margin-bottom: 16px;
  padding: 12px;
  background: white;
  border-radius: 8px;
}

.recommend-title {
  font-weight: 500;
  margin-bottom: 12px;
  padding-bottom: 8px;
  border-bottom: 1px solid #eee;
}

.recommend-materials,
.feed-materials {
  display: flex;
  flex-wrap: wrap;
  gap: 12px;
}

.material-card {
  width: 120px;
  background: #f9f9f9;
  border-radius: 8px;
  overflow: hidden;
}

.material-cover {
  width: 100%;
  height: 160px;
  object-fit: cover;
}

.material-info {
  padding: 8px;
}

.material-name {
  font-size: 12px;
  color: #333;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}

.material-stats {
  font-size: 11px;
  color: #999;
  margin-top: 4px;
}

.more-hint {
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 12px;
  color: #999;
  font-size: 12px;
}
</style>
