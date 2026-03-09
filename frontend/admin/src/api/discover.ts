import axios from 'axios';

// 发现页API使用固定端口4005
const DISCOVER_API_BASE = 'http://localhost:4005/api/v1';

// 发现页配置
export interface DiscoverConfig {
  id: number;
  module: string;
  title: string;
  enabled: boolean;
  sort_order: number;
  created_at: string;
  updated_at: string;
}

// 发现页内容项
export interface DiscoverItem {
  id: number;
  config_id: number;
  module: string;
  item_type: string;
  item_id: number;
  title: string;
  cover_url: string;
  color: string;
  author: string;
  description: string;
  extra_data: string;
  views: number;
  likes: number;
  comments: number;
  fans: number;
  count: number;
  sort_order: number;
  is_enabled: boolean;
  created_at: string;
  updated_at: string;
}

// 获取发现页配置列表
export function getDiscoverConfigList() {
  return axios.get(`${DISCOVER_API_BASE}/discover/config`);
}

// 更新发现页配置
export function updateDiscoverConfig(data: Partial<DiscoverConfig>) {
  return axios.put(`${DISCOVER_API_BASE}/discover/config`, data);
}

// 获取发现页内容列表
export function getDiscoverItemList(params: { module?: string; page?: number; page_size?: number }) {
  return axios.get(`${DISCOVER_API_BASE}/discover/items`, { params });
}

// 创建发现页内容
export function createDiscoverItem(data: Partial<DiscoverItem>) {
  return axios.post(`${DISCOVER_API_BASE}/discover/items`, data);
}

// 更新发现页内容
export function updateDiscoverItem(data: Partial<DiscoverItem>) {
  return axios.put(`${DISCOVER_API_BASE}/discover/items`, data);
}

// 删除发现页内容
export function deleteDiscoverItem(id: number) {
  return axios.delete(`${DISCOVER_API_BASE}/discover/items/${id}`);
}
