# 数据库设计文档

## 1. 数据库基本信息

- **数据库名**: happy
- **字符集**: utf8mb4
- **排序规则**: utf8mb4_unicode_ci
- **存储引擎**: InnoDB

---

## 2. 表结构设计

### 2.1 系统管理模块

#### admin_user (管理员表)

| 字段 | 类型 | 约束 | 说明 |
|-----|------|------|------|
| id | bigint unsigned | PK, AUTO_INCREMENT | 主键 |
| created_at | datetime(3) | | 创建时间 |
| updated_at | datetime(3) | | 更新时间 |
| deleted_at | datetime(3) | INDEX | 删除时间(软删除) |
| username | varchar(50) | UNIQUE, NOT NULL | 用户名 |
| password | varchar(255) | NOT NULL | 密码(bcrypt加密) |
| nickname | varchar(50) | | 昵称 (映射为realname) |
| avatar | varchar(500) | | 头像URL |
| email | varchar(100) | | 邮箱 |
| phone | varchar(20) | | 手机号 |
| status | tinyint | DEFAULT 1 | 状态: 0-禁用 1-启用 |
| role_id | bigint unsigned | | 角色ID(已废弃,使用admin_user_roles) |

**索引**:
- PRIMARY KEY (id)
- UNIQUE KEY idx_username (username)
- KEY idx_deleted_at (deleted_at)

**字段映射**: `nickname` → `realname` (代码中使用realname)

---

#### role (角色表)

| 字段 | 类型 | 约束 | 说明 |
|-----|------|------|------|
| id | bigint unsigned | PK, AUTO_INCREMENT | 主键 |
| created_at | datetime(3) | | 创建时间 |
| updated_at | datetime(3) | | 更新时间 |
| deleted_at | datetime(3) | INDEX | 删除时间 |
| name | varchar(50) | NOT NULL | 角色名称 |
| code | varchar(50) | UNIQUE, NOT NULL | 角色代码 |
| description | varchar(200) | | 描述 |
| status | tinyint | DEFAULT 1 | 状态: 0-禁用 1-启用 |

**初始数据**:
- 超级管理员 (super_admin)
- 管理员 (admin)
- 运营 (operator)

---

#### permission (权限表)

| 字段 | 类型 | 约束 | 说明 |
|-----|------|------|------|
| id | bigint unsigned | PK, AUTO_INCREMENT | 主键 |
| created_at | datetime(3) | | 创建时间 |
| updated_at | datetime(3) | | 更新时间 |
| deleted_at | datetime(3) | INDEX | 删除时间 |
| name | varchar(50) | NOT NULL | 权限名称 |
| code | varchar(100) | UNIQUE, NOT NULL | 权限代码 |
| type | tinyint | NOT NULL | 类型: 1-菜单 2-按钮 |
| parent_id | bigint unsigned | DEFAULT 0 | 父权限ID |
| path | varchar(200) | | 路由路径 |
| icon | varchar(50) | | 图标 |
| status | tinyint | DEFAULT 1 | 状态 |
| sort | int | DEFAULT 0 | 排序 |

**树形结构**: 支持父子关系,用于菜单和权限树

---

#### admin_user_roles (管理员角色关联表)

| 字段 | 类型 | 约束 | 说明 |
|-----|------|------|------|
| id | bigint unsigned | PK, AUTO_INCREMENT | 主键 |
| created_at | datetime(3) | | 创建时间 |
| admin_user_id | bigint unsigned | NOT NULL | 管理员ID |
| role_id | bigint unsigned | NOT NULL | 角色ID |

**索引**:
- UNIQUE KEY uk_admin_role (admin_user_id, role_id)
- KEY idx_admin_user_id (admin_user_id)
- KEY idx_role_id (role_id)

---

#### role_permission (角色权限关联表)

| 字段 | 类型 | 约束 | 说明 |
|-----|------|------|------|
| id | bigint unsigned | PK, AUTO_INCREMENT | 主键 |
| created_at | datetime(3) | | 创建时间 |
| role_id | bigint unsigned | NOT NULL | 角色ID |
| permission_id | bigint unsigned | NOT NULL | 权限ID |

---

### 2.2 频道管理模块

#### channel (频道表)

