package logic

import (
	"context"

	"happy/app/channel/internal/svc"
	"happy/app/channel/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TagListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTagListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TagListLogic {
	return &TagListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TagListLogic) TagList(req *types.TagListRequest) (*types.TagListResponse, error) {
	// 使用Repository接口
	list, total, err := l.svcCtx.TagRepo.List(l.ctx, req)
	if err != nil {
		return nil, err
	}

	return &types.TagListResponse{Total: total, List: list}, nil
}

type TagCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTagCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TagCreateLogic {
	return &TagCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TagCreateLogic) TagCreate(req *types.TagCreateRequest) (interface{}, error) {
	tag := &types.Tag{
		Name: req.Name,
		Type: req.Type,
		Sort: req.Sort,
		Status: 1,
	}

	err := l.svcCtx.TagRepo.Create(l.ctx, tag)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{"id": tag.ID, "message": "创建成功"}, nil
}

type TagUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTagUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TagUpdateLogic {
	return &TagUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TagUpdateLogic) TagUpdate(req *types.TagUpdateRequest) (interface{}, error) {
	tag := &types.Tag{
		ID:   req.ID,
		Name: req.Name,
	}

	if req.Type != nil {
		tag.Type = *req.Type
	}
	if req.Status != nil {
		tag.Status = *req.Status
	}
	if req.Sort != nil {
		tag.Sort = *req.Sort
	}

	err := l.svcCtx.TagRepo.Update(l.ctx, tag)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{"message": "更新成功"}, nil
}

type TagDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTagDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TagDeleteLogic {
	return &TagDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TagDeleteLogic) TagDelete(id uint) (interface{}, error) {
	err := l.svcCtx.TagRepo.Delete(l.ctx, id)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{"message": "删除成功"}, nil
}
