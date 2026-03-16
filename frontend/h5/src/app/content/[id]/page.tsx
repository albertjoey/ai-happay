'use client';

import { useEffect, useState, useRef, useCallback } from 'react';
import { useParams, useRouter } from 'next/navigation';
import Head from 'next/head';
import { getContentDetail, getRelatedContent, getChapterList, getChapterDetail, contentTypeMap, formatNumber, formatDuration, type ContentDetail, type Content, type Chapter, type ChapterDetail } from '@/lib/content';
import InteractionBar from '@/components/InteractionBar';
import CommentSection from '@/components/CommentSection';
import VideoPlayer from '@/components/VideoPlayer';

export default function ContentDetailPage() {
  const params = useParams();
  const router = useRouter();
  const contentId = Number(params.id);
  
  const [content, setContent] = useState<ContentDetail | null>(null);
  const [related, setRelated] = useState<Content[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);
  const [showComments, setShowComments] = useState(false);

  useEffect(() => {
    loadContent();
  }, [contentId]);

  const loadContent = async () => {
    setLoading(true);
    setError(null);
    try {
      const [contentData, relatedData] = await Promise.all([
        getContentDetail(contentId),
        getRelatedContent(contentId, 6),
      ]);
      setContent(contentData);
      setRelated(relatedData);
    } catch (err) {
      setError('加载内容失败');
      console.error(err);
    } finally {
      setLoading(false);
    }
  };

  // 处理返回
  const handleBack = () => {
    // 检查是否有历史记录
    if (window.history.length > 1) {
      router.back();
    } else {
      // 如果没有历史记录,跳转到首页
      router.push('/');
    }
  };

  if (loading) {
    return (
      <div className="min-h-screen bg-white flex items-center justify-center">
        <div className="text-center">
          <div className="w-12 h-12 border-4 border-primary-500 border-t-transparent rounded-full animate-spin mx-auto" />
          <p className="mt-4 text-gray-500">加载中...</p>
        </div>
      </div>
    );
  }

  if (error || !content) {
    return (
      <div className="min-h-screen bg-white flex items-center justify-center">
        <div className="text-center">
          <p className="text-gray-500">{error || '内容不存在'}</p>
          <button 
            onClick={handleBack}
            className="mt-4 px-4 py-2 bg-primary-500 text-white rounded-lg"
          >
            返回
          </button>
        </div>
      </div>
    );
  }

  const typeInfo = contentTypeMap[content.type] || { name: '未知', icon: '📄' };

  // 根据内容类型渲染不同的内容区域
  const renderContent = () => {
    switch (content.type) {
      case 'image_text':
        return <ImageTextContent content={content} />;
      case 'long_video':
      case 'short_video':
      case 'video':
        return <VideoContent content={content} />;
      case 'short_drama':
        return <ShortDramaContent content={content} />;
      case 'manhua':
        return <DramaContent content={content} />;
      case 'novel':
      case 'comic':
        return <NovelContent content={content} />;
      default:
        return <ImageTextContent content={content} />;
    }
  };

  return (
    <div className="min-h-screen bg-white pb-20">
      <Head>
        <title>{content.title} - Happy</title>
        <meta name="description" content={content.description || content.title} />
      </Head>

      {/* 顶部导航 */}
      <header className="sticky top-0 z-50 bg-white border-b border-gray-200">
        <div className="flex items-center justify-between px-4 py-3">
          <button onClick={handleBack} className="p-2 -ml-2">
            <svg className="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M15 19l-7-7 7-7" />
            </svg>
          </button>
          <div className="flex items-center gap-2">
            <span className="text-sm text-gray-500">{typeInfo.icon} {typeInfo.name}</span>
          </div>
          <button className="p-2 -mr-2">
            <svg className="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M12 5v.01M12 12v.01M12 19v.01M12 6a1 1 0 110-2 1 1 0 010 2zm0 7a1 1 0 110-2 1 1 0 010 2zm0 7a1 1 0 110-2 1 1 0 010 2z" />
            </svg>
          </button>
        </div>
      </header>

      {/* 内容区域 */}
      {renderContent()}

      {/* 作者信息 */}
      <div className="px-4 py-3 border-b border-gray-100">
        <div className="flex items-center justify-between">
          <div className="flex items-center gap-3">
            <div className="w-10 h-10 rounded-full bg-gradient-to-br from-primary-400 to-primary-600 flex items-center justify-center text-white font-bold">
              {content.author?.charAt(0) || 'U'}
            </div>
            <div>
              <p className="font-medium text-gray-800">{content.author || '匿名用户'}</p>
              <p className="text-xs text-gray-500">{content.category || '未分类'}</p>
            </div>
          </div>
          <button className="px-4 py-1.5 bg-primary-500 text-white text-sm rounded-full">
            关注
          </button>
        </div>
      </div>

      {/* 描述 */}
      {content.description && (
        <div className="px-4 py-3 border-b border-gray-100">
          <p className="text-gray-700 text-sm leading-relaxed">{content.description}</p>
        </div>
      )}

      {/* 相关推荐 */}
      {related.length > 0 && (
        <div className="px-4 py-4">
          <h3 className="font-semibold text-gray-800 mb-3">相关推荐</h3>
          <div className="grid grid-cols-3 gap-2">
            {related.map((item) => (
              <button
                key={item.id}
                onClick={() => router.push(`/content/${item.id}`)}
                className="text-left"
              >
                <div className="aspect-[3/4] rounded-lg overflow-hidden bg-gray-100">
                  <img
                    src={item.cover_url}
                    alt={item.title}
                    className="w-full h-full object-cover"
                  />
                </div>
                <p className="mt-1 text-xs text-gray-700 line-clamp-2">{item.title}</p>
              </button>
            ))}
          </div>
        </div>
      )}

        {/* 底部互动栏 */}
        <InteractionBar
          contentId={content.id}
          viewCount={content.view_count}
          likeCount={content.like_count}
          commentCount={content.comment_count}
          shareCount={content.share_count}
          collectCount={content.collect_count}
          onComment={() => setShowComments(true)}
        />

        {/* 评论弹窗 */}
        <CommentSection
          contentId={content.id}
          isOpen={showComments}
          onClose={() => setShowComments(false)}
        />
      </div>
    );
  }

  // 图文内容组件
  function ImageTextContent({ content }: { content: ContentDetail }) {
  return (
    <div>
      {/* 封面图 */}
      <div className="aspect-video w-full bg-gray-100">
        <img
          src={content.cover_url}
          alt={content.title}
          className="w-full h-full object-cover"
        />
      </div>
      
      {/* 标题 */}
      <div className="px-4 py-3">
        <h1 className="text-xl font-bold text-gray-900">{content.title}</h1>
        {content.subtitle && (
          <p className="mt-1 text-gray-500 text-sm">{content.subtitle}</p>
        )}
      </div>

      {/* 正文内容（模拟） */}
      <div className="px-4 py-3 prose prose-sm max-w-none">
        <p className="text-gray-700 leading-relaxed">
          {content.description || '这是图文内容的正文部分。在实际应用中，这里会显示完整的图文内容，包括多张图片和文字描述。'}
        </p>
        {/* 模拟多图 */}
        <div className="mt-4 space-y-4">
          <img src={`https://picsum.photos/800/600?random=${content.id}`} alt="内容图片" className="rounded-lg w-full" />
          <p className="text-gray-700 leading-relaxed">
            图文内容支持多张图片展示，用户可以滑动浏览所有图片。每张图片都可以添加文字说明。
          </p>
          <img src={`https://picsum.photos/800/600?random=${content.id + 1}`} alt="内容图片" className="rounded-lg w-full" />
        </div>
      </div>
    </div>
  );
}

