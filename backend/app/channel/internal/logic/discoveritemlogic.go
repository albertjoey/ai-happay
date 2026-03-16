package logic

import (
	"context"

	"happy/app/channel/internal/repository"
	"happy/app/channel/internal/svc"
	"happy/app/channel/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DiscoverItemLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	repo   repository.DiscoverRepository
}

func NewDiscoverItemLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DiscoverItemLogic {
	return &DiscoverItemLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		repo:   repository.NewDiscoverRepository(svcCtx.DB),
	}
}

// GetItemList 获取发现页内容列表
func (l *DiscoverItemLogic) GetItemList(req *types.DiscoverItemListRequest) (*types.DiscoverItemListResponse, error) {
	items, total, err := l.repo.GetContentList(l.ctx, req.Module, req.Page, req.PageSize)
	if err != nil {
		return nil, err
	}

	// 填充内容详情
	for i := range items {
		l.fillItemDetails(&items[i])
	}

	return &types.DiscoverItemListResponse{
		Total: total,
		List:  items,
	}, nil
}

// fillItemDetails 填充内容详情
func (l *DiscoverItemLogic) fillItemDetails(item *types.DiscoverItem) {
	switch item.ItemType {
	case "topic":
		var topic struct {
			Title     string
			CoverURL  string
			ViewCount int
			PostCount int
		}
		err := l.svcCtx.DB.Raw(`
			SELECT name as title, cover as cover_url, 0 as view_count, 0 as post_count
			FROM topic WHERE id = ?
		`, item.ItemID).Scan(&topic).Error
		if err == nil {
			item.Title = topic.Title
			item.CoverURL = topic.CoverURL
			item.Views = topic.ViewCount
			item.Count = topic.PostCount
		}

	case "material":
		var material struct {
			Title     string
			CoverURL  string
			Author    string
			ViewCount int
		}
		err := l.svcCtx.DB.Raw(`
			SELECT title, cover_url, author, view_count
			FROM material WHERE id = ?
		`, item.ItemID).Scan(&material).Error
		if err == nil {
			item.Title = material.Title
			item.CoverURL = material.CoverURL
			item.Author = material.Author
			item.Views = material.ViewCount
		}
	}
}

// CreateItem 创建发现页内容
func (l *DiscoverItemLogic) CreateItem(req *types.DiscoverItemCreateRequest) error {
	// 获取模块ID
	moduleID, err := l.repo.GetModuleIDByType(l.ctx, req.Module)
	if err != nil {
		return err
	}

	return l.repo.CreateContent(l.ctx, moduleID, req.ItemType, req.ItemID, req.SortOrder)
}

// UpdateItem 更新发现页内容
func (l *DiscoverItemLogic) UpdateItem(req *types.DiscoverItemUpdateRequest) error {
	sort := 0
	if req.SortOrder != nil {
		sort = *req.SortOrder
	}

	status := int8(1)
	if req.IsEnabled != nil && !*req.IsEnabled {
		status = 0
	}

	return l.repo.UpdateContent(l.ctx, req.ID, sort, status)
}

// DeleteItem 删除发现页内容
func (l *DiscoverItemLogic) DeleteItem(id uint) error {
	return l.repo.DeleteContent(l.ctx, id)
}
