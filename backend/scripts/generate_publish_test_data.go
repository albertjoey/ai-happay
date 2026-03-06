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

// Material 物料表
type Material struct {
	ID           int64  `json:"id"`
	Type         string `json:"type"`
	Title        string `json:"title"`
	Subtitle     string `json:"subtitle"`
	Description  string `json:"description"`
	CoverURL     string `json:"cover_url"`
	ContentURL   string `json:"content_url"`
	Author       string `json:"author"`
	Category     string `json:"category"`
	Tags         string `json:"tags"`
	ViewCount    int    `json:"view_count"`
	LikeCount    int    `json:"like_count"`
	CommentCount int    `json:"comment_count"`
	ShareCount   int    `json:"share_count"`
	CollectCount int    `json:"collect_count"`
	Duration     int    `json:"duration"`
	WordCount    int    `json:"word_count"`
	ChapterCount int    `json:"chapter_count"`
	Status       int    `json:"status"`
	Sort         int    `json:"sort"`
}

// Episode 剧集
type Episode struct {
	EpisodeNumber int      `json:"episode_number"`
	Title         string   `json:"title"`
	CoverURL      string   `json:"cover_url"`
	VideoURL      string   `json:"video_url"`
	Images        []string `json:"images"`
	Duration      int      `json:"duration"`
}

// Chapter 章节
type Chapter struct {
	ChapterNumber int    `json:"chapter_number"`
	Title         string `json:"title"`
	Content       string `json:"content"`
	WordCount     int    `json:"word_count"`
}

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

	// 生成测试数据
	generateImageTexts(50)      // 图文
	generateShortVideos(50)     // 短视频
	generateLongVideos(30)      // 长视频
	generateShortDramas(20)     // 短剧
	generateDramas(30)          // 漫剧
	generateNovels(20)          // 小说

	fmt.Println("\n🎉 所有测试数据生成完成!")
}

// 生成图文内容
func generateImageTexts(count int) {
	fmt.Printf("\n📝 生成 %d 条图文内容...\n", count)

	titles := []string{
		"今日份的美食分享", "旅行日记：探索未知的美好", "生活小技巧get",
		"周末的惬意时光", "我的健身打卡", "读书笔记分享",
		"城市夜景随拍", "手工DIY教程", "宠物日常",
		"穿搭分享", "美食制作教程", "户外探险记录",
	}

	for i := 0; i < count; i++ {
		title := titles[rand.Intn(len(titles))]
		if count > len(titles) {
			title = fmt.Sprintf("%s #%d", title, i+1)
		}

		// 生成1-9张图片
		imageCount := rand.Intn(9) + 1
		images := make([]string, imageCount)
		for j := 0; j < imageCount; j++ {
			images[j] = fmt.Sprintf("https://picsum.photos/800/600?random=%d", rand.Intn(10000))
		}
		imagesJSON, _ := json.Marshal(images)

		material := Material{
			Type:         "image_text",
			Title:        title,
			Description:  fmt.Sprintf("这是第%d条图文内容，分享生活中的美好瞬间。", i+1),
			CoverURL:     images[0],
			ContentURL:   string(imagesJSON),
			Author:       fmt.Sprintf("用户%d", rand.Intn(100)+1),
			Category:     "图文",
			ViewCount:    rand.Intn(10000),
			LikeCount:    rand.Intn(1000),
			CommentCount: rand.Intn(100),
			ShareCount:   rand.Intn(50),
			CollectCount: rand.Intn(100),
			Status:       1,
			Sort:         i + 1,
		}

		if err := insertMaterial(&material); err != nil {
			log.Printf("插入图文失败: %v", err)
		}
	}

	fmt.Printf("✅ 完成 %d 条图文内容\n", count)
}

// 生成短视频
func generateShortVideos(count int) {
	fmt.Printf("\n📱 生成 %d 条短视频...\n", count)

	titles := []string{
		"今日份的vlog", "搞笑日常", "才艺展示",
		"美食制作", "旅行记录", "运动打卡",
		"宠物萌宠", "生活记录", "创意视频",
	}

	for i := 0; i < count; i++ {
		title := titles[rand.Intn(len(titles))]
		if count > len(titles) {
			title = fmt.Sprintf("%s #%d", title, i+1)
		}

		material := Material{
			Type:         "short_video",
			Title:        title,
			Description:  fmt.Sprintf("第%d个短视频，记录精彩瞬间。", i+1),
			CoverURL:     fmt.Sprintf("https://picsum.photos/720/1280?random=%d", rand.Intn(10000)),
			ContentURL:   "https://test-streams.mux.dev/x36xhzz/x36xhzz.m3u8",
			Author:       fmt.Sprintf("创作者%d", rand.Intn(100)+1),
			Category:     "短视频",
			ViewCount:    rand.Intn(50000),
			LikeCount:    rand.Intn(5000),
			CommentCount: rand.Intn(500),
			ShareCount:   rand.Intn(200),
			CollectCount: rand.Intn(300),
			Duration:     rand.Intn(60) + 10, // 10-70秒
			Status:       1,
			Sort:         i + 1,
		}

		if err := insertMaterial(&material); err != nil {
			log.Printf("插入短视频失败: %v", err)
		}
	}

	fmt.Printf("✅ 完成 %d 条短视频\n", count)
}

