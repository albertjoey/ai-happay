package logic

import (
	"context"

	"happy/app/channel/internal/repository"
	"happy/app/channel/internal/svc"
	"happy/app/channel/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DiscoverConfigLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	repo   repository.DiscoverRepository
}

func NewDiscoverConfigLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DiscoverConfigLogic {
	return &DiscoverConfigLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		repo:   repository.NewDiscoverRepository(svcCtx.DB),
	}
}

// GetConfigList 获取发现页模块配置列表
func (l *DiscoverConfigLogic) GetConfigList() (*types.DiscoverConfigListResponse, error) {
	configs, err := l.repo.GetModuleList(l.ctx)
	if err != nil {
		return nil, err
	}

	return &types.DiscoverConfigListResponse{
		Total: int64(len(configs)),
		List:  configs,
	}, nil
}

// UpdateConfig 更新发现页模块配置
func (l *DiscoverConfigLogic) UpdateConfig(req *types.DiscoverConfigUpdateRequest) error {
	return l.repo.UpdateModule(l.ctx, req)
}
