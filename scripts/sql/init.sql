-- 创建数据库
CREATE DATABASE IF NOT EXISTS `happy` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

USE `happy`;

-- 租户表
CREATE TABLE IF NOT EXISTS `tenant` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` varchar(100) NOT NULL COMMENT '站点名称',
  `code` varchar(50) NOT NULL COMMENT '站点代码',
  `domain` varchar(255) NOT NULL COMMENT '站点域名',
  `description` text COMMENT '站点描述',
  `logo` varchar(255) DEFAULT NULL COMMENT '站点Logo',
  `status` tinyint DEFAULT '1' COMMENT '状态: 0-禁用 1-启用',
  `config` json DEFAULT NULL COMMENT '站点配置JSON',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_name` (`name`),
  UNIQUE KEY `idx_code` (`code`),
  UNIQUE KEY `idx_domain` (`domain`),
  KEY `idx_deleted_at` (`deleted_at`),
  KEY `idx_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='租户/站点表';

-- 用户表
CREATE TABLE IF NOT EXISTS `user` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `tenant_id` bigint unsigned NOT NULL COMMENT '租户ID',
  `username` varchar(50) DEFAULT NULL COMMENT '用户名',
  `email` varchar(100) DEFAULT NULL COMMENT '邮箱',
  `phone` varchar(20) DEFAULT NULL COMMENT '手机号',
  `password` varchar(255) DEFAULT NULL COMMENT '密码',
  `nickname` varchar(50) DEFAULT NULL COMMENT '昵称',
  `avatar` varchar(255) DEFAULT NULL COMMENT '头像',
  `gender` tinyint DEFAULT '0' COMMENT '性别: 0-未知 1-男 2-女',
  `birthday` date DEFAULT NULL COMMENT '生日',
  `bio` text COMMENT '个人简介',
  `status` tinyint DEFAULT '1' COMMENT '状态: 0-禁用 1-启用',
  `role` tinyint DEFAULT '0' COMMENT '角色: 0-普通用户 1-博主 2-管理员',
  `follow_count` int DEFAULT '0' COMMENT '关注数',
  `fans_count` int DEFAULT '0' COMMENT '粉丝数',
  `like_count` int DEFAULT '0' COMMENT '获赞数',
  `third_party_id` varchar(100) DEFAULT NULL COMMENT '第三方登录ID',
  `third_type` varchar(20) DEFAULT NULL COMMENT '第三方登录类型',
  PRIMARY KEY (`id`),
  KEY `idx_deleted_at` (`deleted_at`),
  KEY `idx_tenant_id` (`tenant_id`),
  KEY `idx_username` (`username`),
  KEY `idx_email` (`email`),
  KEY `idx_phone` (`phone`),
  KEY `idx_status` (`status`),
  KEY `idx_role` (`role`),
  KEY `idx_third_party` (`third_party_id`, `third_type`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户表';

-- 内容表
CREATE TABLE IF NOT EXISTS `content` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `tenant_id` bigint unsigned NOT NULL COMMENT '租户ID',
  `user_id` bigint unsigned NOT NULL COMMENT '创建者ID',
  `title` varchar(255) DEFAULT NULL COMMENT '标题',
  `description` text COMMENT '描述',
  `cover` varchar(255) DEFAULT NULL COMMENT '封面',
  `type` tinyint NOT NULL COMMENT '类型: 1-长视频 2-短视频 3-短剧 4-漫剧 5-小说 6-图文',
  `status` tinyint DEFAULT '0' COMMENT '状态: 0-草稿 1-已发布 2-已下架 3-审核中',
  `view_count` int DEFAULT '0' COMMENT '浏览数',
  `like_count` int DEFAULT '0' COMMENT '点赞数',
  `comment_count` int DEFAULT '0' COMMENT '评论数',
  `collect_count` int DEFAULT '0' COMMENT '收藏数',
  `share_count` int DEFAULT '0' COMMENT '分享数',
  `media` json DEFAULT NULL COMMENT '媒体资源JSON',
  `extra` json DEFAULT NULL COMMENT '扩展信息JSON',
  `publish_at` datetime(3) DEFAULT NULL COMMENT '发布时间',
  PRIMARY KEY (`id`),
  KEY `idx_deleted_at` (`deleted_at`),
  KEY `idx_tenant_id` (`tenant_id`),
  KEY `idx_user_id` (`user_id`),
  KEY `idx_title` (`title`),
  KEY `idx_type` (`type`),
  KEY `idx_status` (`status`),
  KEY `idx_publish_at` (`publish_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='内容表';

