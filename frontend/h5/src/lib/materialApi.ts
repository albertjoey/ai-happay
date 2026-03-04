import axios from 'axios';

const API_BASE = process.env.NEXT_PUBLIC_API_BASE || 'http://localhost:4004';

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

// 获取物料列表
export async function getMaterialList(params?: {
  page?: number;
  page_size?: number;
  type?: string;
  category?: string;
  keyword?: string;
}): Promise<{ total: number; list: Material[] }> {
  try {
    const response = await axios.get(`${API_BASE}/api/v1/material/list`, {
      params: { ...params, status: 1 },
    });
    return response.data;
  } catch (error) {
    console.error('获取物料列表失败:', error);
    return { total: 0, list: [] };
  }
}

// 根据类型获取物料
export async function getMaterialsByType(type: string, limit: number = 10): Promise<Material[]> {
  const { list } = await getMaterialList({ type, page_size: limit });
  return list;
}

// 获取推荐物料(多类型)
export async function getRecommendedMaterials(limit: number = 20): Promise<Material[]> {
  const { list } = await getMaterialList({ page_size: limit });
  return list;
}
