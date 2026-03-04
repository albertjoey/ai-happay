import axios from 'axios';

const API_BASE = process.env.NEXT_PUBLIC_API_BASE || 'http://localhost:4004';

// 创建axios实例
const request = axios.create({
  baseURL: `${API_BASE}/api/v1`,
  timeout: 10000,
});

// 互动类型
export type TargetType = 'material' | 'chapter' | 'comment';

// ========== 点赞系统 ==========

// 点赞/取消点赞
export const toggleLike = async (userId: number, targetType: TargetType, targetId: number) => {
  const response = await request.post('/like', {
    user_id: userId,
    target_type: targetType,
    target_id: targetId,
  });
  return response.data;
};

// 获取点赞状态
export const getLikeStatus = async (userId: number, targetType: TargetType, targetId: number) => {
  const response = await request.get('/like/status', {
    params: {
      user_id: userId,
      target_type: targetType,
      target_id: targetId,
    },
  });
  return response.data;
};

// ========== 收藏系统 ==========

// 收藏/取消收藏
export const toggleCollect = async (userId: number, targetType: TargetType, targetId: number) => {
  const response = await request.post('/collect', {
    user_id: userId,
    target_type: targetType,
    target_id: targetId,
  });
  return response.data;
};

// 获取收藏状态
export const getCollectStatus = async (userId: number, targetType: TargetType, targetId: number) => {
  const response = await request.get('/collect/status', {
    params: {
      user_id: userId,
      target_type: targetType,
      target_id: targetId,
    },
  });
  return response.data;
};

// 获取收藏列表
export const getCollectList = async (userId: number, page = 1, pageSize = 20) => {
  const response = await request.get('/collect/list', {
    params: {
      user_id: userId,
      page,
      page_size: pageSize,
    },
  });
  return response.data;
};

// ========== 评论系统 ==========

export interface Comment {
  id: number;
  user_id: number;
  content: string;
  like_count: number;
  created_at: string;
}

// 创建评论
export const createComment = async (userId: number, targetType: TargetType, targetId: number, content: string) => {
  const response = await request.post('/comment', {
    user_id: userId,
    target_type: targetType,
    target_id: targetId,
    content,
  });
  return response.data;
};

// 获取评论列表
export const getCommentList = async (targetType: TargetType, targetId: number, page = 1, pageSize = 20) => {
  const response = await request.get<{ total: number; list: Comment[] }>('/comment/list', {
    params: {
      target_type: targetType,
      target_id: targetId,
      page,
      page_size: pageSize,
    },
  });
  return response.data;
};

// 删除评论
export const deleteComment = async (userId: number, commentId: number) => {
  const response = await request.delete(`/comment/${commentId}?user_id=${userId}`);
  return response.data;
};
