-- 物料表(图文、小说、视频、banner、漫剧、短剧)
CREATE TABLE IF NOT EXISTS `material` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `tenant_id` bigint unsigned NOT NULL DEFAULT 1 COMMENT '租户ID',
  `title` varchar(200) NOT NULL COMMENT '标题',
  `subtitle` varchar(500) DEFAULT '' COMMENT '副标题',
  `type` varchar(20) NOT NULL COMMENT '类型: image_text(图文), novel(小说), video(视频), banner(横幅), comic(漫剧), short_drama(短剧)',
  `cover_url` varchar(500) DEFAULT '' COMMENT '封面图URL',
  `content_url` varchar(500) DEFAULT '' COMMENT '内容URL(视频/音频等)',
  `description` text COMMENT '描述/简介',
  `author` varchar(100) DEFAULT '' COMMENT '作者',
  `tags` json DEFAULT NULL COMMENT '标签JSON数组',
  `category` varchar(50) DEFAULT '' COMMENT '分类',
  `view_count` int unsigned DEFAULT 0 COMMENT '浏览数',
  `like_count` int unsigned DEFAULT 0 COMMENT '点赞数',
  `comment_count` int unsigned DEFAULT 0 COMMENT '评论数',
  `share_count` int unsigned DEFAULT 0 COMMENT '分享数',
  `collect_count` int unsigned DEFAULT 0 COMMENT '收藏数',
  `duration` int unsigned DEFAULT 0 COMMENT '时长(秒)',
  `word_count` int unsigned DEFAULT 0 COMMENT '字数',
  `chapter_count` int unsigned DEFAULT 0 COMMENT '章节数',
  `status` tinyint NOT NULL DEFAULT 1 COMMENT '状态: 0-草稿, 1-已发布, 2-已下架',
  `sort` int NOT NULL DEFAULT 0 COMMENT '排序',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_tenant_id` (`tenant_id`),
  KEY `idx_type` (`type`),
  KEY `idx_status` (`status`),
  KEY `idx_created_at` (`created_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='物料表';

-- 插入测试数据
INSERT INTO `material` (`tenant_id`, `title`, `subtitle`, `type`, `cover_url`, `content_url`, `description`, `author`, `tags`, `category`, `view_count`, `like_count`, `comment_count`, `share_count`, `collect_count`, `duration`, `word_count`, `chapter_count`, `status`, `sort`) VALUES
-- 图文内容
(1, '10个让你变美的小技巧', '简单实用的美容方法', 'image_text', 'https://picsum.photos/400/600?random=1', '', '分享10个简单实用的美容小技巧，让你轻松变美！', '美妆达人', '["美容", "护肤", "变美"]', '美容', 12580, 3560, 892, 456, 1234, 0, 1500, 0, 1, 100),
(1, '旅行必备清单，出发前必看', '再也不怕忘带东西', 'image_text', 'https://picsum.photos/400/600?random=2', '', '超详细的旅行必备清单，涵盖各种场景，让你的旅行更轻松！', '旅行家', '["旅行", "攻略", "清单"]', '旅行', 8920, 2340, 567, 890, 678, 0, 2000, 0, 1, 90),
(1, '美食制作教程：家常菜系列', '简单易学的家常菜', 'image_text', 'https://picsum.photos/400/600?random=3', '', '教你制作美味家常菜，简单易学，新手也能轻松上手！', '美食博主', '["美食", "家常菜", "教程"]', '美食', 15670, 4560, 1230, 678, 2345, 0, 1800, 0, 1, 80),

-- 小说内容
(1, '都市修仙传', '一个普通人的修仙之路', 'novel', 'https://picsum.photos/400/600?random=4', '', '讲述一个普通都市青年意外获得修仙传承，从此踏上修仙之路的故事。', '修仙作者', '["修仙", "都市", "玄幻"]', '玄幻', 45670, 12340, 3456, 2345, 8901, 0, 580000, 1200, 1, 100),
(1, '霸道总裁的小娇妻', '甜宠文来袭', 'novel', 'https://picsum.photos/400/600?random=5', '', '霸道总裁与小白兔的甜蜜爱情故事，全程高甜无虐！', '言情作家', '["言情", "甜宠", "总裁"]', '言情', 34560, 9870, 2345, 1567, 6789, 0, 420000, 890, 1, 95),
(1, '末世重生：我有一个空间', '末世生存指南', 'novel', 'https://picsum.photos/400/600?random=6', '', '重生末世前，获得随身空间，囤积物资，在末世中生存的故事。', '末世作者', '["末世", "重生", "空间"]', '科幻', 28900, 7650, 1890, 1234, 5678, 0, 350000, 750, 1, 90),

