package logic

import (
	"context"
	"happy/app/channel/internal/svc"
	"happy/app/channel/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MaterialListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMaterialListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MaterialListLogic {
	return &MaterialListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MaterialListLogic) MaterialList(req *types.MaterialListRequest) (*types.MaterialListResponse, error) {
	// 使用Repository接口 - 解耦数据库实现
	list, total, err := l.svcCtx.MaterialRepo.List(l.ctx, req)
	if err != nil {
		return nil, err
	}

	return &types.MaterialListResponse{
		Total: total,
		List:  list,
	}, nil
}

type MaterialCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMaterialCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MaterialCreateLogic {
	return &MaterialCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MaterialCreateLogic) MaterialCreate(req *types.MaterialCreateRequest) (interface{}, error) {
	// 构建物料对象
	material := &types.Material{
		Title:       req.Title,
		Subtitle:    req.Subtitle,
		Type:        req.Type,
		CoverURL:    req.CoverURL,
		ContentURL:  req.ContentURL,
		Description: req.Description,
		Author:      req.Author,
		Category:    req.Category,
		Duration:    req.Duration,
		WordCount:   req.WordCount,
		Status:      1,
		Sort:        0,
	}

	// 使用Repository接口创建
	err := l.svcCtx.MaterialRepo.Create(l.ctx, material)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{"id": material.ID, "success": true}, nil
}

type MaterialUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMaterialUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MaterialUpdateLogic {
	return &MaterialUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MaterialUpdateLogic) MaterialUpdate(req *types.MaterialUpdateRequest) (interface{}, error) {
	// 构建物料对象
	material := &types.Material{
		ID:          req.ID,
		Title:       req.Title,
		Subtitle:    req.Subtitle,
		Description: req.Description,
	}
	if req.Status != nil {
		material.Status = *req.Status
	}

	// 使用Repository接口更新
	err := l.svcCtx.MaterialRepo.Update(l.ctx, material)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{"success": true}, nil
}

type MaterialDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMaterialDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MaterialDeleteLogic {
	return &MaterialDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MaterialDeleteLogic) MaterialDelete(id uint) (interface{}, error) {
	// 使用Repository接口删除
	err := l.svcCtx.MaterialRepo.Delete(l.ctx, id)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{"success": true}, nil
}
