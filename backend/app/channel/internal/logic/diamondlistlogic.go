package logic

import (
	"context"
	"happy/app/channel/internal/svc"
	"happy/app/channel/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DiamondListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDiamondListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DiamondListLogic {
	return &DiamondListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DiamondListLogic) DiamondList(req *types.DiamondListRequest) (*types.DiamondListResponse, error) {
	// 使用Repository接口
	list, err := l.svcCtx.DiamondRepo.List(l.ctx, req.ChannelID)
	if err != nil {
		return nil, err
	}

	// 过滤条件
	if req.GroupID != nil || req.Status != nil {
		filtered := make([]types.Diamond, 0)
		for _, d := range list {
			if req.GroupID != nil && d.GroupID != *req.GroupID {
				continue
			}
			if req.Status != nil && d.Status != *req.Status {
				continue
			}
			filtered = append(filtered, d)
		}
		list = filtered
	}

	return &types.DiamondListResponse{List: list}, nil
}
