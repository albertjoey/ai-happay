package logic

import (
	"context"
	"happy/app/channel/internal/svc"
	"happy/app/channel/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChannelUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChannelUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChannelUpdateLogic {
	return &ChannelUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChannelUpdateLogic) ChannelUpdate(req *types.ChannelUpdateRequest) (*types.Channel, error) {
	channel := &types.Channel{
		ID:          req.ID,
		Name:        req.Name,
		Description: req.Description,
		Icon:        req.Icon,
	}
	if req.Status != nil {
		channel.Status = *req.Status
	}
	if req.Sort != nil {
		channel.Sort = *req.Sort
	}

	err := l.svcCtx.ChannelRepo.Update(l.ctx, channel)
	if err != nil {
		return nil, err
	}

	// 返回更新后的数据
	return l.svcCtx.ChannelRepo.FindByID(l.ctx, req.ID)
}