// 视频内容组件
function VideoContent({ content }: { content: ContentDetail }) {
  const videoUrl = content.content_url || 'https://test-streams.mux.dev/x36xhzz/x36xhzz.m3u8';
  
  return (
    <div>
      {/* 视频播放器 */}
      <div className="aspect-video w-full bg-black">
        <VideoPlayer
          src={videoUrl}
          poster={content.cover_url}
        />
      </div>

      {/* 标题和信息 */}
      <div className="px-4 py-3">
        <h1 className="text-xl font-bold text-gray-900">{content.title}</h1>
        {content.subtitle && (
          <p className="mt-1 text-gray-500 text-sm">{content.subtitle}</p>
        )}
        <div className="mt-2 flex items-center gap-4 text-sm text-gray-500">
          {content.duration > 0 && (
            <span>{formatDuration(content.duration)}</span>
          )}
          <span>{formatNumber(content.view_count)} 次播放</span>
        </div>
      </div>
    </div>
  );
}

// 漫剧内容组件
function DramaContent({ content }: { content: ContentDetail }) {
  const [currentEpisode, setCurrentEpisode] = useState(0);
  const [currentImage, setCurrentImage] = useState(0);

  // 模拟剧集数据
  const episodes = content.episodes || [
    { id: 1, title: '第1集', images: Array.from({ length: 10 }, (_, i) => `https://picsum.photos/800/1200?random=${content.id}-${i}`) },
    { id: 2, title: '第2集', images: Array.from({ length: 8 }, (_, i) => `https://picsum.photos/800/1200?random=${content.id + 100}-${i}`) },
  ];

  const currentImages = episodes[currentEpisode]?.images || [];

  return (
    <div>
      {/* 图片阅读器 */}
      <div className="relative aspect-[3/4] w-full bg-black">
        <img
          src={currentImages[currentImage] || content.cover_url}
          alt={`${content.title} - ${currentImage + 1}`}
          className="w-full h-full object-contain"
        />
        
        {/* 图片导航 */}
        <div className="absolute bottom-4 left-0 right-0 flex justify-center gap-2">
          {currentImages.map((_, index) => (
            <button
              key={index}
              onClick={() => setCurrentImage(index)}
              className={`w-2 h-2 rounded-full ${index === currentImage ? 'bg-white' : 'bg-white/50'}`}
            />
          ))}
        </div>

        {/* 图片计数 */}
        <div className="absolute top-4 right-4 bg-black/50 text-white text-xs px-2 py-1 rounded">
          {currentImage + 1} / {currentImages.length}
        </div>
      </div>

      {/* 标题 */}
      <div className="px-4 py-3">
        <h1 className="text-xl font-bold text-gray-900">{content.title}</h1>
        <p className="mt-1 text-gray-500 text-sm">{episodes[currentEpisode]?.title}</p>
      </div>

      {/* 剧集选择 */}
      <div className="px-4 py-3 border-t border-gray-100">
        <h3 className="font-medium text-gray-800 mb-2">选集</h3>
        <div className="flex gap-2 overflow-x-auto pb-2">
          {episodes.map((ep, index) => (
            <button
              key={ep.id}
              onClick={() => { setCurrentEpisode(index); setCurrentImage(0); }}
              className={`flex-shrink-0 px-4 py-2 rounded-lg text-sm ${
                index === currentEpisode 
                  ? 'bg-primary-500 text-white' 
                  : 'bg-gray-100 text-gray-700'
              }`}
            >
              {ep.title}
            </button>
          ))}
        </div>
      </div>
    </div>
  );
}

