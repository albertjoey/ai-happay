-- =============================================
-- Happy内容系统 - 完整测试数据生成脚本
-- 生成600条物料数据 + 频道配置
-- =============================================

SET NAMES utf8mb4;

-- 清理旧数据
DELETE FROM material_chapter;
DELETE FROM material WHERE id > 1;
DELETE FROM feed_config;
DELETE FROM recommend;
DELETE FROM diamond;
DELETE FROM banner;
DELETE FROM channel_config;

-- =============================================
-- 第一部分：生成物料数据（每种类型100条）
-- =============================================

-- 1. 图文内容 100条
INSERT INTO material (tenant_id, title, subtitle, type, cover_url, content_url, description, author, category, view_count, like_count, comment_count, share_count, collect_count, status, sort)
SELECT 
    1,
    CONCAT('图文', LPAD(n, 3, '0'), '：', 
        CASE (n % 10)
            WHEN 1 THEN '美食探店'
            WHEN 2 THEN '旅行攻略'
            WHEN 3 THEN '穿搭分享'
            WHEN 4 THEN '美妆教程'
            WHEN 5 THEN '健身打卡'
            WHEN 6 THEN '家居装修'
            WHEN 7 THEN '数码测评'
            WHEN 8 THEN '读书笔记'
            WHEN 9 THEN '职场干货'
            ELSE '生活记录'
        END
    ),
    CONCAT('精彩图文内容第', n, '期'),
    'image_text',
    CONCAT('https://picsum.photos/seed/img', n, '/400/600'),
    '',
    CONCAT('这是一篇关于', 
        CASE (n % 10)
            WHEN 1 THEN '美食探店'
            WHEN 2 THEN '旅行攻略'
            WHEN 3 THEN '穿搭分享'
            WHEN 4 THEN '美妆教程'
            WHEN 5 THEN '健身打卡'
            WHEN 6 THEN '家居装修'
            WHEN 7 THEN '数码测评'
            WHEN 8 THEN '读书笔记'
            WHEN 9 THEN '职场干货'
            ELSE '生活记录'
        END, '的精彩图文内容，带你发现生活中的美好。'),
    CASE (n % 5)
        WHEN 1 THEN '美食达人'
        WHEN 2 THEN '旅行博主'
        WHEN 3 THEN '时尚达人'
        WHEN 4 THEN '生活家'
        ELSE '内容创作者'
    END,
    CASE (n % 10)
        WHEN 1 THEN '美食'
        WHEN 2 THEN '旅行'
        WHEN 3 THEN '时尚'
        WHEN 4 THEN '美妆'
        WHEN 5 THEN '健身'
        WHEN 6 THEN '家居'
        WHEN 7 THEN '数码'
        WHEN 8 THEN '读书'
        WHEN 9 THEN '职场'
        ELSE '生活'
    END,
    FLOOR(1000 + RAND() * 50000),
    FLOOR(100 + RAND() * 5000),
    FLOOR(10 + RAND() * 500),
    FLOOR(5 + RAND() * 200),
    FLOOR(20 + RAND() * 1000),
    1,
    n
FROM (
    SELECT a.N + b.N * 10 + 1 as n
    FROM 
        (SELECT 0 as N UNION SELECT 1 UNION SELECT 2 UNION SELECT 3 UNION SELECT 4 UNION SELECT 5 UNION SELECT 6 UNION SELECT 7 UNION SELECT 8 UNION SELECT 9) a,
        (SELECT 0 as N UNION SELECT 1 UNION SELECT 2 UNION SELECT 3 UNION SELECT 4 UNION SELECT 5 UNION SELECT 6 UNION SELECT 7 UNION SELECT 8 UNION SELECT 9) b
) numbers
WHERE n <= 100;

