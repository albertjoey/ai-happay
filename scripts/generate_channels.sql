-- =============================================
-- 第三部分：创建频道配置
-- =============================================

SET NAMES utf8mb4;

-- 清理旧频道数据
DELETE FROM banner;
DELETE FROM diamond;
DELETE FROM recommend;
DELETE FROM feed_config;
DELETE FROM channel_config;
DELETE FROM channel WHERE id > 1;

-- 1. 推荐频道（综合频道）
INSERT INTO channel (tenant_id, name, code, description, status, sort) VALUES
(1, '推荐', 'recommend', '综合推荐频道，包含全类型内容', 1, 1);

-- 2. 图文频道
INSERT INTO channel (tenant_id, name, code, description, status, sort) VALUES
(1, '图文', 'image_text', '精选图文内容', 1, 2);

-- 3. 视频频道
INSERT INTO channel (tenant_id, name, code, description, status, sort) VALUES
(1, '视频', 'video', '长视频和短视频', 1, 3);

-- 4. 小说频道
INSERT INTO channel (tenant_id, name, code, description, status, sort) VALUES
(1, '小说', 'novel', '热门小说阅读', 1, 4);

-- 5. 短剧频道
INSERT INTO channel (tenant_id, name, code, description, status, sort) VALUES
(1, '短剧', 'short_drama', '精彩短剧', 1, 5);

-- 6. 漫剧频道
INSERT INTO channel (tenant_id, name, code, description, status, sort) VALUES
(1, '漫剧', 'drama', '热门漫剧', 1, 6);

SELECT '频道创建完成' as status;

-- =============================================
-- 第四部分：配置推荐频道（完整配置）
-- =============================================

-- 获取推荐频道ID
SET @recommend_channel_id = (SELECT id FROM channel WHERE code = 'recommend');

-- Banner配置（5张轮播）
INSERT INTO banner (tenant_id, channel_id, title, image_url, link_type, link_value, content_id, sort, status) VALUES
(1, @recommend_channel_id, '热门小说推荐', 'https://picsum.photos/seed/banner1/800/400', 1, '', (SELECT id FROM material WHERE type = 'novel' ORDER BY view_count DESC LIMIT 1), 1, 1),
(1, @recommend_channel_id, '精彩短剧上线', 'https://picsum.photos/seed/banner2/800/400', 1, '', (SELECT id FROM material WHERE type = 'short_drama' ORDER BY view_count DESC LIMIT 1), 2, 1),
(1, @recommend_channel_id, '漫剧新番', 'https://picsum.photos/seed/banner3/800/400', 1, '', (SELECT id FROM material WHERE type = 'drama' ORDER BY view_count DESC LIMIT 1), 3, 1),
(1, @recommend_channel_id, '热门视频', 'https://picsum.photos/seed/banner4/800/400', 1, '', (SELECT id FROM material WHERE type = 'long_video' ORDER BY view_count DESC LIMIT 1), 4, 1),
(1, @recommend_channel_id, '精选图文', 'https://picsum.photos/seed/banner5/800/400', 1, '', (SELECT id FROM material WHERE type = 'image_text' ORDER BY view_count DESC LIMIT 1), 5, 1);

-- 金刚位配置（8个入口）
INSERT INTO diamond (tenant_id, channel_id, group_id, title, icon, link_type, link_value, material_id, sort, status) VALUES
(1, @recommend_channel_id, 1, '热门小说', '📚', 'channel', 'novel', 0, 1, 1),
(1, @recommend_channel_id, 1, '精选短剧', '🎬', 'channel', 'short_drama', 0, 2, 1),
(1, @recommend_channel_id, 1, '热门漫剧', '📖', 'channel', 'drama', 0, 3, 1),
(1, @recommend_channel_id, 1, '长视频', '🎥', 'channel', 'video', 0, 4, 1),
(1, @recommend_channel_id, 1, '图文精选', '📄', 'channel', 'image_text', 0, 5, 1),
(1, @recommend_channel_id, 1, '排行榜', '🏆', 'page', 'ranking', 0, 6, 1),
(1, @recommend_channel_id, 1, '每日推荐', '⭐', 'recommend', '', 0, 7, 1),
(1, @recommend_channel_id, 1, '更多分类', '📋', 'category', '', 0, 8, 1);

-- 推荐位配置（3种布局）
INSERT INTO recommend (tenant_id, channel_id, title, display_type, source_type, content_ids, filter_rule, sort, status) VALUES
(1, @recommend_channel_id, '热门推荐', 'scroll', 'manual', '[]', '{"types": ["novel", "short_drama"], "order_by": "view_count", "limit": 10}', 1, 1),
(1, @recommend_channel_id, '精选内容', 'grid', 'manual', '[]', '{"types": ["image_text", "short_video"], "order_by": "like_count", "limit": 6}', 2, 1),
(1, @recommend_channel_id, '为你推荐', 'single', 'algorithm', '[]', '{"limit": 5}', 3, 1);