// 小说内容组件
function NovelContent({ content }: { content: ContentDetail }) {
  const [chapters, setChapters] = useState<Chapter[]>([]);
  const [currentChapter, setCurrentChapter] = useState<ChapterDetail | null>(null);
  const [currentChapterIndex, setCurrentChapterIndex] = useState(0);
  const [showReader, setShowReader] = useState(false);
  const [fontSize, setFontSize] = useState(18);
  const [loading, setLoading] = useState(false);
  const readerRef = useRef<HTMLDivElement>(null);
  const currentChapterRef = useRef<ChapterDetail | null>(null);
  const currentChapterIndexRef = useRef(0);

  // 加载章节列表
  useEffect(() => {
    const loadChapters = async () => {
      const list = await getChapterList(content.id);
      setChapters(list);
    };
    loadChapters();
  }, [content.id]);

  // 同步 currentChapter 到 ref
  useEffect(() => {
    currentChapterRef.current = currentChapter;
  }, [currentChapter]);

  // 同步 currentChapterIndex 到 ref
  useEffect(() => {
    currentChapterIndexRef.current = currentChapterIndex;
  }, [currentChapterIndex]);

  // 加载章节内容
  const loadChapter = async (chapterId: number, index: number) => {
    setLoading(true);
    const detail = await getChapterDetail(chapterId);
    if (detail) {
      setCurrentChapter(detail);
      setCurrentChapterIndex(index);
      setShowReader(true);
    }
    setLoading(false);
  };

  // 下一章
  const goNextChapter = () => {
    const chapter = currentChapterRef.current;
    const idx = currentChapterIndexRef.current;
    if (chapter?.next_id) {
      loadChapter(chapter.next_id, idx + 1);
      if (readerRef.current) {
        readerRef.current.scrollTop = 0;
      }
    }
  };

  // 上一章
  const goPrevChapter = () => {
    const chapter = currentChapterRef.current;
    const idx = currentChapterIndexRef.current;
    if (chapter?.prev_id) {
      loadChapter(chapter.prev_id, idx - 1);
      if (readerRef.current) {
        readerRef.current.scrollTop = 0;
      }
    }
  };

  // 处理阅读区域点击 - 点击翻页
  const handleReaderClick = (e: React.MouseEvent<HTMLDivElement>) => {
    const target = e.target as HTMLDivElement;
    const rect = target.getBoundingClientRect();
    const x = e.clientX - rect.left;
    const width = rect.width;
    
    // 点击左侧 1/3 区域：上一章
    if (x < width / 3) {
      if (currentChapter?.prev_id) {
        goPrevChapter();
      }
    }
    // 点击右侧 1/3 区域：下一章
    else if (x > width * 2 / 3) {
      if (currentChapter?.next_id) {
        goNextChapter();
      }
    }
    // 点击中间 1/3 区域：显示/隐藏工具栏（暂不实现）
  };

  return (
    <div>
      {/* 小说封面和信息 */}
      <div className="relative">
        <div className="aspect-[3/4] w-full bg-gradient-to-b from-gray-200 to-gray-300">
          <img
            src={content.cover_url}
            alt={content.title}
            className="w-full h-full object-cover opacity-80"
          />
          <div className="absolute inset-0 bg-gradient-to-t from-black/80 via-black/20 to-transparent" />
        </div>
        
        {/* 标题覆盖层 */}
        <div className="absolute bottom-0 left-0 right-0 p-4 text-white">
          <h1 className="text-2xl font-bold">{content.title}</h1>
          {content.subtitle && (
            <p className="mt-1 text-white/80">{content.subtitle}</p>
          )}
          <div className="mt-2 flex items-center gap-4 text-sm text-white/70">
            <span>{content.word_count || 100000} 字</span>
            <span>{chapters.length} 章</span>
          </div>
        </div>
      </div>

      {/* 开始阅读按钮 */}
      <div className="px-4 py-4">
        <button 
          onClick={() => chapters[0] && loadChapter(chapters[0].id, 0)}
          disabled={chapters.length === 0}
          className="w-full py-3 bg-primary-500 text-white rounded-lg font-medium disabled:bg-gray-300"
        >
          {loading ? '加载中...' : '开始阅读'}
        </button>
      </div>

      {/* 简介 */}
      <div className="px-4 py-3 border-t border-gray-100">
        <h3 className="font-medium text-gray-800 mb-2">简介</h3>
        <p className="text-gray-600 text-sm leading-relaxed">
          {content.description || '这是一部精彩的小说，讲述了主人公的传奇故事。点击开始阅读，探索这个奇妙的世界...'}
        </p>
      </div>

      {/* 章节列表 */}
      <div className="px-4 py-3 border-t border-gray-100">
        <div className="flex items-center justify-between mb-2">
          <h3 className="font-medium text-gray-800">目录</h3>
          <span className="text-sm text-gray-500">共 {chapters.length} 章</span>
        </div>
        <div className="space-y-2 max-h-60 overflow-y-auto">
          {chapters.map((chapter, index) => (
            <button
              key={chapter.id}
              onClick={() => loadChapter(chapter.id, index)}
              className="w-full text-left px-3 py-2 rounded-lg hover:bg-gray-50 flex items-center justify-between"
            >
              <span className="text-gray-700 text-sm">{chapter.title}</span>
              <div className="flex items-center gap-2">
                {chapter.is_free === 0 && <span className="text-xs text-orange-500">付费</span>}
                <span className="text-gray-400 text-xs">{chapter.word_count}字</span>
              </div>
            </button>
          ))}
        </div>
      </div>

      {/* 阅读器弹窗 */}
      {showReader && currentChapter && (
        <div className="fixed inset-0 bg-white z-50">
          <div className="h-full flex flex-col">
            {/* 阅读器头部 */}
            <div className="flex items-center justify-between px-4 py-3 border-b border-gray-200">
              <button onClick={() => setShowReader(false)} className="p-2 -ml-2">
                <svg className="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M6 18L18 6M6 6l12 12" />
                </svg>
              </button>
              <span className="text-sm text-gray-500">{currentChapter.title}</span>
              <div className="w-10" />
            </div>

            {/* 阅读内容 */}
            <div 
              ref={readerRef}
              className="flex-1 overflow-y-auto px-6 py-4 relative" 
              style={{ fontSize: `${fontSize}px`, lineHeight: 1.8 }}
              onClick={handleReaderClick}
            >
              {/* 翻页提示 */}
              <div className="absolute inset-0 pointer-events-none">
                <div className="absolute left-0 top-0 bottom-0 w-1/3 flex items-center justify-center opacity-0 hover:opacity-100 transition-opacity">
                  {currentChapter.prev_id > 0 && (
                    <span className="text-gray-400 text-xs bg-black/10 px-2 py-1 rounded">← 上一章</span>
                  )}
                </div>
                <div className="absolute right-0 top-0 bottom-0 w-1/3 flex items-center justify-center opacity-0 hover:opacity-100 transition-opacity">
                  {currentChapter.next_id > 0 && (
                    <span className="text-gray-400 text-xs bg-black/10 px-2 py-1 rounded">下一章 →</span>
                  )}
                </div>
              </div>
              
              {currentChapter.content.split('\n').map((para, i) => (
                <p key={i} className="text-gray-800 indent-8 mb-4">
                  {para}
                </p>
              ))}
              
              {/* 章节末尾 */}
              <div className="mt-8 text-center">
                {currentChapter.next_id > 0 ? (
                  <button 
                    onClick={(e) => { e.stopPropagation(); goNextChapter(); }}
                    className="px-6 py-3 bg-primary-500 text-white rounded-full text-sm"
                  >
                    下一章
                  </button>
                ) : (
                  <p className="text-gray-400 text-sm">已读完全书</p>
                )}
              </div>
            </div>

            {/* 阅读器底部工具栏 */}
            <div className="border-t border-gray-200 px-4 py-3">
              <div className="flex items-center justify-between">
                <button 
                  onClick={goPrevChapter}
                  disabled={currentChapter.prev_id === 0}
                  className="px-4 py-2 text-gray-600 disabled:text-gray-300"
                >
                  上一章
                </button>
                <div className="flex items-center gap-4">
                  <button 
                    onClick={() => setFontSize(Math.max(14, fontSize - 2))}
                    className="text-gray-600"
                  >
                    A-
                  </button>
                  <span className="text-sm text-gray-500">{fontSize}</span>
                  <button 
                    onClick={() => setFontSize(Math.min(28, fontSize + 2))}
                    className="text-gray-600"
                  >
                    A+
                  </button>
                </div>
                <button 
                  onClick={goNextChapter}
                  disabled={currentChapter.next_id === 0}
                  className="px-4 py-2 text-gray-600 disabled:text-gray-300"
                >
                  下一章
                </button>
              </div>
            </div>
          </div>
        </div>
      )}
    </div>
  );
}

