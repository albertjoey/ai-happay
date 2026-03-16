-- 生成大量物料数据
-- 每种类型200+条,总计1200+条

-- 清空现有数据
DELETE FROM material WHERE id > 0;
ALTER TABLE material AUTO_INCREMENT = 1;

-- 1. 图文内容 (210条)
INSERT INTO material (tenant_id, title, subtitle, type, cover_url, content_url, description, author, category, view_count, like_count, comment_count, share_count, collect_count, word_count, status, sort, created_at, updated_at)
SELECT
    1,
    CONCAT('美食制作教程：第', n, '期精彩内容'),
    '专业分享，值得收藏',
    'image_text',
    CONCAT('https://picsum.photos/seed/it', n, '/400/600'),
    '',
    CONCAT('这是一篇关于美食的精彩内容，包含详细的步骤和实用的技巧，帮助你快速提升！第', n, '期'),
    '美食达人',
    '美食',
    FLOOR(10000 + RAND() * 490000),
    FLOOR(1000 + RAND() * 49000),
    FLOOR(100 + RAND() * 4900),
    FLOOR(50 + RAND() * 1950),
    FLOOR(500 + RAND() * 29500),
    FLOOR(1500 + RAND() * 3500),
    1,
    100 - (n % 10),
    NOW(),
    NOW()
FROM (
    SELECT a.N + b.N * 10 + c.N * 100 + 1 as n
    FROM
        (SELECT 0 AS N UNION SELECT 1 UNION SELECT 2 UNION SELECT 3 UNION SELECT 4 UNION SELECT 5 UNION SELECT 6 UNION SELECT 7 UNION SELECT 8 UNION SELECT 9) a,
        (SELECT 0 AS N UNION SELECT 1 UNION SELECT 2 UNION SELECT 3 UNION SELECT 4 UNION SELECT 5 UNION SELECT 6 UNION SELECT 7 UNION SELECT 8 UNION SELECT 9) b,
        (SELECT 0 AS N UNION SELECT 1) c
) numbers
WHERE n <= 210;

-- 2. 短视频内容 (210条)
INSERT INTO material (tenant_id, title, subtitle, type, cover_url, content_url, description, author, category, view_count, like_count, comment_count, share_count, collect_count, duration, status, sort, created_at, updated_at)
SELECT
    1,
    CONCAT('搞笑视频合集：第', n, '集精彩呈现'),
    '精彩视频，不容错过',
    'video',
    CONCAT('https://picsum.photos/seed/v', n, '/400/600'),
    CONCAT('https://example.com/video/v', n, '.mp4'),
    CONCAT('这是一个关于搞笑的精彩视频，内容丰富有趣，快来看看吧！第', n, '集'),
    '搞笑达人',
    '搞笑',
    FLOOR(50000 + RAND() * 1950000),
    FLOOR(5000 + RAND() * 195000),
    FLOOR(500 + RAND() * 19500),
    FLOOR(100 + RAND() * 9900),
    FLOOR(1000 + RAND() * 99000),
    FLOOR(60 + RAND() * 1740),
    1,
    100 - (n % 10),
    NOW(),
    NOW()
FROM (
    SELECT a.N + b.N * 10 + c.N * 100 + 1 as n
    FROM
        (SELECT 0 AS N UNION SELECT 1 UNION SELECT 2 UNION SELECT 3 UNION SELECT 4 UNION SELECT 5 UNION SELECT 6 UNION SELECT 7 UNION SELECT 8 UNION SELECT 9) a,
        (SELECT 0 AS N UNION SELECT 1 UNION SELECT 2 UNION SELECT 3 UNION SELECT 4 UNION SELECT 5 UNION SELECT 6 UNION SELECT 7 UNION SELECT 8 UNION SELECT 9) b,
        (SELECT 0 AS N UNION SELECT 1) c
) numbers
WHERE n <= 210;

-- 3. 小说内容 (210条)
INSERT INTO material (tenant_id, title, subtitle, type, cover_url, content_url, description, author, category, view_count, like_count, comment_count, share_count, collect_count, word_count, chapter_count, status, sort, created_at, updated_at)
SELECT
    1,
    CONCAT('重生之商业帝国（第', n, '部）'),
    '精彩小说，值得追读',
    'novel',
    CONCAT('https://picsum.photos/seed/n', n, '/400/600'),
    '',
    CONCAT('这是一部关于都市重生的精彩小说，情节跌宕起伏，人物形象鲜明，绝对让你欲罢不能！第', n, '部'),
    '商战大神',
    '都市重生',
    FLOOR(100000 + RAND() * 4900000),
    FLOOR(10000 + RAND() * 490000),
    FLOOR(1000 + RAND() * 99000),
    FLOOR(500 + RAND() * 49500),
    FLOOR(10000 + RAND() * 490000),
    FLOOR(1000000 + RAND() * 5000000),
    FLOOR(500 + RAND() * 2500),
    1,
    100 - (n % 10),
    NOW(),
    NOW()
FROM (
    SELECT a.N + b.N * 10 + c.N * 100 + 1 as n
    FROM
        (SELECT 0 AS N UNION SELECT 1 UNION SELECT 2 UNION SELECT 3 UNION SELECT 4 UNION SELECT 5 UNION SELECT 6 UNION SELECT 7 UNION SELECT 8 UNION SELECT 9) a,
        (SELECT 0 AS N UNION SELECT 1 UNION SELECT 2 UNION SELECT 3 UNION SELECT 4 UNION SELECT 5 UNION SELECT 6 UNION SELECT 7 UNION SELECT 8 UNION SELECT 9) b,
        (SELECT 0 AS N UNION SELECT 1) c
) numbers
WHERE n <= 210;

