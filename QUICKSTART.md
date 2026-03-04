# Happy 项目快速启动指南

## 一、环境准备

### 必需软件
- Docker & Docker Compose
- Go 1.21+ (可选，用于本地开发)
- Node.js 18+ (可选，用于本地开发)

## 二、快速启动

### 方式一：使用Docker Compose（推荐）

```bash
# 1. 启动基础服务
./scripts/start.sh

# 或者手动启动
docker-compose up -d

# 2. 等待服务启动（约10-30秒）
docker-compose ps

# 3. 初始化数据库（自动执行，也可手动）
docker exec -i happy-mysql mysql -uroot -proot123456 < scripts/sql/init.sql
```

### 方式二：本地开发模式

```bash
# 1. 启动基础服务
docker-compose up -d mysql redis kafka elasticsearch

# 2. 启动后端服务
cd backend
go mod tidy
go run app/user/cmd/user.go -f app/user/etc/user.yaml

# 3. 启动H5前端（新终端）
cd frontend/h5
npm install
npm run dev

# 4. 启动管理后台（新终端）
cd frontend/admin
npm install
npm run dev
```

## 三、访问地址

### 基础服务
- **MySQL**: localhost:3306 (root/root123456)
- **Redis**: localhost:6379 (密码: redis123456)
- **Elasticsearch**: http://localhost:9200
- **Kafka**: localhost:9092
- **Kafka UI**: http://localhost:9000
- **Redis Commander**: http://localhost:8081

### 应用服务
- **H5前端**: http://localhost:3000
- **管理后台**: http://localhost:3001
- **用户服务API**: http://localhost:8001
- **内容服务API**: http://localhost:8002

## 四、默认账号

### 管理后台
- 用户名: admin
- 密码: admin123

### 数据库
- 用户名: root
- 密码: root123456

## 五、常用命令

```bash
# 查看服务状态
docker-compose ps

# 查看服务日志
docker-compose logs -f [service_name]

# 停止所有服务
./scripts/stop.sh
# 或
docker-compose down

# 重启服务
docker-compose restart [service_name]

# 清理所有数据（危险操作）
docker-compose down -v
```

## 六、项目结构

```
happy/
├── backend/          # 后端Go代码
├── frontend/
│   ├── h5/          # H5移动端
│   └── admin/       # 管理后台
├── docs/            # 文档
├── scripts/         # 脚本
└── docker-compose.yml
```

## 七、下一步

1. 查看详细文档: [docs/DEVELOPMENT.md](docs/DEVELOPMENT.md)
2. API文档: 访问 http://localhost:8001/swagger (需启用swagger)
3. 开始开发: 参考 [开发指南](docs/DEVELOPMENT.md)

## 八、问题排查

### MySQL连接失败
```bash
# 检查MySQL是否启动
docker-compose ps mysql

# 查看MySQL日志
docker-compose logs mysql

# 重启MySQL
docker-compose restart mysql
```

### 前端无法访问后端API
1. 检查后端服务是否启动
2. 检查端口是否被占用
3. 检查防火墙设置

### Docker服务启动失败
```bash
# 查看详细错误信息
docker-compose up

# 清理并重新启动
docker-compose down -v
docker-compose up -d
```

## 九、技术支持

如遇问题，请查看:
1. [开发文档](docs/DEVELOPMENT.md)
2. [README](README.md)
3. 提交Issue