-- 内容媒体资源表
CREATE TABLE IF NOT EXISTS `content_media` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `content_id` bigint unsigned NOT NULL COMMENT '内容ID',
  `type` tinyint NOT NULL COMMENT '类型: 1-图片 2-视频 3-音频',
  `url` varchar(255) NOT NULL COMMENT '资源URL',
  `thumbnail` varchar(255) DEFAULT NULL COMMENT '缩略图',
  `duration` int DEFAULT '0' COMMENT '时长(秒)',
  `width` int DEFAULT '0' COMMENT '宽度',
  `height` int DEFAULT '0' COMMENT '高度',
  `size` bigint DEFAULT '0' COMMENT '文件大小',
  `sort` int DEFAULT '0' COMMENT '排序',
  PRIMARY KEY (`id`),
  KEY `idx_content_id` (`content_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='内容媒体资源表';

-- 话题表
CREATE TABLE IF NOT EXISTS `topic` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `tenant_id` bigint unsigned NOT NULL COMMENT '租户ID',
  `name` varchar(100) NOT NULL COMMENT '话题名称',
  `description` text COMMENT '话题描述',
  `cover` varchar(255) DEFAULT NULL COMMENT '话题封面',
  `status` tinyint DEFAULT '1' COMMENT '状态: 0-禁用 1-启用',
  `sort` int DEFAULT '0' COMMENT '排序',
  PRIMARY KEY (`id`),
  KEY `idx_deleted_at` (`deleted_at`),
  KEY `idx_tenant_id` (`tenant_id`),
  KEY `idx_name` (`name`),
  KEY `idx_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='话题表';

-- 标签表
CREATE TABLE IF NOT EXISTS `tag` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `tenant_id` bigint unsigned NOT NULL COMMENT '租户ID',
  `name` varchar(50) NOT NULL COMMENT '标签名称',
  `type` tinyint DEFAULT '0' COMMENT '类型: 0-通用 1-长视频 2-短视频 3-短剧 4-漫剧 5-小说 6-图文',
  `status` tinyint DEFAULT '1' COMMENT '状态: 0-禁用 1-启用',
  `sort` int DEFAULT '0' COMMENT '排序',
  PRIMARY KEY (`id`),
  KEY `idx_deleted_at` (`deleted_at`),
  KEY `idx_tenant_id` (`tenant_id`),
  KEY `idx_name` (`name`),
  KEY `idx_type` (`type`),
  KEY `idx_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='标签表';

-- 内容话题关联表
CREATE TABLE IF NOT EXISTS `content_topic` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `content_id` bigint unsigned NOT NULL COMMENT '内容ID',
  `topic_id` bigint unsigned NOT NULL COMMENT '话题ID',
  PRIMARY KEY (`id`),
  KEY `idx_content_id` (`content_id`),
  KEY `idx_topic_id` (`topic_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='内容话题关联表';

-- 内容标签关联表
CREATE TABLE IF NOT EXISTS `content_tag` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `content_id` bigint unsigned NOT NULL COMMENT '内容ID',
  `tag_id` bigint unsigned NOT NULL COMMENT '标签ID',
  PRIMARY KEY (`id`),
  KEY `idx_content_id` (`content_id`),
  KEY `idx_tag_id` (`tag_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='内容标签关联表';

-- 频道表
CREATE TABLE IF NOT EXISTS `channel` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `tenant_id` bigint unsigned NOT NULL COMMENT '租户ID',
  `name` varchar(50) NOT NULL COMMENT '频道名称',
  `code` varchar(50) NOT NULL COMMENT '频道代码',
  `description` text COMMENT '频道描述',
  `icon` varchar(255) DEFAULT NULL COMMENT '频道图标',
  `status` tinyint DEFAULT '1' COMMENT '状态: 0-禁用 1-启用',
  `sort` int DEFAULT '0' COMMENT '排序',
  PRIMARY KEY (`id`),
  KEY `idx_deleted_at` (`deleted_at`),
  KEY `idx_tenant_id` (`tenant_id`),
  KEY `idx_name` (`name`),
  KEY `idx_code` (`code`),
  KEY `idx_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='频道表';

-- Banner表
CREATE TABLE IF NOT EXISTS `banner` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `tenant_id` bigint unsigned NOT NULL COMMENT '租户ID',
  `channel_id` bigint unsigned DEFAULT NULL COMMENT '频道ID',
  `title` varchar(100) DEFAULT NULL COMMENT '标题',
  `image` varchar(255) NOT NULL COMMENT '图片URL',
  `link_type` tinyint DEFAULT '1' COMMENT '链接类型: 1-内容 2-外链',
  `link_url` varchar(255) DEFAULT NULL COMMENT '链接地址',
  `content_id` bigint unsigned DEFAULT NULL COMMENT '内容ID',
  `status` tinyint DEFAULT '1' COMMENT '状态: 0-禁用 1-启用',
  `sort` int DEFAULT '0' COMMENT '排序',
  `start_time` datetime(3) DEFAULT NULL COMMENT '开始时间',
  `end_time` datetime(3) DEFAULT NULL COMMENT '结束时间',
  PRIMARY KEY (`id`),
  KEY `idx_deleted_at` (`deleted_at`),
  KEY `idx_tenant_id` (`tenant_id`),
  KEY `idx_channel_id` (`channel_id`),
  KEY `idx_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='Banner表';

-- 金刚位表
CREATE TABLE IF NOT EXISTS `diamond_position` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `tenant_id` bigint unsigned NOT NULL COMMENT '租户ID',
  `channel_id` bigint unsigned DEFAULT NULL COMMENT '频道ID',
  `name` varchar(50) NOT NULL COMMENT '名称',
  `icon` varchar(255) NOT NULL COMMENT '图标',
  `link_type` tinyint DEFAULT '1' COMMENT '链接类型: 1-频道 2-话题 3-外链',
  `link_url` varchar(255) DEFAULT NULL COMMENT '链接地址',
  `channel_link` bigint unsigned DEFAULT NULL COMMENT '频道链接ID',
  `topic_link` bigint unsigned DEFAULT NULL COMMENT '话题链接ID',
  `status` tinyint DEFAULT '1' COMMENT '状态: 0-禁用 1-启用',
  `sort` int DEFAULT '0' COMMENT '排序',
  PRIMARY KEY (`id`),
  KEY `idx_deleted_at` (`deleted_at`),
  KEY `idx_tenant_id` (`tenant_id`),
  KEY `idx_channel_id` (`channel_id`),
  KEY `idx_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='金刚位表';

-- Feed流配置表
CREATE TABLE IF NOT EXISTS `feed_config` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `tenant_id` bigint unsigned NOT NULL COMMENT '租户ID',
  `channel_id` bigint unsigned NOT NULL COMMENT '频道ID',
  `name` varchar(50) NOT NULL COMMENT '配置名称',
  `layout` tinyint DEFAULT '1' COMMENT '布局: 1-一行两列 2-一行三列',
  `strategy` tinyint DEFAULT '1' COMMENT '策略: 1-算法推荐 2-人工推荐 3-随机',
  `status` tinyint DEFAULT '1' COMMENT '状态: 0-禁用 1-启用',
  `sort` int DEFAULT '0' COMMENT '排序',
  PRIMARY KEY (`id`),
  KEY `idx_deleted_at` (`deleted_at`),
  KEY `idx_tenant_id` (`tenant_id`),
  KEY `idx_channel_id` (`channel_id`),
  KEY `idx_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='Feed流配置表';

-- 广告位表
CREATE TABLE IF NOT EXISTS `ad_position` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `tenant_id` bigint unsigned NOT NULL COMMENT '租户ID',
  `channel_id` bigint unsigned DEFAULT NULL COMMENT '频道ID',
  `name` varchar(50) NOT NULL COMMENT '广告位名称',
  `code` varchar(50) NOT NULL COMMENT '广告位代码',
  `type` tinyint DEFAULT '1' COMMENT '类型: 1-图片 2-视频',
  `image_url` varchar(255) DEFAULT NULL COMMENT '图片URL',
  `video_url` varchar(255) DEFAULT NULL COMMENT '视频URL',
  `link_url` varchar(255) DEFAULT NULL COMMENT '跳转链接',
  `status` tinyint DEFAULT '1' COMMENT '状态: 0-禁用 1-启用',
  `sort` int DEFAULT '0' COMMENT '排序',
  `start_time` datetime(3) DEFAULT NULL COMMENT '开始时间',
  `end_time` datetime(3) DEFAULT NULL COMMENT '结束时间',
  PRIMARY KEY (`id`),
  KEY `idx_deleted_at` (`deleted_at`),
  KEY `idx_tenant_id` (`tenant_id`),
  KEY `idx_channel_id` (`channel_id`),
  KEY `idx_code` (`code`),
  KEY `idx_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='广告位表';

-- 互动表
CREATE TABLE IF NOT EXISTS `interaction` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `user_id` bigint unsigned NOT NULL COMMENT '用户ID',
  `target_id` bigint unsigned NOT NULL COMMENT '目标ID',
  `type` tinyint NOT NULL COMMENT '类型: 1-点赞 2-收藏 3-分享',
  `target_type` tinyint NOT NULL COMMENT '目标类型: 1-内容 2-评论',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_unique` (`user_id`, `target_id`, `type`, `target_type`),
  KEY `idx_user_id` (`user_id`),
  KEY `idx_target_id` (`target_id`),
  KEY `idx_type` (`type`),
  KEY `idx_target_type` (`target_type`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='互动表';

-- 评论表
CREATE TABLE IF NOT EXISTS `comment` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `content_id` bigint unsigned NOT NULL COMMENT '内容ID',
  `user_id` bigint unsigned NOT NULL COMMENT '用户ID',
  `parent_id` bigint unsigned DEFAULT '0' COMMENT '父评论ID',
  `reply_user_id` bigint unsigned DEFAULT NULL COMMENT '回复用户ID',
  `content` text NOT NULL COMMENT '评论内容',
  `like_count` int DEFAULT '0' COMMENT '点赞数',
  `status` tinyint DEFAULT '1' COMMENT '状态: 0-删除 1-正常',
  PRIMARY KEY (`id`),
  KEY `idx_deleted_at` (`deleted_at`),
  KEY `idx_content_id` (`content_id`),
  KEY `idx_user_id` (`user_id`),
  KEY `idx_parent_id` (`parent_id`),
  KEY `idx_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='评论表';

-- 关注表
CREATE TABLE IF NOT EXISTS `follow` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `follower_id` bigint unsigned NOT NULL COMMENT '粉丝ID',
  `following_id` bigint unsigned NOT NULL COMMENT '关注者ID',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_unique` (`follower_id`, `following_id`),
  KEY `idx_follower_id` (`follower_id`),
  KEY `idx_following_id` (`following_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='关注表';

-- 观看历史表
CREATE TABLE IF NOT EXISTS `view_history` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `user_id` bigint unsigned NOT NULL COMMENT '用户ID',
  `content_id` bigint unsigned NOT NULL COMMENT '内容ID',
  `duration` int DEFAULT '0' COMMENT '观看时长',
  `progress` int DEFAULT '0' COMMENT '观看进度',
  PRIMARY KEY (`id`),
  KEY `idx_user_id` (`user_id`),
  KEY `idx_content_id` (`content_id`),
  KEY `idx_created_at` (`created_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='观看历史表';

-- 消息通知表
CREATE TABLE IF NOT EXISTS `message` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `user_id` bigint unsigned NOT NULL COMMENT '接收用户ID',
  `from_user_id` bigint unsigned DEFAULT NULL COMMENT '发送用户ID',
  `type` tinyint NOT NULL COMMENT '类型: 1-点赞 2-评论 3-收藏 4-关注 5-系统',
  `title` varchar(100) DEFAULT NULL COMMENT '标题',
  `content` text COMMENT '内容',
  `target_id` bigint unsigned DEFAULT NULL COMMENT '目标ID',
  `target_type` tinyint DEFAULT NULL COMMENT '目标类型',
  `is_read` tinyint DEFAULT '0' COMMENT '是否已读: 0-未读 1-已读',
  PRIMARY KEY (`id`),
  KEY `idx_deleted_at` (`deleted_at`),
  KEY `idx_user_id` (`user_id`),
  KEY `idx_type` (`type`),
  KEY `idx_is_read` (`is_read`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='消息通知表';

-- 管理员用户表
CREATE TABLE IF NOT EXISTS `admin_user` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `username` varchar(50) NOT NULL COMMENT '用户名',
  `password` varchar(255) NOT NULL COMMENT '密码',
  `nickname` varchar(50) DEFAULT NULL COMMENT '昵称',
  `avatar` varchar(255) DEFAULT NULL COMMENT '头像',
  `email` varchar(100) DEFAULT NULL COMMENT '邮箱',
  `phone` varchar(20) DEFAULT NULL COMMENT '手机号',
  `status` tinyint DEFAULT '1' COMMENT '状态: 0-禁用 1-启用',
  `role_id` bigint unsigned DEFAULT NULL COMMENT '角色ID',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_username` (`username`),
  KEY `idx_deleted_at` (`deleted_at`),
  KEY `idx_status` (`status`),
  KEY `idx_role_id` (`role_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='管理员用户表';

-- 角色表
CREATE TABLE IF NOT EXISTS `role` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` varchar(50) NOT NULL COMMENT '角色名称',
  `code` varchar(50) NOT NULL COMMENT '角色代码',
  `description` text COMMENT '描述',
  `status` tinyint DEFAULT '1' COMMENT '状态: 0-禁用 1-启用',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_code` (`code`),
  KEY `idx_deleted_at` (`deleted_at`),
  KEY `idx_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='角色表';

-- 权限表
CREATE TABLE IF NOT EXISTS `permission` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` varchar(50) NOT NULL COMMENT '权限名称',
  `code` varchar(100) NOT NULL COMMENT '权限代码',
  `type` tinyint NOT NULL COMMENT '类型: 1-菜单 2-按钮',
  `parent_id` bigint unsigned DEFAULT '0' COMMENT '父权限ID',
  `path` varchar(255) DEFAULT NULL COMMENT '路由路径',
  `icon` varchar(50) DEFAULT NULL COMMENT '图标',
  `sort` int DEFAULT '0' COMMENT '排序',
  `status` tinyint DEFAULT '1' COMMENT '状态: 0-禁用 1-启用',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_code` (`code`),
  KEY `idx_deleted_at` (`deleted_at`),
  KEY `idx_parent_id` (`parent_id`),
  KEY `idx_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='权限表';

-- 角色权限关联表
CREATE TABLE IF NOT EXISTS `role_permission` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `role_id` bigint unsigned NOT NULL COMMENT '角色ID',
  `permission_id` bigint unsigned NOT NULL COMMENT '权限ID',
  PRIMARY KEY (`id`),
  KEY `idx_role_id` (`role_id`),
  KEY `idx_permission_id` (`permission_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='角色权限关联表';

-- 插入默认租户
INSERT INTO `tenant` (`name`, `code`, `domain`, `description`, `status`) VALUES
('默认站点', 'default', 'localhost', '默认站点', 1);

-- 插入默认频道
INSERT INTO `channel` (`tenant_id`, `name`, `code`, `description`, `status`, `sort`) VALUES
(1, '推荐', 'recommend', '推荐内容', 1, 1),
(1, '搞笑', 'funny', '搞笑内容', 1, 2),
(1, '热门', 'hot', '热门内容', 1, 3),
(1, '颜值', 'beauty', '颜值内容', 1, 4),
(1, '动漫', 'anime', '动漫内容', 1, 5),
(1, '社区', 'community', '社区内容', 1, 6);

-- 插入默认管理员角色
INSERT INTO `role` (`name`, `code`, `description`, `status`) VALUES
('超级管理员', 'super_admin', '拥有所有权限', 1),
('管理员', 'admin', '普通管理员', 1),
('运营', 'operator', '运营人员', 1);

-- 插入默认管理员 (密码: admin123)
INSERT INTO `admin_user` (`username`, `password`, `nickname`, `status`, `role_id`) VALUES
('admin', '$2a$10$N.zmdr9k7uOCQb376NoUnuTJ8iAt6Z5EHsM8lE9lBOsl7iAt6Z5EH', '超级管理员', 1, 1);
