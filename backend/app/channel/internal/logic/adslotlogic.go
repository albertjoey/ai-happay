package logic

import (
	"context"
	"happy/app/channel/internal/svc"
	"happy/app/channel/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AdSlotListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAdSlotListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdSlotListLogic {
	return &AdSlotListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AdSlotListLogic) AdSlotList(req *types.AdSlotListRequest) (*types.AdSlotListResponse, error) {
	// 使用Repository接口
	list, err := l.svcCtx.AdSlotRepo.List(l.ctx, req.ChannelID)
	if err != nil {
		return nil, err
	}

	// 过滤状态
	if req.Status != nil {
		filtered := make([]types.AdSlot, 0)
		for _, item := range list {
			if item.Status == *req.Status {
				filtered = append(filtered, item)
			}
		}
		list = filtered
	}

	return &types.AdSlotListResponse{List: list}, nil
}

type AdSlotCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAdSlotCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdSlotCreateLogic {
	return &AdSlotCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AdSlotCreateLogic) AdSlotCreate(req *types.AdSlotCreateRequest) (interface{}, error) {
	adSlot := &types.AdSlot{
		ChannelID:   req.ChannelID,
		Name:        req.Name,
		InsertType:  req.InsertType,
		InsertRule:  req.InsertRule,
		AdType:      req.AdType,
		AdContent:   req.AdContent,
		LinkURL:     req.LinkURL,
		Sort:        req.Sort,
		Description: req.Description,
	}

	err := l.svcCtx.AdSlotRepo.Create(l.ctx, adSlot)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{"id": adSlot.ID, "message": "创建成功"}, nil
}

type AdSlotUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAdSlotUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdSlotUpdateLogic {
	return &AdSlotUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AdSlotUpdateLogic) AdSlotUpdate(req *types.AdSlotUpdateRequest) (interface{}, error) {
	adSlot := &types.AdSlot{
		ID:          req.ID,
		Name:        req.Name,
		InsertType:  req.InsertType,
		InsertRule:  req.InsertRule,
		AdType:      req.AdType,
		AdContent:   req.AdContent,
		LinkURL:     req.LinkURL,
		Description: req.Description,
	}
	if req.Sort != nil {
		adSlot.Sort = *req.Sort
	}

	err := l.svcCtx.AdSlotRepo.Update(l.ctx, adSlot)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{"message": "更新成功"}, nil
}

type AdSlotDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAdSlotDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdSlotDeleteLogic {
	return &AdSlotDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AdSlotDeleteLogic) AdSlotDelete(id uint) (interface{}, error) {
	err := l.svcCtx.AdSlotRepo.Delete(l.ctx, id)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{"message": "删除成功"}, nil
}
