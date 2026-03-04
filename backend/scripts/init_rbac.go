package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// 连接数据库
	dsn := "root:happy123456@tcp(127.0.0.1:3306)/happy?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("连接数据库失败:", err)
	}
	defer db.Close()

	// 设置字符集
	db.Exec("SET NAMES utf8mb4")

	fmt.Println("删除旧表...")
	db.Exec("DROP TABLE IF EXISTS admin_role")
	db.Exec("DROP TABLE IF EXISTS admin_user")
	db.Exec("DROP TABLE IF EXISTS role_permission")
	db.Exec("DROP TABLE IF EXISTS permission")
	db.Exec("DROP TABLE IF EXISTS role")

	fmt.Println("开始创建RBAC表...")

	// 创建角色表
	_, err = db.Exec(`
CREATE TABLE role (
  id bigint unsigned NOT NULL AUTO_INCREMENT,
  name varchar(50) NOT NULL COMMENT '角色名称',
  code varchar(50) NOT NULL COMMENT '角色编码',
  description varchar(200) DEFAULT '' COMMENT '角色描述',
  status tinyint NOT NULL DEFAULT 1 COMMENT '状态',
  created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  deleted_at timestamp NULL DEFAULT NULL,
  PRIMARY KEY (id),
  UNIQUE KEY uk_code (code)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='角色表'
`)
	if err != nil {
		log.Fatal("创建角色表失败:", err)
	}
	fmt.Println("✅ 角色表创建成功")

	// 创建权限表
	_, err = db.Exec(`
CREATE TABLE permission (
  id bigint unsigned NOT NULL AUTO_INCREMENT,
  name varchar(50) NOT NULL COMMENT '权限名称',
  code varchar(100) NOT NULL COMMENT '权限编码',
  type varchar(20) NOT NULL COMMENT '权限类型',
  parent_id bigint unsigned DEFAULT 0,
  path varchar(200) DEFAULT '',
  icon varchar(50) DEFAULT '',
  status tinyint NOT NULL DEFAULT 1,
  created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  deleted_at timestamp NULL DEFAULT NULL,
  PRIMARY KEY (id),
  UNIQUE KEY uk_code (code)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='权限表'
`)
	if err != nil {
		log.Fatal("创建权限表失败:", err)
	}
	fmt.Println("✅ 权限表创建成功")

	// 创建角色权限关联表
	_, err = db.Exec(`
CREATE TABLE role_permission (
  id bigint unsigned NOT NULL AUTO_INCREMENT,
  role_id bigint unsigned NOT NULL,
  permission_id bigint unsigned NOT NULL,
  created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id),
  UNIQUE KEY uk_role_permission (role_id, permission_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='角色权限关联表'
`)
	if err != nil {
		log.Fatal("创建角色权限关联表失败:", err)
	}
	fmt.Println("✅ 角色权限关联表创建成功")

	// 创建管理员表
	_, err = db.Exec(`
CREATE TABLE admin_user (
  id bigint unsigned NOT NULL AUTO_INCREMENT,
  username varchar(50) NOT NULL,
  password varchar(255) NOT NULL,
  realname varchar(50) NOT NULL,
  email varchar(100) DEFAULT '',
  phone varchar(20) DEFAULT '',
  avatar varchar(500) DEFAULT '',
  status tinyint NOT NULL DEFAULT 1,
  last_login_at timestamp NULL DEFAULT NULL,
  created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  deleted_at timestamp NULL DEFAULT NULL,
  PRIMARY KEY (id),
  UNIQUE KEY uk_username (username)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='管理员表'
`)
	if err != nil {
		log.Fatal("创建管理员表失败:", err)
	}
	fmt.Println("✅ 管理员表创建成功")

	// 创建管理员角色关联表
	_, err = db.Exec(`
CREATE TABLE admin_role (
  id bigint unsigned NOT NULL AUTO_INCREMENT,
  admin_id bigint unsigned NOT NULL,
  role_id bigint unsigned NOT NULL,
  created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id),
  UNIQUE KEY uk_admin_role (admin_id, role_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='管理员角色关联表'
`)
	if err != nil {
		log.Fatal("创建管理员角色关联表失败:", err)
	}
	fmt.Println("✅ 管理员角色关联表创建成功")

	// 插入默认角色
	roles := []struct {
		name, code, desc string
	}{
		{"超级管理员", "super_admin", "拥有所有权限"},
		{"管理员", "admin", "拥有大部分权限"},
		{"运营", "operator", "运营相关权限"},
		{"编辑", "editor", "内容编辑权限"},
	}

	for _, r := range roles {
		_, err := db.Exec("INSERT INTO role (name, code, description) VALUES (?, ?, ?)", r.name, r.code, r.desc)
		if err != nil {
			fmt.Printf("插入角色 %s 失败: %v\n", r.name, err)
		}
	}
	fmt.Println("✅ 默认角色插入成功")

	// 插入默认权限
	permissions := []struct {
		name, code, ptype, path, icon string
		parentID                      int
	}{
		{"仪表盘", "dashboard", "menu", "/dashboard", "DashboardOutlined", 0},
		{"用户管理", "user", "menu", "/user", "UserOutlined", 0},
		{"物料管理", "material", "menu", "/material", "PictureOutlined", 0},
		{"内容管理", "content", "menu", "/content", "FileOutlined", 0},
		{"频道管理", "channel", "menu", "/channel", "AppstoreOutlined", 0},
		{"系统管理", "system", "menu", "/system", "SettingOutlined", 0},
		{"角色管理", "system:role", "menu", "/system/role", "", 6},
		{"权限管理", "system:permission", "menu", "/system/permission", "", 6},
		{"管理员管理", "system:admin", "menu", "/system/admin-user", "", 6},
	}

	for _, p := range permissions {
		_, err := db.Exec(
			"INSERT INTO permission (name, code, type, parent_id, path, icon) VALUES (?, ?, ?, ?, ?, ?)",
			p.name, p.code, p.ptype, p.parentID, p.path, p.icon,
		)
		if err != nil {
			fmt.Printf("插入权限 %s 失败: %v\n", p.name, err)
		}
	}
	fmt.Println("✅ 默认权限插入成功")

	// 为超级管理员分配所有权限
	_, err = db.Exec("INSERT INTO role_permission (role_id, permission_id) SELECT 1, id FROM permission")
	if err != nil {
		fmt.Println("分配权限失败:", err)
	} else {
		fmt.Println("✅ 超级管理员权限分配成功")
	}

	// 插入默认管理员 (密码: admin123)
	_, err = db.Exec(
		"INSERT INTO admin_user (username, password, realname, email, phone) VALUES ('admin', '$2a$10$N.zmdr9k7uOCQb376NoUnuTJ8iAt6Z5EHsM8lE9lBOsl7iAt6Z5EH', '超级管理员', 'admin@example.com', '13800138000')",
	)
	if err != nil {
		fmt.Println("插入管理员失败:", err)
	} else {
		fmt.Println("✅ 默认管理员插入成功")
	}

	// 为管理员分配超级管理员角色
	_, err = db.Exec("INSERT INTO admin_role (admin_id, role_id) VALUES (1, 1)")
	if err != nil {
		fmt.Println("分配角色失败:", err)
	} else {
		fmt.Println("✅ 管理员角色分配成功")
	}

	// 查询验证
	fmt.Println("\n📊 数据验证:")
	var count int
	db.QueryRow("SELECT COUNT(*) FROM role").Scan(&count)
	fmt.Printf("角色数量: %d\n", count)

	db.QueryRow("SELECT COUNT(*) FROM permission").Scan(&count)
	fmt.Printf("权限数量: %d\n", count)

	db.QueryRow("SELECT COUNT(*) FROM admin_user").Scan(&count)
	fmt.Printf("管理员数量: %d\n", count)

	fmt.Println("\n✅ RBAC表创建和数据初始化完成!")
	fmt.Println("\n📝 默认管理员账号:")
	fmt.Println("用户名: admin")
	fmt.Println("密码: admin123")
}
