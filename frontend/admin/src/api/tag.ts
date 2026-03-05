import request from './request';

// 标签类型定义
export interface Tag {
  id: number;
  name: string;
  type: number;
  status: number;
  sort: number;
  created_at: string;
}

export interface TagListResponse {
  total: number;
  list: Tag[];
}

// 获取标签列表
export const getTagList = (params: {
  page?: number;
  page_size?: number;
  name?: string;
  type?: number;
  status?: number;
}) => {
  return request.get<any, TagListResponse>('/tag/list', { params });
};

// 创建标签
export const createTag = (data: Partial<Tag>) => {
  return request.post('/tag', data);
};

// 更新标签
export const updateTag = (id: number, data: Partial<Tag>) => {
  return request.put(`/tag/${id}`, data);
};

// 删除标签
export const deleteTag = (id: number) => {
  return request.delete(`/tag/${id}`);
};
