'use client';

import { useEffect, useState } from 'react';
import { useRouter } from 'next/navigation';
import { getFeedConfigList, type FeedConfig } from '@/lib/configApi';
import axios from 'axios';

const API_BASE = process.env.NEXT_PUBLIC_API_BASE || 'http://localhost:4004';

interface AdPosition {
  id: number;
  channel_id: number;
  name: string;
  code: string;
  type: number;
  image_url: string;
  link_url: string;
  status: number;
  sort: number;
}

interface FeedSectionProps {
  channelId: number;
  feedId?: number;
  autoLoad?: boolean;
  showTitle?: boolean;
}

export default function FeedSection({ channelId, feedId, autoLoad = true, showTitle = true }: FeedSectionProps) {
  const router = useRouter();
  const [feedConfigs, setFeedConfigs] = useState<FeedConfig[]>([]);
  const [adPositions, setAdPositions] = useState<AdPosition[]>([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    loadFeedItems();
  }, [channelId, feedId]);

  const loadFeedItems = async () => {
    setLoading(true);
    try {
      const [feedData, adData] = await Promise.all([
        getFeedConfigList(channelId),
        loadAdPositions(channelId)
      ]);
      
      // 如果指定了feedId，只显示该feed
      if (feedId) {
        const filtered = feedData.filter(f => f.id === feedId);
        setFeedConfigs(filtered);
      } else {
        setFeedConfigs(feedData);
      }
      
      setAdPositions(adData);
    } catch (error) {
      console.error('加载Feed失败:', error);
    } finally {
      setLoading(false);
    }
  };

  const loadAdPositions = async (channelId: number): Promise<AdPosition[]> => {
    try {
      const response = await axios.get(`${API_BASE}/api/v1/ad-slot/list`, {
        params: { channel_id: channelId, status: 1 }
      });
      return response.data.list || [];
    } catch (error) {
      console.error('加载广告位失败:', error);
      return [];
    }
  };

  // 在Feed流中插入广告
  const insertAdsIntoFeed = (materials: any[], ads: AdPosition[]) => {
    if (ads.length === 0 || materials.length === 0) return materials;
    
    const result: any[] = [];
    const adInterval = 4; // 每4个内容插入1个广告
    
    materials.forEach((material, index) => {
      result.push(material);
      
      // 每4个内容后插入广告
      if ((index + 1) % adInterval === 0 && ads.length > 0) {
        const adIndex = Math.floor(index / adInterval) % ads.length;
        result.push({
          ...ads[adIndex],
          isAd: true,
          id: `ad_${ads[adIndex].id}_${index}`
        });
      }
    });
    
    return result;
  };

  if (loading) {
    return (
      <div className="px-4 py-3">
        <div className="grid grid-cols-2 gap-3">
          {[1, 2, 3, 4].map((i) => (
            <div key={i} className="bg-white rounded-lg overflow-hidden shadow-sm animate-pulse">
              <div className="aspect-[3/4] bg-gray-200" />
              <div className="p-3">
                <div className="h-4 bg-gray-200 rounded mb-2" />
                <div className="h-3 bg-gray-200 rounded w-1/2" />
              </div>
            </div>
          ))}
        </div>
      </div>
    );
  }

  if (feedConfigs.length === 0) {
    return null;
  }

  // 渲染Feed卡片
  const renderFeedCard = (material: any) => (
    <div
      key={material.id}
      className="bg-white rounded-lg overflow-hidden shadow-sm cursor-pointer"
      onClick={() => !material.isAd && router.push(`/content/${material.id}`)}
    >
      <div className="relative aspect-[3/4]">
        <img
          src={material.cover_url || material.image_url}
          alt={material.title || material.name}
          className="w-full h-full object-cover"
        />
        <div className="absolute bottom-0 left-0 right-0 bg-gradient-to-t from-black/60 to-transparent p-3">
          <p className="text-white text-sm font-medium line-clamp-2">
            {material.title || material.name}
          </p>
          {material.isAd && (
            <span className="absolute top-2 right-2 bg-yellow-400 text-xs px-2 py-1 rounded">
              广告
            </span>
          )}
        </div>
      </div>
      <div className="p-3">
        <div className="flex items-center gap-2">
          <div className="w-6 h-6 rounded-full bg-gradient-to-br from-blue-400 to-blue-600 flex items-center justify-center text-white text-xs">
            {material.author?.charAt(0) || 'A'}
          </div>
          <span className="text-xs text-gray-600 flex-1 truncate">
            {material.author || '广告主'}
          </span>
        </div>
        <div className="flex items-center gap-3 mt-2 text-xs text-gray-500">
          <span className="flex items-center gap-1">
            <svg className="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
              <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z" />
            </svg>
            {material.view_count || 0}
          </span>
          <span className="flex items-center gap-1">
            <svg className="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M4.318 6.318a4.5 4.5 0 000 6.364L12 20.364l7.682-7.682a4.5 4.5 0 00-6.364-6.364L12 7.636l-1.318-1.318a4.5 4.5 0 00-6.364 0z" />
            </svg>
            {material.like_count || 0}
          </span>
        </div>
      </div>
    </div>
  );

  return (
    <div className="px-4 py-3">
      {feedConfigs.map((feedConfig) => {
        const materialsWithAds = insertAdsIntoFeed(
          feedConfig.materials || [],
          adPositions
        );
        
        return (
          <div key={feedConfig.id}>
            {showTitle && (
              <h3 className="text-lg font-bold mb-3">{feedConfig.title}</h3>
            )}
            <div className="grid grid-cols-2 gap-3">
              {materialsWithAds.length > 0 ? (
                materialsWithAds.map((material: any) => renderFeedCard(material))
              ) : (
                <div className="col-span-2 text-center text-gray-400 py-8">
                  暂无内容
                </div>
              )}
            </div>
          </div>
        );
      })}
    </div>
  );
}
