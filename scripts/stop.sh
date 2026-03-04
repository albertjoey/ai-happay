#!/bin/bash

# Happy项目停止脚本

echo "================================"
echo "Happy 项目停止脚本"
echo "================================"

echo ""
echo "停止所有服务..."
docker-compose down

echo ""
echo "服务已停止"
echo ""
