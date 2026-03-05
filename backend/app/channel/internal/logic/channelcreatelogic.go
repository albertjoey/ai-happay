package logic

import (
	"context"
	"happy/app/channel/internal/svc"
	"happy/app/channel/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChannelCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChannelCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChannelCreateLogic {
	return &ChannelCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChannelCreateLogic) ChannelCreate(req *types.ChannelCreateRequest) (*types.Channel, error) {
	channel := &types.Channel{
		Name:        req.Name,
		Code:        req.Code,
		Description: req.Description,
		Icon:        req.Icon,
		Status:      1,
		Sort:        req.Sort,
	}

	err := l.svcCtx.ChannelRepo.Create(l.ctx, channel)
	if err != nil {
		return nil, err
	}

	return channel, nil
}
