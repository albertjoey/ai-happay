# 数据库初始化数据说明

## 文件说明

- `init_data.sql` - 包含所有测试数据的SQL文件

## 使用方法

### 方法1: 使用Docker导入

```bash
# 进入MySQL容器
docker exec -it happy-mysql mysql -uroot -proot123456

# 导入数据
source /path/to/init_data.sql
```

### 方法2: 使用docker exec导入

```bash
# 从宿主机导入数据到容器
docker exec -i happy-mysql mysql -uroot -proot123456 happy < backend/sql/init_data.sql
```

### 方法3: 使用MySQL客户端导入

```bash
# 使用mysql命令导入
mysql -h127.0.0.1 -P3306 -uroot -proot123456 happy < backend/sql/init_data.sql
```

## 数据内容

该文件包含以下测试数据:

### 系统管理数据
- **admin_user** - 管理员账号 (admin/admin123)
- **role** - 角色数据 (超级管理员、管理员、运营)
- **permission** - 权限数据
- **admin_user_roles** - 管理员角色关联

### 频道管理数据
- **channel** - 频道数据 (7个频道)
- **banner** - Banner广告数据
- **diamond** - 金刚位数据
- **recommend** - 推荐位数据
- **feed_config** - Feed流配置
- **ad_slot** - 广告位数据

### 内容管理数据
- **material** - 内容素材数据 (2800+条)
- **tag** - 标签数据
- **topic** - 话题数据
- **category** - 分类数据

### 用户数据
- **user** - 用户数据 (5个测试用户)
- **user_like** - 用户点赞记录
- **user_collect** - 用户收藏记录
- **view_history** - 观看历史记录

### 其他数据
- **tenant** - 租户数据
- **discover_config** - 发现页配置
- **discover_content** - 发现页内容

## 注意事项

1. 导入前请确保数据库已创建表结构
2. 该文件只包含数据,不包含表结构
3. 导入会覆盖现有数据,请谨慎操作
4. 建议在开发环境使用,生产环境请使用真实数据

## 默认账号

### 管理员账号
- 用户名: `admin`
- 密码: `admin123`
- 角色: 超级管理员

### 测试用户账号
- user001 / user002 / blogger001 / blogger002 / blogger003
- 密码均为: `123456`

## 更新数据

如果需要更新测试数据,请执行以下命令重新导出:

```bash
docker exec happy-mysql mysqldump -uroot -proot123456 \
  --default-character-set=utf8mb4 \
  --databases happy \
  --no-create-info \
  --complete-insert \
  --skip-add-locks \
  --skip-disable-keys \
  > backend/sql/init_data.sql
```