-- 2. 长视频 100条
INSERT INTO material (tenant_id, title, subtitle, type, cover_url, content_url, description, author, category, view_count, like_count, comment_count, share_count, collect_count, duration, status, sort)
SELECT 
    1,
    CONCAT(
        CASE (n % 8)
            WHEN 1 THEN '电影'
            WHEN 2 THEN '电视剧'
            WHEN 3 THEN '纪录片'
            WHEN 4 THEN '综艺'
            WHEN 5 THEN '动漫'
            WHEN 6 THEN '公开课'
            WHEN 7 THEN '演唱会'
            ELSE '专题片'
        END, '：', 
        CASE (n % 20)
            WHEN 1 THEN '星际迷航'
            WHEN 2 THEN '都市情缘'
            WHEN 3 THEN '历史探秘'
            WHEN 4 THEN '欢乐喜剧人'
            WHEN 5 THEN '火影忍者'
            WHEN 6 THEN '人工智能导论'
            WHEN 7 THEN '周杰伦演唱会'
            WHEN 8 THEN '舌尖上的中国'
            WHEN 9 THEN '复仇者联盟'
            WHEN 10 THEN '三生三世'
            WHEN 11 THEN '宇宙奥秘'
            WHEN 12 THEN '奔跑吧兄弟'
            WHEN 13 THEN '海贼王'
            WHEN 14 THEN '经济学原理'
            WHEN 15 THEN '林俊杰演唱会'
            WHEN 16 THEN '美丽中国'
            WHEN 17 THEN '速度与激情'
            WHEN 18 THEN '琅琊榜'
            WHEN 19 THEN '自然传奇'
            ELSE '向往的生活'
        END, LPAD(n, 3, '0')
    ),
    CONCAT('精彩', 
        CASE (n % 8)
            WHEN 1 THEN '电影'
            WHEN 2 THEN '电视剧'
            WHEN 3 THEN '纪录片'
            WHEN 4 THEN '综艺'
            WHEN 5 THEN '动漫'
            WHEN 6 THEN '公开课'
            WHEN 7 THEN '演唱会'
            ELSE '专题片'
        END, '内容第', n, '期'),
    'long_video',
    CONCAT('https://picsum.photos/seed/longv', n, '/400/600'),
    'https://www.w3schools.com/html/mov_bbb.mp4',
    CONCAT('这是一部精彩绝伦的',
        CASE (n % 8)
            WHEN 1 THEN '电影'
            WHEN 2 THEN '电视剧'
            WHEN 3 THEN '纪录片'
            WHEN 4 THEN '综艺'
            WHEN 5 THEN '动漫'
            WHEN 6 THEN '公开课'
            WHEN 7 THEN '演唱会'
            ELSE '专题片'
        END, '，值得一看！'),
    CASE (n % 6)
        WHEN 1 THEN '华谊兄弟'
        WHEN 2 THEN '爱奇艺'
        WHEN 3 THEN '腾讯视频'
        WHEN 4 THEN '优酷'
        WHEN 5 THEN '哔哩哔哩'
        ELSE '芒果TV'
    END,
    CASE (n % 8)
        WHEN 1 THEN '电影'
        WHEN 2 THEN '电视剧'
        WHEN 3 THEN '纪录片'
        WHEN 4 THEN '综艺'
        WHEN 5 THEN '动漫'
        WHEN 6 THEN '教育'
        WHEN 7 THEN '音乐'
        ELSE '纪实'
    END,
    FLOOR(5000 + RAND() * 100000),
    FLOOR(500 + RAND() * 20000),
    FLOOR(100 + RAND() * 3000),
    FLOOR(50 + RAND() * 1000),
    FLOOR(200 + RAND() * 5000),
    FLOOR(3600 + RAND() * 7200),
    1,
    n
FROM (
    SELECT a.N + b.N * 10 + 1 as n
    FROM 
        (SELECT 0 as N UNION SELECT 1 UNION SELECT 2 UNION SELECT 3 UNION SELECT 4 UNION SELECT 5 UNION SELECT 6 UNION SELECT 7 UNION SELECT 8 UNION SELECT 9) a,
        (SELECT 0 as N UNION SELECT 1 UNION SELECT 2 UNION SELECT 3 UNION SELECT 4 UNION SELECT 5 UNION SELECT 6 UNION SELECT 7 UNION SELECT 8 UNION SELECT 9) b
) numbers
WHERE n <= 100;

