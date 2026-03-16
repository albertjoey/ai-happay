-- 配置频道运营位数据

-- 1. Banner配置 (每个频道3-5个)
DELETE FROM banner;

-- 为每个频道创建Banner
INSERT INTO banner (tenant_id, channel_id, title, image, link_url, link_type, content_id, sort, status, created_at, updated_at)
SELECT 
    1,
    channel_id,
    CONCAT('精彩推荐', n),
    CONCAT('https://picsum.photos/seed/banner', channel_id, '_', n, '/800/400'),
    CONCAT('/material/', (channel_id - 1) * 200 + n),
    1,
    (channel_id - 1) * 200 + n,
    n,
    1,
    NOW(),
    NOW()
FROM 
    (SELECT id as channel_id FROM channel WHERE id <= 7) channels,
    (SELECT 1 as n UNION SELECT 2 UNION SELECT 3 UNION SELECT 4 UNION SELECT 5) nums;

-- 2. 金刚位配置 (每个频道5-8个)
DELETE FROM diamond;

-- 为每个频道创建金刚位
INSERT INTO diamond (tenant_id, channel_id, group_id, name, icon, title, link_type, link_value, sort, status, created_at, updated_at)
SELECT 
    1,
    channel_id,
    1,
    CASE n
        WHEN 1 THEN '热门推荐'
        WHEN 2 THEN '最新上架'
        WHEN 3 THEN '排行榜'
        WHEN 4 THEN '分类浏览'
        WHEN 5 THEN '我的收藏'
        WHEN 6 THEN '历史记录'
        WHEN 7 THEN '签到有礼'
        WHEN 8 THEN '活动中心'
    END,
    CONCAT('https://picsum.photos/seed/diamond', channel_id, '_', n, '/100/100'),
    CASE n
        WHEN 1 THEN '热门推荐'
        WHEN 2 THEN '最新上架'
        WHEN 3 THEN '排行榜'
        WHEN 4 THEN '分类浏览'
        WHEN 5 THEN '我的收藏'
        WHEN 6 THEN '历史记录'
        WHEN 7 THEN '签到有礼'
        WHEN 8 THEN '活动中心'
    END,
    1,
    CASE n
        WHEN 1 THEN 'recommend'
        WHEN 2 THEN 'latest'
        WHEN 3 THEN 'rank'
        WHEN 4 THEN 'category'
        WHEN 5 THEN 'favorite'
        WHEN 6 THEN 'history'
        WHEN 7 THEN 'checkin'
        WHEN 8 THEN 'activity'
    END,
    n,
    1,
    NOW(),
    NOW()
FROM 
    (SELECT id as channel_id FROM channel WHERE id <= 7) channels,
    (SELECT 1 as n UNION SELECT 2 UNION SELECT 3 UNION SELECT 4 UNION SELECT 5 UNION SELECT 6 UNION SELECT 7 UNION SELECT 8) nums;

-- 3. 推荐位配置 (每个频道3-5个)
DELETE FROM recommend;

-- 为每个频道创建推荐位
INSERT INTO recommend (tenant_id, channel_id, title, display_type, source_type, content_ids, sort, status, description, created_at, updated_at)
SELECT 
    1,
    channel_id,
    CASE n
        WHEN 1 THEN '热门推荐'
        WHEN 2 THEN '最新上架'
        WHEN 3 THEN '编辑精选'
        WHEN 4 THEN '猜你喜欢'
        WHEN 5 THEN '人气榜单'
    END,
    CASE n
        WHEN 1 THEN 'scroll'
        WHEN 2 THEN 'grid'
        WHEN 3 THEN 'single'
        WHEN 4 THEN 'waterfall'
        WHEN 5 THEN 'scroll'
    END,
    'manual',
    JSON_ARRAY(
        (channel_id - 1) * 200 + 1,
        (channel_id - 1) * 200 + 2,
        (channel_id - 1) * 200 + 3,
        (channel_id - 1) * 200 + 4,
        (channel_id - 1) * 200 + 5,
        (channel_id - 1) * 200 + 6,
        (channel_id - 1) * 200 + 7,
        (channel_id - 1) * 200 + 8
    ),
    n,
    1,
    CASE n
        WHEN 1 THEN '精选热门内容推荐'
        WHEN 2 THEN '最新上架内容推荐'
        WHEN 3 THEN '编辑精选优质内容'
        WHEN 4 THEN '根据你的喜好推荐'
        WHEN 5 THEN '人气榜单推荐'
    END,
    NOW(),
    NOW()
FROM 
    (SELECT id as channel_id FROM channel WHERE id <= 7) channels,
    (SELECT 1 as n UNION SELECT 2 UNION SELECT 3 UNION SELECT 4 UNION SELECT 5) nums;

-- 4. Feed流配置 (每个频道1个)
DELETE FROM feed_config;

-- 为每个频道创建Feed流
INSERT INTO feed_config (tenant_id, channel_id, name, title, layout_type, content_strategy, content_ids, status, sort, created_at, updated_at)
SELECT 
    1,
    c.id,
    CONCAT(c.name, 'Feed'),
    CONCAT(c.name, 'Feed流'),
    'waterfall',
    'manual',
    JSON_ARRAY(
        (c.id - 1) * 200 + 1,
        (c.id - 1) * 200 + 2,
        (c.id - 1) * 200 + 3,
        (c.id - 1) * 200 + 4,
        (c.id - 1) * 200 + 5,
        (c.id - 1) * 200 + 6,
        (c.id - 1) * 200 + 7,
        (c.id - 1) * 200 + 8,
        (c.id - 1) * 200 + 9,
        (c.id - 1) * 200 + 10,
        (c.id - 1) * 200 + 11,
        (c.id - 1) * 200 + 12,
        (c.id - 1) * 200 + 13,
        (c.id - 1) * 200 + 14,
        (c.id - 1) * 200 + 15,
        (c.id - 1) * 200 + 16,
        (c.id - 1) * 200 + 17,
        (c.id - 1) * 200 + 18,
        (c.id - 1) * 200 + 19,
        (c.id - 1) * 200 + 20
    ),
    1,
    1,
    NOW(),
    NOW()
FROM channel c WHERE c.id <= 7;

-- 5. 广告位配置 (每个频道2-3个)
DELETE FROM ad_position;

-- 为每个频道创建广告位
INSERT INTO ad_position (tenant_id, channel_id, name, code, type, image_url, link_url, status, sort, created_at, updated_at)
SELECT 
    1,
    channel_id,
    CASE n
        WHEN 1 THEN '信息流广告'
        WHEN 2 THEN 'Banner广告'
        WHEN 3 THEN '插屏广告'
    END,
    CONCAT('ad_', channel_id, '_', n),
    CASE n
        WHEN 1 THEN 1
        WHEN 2 THEN 1
        WHEN 3 THEN 2
    END,
    CONCAT('https://picsum.photos/seed/ad', channel_id, '_', n, '/400/300'),
    'https://example.com/ad',
    1,
    n,
    NOW(),
    NOW()
FROM 
    (SELECT id as channel_id FROM channel WHERE id <= 7) channels,
    (SELECT 1 as n UNION SELECT 2 UNION SELECT 3) nums;

-- 查看结果
SELECT 'Banner' as type, COUNT(*) as count FROM banner
UNION ALL
SELECT 'Diamond' as type, COUNT(*) as count FROM diamond
UNION ALL
SELECT 'Recommend' as type, COUNT(*) as count FROM recommend
UNION ALL
SELECT 'Feed Config' as type, COUNT(*) as count FROM feed_config
UNION ALL
SELECT 'Ad Position' as type, COUNT(*) as count FROM ad_position;
