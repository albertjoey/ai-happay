package logic

import (
	"context"
	

	"happy/app/channel/internal/svc"
	"happy/app/channel/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TopicListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTopicListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TopicListLogic {
	return &TopicListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TopicListLogic) TopicList(req *types.TopicListRequest) (*types.TopicListResponse, error) {
	// 使用Repository接口
	list, total, err := l.svcCtx.TopicRepo.List(l.ctx, req)
	if err != nil {
		return nil, err
	}

	return &types.TopicListResponse{Total: total, List: list}, nil
}

type TopicCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTopicCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TopicCreateLogic {
	return &TopicCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TopicCreateLogic) TopicCreate(req *types.TopicCreateRequest) (interface{}, error) {
	topic := &types.Topic{
		Name:        req.Name,
		Description: req.Description,
		Cover:       req.Cover,
		Sort:        req.Sort,
		Status:      1,
	}

	err := l.svcCtx.TopicRepo.Create(l.ctx, topic)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{"id": topic.ID, "message": "创建成功"}, nil
}

type TopicUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTopicUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TopicUpdateLogic {
	return &TopicUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TopicUpdateLogic) TopicUpdate(req *types.TopicUpdateRequest) (interface{}, error) {
	topic := &types.Topic{
		ID:          req.ID,
		Name:        req.Name,
		Description: req.Description,
		Cover:       req.Cover,
	}

	if req.Status != nil {
		topic.Status = *req.Status
	}
	if req.Sort != nil {
		topic.Sort = *req.Sort
	}

	err := l.svcCtx.TopicRepo.Update(l.ctx, topic)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{"message": "更新成功"}, nil
}

type TopicDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTopicDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TopicDeleteLogic {
	return &TopicDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TopicDeleteLogic) TopicDelete(id uint) (interface{}, error) {
	err := l.svcCtx.TopicRepo.Delete(l.ctx, id)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{"message": "删除成功"}, nil
}