-- 3. 短视频 100条
INSERT INTO material (tenant_id, title, subtitle, type, cover_url, content_url, description, author, category, view_count, like_count, comment_count, share_count, collect_count, duration, status, sort)
SELECT 
    1,
    CONCAT(
        CASE (n % 10)
            WHEN 1 THEN '搞笑'
            WHEN 2 THEN '美食'
            WHEN 3 THEN '萌宠'
            WHEN 4 THEN '舞蹈'
            WHEN 5 THEN '音乐'
            WHEN 6 THEN '运动'
            WHEN 7 THEN '旅行'
            WHEN 8 THEN '美妆'
            WHEN 9 THEN '游戏'
            ELSE '生活'
        END, '短视频', LPAD(n, 3, '0')
    ),
    CONCAT('超火的', 
        CASE (n % 10)
            WHEN 1 THEN '搞笑'
            WHEN 2 THEN '美食'
            WHEN 3 THEN '萌宠'
            WHEN 4 THEN '舞蹈'
            WHEN 5 THEN '音乐'
            WHEN 6 THEN '运动'
            WHEN 7 THEN '旅行'
            WHEN 8 THEN '美妆'
            WHEN 9 THEN '游戏'
            ELSE '生活'
        END, '内容'),
    'short_video',
    CONCAT('https://picsum.photos/seed/shortv', n, '/400/600'),
    'https://www.w3schools.com/html/mov_bbb.mp4',
    CONCAT('今日份', 
        CASE (n % 10)
            WHEN 1 THEN '快乐源泉'
            WHEN 2 THEN '美食教程'
            WHEN 3 THEN '萌宠日常'
            WHEN 4 THEN '舞蹈教学'
            WHEN 5 THEN '音乐分享'
            WHEN 6 THEN '运动打卡'
            WHEN 7 THEN '旅行vlog'
            WHEN 8 THEN '美妆教程'
            WHEN 9 THEN '游戏精彩'
            ELSE '生活记录'
        END, '，快来看看吧！'),
    CONCAT('达人', LPAD(n, 3, '0'), '号'),
    CASE (n % 10)
        WHEN 1 THEN '搞笑'
        WHEN 2 THEN '美食'
        WHEN 3 THEN '萌宠'
        WHEN 4 THEN '舞蹈'
        WHEN 5 THEN '音乐'
        WHEN 6 THEN '运动'
        WHEN 7 THEN '旅行'
        WHEN 8 THEN '美妆'
        WHEN 9 THEN '游戏'
        ELSE '生活'
    END,
    FLOOR(10000 + RAND() * 200000),
    FLOOR(1000 + RAND() * 30000),
    FLOOR(200 + RAND() * 5000),
    FLOOR(100 + RAND() * 2000),
    FLOOR(500 + RAND() * 8000),
    FLOOR(15 + RAND() * 180),
    1,
    n
FROM (
    SELECT a.N + b.N * 10 + 1 as n
    FROM 
        (SELECT 0 as N UNION SELECT 1 UNION SELECT 2 UNION SELECT 3 UNION SELECT 4 UNION SELECT 5 UNION SELECT 6 UNION SELECT 7 UNION SELECT 8 UNION SELECT 9) a,
        (SELECT 0 as N UNION SELECT 1 UNION SELECT 2 UNION SELECT 3 UNION SELECT 4 UNION SELECT 5 UNION SELECT 6 UNION SELECT 7 UNION SELECT 8 UNION SELECT 9) b
) numbers
WHERE n <= 100;

