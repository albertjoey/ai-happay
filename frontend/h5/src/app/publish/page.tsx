'use client';

import { useState } from 'react';
import { useRouter } from 'next/navigation';
import Head from 'next/head';
import ImageTextPublisher from '@/components/publish/ImageTextPublisher';
import VideoPublisher from '@/components/publish/VideoPublisher';
import DramaPublisher from '@/components/publish/DramaPublisher';
import NovelPublisher from '@/components/publish/NovelPublisher';

// 发布类型
const publishTypes = [
  { type: 'image_text', name: '图文', icon: '📄', color: 'from-green-400 to-green-600' },
  { type: 'short_video', name: '短视频', icon: '📱', color: 'from-pink-400 to-pink-600' },
  { type: 'long_video', name: '长视频', icon: '🎬', color: 'from-blue-400 to-blue-600' },
  { type: 'short_drama', name: '短剧', icon: '🎭', color: 'from-purple-400 to-purple-600' },
  { type: 'drama', name: '漫剧', icon: '📚', color: 'from-orange-400 to-orange-600' },
  { type: 'novel', name: '小说', icon: '📖', color: 'from-indigo-400 to-indigo-600' },
];

export default function PublishPage() {
  const router = useRouter();
  const [selectedType, setSelectedType] = useState<string | null>(null);

  // 渲染发布表单
  const renderPublisher = () => {
    switch (selectedType) {
      case 'image_text':
        return <ImageTextPublisher onSuccess={() => router.push('/')} />;
      case 'short_video':
      case 'long_video':
        return <VideoPublisher type={selectedType} onSuccess={() => router.push('/')} />;
      case 'short_drama':
      case 'drama':
        return <DramaPublisher type={selectedType} onSuccess={() => router.push('/')} />;
      case 'novel':
        return <NovelPublisher onSuccess={() => router.push('/')} />;
      default:
        return null;
    }
  };

  return (
    <div className="min-h-screen bg-gray-50">
      <Head>
        <title>发布内容 - Happy</title>
      </Head>

      {/* 顶部导航 */}
      <header className="sticky top-0 z-50 bg-white border-b border-gray-200">
        <div className="flex items-center justify-between px-4 py-3">
          <button onClick={() => selectedType ? setSelectedType(null) : router.back()} className="p-2 -ml-2">
            <svg className="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M15 19l-7-7 7-7" />
            </svg>
          </button>
          <h1 className="text-lg font-semibold">
            {selectedType ? publishTypes.find(t => t.type === selectedType)?.name : '发布内容'}
          </h1>
          <div className="w-10" />
        </div>
      </header>

      {/* 内容区域 */}
      <main className="pb-20">
        {!selectedType ? (
          // 类型选择
          <div className="p-4">
            <h2 className="text-lg font-semibold text-gray-800 mb-4">选择发布类型</h2>
            <div className="grid grid-cols-3 gap-4">
              {publishTypes.map((item) => (
                <button
                  key={item.type}
                  onClick={() => setSelectedType(item.type)}
                  className="flex flex-col items-center p-4 bg-white rounded-xl shadow-sm hover:shadow-md transition-shadow"
                >
                  <div className={`w-14 h-14 rounded-full bg-gradient-to-br ${item.color} flex items-center justify-center text-white text-2xl mb-2`}>
                    {item.icon}
                  </div>
                  <span className="text-sm text-gray-700">{item.name}</span>
                </button>
              ))}
            </div>
          </div>
        ) : (
          // 发布表单
          renderPublisher()
        )}
      </main>
    </div>
  );
}