| 字段 | 类型 | 约束 | 说明 |
|-----|------|------|------|
| id | bigint unsigned | PK, AUTO_INCREMENT | 主键 |
| created_at | datetime(3) | | 创建时间 |
| updated_at | datetime(3) | | 更新时间 |
| deleted_at | datetime(3) | INDEX | 删除时间 |
| tenant_id | bigint unsigned | DEFAULT 1 | 租户ID |
| name | varchar(50) | NOT NULL | 频道名称 |
| code | varchar(50) | UNIQUE, NOT NULL | 频道代码 |
| icon | varchar(50) | | 图标(emoji) |
| description | varchar(200) | | 描述 |
| status | tinyint | DEFAULT 1 | 状态 |
| sort | int | DEFAULT 0 | 排序 |

**初始数据**: 7个频道 (搞笑、热门、美食、旅行、生活、科技、娱乐)

---

#### channel_config (频道配置表)

| 字段 | 类型 | 约束 | 说明 |
|-----|------|------|------|
| id | bigint unsigned | PK, AUTO_INCREMENT | 主键 |
| created_at | datetime(3) | | 创建时间 |
| updated_at | datetime(3) | | 更新时间 |
| deleted_at | datetime(3) | INDEX | 删除时间 |
| tenant_id | bigint unsigned | DEFAULT 1 | 租户ID |
| channel_id | bigint unsigned | UNIQUE, NOT NULL | 频道ID |
| content_type | json | | 内容类型配置 |
| display_mode | varchar(50) | DEFAULT 'default' | 显示模式 |
| custom_data | json | | 自定义数据 |
| page_config | json | | 页面配置 |

---

#### banner (Banner表)

| 字段 | 类型 | 约束 | 说明 |
|-----|------|------|------|
| id | bigint unsigned | PK, AUTO_INCREMENT | 主键 |
| created_at | datetime(3) | | 创建时间 |
| updated_at | datetime(3) | | 更新时间 |
| deleted_at | datetime(3) | INDEX | 删除时间 |
| tenant_id | bigint unsigned | DEFAULT 1 | 租户ID |
| channel_id | bigint unsigned | NOT NULL | 频道ID |
| title | varchar(100) | NOT NULL | 标题 |
| image | varchar(500) | NOT NULL | 图片URL |
| link_type | tinyint | DEFAULT 1 | 链接类型: 1-内部 2-外部 |
| link_url | varchar(500) | | 链接URL |
| content_id | bigint unsigned | | 内容ID |
| status | tinyint | DEFAULT 1 | 状态 |
| sort | int | DEFAULT 0 | 排序 |
| start_time | datetime | | 开始时间 |
| end_time | datetime | | 结束时间 |

---

#### diamond (金刚位表)

| 字段 | 类型 | 约束 | 说明 |
|-----|------|------|------|
| id | bigint unsigned | PK, AUTO_INCREMENT | 主键 |
| created_at | datetime(3) | | 创建时间 |
| updated_at | datetime(3) | | 更新时间 |
| deleted_at | datetime(3) | INDEX | 删除时间 |
| tenant_id | bigint unsigned | DEFAULT 1 | 租户ID |
| channel_id | bigint unsigned | NOT NULL | 频道ID |
| group_id | int | DEFAULT 1 | 分组ID |
| sort | int | DEFAULT 0 | 排序 |
| icon | varchar(50) | | 图标(emoji) |
| title | varchar(50) | NOT NULL | 标题 |
| link_type | varchar(20) | NOT NULL | 链接类型: channel/topic/content/external |
| link_value | varchar(500) | NOT NULL | 链接值 |
| status | tinyint | DEFAULT 1 | 状态 |
| description | varchar(200) | | 描述 |

---

#### recommend (推荐位表)

| 字段 | 类型 | 约束 | 说明 |
|-----|------|------|------|
| id | bigint unsigned | PK, AUTO_INCREMENT | 主键 |
| created_at | datetime(3) | | 创建时间 |
| updated_at | datetime(3) | | 更新时间 |
| deleted_at | datetime(3) | INDEX | 删除时间 |
| tenant_id | bigint unsigned | DEFAULT 1 | 租户ID |
| channel_id | bigint unsigned | NOT NULL | 频道ID |
| title | varchar(100) | NOT NULL | 标题 |
| subtitle | varchar(200) | | 副标题 |
| image | varchar(500) | | 图片URL |
| display_type | varchar(20) | DEFAULT 'single' | 展示类型: single/double/triple/list |
| source_type | varchar(20) | DEFAULT 'algorithm' | 内容来源: algorithm/manual/filter |
| content_ids | json | | 内容ID列表 |
| filter_rule | json | | 筛选规则 |
| link_url | varchar(500) | | 链接URL |
| status | tinyint | DEFAULT 1 | 状态 |
| sort | int | DEFAULT 0 | 排序 |
| description | varchar(200) | | 描述 |

