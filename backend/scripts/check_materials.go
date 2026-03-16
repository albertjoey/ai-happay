package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:root123456@tcp(127.0.0.1:3306)/happy?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("连接数据库失败: " + err.Error())
	}

	var total int64
	db.Table("material").Where("deleted_at IS NULL").Count(&total)
	fmt.Printf("物料总数: %d\n", total)

	type Material struct {
		ID    uint
		Type  string
		Title string
	}

	var materials []Material
	db.Table("material").Where("deleted_at IS NULL").Order("id DESC").Limit(10).Find(&materials)

	fmt.Println("\n最新10条物料:")
	for _, m := range materials {
		fmt.Printf("%d: %s - %s\n", m.ID, m.Type, m.Title)
	}
}
