-- 角色表
CREATE TABLE IF NOT EXISTS `role` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `tenant_id` bigint unsigned NOT NULL DEFAULT 1 COMMENT '租户ID',
  `name` varchar(50) NOT NULL COMMENT '角色名称',
  `code` varchar(50) NOT NULL COMMENT '角色编码',
  `description` varchar(200) DEFAULT '' COMMENT '角色描述',
  `status` tinyint NOT NULL DEFAULT 1 COMMENT '状态: 0-禁用 1-启用',
  `sort` int NOT NULL DEFAULT 0 COMMENT '排序',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_tenant_code` (`tenant_id`, `code`),
  KEY `idx_tenant_id` (`tenant_id`),
  KEY `idx_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='角色表';

-- 权限表
CREATE TABLE IF NOT EXISTS `permission` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `tenant_id` bigint unsigned NOT NULL DEFAULT 1 COMMENT '租户ID',
  `name` varchar(50) NOT NULL COMMENT '权限名称',
  `code` varchar(100) NOT NULL COMMENT '权限编码',
  `type` varchar(20) NOT NULL COMMENT '权限类型: menu-菜单 button-按钮 api-接口',
  `parent_id` bigint unsigned DEFAULT 0 COMMENT '父权限ID',
  `path` varchar(200) DEFAULT '' COMMENT '路由路径',
  `method` varchar(10) DEFAULT '' COMMENT 'HTTP方法',
  `icon` varchar(50) DEFAULT '' COMMENT '图标',
  `sort` int NOT NULL DEFAULT 0 COMMENT '排序',
  `status` tinyint NOT NULL DEFAULT 1 COMMENT '状态: 0-禁用 1-启用',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_tenant_code` (`tenant_id`, `code`),
  KEY `idx_tenant_id` (`tenant_id`),
  KEY `idx_parent_id` (`parent_id`),
  KEY `idx_type` (`type`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='权限表';

-- 角色权限关联表
CREATE TABLE IF NOT EXISTS `role_permission` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `tenant_id` bigint unsigned NOT NULL DEFAULT 1 COMMENT '租户ID',
  `role_id` bigint unsigned NOT NULL COMMENT '角色ID',
  `permission_id` bigint unsigned NOT NULL COMMENT '权限ID',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_role_permission` (`role_id`, `permission_id`),
  KEY `idx_tenant_id` (`tenant_id`),
  KEY `idx_role_id` (`role_id`),
  KEY `idx_permission_id` (`permission_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='角色权限关联表';

-- 管理员表
CREATE TABLE IF NOT EXISTS `admin_user` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `tenant_id` bigint unsigned NOT NULL DEFAULT 1 COMMENT '租户ID',
  `username` varchar(50) NOT NULL COMMENT '用户名',
  `password` varchar(255) NOT NULL COMMENT '密码',
  `name` varchar(50) NOT NULL COMMENT '姓名',
  `email` varchar(100) DEFAULT '' COMMENT '邮箱',
  `phone` varchar(20) DEFAULT '' COMMENT '手机号',
  `avatar` varchar(500) DEFAULT '' COMMENT '头像',
  `status` tinyint NOT NULL DEFAULT 1 COMMENT '状态: 0-禁用 1-启用',
  `last_login_at` timestamp NULL DEFAULT NULL COMMENT '最后登录时间',
  `last_login_ip` varchar(50) DEFAULT '' COMMENT '最后登录IP',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_tenant_username` (`tenant_id`, `username`),
  KEY `idx_tenant_id` (`tenant_id`),
  KEY `idx_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='管理员表';

-- 管理员角色关联表
CREATE TABLE IF NOT EXISTS `admin_role` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `tenant_id` bigint unsigned NOT NULL DEFAULT 1 COMMENT '租户ID',
  `admin_id` bigint unsigned NOT NULL COMMENT '管理员ID',
  `role_id` bigint unsigned NOT NULL COMMENT '角色ID',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_admin_role` (`admin_id`, `role_id`),
  KEY `idx_tenant_id` (`tenant_id`),
  KEY `idx_admin_id` (`admin_id`),
  KEY `idx_role_id` (`role_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='管理员角色关联表';

-- 插入默认角色
INSERT INTO `role` (`tenant_id`, `name`, `code`, `description`, `status`, `sort`) VALUES
(1, '超级管理员', 'super_admin', '拥有所有权限', 1, 1),
(1, '管理员', 'admin', '拥有大部分权限', 1, 2),
(1, '运营', 'operator', '运营相关权限', 1, 3),
(1, '编辑', 'editor', '内容编辑权限', 1, 4);

-- 插入默认权限
INSERT INTO `permission` (`tenant_id`, `name`, `code`, `type`, `parent_id`, `path`, `icon`, `sort`, `status`) VALUES
-- 一级菜单
(1, '仪表盘', 'dashboard', 'menu', 0, '/dashboard', 'DashboardOutlined', 1, 1),
(1, '用户管理', 'user', 'menu', 0, '/user', 'UserOutlined', 2, 1),
(1, '物料管理', 'material', 'menu', 0, '/material', 'PictureOutlined', 3, 1),
(1, '内容管理', 'content', 'menu', 0, '/content', 'FileOutlined', 4, 1),
(1, '频道管理', 'channel', 'menu', 0, '/channel', 'AppstoreOutlined', 5, 1),
(1, '系统管理', 'system', 'menu', 0, '/system', 'SettingOutlined', 6, 1),
-- 频道管理子菜单
(1, '频道列表', 'channel:list', 'menu', 5, '/channel/list', '', 1, 1),
(1, '金刚位管理', 'channel:diamond', 'menu', 5, '/channel/diamond', '', 2, 1),
(1, '广告位管理', 'channel:ad-slot', 'menu', 5, '/channel/ad-slot', '', 3, 1),
-- 系统管理子菜单
(1, '角色管理', 'system:role', 'menu', 6, '/system/role', '', 1, 1),
(1, '权限管理', 'system:permission', 'menu', 6, '/system/permission', '', 2, 1),
(1, '管理员管理', 'system:admin', 'menu', 6, '/system/admin-user', '', 3, 1);

-- 插入按钮权限
INSERT INTO `permission` (`tenant_id`, `name`, `code`, `type`, `parent_id`, `sort`, `status`) VALUES
-- 用户管理按钮
(1, '查看用户', 'user:view', 'button', 2, 1, 1),
(1, '创建用户', 'user:create', 'button', 2, 2, 1),
(1, '编辑用户', 'user:edit', 'button', 2, 3, 1),
(1, '删除用户', 'user:delete', 'button', 2, 4, 1),
-- 物料管理按钮
(1, '查看物料', 'material:view', 'button', 3, 1, 1),
(1, '创建物料', 'material:create', 'button', 3, 2, 1),
(1, '编辑物料', 'material:edit', 'button', 3, 3, 1),
(1, '删除物料', 'material:delete', 'button', 3, 4, 1);

-- 为超级管理员分配所有权限
INSERT INTO `role_permission` (`tenant_id`, `role_id`, `permission_id`)
SELECT 1, 1, id FROM permission WHERE tenant_id = 1;

-- 插入默认管理员
INSERT INTO `admin_user` (`tenant_id`, `username`, `password`, `name`, `email`, `phone`, `status`) VALUES
(1, 'admin', '$2a$10$N.zmdr9k7uOCQb376NoUnuTJ8iAt6Z5EHsM8lE9lBOsl7iAt6Z5EH', '超级管理员', 'admin@example.com', '13800138000', 1);

-- 为管理员分配超级管理员角色
INSERT INTO `admin_role` (`tenant_id`, `admin_id`, `role_id`) VALUES (1, 1, 1);

-- 查询验证
SELECT * FROM role WHERE tenant_id = 1;
SELECT * FROM permission WHERE tenant_id = 1 AND type = 'menu';
SELECT * FROM admin_user WHERE tenant_id = 1;