---

#### feed_config (Feed流配置表)

| 字段 | 类型 | 约束 | 说明 |
|-----|------|------|------|
| id | bigint unsigned | PK, AUTO_INCREMENT | 主键 |
| created_at | datetime(3) | | 创建时间 |
| updated_at | datetime(3) | | 更新时间 |
| deleted_at | datetime(3) | INDEX | 删除时间 |
| tenant_id | bigint unsigned | DEFAULT 1 | 租户ID |
| channel_id | bigint unsigned | NOT NULL | 频道ID |
| title | varchar(100) | NOT NULL | 标题 |
| layout_type | varchar(20) | DEFAULT 'grid2' | 布局类型: grid2/grid3/list/waterfall |
| content_strategy | varchar(20) | DEFAULT 'algorithm' | 内容策略: algorithm/manual/random/tag/topic |
| content_ids | json | | 内容ID列表 |
| filter_rule | json | | 筛选规则 |
| auto_load | tinyint | DEFAULT 1 | 自动加载 |
| show_title | tinyint | DEFAULT 1 | 显示标题 |
| status | tinyint | DEFAULT 1 | 状态 |
| sort | int | DEFAULT 0 | 排序 |
| description | varchar(200) | | 描述 |

---

#### ad_slot (广告位表)

| 字段 | 类型 | 约束 | 说明 |
|-----|------|------|------|
| id | bigint unsigned | PK, AUTO_INCREMENT | 主键 |
| created_at | datetime(3) | | 创建时间 |
| updated_at | datetime(3) | | 更新时间 |
| deleted_at | datetime(3) | INDEX | 删除时间 |
| tenant_id | bigint unsigned | DEFAULT 1 | 租户ID |
| channel_id | bigint unsigned | NOT NULL | 频道ID |
| name | varchar(100) | NOT NULL | 名称 |
| code | varchar(50) | | 代码 |
| insert_type | varchar(20) | DEFAULT 'feed' | 插入方式: fixed/interval/random |
| insert_rule | json | | 插入规则 |
| ad_type | varchar(20) | DEFAULT 'image' | 广告类型: image/video |
| ad_content | json | | 广告内容 |
| description | varchar(200) | | 描述 |
| type | tinyint | DEFAULT 1 | 类型: 1-Banner 2-插屏 |
| image_url | varchar(500) | | 图片URL |
| video_url | varchar(500) | | 视频URL |
| link_url | varchar(500) | | 链接URL |
| status | tinyint | DEFAULT 1 | 状态 |
| sort | int | DEFAULT 0 | 排序 |
| start_time | datetime | | 开始时间 |
| end_time | datetime | | 结束时间 |

---

### 2.3 内容管理模块

#### material (物料表)

| 字段 | 类型 | 约束 | 说明 |
|-----|------|------|------|
| id | bigint unsigned | PK, AUTO_INCREMENT | 主键 |
| created_at | datetime(3) | | 创建时间 |
| updated_at | datetime(3) | | 更新时间 |
| deleted_at | datetime(3) | INDEX | 删除时间 |
| tenant_id | bigint unsigned | DEFAULT 1 | 租户ID |
| title | varchar(200) | NOT NULL | 标题 |
| type | varchar(20) | NOT NULL | 类型: article/image/video/novel/comic |
| cover | varchar(500) | | 封面URL |
| content | text | | 内容 |
| author_id | bigint unsigned | | 作者ID |
| author_name | varchar(50) | | 作者名称 |
| category_id | bigint unsigned | | 分类ID |
| tags | json | | 标签列表 |
| topic_id | bigint unsigned | | 话题ID |
| view_count | int | DEFAULT 0 | 浏览数 |
| like_count | int | DEFAULT 0 | 点赞数 |
| comment_count | int | DEFAULT 0 | 评论数 |
| share_count | int | DEFAULT 0 | 分享数 |
| collect_count | int | DEFAULT 0 | 收藏数 |
| duration | int | | 时长(秒) |
| chapter_count | int | DEFAULT 0 | 章节数 |
| status | tinyint | DEFAULT 1 | 状态: 0-草稿 1-发布 2-下架 |
| publish_time | datetime | | 发布时间 |
| is_original | tinyint | DEFAULT 1 | 是否原创 |
| source | varchar(100) | | 来源 |
| description | varchar(500) | | 描述 |

