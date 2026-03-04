-- 第三阶段数据表创建脚本

-- 金刚位表
CREATE TABLE IF NOT EXISTS `diamond` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `tenant_id` bigint unsigned NOT NULL DEFAULT '1',
  `channel_id` bigint unsigned NOT NULL,
  `group_id` int NOT NULL DEFAULT '1' COMMENT '分组ID (1-5组)',
  `sort` int NOT NULL DEFAULT '0' COMMENT '组内排序',
  `title` varchar(50) NOT NULL COMMENT '标题',
  `icon` varchar(255) DEFAULT NULL COMMENT '图标',
  `link_type` varchar(20) NOT NULL COMMENT '链接类型: channel/topic/content/external',
  `link_value` varchar(500) DEFAULT NULL COMMENT '链接值',
  `status` tinyint NOT NULL DEFAULT '1' COMMENT '状态: 0-禁用 1-启用',
  `description` varchar(200) DEFAULT NULL COMMENT '描述',
  PRIMARY KEY (`id`),
  KEY `idx_diamond_deleted_at` (`deleted_at`),
  KEY `idx_diamond_tenant_id` (`tenant_id`),
  KEY `idx_diamond_channel_id` (`channel_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='金刚位配置表';

-- 推荐位表
CREATE TABLE IF NOT EXISTS `recommend` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `tenant_id` bigint unsigned NOT NULL DEFAULT '1',
  `channel_id` bigint unsigned NOT NULL,
  `title` varchar(100) NOT NULL COMMENT '标题',
  `display_type` varchar(20) NOT NULL COMMENT '展示类型: single/scroll/grid',
  `source_type` varchar(20) NOT NULL COMMENT '内容来源: manual/algorithm/filter',
  `content_ids` json DEFAULT NULL COMMENT '内容ID列表(人工选择)',
  `filter_rule` json DEFAULT NULL COMMENT '筛选规则(条件筛选)',
  `sort` int NOT NULL DEFAULT '0' COMMENT '排序',
  `status` tinyint NOT NULL DEFAULT '1' COMMENT '状态: 0-禁用 1-启用',
  `description` varchar(200) DEFAULT NULL COMMENT '描述',
  PRIMARY KEY (`id`),
  KEY `idx_recommend_deleted_at` (`deleted_at`),
  KEY `idx_recommend_tenant_id` (`tenant_id`),
  KEY `idx_recommend_channel_id` (`channel_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='推荐位配置表';

-- Feed流配置表
CREATE TABLE IF NOT EXISTS `feed_config` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `tenant_id` bigint unsigned NOT NULL DEFAULT '1',
  `channel_id` bigint unsigned NOT NULL,
  `title` varchar(100) NOT NULL COMMENT '标题',
  `layout_type` varchar(20) NOT NULL COMMENT '布局类型: two_col/three_col/big/list/mixed',
  `content_strategy` varchar(20) NOT NULL COMMENT '内容策略: algorithm/manual/time/hot/random',
  `content_ids` json DEFAULT NULL COMMENT '内容ID列表(人工推荐)',
  `filter_rule` json DEFAULT NULL COMMENT '筛选规则',
  `sort` int NOT NULL DEFAULT '0' COMMENT '排序',
  `status` tinyint NOT NULL DEFAULT '1' COMMENT '状态: 0-禁用 1-启用',
  `description` varchar(200) DEFAULT NULL COMMENT '描述',
  PRIMARY KEY (`id`),
  KEY `idx_feed_config_deleted_at` (`deleted_at`),
  KEY `idx_feed_config_tenant_id` (`tenant_id`),
  KEY `idx_feed_config_channel_id` (`channel_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='Feed流配置表';

-- 插入测试数据

-- 金刚位测试数据 (推荐频道)
INSERT INTO `diamond` (`tenant_id`, `channel_id`, `group_id`, `sort`, `title`, `icon`, `link_type`, `link_value`, `status`, `description`) VALUES
(1, 1, 1, 1, '热门话题', '🔥', 'topic', '1', 1, '热门话题入口'),
(1, 1, 1, 2, '新人推荐', '⭐', 'channel', '2', 1, '新人推荐频道'),
(1, 1, 1, 3, '排行榜', '📊', 'external', '/rank', 1, '排行榜页面'),
(1, 1, 1, 4, '活动中心', '🎁', 'external', '/activity', 1, '活动中心页面'),
(1, 1, 2, 1, '每日签到', '✅', 'external', '/checkin', 1, '每日签到'),
(1, 1, 2, 2, '会员中心', '👑', 'external', '/vip', 1, '会员中心');

-- 推荐位测试数据
INSERT INTO `recommend` (`tenant_id`, `channel_id`, `title`, `display_type`, `source_type`, `content_ids`, `filter_rule`, `sort`, `status`, `description`) VALUES
(1, 1, '今日推荐', 'single', 'algorithm', NULL, '{"type": "video", "limit": 1}', 1, 1, '今日推荐内容'),
(1, 1, '热门精选', 'scroll', 'filter', NULL, '{"type": "video", "order": "hot", "limit": 10}', 2, 1, '热门精选内容'),
(1, 1, '编辑推荐', 'grid', 'manual', '[1, 2, 3, 4, 5, 6]', NULL, 3, 1, '编辑推荐内容');

-- Feed流配置测试数据
INSERT INTO `feed_config` (`tenant_id`, `channel_id`, `title`, `layout_type`, `content_strategy`, `content_ids`, `filter_rule`, `sort`, `status`, `description`) VALUES
(1, 1, '推荐Feed', 'two_col', 'algorithm', NULL, '{"type": "all", "limit": 20}', 1, 1, '推荐Feed流'),
(1, 2, '搞笑Feed', 'waterfall', 'hot', NULL, '{"type": "video", "category": "funny", "limit": 20}', 1, 1, '搞笑Feed流');
