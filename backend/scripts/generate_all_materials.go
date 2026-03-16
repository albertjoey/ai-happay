package main

import (
	"fmt"
	"math/rand"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Material struct {
	ID            uint      `gorm:"primaryKey"`
	TenantID      uint      `gorm:"column:tenant_id"`
	Title         string    `gorm:"column:title"`
	Subtitle      string    `gorm:"column:subtitle"`
	Type          string    `gorm:"column:type"`
	CoverURL      string    `gorm:"column:cover_url"`
	ContentURL    string    `gorm:"column:content_url"`
	Description   string    `gorm:"column:description"`
	Author        string    `gorm:"column:author"`
	Category      string    `gorm:"column:category"`
	ViewCount     int       `gorm:"column:view_count"`
	LikeCount     int       `gorm:"column:like_count"`
	CommentCount  int       `gorm:"column:comment_count"`
	ShareCount    int       `gorm:"column:share_count"`
	CollectCount  int       `gorm:"column:collect_count"`
	Duration      int       `gorm:"column:duration"`
	WordCount     int       `gorm:"column:word_count"`
	ChapterCount  int       `gorm:"column:chapter_count"`
	Status        int       `gorm:"column:status"`
	Sort          int       `gorm:"column:sort"`
	CreatedAt     time.Time `gorm:"column:created_at"`
	UpdatedAt     time.Time `gorm:"column:updated_at"`
}

func main() {
	dsn := "root:root123456@tcp(127.0.0.1:3306)/happy?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("连接数据库失败: " + err.Error())
	}

	rand.Seed(time.Now().UnixNano())
	now := time.Now()

	// 清空现有数据
	db.Exec("DELETE FROM material")
	db.Exec("ALTER TABLE material AUTO_INCREMENT = 1")

	// 生成数据
	totalCount := 0

	// 1. 图文内容 (200+条)
	imageTextCount := generateImageText(db, now)
	totalCount += imageTextCount
	fmt.Printf("✅ 生成图文内容: %d 条\n", imageTextCount)

	// 2. 短视频内容 (200+条)
	videoCount := generateVideo(db, now)
	totalCount += videoCount
	fmt.Printf("✅ 生成短视频内容: %d 条\n", videoCount)

	// 3. 小说内容 (200+条)
	novelCount := generateNovel(db, now)
	totalCount += novelCount
	fmt.Printf("✅ 生成小说内容: %d 条\n", novelCount)

	// 4. 漫画内容 (200+条)
	comicCount := generateComic(db, now)
	totalCount += comicCount
	fmt.Printf("✅ 生成漫画内容: %d 条\n", comicCount)

	// 5. 漫剧内容 (200+条)
	manhuaCount := generateManhua(db, now)
	totalCount += manhuaCount
	fmt.Printf("✅ 生成漫剧内容: %d 条\n", manhuaCount)

	// 6. 短剧内容 (200+条)
	shortDramaCount := generateShortDrama(db, now)
	totalCount += shortDramaCount
	fmt.Printf("✅ 生成短剧内容: %d 条\n", shortDramaCount)

	fmt.Printf("\n🎉 总计生成物料数据: %d 条\n", totalCount)
}

func generateImageText(db *gorm.DB, now time.Time) int {
	titles := []string{
		"美食制作教程", "旅行攻略分享", "健身训练计划", "美妆护肤技巧",
		"家居装修指南", "摄影技巧分享", "职场提升方法", "学习方法总结",
		"穿搭搭配技巧", "美食探店记录", "旅行日记分享", "健身打卡记录",
	}

	categories := []string{"美食", "旅行", "健身", "美妆", "家居", "摄影", "职场", "学习", "穿搭", "生活"}
	authors := []string{"美食达人", "旅行博主", "健身教练", "美妆博主", "家居达人", "摄影师", "职场导师", "学习博主", "穿搭达人", "生活家"}

	materials := make([]Material, 210)
	for i := 0; i < 210; i++ {
		title := fmt.Sprintf("%s：第%d期精彩内容", titles[i%len(titles)], i+1)
		materials[i] = Material{
			TenantID:     1,
			Title:        title,
			Subtitle:     fmt.Sprintf("专业分享，值得收藏"),
			Type:         "image_text",
			CoverURL:     fmt.Sprintf("https://picsum.photos/seed/it%d/400/600", i+1),
			ContentURL:   "",
			Description:  fmt.Sprintf("这是一篇关于%s的精彩内容，包含详细的步骤和实用的技巧，帮助你快速提升！", categories[i%len(categories)]),
			Author:       authors[i%len(authors)],
			Category:     categories[i%len(categories)],
			ViewCount:    randInt(10000, 500000),
			LikeCount:    randInt(1000, 50000),
			CommentCount: randInt(100, 5000),
			ShareCount:   randInt(50, 2000),
			CollectCount: randInt(500, 30000),
			WordCount:    randInt(1500, 5000),
			Status:       1,
			Sort:         100 - (i % 10),
			CreatedAt:    now,
			UpdatedAt:    now,
		}
	}

	db.Table("material").Create(&materials)
	return len(materials)
}

