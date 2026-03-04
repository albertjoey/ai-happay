'use client';

import { useState, useEffect } from 'react';
import { formatNumber } from '@/lib/content';
import { toggleLike, toggleCollect, getLikeStatus, getCollectStatus } from '@/lib/interaction';

interface InteractionBarProps {
  contentId: number;
  viewCount: number;
  likeCount: number;
  commentCount: number;
  shareCount: number;
  collectCount: number;
  userId?: number; // 用户ID，未登录时为undefined
  onComment?: () => void;
  onShare?: () => void;
}

export default function InteractionBar({
  contentId,
  viewCount,
  likeCount: initialLikeCount,
  commentCount,
  shareCount,
  collectCount: initialCollectCount,
  userId = 1, // 默认用户ID，实际应从登录状态获取
  onComment,
  onShare,
}: InteractionBarProps) {
  const [liked, setLiked] = useState(false);
  const [collected, setCollected] = useState(false);
  const [likeCount, setLikeCount] = useState(initialLikeCount);
  const [collectCount, setCollectCount] = useState(initialCollectCount);
  const [animating, setAnimating] = useState<'like' | 'collect' | null>(null);
  const [loading, setLoading] = useState(false);

  // 初始化时获取点赞和收藏状态
  useEffect(() => {
    const fetchStatus = async () => {
      try {
        const [likeRes, collectRes] = await Promise.all([
          getLikeStatus(userId, 'material', contentId),
          getCollectStatus(userId, 'material', contentId),
        ]);
        setLiked(likeRes.liked || false);
        setCollected(collectRes.collected || false);
      } catch (error) {
        console.error('获取互动状态失败:', error);
      }
    };
    
    if (userId) {
      fetchStatus();
    }
  }, [userId, contentId]);

  const handleLike = async () => {
    if (animating || loading) return;
    
    setLoading(true);
    setAnimating('like');
    
    try {
      const res = await toggleLike(userId, 'material', contentId);
      setLiked(res.liked);
      setLikeCount(prev => res.liked ? prev + 1 : prev - 1);
    } catch (error) {
      console.error('点赞失败:', error);
    }
    
    setTimeout(() => {
      setAnimating(null);
      setLoading(false);
    }, 300);
  };

  const handleCollect = async () => {
    if (animating || loading) return;
    
    setLoading(true);
    setAnimating('collect');
    
    try {
      const res = await toggleCollect(userId, 'material', contentId);
      setCollected(res.collected);
      setCollectCount(prev => res.collected ? prev + 1 : prev - 1);
    } catch (error) {
      console.error('收藏失败:', error);
    }
    
    setTimeout(() => {
      setAnimating(null);
      setLoading(false);
    }, 300);
  };

  const handleShare = () => {
    // 复制链接到剪贴板
    if (navigator.clipboard) {
      navigator.clipboard.writeText(window.location.href);
      alert('链接已复制到剪贴板');
    }
    onShare?.();
  };

  return (
    <div className="fixed bottom-0 left-0 right-0 bg-white border-t border-gray-200 z-50">
      <div className="flex items-center justify-around py-2 px-4">
        {/* 浏览量 */}
        <div className="flex flex-col items-center text-gray-500">
          <svg className="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
            <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z" />
          </svg>
          <span className="text-xs mt-0.5">{formatNumber(viewCount)}</span>
        </div>

        {/* 点赞 */}
        <button 
          className={`flex flex-col items-center transition-transform ${animating === 'like' ? 'scale-125' : ''} ${liked ? 'text-red-500' : 'text-gray-500'}`}
          onClick={handleLike}
          disabled={loading}
        >
          <svg className="w-6 h-6" fill={liked ? 'currentColor' : 'none'} stroke="currentColor" viewBox="0 0 24 24">
            <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M4.318 6.318a4.5 4.5 0 000 6.364L12 20.364l7.682-7.682a4.5 4.5 0 00-6.364-6.364L12 7.636l-1.318-1.318a4.5 4.5 0 00-6.364 0z" />
          </svg>
          <span className="text-xs mt-0.5">{formatNumber(likeCount)}</span>
        </button>

        {/* 评论 */}
        <button 
          className="flex flex-col items-center text-gray-500"
          onClick={onComment}
        >
          <svg className="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M8 12h.01M12 12h.01M16 12h.01M21 12c0 4.418-4.03 8-9 8a9.863 9.863 0 01-4.255-.949L3 20l1.395-3.72C3.512 15.042 3 13.574 3 12c0-4.418 4.03-8 9-8s9 3.582 9 8z" />
          </svg>
          <span className="text-xs mt-0.5">{formatNumber(commentCount)}</span>
        </button>

        {/* 收藏 */}
        <button 
          className={`flex flex-col items-center transition-transform ${animating === 'collect' ? 'scale-125' : ''} ${collected ? 'text-yellow-500' : 'text-gray-500'}`}
          onClick={handleCollect}
          disabled={loading}
        >
          <svg className="w-6 h-6" fill={collected ? 'currentColor' : 'none'} stroke="currentColor" viewBox="0 0 24 24">
            <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M11.049 2.927c.3-.921 1.603-.921 1.902 0l1.519 4.674a1 1 0 00.95.69h4.915c.969 0 1.371 1.24.588 1.81l-3.976 2.888a1 1 0 00-.363 1.118l1.518 4.674c.3.922-.755 1.688-1.538 1.118l-3.976-2.888a1 1 0 00-1.176 0l-3.976 2.888c-.783.57-1.838-.197-1.538-1.118l1.518-4.674a1 1 0 00-.363-1.118l-3.976-2.888c-.784-.57-.38-1.81.588-1.81h4.914a1 1 0 00.951-.69l1.519-4.674z" />
          </svg>
          <span className="text-xs mt-0.5">{formatNumber(collectCount)}</span>
        </button>

        {/* 分享 */}
        <button 
          className="flex flex-col items-center text-gray-500"
          onClick={handleShare}
        >
          <svg className="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M8.684 13.342C8.886 12.938 9 12.482 9 12c0-.482-.114-.938-.316-1.342m0 2.684a3 3 0 110-2.684m0 2.684l6.632 3.316m-6.632-6l6.632-3.316m0 0a3 3 0 105.367-2.684 3 3 0 00-5.367 2.684zm0 9.316a3 3 0 105.368 2.684 3 3 0 00-5.368-2.684z" />
          </svg>
          <span className="text-xs mt-0.5">{formatNumber(shareCount)}</span>
        </button>
      </div>
    </div>
  );
}
