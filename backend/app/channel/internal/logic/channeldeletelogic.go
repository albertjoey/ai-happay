package logic

import (
	"context"
	"errors"
	"happy/app/channel/internal/svc"
	"happy/common/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChannelDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChannelDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChannelDeleteLogic {
	return &ChannelDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChannelDeleteLogic) ChannelDelete(id uint) error {
	var channel model.Channel
	err := l.svcCtx.DB.Where("id = ? AND tenant_id = ?", id, 1).First(&channel).Error
	if err != nil {
		return errors.New("频道不存在")
	}

	// 软删除
	return l.svcCtx.DB.Delete(&channel).Error
}
