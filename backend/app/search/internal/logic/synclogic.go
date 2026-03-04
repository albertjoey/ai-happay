package logic

import (
	"context"
	"fmt"
	"time"

	"happy/app/search/internal/svc"
	"happy/app/search/internal/types"
	"happy/common/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type SyncLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSyncLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SyncLogic {
	return &SyncLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// SyncContent 同步单个内容到ES
func (l *SyncLogic) SyncContent(contentID uint) error {
	var content model.Content
	if err := l.svcCtx.DB.First(&content, contentID).Error; err != nil {
		return err
	}

	// 只同步已发布的内容
	if content.Status != 1 {
		return nil
	}

	// 获取作者信息
	var user model.User
	l.svcCtx.DB.First(&user, content.UserID)

	// 获取标签
	var tags []string
	l.svcCtx.DB.Table("tag").
		Select("tag.name").
		Joins("JOIN content_tag ON content_tag.tag_id = tag.id").
		Where("content_tag.content_id = ?", contentID).
		Pluck("tag.name", &tags)

	// 获取话题
	var topics []string
	l.svcCtx.DB.Table("topic").
		Select("topic.name").
		Joins("JOIN content_topic ON content_topic.topic_id = topic.id").
		Where("content_topic.content_id = ?", contentID).
		Pluck("topic.name", &topics)

	// 构建索引内容
	indexContent := &types.IndexContent{
		ID:          content.ID,
		Title:       content.Title,
		Description: content.Description,
		Type:        content.Type,
		Tags:        tags,
		Topics:      topics,
		AuthorName:  user.Nickname,
		ViewCount:   content.ViewCount,
		LikeCount:   content.LikeCount,
		Status:      content.Status,
		TenantID:    content.TenantID,
		CreatedAt:   content.CreatedAt.Format(time.RFC3339),
	}

	// 索引到ES
	return l.svcCtx.ES.IndexContent(indexContent)
}

// DeleteContent 从ES删除内容
func (l *SyncLogic) DeleteContent(contentID uint) error {
	return l.svcCtx.ES.DeleteContent(contentID)
}

// SyncAll 全量同步
func (l *SyncLogic) SyncAll(tenantID uint) error {
	var contents []model.Content
	if err := l.svcCtx.DB.Where("status = ? AND tenant_id = ?", 1, tenantID).Find(&contents).Error; err != nil {
		return err
	}

	successCount := 0
	failCount := 0

	for _, content := range contents {
		if err := l.SyncContent(content.ID); err != nil {
			logx.Errorf("同步内容失败 ID=%d, Error=%v", content.ID, err)
			failCount++
		} else {
			successCount++
		}
	}

	return fmt.Errorf("同步完成: 成功%d, 失败%d", successCount, failCount)
}