-- 4. 漫画内容 (210条)
INSERT INTO material (tenant_id, title, subtitle, type, cover_url, content_url, description, author, category, view_count, like_count, comment_count, share_count, collect_count, chapter_count, status, sort, created_at, updated_at)
SELECT
    1,
    CONCAT('斗罗大陆（第', n, '季）'),
    '精彩漫画，热血呈现',
    'comic',
    CONCAT('https://picsum.photos/seed/c', n, '/400/600'),
    '',
    CONCAT('这是一部关于玄幻热血的精彩漫画，画风精美，剧情热血，绝对让你爱不释手！第', n, '季'),
    '唐家三少',
    '玄幻热血',
    FLOOR(500000 + RAND() * 4500000),
    FLOOR(50000 + RAND() * 450000),
    FLOOR(5000 + RAND() * 95000),
    FLOOR(2000 + RAND() * 48000),
    FLOOR(50000 + RAND() * 450000),
    FLOOR(300 + RAND() * 700),
    1,
    100 - (n % 10),
    NOW(),
    NOW()
FROM (
    SELECT a.N + b.N * 10 + c.N * 100 + 1 as n
    FROM
        (SELECT 0 AS N UNION SELECT 1 UNION SELECT 2 UNION SELECT 3 UNION SELECT 4 UNION SELECT 5 UNION SELECT 6 UNION SELECT 7 UNION SELECT 8 UNION SELECT 9) a,
        (SELECT 0 AS N UNION SELECT 1 UNION SELECT 2 UNION SELECT 3 UNION SELECT 4 UNION SELECT 5 UNION SELECT 6 UNION SELECT 7 UNION SELECT 8 UNION SELECT 9) b,
        (SELECT 0 AS N UNION SELECT 1) c
) numbers
WHERE n <= 210;

-- 5. 漫剧内容 (210条)
INSERT INTO material (tenant_id, title, subtitle, type, cover_url, content_url, description, author, category, view_count, like_count, comment_count, share_count, collect_count, chapter_count, status, sort, created_at, updated_at)
SELECT
    1,
    CONCAT('霸道总裁爱上我（第', n, '集）'),
    '精彩漫剧，不容错过',
    'manhua',
    CONCAT('https://picsum.photos/seed/m', n, '/400/600'),
    '',
    CONCAT('这是一部关于都市甜宠的精彩漫剧，剧情精彩，人物生动，绝对让你欲罢不能！第', n, '集'),
    '甜宠剧场',
    '都市甜宠',
    FLOOR(200000 + RAND() * 1800000),
    FLOOR(20000 + RAND() * 180000),
    FLOOR(2000 + RAND() * 48000),
    FLOOR(500 + RAND() * 19500),
    FLOOR(20000 + RAND() * 180000),
    FLOOR(50 + RAND() * 250),
    1,
    100 - (n % 10),
    NOW(),
    NOW()
FROM (
    SELECT a.N + b.N * 10 + c.N * 100 + 1 as n
    FROM
        (SELECT 0 AS N UNION SELECT 1 UNION SELECT 2 UNION SELECT 3 UNION SELECT 4 UNION SELECT 5 UNION SELECT 6 UNION SELECT 7 UNION SELECT 8 UNION SELECT 9) a,
        (SELECT 0 AS N UNION SELECT 1 UNION SELECT 2 UNION SELECT 3 UNION SELECT 4 UNION SELECT 5 UNION SELECT 6 UNION SELECT 7 UNION SELECT 8 UNION SELECT 9) b,
        (SELECT 0 AS N UNION SELECT 1) c
) numbers
WHERE n <= 210;

-- 6. 短剧内容 (210条)
INSERT INTO material (tenant_id, title, subtitle, type, cover_url, content_url, description, author, category, view_count, like_count, comment_count, share_count, collect_count, duration, status, sort, created_at, updated_at)
SELECT
    1,
    CONCAT('逆袭人生（第', n, '集）'),
    '精彩短剧，值得追看',
    'short_drama',
    CONCAT('https://picsum.photos/seed/sd', n, '/400/600'),
    CONCAT('https://example.com/drama/sd', n, '.mp4'),
    CONCAT('这是一部关于励志的精彩短剧，剧情紧凑，演技在线，绝对让你看得停不下来！第', n, '集'),
    '导演A',
    '励志',
    FLOOR(300000 + RAND() * 2700000),
    FLOOR(30000 + RAND() * 270000),
    FLOOR(3000 + RAND() * 57000),
    FLOOR(1000 + RAND() * 29000),
    FLOOR(30000 + RAND() * 270000),
    FLOOR(60 + RAND() * 240),
    1,
    100 - (n % 10),
    NOW(),
    NOW()
FROM (
    SELECT a.N + b.N * 10 + c.N * 100 + 1 as n
    FROM
        (SELECT 0 AS N UNION SELECT 1 UNION SELECT 2 UNION SELECT 3 UNION SELECT 4 UNION SELECT 5 UNION SELECT 6 UNION SELECT 7 UNION SELECT 8 UNION SELECT 9) a,
        (SELECT 0 AS N UNION SELECT 1 UNION SELECT 2 UNION SELECT 3 UNION SELECT 4 UNION SELECT 5 UNION SELECT 6 UNION SELECT 7 UNION SELECT 8 UNION SELECT 9) b,
        (SELECT 0 AS N UNION SELECT 1) c
) numbers
WHERE n <= 210;

-- 查看结果
SELECT type, COUNT(*) as count FROM material GROUP BY type ORDER BY type;
SELECT COUNT(*) as total FROM material;
