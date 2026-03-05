'use client';

import { useState } from 'react';

interface DramaPublisherProps {
  type: 'short_drama' | 'drama';
  onSuccess?: () => void;
}

interface Episode {
  title: string;
  videoUrl?: string;
  images?: string[];
}

export default function DramaPublisher({ type, onSuccess }: DramaPublisherProps) {
  const [title, setTitle] = useState('');
  const [description, setDescription] = useState('');
  const [coverUrl, setCoverUrl] = useState('');
  const [episodes, setEpisodes] = useState<Episode[]>([{ title: '第1集' }]);
  const [submitting, setSubmitting] = useState(false);

  const isDrama = type === 'drama';

  // 添加剧集
  const addEpisode = () => {
    setEpisodes([...episodes, { title: `第${episodes.length + 1}集` }]);
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
          cover_url: coverUrl || `https://picsum.photos/400/600?random=${Date.now()}`,
          author: '用户发布',
          category: isDrama ? '漫剧' : '短剧',
          chapter_count: episodes.length,
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
      {/* 封面 */}
      <div className="mb-4">
        <label className="block text-sm font-medium text-gray-700 mb-2">封面</label>
        <div className="w-32 h-44 bg-gray-100 rounded-lg flex items-center justify-center overflow-hidden">
          {coverUrl ? (
            <img src={coverUrl} alt="" className="w-full h-full object-cover" />
          ) : (
            <div className="text-center text-gray-400">
              <div className="text-2xl">📷</div>
              <div className="text-xs mt-1">上传封面</div>
            </div>
          )}
        </div>
        <button
          onClick={() => setCoverUrl(`https://picsum.photos/400/600?random=${Date.now()}`)}
          className="mt-2 px-3 py-1 border border-gray-300 rounded text-sm"
        >
          选择封面
        </button>
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
            className="text-blue-500 text-sm"
          >
            + 添加{isDrama ? '话' : '集'}
          </button>
        </div>
        <div className="space-y-2">
          {episodes.map((ep, index) => (
            <div key={index} className="flex items-center gap-2 p-2 bg-gray-50 rounded-lg">
              <input
                type="text"
                value={ep.title}
                onChange={(e) => updateEpisode(index, 'title', e.target.value)}
                className="flex-1 px-3 py-2 border border-gray-200 rounded text-sm"
              />
              <button
                onClick={() => removeEpisode(index)}
                className="text-red-500 text-sm"
                disabled={episodes.length === 1}
              >
                删除
              </button>
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
