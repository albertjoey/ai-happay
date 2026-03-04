'use client';

import { useEffect, useState } from 'react';
import { useRouter } from 'next/navigation';
import { getFeedConfigList, getAdSlotList, type FeedConfig, type AdSlot } from '@/lib/configApi';

interface FeedSectionProps {
  channelId: number;
  feedId?: number;
  autoLoad?: boolean;
  showTitle?: boolean;
}

export default function FeedSection({ channelId, feedId, autoLoad = true, showTitle = true }: FeedSectionProps) {
  const router = useRouter();
  const [feedConfigs, setFeedConfigs] = useState<FeedConfig[]>([]);
  const [adSlots, setAdSlots] = useState<AdSlot[]>([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    loadFeedConfigs();
  }, [channelId, feedId]);

  const loadFeedConfigs = async () => {
    setLoading(true);
    const [feedData, adData] = await Promise.all([
      getFeedConfigList(channelId),
      getAdSlotList(channelId)
    ]);
    
    // 如果指定了Feed配置ID，只显示该配置
    if (feedId) {
      const filtered = feedData.filter(f => f.id === feedId);
      setFeedConfigs(filtered);
    } else {
      setFeedConfigs(feedData);
    }
    setAdSlots(adData);
    setLoading(false);
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

  // 在物料列表中插入广告
  const insertAdsIntoMaterials = (materials: any[], adSlots: AdSlot[]) => {
    if (adSlots.length === 0 || materials.length === 0) {
      return materials;
    }

    const result: any[] = [];
    let adIndex = 0;

    materials.forEach((material, index) => {
      result.push({ ...material, type: 'material' });

      // 检查是否需要插入广告
      adSlots.forEach((ad) => {
        if (ad.status !== 1) return;

        if (ad.insert_type === 'fixed') {
          // 固定位置插入
          const position = (ad.insert_rule as any).fixed_position || 3;
          if (index + 1 === position && adIndex < (ad.insert_rule as any).max_count) {
            result.push({ ...ad, type: 'ad' });
            adIndex++;
          }
        } else if (ad.insert_type === 'interval') {
          // 间隔插入
          const interval = (ad.insert_rule as any).interval || 5;
          if ((index + 1) % interval === 0 && adIndex < (ad.insert_rule as any).max_count) {
            result.push({ ...ad, type: 'ad' });
            adIndex++;
          }
        } else if (ad.insert_type === 'random') {
          // 随机插入 (简化实现: 每3-5条随机插入)
          const randomInterval = Math.floor(Math.random() * 3) + 3;
          if ((index + 1) % randomInterval === 0 && adIndex < (ad.insert_rule as any).max_count) {
            result.push({ ...ad, type: 'ad' });
            adIndex++;
          }
        }
      });
    });

    return result;
  };

  // 渲染物料卡片
  const renderMaterialCard = (material: any) => (
    <div 
      className="bg-white rounded-lg overflow-hidden shadow-sm cursor-pointer"
      onClick={() => router.push(`/content/${material.id}`)}
    >
      <div className="relative aspect-[3/4]">
        <img
          src={material.cover_url}
          alt={material.title}
          className="w-full h-full object-cover"
        />
        {material.type === 'video' && (
          <div className="absolute top-2 right-2 bg-black/60 text-white text-xs px-2 py-1 rounded">
            {Math.floor(material.duration / 60)}:{String(material.duration % 60).padStart(2, '0')}
          </div>
        )}
        <div className="absolute bottom-0 left-0 right-0 bg-gradient-to-t from-black/60 to-transparent p-3">
          <p className="text-white text-sm font-medium line-clamp-2">{material.title}</p>
        </div>
      </div>
      <div className="p-3">
        <div className="flex items-center gap-2">
          <div className="w-6 h-6 rounded-full bg-gray-300 flex items-center justify-center">
            <span className="text-xs text-gray-600">{material.author?.charAt(0)}</span>
          </div>
          <span className="text-xs text-gray-600 flex-1 truncate">
            {material.author}
          </span>
        </div>
        <div className="flex items-center gap-3 mt-2 text-xs text-gray-500">
          <span className="flex items-center gap-1">
            <svg className="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
              <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z" />
            </svg>
            {material.view_count}
          </span>
          <span className="flex items-center gap-1">
            <svg className="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M4.318 6.318a4.5 4.5 0 000 6.364L12 20.364l7.682-7.682a4.5 4.5 0 00-6.364-6.364L12 7.636l-1.318-1.318a4.5 4.5 0 00-6.364 0z" />
            </svg>
            {material.like_count}
          </span>
        </div>
      </div>
    </div>
  );

  // 渲染广告卡片
  const renderAdCard = (ad: any) => (
    <div className="bg-gradient-to-br from-blue-50 to-purple-50 rounded-lg overflow-hidden shadow-sm border-2 border-blue-200">
      <div className="relative aspect-[3/4]">
        <img
          src={(ad.ad_content as any).image_url}
          alt={(ad.ad_content as any).title}
          className="w-full h-full object-cover"
        />
        <div className="absolute top-2 right-2 bg-red-500 text-white text-xs px-2 py-1 rounded">
          广告
        </div>
        <div className="absolute bottom-0 left-0 right-0 bg-gradient-to-t from-black/60 to-transparent p-3">
          <p className="text-white text-sm font-medium">{(ad.ad_content as any).title}</p>
        </div>
      </div>
      <div className="p-3 bg-blue-50">
        <p className="text-xs text-blue-600 text-center">点击查看详情</p>
      </div>
    </div>
  );

  // 根据layout_type渲染不同布局
  const renderFeed = (config: FeedConfig) => {
    switch (config.layout_type) {
      case 'two_col':
        return renderTwoColumnFeed(config);
      case 'three_col':
        return renderThreeColumnFeed(config);
      case 'big':
        return renderBigFeed(config);
      case 'list':
        return renderListFeed(config);
      case 'mixed':
        return renderMixedFeed(config);
      default:
        return renderTwoColumnFeed(config);
    }
  };

  // 两列布局
  const renderTwoColumnFeed = (config: FeedConfig) => {
    const materials = config.materials || [];
    const itemsWithAds = insertAdsIntoMaterials(materials, adSlots);

    return (
      <div key={config.id} className="mb-4">
        {showTitle && <h3 className="text-base font-semibold text-gray-800 mb-3 px-4">{config.title}</h3>}
        <div className="grid grid-cols-2 gap-3 px-4">
          {itemsWithAds.map((item, index) => (
            <div key={`${item.type}-${item.id || index}`}>
              {item.type === 'ad' ? renderAdCard(item) : renderMaterialCard(item)}
            </div>
          ))}
        </div>
      </div>
    );
  };

  // 三列布局
  const renderThreeColumnFeed = (config: FeedConfig) => {
    const materials = config.materials || [];
    const itemsWithAds = insertAdsIntoMaterials(materials, adSlots);

    return (
      <div key={config.id} className="mb-4">
        {showTitle && <h3 className="text-base font-semibold text-gray-800 mb-3 px-4">{config.title}</h3>}
        <div className="grid grid-cols-3 gap-2 px-4">
          {itemsWithAds.map((item, index) => (
            <div key={`${item.type}-${item.id || index}`}>
              {item.type === 'ad' ? (
                <div className="bg-gradient-to-br from-blue-50 to-purple-50 rounded-lg overflow-hidden shadow-sm border border-blue-200 aspect-square flex items-center justify-center">
                  <div className="text-center p-2">
                    <div className="bg-red-500 text-white text-xs px-2 py-1 rounded inline-block mb-2">
                      广告
                    </div>
                    <p className="text-xs text-gray-700">{(item.ad_content as any).title}</p>
                  </div>
                </div>
              ) : (
                <div className="bg-white rounded-lg overflow-hidden shadow-sm cursor-pointer" onClick={() => router.push(`/content/${item.id}`)}>
                  <div className="relative aspect-square">
                    <img
                      src={item.cover_url}
                      alt={item.title}
                      className="w-full h-full object-cover"
                    />
                  </div>
                  <div className="p-2">
                    <p className="text-xs text-gray-800 line-clamp-2">{item.title}</p>
                  </div>
                </div>
              )}
            </div>
          ))}
        </div>
      </div>
    );
  };

  // 大图模式
  const renderBigFeed = (config: FeedConfig) => {
    const materials = config.materials || [];
    const itemsWithAds = insertAdsIntoMaterials(materials, adSlots);

    return (
      <div key={config.id} className="mb-4 px-4">
        {showTitle && <h3 className="text-base font-semibold text-gray-800 mb-3">{config.title}</h3>}
        {itemsWithAds.map((item, index) => (
          <div key={`${item.type}-${item.id || index}`} className="mb-3">
            {item.type === 'ad' ? (
              <div className="bg-gradient-to-br from-blue-50 to-purple-50 rounded-lg overflow-hidden shadow-sm border-2 border-blue-200">
                <div className="relative aspect-[16/9]">
                  <img
                    src={(item.ad_content as any).image_url}
                    alt={(item.ad_content as any).title}
                    className="w-full h-full object-cover"
                  />
                  <div className="absolute top-2 right-2 bg-red-500 text-white text-xs px-2 py-1 rounded">
                    广告
                  </div>
                </div>
                <div className="p-4 bg-blue-50">
                  <p className="text-sm font-medium text-blue-700">{(item.ad_content as any).title}</p>
                  <p className="text-xs text-blue-600 mt-1">点击查看详情</p>
                </div>
              </div>
            ) : (
              <div className="bg-white rounded-lg overflow-hidden shadow-sm cursor-pointer" onClick={() => router.push(`/content/${item.id}`)}>
                <div className="relative aspect-[16/9]">
                  <img
                    src={item.cover_url}
                    alt={item.title}
                    className="w-full h-full object-cover"
                  />
                </div>
                <div className="p-4">
                  <p className="text-sm font-medium text-gray-800 mb-2">{item.title}</p>
                  <div className="flex items-center gap-2">
                    <div className="w-6 h-6 rounded-full bg-gray-300 flex items-center justify-center">
                      <span className="text-xs text-gray-600">{item.author?.charAt(0)}</span>
                    </div>
                    <span className="text-xs text-gray-600">{item.author}</span>
                  </div>
                </div>
              </div>
            )}
          </div>
        ))}
      </div>
    );
  };

  // 列表模式
  const renderListFeed = (config: FeedConfig) => {
    const materials = config.materials || [];
    const itemsWithAds = insertAdsIntoMaterials(materials, adSlots);

    return (
      <div key={config.id} className="mb-4 px-4">
        {showTitle && <h3 className="text-base font-semibold text-gray-800 mb-3">{config.title}</h3>}
        {itemsWithAds.map((item, index) => (
          <div key={`${item.type}-${item.id || index}`} className="mb-3">
            {item.type === 'ad' ? (
              <div className="bg-gradient-to-r from-blue-50 to-purple-50 rounded-lg overflow-hidden shadow-sm border border-blue-200 flex">
                <div className="w-32 h-24 flex-shrink-0">
                  <img
                    src={(item.ad_content as any).image_url}
                    alt={(item.ad_content as any).title}
                    className="w-full h-full object-cover"
                  />
                </div>
                <div className="flex-1 p-3 flex flex-col justify-center">
                  <div className="flex items-center gap-2 mb-1">
                    <span className="bg-red-500 text-white text-xs px-2 py-0.5 rounded">广告</span>
                    <p className="text-sm font-medium text-gray-800">{(item.ad_content as any).title}</p>
                  </div>
                  <p className="text-xs text-blue-600">点击查看详情</p>
                </div>
              </div>
            ) : (
              <div className="bg-white rounded-lg overflow-hidden shadow-sm flex">
                <div className="w-32 h-24 flex-shrink-0">
                  <img
                    src={item.cover_url}
                    alt={item.title}
                    className="w-full h-full object-cover"
                  />
                </div>
                <div className="flex-1 p-3">
                  <p className="text-sm font-medium text-gray-800 line-clamp-2 mb-2">{item.title}</p>
                  <div className="flex items-center gap-2">
                    <div className="w-5 h-5 rounded-full bg-gray-300 flex items-center justify-center">
                      <span className="text-xs text-gray-600">{item.author?.charAt(0)}</span>
                    </div>
                    <span className="text-xs text-gray-600">{item.author}</span>
                  </div>
                </div>
              </div>
            )}
          </div>
        ))}
      </div>
    );
  };

  // 混合布局
  const renderMixedFeed = (config: FeedConfig) => {
    const materials = config.materials || [];
    const itemsWithAds = insertAdsIntoMaterials(materials, adSlots);

    return (
      <div key={config.id} className="mb-4 px-4">
        {showTitle && <h3 className="text-base font-semibold text-gray-800 mb-3">{config.title}</h3>}
        <div className="space-y-3">
          {itemsWithAds.map((item, index) => {
            // 第一个用大图
            if (index === 0) {
              return (
                <div key={`${item.type}-${item.id || index}`} className="bg-white rounded-lg overflow-hidden shadow-sm">
                  <div className="relative aspect-[16/9]">
                    <img
                      src={item.type === 'ad' ? (item.ad_content as any).image_url : item.cover_url}
                      alt={item.type === 'ad' ? (item.ad_content as any).title : item.title}
                      className="w-full h-full object-cover"
                    />
                    {item.type === 'ad' && (
                      <div className="absolute top-2 right-2 bg-red-500 text-white text-xs px-2 py-1 rounded">
                        广告
                      </div>
                    )}
                  </div>
                  <div className="p-4">
                    <p className="text-sm font-medium text-gray-800">
                      {item.type === 'ad' ? (item.ad_content as any).title : item.title}
                    </p>
                  </div>
                </div>
              );
            }

            // 其他用两列
            if (index === 1) {
              const remainingItems = itemsWithAds.slice(1);
              return (
                <div key="grid-section" className="grid grid-cols-2 gap-3">
                  {remainingItems.map((subItem, subIndex) => (
                    <div key={`${subItem.type}-${subItem.id || subIndex}`}>
                      {subItem.type === 'ad' ? renderAdCard(subItem) : (
                        <div className="bg-white rounded-lg overflow-hidden shadow-sm cursor-pointer" onClick={() => router.push(`/content/${item.id}`)}>
                          <div className="relative aspect-[3/4]">
                            <img
                              src={subItem.cover_url}
                              alt={subItem.title}
                              className="w-full h-full object-cover"
                            />
                          </div>
                          <div className="p-2">
                            <p className="text-xs text-gray-800 line-clamp-2">{subItem.title}</p>
                          </div>
                        </div>
                      )}
                    </div>
                  ))}
                </div>
              );
            }

            return null;
          })}
        </div>
      </div>
    );
  };

  return <div>{feedConfigs.map(renderFeed)}</div>;
}
