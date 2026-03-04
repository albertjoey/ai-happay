#!/bin/bash

# Happy项目快速启动脚本

echo "================================"
echo "Happy 项目启动脚本"
echo "================================"

# 检查Docker是否安装
if ! command -v docker &> /dev/null; then
    echo "错误: Docker未安装，请先安装Docker"
    exit 1
fi

# 检查Docker Compose是否安装
if ! command -v docker-compose &> /dev/null; then
    echo "错误: Docker Compose未安装，请先安装Docker Compose"
    exit 1
fi

echo ""
echo "1. 启动基础服务 (MySQL, Redis, Kafka, Elasticsearch)..."
docker-compose up -d mysql redis zookeeper kafka elasticsearch

echo ""
echo "等待服务启动..."
sleep 10

echo ""
echo "2. 检查服务状态..."
docker-compose ps

echo ""
echo "3. 初始化数据库..."
# 等待MySQL完全启动
echo "等待MySQL启动..."
until docker exec happy-mysql mysqladmin ping -h localhost -uroot -proot123456 --silent; do
    echo "MySQL未就绪，等待中..."
    sleep 2
done

echo "MySQL已就绪，执行初始化脚本..."
docker exec -i happy-mysql mysql -uroot -proot123456 < scripts/sql/init.sql

echo ""
echo "================================"
echo "基础服务启动完成！"
echo "================================"
echo ""
echo "服务地址:"
echo "  MySQL:           localhost:3306 (root/root123456)"
echo "  Redis:           localhost:6379 (密码: redis123456)"
echo "  Elasticsearch:   http://localhost:9200"
echo "  Kafka:           localhost:9092"
echo "  Kafka UI:        http://localhost:9000"
echo "  Redis Commander: http://localhost:8081"
echo ""
echo "下一步:"
echo "  1. 启动后端服务: cd backend && go run app/user/cmd/user.go -f app/user/etc/user.yaml"
echo "  2. 启动H5前端:   cd frontend/h5 && npm install && npm run dev"
echo "  3. 启动管理后台:  cd frontend/admin && npm install && npm run dev"
echo ""
