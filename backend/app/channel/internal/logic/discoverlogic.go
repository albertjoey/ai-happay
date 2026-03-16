package logic

import (
	"context"

	"happy/app/channel/internal/repository"
	"happy/app/channel/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DiscoverLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	repo   repository.DiscoverRepository
}

func NewDiscoverLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DiscoverLogic {
	return &DiscoverLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		repo:   repository.NewDiscoverRepository(svcCtx.DB),
	}
}

func (l *DiscoverLogic) Discover() (interface{}, error) {
	// 从repository获取模块列表
	modules, err := l.repo.GetDiscoverModules(l.ctx)
	if err != nil {
		return nil, err
	}

	// 构建返回数据
	result := map[string]interface{}{
		"modules": make([]map[string]interface{}, 0),
	}

	for _, module := range modules {
		moduleID := module["id"].(uint)
		moduleType := module["module"].(string)

		// 获取模块内容
		items, err := l.repo.GetModuleContents(l.ctx, moduleID, moduleType)
		if err != nil {
			items = []map[string]interface{}{}
		}

		moduleData := map[string]interface{}{
			"module": moduleType,
			"title":  module["title"],
			"items":  items,
		}
		result["modules"] = append(result["modules"].([]map[string]interface{}), moduleData)
	}

	return result, nil
}
