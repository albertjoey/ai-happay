import request from './request';

export interface Chapter {
  id: number;
  material_id: number;
  chapter_type: string;
  title: string;
  content?: string;
  images?: string[];
  video_url?: string;
  word_count: number;
  duration: number;
  sort: number;
  is_free: number;
  price: number;
}

export interface ChapterListResponse {
  list: Chapter[];
  total: number;
}

export interface ChapterCreateRequest {
  material_id: number;
  chapter_type: string;
  title: string;
  content?: string;
  images?: string[];
  video_url?: string;
  word_count?: number;
  duration?: number;
  sort?: number;
  is_free?: number;
  price?: number;
}

export interface ChapterUpdateRequest {
  chapter_type: string;
  title: string;
  content?: string;
  images?: string[];
  video_url?: string;
  word_count?: number;
  duration?: number;
  sort?: number;
  is_free?: number;
  price?: number;
}

// 获取章节列表
export function getChapterList(materialId: number): Promise<ChapterListResponse> {
  return request.get(`/chapter/list/${materialId}`);
}

// 获取章节详情
export function getChapterDetail(id: number): Promise<Chapter> {
  return request.get(`/chapter/${id}`);
}

// 创建章节
export function createChapter(data: ChapterCreateRequest): Promise<{ id: number; success: boolean }> {
  return request.post('/chapter', {
    ...data,
    word_count: data.word_count || 0,
    duration: data.duration || 0,
    sort: data.sort || 0,
    is_free: data.is_free ?? 1,
    price: data.price || 0,
  });
}

// 更新章节
export function updateChapter(id: number, data: ChapterUpdateRequest): Promise<{ success: boolean }> {
  return request.put(`/chapter/${id}`, {
    ...data,
    word_count: data.word_count || 0,
    duration: data.duration || 0,
    sort: data.sort || 0,
    is_free: data.is_free ?? 1,
    price: data.price || 0,
  });
}

// 删除章节
export function deleteChapter(id: number): Promise<{ success: boolean }> {
  return request.delete(`/chapter/${id}`);
}
