import request from './request';

// 推荐位类型定义
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

export interface RecommendListResponse {
  list: Recommend[];
}

// 获取推荐位列表
export const getRecommendList = (params: {
  channel_id: number;
  status?: number;
}) => {
  return request.get<any, RecommendListResponse>('/recommend/list', { params });
};

// 创建推荐位
export const createRecommend = (data: Partial<Recommend>) => {
  return request.post('/recommend', data);
};

// 更新推荐位
export const updateRecommend = (id: number, data: Partial<Recommend>) => {
  return request.put(`/recommend/${id}`, data);
};

// 删除推荐位
export const deleteRecommend = (id: number) => {
  return request.delete(`/recommend/${id}`);
};
