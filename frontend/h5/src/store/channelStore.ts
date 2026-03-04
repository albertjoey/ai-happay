import { create } from 'zustand';
import { Channel, ChannelConfig, getChannelList, getChannelConfig } from '@/lib/channel';

interface ChannelState {
  channels: Channel[];
  currentChannel: Channel | null;
  channelConfig: ChannelConfig | null;
  loading: boolean;
  error: string | null;
  lastFetchTime: number;

  // Actions
  fetchChannels: () => Promise<void>;
  setCurrentChannel: (channel: Channel) => void;
  fetchChannelConfig: (channelId: number) => Promise<void>;
  refreshChannels: () => Promise<void>;
}

const CACHE_DURATION = 60 * 1000; // 1分钟缓存

// 安全的localStorage操作
const safeGetItem = (key: string): string | null => {
  try {
    if (typeof window !== 'undefined' && window.localStorage) {
      return localStorage.getItem(key);
    }
  } catch (e) {
    console.warn('localStorage not available');
  }
  return null;
};

const safeSetItem = (key: string, value: string): void => {
  try {
    if (typeof window !== 'undefined' && window.localStorage) {
      localStorage.setItem(key, value);
    }
  } catch (e) {
    console.warn('localStorage not available');
  }
};

export const useChannelStore = create<ChannelState>()((set, get) => ({
  channels: [],
  currentChannel: null,
  channelConfig: null,
  loading: false,
  error: null,
  lastFetchTime: 0,

  fetchChannels: async () => {
    const { lastFetchTime } = get();
    const now = Date.now();

    // 如果缓存未过期,直接返回
    if (now - lastFetchTime < CACHE_DURATION && get().channels.length > 0) {
      return;
    }

    set({ loading: true, error: null });
    try {
      const response = await getChannelList();
      const channels = response.list || [];

      set({
        channels,
        lastFetchTime: now,
        loading: false,
      });

      // 如果没有当前频道,默认选择第一个
      if (!get().currentChannel && channels.length > 0) {
        set({ currentChannel: channels[0] });
        // 保存到localStorage
        safeSetItem('currentChannelId', String(channels[0].id));
      }
    } catch (error) {
      set({
        error: '加载频道失败',
        loading: false,
      });
    }
  },

  setCurrentChannel: (channel: Channel) => {
    set({ currentChannel: channel });
    // 保存到localStorage
    safeSetItem('currentChannelId', String(channel.id));
    // 切换频道时加载配置
    get().fetchChannelConfig(channel.id);
  },

  fetchChannelConfig: async (channelId: number) => {
    try {
      const config = await getChannelConfig(channelId);
      set({ channelConfig: config });
    } catch (error) {
      console.error('加载频道配置失败:', error);
    }
  },

  refreshChannels: async () => {
    set({ lastFetchTime: 0 });
    await get().fetchChannels();
  },
}));

// 初始化时尝试恢复当前频道
if (typeof window !== 'undefined') {
  const savedChannelId = safeGetItem('currentChannelId');
  if (savedChannelId) {
    // 延迟执行，等待store初始化完成
    setTimeout(() => {
      const state = useChannelStore.getState();
      const channel = state.channels.find(c => c.id === parseInt(savedChannelId));
      if (channel) {
        state.setCurrentChannel(channel);
      }
    }, 100);
  }
}
