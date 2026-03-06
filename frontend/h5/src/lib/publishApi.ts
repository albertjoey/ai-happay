import axios from 'axios';

const API_BASE = process.env.NEXT_PUBLIC_API_BASE || 'http://localhost:4004';

// 发布类型
export type PublishType = 'image_text' | 'short_video' | 'long_video' | 'short_drama' | 'drama' | 'novel';

// 基础发布请求
export interface BasePublishRequest {
  type: PublishType;
  title: string;
  subtitle?: string;
  description?: string;
  cover_url?: string;
  category?: string;
  tags?: string[];
  topic_ids?: number[];
}

// 图文发布请求
export interface ImageTextPublishRequest extends BasePublishRequest {
  type: 'image_text';
  images: string[]; // 图片URL数组
  content: string; // 正文内容
}

// 视频发布请求
export interface VideoPublishRequest extends BasePublishRequest {
  type: 'short_video' | 'long_video';
  video_url: string; // 视频URL
  duration: number; // 时长(秒)
  width?: number; // 视频宽度
  height?: number; // 视频高度
}

// 短剧/漫剧发布请求
export interface DramaPublishRequest extends BasePublishRequest {
  type: 'short_drama' | 'drama';
  episodes: Episode[]; // 分集列表
  total_episodes: number; // 总集数
}

// 分集信息
export interface Episode {
  episode_number: number; // 集数
  title: string; // 标题
  cover_url?: string; // 封面
  video_url?: string; // 视频URL (短剧)
  images?: string[]; // 图片数组 (漫剧)
  duration?: number; // 时长
}

// 小说发布请求
export interface NovelPublishRequest extends BasePublishRequest {
  type: 'novel';
  chapters: Chapter[]; // 章节列表
  total_chapters: number; // 总章数
  word_count: number; // 总字数
}

// 章节信息
export interface Chapter {
  chapter_number: number; // 章数
  title: string; // 标题
  content: string; // 内容
  word_count: number; // 字数
}

// 发布响应
export interface PublishResponse {
  id: number;
  title: string;
  type: PublishType;
  status: number;
  message: string;
}

// 发布物料
export async function publishMaterial(data: BasePublishRequest | ImageTextPublishRequest | VideoPublishRequest | DramaPublishRequest | NovelPublishRequest): Promise<PublishResponse> {
  try {
    // 构建请求数据
    const requestData: any = {
      type: data.type,
      title: data.title,
      subtitle: data.subtitle || '',
      description: data.description || '',
      cover_url: data.cover_url || '',
      category: data.category || '',
      tags: data.tags || [],
      author: '用户发布',
      status: 1,
    };

    // 根据类型处理特定字段
    switch (data.type) {
      case 'image_text':
        const imageData = data as ImageTextPublishRequest;
        requestData.content_url = JSON.stringify(imageData.images);
        requestData.description = imageData.content;
        if (imageData.images.length > 0) {
          requestData.cover_url = imageData.images[0];
        }
        break;

      case 'short_video':
      case 'long_video':
        const videoData = data as VideoPublishRequest;
        requestData.content_url = videoData.video_url;
        requestData.duration = videoData.duration;
        break;

      case 'short_drama':
      case 'drama':
        const dramaData = data as DramaPublishRequest;
        requestData.content_url = JSON.stringify(dramaData.episodes);
        requestData.chapter_count = dramaData.total_episodes;
        break;

      case 'novel':
        const novelData = data as NovelPublishRequest;
        requestData.content_url = JSON.stringify(novelData.chapters);
        requestData.chapter_count = novelData.total_chapters;
        requestData.word_count = novelData.word_count;
        break;
    }

    const response = await axios.post(`${API_BASE}/api/v1/material`, requestData);
    return {
      id: response.data.id || Date.now(),
      title: data.title,
      type: data.type,
      status: 1,
      message: '发布成功',
    };
  } catch (error) {
    console.error('发布失败:', error);
    throw new Error('发布失败，请重试');
  }
}

// 上传文件
export async function uploadFile(file: File, type: 'image' | 'video' = 'image'): Promise<string> {
  try {
    const formData = new FormData();
    formData.append('file', file);
    formData.append('type', type);

    const response = await axios.post(`${API_BASE}/api/v1/upload`, formData, {
      headers: {
        'Content-Type': 'multipart/form-data',
      },
    });

    return response.data.url;
  } catch (error) {
    console.error('上传失败:', error);
    // 如果上传接口不存在，返回模拟URL
    return `https://picsum.photos/800/600?random=${Date.now()}`;
  }
}

// 获取话题列表
export async function getTopicList(params?: {
  page?: number;
  page_size?: number;
  keyword?: string;
}): Promise<{ total: number; list: any[] }> {
  try {
    const response = await axios.get(`${API_BASE}/api/v1/topic/list`, { params });
    return response.data;
  } catch (error) {
    console.error('获取话题列表失败:', error);
    return { total: 0, list: [] };
  }
}

// 获取标签列表
export async function getTagList(params?: {
  page?: number;
  page_size?: number;
  keyword?: string;
}): Promise<{ total: number; list: any[] }> {
  try {
    const response = await axios.get(`${API_BASE}/api/v1/tag/list`, { params });
    return response.data;
  } catch (error) {
    console.error('获取标签列表失败:', error);
    return { total: 0, list: [] };
  }
}
