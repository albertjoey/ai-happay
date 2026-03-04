import request from './request';

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
  tags: string[];
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

// 物料列表响应
export interface MaterialListResponse {
  total: number;
  list: Material[];
}

// 获取物料列表
export const getMaterialList = (params: {
  page?: number;
  page_size?: number;
  type?: string;
  category?: string;
  keyword?: string;
  status?: number;
}) => {
  return request.get<any, MaterialListResponse>('/material/list', { params });
};

// 创建物料
export const createMaterial = (data: Partial<Material>) => {
  return request.post('/material', data);
};

// 更新物料
export const updateMaterial = (id: number, data: Partial<Material>) => {
  return request.put(`/material/${id}`, data);
};

// 删除物料
export const deleteMaterial = (id: number) => {
  return request.delete(`/material/${id}`);
};
