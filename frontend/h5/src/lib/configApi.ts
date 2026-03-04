import axios from 'axios';

const API_BASE = process.env.NEXT_PUBLIC_API_BASE || 'http://localhost:4004';

// 金刚位类型
export interface Diamond {
  id: number;
  channel_id: number;
  group_id: number;
  sort: number;
  title: string;
  icon: string;
  link_type: string;
  link_value: string;
  status: number;
  description: string;
  material_id: number;
  material?: {
    id: number;
    title: string;
    type: string;
    cover_url: string;
    author: string;
    view_count: number;
  };
}

// 推荐位类型
export interface Recommend {
  id: number;
  channel_id: number;
  title: string;
  display_type: string;
  source_type: string;
  content_ids: number[];
  filter_rule: Record<string, any>;
  sort: number;
  status: number;
  description: string;
}

// Feed流配置类型
export interface FeedConfig {
  id: number;
  channel_id: number;
  title: string;
  layout_type: string;
  content_strategy: string;
  content_ids: number[];
  filter_rule: Record<string, any>;
  sort: number;
  status: number;
  description: string;
  materials?: any[]; // 物料列表
}

// 获取金刚位列表
export async function getDiamondList(channelId: number): Promise<Diamond[]> {
  try {
    const response = await axios.get(`${API_BASE}/api/v1/diamond/list`, {
      params: { channel_id: channelId, status: 1 },
    });
    return response.data.list || [];
  } catch (error) {
    console.error('获取金刚位列表失败:', error);
    return [];
  }
}

// 获取推荐位列表
export async function getRecommendList(channelId: number): Promise<Recommend[]> {
  try {
    const response = await axios.get(`${API_BASE}/api/v1/recommend/list`, {
      params: { channel_id: channelId, status: 1 },
    });
    return response.data.list || [];
  } catch (error) {
    console.error('获取推荐位列表失败:', error);
    return [];
  }
}

// 获取Feed流配置列表
export async function getFeedConfigList(channelId: number): Promise<FeedConfig[]> {
  try {
    const response = await axios.get(`${API_BASE}/api/v1/feed-config/list`, {
      params: { channel_id: channelId, status: 1 },
    });
    return response.data.list || [];
  } catch (error) {
    console.error('获取Feed流配置列表失败:', error);
    return [];
  }
}

// 物料类型
export interface Material {
  id: number;
  title: string;
  subtitle: string;
  type: string;
  cover_url: string;
  content_url: string;
  description: string;
  author: string;
  category: string;
  view_count: number;
  like_count: number;
  comment_count: number;
  share_count: number;
  collect_count: number;
  duration: number;
  word_count: number;
  chapter_count: number;
  status: number;
  sort: number;
}

// 广告位类型
export interface AdSlot {
  id: number;
  channel_id: number;
  name: string;
  insert_type: string;
  insert_rule: Record<string, any>;
  ad_type: string;
  ad_content: Record<string, any>;
  link_url: string;
  status: number;
  sort: number;
  description: string;
}

// 获取广告位列表
export async function getAdSlotList(channelId: number): Promise<AdSlot[]> {
  try {
    const response = await axios.get(`${API_BASE}/api/v1/ad-slot/list`, {
      params: { channel_id: channelId, status: 1 },
    });
    return response.data.list || [];
  } catch (error) {
    console.error('获取广告位列表失败:', error);
    return [];
  }
}
