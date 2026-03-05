package logic

import (
	"context"
	"happy/app/channel/internal/svc"
	"happy/app/channel/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChannelListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChannelListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChannelListLogic {
	return &ChannelListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChannelListLogic) ChannelList(req *types.ChannelListRequest) (*types.ChannelListResponse, error) {
	// 使用Repository接口 - 解耦数据库实现
	list, total, err := l.svcCtx.ChannelRepo.List(l.ctx, req)
	if err != nil {
		return nil, err
	}

	return &types.ChannelListResponse{
		Total: total,
		List:  list,
	}, nil
}
