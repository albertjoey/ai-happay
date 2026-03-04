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

	// 更新金刚位表
	_, err = db.Exec("ALTER TABLE diamond ADD COLUMN material_id bigint unsigned DEFAULT 0 COMMENT '关联物料ID' AFTER channel_id")
	if err != nil {
		fmt.Println("添加material_id字段(可能已存在):", err)
	} else {
		fmt.Println("✅ 金刚位表添加material_id字段成功")
	}

	// 更新Feed流配置表
	_, err = db.Exec("ALTER TABLE feed_config ADD COLUMN material_ids json DEFAULT NULL COMMENT '物料ID列表' AFTER filter_rule")
	if err != nil {
		fmt.Println("添加material_ids字段(可能已存在):", err)
	} else {
		fmt.Println("✅ Feed流配置表添加material_ids字段成功")
	}

	// 更新金刚位数据
	for i := 1; i <= 6; i++ {
		_, err := db.Exec("UPDATE diamond SET material_id = ? WHERE id = ?", i, i)
		if err != nil {
			log.Printf("更新金刚位%d失败: %v", i, err)
		}
	}
	fmt.Println("✅ 金刚位数据更新成功")

	// 更新推荐位数据
	_, err = db.Exec("UPDATE recommend SET content_ids = '[1, 2, 3, 4, 5, 6]' WHERE id = 1")
	if err != nil {
		log.Printf("更新推荐位1失败: %v", err)
	}
	_, err = db.Exec("UPDATE recommend SET content_ids = '[7, 8, 9, 10, 11, 12]' WHERE id = 2")
	if err != nil {
		log.Printf("更新推荐位2失败: %v", err)
	}
	fmt.Println("✅ 推荐位数据更新成功")

	// 更新Feed流配置数据
	_, err = db.Exec("UPDATE feed_config SET material_ids = '[1, 2, 3, 4, 5, 6, 7, 8, 9, 10]' WHERE id = 1")
	if err != nil {
		log.Printf("更新Feed流配置1失败: %v", err)
	}
	_, err = db.Exec("UPDATE feed_config SET material_ids = '[11, 12, 13, 14, 15, 16, 17]' WHERE id = 2")
	if err != nil {
		log.Printf("更新Feed流配置2失败: %v", err)
	}
	fmt.Println("✅ Feed流配置数据更新成功")

	// 查询验证
	fmt.Println("\n📊 金刚位数据预览:")
	rows, err := db.Query("SELECT id, title, material_id FROM diamond LIMIT 5")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id, materialID int
		var title string
		rows.Scan(&id, &title, &materialID)
		fmt.Printf("ID: %d, 标题: %s, 物料ID: %d\n", id, title, materialID)
	}

	fmt.Println("\n✅ 所有更新完成!")
}
