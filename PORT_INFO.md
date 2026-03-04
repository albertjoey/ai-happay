# Happy 项目端口配置说明

## 🚀 服务端口列表（已更新到4000段）

### 应用服务端口

| 服务名称 | 端口 | 访问地址 | 说明 |
|---------|------|---------|------|
| H5移动端 | **4000** | http://localhost:4000 | Next.js前端应用 |
| 用户服务API | **4001** | http://localhost:4001 | 后端微服务 |
| 管理后台 | **4002** | http://localhost:4002 | Vue3管理后台 |

### 基础服务端口

| 服务名称 | 端口 | 访问地址 | 说明 |
|---------|------|---------|------|
| MySQL | **3306** | localhost:3306 | 数据库服务 |
| Redis | **6379** | localhost:6379 | 缓存服务 |

## 快速访问

### 1. H5移动端
- **地址**: http://localhost:4000
- **功能**: 移动端H5页面，响应式设计
- **特点**: 频道切换、内容瀑布流、底部导航

### 2. 用户服务API
- **地址**: http://localhost:4001
- **测试接口**:
  ```bash
  # 获取用户列表
  curl "http://localhost:4001/api/v1/user/list?page=1&page_size=10"

  # 用户注册
  curl -X POST http://localhost:4001/api/v1/user/register \
    -H "Content-Type: application/json" \
    -d '{"username":"test","email":"test@example.com","password":"password123"}'

  # 用户登录
  curl -X POST http://localhost:4001/api/v1/user/login \
    -H "Content-Type: application/json" \
    -d '{"email":"test@example.com","password":"password123"}'
  ```

### 3. 管理后台
- **地址**: http://localhost:4002
- **登录账号**: admin / admin123
- **功能模块**:
  - 仪表盘
  - 用户管理
  - 内容管理
  - 话题管理
  - 标签管理
  - 频道管理
  - 系统管理

## 配置文件位置

### 后端服务配置
- **文件**: `backend/app/user/etc/user.yaml`
- **端口配置**: `Port: 4001`

### H5前端配置
- **文件**: `frontend/h5/package.json`
- **启动命令**: `"dev": "next dev -p 4000"`
- **API代理**: `frontend/h5/next.config.js` (代理到4001)

### 管理后台配置
- **文件**: `frontend/admin/vite.config.ts`
- **端口配置**: `port: 4002`
- **API代理**: 代理到4001

## 服务管理命令

### 启动服务
```bash
# 启动后端服务
cd backend && go run app/user/cmd/user.go -f app/user/etc/user.yaml &

# 启动H5前端
cd frontend/h5 && npm run dev &

# 启动管理后台
cd frontend/admin && npm run dev &
```

### 停止服务
```bash
# 停止后端服务
lsof -ti :4001 | xargs kill -9

# 停止H5前端
lsof -ti :4000 | xargs kill -9

# 停止管理后台
lsof -ti :4002 | xargs kill -9
```

### 查看服务状态
```bash
# 查看所有服务端口
lsof -i :4000
lsof -i :4001
lsof -i :4002
```

## 数据库连接信息

### MySQL
- Host: localhost
- Port: 3306
- User: root
- Password: happy123456
- Database: happy

### Redis
- Host: localhost
- Port: 6379
- Password: (无密码)

## 测试数据

已填充的测试数据：
- 5个测试用户
- 10条测试内容
- 5个热门话题
- 10个常用标签
- 5条评论
- 5条关注关系
- 5条互动数据

## 注意事项

1. 所有应用服务端口已统一到4000段
2. 基础服务（MySQL、Redis）保持原端口不变
3. 前端应用已配置API代理，自动转发到后端4001端口
4. 如需修改端口，请同步更新相关配置文件

---

**更新时间**: 2026-02-28
**版本**: v1.0.0
**状态**: ✅ 所有服务运行正常
