.PHONY: help start stop restart build clean dev test

help: ## 显示帮助信息
	@echo "Happy 项目管理命令"
	@echo ""
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'

start: ## 启动所有服务
	@echo "启动基础服务..."
	docker-compose up -d
	@echo "等待服务启动..."
	@sleep 10
	@echo "初始化数据库..."
	@docker exec -i happy-mysql mysql -uroot -proot123456 < scripts/sql/init.sql 2>/dev/null || true
	@echo "服务启动完成！"

stop: ## 停止所有服务
	@echo "停止所有服务..."
	docker-compose down
	@echo "服务已停止"

restart: stop start ## 重启所有服务

build-backend: ## 构建后端服务
	@echo "构建后端服务..."
	cd backend && docker build -t happy/user-service:latest .

build-h5: ## 构建H5前端
	@echo "构建H5前端..."
	cd frontend/h5 && docker build -t happy/h5-frontend:latest .

build-admin: ## 构建管理后台
	@echo "构建管理后台..."
	cd frontend/admin && docker build -t happy/admin-frontend:latest .

build: build-backend build-h5 build-admin ## 构建所有服务

dev-backend: ## 启动后端开发环境
	@echo "启动后端开发环境..."
	cd backend && go mod tidy && go run app/user/cmd/user.go -f app/user/etc/user.yaml

dev-h5: ## 启动H5前端开发环境
	@echo "启动H5前端开发环境..."
	cd frontend/h5 && npm install && npm run dev

dev-admin: ## 启动管理后台开发环境
	@echo "启动管理后台开发环境..."
	cd frontend/admin && npm install && npm run dev

clean: ## 清理所有容器和数据
	@echo "清理所有容器和数据..."
	docker-compose down -v
	@echo "清理完成"

logs: ## 查看服务日志
	docker-compose logs -f

ps: ## 查看服务状态
	docker-compose ps

test: ## 运行测试
	@echo "运行测试..."
	cd backend && go test ./...

install: ## 安装依赖
	@echo "安装后端依赖..."
	cd backend && go mod tidy
	@echo "安装H5前端依赖..."
	cd frontend/h5 && npm install
	@echo "安装管理后台依赖..."
	cd frontend/admin && npm install
	@echo "依赖安装完成"

k8s-deploy: ## 部署到Kubernetes
	@echo "部署到Kubernetes..."
	kubectl apply -f backend/deploy/k8s/

k8s-delete: ## 从Kubernetes删除
	@echo "从Kubernetes删除..."
	kubectl delete -f backend/deploy/k8s/

k8s-ps: ## 查看Kubernetes部署状态
	kubectl get pods -n happy
	kubectl get services -n happy
