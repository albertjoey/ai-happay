package logic

import (
	"context"
	"encoding/json"
	"time"
	"happy/app/content/internal/svc"
	"happy/app/content/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ContentUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewContentUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ContentUpdateLogic {
	return &ContentUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ContentUpdateLogic) ContentUpdate(req *types.UpdateContentRequest) (*types.ContentResponse, error) {
	now := time.Now()
	
	// 构建更新SQL
	updateSQL := "UPDATE content SET updated_at = ?"
	args := []interface{}{now}

	if req.Title != "" {
		updateSQL += ", title = ?"
		args = append(args, req.Title)
	}
	if req.Description != "" {
		updateSQL += ", description = ?"
		args = append(args, req.Description)
	}
	if req.Cover != "" {
		updateSQL += ", cover = ?"
		args = append(args, req.Cover)
	}
	if len(req.Media) > 0 {
		mediaJSON, err := json.Marshal(req.Media)
		if err != nil {
			return nil, err
		}
		updateSQL += ", media = ?"
		args = append(args, string(mediaJSON))
	}

	updateSQL += " WHERE id = ? AND deleted_at IS NULL"
	args = append(args, req.ID)

	result := l.svcCtx.DB.Exec(updateSQL, args...)
	if result.Error != nil {
		return nil, result.Error
	}

	// 更新话题关联
	if len(req.Topics) > 0 {
		// 删除旧关联
		l.svcCtx.DB.Exec("DELETE FROM content_topic WHERE content_id = ?", req.ID)
		// 插入新关联
		for _, topicID := range req.Topics {
			l.svcCtx.DB.Exec("INSERT INTO content_topic (content_id, topic_id, created_at) VALUES (?, ?, ?)", req.ID, topicID, now)
		}
	}

	// 更新标签关联
	if len(req.Tags) > 0 {
		// 删除旧关联
		l.svcCtx.DB.Exec("DELETE FROM content_tag WHERE content_id = ?", req.ID)
		// 插入新关联
		for _, tagID := range req.Tags {
			l.svcCtx.DB.Exec("INSERT INTO content_tag (content_id, tag_id, created_at) VALUES (?, ?, ?)", req.ID, tagID, now)
		}
	}

	// 返回更新后的内容
	detailLogic := NewContentDetailLogic(l.ctx, l.svcCtx)
	return detailLogic.ContentDetail(req.ID)
}
