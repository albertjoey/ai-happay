'use client';

import { useState } from 'react';

// 热门话题数据
const hotTopics = [
  { id: 1, name: '周末美食打卡', count: '12.5万', color: '#FF6B6B' },
  { id: 2, name: '旅行日记', count: '8.3万', color: '#4ECDC4' },
  { id: 3, name: '宠物萌照', count: '6.7万', color: '#45B7D1' },
  { id: 4, name: '健身打卡', count: '5.2万', color: '#96CEB4' },
  { id: 5, name: '穿搭分享', count: '4.8万', color: '#FFEAA7' },
  { id: 6, name: '游戏日常', count: '3.9万', color: '#DDA0DD' },
];

// 热门榜单数据
const hotRankItems = [
  { id: 1, title: '今日份美食分享', author: '美食达人小王', views: '128万', color: '#FF6B6B' },
  { id: 2, title: '周末旅行vlog', author: '旅行者小李', views: '96万', color: '#4ECDC4' },
  { id: 3, title: '我家猫咪的一天', author: '猫奴小张', views: '85万', color: '#45B7D1' },
  { id: 4, title: '健身30天挑战', author: '健身教练', views: '72万', color: '#96CEB4' },
  { id: 5, title: '春季穿搭指南', author: '时尚博主', views: '68万', color: '#FFEAA7' },
];

// 推荐创作者
const recommendCreators = [
  { id: 1, name: '美食达人小王', avatar: '👨‍🍳', fans: '52.3万', desc: '分享美食日常', color: '#FF6B6B' },
  { id: 2, name: '旅行者小李', avatar: '🌍', fans: '38.7万', desc: '环游世界ing', color: '#4ECDC4' },
  { id: 3, name: '猫奴小张', avatar: '🐱', fans: '29.1万', desc: '两只猫的铲屎官', color: '#45B7D1' },
  { id: 4, name: '健身教练', avatar: '💪', fans: '45.6万', desc: '专业健身指导', color: '#96CEB4' },
];

// 猜你喜欢数据
const guessYouLike = [
  { id: 1, title: '超简单的家常菜做法', author: '美食达人', views: '23万', color: '#FF6B6B' },
  { id: 2, title: '一个人的旅行', author: '旅行者', views: '18万', color: '#4ECDC4' },
  { id: 3, title: '猫咪搞笑瞬间', author: '猫奴', views: '45万', color: '#45B7D1' },
  { id: 4, title: '居家健身教程', author: '健身教练', views: '32万', color: '#96CEB4' },
  { id: 5, title: '春季穿搭推荐', author: '时尚博主', views: '28万', color: '#FFEAA7' },
  { id: 6, title: '游戏精彩操作', author: '游戏玩家', views: '56万', color: '#DDA0DD' },
];

export default function DiscoverPage() {
  const [activeTab, setActiveTab] = useState<'day' | 'week'>('day');
  const [followedCreators, setFollowedCreators] = useState<number[]>([]);

  const toggleFollow = (creatorId: number) => {
    if (followedCreators.includes(creatorId)) {
      setFollowedCreators(followedCreators.filter(id => id !== creatorId));
    } else {
      setFollowedCreators([...followedCreators, creatorId]);
    }
  };

  return (
    <div className="min-h-screen bg-gray-50 pb-20">
      {/* 顶部标题 */}
      <header className="sticky top-0 bg-white z-40 px-4 py-3 border-b border-gray-100">
        <h1 className="text-xl font-bold text-center">发现</h1>
      </header>

      {/* 热门话题 */}
      <section className="bg-white mt-2 p-4">
        <div className="flex items-center justify-between mb-3">
          <h2 className="text-lg font-bold">🔥 热门话题</h2>
          <button className="text-sm text-gray-500">更多 &gt;</button>
        </div>
        <div className="flex gap-3 overflow-x-auto pb-2 [-ms-overflow-style:none] [scrollbar-width:none] [&::-webkit-scrollbar]:hidden">
          {hotTopics.map(topic => (
            <div key={topic.id} className="flex-shrink-0 w-24">
              <div 
                className="w-24 h-24 rounded-lg mb-1 flex items-center justify-center text-white text-2xl font-bold"
                style={{ backgroundColor: topic.color }}
              >
                {topic.name.slice(0, 2)}
              </div>
              <p className="text-xs font-medium truncate">{topic.name}</p>
              <p className="text-xs text-gray-400">{topic.count}参与</p>
            </div>
          ))}
        </div>
      </section>

      {/* 热门榜单 */}
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
          {hotRankItems.map((item, index) => (
            <div key={item.id} className="flex items-center gap-3">
              <span className={`w-6 h-6 rounded-full flex items-center justify-center text-xs font-bold ${
                index < 3 ? 'bg-gradient-to-br from-yellow-400 to-orange-500 text-white' : 'bg-gray-200 text-gray-600'
              }`}>
                {index + 1}
              </span>
              <div 
                className="w-16 h-20 rounded flex-shrink-0 flex items-center justify-center text-white text-lg font-bold"
                style={{ backgroundColor: item.color }}
              >
                {item.title.slice(0, 2)}
              </div>
              <div className="flex-1 min-w-0">
                <p className="font-medium truncate">{item.title}</p>
                <p className="text-xs text-gray-500">{item.author}</p>
                <p className="text-xs text-gray-400">{item.views}次播放</p>
              </div>
            </div>
          ))}
        </div>
      </section>

      {/* 推荐创作者 */}
      <section className="bg-white mt-2 p-4">
        <div className="flex items-center justify-between mb-3">
          <h2 className="text-lg font-bold">⭐ 推荐创作者</h2>
          <button className="text-sm text-gray-500">换一批</button>
        </div>
        <div className="flex gap-4 overflow-x-auto pb-2 [-ms-overflow-style:none] [scrollbar-width:none] [&::-webkit-scrollbar]:hidden">
          {recommendCreators.map(creator => (
            <div key={creator.id} className="flex-shrink-0 w-28 text-center">
              <div 
                className="w-16 h-16 rounded-full mx-auto mb-2 flex items-center justify-center text-3xl"
                style={{ backgroundColor: creator.color + '30', border: `2px solid ${creator.color}` }}
              >
                {creator.avatar}
              </div>
              <p className="text-sm font-medium truncate">{creator.name}</p>
              <p className="text-xs text-gray-400 truncate">{creator.desc}</p>
              <p className="text-xs text-gray-400">{creator.fans}粉丝</p>
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

      {/* 猜你喜欢 */}
      <section className="bg-white mt-2 p-4">
        <div className="flex items-center justify-between mb-3">
          <h2 className="text-lg font-bold">💡 猜你喜欢</h2>
          <button className="text-sm text-gray-500">刷新</button>
        </div>
        <div className="grid grid-cols-2 gap-3">
          {guessYouLike.map(item => (
            <div key={item.id} className="bg-gray-50 rounded-lg overflow-hidden">
              <div 
                className="aspect-[3/4] flex items-center justify-center text-white text-xl font-bold"
                style={{ backgroundColor: item.color }}
              >
                {item.title.slice(0, 2)}
              </div>
              <div className="p-2">
                <p className="text-sm font-medium truncate">{item.title}</p>
                <div className="flex items-center justify-between mt-1">
                  <span className="text-xs text-gray-500">{item.author}</span>
                  <span className="text-xs text-gray-400">{item.views}次</span>
                </div>
              </div>
            </div>
          ))}
        </div>
      </section>

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
