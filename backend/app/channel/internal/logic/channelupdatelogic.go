package logic

import (
	"context"
	"errors"
	"happy/app/channel/internal/svc"
	"happy/app/channel/internal/types"
	"happy/common/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChannelUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChannelUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChannelUpdateLogic {
	return &ChannelUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChannelUpdateLogic) ChannelUpdate(req *types.ChannelUpdateRequest) (*types.Channel, error) {
	var channel model.Channel
	err := l.svcCtx.DB.Where("id = ? AND tenant_id = ?", req.ID, 1).First(&channel).Error
	if err != nil {
		return nil, errors.New("频道不存在")
	}

	// 更新字段
	updates := make(map[string]interface{})
	if req.Name != "" {
		updates["name"] = req.Name
	}
	if req.Description != "" {
		updates["description"] = req.Description
	}
	if req.Icon != "" {
		updates["icon"] = req.Icon
	}
	if req.Status != nil {
		updates["status"] = *req.Status
	}
	if req.Sort != nil {
		updates["sort"] = *req.Sort
	}

	if len(updates) > 0 {
		err = l.svcCtx.DB.Model(&channel).Updates(updates).Error
		if err != nil {
			return nil, err
		}
	}

	// 重新查询
	l.svcCtx.DB.Where("id = ?", req.ID).First(&channel)

	return &types.Channel{
		ID:          channel.ID,
		Name:        channel.Name,
		Code:        channel.Code,
		Description: channel.Description,
		Icon:        channel.Icon,
		Status:      channel.Status,
		Sort:        channel.Sort,
	}, nil
}
