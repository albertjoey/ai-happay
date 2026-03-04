package logic

import (
	"context"
	"time"
	"happy/app/content/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ContentDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewContentDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ContentDeleteLogic {
	return &ContentDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ContentDeleteLogic) ContentDelete(id uint) error {
	now := time.Now()
	
	// 软删除内容
	deleteSQL := "UPDATE content SET deleted_at = ? WHERE id = ? AND deleted_at IS NULL"
	result := l.svcCtx.DB.Exec(deleteSQL, now, id)
	if result.Error != nil {
		return result.Error
	}

	// 删除话题关联
	l.svcCtx.DB.Exec("DELETE FROM content_topic WHERE content_id = ?", id)

	// 删除标签关联
	l.svcCtx.DB.Exec("DELETE FROM content_tag WHERE content_id = ?", id)

	return nil
}
