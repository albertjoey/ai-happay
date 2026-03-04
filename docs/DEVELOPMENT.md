# Happy 项目开发指南

## 目录结构

```
happy/
├── backend/                 # 后端代码
│   ├── app/                # 微服务应用
│   │   ├── user/          # 用户服务
│   │   │   ├── cmd/       # 启动入口
│   │   │   ├── internal/  # 内部代码
│   │   │   │   ├── config/    # 配置
│   │   │   │   ├── handler/   # 处理器
│   │   │   │   ├── logic/     # 业务逻辑
│   │   │   │   ├── model/     # 数据模型
│   │   │   │   ├── svc/       # 服务上下文
│   │   │   │   └── types/     # 类型定义
│   │   │   └── etc/       # 配置文件
│   │   ├── content/       # 内容服务
│   │   ├── interaction/   # 互动服务
│   │   ├── search/        # 搜索服务
│   │   └── admin/         # 管理服务
│   ├── common/            # 公共代码
│   │   ├── config/        # 配置定义
│   │   ├── errors/        # 错误定义
│   │   ├── middleware/    # 中间件
│   │   ├── model/         # 公共模型
│   │   └── utils/         # 工具函数
│   ├── proto/             # Protobuf定义
│   └── deploy/            # 部署配置
│       └── k8s/           # Kubernetes配置
├── frontend/              # 前端代码
│   ├── h5/               # H5移动端
│   │   ├── src/
│   │   │   ├── app/      # Next.js App Router
│   │   │   ├── components/   # 组件
│   │   │   ├── lib/      # 工具库
│   │   │   ├── store/    # 状态管理
│   │   │   └── types/    # 类型定义
│   │   └── public/       # 静态资源
│   └── admin/            # 运营管理后台
│       ├── src/
│       │   ├── views/    # 页面
│       │   ├── components/   # 组件
│       │   ├── router/   # 路由
│       │   ├── store/    # 状态管理
│       │   ├── api/      # API接口
│       │   └── utils/    # 工具函数
│       └── public/       # 静态资源
├── docs/                  # 文档
└── scripts/              # 脚本工具
    └── sql/              # SQL脚本
```

## 技术栈

### 后端
- **语言**: Go 1.21+
- **框架**: go-zero (微服务)
- **数据库**: MySQL 8.0, Elasticsearch 8.x
- **缓存**: Redis 7.x
- **消息队列**: Kafka 3.x
- **部署**: Docker, Kubernetes

### 前端
- **H5/Web**: Next.js 14 (SSR, SEO优化)
- **管理后台**: Vue 3 + vxe-table
- **UI框架**: Tailwind CSS, Ant Design

## 开发环境搭建

### 1. 安装依赖工具

```bash
# Go
brew install go

# Node.js
brew install node

# Docker
brew install docker docker-compose

# Kubernetes (可选)
brew install kubectl minikube
```

### 2. 启动基础服务

```bash
# 启动MySQL、Redis、Kafka、Elasticsearch
docker-compose up -d

# 查看服务状态
docker-compose ps

# 查看日志
docker-compose logs -f
```

### 3. 初始化数据库

```bash
# 数据库会自动初始化，也可以手动执行
mysql -h 127.0.0.1 -P 3306 -u root -proot123456 < scripts/sql/init.sql
```

### 4. 启动后端服务

```bash
cd backend

# 安装依赖
go mod tidy

# 启动用户服务
go run app/user/cmd/user.go -f app/user/etc/user.yaml

# 启动内容服务
go run app/content/cmd/content.go -f app/content/etc/content.yaml
```

### 5. 启动前端服务

```bash
# H5移动端
cd frontend/h5
npm install
npm run dev

# 管理后台
cd frontend/admin
npm install
npm run dev
```

## 服务端口

| 服务 | 端口 | 说明 |
|------|------|------|
| MySQL | 3306 | 数据库 |
| Redis | 6379 | 缓存 |
| Elasticsearch | 9200 | 搜索引擎 |
| Kafka | 9092 | 消息队列 |
| Kafka UI | 9000 | Kafka管理界面 |
| Redis Commander | 8081 | Redis管理界面 |
| User Service | 8001 | 用户服务 |
| Content Service | 8002 | 内容服务 |
| H5 Frontend | 3000 | H5前端 |
| Admin Frontend | 3001 | 管理后台 |

## API文档

