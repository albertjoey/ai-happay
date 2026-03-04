package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:happy123456@tcp(127.0.0.1:3306)/happy?charset=utf8mb4")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, title, material_id FROM diamond WHERE id IN (1,2,3,4,5,6)")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	fmt.Println("金刚位数据:")
	for rows.Next() {
		var id, materialID int
		var title string
		rows.Scan(&id, &title, &materialID)
		fmt.Printf("ID: %d, 标题: %s, 物料ID: %d\n", id, title, materialID)
	}
}
