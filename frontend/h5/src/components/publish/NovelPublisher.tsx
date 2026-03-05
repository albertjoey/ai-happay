'use client';

import { useState } from 'react';

interface NovelPublisherProps {
  onSuccess?: () => void;
}

interface Chapter {
  title: string;
  content: string;
}

export default function NovelPublisher({ onSuccess }: NovelPublisherProps) {
  const [title, setTitle] = useState('');
  const [description, setDescription] = useState('');
  const [author, setAuthor] = useState('');
  const [coverUrl, setCoverUrl] = useState('');
  const [chapters, setChapters] = useState<Chapter[]>([{ title: '第一章', content: '' }]);
  const [currentChapter, setCurrentChapter] = useState(0);
  const [submitting, setSubmitting] = useState(false);

  // 添加章节
  const addChapter = () => {
    setChapters([...chapters, { title: `第${chapters.length + 1}章`, content: '' }]);
    setCurrentChapter(chapters.length);
  };

  // 删除章节
  const removeChapter = (index: number) => {
    if (chapters.length > 1) {
      const newChapters = chapters.filter((_, i) => i !== index);
      setChapters(newChapters);
      if (currentChapter >= newChapters.length) {
        setCurrentChapter(newChapters.length - 1);
      }
    }
  };

  // 更新章节
  const updateChapter = (index: number, field: keyof Chapter, value: string) => {
    const newChapters = [...chapters];
    newChapters[index] = { ...newChapters[index], [field]: value };
    setChapters(newChapters);
  };

  // 计算总字数
  const totalWords = chapters.reduce((sum, ch) => sum + ch.content.length, 0);

  // 提交发布
  const handleSubmit = async () => {
    if (!title.trim()) {
      alert('请输入书名');
      return;
    }
    if (totalWords < 100) {
      alert('内容太短，请至少写100字');
      return;
    }

    setSubmitting(true);
    try {
      const response = await fetch('http://localhost:4004/api/v1/material', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({
          type: 'novel',
          title,
          description,
          cover_url: coverUrl || `https://picsum.photos/400/600?random=${Date.now()}`,
          author: author || '匿名作者',
          category: '小说',
          word_count: totalWords,
          chapter_count: chapters.length,
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
        <div className="w-24 h-32 bg-gray-100 rounded flex items-center justify-center overflow-hidden">
          {coverUrl ? (
            <img src={coverUrl} alt="" className="w-full h-full object-cover" />
          ) : (
            <div className="text-center text-gray-400">
              <div className="text-2xl">📖</div>
              <div className="text-xs mt-1">封面</div>
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

      {/* 书名 */}
      <div className="mb-4">
        <label className="block text-sm font-medium text-gray-700 mb-1">书名</label>
        <input
          type="text"
          value={title}
          onChange={(e) => setTitle(e.target.value)}
          placeholder="请输入书名"
          className="w-full px-4 py-3 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
        />
      </div>

      {/* 作者 */}
      <div className="mb-4">
        <label className="block text-sm font-medium text-gray-700 mb-1">作者</label>
        <input
          type="text"
          value={author}
          onChange={(e) => setAuthor(e.target.value)}
          placeholder="请输入作者名"
          className="w-full px-4 py-3 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
        />
      </div>

      {/* 简介 */}
      <div className="mb-4">
        <label className="block text-sm font-medium text-gray-700 mb-1">简介</label>
        <textarea
          value={description}
          onChange={(e) => setDescription(e.target.value)}
          placeholder="介绍一下你的小说..."
          className="w-full px-4 py-3 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 min-h-[80px]"
        />
      </div>

      {/* 章节管理 */}
      <div className="mb-4">
        <div className="flex items-center justify-between mb-2">
          <label className="text-sm font-medium text-gray-700">
            章节管理 ({chapters.length}章, 共{totalWords}字)
          </label>
          <button onClick={addChapter} className="text-blue-500 text-sm">
            + 添加章节
          </button>
        </div>

        {/* 章节列表 */}
        <div className="flex gap-2 overflow-x-auto pb-2">
          {chapters.map((ch, index) => (
            <button
              key={index}
              onClick={() => setCurrentChapter(index)}
              className={`flex-shrink-0 px-3 py-1 rounded text-sm ${
                currentChapter === index
                  ? 'bg-blue-500 text-white'
                  : 'bg-gray-100 text-gray-700'
              }`}
            >
              {ch.title || `第${index + 1}章`}
            </button>
          ))}
        </div>
      </div>

      {/* 当前章节编辑 */}
      <div className="mb-4 p-3 bg-gray-50 rounded-lg">
        <div className="flex items-center justify-between mb-2">
          <input
            type="text"
            value={chapters[currentChapter]?.title || ''}
            onChange={(e) => updateChapter(currentChapter, 'title', e.target.value)}
            className="flex-1 px-2 py-1 border border-gray-200 rounded text-sm"
            placeholder="章节标题"
          />
          <button
            onClick={() => removeChapter(currentChapter)}
            className="ml-2 text-red-500 text-sm"
            disabled={chapters.length === 1}
          >
            删除
          </button>
        </div>
        <textarea
          value={chapters[currentChapter]?.content || ''}
          onChange={(e) => updateChapter(currentChapter, 'content', e.target.value)}
          placeholder="开始写作..."
          className="w-full px-3 py-2 border border-gray-200 rounded text-sm min-h-[200px]"
        />
        <div className="text-right text-xs text-gray-400 mt-1">
          {chapters[currentChapter]?.content.length || 0}字
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
