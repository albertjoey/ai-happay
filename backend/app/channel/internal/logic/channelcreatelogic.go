package logic

import (
	"context"
	"errors"
	"happy/app/channel/internal/svc"
	"happy/app/channel/internal/types"
	"happy/common/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChannelCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChannelCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChannelCreateLogic {
	return &ChannelCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChannelCreateLogic) ChannelCreate(req *types.ChannelCreateRequest) (*types.Channel, error) {
	// 检查code是否已存在
	var count int64
	l.svcCtx.DB.Model(&model.Channel{}).
		Where("tenant_id = ? AND code = ?", 1, req.Code).
		Count(&count)

	if count > 0 {
		return nil, errors.New("频道代码已存在")
	}

	// 如果没有指定排序，设置为最大值+1
	sort := req.Sort
	if sort == 0 {
		var maxSort int
		l.svcCtx.DB.Model(&model.Channel{}).
			Where("tenant_id = ?", 1).
			Select("COALESCE(MAX(sort), 0)").
			Scan(&maxSort)
		sort = maxSort + 1
	}

	// 创建频道
	channel := model.Channel{
		TenantID:    1,
		Name:        req.Name,
		Code:        req.Code,
		Description: req.Description,
		Icon:        req.Icon,
		Status:      1,
		Sort:        sort,
	}

	err := l.svcCtx.DB.Create(&channel).Error
	if err != nil {
		return nil, err
	}

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