**索引**:
- KEY idx_type (type)
- KEY idx_author (author_id)
- KEY idx_category (category_id)
- KEY idx_topic (topic_id)
- KEY idx_status (status)

---

#### chapter (章节表)

| 字段 | 类型 | 约束 | 说明 |
|-----|------|------|------|
| id | bigint unsigned | PK, AUTO_INCREMENT | 主键 |
| created_at | datetime(3) | | 创建时间 |
| updated_at | datetime(3) | | 更新时间 |
| deleted_at | datetime(3) | INDEX | 删除时间 |
| tenant_id | bigint unsigned | DEFAULT 1 | 租户ID |
| material_id | bigint unsigned | NOT NULL | 物料ID |
| title | varchar(200) | NOT NULL | 章节标题 |
| content | longtext | | 章节内容 |
| sort | int | DEFAULT 0 | 排序 |
| word_count | int | DEFAULT 0 | 字数 |
| duration | int | | 时长(秒) |
| status | tinyint | DEFAULT 1 | 状态 |
| publish_time | datetime | | 发布时间 |

---

#### tag (标签表)

| 字段 | 类型 | 约束 | 说明 |
|-----|------|------|------|
| id | bigint unsigned | PK, AUTO_INCREMENT | 主键 |
| created_at | datetime(3) | | 创建时间 |
| updated_at | datetime(3) | | 更新时间 |
| deleted_at | datetime(3) | INDEX | 删除时间 |
| tenant_id | bigint unsigned | DEFAULT 1 | 租户ID |
| name | varchar(50) | NOT NULL | 标签名称 |
| type | tinyint | DEFAULT 0 | 类型: 0-通用 1-长视频 2-短视频 3-短剧 4-漫剧 5-小说 6-图文 |
| status | tinyint | DEFAULT 1 | 状态 |
| sort | int | DEFAULT 0 | 排序 |

---

#### topic (话题表)

| 字段 | 类型 | 约束 | 说明 |
|-----|------|------|------|
| id | bigint unsigned | PK, AUTO_INCREMENT | 主键 |
| created_at | datetime(3) | | 创建时间 |
| updated_at | datetime(3) | | 更新时间 |
| deleted_at | datetime(3) | INDEX | 删除时间 |
| tenant_id | bigint unsigned | DEFAULT 1 | 租户ID |
| name | varchar(100) | NOT NULL | 话题名称 |
| description | varchar(500) | | 描述 |
| cover | varchar(500) | | 封面URL |
| status | tinyint | DEFAULT 1 | 状态 |
| sort | int | DEFAULT 0 | 排序 |

---

### 2.4 发现页模块

#### discover_module (发现页模块表)

| 字段 | 类型 | 约束 | 说明 |
|-----|------|------|------|
| id | bigint unsigned | PK, AUTO_INCREMENT | 主键 |
| created_at | datetime(3) | | 创建时间 |
| updated_at | datetime(3) | | 更新时间 |
| deleted_at | datetime(3) | INDEX | 删除时间 |
| tenant_id | bigint unsigned | DEFAULT 1 | 租户ID |
| name | varchar(100) | NOT NULL | 模块名称 |
| type | varchar(50) | NOT NULL | 模块类型: banner/diamond/recommend/feed |
| status | tinyint | DEFAULT 1 | 状态 |
| sort | int | DEFAULT 0 | 排序 |

---

#### discover_content (发现页内容表)

| 字段 | 类型 | 约束 | 说明 |
|-----|------|------|------|
| id | bigint unsigned | PK, AUTO_INCREMENT | 主键 |
| created_at | datetime(3) | | 创建时间 |
| updated_at | datetime(3) | | 更新时间 |
| deleted_at | datetime(3) | INDEX | 删除时间 |
| tenant_id | bigint unsigned | DEFAULT 1 | 租户ID |
| module_id | bigint unsigned | NOT NULL | 模块ID |
| content_type | varchar(50) | NOT NULL | 内容类型: material/banner/diamond |
| content_id | bigint unsigned | NOT NULL | 内容ID |
| sort | int | DEFAULT 0 | 排序 |
| status | tinyint | DEFAULT 1 | 状态 |

---

### 2.5 用户模块

#### user (用户表)

