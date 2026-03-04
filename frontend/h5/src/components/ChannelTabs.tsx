'use client';

import { useEffect, useState } from 'react';
import { useChannelStore } from '@/store/channelStore';

export default function ChannelTabs() {
  const {
    channels,
    currentChannel,
    loading,
    error,
    fetchChannels,
    setCurrentChannel,
    refreshChannels,
  } = useChannelStore();

  const [mounted, setMounted] = useState(false);

  useEffect(() => {
    setMounted(true);
    fetchChannels();
  }, [fetchChannels]);

  // 下拉刷新
  const handleRefresh = async () => {
    await refreshChannels();
  };

  // 等待客户端挂载
  if (!mounted) {
    return (
      <div className="flex overflow-x-auto px-4 pb-3 gap-4">
        {[1, 2, 3, 4, 5].map((i) => (
          <div
            key={i}
            className="flex-shrink-0 px-4 py-2 rounded-full bg-gray-200 animate-pulse"
            style={{ width: '80px' }}
          />
        ))}
      </div>
    );
  }

  if (loading && channels.length === 0) {
    return (
      <div className="flex overflow-x-auto px-4 pb-3 gap-4">
        {[1, 2, 3, 4, 5].map((i) => (
          <div
            key={i}
            className="flex-shrink-0 px-4 py-2 rounded-full bg-gray-200 animate-pulse"
            style={{ width: '80px' }}
          />
        ))}
      </div>
    );
  }

  if (error && channels.length === 0) {
    return (
      <div className="px-4 pb-3">
        <div className="text-red-500 text-sm">{error}</div>
        <button
          onClick={handleRefresh}
          className="text-blue-500 text-sm underline mt-1"
        >
          重试
        </button>
      </div>
    );
  }

  if (channels.length === 0) {
    return (
      <div className="px-4 pb-3">
        <div className="text-gray-500 text-sm">暂无频道数据</div>
        <button
          onClick={handleRefresh}
          className="text-blue-500 text-sm underline mt-1"
        >
          刷新
        </button>
      </div>
    );
  }

  return (
    <div className="flex overflow-x-auto px-4 pb-3 gap-4 scrollbar-hide">
      {channels.map((channel) => (
        <button
          key={channel.id}
          onClick={() => setCurrentChannel(channel)}
          className={`flex-shrink-0 px-4 py-2 rounded-full text-sm font-medium transition-colors whitespace-nowrap ${
            currentChannel?.id === channel.id
              ? 'bg-blue-500 text-white'
              : 'bg-gray-100 text-gray-700 hover:bg-gray-200'
          }`}
        >
          <span className="mr-1">{channel.icon}</span>
          {channel.name}
        </button>
      ))}
    </div>
  );
}
