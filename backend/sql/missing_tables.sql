-- 发现页模块配置表
CREATE TABLE IF NOT EXISTS `discover_module` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `tenant_id` bigint unsigned DEFAULT '1',
  `name` varchar(100) NOT NULL COMMENT '模块名称',
  `type` varchar(50) NOT NULL COMMENT '模块类型: banner/diamond/recommend/feed',
  `status` tinyint DEFAULT '1' COMMENT '状态: 0-禁用 1-启用',
  `sort` int DEFAULT '0' COMMENT '排序',
  PRIMARY KEY (`id`),
  KEY `idx_tenant_id` (`tenant_id`),
  KEY `idx_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='发现页模块配置表';

-- 发现页内容表
CREATE TABLE IF NOT EXISTS `discover_content` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `tenant_id` bigint unsigned DEFAULT '1',
  `module_id` bigint unsigned NOT NULL COMMENT '模块ID',
  `content_type` varchar(50) NOT NULL COMMENT '内容类型: material/banner/diamond',
  `content_id` bigint unsigned NOT NULL COMMENT '内容ID',
  `sort` int DEFAULT '0' COMMENT '排序',
  `status` tinyint DEFAULT '1' COMMENT '状态: 0-禁用 1-启用',
  PRIMARY KEY (`id`),
  KEY `idx_tenant_id` (`tenant_id`),
  KEY `idx_module_id` (`module_id`),
  KEY `idx_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='发现页内容表';

-- 用户表
CREATE TABLE IF NOT EXISTS `user` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `tenant_id` bigint unsigned DEFAULT '1',
  `username` varchar(50) NOT NULL COMMENT '用户名',
  `email` varchar(100) DEFAULT NULL COMMENT '邮箱',
  `phone` varchar(20) DEFAULT NULL COMMENT '手机号',
  `password` varchar(255) NOT NULL COMMENT '密码',
  `nickname` varchar(50) DEFAULT NULL COMMENT '昵称',
  `avatar` varchar(500) DEFAULT NULL COMMENT '头像',
  `gender` tinyint DEFAULT '0' COMMENT '性别: 0-未知 1-男 2-女',
  `birthday` date DEFAULT NULL COMMENT '生日',
  `bio` varchar(500) DEFAULT NULL COMMENT '个人简介',
  `status` tinyint DEFAULT '1' COMMENT '状态: 0-禁用 1-启用',
  `role` tinyint DEFAULT '0' COMMENT '角色: 0-普通用户 1-创作者',
  `follow_count` int DEFAULT '0' COMMENT '关注数',
  `fans_count` int DEFAULT '0' COMMENT '粉丝数',
  `like_count` int DEFAULT '0' COMMENT '获赞数',
  `third_party_id` varchar(100) DEFAULT NULL COMMENT '第三方平台ID',
  `third_type` varchar(50) DEFAULT NULL COMMENT '第三方平台类型',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_username` (`username`),
  KEY `idx_tenant_id` (`tenant_id`),
  KEY `idx_email` (`email`),
  KEY `idx_phone` (`phone`),
  KEY `idx_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户表';

-- 互动表(点赞、收藏、分享)
CREATE TABLE IF NOT EXISTS `interaction` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `tenant_id` bigint unsigned DEFAULT '1',
  `user_id` bigint unsigned NOT NULL COMMENT '用户ID',
  `target_id` bigint unsigned NOT NULL COMMENT '目标ID',
  `target_type` varchar(50) NOT NULL COMMENT '目标类型: material/comment/user',
  `type` varchar(20) NOT NULL COMMENT '互动类型: like/collect/share',
  `status` tinyint DEFAULT '1' COMMENT '状态: 0-取消 1-有效',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_user_target` (`user_id`, `target_id`, `target_type`, `type`),
  KEY `idx_tenant_id` (`tenant_id`),
  KEY `idx_target` (`target_id`, `target_type`),
  KEY `idx_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='互动表';

-- 搜索索引表
CREATE TABLE IF NOT EXISTS `search_index` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `tenant_id` bigint unsigned DEFAULT '1',
  `content_type` varchar(50) NOT NULL COMMENT '内容类型: material/user/topic',
  `content_id` bigint unsigned NOT NULL COMMENT '内容ID',
  `title` varchar(500) DEFAULT NULL COMMENT '标题',
  `content` text COMMENT '内容',
  `tags` varchar(500) DEFAULT NULL COMMENT '标签(JSON数组)',
  `status` tinyint DEFAULT '1' COMMENT '状态: 0-禁用 1-启用',
  `score` int DEFAULT '0' COMMENT '热度分数',
  PRIMARY KEY (`id`),
  KEY `idx_tenant_id` (`tenant_id`),
  KEY `idx_content` (`content_type`, `content_id`),
  KEY `idx_deleted_at` (`deleted_at`),
  FULLTEXT KEY `ft_title_content` (`title`, `content`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='搜索索引表';

-- 插入发现页模块初始数据
INSERT INTO `discover_module` (`name`, `type`, `status`, `sort`) VALUES
('Banner轮播', 'banner', 1, 1),
('金刚位', 'diamond', 1, 2),
('推荐位', 'recommend', 1, 3),
('信息流', 'feed', 1, 4);