// 短剧内容组件 - 视频剧集
function ShortDramaContent({ content }: { content: ContentDetail }) {
  const [currentEpisode, setCurrentEpisode] = useState(0);

  // 模拟剧集数据
  const episodes = content.episodes || Array.from({ length: content.chapter_count || 10 }, (_, i) => ({
    id: i + 1,
    title: `第${i + 1}集`,
    video_url: `https://test-streams.mux.dev/x36xhzz/x36xhzz.m3u8`,
    duration: 180 + Math.floor(Math.random() * 120)
  }));

  const currentVideo = episodes[currentEpisode];

  return (
    <div>
      {/* 视频播放器 */}
      <div className="aspect-video w-full bg-black">
        <VideoPlayer
          src={currentVideo?.video_url || 'https://test-streams.mux.dev/x36xhzz/x36xhzz.m3u8'}
          poster={content.cover_url}
        />
      </div>

      {/* 标题和信息 */}
      <div className="px-4 py-3">
        <h1 className="text-xl font-bold text-gray-900">{content.title}</h1>
        <p className="mt-1 text-gray-500 text-sm">{currentVideo?.title}</p>
        <div className="mt-2 flex items-center gap-4 text-sm text-gray-500">
          <span>{formatNumber(content.view_count)} 次播放</span>
          <span>共 {episodes.length} 集</span>
        </div>
      </div>

      {/* 剧集选择 */}
      <div className="px-4 py-3 border-t border-gray-100">
        <h3 className="font-medium text-gray-800 mb-2">选集</h3>
        <div className="grid grid-cols-5 gap-2">
          {episodes.map((ep, index) => (
            <button
              key={ep.id}
              onClick={() => setCurrentEpisode(index)}
              className={`px-3 py-2 rounded-lg text-sm ${
                index === currentEpisode 
                  ? 'bg-primary-500 text-white' 
                  : 'bg-gray-100 text-gray-700'
              }`}
            >
              {ep.title}
            </button>
          ))}
        </div>
      </div>
    </div>
  );
}
