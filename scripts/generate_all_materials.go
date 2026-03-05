package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// 物料类型
	types := []string{"image_text", "short_video", "long_video", "short_drama", "drama", "novel"}
	typeNames := map[string]string{
		"image_text":   "图文",
		"short_video":  "短视频",
		"long_video":   "长视频",
		"short_drama":  "短剧",
		"drama":        "漫剧",
		"novel":        "小说",
	}
	categories := []string{"搞笑", "美食", "旅行", "音乐", "舞蹈", "游戏", "宠物", "健身", "科技", "时尚"}

	// 每种类型生成的数量
	counts := map[string]int{
		"image_text":   50,
		"short_video":  50,
		"long_video":   30,
		"short_drama":  20,
		"drama":        30,
		"novel":        20,
	}

	fmt.Println("-- 生成6种物料类型的测试数据")
	fmt.Println("-- 生成时间:", time.Now().Format("2006-01-02 15:04:05"))
	fmt.Println()

	for _, t := range types {
		fmt.Printf("-- %s (%s)\n", typeNames[t], t)
		for i := 1; i <= counts[t]; i++ {
			title := fmt.Sprintf("%s作品%03d", typeNames[t], i)
			subtitle := fmt.Sprintf("精彩%s内容", typeNames[t])
			category := categories[rand.Intn(len(categories))]
			author := fmt.Sprintf("创作者%03d号", rand.Intn(100)+1)
			viewCount := rand.Intn(500000) + 10000
			likeCount := rand.Intn(50000) + 1000
			commentCount := rand.Intn(5000) + 100
			shareCount := rand.Intn(5000) + 100
			collectCount := rand.Intn(20000) + 1000
			chapterCount := 0
			wordCount := 0
			duration := 0

			switch t {
			case "novel":
				wordCount = rand.Intn(500000) + 100000
				chapterCount = rand.Intn(300) + 50
			case "drama", "short_drama":
				chapterCount = rand.Intn(100) + 20
			case "long_video":
				duration = rand.Intn(7200) + 1800 // 30分钟-2小时
			case "short_video":
				duration = rand.Intn(60) + 15 // 15-75秒
			}

			fmt.Printf(`INSERT INTO material (tenant_id, title, subtitle, type