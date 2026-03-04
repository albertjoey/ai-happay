'use client';

import { type Material } from '@/lib/materialApi';

interface MaterialCardProps {
  material: Material;
  layout?: 'vertical' | 'horizontal' | 'grid';
}

export default function MaterialCard({ material, layout = 'vertical' }: MaterialCardProps) {
  const handleClick = () => {
    // 跳转到详情页
    console.log('查看物料详情:', material.id);
  };

  // 格式化数字
  const formatNumber = (num: number) => {
    if (num >= 10000) {
      return (num / 10000).toFixed(1) + 'w';
    }
    if (num >= 1000) {
      return (num / 1000).toFixed(1) + 'k';
    }
    return num.toString();
  };

  // 格式化时长
  const formatDuration = (seconds: number) => {
    const minutes = Math.floor(seconds / 60);
    const secs = seconds % 60;
    return `${minutes}:${secs.toString().padStart(2, '0')}`;
  };

  // 获取类型标签
  const getTypeTag = () => {
    const typeMap: Record<string, { label: string; color: string }> = {
      image_text: { label: '图文', color: 'bg-blue-500' },
      novel: { label: '小说', color: 'bg-green-500' },
      video: { label: '视频', color: 'bg-orange-500' },
      banner: { label: 'Banner', color: 'bg-purple-500' },
      comic: { label: '漫剧', color: 'bg-cyan-500' },
      short_drama: { label: '短剧', color: 'bg-pink-500' },
    };
    return typeMap[material.type] || { label: material.type, color: 'bg-gray-500' };
  };

  const typeTag = getTypeTag();

  if (layout === 'horizontal') {
    return (
      <div className="bg-white rounded-lg overflow-hidden shadow-sm flex" onClick={handleClick}>
        <div className="relative w-32 h-24 flex-shrink-0">
          <img
            src={material.cover_url}
            alt={material.title}
            className="w-full h-full object-cover"
          />
          {material.type === 'video' && material.duration > 0 && (
            <div className="absolute bottom-1 right-1 bg-black/70 text-white text-xs px-1 rounded">
              {formatDuration(material.duration)}
            </div>
          )}
        </div>
        <div className="flex-1 p-3 flex flex-col justify-between">
          <div>
            <h3 className="text-sm font-medium line-clamp-2">{material.title}</h3>
            <p className="text-xs text-gray-500 mt-1">{material.author}</p>
          </div>
          <div className="flex items-center justify-between text-xs text-gray-400">
            <span>{formatNumber(material.view_count)} 浏览</span>
            <span>{formatNumber(material.like_count)} 点赞</span>
          </div>
        </div>
      </div>
    );
  }

  if (layout === 'grid') {
    return (
      <div className="bg-white rounded-lg overflow-hidden shadow-sm" onClick={handleClick}>
        <div className="relative aspect-square">
          <img
            src={material.cover_url}
            alt={material.title}
            className="w-full h-full object-cover"
          />
          <div className={`absolute top-2 left-2 ${typeTag.color} text-white text-xs px-2 py-0.5 rounded`}>
            {typeTag.label}
          </div>
          {material.type === 'video' && material.duration > 0 && (
            <div className="absolute bottom-2 right-2 bg-black/70 text-white text-xs px-2 py-0.5 rounded">
              {formatDuration(material.duration)}
            </div>
          )}
        </div>
        <div className="p-2">
          <h3 className="text-sm font-medium line-clamp-2">{material.title}</h3>
          <div className="flex items-center justify-between mt-2 text-xs text-gray-400">
            <span>{formatNumber(material.view_count)}</span>
            <span>{formatNumber(material.like_count)}</span>
          </div>
        </div>
      </div>
    );
  }

  // 默认垂直布局
  return (
    <div className="bg-white rounded-lg overflow-hidden shadow-sm" onClick={handleClick}>
      <div className="relative aspect-[3/4]">
        <img
          src={material.cover_url}
          alt={material.title}
          className="w-full h-full object-cover"
        />
        <div className={`absolute top-2 left-2 ${typeTag.color} text-white text-xs px-2 py-0.5 rounded`}>
          {typeTag.label}
        </div>
        {material.type === 'video' && material.duration > 0 && (
          <div className="absolute bottom-2 right-2 bg-black/70 text-white text-xs px-2 py-0.5 rounded">
            {formatDuration(material.duration)}
          </div>
        )}
        {material.type === 'novel' && material.word_count > 0 && (
          <div className="absolute bottom-2 right-2 bg-black/70 text-white text-xs px-2 py-0.5 rounded">
            {(material.word_count / 10000).toFixed(0)}万字
          </div>
        )}
      </div>
      <div className="p-3">
        <h3 className="text-sm font-medium line-clamp-2">{material.title}</h3>
        <p className="text-xs text-gray-500 mt-1">{material.author}</p>
        <div className="flex items-center justify-between mt-2 text-xs text-gray-400">
          <span>{formatNumber(material.view_count)} 浏览</span>
          <span>{formatNumber(material.like_count)} 点赞</span>
        </div>
      </div>
    </div>
  );
}
