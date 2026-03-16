import axios from 'axios';

// 发现页API使用channel服务端口4004
const DISCOVER_API_BASE = process.env.NEXT_PUBLIC_API_BASE || 'http://localhost:4004';

// 发现页模块数据
export interface DiscoverModule {
  module: string;
  title: string;
  items: any[];
}

// 发现页响应
export interface DiscoverResponse {
  modules: DiscoverModule[];
}

// 获取发现页完整数据
export async function getDiscoverPage(): Promise<DiscoverResponse> {
  try {
    console.log('调用发现页API:', `${DISCOVER_API_BASE}/api/v1/discover`);
    const response = await axios.get(`${DISCOVER_API_BASE}/api/v1/discover`, {
      timeout: 10000,
      headers: {
        'Content-Type': 'application/json',
      }
    });
    console.log('API响应:', response.data);
    return response.data;
  } catch (error: any) {
    console.error('获取发现页数据失败:', error);
    console.error('错误详情:', {
      message: error.message,
      code: error.code,
      response: error.response
    });
    throw new Error(`API调用失败: ${error.message}`);
  }
}

// 获取发现页模块内容
export async function getDiscoverItems(module: string): Promise<any[]> {
  try {
    const response = await axios.get(`${DISCOVER_API_BASE}/api/v1/discover/items`, {
      params: { module }
    });
    return response.data.list || [];
  } catch (error) {
    console.error('获取发现页内容失败:', error);
    return [];
  }
}