-- 4. 小说 100条
INSERT INTO material (tenant_id, title, subtitle, type, cover_url, content_url, description, author, category, view_count, like_count, comment_count, share_count, collect_count, word_count, chapter_count, status, sort)
SELECT 
    1,
    CONCAT(
        CASE (n % 12)
            WHEN 1 THEN '斗破苍穹'
            WHEN 2 THEN '诡秘之主'
            WHEN 3 THEN '凡人修仙传'
            WHEN 4 THEN '庆余年'
            WHEN 5 THEN '赘婿'
            WHEN 6 THEN '大奉打更人'
            WHEN 7 THEN '遮天'
            WHEN 8 THEN '完美世界'
            WHEN 9 THEN '星辰变'
            WHEN 10 THEN '盘龙'
            WHEN 11 THEN '雪中悍刀行'
            ELSE '剑来'
        END, '之', 
        CASE (n % 8)
            WHEN 1 THEN '觉醒篇'
            WHEN 2 THEN '崛起篇'
            WHEN 3 THEN '争霸篇'
            WHEN 4 THEN '巅峰篇'
            WHEN 5 THEN '传承篇'
            WHEN 6 THEN '复仇篇'
            WHEN 7 THEN '重生篇'
            ELSE '终章篇'
        END, LPAD(n, 3, '0')
    ),
    CONCAT(
        CASE (n % 6)
            WHEN 1 THEN '热血爽文'
            WHEN 2 THEN '玄幻巨制'
            WHEN 3 THEN '都市传奇'
            WHEN 4 THEN '仙侠经典'
            WHEN 5 THEN '历史架空'
            ELSE '科幻未来'
        END, '，精彩纷呈'),
    'novel',
    CONCAT('https://picsum.photos/seed/novel', n, '/400/600'),
    '',
    CONCAT('这是一部',
        CASE (n % 6)
            WHEN 1 THEN '热血爽文'
            WHEN 2 THEN '玄幻巨制'
            WHEN 3 THEN '都市传奇'
            WHEN 4 THEN '仙侠经典'
            WHEN 5 THEN '历史架空'
            ELSE '科幻未来'
        END, '，讲述主角从默默无闻到傲视天下的传奇故事。跌宕起伏的剧情，精彩绝伦的打斗，让人欲罢不能！'),
    CONCAT('作家', LPAD(n, 3, '0'), '号'),
    CASE (n % 6)
        WHEN 1 THEN '玄幻'
        WHEN 2 THEN '奇幻'
        WHEN 3 THEN '都市'
        WHEN 4 THEN '仙侠'
        WHEN 5 THEN '历史'
        ELSE '科幻'
    END,
    FLOOR(50000 + RAND() * 500000),
    FLOOR(5000 + RAND() * 50000),
    FLOOR(1000 + RAND() * 10000),
    FLOOR(500 + RAND() * 5000),
    FLOOR(2000 + RAND() * 20000),
    FLOOR(500000 + RAND() * 2000000),
    FLOOR(100 + RAND() * 500),
    1,
    n
FROM (
    SELECT a.N + b.N * 10 + 1 as n
    FROM 
        (SELECT 0 as N UNION SELECT 1 UNION SELECT 2 UNION SELECT 3 UNION SELECT 4 UNION SELECT 5 UNION SELECT 6 UNION SELECT 7 UNION SELECT 8 UNION SELECT 9) a,
        (SELECT 0 as N UNION SELECT 1 UNION SELECT 2 UNION SELECT 3 UNION SELECT 4 UNION SELECT 5 UNION SELECT 6 UNION SELECT 7 UNION SELECT 8 UNION SELECT 9) b
) numbers
WHERE n <= 100;

-- 5. 短剧 100条
INSERT INTO material (tenant_id, title, subtitle, type, cover_url, content_url, description, author, category, view_count, like_count, comment_count, share_count, collect_count, duration, status, sort)
SELECT 
    1,
    CONCAT(
        CASE (n % 15)
            WHEN 1 THEN '霸道总裁爱上我'
            WHEN 2 THEN '重生之复仇女王'
            WHEN 3 THEN '穿越之古代奇缘'
            WHEN 4 THEN '逆袭人生'
            WHEN 5 THEN '都市情缘'
            WHEN 6 THEN '悬疑迷局'
            WHEN 7 THEN '甜蜜恋爱'
            WHEN 8 THEN '豪门恩怨'
            WHEN 9 THEN '校园青春'
            WHEN 10 THEN '职场逆袭'
            WHEN 11 THEN '仙侠情缘'
            WHEN 12 THEN '民国奇闻'
            WHEN 13 THEN '科幻未来'
            WHEN 14 THEN '家庭伦理'
            ELSE '古装宫斗'
        END, LPAD(n, 3, '0')
    ),
    CONCAT(
        CASE (n % 5)
            WHEN 1 THEN '甜宠短剧'
            WHEN 2 THEN '逆袭爽剧'
            WHEN 3 THEN '悬疑烧脑'
            WHEN 4 THEN '古装言情'
            ELSE '都市情感'
        END, '，每集3分钟'),
    'short_drama',
    CONCAT('https://picsum.photos/seed/sdrama', n, '/400/600'),
    'https://www.w3schools.com/html/mov_bbb.mp4',
    CONCAT('超火的',
        CASE (n % 5)
            WHEN 1 THEN '甜宠短剧'
            WHEN 2 THEN '逆袭爽剧'
            WHEN 3 THEN '悬疑烧脑'
            WHEN 4 THEN '古装言情'
            ELSE '都市情感'
        END, '，剧情紧凑，高潮迭起，一口气看完！'),
    CONCAT('短剧工场', LPAD(n, 3, '0')),
    CASE (n % 5)
        WHEN 1 THEN '甜宠'
        WHEN 2 THEN '逆袭'
        WHEN 3 THEN '悬疑'
        WHEN 4 THEN '古装'
        ELSE '都市'
    END,
    FLOOR(20000 + RAND() * 300000),
    FLOOR(2000 + RAND() * 40000),
    FLOOR(500 + RAND() * 8000),
    FLOOR(200 + RAND() * 3000),
    FLOOR(1000 + RAND() * 15000),
    FLOOR(180 + RAND() * 600),
    1,
    n
