'use client';

import { useState } from 'react';
import { useRouter } from 'next/navigation';

interface ImageTextPublisherProps {
  onSuccess?: () => void;
}

export default function ImageTextPublisher({ onSuccess }: ImageTextPublisherProps) {
  const router = useRouter();
  const [title, setTitle] = useState('');
  const [content, setContent] = useState('');
  const [images, setImages] = useState<string[]>([]);
  const [submitting, setSubmitting] = useState(false);

  // 添加图片
  const handleAddImage = () => {
    if (images.length >= 9) {
      alert('最多只能添加9张图片');
      return;
    }
    // 模拟添加图片
    const newImage = `https://picsum.photos/800/600?random=${Date.now()}`;
    setImages([...images, newImage]);
  };

  // 删除图片
  const handleRemoveImage = (index: number) => {
    setImages(images.filter((_, i) => i !== index));
  };

  // 提交发布
  const handleSubmit = async () => {
    if (!title.trim()) {
      alert('请输入标题');
      return;
    }
    if (images.length === 0) {
      alert('请至少添加一张图片');
      return;
    }

    setSubmitting(true);
    try {
      // 调用发布API
      const response = await fetch('http://localhost:4004/api/v1/material', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({
          type: 'image_text',
          title,
          description: content,
          cover_url: images[0],
          content_url: JSON.stringify(images),
          author: '用户发布',
          category: '图文',
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
        <div className="text-right text-xs text-gray-400 mt-1">{title.length}/50</div>
      </div>

      {/* 内容输入 */}
      <div className="mb-4">
        <label className="block text-sm font-medium text-gray-700 mb-1">正文</label>
        <textarea
          value={content}
          onChange={(e) => setContent(e.target.value)}
          placeholder="分享你的故事..."
          className="w-full px-4 py-3 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 min-h-[120px]"
          maxLength={2000}
        />
        <div className="text-right text-xs text-gray-400 mt-1">{content.length}/2000</div>
      </div>

      {/* 图片上传 */}
      <div className="mb-6">
        <label className="block text-sm font-medium text-gray-700 mb-2">图片 ({images.length}/9)</label>
        <div className="grid grid-cols-3 gap-2">
          {images.map((img, index) => (
            <div key={index} className="relative aspect-square">
              <img src={img} alt="" className="w-full h-full object-cover rounded-lg" />
              <button
                onClick={() => handleRemoveImage(index)}
                className="absolute top-1 right-1 w-6 h-6 bg-black/50 rounded-full flex items-center justify-center text-white text-sm"
              >
                ×
              </button>
            </div>
          ))}
          {images.length < 9 && (
            <button
              onClick={handleAddImage}
              className="aspect-square border-2 border-dashed border-gray-300 rounded-lg flex flex-col items-center justify-center text-gray-400"
            >
              <svg className="w-8 h-8" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M12 4v16m8-8H4" />
              </svg>
              <span className="text-xs mt-1">添加图片</span>
            </button>
          )}
        </div>
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
