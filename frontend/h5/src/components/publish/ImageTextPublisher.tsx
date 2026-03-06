'use client';

import { useState, useRef } from 'react';
import { useRouter } from 'next/navigation';
import { publishMaterial, uploadFile, ImageTextPublishRequest } from '@/lib/publishApi';

interface ImageTextPublisherProps {
  onSuccess?: () => void;
}

export default function ImageTextPublisher({ onSuccess }: ImageTextPublisherProps) {
  const router = useRouter();
  const fileInputRef = useRef<HTMLInputElement>(null);
  const [title, setTitle] = useState('');
  const [content, setContent] = useState('');
  const [images, setImages] = useState<string[]>([]);
  const [submitting, setSubmitting] = useState(false);
  const [uploading, setUploading] = useState(false);

  // 选择图片
  const handleSelectImage = () => {
    if (images.length >= 9) {
      alert('最多只能添加9张图片');
      return;
    }
    fileInputRef.current?.click();
  };

  // 处理文件选择
  const handleFileChange = async (e: React.ChangeEvent<HTMLInputElement>) => {
    const files = e.target.files;
    if (!files || files.length === 0) return;

    const remainingSlots = 9 - images.length;
    const filesToUpload = Array.from(files).slice(0, remainingSlots);

    setUploading(true);
    try {
      const uploadPromises = filesToUpload.map(file => uploadFile(file, 'image'));
      const uploadedUrls = await Promise.all(uploadPromises);
      setImages([...images, ...uploadedUrls]);
    } catch (error) {
      alert('图片上传失败，请重试');
    } finally {
      setUploading(false);
      if (fileInputRef.current) {
        fileInputRef.current.value = '';
      }
    }
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
      const publishData: ImageTextPublishRequest = {
        type: 'image_text',
        title,
        content,
        images,
        category: '图文',
      };

      await publishMaterial(publishData);
      alert('发布成功！');
      onSuccess?.();
    } catch (error) {
      alert('发布失败，请重试');
    } finally {
      setSubmitting(false);
    }
  };

  return (
    <div className="p-4">
      {/* 隐藏的文件输入 */}
      <input
        ref={fileInputRef}
        type="file"
        accept="image/*"
        multiple
        onChange={handleFileChange}
        className="hidden"
      />

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
              onClick={handleSelectImage}
              disabled={uploading}
              className="aspect-square border-2 border-dashed border-gray-300 rounded-lg flex flex-col items-center justify-center text-gray-400 disabled:opacity-50"
            >
              {uploading ? (
                <div className="w-6 h-6 border-2 border-gray-400 border-t-transparent rounded-full animate-spin" />
              ) : (
                <>
                  <svg className="w-8 h-8" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M12 4v16m8-8H4" />
                  </svg>
                  <span className="text-xs mt-1">添加图片</span>
                </>
              )}
            </button>
          )}
        </div>
      </div>

      {/* 发布按钮 */}
      <button
        onClick={handleSubmit}
        disabled={submitting || uploading}
        className="w-full py-3 bg-blue-500 text-white rounded-lg font-medium disabled:bg-gray-300"
      >
        {submitting ? '发布中...' : '发布'}
      </button>
    </div>
  );
}