FROM (
    SELECT a.N + b.N * 10 + 1 as n
    FROM 
        (SELECT 0 as N UNION SELECT 1 UNION SELECT 2 UNION SELECT 3 UNION SELECT 4 UNION SELECT 5 UNION SELECT 6 UNION SELECT 7 UNION SELECT 8 UNION SELECT 9) a,
        (SELECT 0 as N UNION SELECT 1 UNION SELECT 2 UNION SELECT 3 UNION SELECT 4 UNION SELECT 5 UNION SELECT 6 UNION SELECT 7 UNION SELECT 8 UNION SELECT 9) b
) numbers
WHERE n <= 100;

-- 6. 漫剧 100条
INSERT INTO material (tenant_id, title, subtitle, type, cover_url, content_url, description, author, category, view_count, like_count, comment_count, share_count, collect_count, chapter_count, status, sort)
SELECT 
    1,
    CONCAT(
        CASE (n % 12)
            WHEN 1 THEN '绝世武神'
            WHEN 2 THEN '都市最强战神'
            WHEN 3 THEN '甜蜜恋爱日常'
            WHEN 4 THEN '修仙从打工开始'
            WHEN 5 THEN '重生之都市仙尊'
            WHEN 6 THEN '我的美女总裁'
            WHEN 7 THEN '龙王传说'
            WHEN 8 THEN '万界独尊'
            WHEN 9 THEN '绝世唐门'
            WHEN 10 THEN '斗罗大陆'
            WHEN 11 THEN '武动乾坤'
            ELSE '大主宰'
        END, LPAD(n, 3, '0')
    ),
    CONCAT(
        CASE (n % 4)
            WHEN 1 THEN '热血漫剧'
            WHEN 2 THEN '恋爱漫剧'
            WHEN 3 THEN '玄幻漫剧'
            ELSE '都市漫剧'
        END, '，精彩不断'),
    'drama',
    CONCAT('https://picsum.photos/seed/drama', n, '/400/600'),
    '',
    CONCAT('超人气',
        CASE (n % 4)
            WHEN 1 THEN '热血漫剧'
            WHEN 2 THEN '恋爱漫剧'
            WHEN 3 THEN '玄幻漫剧'
            ELSE '都市漫剧'
        END, '，画风精美，剧情精彩，追更必备！'),
    CONCAT('漫画工作室', LPAD(n, 3, '0')),
    CASE (n % 4)
        WHEN 1 THEN '热血'
        WHEN 2 THEN '恋爱'
        WHEN 3 THEN '玄幻'
        ELSE '都市'
    END,
    FLOOR(30000 + RAND() * 400000),
    FLOOR(3000 + RAND() * 50000),
    FLOOR(800 + RAND() * 10000),
    FLOOR(400 + RAND() * 4000),
    FLOOR(1500 + RAND() * 20000),
    FLOOR(50 + RAND() * 200),
    1,
    n
FROM (
    SELECT a.N + b.N * 10 + 1 as n
    FROM 
        (SELECT 0 as N UNION SELECT 1 UNION SELECT 2 UNION SELECT 3 UNION SELECT 4 UNION SELECT 5 UNION SELECT 6 UNION SELECT 7 UNION SELECT 8 UNION SELECT 9) a,
        (SELECT 0 as N UNION SELECT 1 UNION SELECT 2 UNION SELECT 3 UNION SELECT 4 UNION SELECT 5 UNION SELECT 6 UNION SELECT 7 UNION SELECT 8 UNION SELECT 9) b
) numbers
WHERE n <= 100;

SELECT '物料数据生成完成' as status;
