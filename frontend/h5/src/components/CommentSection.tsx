'use client';

import { useState, useEffect } from 'react';
import { getCommentList, createComment, deleteComment, Comment } from '@/lib/interaction';

interface CommentSectionProps {
  contentId: number;
  userId?: number;
  isOpen: boolean;
  onClose: () => void;
}

export default function CommentSection({
  contentId,
  userId = 1,
  isOpen,
  onClose,
}: CommentSectionProps) {
  const [comments, setComments] = useState<Comment[]>([]);
  const [total, setTotal] = useState(0);
  const [page, setPage] = useState(1);
  const [loading, setLoading] = useState(false);
  const [inputValue, setInputValue] = useState('');
  const [submitting, setSubmitting] = useState(false);

  // 加载评论列表
  useEffect(() => {
    if (isOpen) {
      fetchComments();
    }
  }, [isOpen, contentId, page]);

  const fetchComments = async () => {
    setLoading(true);
    try {
      const res = await getCommentList('material', contentId, page, 20);
      setComments(res.list || []);
      setTotal(res.total || 0);
    } catch (error) {
      console.error('获取评论失败:', error);
    }
    setLoading(false);
  };

  const handleSubmit = async () => {
    if (!inputValue.trim() || submitting) return;
    
    setSubmitting(true);
    try {
      await createComment(userId, 'material', contentId, inputValue.trim());
      setInputValue('');
      setPage(1);
      fetchComments();
    } catch (error) {
      console.error('发表评论失败:', error);
      alert('发表失败，请重试');
    }
    setSubmitting(false);
  };

  const handleDelete = async (commentId: number) => {
    if (!confirm('确定删除这条评论吗？')) return;
    
    try {
      await deleteComment(userId, commentId);
      fetchComments();
    } catch (error) {
      console.error('删除评论失败:', error);
    }
  };

  if (!isOpen) return null;

  return (
    <div className="fixed inset-0 z-50 bg-black/50" onClick={onClose}>
      <div 
        className="absolute bottom-0 left-0 right-0 bg-white rounded-t-2xl max-h-[70vh] flex flex-col"
        onClick={e => e.stopPropagation()}
      >
        {/* 头部 */}
        <div className="flex items-center justify-between p-4 border-b">
          <h3 className="text-lg font-medium">评论 ({total})</h3>
          <button onClick={onClose} className="text-gray-500">
            <svg className="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M6 18L18 6M6 6l12 12" />
            </svg>
          </button>
        </div>

        {/* 评论列表 */}
        <div className="flex-1 overflow-y-auto p-4 space-y-4">
          {loading ? (
            <div className="text-center text-gray-500 py-8">加载中...</div>
          ) : comments.length === 0 ? (
            <div className="text-center text-gray-500 py-8">暂无评论，快来抢沙发吧~</div>
          ) : (
            comments.map(comment => (
              <div key={comment.id} className="flex gap-3">
                <div className="w-10 h-10 rounded-full bg-gray-200 flex-shrink-0 flex items-center justify-center">
                  <span className="text-gray-500 text-sm">{comment.user_id}</span>
                </div>
                <div className="flex-1">
                  <div className="flex items-center gap-2">
                    <span className="text-sm font-medium text-gray-700">用户{comment.user_id}</span>
                    <span className="text-xs text-gray-400">{comment.created_at}</span>
                  </div>
                  <p className="text-gray-800 mt-1">{comment.content}</p>
                  <div className="flex items-center gap-4 mt-2 text-xs text-gray-500">
                    <span>❤️ {comment.like_count}</span>
                    {comment.user_id === userId && (
                      <button 
                        onClick={() => handleDelete(comment.id)}
                        className="text-red-500"
                      >
                        删除
                      </button>
                    )}
                  </div>
                </div>
              </div>
            ))
          )}
          
          {total > 20 && page * 20 < total && (
            <button 
              onClick={() => setPage(p => p + 1)}
              className="w-full py-2 text-center text-blue-500"
            >
              加载更多
            </button>
          )}
        </div>

        {/* 输入框 */}
        <div className="p-4 border-t flex gap-2">
          <input
            type="text"
            value={inputValue}
            onChange={e => setInputValue(e.target.value)}
            placeholder="说点什么..."
            className="flex-1 px-4 py-2 border rounded-full focus:outline-none focus:border-blue-500"
            maxLength={500}
          />
          <button
            onClick={handleSubmit}
            disabled={!inputValue.trim() || submitting}
            className={`px-6 py-2 rounded-full ${
              inputValue.trim() && !submitting
                ? 'bg-blue-500 text-white'
                : 'bg-gray-200 text-gray-400'
            }`}
          >
            发送
          </button>
        </div>
      </div>
    </div>
  );
}
