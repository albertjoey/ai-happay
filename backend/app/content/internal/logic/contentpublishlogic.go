package logic

import (
	"context"
	"time"
	"happy/app/content/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ContentPublishLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewContentPublishLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ContentPublishLogic {
	return &ContentPublishLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ContentPublishLogic) ContentPublish(id uint) error {
	now := time.Now()
	
	// 更新状态为已发布
	updateSQL := "UPDATE content SET status = 1, publish_at = ?, updated_at = ? WHERE id = ? AND deleted_at IS NULL"
	result := l.svcCtx.DB.Exec(updateSQL, now, now, id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

type ContentUnpublishLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewContentUnpublishLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ContentUnpublishLogic {
	return &ContentUnpublishLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ContentUnpublishLogic) ContentUnpublish(id uint) error {
	now := time.Now()
	
	// 更新状态为已下架
	updateSQL := "UPDATE content SET status = 2, updated_at = ? WHERE id = ? AND deleted_at IS NULL"
	result := l.svcCtx.DB.Exec(updateSQL, now, id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