// 生成长视频
func generateLongVideos(count int) {
	fmt.Printf("\n🎬 生成 %d 条长视频...\n", count)

	titles := []string{
		"深度解析：科技发展趋势", "旅行纪录片", "美食探店",
		"历史故事讲解", "科学知识科普", "电影解说",
		"游戏实况", "音乐MV", "纪录片分享",
	}

	for i := 0; i < count; i++ {
		title := titles[rand.Intn(len(titles))]
		if count > len(titles) {
			title = fmt.Sprintf("%s #%d", title, i+1)
		}

		material := Material{
			Type:         "long_video",
			Title:        title,
			Description:  fmt.Sprintf("第%d个长视频，深度内容分享。", i+1),
			CoverURL:     fmt.Sprintf("https://picsum.photos/1280/720?random=%d", rand.Intn(10000)),
			ContentURL:   "https://test-streams.mux.dev/x36xhzz/x36xhzz.m3u8",
			Author:       fmt.Sprintf("UP主%d", rand.Intn(100)+1),
			Category:     "长视频",
			ViewCount:    rand.Intn(100000),
			LikeCount:    rand.Intn(10000),
			CommentCount: rand.Intn(1000),
			ShareCount:   rand.Intn(500),
			CollectCount: rand.Intn(800),
			Duration:     rand.Intn(3600) + 600, // 10-70分钟
			Status:       1,
			Sort:         i + 1,
		}

		if err := insertMaterial(&material); err != nil {
			log.Printf("插入长视频失败: %v", err)
		}
	}

	fmt.Printf("✅ 完成 %d 条长视频\n", count)
}

// 生成短剧
func generateShortDramas(count int) {
	fmt.Printf("\n🎭 生成 %d 部短剧...\n", count)

	titles := []string{
		"霸道总裁爱上我", "重生之商业帝国", "穿越时空的爱恋",
		"都市情缘", "豪门恩怨", "青春校园恋",
		"职场逆袭记", "神秘千金", "契约婚姻",
	}

	for i := 0; i < count; i++ {
		title := titles[rand.Intn(len(titles))]
		if count > len(titles) {
			title = fmt.Sprintf("%s #%d", title, i+1)
		}

		// 生成5-10集（减少集数以避免数据过长）
		episodeCount := rand.Intn(6) + 5
		episodes := make([]Episode, episodeCount)
		for j := 0; j < episodeCount; j++ {
			episodes[j] = Episode{
				EpisodeNumber: j + 1,
				Title:         fmt.Sprintf("第%d集", j+1),
				CoverURL:      fmt.Sprintf("https://picsum.photos/720/1280?random=%d", rand.Intn(10000)),
				VideoURL:      "https://test-streams.mux.dev/x36xhzz/x36xhzz.m3u8",
				Duration:      rand.Intn(180) + 60, // 1-4分钟
			}
		}
		episodesJSON, _ := json.Marshal(episodes)

		material := Material{
			Type:         "short_drama",
			Title:        title,
			Subtitle:     fmt.Sprintf("共%d集", episodeCount),
			Description:  fmt.Sprintf("精彩短剧《%s》，共%d集，每集2-3分钟。", title, episodeCount),
			CoverURL:     fmt.Sprintf("https://picsum.photos/720/1280?random=%d", rand.Intn(10000)),
			ContentURL:   string(episodesJSON),
			Author:       fmt.Sprintf("导演%d", rand.Intn(50)+1),
			Category:     "短剧",
			ViewCount:    rand.Intn(500000),
			LikeCount:    rand.Intn(50000),
			CommentCount: rand.Intn(5000),
			ShareCount:   rand.Intn(2000),
			CollectCount: rand.Intn(3000),
			ChapterCount: episodeCount,
			Status:       1,
			Sort:         i + 1,
		}

		if err := insertMaterial(&material); err != nil {
			log.Printf("插入短剧失败: %v", err)
		}
	}

	fmt.Printf("✅ 完成 %d 部短剧\n", count)
}

