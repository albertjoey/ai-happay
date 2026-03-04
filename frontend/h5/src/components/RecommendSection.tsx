'use client';

import { useEffect, useState } from 'react';
import { useRouter } from 'next/navigation';
import { getRecommendList, type Recommend } from '@/lib/configApi';

interface RecommendSectionProps {
  channelId: number;
  recommendId?: number;
  customTitle?: string;
}

export default function RecommendSection({ channelId, recommendId, customTitle }: RecommendSectionProps) {
  const router = useRouter();
  const [recommends, setRecommends] = useState<Recommend[]>([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    loadRecommends();
  }, [channelId, recommendId]);

  const loadRecommends = async () => {
    setLoading(true);
    const data = await getRecommendList(channelId);
    
    // 如果指定了推荐位ID，只显示该推荐位
    if (recommendId) {
      const filtered = data.filter(r => r.id === recommendId);
      setRecommends(filtered);
    } else {
      setRecommends(data);
    }
    setLoading(false);
  };

  if (loading) {
    return (
      <div className="px-4 py-3">
        <div className="h-6 bg-gray-200 rounded w-24 mb-3 animate-pulse" />
        <div className="h-48 bg-gray-200 rounded animate-pulse" />
      </div>
    );
  }

  if (recommends.length === 0) {
    return null;
  }

  // 根据display_type渲染不同样式
  const renderRecommend = (recommend: Recommend) => {
    switch (recommend.display_type) {
      case 'single':
        return renderSingleRecommend(recommend);
      case 'scroll':
        return renderScrollRecommend(recommend);
      case 'grid':
        return renderGridRecommend(recommend);
      default:
        return null;
    }
  };

  // 单个大图
  const renderSingleRecommend = (recommend: Recommend) => {
    const material = (recommend as any).materials?.[0];
    return (
      <div key={recommend.id} className="mb-4 px-4">
        <h3 className="text-base font-semibold text-gray-800 mb-3">{customTitle || recommend.title}</h3>
        <div 
          className="relative aspect-[16/9] rounded-lg overflow-hidden bg-gray-200 cursor-pointer"
          onClick={() => material && router.push(`/content/${material.id}`)}
        >
          {material ? (
            <>
              <img src={material.cover_url} alt={material.title} className="w-full h-full object-cover" />
              <div className="absolute bottom-0 left-0 right-0 bg-gradient-to-t from-black/60 to-transparent p-4">
                <p className="text-white font-medium">{material.title}</p>
              </div>
            </>
          ) : (
            <div className="w-full h-full bg-gradient-to-r from-primary-400 to-primary-600 flex items-center justify-center text-white">
              <p className="text-lg font-medium">{recommend.description}</p>
            </div>
          )}
        </div>
      </div>
    );
  };

  // 横向滑动
  const renderScrollRecommend = (recommend: Recommend) => {
    const materials = (recommend as any).materials || [];
    return (
      <div key={recommend.id} className="mb-4">
        <h3 className="text-base font-semibold text-gray-800 mb-3 px-4">{customTitle || recommend.title}</h3>
        <div className="flex overflow-x-auto gap-3 pb-2 px-4 scrollbar-hide">
          {materials.map((material: any) => (
            <div
              key={material.id}
              className="flex-shrink-0 w-32 aspect-[3/4] rounded-lg overflow-hidden bg-gray-200 cursor-pointer"
              onClick={() => router.push(`/content/${material.id}`)}
            >
              <div className="relative w-full h-full">
                <img 
                  src={material.cover_url} 
                  alt={material.title}
                  className="w-full h-full object-cover"
                />
                <div className="absolute bottom-0 left-0 right-0 bg-gradient-to-t from-black/60 to-transparent p-2">
                  <p className="text-white text-xs line-clamp-2">{material.title}</p>
                </div>
              </div>
            </div>
          ))}
          {materials.length === 0 && [1, 2, 3, 4, 5].map((i) => (
            <div
              key={i}
              className="flex-shrink-0 w-32 aspect-[3/4] rounded-lg overflow-hidden bg-gray-200"
            >
              <div className="w-full h-full bg-gradient-to-br from-blue-400 to-blue-600 flex items-center justify-center text-white text-sm">
                内容 {i}
              </div>
            </div>
          ))}
        </div>
      </div>
    );
  };

  // 网格布局
  const renderGridRecommend = (recommend: Recommend) => {
    const materials = (recommend as any).materials || [];
    return (
      <div key={recommend.id} className="mb-4">
        <h3 className="text-base font-semibold text-gray-800 mb-3 px-4">{customTitle || recommend.title}</h3>
        <div className="grid grid-cols-3 gap-2 px-4">
          {materials.map((material: any) => (
            <div
              key={material.id}
              className="aspect-square rounded-lg overflow-hidden bg-gray-200 cursor-pointer"
              onClick={() => router.push(`/content/${material.id}`)}
            >
              <div className="relative w-full h-full">
                <img 
                  src={material.cover_url} 
                  alt={material.title}
                  className="w-full h-full object-cover"
                />
                <div className="absolute bottom-0 left-0 right-0 bg-gradient-to-t from-black/60 to-transparent p-1">
                  <p className="text-white text-xs line-clamp-1">{material.title}</p>
                </div>
              </div>
            </div>
          ))}
          {materials.length === 0 && [1, 2, 3, 4, 5, 6].map((i) => (
            <div
              key={i}
              className="aspect-square rounded-lg overflow-hidden bg-gray-200"
            >
              <div className="w-full h-full bg-gradient-to-br from-green-400 to-green-600 flex items-center justify-center text-white text-sm">
                {i}
              </div>
            </div>
          ))}
        </div>
      </div>
    );
  };

  return <div className="py-3">{recommends.map(renderRecommend)}</div>;
}
