'use client';

import { useState, useEffect } from 'react';
import Link from 'next/link';
import { getDiscoverPage, DiscoverModule } from '@/lib/discoverApi';

export default function DiscoverPage() {
  const [activeTab, setActiveTab] = useState<'day' | 'week'>('day');
  const [followedCreators, setFollowedCreators] = useState<number[]>([]);
  const [modules, setModules] = useState<DiscoverModule[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string>('');

  useEffect(() => {
    loadDiscoverData();
  }, []);

  const loadDiscoverData = async () => {
    setLoading(true);
    setError('');
    try {
      console.log('开始加载发现页数据...');
      const data = await getDiscoverPage();
      console.log('发现页数据:', data);
      console.log('模块数量:', data.modules?.length);
      
      if (data && data.modules && data.modules.length > 0) {
        console.log('设置modules状态,数量:', data.modules.length);
        setModules(data.modules);
        console.log('数据设置完成');
      } else {
        console.log('没有获取到数据或数据为空');
        setError('没有获取到数据');
      }
    } catch (error: any) {
      console.error('加载发现页数据失败:', error);
      setError('加载失败: ' + (error.message || '未知错误'));
    } finally {
      setLoading(false);
      console.log('加载完成,loading设置为false');
    }
  };

  const toggleFollow = (creatorId: number) => {
    if (followedCreators.includes(creatorId)) {
      setFollowedCreators(followedCreators.filter(id => id !== creatorId));
    } else {
      setFollowedCreators([...followedCreators, creatorId]);
    }
  };

  // 根据模块名获取数据
  const getModuleData = (moduleName: string) => {
    const module = modules.find(m => m.module === moduleName);
    return module?.items || [];
  };

  const hotTopics = getModuleData('hot_topics');
  const hotRankItems = getModuleData('hot_rank');
  const recommendCreators = getModuleData('recommend_creators');
  const guessYouLike = getModuleData('guess_you_like');

  return (
    <div className="min-h-screen bg-gray-50 pb-20">
      {/* 顶部标题 */}
      <header className="sticky top-0 bg-white z-40 px-4 py-3 border-b border-gray-100">
        <h1 className="text-xl font-bold text-center">发现</h1>
      </header>

      {loading ? (
        <div className="flex items-center justify-center py-20">
          <div className="text-gray-400">加载中...</div>
        </div>
      ) : modules.length === 0 ? (
        <div className="flex flex-col items-center justify-center py-20">
          <div className="text-gray-400 mb-4">暂无内容</div>
          <button
            onClick={loadDiscoverData}
            className="px-4 py-2 bg-blue-500 text-white rounded"
          >
            重新加载
          </button>
        </div>
      ) : (
        <>
          {/* 热门话题 */}
          {hotTopics.length > 0 && (
            <section className="bg-white mt-2 p-4">
              <div className="flex items-center justify-between mb-3">
                <h2 className="text-lg font-bold">🔥 热门话题</h2>
                <button className="text-sm text-gray-500">更多 &gt;</button>
              </div>
              <div className="flex gap-3 overflow-x-auto pb-2 [-ms-overflow-style:none] [scrollbar-width:none] [&::-webkit-scrollbar]:hidden">
                {hotTopics.map((topic: any) => (
                  <Link 
                    key={topic.id} 
                    href={`/topic/${topic.id}`}
                    className="flex-shrink-0 w-24 cursor-pointer"
                  >
                    <div 
                      className="w-24 h-24 rounded-lg mb-1 overflow-hidden relative"
                    >
                      {topic.cover_url ? (
                        <img 
                          src={topic.cover_url} 
                          alt={topic.title}
                          className="w-full h-full object-cover"
                        />
                      ) : (
                        <div 
                          className="w-full h-full flex items-center justify-center text-white text-2xl font-bold"
                          style={{ backgroundColor: topic.color || '#FF6B6B' }}
                        >
                          {topic.title.slice(0, 2)}
                        </div>
                      )}
                    </div>
                    <p className="text-xs font-medium truncate">{topic.title}</p>
                    <p className="text-xs text-gray-400">{topic.count || '0'}参与</p>
                  </Link>
                ))}
              </div>
            </section>
          )}

          {/* 热门榜单 */}
          {hotRankItems.length > 0 && (
            <section className="bg-white mt-2 p-4">
              <div className="flex items-center justify-between mb-3">
                <h2 className="text-lg font-bold">🏆 热门榜单</h2>
                <div className="flex gap-2">
                  <button 
                    onClick={() => setActiveTab('day')}
                    className={`px-3 py-1 text-xs rounded-full ${activeTab === 'day' ? 'bg-blue-500 text-white' : 'bg-gray-100 text-gray-600'}`}
                  >
                    日榜
                  </button>
                  <button 
                    onClick={() => setActiveTab('week')}
                    className={`px-3 py-1 text-xs rounded-full ${activeTab === 'week' ? 'bg-blue-500 text-white' : 'bg-gray-100 text-gray-600'}`}
                  >
                    周榜
                  </button>
                </div>
              </div>
              <div className="space-y-3">
                {hotRankItems.map((item: any, index: number) => (
                  <Link 
                    key={item.id} 
                    href={`/content/${item.id}`}
                    className="flex items-center gap-3 cursor-pointer"
                  >
                    <span className={`w-6 h-6 rounded-full flex items-center justify-center text-xs font-bold ${
                      index < 3 ? 'bg-gradient-to-br from-yellow-400 to-orange-500 text-white' : 'bg-gray-200 text-gray-600'
                    }`}>
                      {index + 1}
                    </span>
                    <div className="w-16 h-20 rounded flex-shrink-0 overflow-hidden">
                      {item.cover_url ? (
                        <img 
                          src={item.cover_url} 
                          alt={item.title}
                          className="w-full h-full object-cover"
                        />
                      ) : (
                        <div 
                          className="w-full h-full flex items-center justify-center text-white text-lg font-bold"
                          style={{ backgroundColor: item.color || '#FF6B6B' }}
                        >
                          {item.title.slice(0, 2)}
                        </div>
                      )}
                    </div>
                    <div className="flex-1 min-w-0">
                      <p className="font-medium truncate">{item.title}</p>
                      <p className="text-xs text-gray-500">{item.author || '未知作者'}</p>
                      <p className="text-xs text-gray-400">{item.views || '0'}次播放</p>
                    </div>
                  </Link>
                ))}
              </div>
            </section>
          )}

          {/* 推荐创作者 */}
          {recommendCreators.length > 0 && (
            <section className="bg-white mt-2 p-4">
              <div className="flex items-center justify-between mb-3">
                <h2 className="text-lg font-bold">⭐ 推荐创作者</h2>
                <button className="text-sm text-gray-500">换一批</button>
              </div>
              <div className="flex gap-4 overflow-x-auto pb-2 [-ms-overflow-style:none] [scrollbar-width:none] [&::-webkit-scrollbar]:hidden">
                {recommendCreators.map((creator: any) => (
                  <div key={creator.id} className="flex-shrink-0 w-28 text-center">
                    <div className="w-16 h-16 rounded-full mx-auto mb-2 overflow-hidden">
                      {creator.cover_url ? (
                        <img 
                          src={creator.cover_url} 
                          alt={creator.title}
                          className="w-full h-full object-cover"
                        />
                      ) : (
                        <div 
                          className="w-full h-full flex items-center justify-center text-3xl"
                          style={{ backgroundColor: (creator.color || '#FF6B6B') + '30', border: `2px solid ${creator.color || '#FF6B6B'}` }}
                        >
                          {creator.avatar || '👤'}
                        </div>
                      )}
                    </div>
                    <p className="text-sm font-medium truncate">{creator.title}</p>
                    <p className="text-xs text-gray-400 truncate">{creator.desc || ''}</p>
                    <p className="text-xs text-gray-400">{creator.fans || '0'}粉丝</p>
                    <button 
                      onClick={() => toggleFollow(creator.id)}
                      className={`mt-2 px-4 py-1 text-xs rounded-full ${
                        followedCreators.includes(creator.id) 
                          ? 'bg-gray-200 text-gray-600' 
                          : 'bg-blue-500 text-white'
                      }`}
                    >
                      {followedCreators.includes(creator.id) ? '已关注' : '关注'}
                    </button>
                  </div>
                ))}
              </div>
            </section>
          )}

          {/* 猜你喜欢 */}
          {guessYouLike.length > 0 && (
            <section className="bg-white mt-2 p-4">
              <div className="flex items-center justify-between mb-3">
                <h2 className="text-lg font-bold">💡 猜你喜欢</h2>
                <button onClick={loadDiscoverData} className="text-sm text-gray-500">刷新</button>
              </div>
              <div className="grid grid-cols-2 gap-3">
                {guessYouLike.map((item: any) => (
                  <Link 
                    key={item.id} 
                    href={`/content/${item.id}`}
                    className="bg-gray-50 rounded-lg overflow-hidden cursor-pointer"
                  >
                    <div className="aspect-[3/4] overflow-hidden">
                      {item.cover_url ? (
                        <img 
                          src={item.cover_url} 
                          alt={item.title}
                          className="w-full h-full object-cover"
                        />
                      ) : (
                        <div 
                          className="w-full h-full flex items-center justify-center text-white text-xl font-bold"
                          style={{ backgroundColor: item.color || '#FF6B6B' }}
                        >
                          {item.title.slice(0, 2)}
                        </div>
                      )}
                    </div>
                    <div className="p-2">
                      <p className="text-sm font-medium truncate">{item.title}</p>
                      <div className="flex items-center justify-between mt-1">
                        <span className="text-xs text-gray-500">{item.author || '未知'}</span>
                        <span className="text-xs text-gray-400">{item.views || '0'}次</span>
                      </div>
                    </div>
                  </Link>
                ))}
              </div>
            </section>
          )}
        </>
      )}

      {/* 底部导航 */}
      <nav className="fixed bottom-0 left-0 right-0 bg-white border-t border-gray-200 z-50">
        <div className="flex items-center justify-around py-2">
          <button 
            onClick={() => window.location.href = '/'}
            className="flex flex-col items-center text-gray-400"
          >
            <svg className="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M3 12l2-2m0 0l7-7 7 7M5 10v10a1 1 0 001 1h3m10-11l2 2m-2-2v10a1 1 0 01-1 1h-3m-6 0a1 1 0 001-1v-4a1 1 0 011-1h2a1 1 0 011 1v4a1 1 0 001 1m-6 0h6" />
            </svg>
            <span className="text-xs mt-1">首页</span>
          </button>
          <button className="flex flex-col items-center text-blue-500">
            <svg className="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
            </svg>
            <span className="text-xs mt-1">发现</span>
          </button>
          <button 
            onClick={() => window.location.href = '/publish'}
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
              <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M8 12h.01M12 12h.01M16 12h.01M21 12c0 4.418-4.03 8-9 8a9.863 9.863 0 01-4.255-.949L3 20l1.395-3.72C3.512 15.042 3 13.574 3 12c0-4.418 4.03-8 9-8s9 3.582 9 8z" />
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
