package main

import (
	"encoding/json"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

func main() {
	// 连接数据库
	dsn := "root:happy123456@tcp(127.0.0.1:3306)/happy?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("连接数据库失败: " + err.Error())
	}

	now := time.Now()

	// 为每个频道添加配置
	channels := []int{2, 3, 4, 5, 6} // 搞笑、热门、颜值、动漫、社区

	for _, channelID := range channels {
		fmt.Printf("\n=== 为频道 %d 添加配置 ===\n", channelID)

		// 1. 添加金刚位 (5个一组)
		diamonds := []map[string]interface{}{
			{
				"tenant_id":   1,
				"channel_id":  channelID,
				"group_id":    1,
				"sort":        1,
				"title":       "热门话题",
				"icon":        "🔥",
				"link_type":   "topic",
				"link_value":  "1",
				"status":      1,
				"description": "热门话题入口",
				"created_at":  now,
				"updated_at":  now,
			},
			{
				"tenant_id":   1,
				"channel_id":  channelID,
				"group_id":    1,
				"sort":        2,
				"title":       "精选内容",
				"icon":        "⭐",
				"link_type":   "content",
				"link_value":  "1",
				"status":      1,
				"description": "精选内容入口",
				"created_at":  now,
				"updated_at":  now,
			},
			{
				"tenant_id":   1,
				"channel_id":  channelID,
				"group_id":    1,
				"sort":        3,
				"title":       "排行榜",
				"icon":        "📊",
				"link_type":   "rank",
				"link_value":  "1",
				"status":      1,
				"description": "排行榜入口",
				"created_at":  now,
				"updated_at":  now,
			},
			{
				"tenant_id":   1,
				"channel_id":  channelID,
				"group_id":    1,
				"sort":        4,
				"title":       "活动中心",
				"icon":        "🎁",
				"link_type":   "activity",
				"link_value":  "1",
				"status":      1,
				"description": "活动中心入口",
				"created_at":  now,
				"updated_at":  now,
			},
			{
				"tenant_id":   1,
				"channel_id":  channelID,
				"group_id":    1,
				"sort":        5,
				"title":       "更多",
				"icon":        "➕",
				"link_type":   "more",
				"link_value":  "",
				"status":      1,
				"description": "更多入口",
				"created_at":  now,
				"updated_at":  now,
			},
		}

		for _, diamond := range diamonds {
			result := db.Table("diamond").Create(&diamond)
			if result.Error != nil {
				fmt.Printf("添加金刚位失败: %v\n", result.Error)
			}
		}
		fmt.Printf("✅ 添加了 %d 个金刚位\n", len(diamonds))

		// 2. 添加推荐位 (3个不同展示类型)
		contentIDs, _ := json.Marshal([]int{1, 2, 3, 4, 5, 6})
		filterRule, _ := json.Marshal(map[string]interface{}{"limit": 6, "type": "all"})

		recommends := []map[string]interface{}{
			{
				"tenant_id":    1,
				"channel_id":   channelID,
				"title":        "今日推荐",
				"display_type": "single",
				"source_type":  "algorithm",
				"content_ids":  contentIDs,
				"filter_rule":  filterRule,
				"sort":         1,
				"status":       1,
				"description":  "今日推荐内容",
				"created_at":   now,
				"updated_at":   now,
			},
			{
				"tenant_id":    1,
				"channel_id":   channelID,
				"title":        "热门精选",
				"display_type": "scroll",
				"source_type":  "filter",
				"content_ids":  contentIDs,
				"filter_rule":  filterRule,
				"sort":         2,
				"status":       1,
				"description":  "热门精选内容",
				"created_at":   now,
				"updated_at":   now,
			},
			{
				"tenant_id":    1,
				"channel_id":   channelID,
				"title":        "编辑推荐",
				"display_type": "grid",
				"source_type":  "manual",
				"content_ids":  contentIDs,
				"filter_rule":  filterRule,
				"sort":         3,
				"status":       1,
				"description":  "编辑推荐内容",
				"created_at":   now,
				"updated_at":   now,
			},
		}

		for _, recommend := range recommends {
			result := db.Table("recommend").Create(&recommend)
			if result.Error != nil {
				fmt.Printf("添加推荐位失败: %v\n", result.Error)
			}
		}
		fmt.Printf("✅ 添加了 %d 个推荐位\n", len(recommends))

		// 3. 添加Feed流配置
		feedFilterRule, _ := json.Marshal(map[string]interface{}{"limit": 20, "type": "all"})

		feedConfigs := []map[string]interface{}{
			{
				"tenant_id":        1,
				"channel_id":       channelID,
				"title":            "推荐Feed",
				"layout_type":      "two_col",
				"content_strategy": "algorithm",
				"content_ids":      contentIDs,
				"filter_rule":      feedFilterRule,
				"sort":             1,
				"status":           1,
				"description":      "推荐Feed流",
				"created_at":       now,
				"updated_at":       now,
			},
		}

		for _, feed := range feedConfigs {
			result := db.Table("feed_config").Create(&feed)
			if result.Error != nil {
				fmt.Printf("添加Feed配置失败: %v\n", result.Error)
			}
		}
		fmt.Printf("✅ 添加了 %d 个Feed配置\n", len(feedConfigs))

		// 4. 添加广告位 (2个不同插入方式)
		adContent1, _ := json.Marshal(map[string]interface{}{
			"image_url": fmt.Sprintf("https://picsum.photos/400/600?ad=%d-1", channelID),
			"title":     "精彩广告",
		})
		adContent2, _ := json.Marshal(map[string]interface{}{
			"image_url": fmt.Sprintf("https://picsum.photos/400/600?ad=%d-2", channelID),
			"title":     "推荐广告",
		})

		insertRule1, _ := json.Marshal(map[string]interface{}{"fixed_position": 3, "max_count": 1})
		insertRule2, _ := json.Marshal(map[string]interface{}{"interval": 5, "max_count": 2})

		adSlots := []map[string]interface{}{
			{
				"tenant_id":   1,
				"channel_id":  channelID,
				"name":        fmt.Sprintf("频道%d固定广告", channelID),
				"insert_type": "fixed",
				"insert_rule": insertRule1,
				"ad_type":     "image",
				"ad_content":  adContent1,
				"link_url":    fmt.Sprintf("https://example.com/ad/%d-1", channelID),
				"status":      1,
				"sort":        1,
				"description": "第3条内容后插入",
				"created_at":  now,
				"updated_at":  now,
			},
			{
				"tenant_id":   1,
				"channel_id":  channelID,
				"name":        fmt.Sprintf("频道%d间隔广告", channelID),
				"insert_type": "interval",
				"insert_rule": insertRule2,
				"ad_type":     "image",
				"ad_content":  adContent2,
				"link_url":    fmt.Sprintf("https://example.com/ad/%d-2", channelID),
				"status":      1,
				"sort":        2,
				"description": "每5条插入一个",
				"created_at":  now,
				"updated_at":  now,
			},
		}

		for _, ad := range adSlots {
			result := db.Table("ad_slot").Create(&ad)
			if result.Error != nil {
				fmt.Printf("添加广告位失败: %v\n", result.Error)
			}
		}
		fmt.Printf("✅ 添加了 %d 个广告位\n", len(adSlots))
	}

	fmt.Println("\n=== 所有频道配置添加完成 ===")
}
