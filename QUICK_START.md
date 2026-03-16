# 快速启动指南 (Quick Start Guide)

> 本文档说明如何让AI一次性实现完整项目功能

---

## 📋 需要提供的完整信息清单

### 1. 项目规范文档 ✅
**文件**: `PROJECT_SPEC.md`

**内容**:
- 架构规范 (Handler → Logic → Repository)
- 技术栈约束 (Vue 3 + Ant Design Vue)
- 数据库规范 (字段映射、字符集)
- 开发流程规范
- 常见问题解决方案

**作用**: 让AI理解项目的技术标准和约束

---

### 2. 数据库设计文档 ⚠️ (必需)

**文件**: `DATABASE_DESIGN.md`

**内容**:
```markdown
# 数据库设计

## 1. 表结构设计

### channel (频道表)
| 字段 | 类型 | 说明 |
|-----|------|------|
| id | bigint | 主键 |
| name | varchar(50) | 频道名称 |
| code | varchar(50) | 频道代码 |
| status | tinyint | 状态 |

### admin_user (管理员表)
| 字段 | 类型 | 说明 |
|-----|------|------|
| id | bigint | 主键 |
| username | varchar(50) | 用户名 |
| nickname | varchar(50) | 昵称 (映射为realname) |
| password | varchar(255) | 密码 |

## 2. 表关系
- admin_user_roles: admin_user ↔ role
- channel_config: channel的配置

## 3. 初始化数据
- 管理员: admin/admin123
- 角色: 超级管理员、管理员、运营
- 频道: 7个测试频道
```

**作用**: 让AI知道数据库结构,避免字段映射错误

---

### 3. 功能需求文档 ⚠️ (必需)

**文件**: `REQUIREMENTS.md`

**内容**:
```markdown
# 功能需求

## 1. 系统管理模块

### 1.1 角色管理
- 功能: 角色的增删改查
- 字段: name, code, description, status
- 权限: 可以分配权限给角色

### 1.2 管理员管理
- 功能: 管理员的增删改查
- 字段: username, nickname, email, phone, status
- 角色: 可以分配角色给管理员

### 1.3 权限管理
- 功能: 权限的树形管理
- 字段: name, code, type, parent_id, path
- 结构: 树形结构,支持父子关系

## 2. 频道管理模块

### 2.1 频道列表
- 功能: 频道的增删改查
- 字段: name, code, description, status

### 2.2 频道配置
- 功能: 配置频道的内容类型、显示模式
- 配置项: Banner、金刚位、推荐位、Feed流

### 2.3 金刚位管理
- 功能: 金刚位的增删改查
- 字段: name, icon, link_type, link_url
- 关联: channel_id

### 2.4 推荐位管理
- 功能: 推荐位的增删改查
- 字段: title, subtitle, image, link_url
- 关联: channel_id

### 2.5 Feed流配置
- 功能: Feed流的配置管理
- 字段: name, type, config
- 关联: channel_id

### 2.6 广告位管理
- 功能: 广告位的增删改查
- 字段: name, insert_type, insert_rule, ad_type
- 关联: channel_id

## 3. 发现页管理模块

### 3.1 模块配置
- 功能: 发现页模块的配置
- 字段: name, type, status, sort

### 3.2 内容管理
- 功能: 模块内容的管理
- 字段: module_id, content_type, content_id, sort
```

**作用**: 让AI知道要实现哪些功能

---

### 4. API接口文档 ⚠️ (可选,但推荐)

**文件**: `API_DESIGN.md`

**内容**:
```markdown
# API接口设计

## 1. 频道管理API

### GET /api/v1/channel/list
- 说明: 获取频道列表
- 参数: page, page_size
- 响应: { list: [], total: 0 }

### POST /api/v1/channel
- 说明: 创建频道
- 参数: { name, code, description }
- 响应: { id: 1 }

## 2. 管理员管理API

### GET /api/v1/admin-user/list
- 说明: 获取管理员列表
- 参数: page, page_size
- 响应: { list: [], total: 0 }

### GET /api/v1/admin-user/:id/roles
- 说明: 获取管理员角色
- 响应: [1, 2, 3]

### POST /api/v1/admin-user/:id/roles
- 说明: 分配管理员角色
- 参数: [1, 2, 3]
```

**作用**: 让AI知道API的设计规范

---

### 5. 页面设计文档 ⚠️ (可选,但推荐)

**文件**: `UI_DESIGN.md`

**内容**:
```markdown
# 页面设计

## 1. 管理后台布局

### 1.1 整体布局
- 左侧: 菜单导航
- 右侧: 内容区域
- 顶部: 用户信息、退出按钮

### 1.2 菜单结构
- 仪表盘
- 发现页管理
  - 模块配置
  - 内容管理
- 频道管理
  - 频道列表
  - 金刚位管理
  - 推荐位管理
  - Feed流配置
  - 广告位管理
  - 频道配置
- 系统管理
  - 角色管理
  - 权限管理
  - 管理员管理

## 2. 页面组件

### 2.1 列表页面
- 顶部: 频道选择器 + 添加按钮
- 中间: 数据表格 (使用a-table)
- 底部: 分页器

### 2.2 表单页面
- 使用a-modal弹窗
- 使用a-form表单
- 支持添加和编辑
```