-- 视频内容
(1, '搞笑视频合集：笑到肚子疼', '精选搞笑瞬间', 'video', 'https://picsum.photos/400/600?random=7', 'https://example.com/video1.mp4', '精选网络搞笑视频，让你笑到停不下来！', '搞笑达人', '["搞笑", "娱乐", "视频"]', '娱乐', 89000, 23450, 6780, 4567, 12340, 180, 0, 0, 1, 100),
(1, '健身教程：30天练出马甲线', '科学健身方法', 'video', 'https://picsum.photos/400/600?random=8', 'https://example.com/video2.mp4', '专业健身教练教你30天练出马甲线，科学有效！', '健身教练', '["健身", "减肥", "马甲线"]', '健身', 56780, 15670, 4560, 2345, 8901, 1200, 0, 0, 1, 95),
(1, '旅行Vlog：云南之旅', '记录美好时光', 'video', 'https://picsum.photos/400/600?random=9', 'https://example.com/video3.mp4', '带你领略云南的美景美食，感受旅行的魅力！', '旅行博主', '["旅行", "云南", "Vlog"]', '旅行', 34560, 9870, 2345, 1567, 5678, 600, 0, 0, 1, 90),

-- Banner内容
(1, '双十一大促', '全场五折起', 'banner', 'https://picsum.photos/800/400?random=10', 'https://example.com/promotion', '双十一购物狂欢节，全场商品五折起，错过等一年！', '官方活动', '["促销", "双十一", "购物"]', '活动', 123456, 34567, 8901, 6789, 23456, 0, 0, 0, 1, 100),
(1, '新年特惠活动', '新年新气象', 'banner', 'https://picsum.photos/800/400?random=11', 'https://example.com/newyear', '新年特惠活动，精选商品限时折扣，快来抢购！', '官方活动', '["新年", "特惠", "活动"]', '活动', 98760, 23456, 6789, 4567, 19012, 0, 0, 0, 1, 95),

-- 漫剧内容
(1, '霸道总裁爱上我', '甜宠漫剧', 'comic', 'https://picsum.photos/400/600?random=12', '', '霸道总裁与职场小白的甜蜜爱情故事，画风精美，剧情甜宠！', '漫画家A', '["漫剧", "甜宠", "总裁"]', '言情', 67890, 19012, 5678, 3456, 12345, 0, 0, 120, 1, 100),
(1, '修仙之路', '玄幻漫剧', 'comic', 'https://picsum.photos/400/600?random=13', '', '少年踏上修仙之路，历经磨难，最终成为一代强者！', '漫画家B', '["漫剧", "修仙", "玄幻"]', '玄幻', 56780, 15670, 4560, 2345, 10901, 0, 0, 150, 1, 95),
(1, '校园恋爱物语', '青春漫剧', 'comic', 'https://picsum.photos/400/600?random=14', '', '校园青春恋爱故事，纯真美好的校园时光！', '漫画家C', '["漫剧", "校园", "恋爱"]', '青春', 45670, 12340, 3456, 1890, 9876, 0, 0, 100, 1, 90),

-- 短剧内容
(1, '逆袭人生', '励志短剧', 'short_drama', 'https://picsum.photos/400/600?random=15', 'https://example.com/drama1.mp4', '讲述一个普通人通过努力逆袭成功的故事，励志感人！', '导演A', '["短剧", "励志", "逆袭"]', '励志', 78900, 23456, 6789, 4567, 19012, 3600, 0, 30, 1, 100),
(1, '都市情缘', '都市爱情短剧', 'short_drama', 'https://picsum.photos/400/600?random=16', 'https://example.com/drama2.mp4', '都市男女的爱情故事，浪漫温馨，让人心动！', '导演B', '["短剧", "都市", "爱情"]', '爱情', 67890, 19012, 5678, 3456, 15678, 2700, 0, 25, 1, 95),
(1, '悬疑迷局', '悬疑短剧', 'short_drama', 'https://picsum.photos/400/600?random=17', 'https://example.com/drama3.mp4', '扑朔迷离的悬疑案件，烧脑剧情，让你欲罢不能！', '导演C', '["短剧", "悬疑", "烧脑"]', '悬疑', 56780, 15670, 4560, 2345, 12345, 4500, 0, 40, 1, 90);

-- 查询数据验证
SELECT * FROM material WHERE status = 1 ORDER BY sort DESC, created_at DESC LIMIT 20;
