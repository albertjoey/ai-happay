'use client';

import Head from 'next/head';
import { useEffect, useState } from 'react';
import ChannelTabs from '@/components/ChannelTabs';
import BannerSection from '@/components/BannerSection';
import DiamondGrid from '@/components/DiamondGrid';
import RecommendSection from '@/components/RecommendSection';
import FeedSection from '@/components/FeedSection';
import { useChannelStore } from '@/store/channelStore';

export default function Home() {
  const { currentChannel, channelConfig, fetchChannelConfig, channels, fetchChannels, setCurrentChannel } = useChannelStore();
  const [mounted, setMounted] = useState(false);

  // 初始化加载频道
  useEffect(() => {
    setMounted(true);
    fetchChannels();
  }, [fetchChannels]);

  // 当频道列表加载完成后，自动选择第一个频道
  useEffect(() => {
    if (channels.length > 0 && !currentChannel) {
      setCurrentChannel(channels[0]);
    }
  }, [channels, currentChannel, setCurrentChannel]);

  // 当频道切换时加载配置
  useEffect(() => {
    if (currentChannel) {
      fetchChannelConfig(currentChannel.id);
    }
  }, [currentChannel?.id, fetchChannelConfig]);

  // 获取页面配置
  const pageConfig = channelConfig?.page_config;
  const showBanner = pageConfig?.banner?.enabled !== false;
  const showDiamond = pageConfig?.diamond?.enabled !== false;
  const showFeed = pageConfig?.feed?.enabled !== false;
  const recommends = pageConfig?.recommends || [];
  const feedConfig = pageConfig?.feed;

  // 服务端渲染时不显示内容
  if (!mounted) {
    return (
      <div className="min-h-screen bg-gray-50 flex items-center justify-center">
        <div className="text-gray-400">加载中...</div>
      </div>
    );
  }

  return (
    <div className="min-h-screen bg-gray-50">
      <Head>
        <title>Happy - 多内容平台</title>
        <meta name="description" content="包含长短视频、短剧、漫剧、小说、图文等多种内容形态" />
      </Head>

      {/* 顶部搜索栏 */}
      <header className="sticky top-0 z-50 bg-white shadow-sm">
        <div className="px-4 py-3">
          <div className="flex items-center gap-3">
            <div className="flex-1">
              <input
                type="text"
                placeholder="搜索视频、短剧、漫剧、小说..."
                className="w-full px-4 py-2 bg-gray-100 rounded-full text-sm focus:outline-none"
              />
            </div>
            <button className="p-2">
              <svg className="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
              </svg>
            </button>
          </div>
        </div>

        {/* 动态频道列表 */}
        <ChannelTabs />
      </header>

      {/* 内容区域 - 根据频道配置动态显示 */}
      <main className="pb-16">
        {currentChannel ? (
          <>
            {/* Banner区域 - 根据配置显示 */}
            {showBanner && (
              <BannerSection 
                channelId={currentChannel.id}
                bannerIds={pageConfig?.banner?.banner_ids}
              />
            )}

            {/* 金刚位区域 - 根据配置显示 */}
            {showDiamond && (
              <div className="bg-white">
                <DiamondGrid 
                  channelId={currentChannel.id} 
                  groupIds={pageConfig?.diamond?.group_ids}
                />
              </div>
            )}

            {/* 推荐位区域 - 根据配置显示多个推荐位 */}
            {recommends.length > 0 ? (
              recommends.map((rec, index) => (
                <div key={rec.id || index} className="mt-2">
                  <RecommendSection 
                    channelId={currentChannel.id}
                    recommendId={rec.id}
                    customTitle={rec.title}
                  />
                </div>
              ))
            ) : (
              /* 默认推荐位 */
              <div className="mt-2">
                <RecommendSection channelId={currentChannel.id} />
              </div>
            )}

            {/* Feed流区域 - 根据配置显示 */}
            {showFeed && (
              <div className="mt-2">
                <FeedSection 
                  channelId={currentChannel.id}
                  feedId={feedConfig?.feed_id}
                  autoLoad={feedConfig?.auto_load}
                  showTitle={feedConfig?.show_title}
                />
              </div>
            )}
          </>
        ) : (
          /* 默认展示 - 当没有选择频道时 */
          <div className="px-4 py-8">
            <div className="text-center text-gray-400">
              <p>加载频道中...</p>
            </div>
          </div>
        )}
      </main>

        {/* 底部导航 */}
        <nav className="fixed bottom-0 left-0 right-0 bg-white border-t border-gray-200 z-50">
          <div className="flex items-center justify-around py-2">
            <button className="flex flex-col items-center text-blue-500">
              <svg className="w-6 h-6" fill="currentColor" viewBox="0 0 24 24">
                <path d="M10 20v-6h4v6h5v-8h3L12 3 2 12h3v8z" />
              </svg>
              <span className="text-xs mt-1">首页</span>
            </button>
            <button className="flex flex-col items-center text-gray-400">
              <svg className="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M7 4v16M17 4v16M3 8h4m10 0h4M3 12h18M3 16h4m10 0h4M4 20h16a1 1 0 001-1V5a1 1 0 00-1-1H4a1 1 0 00-1 1v14a1 1 0 001 1z" />
              </svg>
              <span className="text-xs mt-1">分类</span>
            </button>
            {/* 发布按钮 */}
            <button 
              onClick={() => window.location.href = "/publish"}
              className="flex flex-col items-center -mt-4"
            >
              <div className="w-12 h-12 bg-gradient-to-br from-blue-400 to-blue-600 rounded-full flex items-center justify-center shadow-lg">
                <svg className="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M12 4v16m8-8H4" />
                </svg>
              </div>
              <span className="text-xs mt-1 text-gray-600">发布</span>
            </button>
            <button className="flex flex-col items-center text-gray-400">
              <svg className="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M12 8v13m0-13V6a2 2 0 112 2h-2zm0 0V5.5A2.5 2.5 0 109.5 8H12zm-7 4h14M5 12a2 2 0 110-4h14a2 2 0 110 4M5 12v7a2 2 0 002 2h10a2 2 0 002-2v-7" />
              </svg>
              <span className="text-xs mt-1">消息</span>
            </button>
            <button className="flex flex-col items-center text-gray-400">
              <svg className="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z" />
              </svg>
              <span className="text-xs mt-1">我的</span>
            </button>
          </div>
        </nav>
    </div>
  );
}
