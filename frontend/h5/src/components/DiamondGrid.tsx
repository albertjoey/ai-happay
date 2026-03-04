'use client';

import { useEffect, useState } from 'react';
import { useRouter } from 'next/navigation';
import { getDiamondList, type Diamond } from '@/lib/configApi';

interface DiamondGridProps {
  channelId: number;
  groupIds?: number[];
}

export default function DiamondGrid({ channelId, groupIds }: DiamondGridProps) {
  const router = useRouter();
  const [diamonds, setDiamonds] = useState<Diamond[]>([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    loadDiamonds();
  }, [channelId, groupIds]);

  const loadDiamonds = async () => {
    setLoading(true);
    const data = await getDiamondList(channelId);
    
    // 如果指定了分组ID，只显示这些分组
    if (groupIds && groupIds.length > 0) {
      const filtered = data.filter(d => groupIds.includes(d.group_id));
      setDiamonds(filtered);
    } else {
      setDiamonds(data);
    }
    setLoading(false);
  };

  const handleDiamondClick = (diamond: Diamond) => {
    // 如果有关联物料,跳转到物料详情页
    if (diamond.material) {
      router.push(`/content/${diamond.material.id}`);
      return;
    }
    
    // 根据link_type和link_value进行跳转
    switch (diamond.link_type) {
      case 'channel':
        // 切换频道
        console.log('Navigate to channel:', diamond.link_value);
        break;
      case 'topic':
        // 跳转到话题页
        router.push(`/topic/${diamond.link_value}`);
        break;
      case 'content':
        // 跳转到内容详情页
        router.push(`/content/${diamond.link_value}`);
        break;
      case 'external':
        // 外部链接
        if (diamond.link_value) {
          window.open(diamond.link_value, '_blank');
        }
        break;
      default:
        console.log('Unknown link type:', diamond.link_type);
    }
  };

  if (loading) {
    return (
      <div className="px-4 py-3">
        <div className="grid grid-cols-5 gap-3">
          {Array.from({ length: 10 }).map((_, i) => (
            <div key={i} className="animate-pulse">
              <div className="w-12 h-12 mx-auto bg-gray-200 rounded-lg"></div>
              <div className="mt-2 h-3 bg-gray-200 rounded w-12 mx-auto"></div>
            </div>
          ))}
        </div>
      </div>
    );
  }

  if (diamonds.length === 0) {
    return null;
  }

  // 按分组整理金刚位
  const groupedDiamonds: Record<number, Diamond[]> = {};
  diamonds.forEach((diamond) => {
    if (!groupedDiamonds[diamond.group_id]) {
      groupedDiamonds[diamond.group_id] = [];
    }
    groupedDiamonds[diamond.group_id].push(diamond);
  });

  return (
    <div className="px-4 py-3">
      {Object.entries(groupedDiamonds).map(([groupId, groupDiamonds]) => (
        <div key={groupId} className="mb-4">
          <div className="grid grid-cols-5 gap-3">
            {groupDiamonds.map((diamond) => (
              <button
                key={diamond.id}
                onClick={() => handleDiamondClick(diamond)}
                className="flex flex-col items-center"
              >
                {/* 图标或物料封面 */}
                <div className="w-12 h-12 rounded-lg overflow-hidden bg-gradient-to-br from-primary-400 to-primary-600 flex items-center justify-center">
                  {diamond.material ? (
                    <img
                      src={diamond.material.cover_url}
                      alt={diamond.material.title}
                      className="w-full h-full object-cover"
                    />
                  ) : (
                    <span className="text-white text-lg font-bold">
                      {diamond.title.charAt(0)}
                    </span>
                  )}
                </div>
                {/* 标题 */}
                <span className="mt-2 text-xs text-gray-700 line-clamp-1">
                  {diamond.title}
                </span>
                {/* 如果有物料,显示浏览量 */}
                {diamond.material && (
                  <span className="text-xs text-gray-400">
                    {formatNumber(diamond.material.view_count)}
                  </span>
                )}
              </button>
            ))}
          </div>
        </div>
      ))}
    </div>
  );
}

// 格式化数字
function formatNumber(num: number) {
  if (num >= 10000) {
    return (num / 10000).toFixed(1) + 'w';
  }
  if (num >= 1000) {
    return (num / 1000).toFixed(1) + 'k';
  }
  return num.toString();
}