func generateVideo(db *gorm.DB, now time.Time) int {
	titles := []string{
		"搞笑视频合集", "美食制作视频", "旅行Vlog", "健身教学视频",
		"美妆教程视频", "游戏实况视频", "音乐翻唱视频", "萌宠日常视频",
	}

	categories := []string{"搞笑", "美食", "旅行", "健身", "美妆", "游戏", "音乐", "萌宠"}
	authors := []string{"搞笑达人", "美食博主", "旅行博主", "健身教练", "美妆博主", "游戏主播", "音乐人", "萌宠博主"}

	materials := make([]Material, 210)
	for i := 0; i < 210; i++ {
		title := fmt.Sprintf("%s：第%d集精彩呈现", titles[i%len(titles)], i+1)
		materials[i] = Material{
			TenantID:     1,
			Title:        title,
			Subtitle:     fmt.Sprintf("精彩视频，不容错过"),
			Type:         "video",
			CoverURL:     fmt.Sprintf("https://picsum.photos/seed/v%d/400/600", i+1),
			ContentURL:   fmt.Sprintf("https://example.com/video/v%d.mp4", i+1),
			Description:  fmt.Sprintf("这是一个关于%s的精彩视频，内容丰富有趣，快来看看吧！", categories[i%len(categories)]),
			Author:       authors[i%len(authors)],
			Category:     categories[i%len(categories)],
			ViewCount:    randInt(50000, 2000000),
			LikeCount:    randInt(5000, 200000),
			CommentCount: randInt(500, 20000),
			ShareCount:   randInt(100, 10000),
			CollectCount: randInt(1000, 100000),
			Duration:     randInt(60, 1800),
			Status:       1,
			Sort:         100 - (i % 10),
			CreatedAt:    now,
			UpdatedAt:    now,
		}
	}

	db.Create(&materials)
	return len(materials)
}

func generateNovel(db *gorm.DB, now time.Time) int {
	titles := []string{
		"重生之商业帝国", "仙武帝尊", "霸道总裁的小娇妻", "末世重生：我有一个空间",
		"都市修仙传", "穿越之古代奇缘", "玄幻世界冒险", "都市异能传说",
	}

	categories := []string{"都市重生", "玄幻修仙", "现代言情", "末世科幻", "都市修仙", "古装穿越", "玄幻热血", "都市异能"}
	authors := []string{"商战大神", "玄幻至尊", "甜文小公主", "末世作者", "修仙作者", "穿越作家", "玄幻作家", "异能作者"}

	materials := make([]Material, 210)
	for i := 0; i < 210; i++ {
		title := fmt.Sprintf("%s（第%d部）", titles[i%len(titles)], i+1)
		materials[i] = Material{
			TenantID:     1,
			Title:        title,
			Subtitle:     fmt.Sprintf("精彩小说，值得追读"),
			Type:         "novel",
			CoverURL:     fmt.Sprintf("https://picsum.photos/seed/n%d/400/600", i+1),
			ContentURL:   "",
			Description:  fmt.Sprintf("这是一部关于%s的精彩小说，情节跌宕起伏，人物形象鲜明，绝对让你欲罢不能！", categories[i%len(categories)]),
			Author:       authors[i%len(authors)],
			Category:     categories[i%len(categories)],
			ViewCount:    randInt(100000, 5000000),
			LikeCount:    randInt(10000, 500000),
			CommentCount: randInt(1000, 100000),
			ShareCount:   randInt(500, 50000),
			CollectCount: randInt(10000, 500000),
			WordCount:    randInt(1000000, 6000000),
			ChapterCount: randInt(500, 3000),
			Status:       1,
			Sort:         100 - (i % 10),
			CreatedAt:    now,
			UpdatedAt:    now,
		}
	}

	db.Create(&materials)
	return len(materials)
}

