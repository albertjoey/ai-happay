'use client';

import { useEffect } from 'react';
import { useChannelStore } from '@/store/channelStore';
import clsx from 'clsx';

export default function ChannelNav() {
  const { channels, currentChannel, loading, fetchChannels, setCurrentChannel } = useChannelStore();

  useEffect(() => {
    fetchChannels();
  }, [fetchChannels]);

  if (loading && channels.length === 0) {
    return (
      <div className="bg-white border-b sticky top-0 z-10">
        <div className="flex items-center px-4 py-3 space-x-6 overflow-x-auto">
          {[1, 2, 3, 4, 5].map((i) => (
            <div key={i} className="h-6 w-16 bg-gray-200 rounded animate-pulse" />
          ))}
        </div>
      </div>
    );
  }

  return (
    <div className="bg-white border-b sticky top-0 z-10">
      <div className="flex items-center px-4 py-3 space-x-6 overflow-x-auto scrollbar-hide">
        {channels.map((channel) => (
          <button
            key={channel.id}
            onClick={() => setCurrentChannel(channel)}
            className={clsx(
              'flex-shrink-0 text-base font-medium transition-colors whitespace-nowrap',
              currentChannel?.id === channel.id
                ? 'text-blue-600 border-b-2 border-blue-600 pb-1'
                : 'text-gray-600 hover:text-gray-900'
            )}
          >
            {channel.name}
          </button>
        ))}
      </div>
    </div>
  );
}
