package repository

import (
	"context"

	"happy/app/channel/internal/types"
)

// DiscoverRepository 发现页仓储接口
type DiscoverRepository interface {
	// 模块配置相关
	GetModuleList(ctx context.Context) ([]types.DiscoverConfig, error)
	UpdateModule(ctx context.Context, req *types.DiscoverConfigUpdateRequest) error
	GetModuleIDByType(ctx context.Context, moduleType string) (uint, error)

	// 内容管理相关
	GetContentList(ctx context.Context, module string, page, pageSize int) ([]types.DiscoverItem, int64, error)
	CreateContent(ctx context.Context, moduleID uint, contentType string, contentID uint, sort int) error
	UpdateContent(ctx context.Context, id uint, sort int, status int8) error
	DeleteContent(ctx context.Context, id uint) error

	// 发现页数据获取
	GetDiscoverModules(ctx context.Context) ([]map[string]interface{}, error)
	GetModuleContents(ctx context.Context, moduleID uint, moduleType string) ([]map[string]interface{}, error)
}
