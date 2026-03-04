import request from './request';

// 金刚位类型定义
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
}

export interface DiamondListResponse {
  list: Diamond[];
}

// 获取金刚位列表
export const getDiamondList = (params: {
  channel_id: number;
  group_id?: number;
  status?: number;
}) => {
  return request.get<any, DiamondListResponse>('/diamond/list', { params });
};

// 创建金刚位
export const createDiamond = (data: Partial<Diamond>) => {
  return request.post('/diamond', data);
};

// 更新金刚位
export const updateDiamond = (id: number, data: Partial<Diamond>) => {
  return request.put(`/diamond/${id}`, data);
};

// 删除金刚位
export const deleteDiamond = (id: number) => {
  return request.delete(`/diamond/${id}`);
};
