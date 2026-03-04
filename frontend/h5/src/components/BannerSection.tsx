'use client';

import { useEffect, useState } from 'react';
import { useRouter } from 'next/navigation';
import axios from 'axios';

const API_BASE = process.env.NEXT_PUBLIC_API_BASE || 'http://localhost:4004';

interface Banner {
  id: number;
  title: string;
  image: string;
  link_type: number;
  link_url: string;
  content_id?: number;
  sort: number;
}

interface BannerSectionProps {
  channelId: number;
  bannerIds?: number[];
}

export default function BannerSection({ channelId, bannerIds }: BannerSectionProps) {
  const router = useRouter();
  const [banners, setBanners] = useState<Banner[]>([]);
  const [loading, setLoading] = useState(true);
  const [currentIndex, setCurrentIndex] = useState(0);

  useEffect(() => {
    loadBanners();
  }, [channelId, bannerIds]);

  // 自动轮播
  useEffect(() => {
    if (banners.length <= 1) return;
    
    const timer = setInterval(() => {
      setCurrentIndex((prev) => (prev + 1) % banners.length);
    }, 3000);
    
    return () => clearInterval(timer);
  }, [banners.length]);

  const loadBanners = async () => {
    setLoading(true);
    try {
      const response = await axios.get(`${API_BASE}/api/v1/banner/list`, {
        params: {
          channel_id: channelId,
          status: 1,
        },
      });
      
      let bannerList = response.data.list || [];
      
      // 如果指定了Banner IDs，只显示这些Banner
      if (bannerIds && bannerIds.length > 0) {
        bannerList = bannerList.filter((b: Banner) => bannerIds.includes(b.id));
      }
      
      setBanners(bannerList);
    } catch (error) {
      console.error('加载Banner失败:', error);
    } finally {
      setLoading(false);
    }
  };

  if (loading) {
    return (
      <div className="px-4 py-3">
        <div className="aspect-[2/1] bg-gray-200 rounded-lg animate-pulse" />
      </div>
    );
  }

  if (banners.length === 0) {
    return null;
  }

  const handleBannerClick = (banner: Banner) => {
    // link_type: 1=内容, 2=外部链接, 3=话题
    switch (banner.link_type) {
      case 1:
        // 跳转到内容详情页
        if (banner.content_id) {
          router.push(`/content/${banner.content_id}`);
        } else if (banner.link_url) {
          // 如果link_url是数字，当作内容ID
          const contentId = parseInt(banner.link_url);
          if (!isNaN(contentId)) {
            router.push(`/content/${contentId}`);
          }
        }
        break;
      case 2:
        // 外部链接
        if (banner.link_url) {
          window.open(banner.link_url, '_blank');
        }
        break;
      case 3:
        // 跳转到话题页
        if (banner.link_url) {
          router.push(`/topic/${banner.link_url}`);
        }
        break;
      default:
        console.log('Unknown link type:', banner.link_type);
    }
  };

  return (
    <div className="px-4 py-3">
      <div className="relative overflow-hidden rounded-lg">
        {/* Banner图片 */}
        <div 
          className="flex transition-transform duration-300 ease-out"
          style={{ transform: `translateX(-${currentIndex * 100}%)` }}
        >
          {banners.map((banner) => (
            <div
              key={banner.id}
              className="flex-shrink-0 w-full aspect-[2/1] cursor-pointer"
              onClick={() => handleBannerClick(banner)}
            >
              <img
                src={banner.image}
                alt={banner.title}
                className="w-full h-full object-cover"
              />
            </div>
          ))}
        </div>

        {/* 指示器 */}
        {banners.length > 1 && (
          <div className="absolute bottom-2 left-0 right-0 flex justify-center gap-1.5">
            {banners.map((_, index) => (
              <button
                key={index}
                className={`w-2 h-2 rounded-full transition-all ${
                  index === currentIndex 
                    ? 'bg-white w-4' 
                    : 'bg-white/50'
                }`}
                onClick={() => setCurrentIndex(index)}
              />
            ))}
          </div>
        )}
      </div>
    </div>
  );
}
