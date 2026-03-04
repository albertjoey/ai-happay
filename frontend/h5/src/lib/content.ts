import axios from 'axios';

const API_BASE = process.env.NEXT_PUBLIC_API_BASE || 'http://localhost:4004';

export interface Content {
  id: number;
  title: string;
  subtitle: string;
  type: 'image_text' | 'long_video' | 'short_video' | 'drama' | 'short_drama' | 'novel';
  cover_url: string;
  content_url: string;
  description: string;
  author: string;
  tags: string[] | null;
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
  created_at?: string;
}

export interface ContentDetail extends Content {
  content?: string; // 正文内容
  images?: string[]; // 图文图片列表
  episodes?: Episode[]; // 剧集列表
  chapters?: Chapter[]; // 章节列表
}

export interface Episode {
  id: number;
  title: string;
  cover_url: string;
  video_url: string;
  duration: number;
  sort: number;
  images?: string[]; // 漫剧图片列表
}

export interface Chapter {
  id: number;
  material_id: number;
  title: string;
  word_count: number;
  sort: number;
  is_free: number;
  price: number;
}

export interface ChapterDetail extends Chapter {
  content: string;
  prev_id: number;
  next_id: number;
}

// 获取章节列表
export async function getChapterList(materialId: number): Promise<Chapter[]> {
  try {
    const response = await axios.get(`${API_BASE}/api/v1/chapter/list/${materialId}`);
    return response.data.list || [];
  } catch (error) {
    console.error('获取章节列表失败:', error);
    return [];
  }
}

// 获取章节详情
export async function getChapterDetail(chapterId: number): Promise<ChapterDetail | null> {
  try {
    const response = await axios.get(`${API_BASE}/api/v1/chapter/${chapterId}`);
    return response.data;
  } catch (error) {
    console.error('获取章节详情失败:', error);
    return null;
  }
}

// 获取内容详情
export async function getContentDetail(id: number): Promise<ContentDetail> {
  try {
    const response = await axios.get(`${API_BASE}/api/v1/material/${id}`);
    return response.data;
  } catch (error) {
    console.error('获取内容详情失败:', error);
    throw error;
  }
}

// 获取相关推荐
export async function getRelatedContent(id: number, limit: number = 10): Promise<Content[]> {
  try {
    const response = await axios.get(`${API_BASE}/api/v1/material/list`, {
      params: {
        page: 1,
        page_size: limit,
        exclude_id: id,
      },
    });
    return response.data.list || [];
  } catch (error) {
    console.error('获取相关推荐失败:', error);
    return [];
  }
}

// 获取章节内容（小说）
export async function getChapterContent(contentId: number, chapterId: number): Promise<string> {
  try {
    const response = await axios.get(`${API_BASE}/api/v1/material/${contentId}/chapter/${chapterId}`);
    return response.data.content || '';
  } catch (error) {
    console.error('获取章节内容失败:', error);
    return '';
  }
}

// 记录浏览
export async function recordView(contentId: number): Promise<void> {
  try {
    await axios.post(`${API_BASE}/api/v1/material/${contentId}/view`);
  } catch (error) {
    console.error('记录浏览失败:', error);
  }
}

// 点赞
export async function likeContent(contentId: number): Promise<{ liked: boolean; like_count: number }> {
  try {
    const response = await axios.post(`${API_BASE}/api/v1/material/${contentId}/like`);
    return response.data;
  } catch (error) {
    console.error('点赞失败:', error);
    throw error;
  }
}

// 收藏
export async function collectContent(contentId: number): Promise<{ collected: boolean; collect_count: number }> {
  try {
    const response = await axios.post(`${API_BASE}/api/v1/material/${contentId}/collect`);
    return response.data;
  } catch (error) {
    console.error('收藏失败:', error);
    throw error;
  }
}

// 内容类型映射
export const contentTypeMap: Record<string, { name: string; icon: string }> = {
  image_text: { name: '图文', icon: '📄' },
  long_video: { name: '长视频', icon: '🎬' },
  short_video: { name: '短视频', icon: '📱' },
  drama: { name: '漫剧', icon: '📚' },
  short_drama: { name: '短剧', icon: '🎭' },
  novel: { name: '小说', icon: '📖' },
};

// 格式化数字
export function formatNumber(num: number): string {
  if (num >= 10000) {
    return (num / 10000).toFixed(1) + 'w';
  }
  if (num >= 1000) {
    return (num / 1000).toFixed(1) + 'k';
  }
  return num.toString();
}

// 格式化时长
export function formatDuration(seconds: number): string {
  const hours = Math.floor(seconds / 3600);
  const minutes = Math.floor((seconds % 3600) / 60);
  const secs = seconds % 60;
  
  if (hours > 0) {
    return `${hours}:${String(minutes).padStart(2, '0')}:${String(secs).padStart(2, '0')}`;
  }
  return `${minutes}:${String(secs).padStart(2, '0')}`;
}
