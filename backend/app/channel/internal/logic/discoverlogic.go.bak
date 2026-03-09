package logic

import (
	"context"
	"encoding/json"
	"errors"

	"happy/app/channel/internal/svc"
	"happy/app/channel/internal/types"
	"happy/common/xerr"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type DiscoverLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDiscoverLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DiscoverLogic {
	return &DiscoverLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// GetDiscoverPage 获取发现页完整数据
func (l *DiscoverLogic) GetDiscoverPage(req *types.DiscoverPageRequest) (resp *types.DiscoverPageResponse, err error) {
	resp = &types.DiscoverPageResponse{
		Modules: make([]types.DiscoverModuleData, 0),
	}

	// 获取所有启用的配置
	configs, err := l.svcCtx.DB.DiscoverConfig.WithContext(l.ctx).
		Where("is_enabled = ?", true).
		Order("sort_order ASC").
		Find()

	if err != nil {
		return nil, xerr.NewErrMsg("获取发现页配置失败")
	}

	// 遍历每个模块,获取对应的内容
	for _, config := range configs {
		items, err := l.svcCtx.DB.DiscoverItem.WithContext(l.ctx).
			Where("module = ?", config.Module).
			Order("sort_order ASC").
			Find()

		if err != nil {
			continue
		}

		moduleData := types.DiscoverModuleData{
			Module: config.Module,
			Title:  config.Title,
			Items:  make([]interface{}, 0),
		}

		// 将内容转换为interface{}
		for _, item := range items {
			itemMap := map[string]interface{}{
				"id":         item.ID,
				"item_type":  item.ItemType,
				"item_id":    item.ItemID,
				"title":      item.Title,
				"cover_url":  item.CoverURL,
				"sort_order": item.SortOrder,
			}

			// 解析extra_data
			if item.ExtraData != "" {
				var extraData map[string]interface{}
				if err := json.Unmarshal([]byte(item.ExtraData), &extraData); err == nil {
					for k, v := range extraData {
						itemMap[k] = v
					}
				}
			}

			moduleData.Items = append(moduleData.Items, itemMap)
		}

		resp.Modules = append(resp.Modules, moduleData)
	}

	return resp, nil
}

// GetDiscoverConfigList 获取发现页配置列表
func (l *DiscoverLogic) GetDiscoverConfigList(req *types.DiscoverConfigListRequest) (resp *types.DiscoverConfigListResponse, err error) {
	resp = &types.DiscoverConfigListResponse{}

	query := l.svcCtx.DB.DiscoverConfig.WithContext(l.ctx)

	if req.Module != "" {
		query = query.Where("module = ?", req.Module)
	}

	total, err := query.Count()
	if err != nil {
		return nil, xerr.NewErrMsg("获取配置总数失败")
	}

	configs, err := query.Offset((req.Page - 1) * req.PageSize).Limit(req.PageSize).Find()
	if err != nil {
		return nil, xerr.NewErrMsg("获取配置列表失败")
	}

	resp.Total = total
	resp.List = make([]types.DiscoverConfig, 0)
	_ = copier.Copy(&resp.List, configs)

	return resp, nil
}

// UpdateDiscoverConfig 更新发现页配置
func (l *DiscoverLogic) UpdateDiscoverConfig(req *types.DiscoverConfigUpdateRequest) error {
	updates := make(map[string]interface{})

	if req.Title != "" {
		updates["title"] = req.Title
	}
	if req.IsEnabled != nil {
		updates["is_enabled"] = *req.IsEnabled
	}
	if req.SortOrder != nil {
		updates["sort_order"] = *req.SortOrder
	}
	if req.Config != nil {
		configJSON, _ := json.Marshal(req.Config)
		updates["config"] = string(configJSON)
	}

	err := l.svcCtx.DB.DiscoverConfig.WithContext(l.ctx).
		Where("id = ?", req.ID).
		Updates(updates)

	if err != nil {
		return xerr.NewErrMsg("更新配置失败")
	}

	return nil
}

// GetDiscoverItemList 获取发现页内容列表
func (l *DiscoverLogic) GetDiscoverItemList(req *types.DiscoverItemListRequest) (resp *types.DiscoverItemListResponse, err error) {
	resp = &types.DiscoverItemListResponse{}

	query := l.svcCtx.DB.DiscoverItem.WithContext(l.ctx).
		Where("module = ?", req.Module)

	total, err := query.Count()
	if err != nil {
		return nil, xerr.NewErrMsg("获取内容总数失败")
	}

	items, err := query.Offset((req.Page - 1) * req.PageSize).Limit(req.PageSize).Find()
	if err != nil {
		return nil, xerr.NewErrMsg("获取内容列表失败")
	}

	resp.Total = total
	resp.List = make([]types.DiscoverItem, 0)
	_ = copier.Copy(&resp.List, items)

	return resp, nil
}

// CreateDiscoverItem 创建发现页内容
func (l *DiscoverLogic) CreateDiscoverItem(req *types.DiscoverItemCreateRequest) error {
	extraDataJSON, _ := json.Marshal(req.ExtraData)

	item := &model.DiscoverItem{
		Module:    req.Module,
		ItemType:  req.ItemType,
		ItemID:    req.ItemID,
		Title:     req.Title,
		CoverURL:  req.CoverURL,
		ExtraData: string(extraDataJSON),
		SortOrder: req.SortOrder,
	}

	err := l.svcCtx.DB.DiscoverItem.WithContext(l.ctx).Create(item)
	if err != nil {
		return xerr.NewErrMsg("创建内容失败")
	}

	return nil
}

// UpdateDiscoverItem 更新发现页内容
func (l *DiscoverLogic) UpdateDiscoverItem(req *types.DiscoverItemUpdateRequest) error {
	updates := make(map[string]interface{})

	if req.Title != "" {
		updates["title"] = req.Title
	}
	if req.CoverURL != "" {
		updates["cover_url"] = req.CoverURL
	}
	if req.ExtraData != nil {
		extraDataJSON, _ := json.Marshal(req.ExtraData)
		updates["extra_data"] = string(extraDataJSON)
	}
	if req.SortOrder != nil {
		updates["sort_order"] = *req.SortOrder
	}
	if req.IsEnabled != nil {
		updates["is_enabled"] = *req.IsEnabled
	}

	err := l.svcCtx.DB.DiscoverItem.WithContext(l.ctx).
		Where("id = ?", req.ID).
		Updates(updates)

	if err != nil {
		return xerr.NewErrMsg("更新内容失败")
	}

	return nil
}

// DeleteDiscoverItem 删除发现页内容
func (l *DiscoverLogic) DeleteDiscoverItem(id uint) error {
	err := l.svcCtx.DB.DiscoverItem.WithContext(l.ctx).
		Where("id = ?", id).
		Delete()

	if err != nil {
		return xerr.NewErrMsg("删除内容失败")
	}

	return nil
}
