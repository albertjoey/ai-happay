import request from './request';

// 广告位类型定义
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

export interface AdSlotListResponse {
  list: AdSlot[];
}

// 获取广告位列表
export const getAdSlotList = (params: {
  channel_id: number;
  status?: number;
}) => {
  return request.get<any, AdSlotListResponse>('/ad-slot/list', { params });
};

// 创建广告位
export const createAdSlot = (data: Partial<AdSlot>) => {
  return request.post('/ad-slot', data);
};

// 更新广告位
export const updateAdSlot = (id: number, data: Partial<AdSlot>) => {
  return request.put(`/ad-slot/${id}`, data);
};

// 删除广告位
export const deleteAdSlot = (id: number) => {
  return request.delete(`/ad-slot/${id}`);
};
