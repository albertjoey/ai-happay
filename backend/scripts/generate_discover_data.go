package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	dbHost     = "localhost"
	dbPort     = 3306
	dbUser     = "root"
	dbPassword = "happy123456"
	dbName     = "happy"
)

var db *sql.DB

func main() {
	// 连接数据库
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPassword, dbHost, dbPort, dbName)

	var err error
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("连接数据库失败:", err)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Fatal("数据库连接失败:", err)
	}

	fmt.Println("✅ 数据库连接成功")

	// 生成发现页测试数据
	generateHotTopics()         // 热门话题
	generateHotRank()           // 热门榜单
	generateRecommendCreators() // 推荐创作者
	generateGuessYouLike()      // 猜你喜欢

	fmt.Println("\n🎉 发现页测试数据生成完成!")
}

// 生成热门话题
func generateHotTopics() {
	fmt.Println("\n🔥 生成热门话题...")

	topics := []map[string]interface{}{
		{"name": "周末美食打卡", "count": "12.5万", "color": "#FF6B6B"},
		{"name": "旅行日记", "count": "8.3万", "color": "#4ECDC4"},
		{"name": "宠物萌照", "count": "6.7万", "color": "#45B7D1"},
		{"name": "健身打卡", "count": "5.2万", "color": "#96CEB4"},
		{"name": "穿搭分享", "count": "4.8万", "color": "#FFEAA7"},
		{"name": "游戏日常", "count": "3.9万", "color": "#DDA0DD"},
	}

	for i, topic := range topics {
		extraData, _ := json.Marshal(map[string]interface{}{
			"count": topic["count"],
			"color": topic["color"],
		})

		_, err := db.Exec(`
			INSERT INTO discover_item (tenant_id, module, item_type, item_id, title, cover_url, extra_data, sort_order, is_enabled, created_at, updated_at)
			VALUES (1, 'hot_topics', 'topic', ?, ?, '', ?, ?, 1, NOW(), NOW())
		`, i+1, topic["name"], string(extraData), i+1)

		if err != nil {
			log.Printf("插入热门话题失败: %v", err)
		}
	}

	fmt.Println("✅ 完成 6 个热门话题")
}

// 生成热门榜单
func generateHotRank() {
	fmt.Println("\n🏆 生成热门榜单...")

	items := []map[string]interface{}{
		{"title": "今日份美食分享", "author": "美食达人小王", "views": "128万", "color": "#FF6B6B"},
		{"title": "周末旅行vlog", "author": "旅行者小李", "views": "96万", "color": "#4ECDC4"},
		{"title": "我家猫咪的一天", "author": "猫奴小张", "views": "85万", "color": "#45B7D1"},
		{"title": "健身30天挑战", "author": "健身教练", "views": "72万", "color": "#96CEB4"},
		{"title": "春季穿搭指南", "author": "时尚博主", "views": "68万", "color": "#FFEAA7"},
	}

	for i, item := range items {
		extraData, _ := json.Marshal(map[string]interface{}{
			"author": item["author"],
			"views":  item["views"],
			"color":  item["color"],
			"rank":   i + 1,
		})

		_, err := db.Exec(`
			INSERT INTO discover_item (tenant_id, module, item_type, item_id, title, cover_url, extra_data, sort_order, is_enabled, created_at, updated_at)
			VALUES (1, 'hot_rank', 'content', ?, ?, '', ?, ?, 1, NOW(), NOW())
		`, i+1, item["title"], string(extraData), i+1)

		if err != nil {
			log.Printf("插入热门榜单失败: %v", err)
		}
	}

	fmt.Println("✅ 完成 5 个热门榜单")
}

// 生成推荐创作者
func generateRecommendCreators() {
	fmt.Println("\n⭐ 生成推荐创作者...")

	creators := []map[string]interface{}{
		{"name": "美食达人小王", "avatar": "👨‍🍳", "fans": "52.3万", "desc": "分享美食日常", "color": "#FF6B6B"},
		{"name": "旅行者小李", "avatar": "🌍", "fans": "38.7万", "desc": "环游世界ing", "color": "#4ECDC4"},
		{"name": "猫奴小张", "avatar": "🐱", "fans": "29.1万", "desc": "两只猫的铲屎官", "color": "#45B7D1"},
		{"name": "健身教练", "avatar": "💪", "fans": "45.6万", "desc": "专业健身指导", "color": "#96CEB4"},
	}

	for i, creator := range creators {
		extraData, _ := json.Marshal(map[string]interface{}{
			"avatar": creator["avatar"],
			"fans":   creator["fans"],
			"desc":   creator["desc"],
			"color":  creator["color"],
		})

		_, err := db.Exec(`
			INSERT INTO discover_item (tenant_id, module, item_type, item_id, title, cover_url, extra_data, sort_order, is_enabled, created_at, updated_at)
			VALUES (1, 'recommend_creators', 'creator', ?, ?, '', ?, ?, 1, NOW(), NOW())
		`, i+1, creator["name"], string(extraData), i+1)

		if err != nil {
			log.Printf("插入推荐创作者失败: %v", err)
		}
	}

	fmt.Println("✅ 完成 4 个推荐创作者")
}

// 生成猜你喜欢
func generateGuessYouLike() {
	fmt.Println("\n💡 生成猜你喜欢...")

	items := []map[string]interface{}{
		{"title": "超简单的家常菜做法", "author": "美食达人", "views": "23万", "color": "#FF6B6B"},
		{"title": "一个人的旅行", "author": "旅行者", "views": "18万", "color": "#4ECDC4"},
		{"title": "猫咪搞笑瞬间", "author": "猫奴", "views": "45万", "color": "#45B7D1"},
		{"title": "居家健身教程", "author": "健身教练", "views": "32万", "color": "#96CEB4"},
		{"title": "春季穿搭推荐", "author": "时尚博主", "views": "28万", "color": "#FFEAA7"},
		{"title": "游戏精彩操作", "author": "游戏玩家", "views": "56万", "color": "#DDA0DD"},
	}

	for i, item := range items {
		extraData, _ := json.Marshal(map[string]interface{}{
			"author": item["author"],
			"views":  item["views"],
			"color":  item["color"],
		})

		_, err := db.Exec(`
			INSERT INTO discover_item (tenant_id, module, item_type, item_id, title, cover_url, extra_data, sort_order, is_enabled, created_at, updated_at)
			VALUES (1, 'guess_you_like', 'content', ?, ?, '', ?, ?, 1, NOW(), NOW())
		`, i+1, item["title"], string(extraData), i+1)

		if err != nil {
			log.Printf("插入猜你喜欢失败: %v", err)
		}
	}

	fmt.Println("✅ 完成 6 个猜你喜欢")
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