### 用户服务 API

#### 注册
```
POST /api/v1/user/register
Content-Type: application/json

{
  "username": "test",
  "email": "test@example.com",
  "password": "password123",
  "nickname": "测试用户"
}
```

#### 登录
```
POST /api/v1/user/login
Content-Type: application/json

{
  "username": "test",
  "password": "password123"
}
```

#### 获取用户信息
```
GET /api/v1/user/info
Authorization: Bearer {token}
```

#### 更新用户信息
```
PUT /api/v1/user/info
Authorization: Bearer {token}
Content-Type: application/json

{
  "nickname": "新昵称",
  "avatar": "https://example.com/avatar.jpg",
  "bio": "个人简介"
}
```

#### 关注用户
```
POST /api/v1/user/follow
Authorization: Bearer {token}
Content-Type: application/json

{
  "user_id": 2
}
```

### 内容服务 API

#### 创建内容
```
POST /api/v1/content
Authorization: Bearer {token}
Content-Type: application/json

{
  "title": "内容标题",
  "description": "内容描述",
  "type": 2,
  "media": [
    {
      "type": 2,
      "url": "https://example.com/video.mp4",
      "thumbnail": "https://example.com/thumb.jpg",
      "duration": 60
    }
  ]
}
```

#### 获取内容列表
```
GET /api/v1/content/list?page=1&page_size=20&type=2
Authorization: Bearer {token}
```

#### 获取内容详情
```
GET /api/v1/content/{id}
Authorization: Bearer {token}
```

## 数据库设计

### 核心表

- **tenant**: 租户/站点表
- **user**: 用户表
- **content**: 内容表
- **content_media**: 内容媒体资源表
- **topic**: 话题表
- **tag**: 标签表
- **channel**: 频道表
- **banner**: Banner表
- **diamond_position**: 金刚位表
- **feed_config**: Feed流配置表
- **ad_position**: 广告位表
- **interaction**: 互动表
- **comment**: 评论表
- **follow**: 关注表
- **view_history**: 观看历史表
- **message**: 消息通知表
- **admin_user**: 管理员用户表
- **role**: 角色表
- **permission**: 权限表

## 部署指南

### Docker部署

```bash
# 构建镜像
docker build -t happy/user-service:latest -f backend/Dockerfile backend/
docker build -t happy/h5-frontend:latest -f frontend/h5/Dockerfile frontend/h5/
docker build -t happy/admin-frontend:latest -f frontend/admin/Dockerfile frontend/admin/

# 运行容器
docker run -d -p 8001:8001 happy/user-service:latest
docker run -d -p 3000:3000 happy/h5-frontend:latest
docker run -d -p 80:80 happy/admin-frontend:latest
```

### Kubernetes部署

```bash
# 创建命名空间
kubectl apply -f backend/deploy/k8s/backend-deployment.yaml

# 部署数据库
kubectl apply -f backend/deploy/k8s/database-deployment.yaml

# 部署后端服务
kubectl apply -f backend/deploy/k8s/backend-deployment.yaml

# 部署前端服务
kubectl apply -f backend/deploy/k8s/frontend-deployment.yaml

# 查看部署状态
kubectl get pods -n happy
kubectl get services -n happy
kubectl get ingress -n happy
```

## 开发规范

### Git提交规范

```
feat: 新功能
fix: 修复bug
docs: 文档更新
style: 代码格式调整
refactor: 重构
test: 测试相关
chore: 构建/工具相关
```

### 代码规范

- Go代码遵循 [Effective Go](https://golang.org/doc/effective_go)
- TypeScript代码遵循 [Airbnb JavaScript Style Guide](https://github.com/airbnb/javascript)
- 使用ESLint和Prettier进行代码格式化

### 分支管理

- `main`: 主分支，生产环境代码
- `develop`: 开发分支
- `feature/*`: 功能分支
- `hotfix/*`: 紧急修复分支

## 常见问题

### 1. 数据库连接失败
检查MySQL是否启动，端口是否正确，用户名密码是否正确。

### 2. Redis连接失败
检查Redis是否启动，密码是否正确。

### 3. 前端跨域问题
开发环境已配置代理，生产环境需要配置Nginx或Ingress。

### 4. Kafka连接失败
检查Kafka和Zookeeper是否启动，端口是否正确。

## 联系方式

如有问题，请提交Issue或联系开发团队。
