'use client';

import { type AdSlot } from '@/lib/adApi';

interface AdCardProps {
  ad: AdSlot;
}

export default function AdCard({ ad }: AdCardProps) {
  const handleClick = () => {
    if (ad.link_url) {
      window.open(ad.link_url, '_blank');
    }
  };

  return (
    <div className="bg-white rounded-lg overflow-hidden shadow-sm relative">
      {/* 广告标识 */}
      <div className="absolute top-2 right-2 z-10">
        <span className="bg-black/60 text-white text-xs px-2 py-1 rounded">
          广告
        </span>
      </div>

      {ad.ad_type === 'image' && ad.ad_content.image_url ? (
        // 图片广告
        <div className="relative aspect-[3/4] cursor-pointer" onClick={handleClick}>
          <img
            src={ad.ad_content.image_url}
            alt={ad.ad_content.title}
            className="w-full h-full object-cover"
          />
          <div className="absolute bottom-0 left-0 right-0 bg-gradient-to-t from-black/60 to-transparent p-3">
            <p className="text-white text-sm font-medium line-clamp-2">
              {ad.ad_content.title}
            </p>
          </div>
        </div>
      ) : ad.ad_type === 'video' && ad.ad_content.video_url ? (
        // 视频广告
        <div className="relative aspect-[3/4] bg-black cursor-pointer" onClick={handleClick}>
          <video
            src={ad.ad_content.video_url}
            className="w-full h-full object-cover"
            autoPlay
            muted
            loop
            playsInline
          />
          <div className="absolute bottom-0 left-0 right-0 bg-gradient-to-t from-black/60 to-transparent p-3">
            <p className="text-white text-sm font-medium line-clamp-2">
              {ad.ad_content.title}
            </p>
            {ad.ad_content.duration && (
              <p className="text-white/80 text-xs mt-1">
                {ad.ad_content.duration}秒
              </p>
            )}
          </div>
          {/* 播放图标 */}
          <div className="absolute inset-0 flex items-center justify-center">
            <div className="w-16 h-16 bg-white/30 rounded-full flex items-center justify-center">
              <svg className="w-8 h-8 text-white" fill="currentColor" viewBox="0 0 24 24">
                <path d="M8 5v14l11-7z" />
              </svg>
            </div>
          </div>
        </div>
      ) : (
        // 默认广告占位
        <div className="aspect-[3/4] bg-gradient-to-br from-gray-200 to-gray-300 flex items-center justify-center">
          <div className="text-center text-gray-500">
            <p className="text-sm">{ad.name}</p>
            <p className="text-xs mt-1">{ad.description}</p>
          </div>
        </div>
      )}

      {/* 广告信息 */}
      <div className="p-3">
        <div className="flex items-center justify-between">
          <span className="text-xs text-gray-500">{ad.name}</span>
          <span className="text-xs text-primary-500" onClick={handleClick}>
            了解详情
          </span>
        </div>
      </div>
    </div>
  );
}