**作用**: 让AI知道页面的设计要求

---

### 6. 技术栈说明 ⚠️ (必需)

**文件**: `TECH_STACK.md`

**内容**:
```markdown
# 技术栈

## 1. 后端技术栈
- 框架: Go-Zero
- ORM: GORM
- 数据库: MySQL 8.0
- 字符集: utf8mb4

## 2. 前端技术栈
- 框架: Vue 3
- UI库: Ant Design Vue
- 路由: Vue Router
- 状态管理: Pinia
- 构建工具: Vite

## 3. 禁止使用
- ❌ vxe-table (存在缓存问题)
- ❌ Logic层直接写SQL
- ❌ 前端直接操作DOM

## 4. 必须使用
- ✅ Repository模式
- ✅ a-table组件
- ✅ router-view的key属性
```

**作用**: 让AI知道技术选型和约束

---

## 🚀 完整启动流程

### 步骤1: 准备文档

在项目根目录创建以下文件:
```
project/
├── PROJECT_SPEC.md        # 项目规范
├── DATABASE_DESIGN.md     # 数据库设计
├── REQUIREMENTS.md        # 功能需求
├── API_DESIGN.md          # API设计 (可选)
├── UI_DESIGN.md           # 页面设计 (可选)
└── TECH_STACK.md          # 技术栈
```

### 步骤2: 向AI提问

**提问模板**:
```markdown
参考以下文档,实现完整的项目功能:

1. PROJECT_SPEC.md - 项目规范
2. DATABASE_DESIGN.md - 数据库设计
3. REQUIREMENTS.md - 功能需求
4. API_DESIGN.md - API设计
5. UI_DESIGN.md - 页面设计
6. TECH_STACK.md - 技术栈

请按照以下顺序实现:
1. 创建数据库和表结构
2. 实现后端API (Handler → Logic → Repository)
3. 实现前端页面 (API → 组件 → 路由)
4. 导入测试数据
5. 测试功能是否正常

要求:
- 严格按照PROJECT_SPEC.md中的架构规范
- 处理好字段映射关系
- 处理好空值和错误情况
- 确保路由切换正常
```

### 步骤3: AI实现

AI会按照以下顺序实现:
1. ✅ 创建数据库表结构
2. ✅ 创建Repository层
3. ✅ 创建Logic层
4. ✅ 创建Handler层
5. ✅ 创建前端API
6. ✅ 创建前端组件
7. ✅ 注册路由
8. ✅ 导入测试数据
9. ✅ 测试功能

---

## 📊 文档重要性排序

| 文档 | 重要性 | 说明 |
|-----|--------|------|
| PROJECT_SPEC.md | ⭐⭐⭐⭐⭐ | 必需,定义架构规范 |
| DATABASE_DESIGN.md | ⭐⭐⭐⭐⭐ | 必需,定义数据结构 |
| REQUIREMENTS.md | ⭐⭐⭐⭐⭐ | 必需,定义功能需求 |
| TECH_STACK.md | ⭐⭐⭐⭐ | 必需,定义技术栈 |
| API_DESIGN.md | ⭐⭐⭐ | 推荐,定义API规范 |
| UI_DESIGN.md | ⭐⭐⭐ | 推荐,定义UI规范 |

---

## 💡 最小化文档集

如果只想提供最少的文档,至少需要:

### 必需文档 (3个)
1. **PROJECT_SPEC.md** - 架构规范
2. **DATABASE_DESIGN.md** - 数据库设计
3. **REQUIREMENTS.md** - 功能需求

### 提问示例
```markdown
参考 PROJECT_SPEC.md、DATABASE_DESIGN.md、REQUIREMENTS.md,
实现完整的项目功能。

技术栈:
- 后端: Go-Zero + GORM + MySQL
- 前端: Vue 3 + Ant Design Vue
- 禁止: vxe-table、Logic层写SQL

请一次性实现所有功能,包括:
1. 数据库表结构
2. 后端API (Repository模式)
3. 前端页面 (a-table)
4. 测试数据
```

---

## 🎯 总结

### 单独一个文件不够的原因

1. **PROJECT_SPEC.md** 只定义了规范,没有具体需求
2. 需要知道**数据库结构**才能正确映射字段
3. 需要知道**功能需求**才能实现具体功能
4. 需要知道**技术栈**才能选择正确的组件

### 最小文档集

- ✅ PROJECT_SPEC.md (架构规范)
- ✅ DATABASE_DESIGN.md (数据结构)
- ✅ REQUIREMENTS.md (功能需求)

### 完整文档集

- ✅ PROJECT_SPEC.md
- ✅ DATABASE_DESIGN.md
- ✅ REQUIREMENTS.md
- ✅ TECH_STACK.md
- ✅ API_DESIGN.md
- ✅ UI_DESIGN.md

提供完整文档集,AI就能一次性实现所有功能,无需反复沟通! 🎯
