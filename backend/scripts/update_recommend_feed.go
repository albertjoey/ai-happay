package main

import (
	"encoding/json"
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

	// 获取不同类型的物料ID
	type MaterialID struct {
		ID   uint
		Type string
	}

	var videoIDs []uint
	var novelIDs []uint
	var imageTextIDs []uint
	var shortDramaIDs []uint
	var comicIDs []uint

	db.Table("material").Where("type = ? AND deleted_at IS NULL", "video").Pluck("id", &videoIDs)
	db.Table("material").Where("type = ? AND deleted_at IS NULL", "novel").Pluck("id", &novelIDs)
	db.Table("material").Where("type = ? AND deleted_at IS NULL", "image_text").Pluck("id", &imageTextIDs)
	db.Table("material").Where("type = ? AND deleted_at IS NULL", "short_drama").Pluck("id", &shortDramaIDs)
	db.Table("material").Where("type = ? AND deleted_at IS NULL", "comic").Pluck("id", &comicIDs)

	fmt.Printf("视频: %d个, 小说: %d个, 图文: %d个, 短剧: %d个, 漫剧: %d个\n", 
		len(videoIDs), len(novelIDs), len(imageTextIDs), len(shortDramaIDs), len(comicIDs))

	// 更新所有推荐位的content_ids
	allIDs := append(append(append(append(videoIDs, novelIDs...), imageTextIDs...), shortDramaIDs...), comicIDs...)
	contentIDsJSON, _ := json.Marshal(allIDs[:min(20, len(allIDs))]) // 取前20个

	result := db.Table("recommend").Where("1=1").Update("content_ids", contentIDsJSON)
	fmt.Printf("✅ 更新了 %d 个推荐位的content_ids\n", result.RowsAffected)

	// 更新所有Feed流的content_ids
	result = db.Table("feed_config").Where("1=1").Update("content_ids", contentIDsJSON)
	fmt.Printf("✅ 更新了 %d 个Feed流的content_ids\n", result.RowsAffected)

	// 为每个频道创建不同类型的推荐位
	channels := []int{1, 2, 3, 4, 5, 6}
	for _, channelID := range channels {
		// 删除旧的推荐位
		db.Table("recommend").Where("channel_id = ?", channelID).Delete(nil)

		// 创建新的推荐位
		recommends := []map[string]interface{}{
			{
				"tenant_id":    1,
				"channel_id":   channelID,
				"title":        "热门视频",
				"display_type": "scroll",
				"source_type":  "manual",
				"content_ids":  mustMarshal(videoIDs[:min(8, len(videoIDs))]),
				"sort":         1,
				"status":       1,
				"description":  "精选热门视频推荐",
			},
			{
				"tenant_id":    1,
				"channel_id":   channelID,
				"title":        "热门小说",
				"display_type": "grid",
				"source_type":  "manual",
				"content_ids":  mustMarshal(novelIDs[:min(5, len(novelIDs))]),
				"sort":         2,
				"status":       1,
				"description":  "精选热门小说推荐",
			},
			{
				"tenant_id":    1,
				"channel_id":   channelID,
				"title":        "精彩图文",
				"display_type": "single",
				"source_type":  "manual",
				"content_ids":  mustMarshal(imageTextIDs[:min(5, len(imageTextIDs))]),
				"sort":         3,
				"status":       1,
				"description":  "精选图文内容推荐",
			},
		}

		for _, rec := range recommends {
			db.Table("recommend").Create(&rec)
		}
	}
	fmt.Printf("✅ 为 %d 个频道创建了新的推荐位\n", len(channels))

	// 更新Feed流配置
	for _, channelID := range channels {
		db.Table("feed_config").Where("channel_id = ?", channelID).Update("content_ids", contentIDsJSON)
	}
	fmt.Printf("✅ 更新了 %d 个频道的Feed流配置\n", len(channels))
}

func mustMarshal(v interface{}) []byte {
	data, _ := json.Marshal(v)
	return data
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