// 生成漫剧
func generateDramas(count int) {
	fmt.Printf("\n📚 生成 %d 部漫剧...\n", count)

	titles := []string{
		"绝世武神", "都市修仙", "霸道王爷",
		"神医嫡女", "重生女王", "天才宝宝",
		"豪门宠婚", "异能觉醒", "古风言情",
	}

	for i := 0; i < count; i++ {
		title := titles[rand.Intn(len(titles))]
		if count > len(titles) {
			title = fmt.Sprintf("%s #%d", title, i+1)
		}

		// 生成5-10话（减少话数以避免数据过长）
		episodeCount := rand.Intn(6) + 5
		episodes := make([]Episode, episodeCount)
		for j := 0; j < episodeCount; j++ {
			// 每话3-5张图片（减少图片数量）
			imageCount := rand.Intn(3) + 3
			images := make([]string, imageCount)
			for k := 0; k < imageCount; k++ {
				images[k] = fmt.Sprintf("https://picsum.photos/800/1200?random=%d", rand.Intn(10000))
			}

			episodes[j] = Episode{
				EpisodeNumber: j + 1,
				Title:         fmt.Sprintf("第%d话", j+1),
				CoverURL:      images[0],
				Images:        images,
			}
		}
		episodesJSON, _ := json.Marshal(episodes)

		material := Material{
			Type:         "drama",
			Title:        title,
			Subtitle:     fmt.Sprintf("共%d话", episodeCount),
			Description:  fmt.Sprintf("精彩漫剧《%s》，共%d话，条漫形式呈现。", title, episodeCount),
			CoverURL:     fmt.Sprintf("https://picsum.photos/720/960?random=%d", rand.Intn(10000)),
			ContentURL:   string(episodesJSON),
			Author:       fmt.Sprintf("画师%d", rand.Intn(50)+1),
			Category:     "漫剧",
			ViewCount:    rand.Intn(300000),
			LikeCount:    rand.Intn(30000),
			CommentCount: rand.Intn(3000),
			ShareCount:   rand.Intn(1000),
			CollectCount: rand.Intn(2000),
			ChapterCount: episodeCount,
			Status:       1,
			Sort:         i + 1,
		}

		if err := insertMaterial(&material); err != nil {
			log.Printf("插入漫剧失败: %v", err)
		}
	}

	fmt.Printf("✅ 完成 %d 部漫剧\n", count)
}

// 生成小说
func generateNovels(count int) {
	fmt.Printf("\n📖 生成 %d 部小说...\n", count)

	titles := []string{
		"修仙之路", "都市神医", "重生之巅峰人生",
		"星际争霸", "玄幻世界", "武侠江湖",
		"都市异能", "历史穿越", "科幻未来",
	}

	contents := []string{
		"这是一个关于成长的故事。主人公从小镇出发，踏上了寻找真相的旅程。",
		"命运的齿轮开始转动，一场惊心动魄的冒险即将展开。",
		"在这个充满未知的世界里，每一步都可能是生死的考验。",
		"友情、爱情、亲情，在命运的洪流中交织成一幅壮丽的画卷。",
		"当真相浮出水面，一切都变得不再简单。",
	}

	for i := 0; i < count; i++ {
		title := titles[rand.Intn(len(titles))]
		if count > len(titles) {
			title = fmt.Sprintf("%s #%d", title, i+1)
		}

		// 生成5-10章（减少章节数以避免数据过长）
		chapterCount := rand.Intn(6) + 5
		chapters := make([]Chapter, chapterCount)
		totalWords := 0
		for j := 0; j < chapterCount; j++ {
			content := contents[rand.Intn(len(contents))]
			// 每章2000-5000字
			wordCount := rand.Intn(3001) + 2000
			totalWords += wordCount

			chapters[j] = Chapter{
				ChapterNumber: j + 1,
				Title:         fmt.Sprintf("第%d章：新的开始", j+1),
				Content:       content,
				WordCount:     wordCount,
			}
		}
		chaptersJSON, _ := json.Marshal(chapters)

		material := Material{
			Type:         "novel",
			Title:        title,
			Subtitle:     fmt.Sprintf("共%d章", chapterCount),
			Description:  fmt.Sprintf("长篇小说《%s》，共%d章，%d万字。精彩纷呈，不容错过！", title, chapterCount, totalWords/10000),
			CoverURL:     fmt.Sprintf("https://picsum.photos/400/600?random=%d", rand.Intn(10000)),
			ContentURL:   string(chaptersJSON),
			Author:       fmt.Sprintf("作家%d", rand.Intn(50)+1),
			Category:     "小说",
			ViewCount:    rand.Intn(200000),
			LikeCount:    rand.Intn(20000),
			CommentCount: rand.Intn(2000),
			ShareCount:   rand.Intn(800),
			CollectCount: rand.Intn(1500),
			WordCount:    totalWords,
			ChapterCount: chapterCount,
			Status:       1,
			Sort:         i + 1,
		}

		if err := insertMaterial(&material); err != nil {
			log.Printf("插入小说失败: %v", err)
		}
	}

	fmt.Printf("✅ 完成 %d 部小说\n", count)
}

// 插入物料
func insertMaterial(m *Material) error {
	// 处理tags字段，确保是有效的JSON
	tags := m.Tags
	if tags == "" {
		tags = "[]" // 空数组而不是空字符串
	}

	query := `INSERT INTO material
		(type, title, subtitle, description, cover_url, content_url, author, category,
		tags, view_count, like_count, comment_count, share_count, collect_count,
		duration, word_count, chapter_count, status, sort, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, NOW(), NOW())`

	_, err := db.Exec(query,
		m.Type, m.Title, m.Subtitle, m.Description, m.CoverURL, m.ContentURL,
		m.Author, m.Category, tags, m.ViewCount, m.LikeCount, m.CommentCount,
		m.ShareCount, m.CollectCount, m.Duration, m.WordCount, m.ChapterCount,
		m.Status, m.Sort)

	return err
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
