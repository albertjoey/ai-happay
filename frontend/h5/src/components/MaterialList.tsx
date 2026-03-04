'use client';

import { useState, useEffect } from 'react';
import { getMaterialList, type Material } from '@/lib/materialApi';
import MaterialCard from './MaterialCard';

interface MaterialListProps {
  type?: string;
  layout?: 'vertical' | 'horizontal' | 'grid';
  columns?: number;
  limit?: number;
  title?: string;
}

export default function MaterialList({
  type,
  layout = 'vertical',
  columns = 2,
  limit = 10,
  title,
}: MaterialListProps) {
  const [materials, setMaterials] = useState<Material[]>([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    loadMaterials();
  }, [type, limit]);

  const loadMaterials = async () => {
    setLoading(true);
    try {
      const { list } = await getMaterialList({
        type,
        page_size: limit,
      });
      setMaterials(list);
    } catch (error) {
      console.error('加载物料失败:', error);
    } finally {
      setLoading(false);
    }
  };

  if (loading) {
    return (
      <div className="py-4">
        {title && <h2 className="text-lg font-bold mb-4">{title}</h2>}
        <div className="grid gap-4" style={{ gridTemplateColumns: `repeat(${columns}, 1fr)` }}>
          {Array.from({ length: limit }).map((_, i) => (
            <div key={i} className="animate-pulse">
              <div className="bg-gray-200 aspect-[3/4] rounded-lg"></div>
              <div className="mt-2 h-4 bg-gray-200 rounded w-3/4"></div>
              <div className="mt-1 h-3 bg-gray-200 rounded w-1/2"></div>
            </div>
          ))}
        </div>
      </div>
    );
  }

  if (materials.length === 0) {
    return null;
  }

  return (
    <div className="py-4">
      {title && (
        <div className="flex items-center justify-between mb-4">
          <h2 className="text-lg font-bold">{title}</h2>
          <button className="text-sm text-primary-500">查看更多</button>
        </div>
      )}
      <div
        className="grid gap-4"
        style={{ gridTemplateColumns: `repeat(${columns}, 1fr)` }}
      >
        {materials.map((material) => (
          <MaterialCard key={material.id} material={material} layout={layout} />
        ))}
      </div>
    </div>
  );
}
