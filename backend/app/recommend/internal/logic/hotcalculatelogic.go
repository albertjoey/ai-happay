package logic

import (
	"math"
	"time"
)

// Content 内容结构（用于热度计算）
type Content struct {
	ID           uint
	ViewCount    int
	LikeCount    int
	CommentCount int
	CollectCount int
	ShareCount   int
	PublishAt    *time.Time
	CreatedAt    time.Time
}

// CalculateHotScore 计算热度分数
// 热度分数 = (浏览×1 + 点赞×5 + 评论×10 + 收藏×8 + 分享×15) / (时间衰减)
// 时间衰减 = (当前时间 - 发布时间 + 1)^0.5
func CalculateHotScore(content *Content) float64 {
	// 基础分数
	baseScore := float64(content.ViewCount*1 +
		content.LikeCount*5 +
		content.CommentCount*10 +
		content.CollectCount*8 +
		content.ShareCount*15)

	// 时间衰减
	var publishTime time.Time
	if content.PublishAt != nil {
		publishTime = *content.PublishAt
	} else {
		publishTime = content.CreatedAt
	}

	hoursSincePublish := time.Since(publishTime).Hours()
	if hoursSincePublish < 0 {
		hoursSincePublish = 0
	}

	// 时间衰减因子（避免除零）
	timeDecay := math.Pow(hoursSincePublish+1, 0.5)

	// 最终热度分数
	hotScore := baseScore / timeDecay

	return hotScore
}

// CalculateHotScoreFromData 从数据库数据计算热度分数
func CalculateHotScoreFromData(viewCount, likeCount, commentCount, collectCount, shareCount int, publishAt, createdAt time.Time) float64 {
	content := &Content{
		ViewCount:    viewCount,
		LikeCount:    likeCount,
		CommentCount: commentCount,
		CollectCount: collectCount,
		ShareCount:   shareCount,
		CreatedAt:    createdAt,
	}

	if !publishAt.IsZero() {
		content.PublishAt = &publishAt
	}

	return CalculateHotScore(content)
}
