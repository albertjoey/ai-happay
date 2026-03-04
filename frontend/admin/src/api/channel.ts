import request from './request';

// 频道类型定义
export interface Channel {
  id: number;
  name: string;
  icon: string;
  description: string;
  sort: number;
  status: number;
  tenant_id: number;
  created_at: string;
  updated_at: string;
}

export interface ChannelConfig {
  channel_id: number;
  content_type: {
    video?: boolean;
    article?: boolean;
    image?: boolean;
  };
  display_mode: string;
  custom_data: Record<string, any>;
}

export interface ChannelListResponse {
  list: Channel[];
  total: number;
}

export interface ChannelSortItem {
  id: number;
  sort: number;
}

// 获取频道列表
export const getChannelList = (params: {
  page?: number;
  page_size?: number;
}) => {
  return request.get<any, ChannelListResponse>('/channel/list', { params });
};

// 创建频道
export const createChannel = (data: Partial<Channel>) => {
  return request.post('/channel', data);
};

// 更新频道
export const updateChannel = (id: number, data: Partial<Channel>) => {
  return request.put(`/channel/${id}`, data);
};

// 删除频道
export const deleteChannel = (id: number) => {
  return request.delete(`/channel/${id}`);
};

// 频道排序
export const sortChannels = (data: ChannelSortItem[]) => {
  return request.post('/channel/sort', { items: data });
};

// 获取频道配置
export const getChannelConfig = (id: number) => {
  return request.get<any, ChannelConfig>(`/channel/config/${id}`);
};

// 更新频道配置
export const updateChannelConfig = (id: number, data: Partial<ChannelConfig>) => {
  return request.put(`/channel/config/${id}`, data);
};
