package logic

import (
	"context"
	"happy/app/channel/internal/svc"
	"happy/app/channel/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChannelSortLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChannelSortLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChannelSortLogic {
	return &ChannelSortLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChannelSortLogic) ChannelSort(req *types.ChannelSortRequest) error {
	// 批量更新排序
	for _, item := range req.Items {
		l.svcCtx.ChannelRepo.UpdateSort(l.ctx, item.ID, item.Sort)
	}
	return nil
}