-- Feed流配置（两列瀑布流）
INSERT INTO feed_config (tenant_id, channel_id, title, layout_type, content_strategy, content_ids, filter_rule, sort, status) VALUES
(1, @recommend_channel_id, '推荐流', 'waterfall_2', 'algorithm', '[]', '{"limit": 50}', 1, 1);

SELECT '推荐频道配置完成' as status;

-- =============================================
-- 第五部分：配置其他频道
-- =============================================

-- 图文频道配置
SET @image_text_channel_id = (SELECT id FROM channel WHERE code = 'image_text');
INSERT INTO banner (tenant_id, channel_id, title, image_url, link_type, link_value, content_id, sort, status) VALUES
(1, @image_text_channel_id, '图文精选', 'https://picsum.photos/seed/it_banner1/800/400', 1, '', (SELECT id FROM material WHERE type = 'image_text' ORDER BY id LIMIT 1), 1, 1);

INSERT INTO recommend (tenant_id, channel_id, title, display_type, source_type, content_ids, filter_rule, sort, status) VALUES
(1, @image_text_channel_id, '热门图文', 'scroll', 'filter', '[]', '{"types": ["image_text"], "order_by": "view_count", "limit": 10}', 1, 1),
(1, @image_text_channel_id, '最新图文', 'grid', 'filter', '[]', '{"types": ["image_text"], "order_by": "created_at", "limit": 9}', 2, 1);

INSERT INTO feed_config (tenant_id, channel_id, title, layout_type, content_strategy, content_ids, filter_rule, sort, status) VALUES
(1, @image_text_channel_id, '图文流', 'waterfall_2', 'filter', '[]', '{"types": ["image_text"], "limit": 50}', 1, 1);

-- 视频频道配置
SET @video_channel_id = (SELECT id FROM channel WHERE code = 'video');
INSERT INTO banner (tenant_id, channel_id, title, image_url, link_type, link_value, content_id, sort, status) VALUES
(1, @video_channel_id, '热门视频', 'https://picsum.photos/seed/v_banner1/800/400', 1, '', (SELECT id FROM material WHERE type = 'long_video' ORDER BY id LIMIT 1), 1, 1);

INSERT INTO diamond (tenant_id, channel_id, group_id, title, icon, link_type, link_value, material_id, sort, status) VALUES
(1, @video_channel_id, 1, '电影', '🎬', 'filter', '{"type": "long_video", "category": "电影"}', 0, 1, 1),
(1, @video_channel_id, 1, '电视剧', '📺', 'filter', '{"type": "long_video", "category": "电视剧"}', 0, 2, 1),
(1, @video_channel_id, 1, '综艺', '🎭', 'filter', '{"type": "long_video", "category": "综艺"}', 0, 3, 1),
(1, @video_channel_id, 1, '短视频', '📱', 'filter', '{"type": "short_video"}', 0, 4, 1);

INSERT INTO recommend (tenant_id, channel_id, title, display_type, source_type, content_ids, filter_rule, sort, status) VALUES
(1, @video_channel_id, '热门长视频', 'scroll', 'filter', '[]', '{"types": ["long_video"], "order_by": "view_count", "limit": 10}', 1, 1),
(1, @video_channel_id, '精选短视频', 'grid', 'filter', '[]', '{"types": ["short_video"], "order_by": "like_count", "limit": 6}', 2, 1);

INSERT INTO feed_config (tenant_id, channel_id, title, layout_type, content_strategy, content_ids, filter_rule, sort, status) VALUES
(1, @video_channel_id, '视频流', 'list', 'filter', '[]', '{"types": ["long_video", "short_video"], "limit": 50}', 1, 1);

-- 小说频道配置
SET @novel_channel_id = (SELECT id FROM channel WHERE code = 'novel');
INSERT INTO banner (tenant_id, channel_id, title, image_url, link_type, link_value, content_id, sort, status) VALUES
(1, @novel_channel_id, '热门小说', 'https://picsum.photos/seed/n_banner1/800/400', 1, '', (SELECT id FROM material WHERE type = 'novel' ORDER BY id LIMIT 1), 1, 1);

INSERT INTO diamond (tenant_id, channel_id, group_id, title, icon, link_type, link_value, material_id, sort, status) VALUES
(1, @novel_channel_id, 1, '玄幻', '⚡', 'filter', '{"type": "novel", "category": "玄幻"}', 0, 1, 1),
(1, @novel_channel_id, 1, '都市', '🏙️', 'filter', '{"type": "novel", "category": "都市"}', 0, 2, 1),
(1, @novel_channel_id, 1, '仙侠', '🗡️', 'filter', '{"type": "novel", "category": "仙侠"}', 0, 3, 1),
(1, @novel_channel_id, 1, '历史', '📜', 'filter', '{"type": "novel", "category": "历史"}', 0, 4, 1);

