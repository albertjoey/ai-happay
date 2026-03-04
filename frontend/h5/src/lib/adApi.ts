import axios from 'axios';

const API_BASE = process.env.NEXT_PUBLIC_API_BASE || 'http://localhost:4004';

// 广告位类型
export interface AdSlot {
  id: number;
  channel_id: number;
  name: string;
  insert_type: string;
  insert_rule: {
    fixed_position?: number;
    interval?: number;
    max_count: number;
  };
  ad_type: string;
  ad_content: {
    image_url?: string;
    video_url?: string;
    title: string;
    duration?: number;
  };
  link_url: string;
  status: number;
  sort: number;
  description: string;
}

// 获取广告位列表
export async function getAdSlotList(channelId: number): Promise<AdSlot[]> {
  try {
    const response = await axios.get(`${API_BASE}/api/v1/ad-slot/list`, {
      params: { channel_id: channelId, status: 1 },
    });
    return response.data.list || [];
  } catch (error) {
    console.error('获取广告位列表失败:', error);
    return [];
  }
}

// 根据插入规则计算广告位置
export function calculateAdPositions(
  adSlots: AdSlot[],
  contentCount: number
): Map<number, AdSlot[]> {
  const adPositions = new Map<number, AdSlot[]>();

  adSlots.forEach((ad) => {
    const positions: number[] = [];

    switch (ad.insert_type) {
      case 'fixed':
        // 固定位置插入
        if (ad.insert_rule.fixed_position && ad.insert_rule.fixed_position <= contentCount) {
          positions.push(ad.insert_rule.fixed_position);
        }
        break;

      case 'interval':
        // 间隔插入
        if (ad.insert_rule.interval) {
          let count = 0;
          for (let i = ad.insert_rule.interval; i <= contentCount && count < ad.insert_rule.max_count; i += ad.insert_rule.interval) {
            positions.push(i);
            count++;
          }
        }
        break;

      case 'random':
        // 随机插入
        const randomPositions: number[] = [];
        for (let i = 0; i < ad.insert_rule.max_count; i++) {
          const pos = Math.floor(Math.random() * contentCount) + 1;
          if (!randomPositions.includes(pos)) {
            randomPositions.push(pos);
          }
        }
        randomPositions.sort((a, b) => a - b).forEach(pos => positions.push(pos));
        break;
    }

    positions.forEach((pos) => {
      if (!adPositions.has(pos)) {
        adPositions.set(pos, []);
      }
      adPositions.get(pos)!.push(ad);
    });
  });

  return adPositions;
}
