'use client';

import { useChannelStore } from '@/store/channelStore';
import { useEffect, useState } from 'react';
import axios from 'axios';

interface Content {
  id: number;
  title: string;
  description: string;
  cover: string;
  type: number;
  view_count: number;
  like_count: number;
  author_name: string;
  author_avatar: string;
}

export default function ContentList() {
  const { currentChannel, channelConfig } = useChannelStore();
  const [contents, setContents] = useState<Content[]>([]);
  const [loading, setLoading] = useState(false);

  useEffect(() => {
    if (currentChannel) {
      loadContent();
    }
  }, [currentChannel]);

  const loadContent = async () => {
    setLoading(true);
    try {
      // 根据频道配置过滤内容类型
      const contentTypes = channelConfig?.content_type || {};
      const types = Object.keys(contentTypes).filter((key) => contentTypes[key]);

      const response = await axios.get('http://localhost:4001/api/v1/content/list', {
        params: {
          page: 1,
          page_size: 20,
          channel_id: currentChannel?.id,
        },
      });

      setContents(response.data.list || []);
    } catch (error) {
      console.error('加载内容失败:', error);
      // 使用模拟数据
      setContents(getMockContent());
    } finally {
      setLoading(false);
    }
  };

  const getMockContent = (): Content[] => {
    return Array.from({ length: 10 }, (_, i) => ({
      id: i + 1,
      title: `${currentChannel?.name || '推荐'}内容 ${i + 1}`,
      description: '这是内容的描述信息，展示内容的详细信息',
      cover: `https://picsum.photos/seed/${currentChannel?.code || 'content'}${i}/400/600`,
      type: [1, 2, 3][i % 3],
      view_count: Math.floor(Math.random() * 10000),
      like_count: Math.floor(Math.random() * 1000),
      author_name: '创作者',
      author_avatar: 'https://picsum.photos/seed/avatar/100/100',
    }));
  };

  // 根据展示模式渲染不同布局
  const renderContent = () => {
    const displayMode = channelConfig?.display_mode || 'default';

    switch (displayMode) {
      case 'waterfall':
        return <WaterfallLayout contents={contents} loading={loading} />;
      case 'grid':
        return <GridLayout contents={contents} loading={loading} />;
      case 'list':
        return <ListLayout contents={contents} loading={loading} />;
      default:
        return <DefaultLayout contents={contents} loading={loading} />;
    }
  };

  return <div className="min-h-screen bg-gray-50">{renderContent()}</div>;
}

// 瀑布流布局
function WaterfallLayout({ contents, loading }: { contents: Content[]; loading: boolean }) {
  if (loading) {
    return <div className="p-4 text-center text-gray-500">加载中...</div>;
  }

  return (
    <div className="grid grid-cols-2 gap-2 p-2">
      {contents.map((content) => (
        <div key={content.id} className="bg-white rounded-lg overflow-hidden shadow-sm">
          <img src={content.cover} alt={content.title} className="w-full h-48 object-cover" />
          <div className="p-3">
            <h3 className="text-sm font-medium line-clamp-2">{content.title}</h3>
            <div className="flex items-center justify-between mt-2 text-xs text-gray-500">
              <span>{content.author_name}</span>
              <span>❤️ {content.like_count}</span>
            </div>
          </div>
        </div>
      ))}
    </div>
  );
}

// 网格布局
function GridLayout({ contents, loading }: { contents: Content[]; loading: boolean }) {
  if (loading) {
    return <div className="p-4 text-center text-gray-500">加载中...</div>;
  }

  return (
    <div className="grid grid-cols-3 gap-2 p-2">
      {contents.map((content) => (
        <div key={content.id} className="bg-white rounded-lg overflow-hidden shadow-sm">
          <img src={content.cover} alt={content.title} className="w-full h-32 object-cover" />
          <div className="p-2">
            <h3 className="text-xs font-medium line-clamp-1">{content.title}</h3>
          </div>
        </div>
      ))}
    </div>
  );
}

// 列表布局
function ListLayout({ contents, loading }: { contents: Content[]; loading: boolean }) {
  if (loading) {
    return <div className="p-4 text-center text-gray-500">加载中...</div>;
  }

  return (
    <div className="space-y-2 p-2">
      {contents.map((content) => (
        <div key={content.id} className="bg-white rounded-lg p-3 shadow-sm flex space-x-3">
          <img src={content.cover} alt={content.title} className="w-24 h-24 object-cover rounded" />
          <div className="flex-1">
            <h3 className="text-sm font-medium line-clamp-2">{content.title}</h3>
            <p className="text-xs text-gray-500 mt-1 line-clamp-1">{content.description}</p>
            <div className="flex items-center justify-between mt-2 text-xs text-gray-500">
              <span>{content.author_name}</span>
              <span>👁 {content.view_count}</span>
            </div>
          </div>
        </div>
      ))}
    </div>
  );
}

// 默认布局
function DefaultLayout({ contents, loading }: { contents: Content[]; loading: boolean }) {
  if (loading) {
    return <div className="p-4 text-center text-gray-500">加载中...</div>;
  }

  return (
    <div className="space-y-3 p-3">
      {contents.map((content) => (
        <div key={content.id} className="bg-white rounded-lg overflow-hidden shadow-sm">
          <img src={content.cover} alt={content.title} className="w-full h-48 object-cover" />
          <div className="p-3">
            <h3 className="text-base font-medium">{content.title}</h3>
            <p className="text-sm text-gray-500 mt-1">{content.description}</p>
            <div className="flex items-center justify-between mt-2 text-sm text-gray-500">
              <div className="flex items-center space-x-2">
                <img src={content.author_avatar} alt="" className="w-6 h-6 rounded-full" />
                <span>{content.author_name}</span>
              </div>
              <div className="flex items-center space-x-3">
                <span>👁 {content.view_count}</span>
                <span>❤️ {content.like_count}</span>
              </div>
            </div>
          </div>
        </div>
      ))}
    </div>
  );
}
