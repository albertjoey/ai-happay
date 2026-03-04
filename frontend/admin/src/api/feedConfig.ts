import request from './request';

// Feed流配置类型定义
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
}

export interface FeedConfigListResponse {
  list: FeedConfig[];
}

// 获取Feed流配置列表
export const getFeedConfigList = (params: {
  channel_id: number;
  status?: number;
}) => {
  return request.get<any, FeedConfigListResponse>('/feed-config/list', { params });
};

// 创建Feed流配置
export const createFeedConfig = (data: Partial<FeedConfig>) => {
  return request.post('/feed-config', data);
};

// 更新Feed流配置
export const updateFeedConfig = (id: number, data: Partial<FeedConfig>) => {
  return request.put(`/feed-config/${id}`, data);
};

// 删除Feed流配置
export const deleteFeedConfig = (id: number) => {
  return request.delete(`/feed-config/${id}`);
};