INSERT INTO recommend (tenant_id, channel_id, title, display_type, source_type, content_ids, filter_rule, sort, status) VALUES
(1, @novel_channel_id, '热门小说', 'scroll', 'filter', '[]', '{"types": ["novel"], "order_by": "view_count", "limit": 10}', 1, 1),
(1, @novel_channel_id, '新书推荐', 'grid', 'filter', '[]', '{"types": ["novel"], "order_by": "created_at", "limit": 6}', 2, 1),
(1, @novel_channel_id, '完结精品', 'single', 'filter', '[]', '{"types": ["novel"], "limit": 3}', 3, 1);

INSERT INTO feed_config (tenant_id, channel_id, title, layout_type, content_strategy, content_ids, filter_rule, sort, status) VALUES
(1, @novel_channel_id, '小说流', 'waterfall_3', 'filter', '[]', '{"types": ["novel"], "limit": 50}', 1, 1);

-- 短剧频道配置
SET @short_drama_channel_id = (SELECT id FROM channel WHERE code = 'short_drama');
INSERT INTO banner (tenant_id, channel_id, title, image_url, link_type, link_value, content_id, sort, status) VALUES
(1, @short_drama_channel_id, '热门短剧', 'https://picsum.photos/seed/sd_banner1/800/400', 1, '', (SELECT id FROM material WHERE type = 'short_drama' ORDER BY id LIMIT 1), 1, 1);

INSERT INTO recommend (tenant_id, channel_id, title, display_type, source_type, content_ids, filter_rule, sort, status) VALUES
(1, @short_drama_channel_id, '热播短剧', 'scroll', 'filter', '[]', '{"types": ["short_drama"], "order_by": "view_count", "limit": 10}', 1, 1),
(1, @short_drama_channel_id, '甜宠专区', 'grid', 'filter', '[]', '{"types": ["short_drama"], "category": "甜宠", "limit": 6}', 2, 1);

INSERT INTO feed_config (tenant_id, channel_id, title, layout_type, content_strategy, content_ids, filter_rule, sort, status) VALUES
(1, @short_drama_channel_id, '短剧流', 'waterfall_2', 'filter', '[]', '{"types": ["short_drama"], "limit": 50}', 1, 1);

-- 漫剧频道配置
SET @drama_channel_id = (SELECT id FROM channel WHERE code = 'drama');
INSERT INTO banner (tenant_id, channel_id, title, image_url, link_type, link_value, content_id, sort, status) VALUES
(1, @drama_channel_id, '热门漫剧', 'https://picsum.photos/seed/d_banner1/800/400', 1, '', (SELECT id FROM material WHERE type = 'drama' ORDER BY id LIMIT 1), 1, 1);

INSERT INTO diamond (tenant_id, channel_id, group_id, title, icon, link_type, link_value, material_id, sort, status) VALUES
(1, @drama_channel_id, 1, '热血', '🔥', 'filter', '{"type": "drama", "category": "热血"}', 0, 1, 1),
(1, @drama_channel_id, 1, '恋爱', '💕', 'filter', '{"type": "drama", "category": "恋爱"}', 0, 2, 1),
(1, @drama_channel_id, 1, '玄幻', '✨', 'filter', '{"type": "drama", "category": "玄幻"}', 0, 3, 1),
(1, @drama_channel_id, 1, '都市', '🌆', 'filter', '{"type": "drama", "category": "都市"}', 0, 4, 1);

INSERT INTO recommend (tenant_id, channel_id, title, display_type, source_type, content_ids, filter_rule, sort, status) VALUES
(1, @drama_channel_id, '热门漫剧', 'scroll', 'filter', '[]', '{"types": ["drama"], "order_by": "view_count", "limit": 10}', 1, 1),
(1, @drama_channel_id, '新作推荐', 'grid', 'filter', '[]', '{"types": ["drama"], "order_by": "created_at", "limit": 6}', 2, 1);

INSERT INTO feed_config (tenant_id, channel_id, title, layout_type, content_strategy, content_ids, filter_rule, sort, status) VALUES
(1, @drama_channel_id, '漫剧流', 'waterfall_2', 'filter', '[]', '{"types": ["drama"], "limit": 50}', 1, 1);

SELECT '所有频道配置完成' as status;

-- 统计数据
SELECT '=== 数据统计 ===' as info;
SELECT type, COUNT(*) as count FROM material GROUP BY type;
SELECT COUNT(*) as chapter_count FROM material_chapter;
SELECT COUNT(*) as channel_count FROM channel WHERE deleted_at IS NULL;
SELECT COUNT(*) as banner_count FROM banner WHERE deleted_at IS NULL;
SELECT COUNT(*) as diamond_count FROM diamond WHERE deleted_at IS NULL;
SELECT COUNT(*) as recommend_count FROM recommend WHERE deleted_at IS NULL;
SELECT COUNT(*) as feed_count FROM feed_config WHERE deleted_at IS NULL;
