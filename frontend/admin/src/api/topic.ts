import request from './request';

// 话题类型定义
export interface Topic {
  id: number;
  name: string;
  description: string;
  cover: string;
  status: number;
  sort: number;
  created_at: string;
}

export interface TopicListResponse {
  total: number;
  list: Topic[];
}

// 获取话题列表
export const getTopicList = (params: {
  page?: number;
  page_size?: number;
  name?: string;
  status?: number;
}) => {
  return request.get<any, TopicListResponse>('/topic/list', { params });
};

// 创建话题
export const createTopic = (data: Partial<Topic>) => {
  return request.post('/topic', data);
};

// 更新话题
export const updateTopic = (id: number, data: Partial<Topic>) => {
  return request.put(`/topic/${id}`, data);
};

// 删除话题
export const deleteTopic = (id: number) => {
  return request.delete(`/topic/${id}`);
};