| 字段 | 类型 | 约束 | 说明 |
|-----|------|------|------|
| id | bigint unsigned | PK, AUTO_INCREMENT | 主键 |
| created_at | datetime(3) | | 创建时间 |
| updated_at | datetime(3) | | 更新时间 |
| deleted_at | datetime(3) | INDEX | 删除时间 |
| tenant_id | bigint unsigned | DEFAULT 1 | 租户ID |
| username | varchar(50) | UNIQUE, NOT NULL | 用户名 |
| email | varchar(100) | | 邮箱 |
| phone | varchar(20) | | 手机号 |
| password | varchar(255) | NOT NULL | 密码 |
| nickname | varchar(50) | | 昵称 |
| avatar | varchar(500) | | 头像URL |
| gender | tinyint | DEFAULT 0 | 性别: 0-未知 1-男 2-女 |
| birthday | date | | 生日 |
| bio | varchar(500) | | 个人简介 |
| status | tinyint | DEFAULT 1 | 状态 |
| role | tinyint | DEFAULT 0 | 角色: 0-普通用户 1-创作者 |
| follow_count | int | DEFAULT 0 | 关注数 |
| fans_count | int | DEFAULT 0 | 粉丝数 |
| like_count | int | DEFAULT 0 | 获赞数 |
| third_party_id | varchar(100) | | 第三方平台ID |
| third_type | varchar(50) | | 第三方平台类型 |

---

### 2.6 互动模块

#### interaction (互动表)

| 字段 | 类型 | 约束 | 说明 |
|-----|------|------|------|
| id | bigint unsigned | PK, AUTO_INCREMENT | 主键 |
| created_at | datetime(3) | | 创建时间 |
| updated_at | datetime(3) | | 更新时间 |
| deleted_at | datetime(3) | INDEX | 删除时间 |
| tenant_id | bigint unsigned | DEFAULT 1 | 租户ID |
| user_id | bigint unsigned | NOT NULL | 用户ID |
| target_id | bigint unsigned | NOT NULL | 目标ID |
| target_type | varchar(50) | NOT NULL | 目标类型: material/comment/user |
| type | varchar(20) | NOT NULL | 互动类型: like/collect/share |
| status | tinyint | DEFAULT 1 | 状态: 0-取消 1-有效 |

**索引**:
- UNIQUE KEY uk_user_target (user_id, target_id, target_type, type)
- KEY idx_target (target_id, target_type)

---

## 3. 表关系

```
tenant (租户)
  ├── channel (频道) [1:N]
  │     ├── banner [1:N]
  │     ├── diamond [1:N]
  │     ├── recommend [1:N]
  │     ├── feed_config [1:N]
  │     ├── ad_slot [1:N]
  │     └── channel_config [1:1]
  │
  ├── material (物料) [1:N]
  │     ├── chapter (章节) [1:N]
  │     ├── tag (标签) [N:M]
  │     └── topic (话题) [N:M]
  │
  ├── user (用户) [1:N]
  │     ├── interaction (互动) [1:N]
  │     └── follow (关注) [1:N]
  │
  └── discover_module (发现页模块) [1:N]
        └── discover_content (发现页内容) [1:N]

admin_user (管理员)
  └── role (角色) [N:M]
        └── permission (权限) [N:M]
```

---

## 4. 数据库初始化

### 4.1 初始化脚本

```bash
# 导入表结构
mysql -uroot -proot123456 happy < backend/sql/rbac_tables.sql
mysql -uroot -proot123456 happy < backend/sql/material.sql
mysql -uroot -proot123456 happy < backend/sql/chapter.sql
mysql -uroot -proot123456 happy < backend/sql/missing_tables.sql

# 导入测试数据
mysql -uroot -proot123456 happy < backend/sql/init_data.sql
```

### 4.2 默认账号

**管理员账号**:
- 用户名: `admin`
- 密码: `admin123`
- 角色: 超级管理员

**测试用户**:
- user001, user002, blogger001, blogger002, blogger003
- 密码: `123456`

---

## 5. 注意事项

### 5.1 字段映射

- `admin_user.nickname` → 代码中使用 `realname`
- Repository层SQL必须使用: `SELECT nickname as realname`

### 5.2 字符集

- 所有表必须使用 `utf8mb4` 字符集
- 导入数据前执行: `SET NAMES utf8mb4`

### 5.3 软删除

- 使用 `deleted_at` 字段实现软删除
- 查询时必须过滤: `WHERE deleted_at IS NULL`

### 5.4 JSON字段

- `content_type`, `insert_rule`, `ad_content` 等使用JSON类型
- MySQL 5.7+ 支持JSON类型
- 查询时使用: `JSON_EXTRACT()`, `JSON_UNQUOTE()`
