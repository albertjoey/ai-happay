import axios from 'axios';

const API_BASE = process.env.NEXT_PUBLIC_API_BASE || 'http://localhost:4004';

export interface Channel {
  id: number;
  name: string;
  code: string;
  description: string;
  icon: string;
  status: number;
  sort: number;
}

// 页面配置相关类型
export interface BannerConfig {
  enabled: boolean;
  banner_ids?: number[];
}

export interface DiamondConfig {
  enabled: boolean;
  group_ids?: number[];
}

export interface RecommendConfig {
  id: number;
  title?: string;
  sort?: number;
}

export interface FeedConfigItem {
  enabled: boolean;
  feed_id?: number;
  auto_load?: boolean;
  show_title?: boolean;
}

export interface ChannelPageConfig {
  banner?: BannerConfig;
  diamond?: DiamondConfig;
  recommends?: RecommendConfig[];
  feed?: FeedConfigItem;
}

export interface ChannelConfig {
  channel_id: number;
  content_type: Record<string, boolean>;
  display_mode: string;
  custom_data: Record<string, any>;
  page_config?: ChannelPageConfig;
}

export interface ChannelListResponse {
  total: number;
  list: Channel[];
}

// 获取频道列表
export async function getChannelList(): Promise<ChannelListResponse> {
  try {
    const response = await axios.get(`${API_BASE}/api/v1/channel/list`);
    // 从channel/list接口获取channels
    const channels = response.data.list || [];
    return {
      total: channels.length,
      list: channels.map((ch: any) => ({
        id: ch.id,
        name: ch.name,
        code: ch.code,
        description: ch.description || ch.code,
        icon: ch.icon || '',
        status: ch.status || 1,
        sort: ch.sort || ch.id,
      }))
    };
  } catch (error) {
    console.error('获取频道列表失败:', error);
    return { total: 0, list: [] };
  }
}

// 获取频道配置
export async function getChannelConfig(channelId: number): Promise<ChannelConfig> {
  try {
    const response = await axios.get(`${API_BASE}/api/v1/channel/config/${channelId}`);
    return response.data;
  } catch (error) {
    console.error('获取频道配置失败:', error);
    // 返回默认配置
    return {
      channel_id: channelId,
      content_type: { video: true, image: true, article: true },
      display_mode: 'default',
      custom_data: {},
      page_config: {
        banner: { enabled: true },
        diamond: { enabled: true },
        recommends: [],
        feed: { enabled: true, auto_load: true, show_title: true },
      },
    };
  }
}
