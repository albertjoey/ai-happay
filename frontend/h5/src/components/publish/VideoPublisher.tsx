'use client';

import { useState, useRef } from 'react';
import { publishMaterial, uploadFile, VideoPublishRequest } from '@/lib/publishApi';

interface VideoPublisherProps {
  type: 'short_video' | 'long_video';
  onSuccess?: () => void;
}

export default function VideoPublisher({ type, onSuccess }: VideoPublisherProps) {
  const videoInputRef = useRef<HTMLInputElement>(null);
  const coverInputRef = useRef<HTMLInputElement>(null);
  const [title, setTitle] = useState('');
  const [description, setDescription] = useState('');
  const [coverUrl, setCoverUrl] = useState('');
  const [videoUrl, setVideoUrl] = useState('');
  const [duration, setDuration] = useState(0);
  const [submitting, setSubmitting] = useState(false);
  const [uploadingVideo, setUploadingVideo] = useState(false);
  const [uploadProgress, setUploadProgress] = useState(0);

  const isShort = type === 'short_video';

  // 选择视频
  const handleSelectVideo = () => {
    videoInputRef.current?.click();
  };

  // 处理视频选择
  const handleVideoChange = async (e: React.ChangeEvent<HTMLInputElement>) => {
    const file = e.target.files?.[0];
    if (!file) return;

    setUploadingVideo(true);
    setUploadProgress(0);

    try {
      // 模拟上传进度
      const progressInterval = setInterval(() => {
        setUploadProgress(prev => Math.min(prev + 10, 90));
      }, 200);

      const uploadedUrl = await uploadFile(file, 'video');
      setVideoUrl(uploadedUrl);

      // 获取视频时长
      const video = document.createElement('video');
      video.src = URL.createObjectURL(file);
      video.onloadedmetadata = () => {
        setDuration(Math.floor(video.duration));
        URL.revokeObjectURL(video.src);
      };

      clearInterval(progressInterval);
      setUploadProgress(100);
    } catch (error) {
      alert('视频上传失败，请重试');
    } finally {
      setUploadingVideo(false);
      if (videoInputRef.current) {
        videoInputRef.current.value = '';
      }
    }
  };

  // 选择封面
  const handleSelectCover = () => {
    coverInputRef.current?.click();
  };

  // 处理封面上传
  const handleCoverChange = async (e: React.ChangeEvent<HTMLInputElement>) => {
    const file = e.target.files?.[0];
    if (!file) return;

    try {
      const uploadedUrl = await uploadFile(file, 'image');
      setCoverUrl(uploadedUrl);
    } catch (error) {
      alert('封面上传失败，请重试');
    } finally {
      if (coverInputRef.current) {
        coverInputRef.current.value = '';
      }
    }
  };

  // 提交发布
  const handleSubmit = async () => {
    if (!title.trim()) {
      alert('请输入标题');
      return;
    }
    if (!videoUrl) {
      alert('请上传视频');
      return;
    }

    setSubmitting(true);
    try {
      const publishData: VideoPublishRequest = {
        type,
        title,
        description,
        cover_url: coverUrl || `https://picsum.photos/800/450?random=${Date.now()}`,
        video_url: videoUrl,
        duration: duration || (isShort ? 60 : 1800),
        category: isShort ? '短视频' : '长视频',
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
        ref={videoInputRef}
        type="file"
        accept="video/*"
        onChange={handleVideoChange}
        className="hidden"
      />
      <input
        ref={coverInputRef}
        type="file"
        accept="image/*"
        onChange={handleCoverChange}
        className="hidden"
      />

      {/* 视频上传区域 */}
      <div className="mb-4">
        <label className="block text-sm font-medium text-gray-700 mb-2">
          {isShort ? '短视频' : '视频文件'}
        </label>
        <div
          onClick={handleSelectVideo}
          className="border-2 border-dashed border-gray-300 rounded-lg p-8 text-center cursor-pointer hover:border-blue-400 transition-colors"
        >
          {uploadingVideo ? (
            <div>
              <div className="w-12 h-12 mx-auto mb-2 border-4 border-blue-500 border-t-transparent rounded-full animate-spin" />
              <p className="text-gray-600 text-sm">上传中... {uploadProgress}%</p>
            </div>
          ) : videoUrl ? (
            <div>
              <div className="text-4xl mb-2">✅</div>
              <p className="text-green-600 text-sm font-medium">视频已上传</p>
              {duration > 0 && (
                <p className="text-gray-500 text-xs mt-1">时长: {Math.floor(duration / 60)}:{(duration % 60).toString().padStart(2, '0')}</p>
              )}
            </div>
          ) : (
            <div>
              <div className="text-4xl mb-2">🎬</div>
              <p className="text-gray-500 text-sm">点击上传视频</p>
              <p className="text-gray-400 text-xs mt-1">
                {isShort ? '支持竖屏视频，时长60秒以内' : '支持MP4格式，最大2GB'}
              </p>
            </div>
          )}
        </div>
      </div>

      {/* 封面设置 */}
      <div className="mb-4">
        <label className="block text-sm font-medium text-gray-700 mb-2">封面</label>
        <div className="flex gap-2">
          <div
            onClick={handleSelectCover}
            className="w-24 h-16 bg-gray-100 rounded flex items-center justify-center text-gray-400 cursor-pointer hover:bg-gray-200 transition-colors overflow-hidden"
          >
            {coverUrl ? (
              <img src={coverUrl} alt="" className="w-full h-full object-cover" />
            ) : (
              <span className="text-xs">选择封面</span>
            )}
          </div>
          <button
            onClick={() => setCoverUrl(`https://picsum.photos/800/450?random=${Date.now()}`)}
            className="px-3 py-1 border border-gray-300 rounded text-sm hover:bg-gray-50"
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
        disabled={submitting || uploadingVideo}
        className="w-full py-3 bg-blue-500 text-white rounded-lg font-medium disabled:bg-gray-300"
      >
        {submitting ? '发布中...' : '发布'}
      </button>
    </div>
  );
}
