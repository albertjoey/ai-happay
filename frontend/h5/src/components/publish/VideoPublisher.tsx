'use client';

import { useState } from 'react';

interface VideoPublisherProps {
  type: 'short_video' | 'long_video';
  onSuccess?: () => void;
}

export default function VideoPublisher({ type, onSuccess }: VideoPublisherProps) {
  const [title, setTitle] = useState('');
  const [description, setDescription] = useState('');
  const [coverUrl, setCoverUrl] = useState('');
  const [videoUrl, setVideoUrl] = useState('');
  const [submitting, setSubmitting] = useState(false);

  const isShort = type === 'short_video';

  // 提交发布
  const handleSubmit = async () => {
    if (!title.trim()) {
      alert('请输入标题');
      return;
    }

    setSubmitting(true);
    try {
      const response = await fetch('http://localhost:4004/api/v1/material', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({
          type,
          title,
          description,
          cover_url: coverUrl || `https://picsum.photos/800/450?random=${Date.now()}`,
          content_url: videoUrl || 'https://test-streams.mux.dev/x36xhzz/x36xhzz.m3u8',
          author: '用户发布',
          category: isShort ? '短视频' : '长视频',
          duration: isShort ? 60 : 1800,
        }),
      });

      if (response.ok) {
        alert('发布成功！');
        onSuccess?.();
      } else {
        throw new Error('发布失败');
      }
    } catch (error) {
      alert('发布失败，请重试');
    } finally {
      setSubmitting(false);
    }
  };

  return (
    <div className="p-4">
      {/* 视频上传区域 */}
      <div className="mb-4">
        <label className="block text-sm font-medium text-gray-700 mb-2">
          {isShort ? '短视频' : '视频文件'}
        </label>
        <div className="border-2 border-dashed border-gray-300 rounded-lg p-8 text-center">
          <div className="text-4xl mb-2">🎬</div>
          <p className="text-gray-500 text-sm">点击上传视频</p>
          <p className="text-gray-400 text-xs mt-1">
            {isShort ? '支持竖屏视频，时长60秒以内' : '支持MP4格式，最大2GB'}
          </p>
          <button className="mt-4 px-4 py-2 bg-blue-500 text-white rounded-lg text-sm">
            选择视频
          </button>
        </div>
        {videoUrl && (
          <div className="mt-2 p-2 bg-gray-100 rounded text-sm text-gray-600">
            已选择视频文件
          </div>
        )}
      </div>

      {/* 封面设置 */}
      <div className="mb-4">
        <label className="block text-sm font-medium text-gray-700 mb-2">封面</label>
        <div className="flex gap-2">
          <div className="w-24 h-16 bg-gray-100 rounded flex items-center justify-center text-gray-400">
            {coverUrl ? (
              <img src={coverUrl} alt="" className="w-full h-full object-cover rounded" />
            ) : (
              <span className="text-xs">选择封面</span>
            )}
          </div>
          <button
            onClick={() => setCoverUrl(`https://picsum.photos/800/450?random=${Date.now()}`)}
            className="px-3 py-1 border border-gray-300 rounded text-sm"
          >
            自动生成
          </button>
        </div>
      </div>

      {/* 标题输入 */}
      <div className="mb-4">
        <label className="block text-sm font-medium text-gray-700 mb-1">标题</label>
        <input
          type="text"
          value={title}
          onChange={(e) => setTitle(e.target.value)}
          placeholder="请输入标题"
          className="w-full px-4 py-3 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
          maxLength={50}
        />
      </div>

      {/* 描述输入 */}
      <div className="mb-6">
        <label className="block text-sm font-medium text-gray-700 mb-1">描述</label>
        <textarea
          value={description}
          onChange={(e) => setDescription(e.target.value)}
          placeholder="介绍一下你的视频..."
          className="w-full px-4 py-3 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 min-h-[80px]"
          maxLength={500}
        />
      </div>

      {/* 发布按钮 */}
      <button
        onClick={handleSubmit}
        disabled={submitting}
        className="w-full py-3 bg-blue-500 text-white rounded-lg font-medium disabled:bg-gray-300"
      >
        {submitting ? '发布中...' : '发布'}
      </button>
    </div>
  );
}
