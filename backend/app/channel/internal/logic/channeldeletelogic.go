package logic

import (
	"context"
	"happy/app/channel/internal/svc"

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
	return l.svcCtx.ChannelRepo.Delete(l.ctx, id)
}
