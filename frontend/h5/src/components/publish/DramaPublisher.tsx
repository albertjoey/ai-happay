'use client';

import { useState, useRef } from 'react';
import { publishMaterial, uploadFile, DramaPublishRequest, Episode as EpisodeType } from '@/lib/publishApi';

interface DramaPublisherProps {
  type: 'short_drama' | 'drama';
  onSuccess?: () => void;
}

interface Episode {
  episode_number: number;
  title: string;
  cover_url?: string;
  video_url?: string;
  images?: string[];
  duration?: number;
}

export default function DramaPublisher({ type, onSuccess }: DramaPublisherProps) {
  const coverInputRef = useRef<HTMLInputElement>(null);
  const [title, setTitle] = useState('');
  const [description, setDescription] = useState('');
  const [coverUrl, setCoverUrl] = useState('');
  const [episodes, setEpisodes] = useState<Episode[]>([{ episode_number: 1, title: '第1集' }]);
  const [submitting, setSubmitting] = useState(false);

  const isDrama = type === 'drama';

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

  // 添加剧集
  const addEpisode = () => {
    setEpisodes([...episodes, {
      episode_number: episodes.length + 1,
      title: `第${episodes.length + 1}集`
    }]);
  };

  // 删除剧集
  const removeEpisode = (index: number) => {
    if (episodes.length > 1) {
      setEpisodes(episodes.filter((_, i) => i !== index));
    }
  };

  // 更新剧集
  const updateEpisode = (index: number, field: keyof Episode, value: any) => {
    const newEpisodes = [...episodes];
    newEpisodes[index] = { ...newEpisodes[index], [field]: value };
    setEpisodes(newEpisodes);
  };

  // 上传剧集视频/图片
  const handleEpisodeUpload = async (index: number, file: File) => {
    try {
      const uploadedUrl = await uploadFile(file, isDrama ? 'image' : 'video');

      if (isDrama) {
        const currentImages = episodes[index].images || [];
        updateEpisode(index, 'images', [...currentImages, uploadedUrl]);
      } else {
        updateEpisode(index, 'video_url', uploadedUrl);
      }
    } catch (error) {
      alert('上传失败，请重试');
    }
  };

  // 提交发布
  const handleSubmit = async () => {
    if (!title.trim()) {
      alert('请输入标题');
      return;
    }

    setSubmitting(true);
    try {
      const publishData: DramaPublishRequest = {
        type,
        title,
        description,
        cover_url: coverUrl || `https://picsum.photos/400/600?random=${Date.now()}`,
        episodes: episodes.map(ep => ({
          episode_number: ep.episode_number,
          title: ep.title,
          cover_url: ep.cover_url,
          video_url: ep.video_url,
          images: ep.images,
          duration: ep.duration,
        })),
        total_episodes: episodes.length,
        category: isDrama ? '漫剧' : '短剧',
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
        ref={coverInputRef}
        type="file"
        accept="image/*"
        onChange={handleCoverChange}
        className="hidden"
      />

      {/* 封面 */}
      <div className="mb-4">
        <label className="block text-sm font-medium text-gray-700 mb-2">封面</label>
        <div
          onClick={handleSelectCover}
          className="w-32 h-44 bg-gray-100 rounded-lg flex items-center justify-center overflow-hidden cursor-pointer hover:bg-gray-200 transition-colors"
        >
          {coverUrl ? (
            <img src={coverUrl} alt="" className="w-full h-full object-cover" />
          ) : (
            <div className="text-center text-gray-400">
              <div className="text-2xl">📷</div>
              <div className="text-xs mt-1">上传封面</div>
            </div>
          )}
        </div>
      </div>

      {/* 标题 */}
      <div className="mb-4">
        <label className="block text-sm font-medium text-gray-700 mb-1">
          {isDrama ? '漫剧名称' : '短剧名称'}
        </label>
        <input
          type="text"
          value={title}
          onChange={(e) => setTitle(e.target.value)}
          placeholder="请输入名称"
          className="w-full px-4 py-3 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
        />
      </div>

      {/* 描述 */}
      <div className="mb-4">
        <label className="block text-sm font-medium text-gray-700 mb-1">简介</label>
        <textarea
          value={description}
          onChange={(e) => setDescription(e.target.value)}
          placeholder="介绍一下你的作品..."
          className="w-full px-4 py-3 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 min-h-[80px]"
        />
      </div>

      {/* 剧集管理 */}
      <div className="mb-6">
        <div className="flex items-center justify-between mb-2">
          <label className="text-sm font-medium text-gray-700">
            {isDrama ? '分话管理' : '分集管理'} ({episodes.length}集)
          </label>
          <button
            onClick={addEpisode}
            className="text-blue-500 text-sm hover:text-blue-600"
          >
            + 添加{isDrama ? '话' : '集'}
          </button>
        </div>
        <div className="space-y-3">
          {episodes.map((ep, index) => (
            <div key={index} className="p-3 bg-gray-50 rounded-lg">
              <div className="flex items-center gap-2 mb-2">
                <input
                  type="text"
                  value={ep.title}
                  onChange={(e) => updateEpisode(index, 'title', e.target.value)}
                  className="flex-1 px-3 py-2 border border-gray-200 rounded text-sm"
                  placeholder="集名"
                />
                <button
                  onClick={() => removeEpisode(index)}
                  className="text-red-500 text-sm hover:text-red-600 disabled:opacity-50"
                  disabled={episodes.length === 1}
                >
                  删除
                </button>
              </div>

              {/* 上传按钮 */}
              <label className="block">
                <input
                  type="file"
                  accept={isDrama ? 'image/*' : 'video/*'}
                  multiple={isDrama}
                  onChange={(e) => {
                    const files = e.target.files;
                    if (files) {
                      Array.from(files).forEach(file => handleEpisodeUpload(index, file));
                    }
                  }}
                  className="hidden"
                />
                <div className="px-3 py-2 border border-dashed border-gray-300 rounded text-center text-sm text-gray-500 cursor-pointer hover:border-blue-400">
                  {isDrama ? '📷 上传图片' : '🎬 上传视频'}
                </div>
              </label>

              {/* 已上传内容显示 */}
              {isDrama && ep.images && ep.images.length > 0 && (
                <div className="mt-2 flex gap-1 flex-wrap">
                  {ep.images.map((img, imgIndex) => (
                    <div key={imgIndex} className="w-12 h-12 rounded overflow-hidden">
                      <img src={img} alt="" className="w-full h-full object-cover" />
                    </div>
                  ))}
                </div>
              )}
              {!isDrama && ep.video_url && (
                <div className="mt-2 text-xs text-green-600">✓ 视频已上传</div>
              )}
            </div>
          ))}
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