func generateComic(db *gorm.DB, now time.Time) int {
	titles := []string{
		"斗罗大陆", "一人之下", "狐妖小红娘", "武动乾坤",
		"大主宰", "完美世界", "绝世唐门", "龙王传说",
	}

	categories := []string{"玄幻热血", "都市异能", "爱情玄幻", "玄幻热血", "玄幻热血", "玄幻热血", "玄幻热血", "玄幻热血"}
	authors := []string{"唐家三少", "米二", "小新", "天蚕土豆", "天蚕土豆", "辰东", "唐家三少", "唐家三少"}

	materials := make([]Material, 210)
	for i := 0; i < 210; i++ {
		title := fmt.Sprintf("%s（第%d季）", titles[i%len(titles)], i+1)
		materials[i] = Material{
			TenantID:     1,
			Title:        title,
			Subtitle:     fmt.Sprintf("精彩漫画，热血呈现"),
			Type:         "comic",
			CoverURL:     fmt.Sprintf("https://picsum.photos/seed/c%d/400/600", i+1),
			ContentURL:   "",
			Description:  fmt.Sprintf("这是一部关于%s的精彩漫画，画风精美，剧情热血，绝对让你爱不释手！", categories[i%len(categories)]),
			Author:       authors[i%len(authors)],
			Category:     categories[i%len(categories)],
			ViewCount:    randInt(500000, 5000000),
			LikeCount:    randInt(50000, 500000),
			CommentCount: randInt(5000, 100000),
			ShareCount:   randInt(2000, 50000),
			CollectCount: randInt(50000, 500000),
			ChapterCount: randInt(300, 1000),
			Status:       1,
			Sort:         100 - (i % 10),
			CreatedAt:    now,
			UpdatedAt:    now,
		}
	}

	db.Create(&materials)
	return len(materials)
}

func generateManhua(db *gorm.DB, now time.Time) int {
	titles := []string{
		"霸道总裁爱上我", "重生之复仇女王", "穿越之古代奇缘", "都市甜宠故事",
		"校园恋爱物语", "玄幻修仙传", "都市异能者", "古代宫廷剧",
	}

	categories := []string{"都市甜宠", "都市复仇", "古装穿越", "都市甜宠", "校园恋爱", "玄幻修仙", "都市异能", "古装宫廷"}
	authors := []string{"甜宠剧场", "爽剧工厂", "古装剧场", "甜宠剧场", "校园剧场", "玄幻剧场", "异能剧场", "宫廷剧场"}

	materials := make([]Material, 210)
	for i := 0; i < 210; i++ {
		title := fmt.Sprintf("%s（第%d集）", titles[i%len(titles)], i+1)
		materials[i] = Material{
			TenantID:     1,
			Title:        title,
			Subtitle:     fmt.Sprintf("精彩漫剧，不容错过"),
			Type:         "manhua",
			CoverURL:     fmt.Sprintf("https://picsum.photos/seed/m%d/400/600", i+1),
			ContentURL:   "",
			Description:  fmt.Sprintf("这是一部关于%s的精彩漫剧，剧情精彩，人物生动，绝对让你欲罢不能！", categories[i%len(categories)]),
			Author:       authors[i%len(authors)],
			Category:     categories[i%len(categories)],
			ViewCount:    randInt(200000, 2000000),
			LikeCount:    randInt(20000, 200000),
			CommentCount: randInt(2000, 50000),
			ShareCount:   randInt(500, 20000),
			CollectCount: randInt(20000, 200000),
			ChapterCount: randInt(50, 300),
			Status:       1,
			Sort:         100 - (i % 10),
			CreatedAt:    now,
			UpdatedAt:    now,
		}
	}

	db.Create(&materials)
	return len(materials)
}

func generateShortDrama(db *gorm.DB, now time.Time) int {
	titles := []string{
		"逆袭人生", "都市情缘", "悬疑迷局", "霸道总裁爱上我",
		"重生之复仇女王", "穿越之古代奇缘", "都市甜宠故事", "校园恋爱物语",
	}

	categories := []string{"励志", "爱情", "悬疑", "甜宠", "复仇", "穿越", "甜宠", "校园"}
	authors := []string{"导演A", "导演B", "导演C", "甜宠剧场", "爽剧工厂", "古装剧场", "甜宠剧场", "校园剧场"}

	materials := make([]Material, 210)
	for i := 0; i < 210; i++ {
		title := fmt.Sprintf("%s（第%d集）", titles[i%len(titles)], i+1)
		materials[i] = Material{
			TenantID:     1,
			Title:        title,
			Subtitle:     fmt.Sprintf("精彩短剧，值得追看"),
			Type:         "short_drama",
			CoverURL:     fmt.Sprintf("https://picsum.photos/seed/sd%d/400/600", i+1),
			ContentURL:   fmt.Sprintf("https://example.com/drama/sd%d.mp4", i+1),
			Description:  fmt.Sprintf("这是一部关于%s的精彩短剧，剧情紧凑，演技在线，绝对让你看得停不下来！", categories[i%len(categories)]),
			Author:       authors[i%len(authors)],
			Category:     categories[i%len(categories)],
			ViewCount:    randInt(300000, 3000000),
			LikeCount:    randInt(30000, 300000),
			CommentCount: randInt(3000, 60000),
			ShareCount:   randInt(1000, 30000),
			CollectCount: randInt(30000, 300000),
			Duration:     randInt(60, 300),
			Status:       1,
			Sort:         100 - (i % 10),
			CreatedAt:    now,
			UpdatedAt:    now,
		}
	}

	db.Create(&materials)
	return len(materials)
}

func randInt(min, max int) int {
	return rand.Intn(max-min) + min
}
