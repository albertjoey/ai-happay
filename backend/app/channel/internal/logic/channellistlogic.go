package logic

import (
	"context"
	"happy/app/channel/internal/svc"
	"happy/app/channel/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChannelListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChannelListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChannelListLogic {
	return &ChannelListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChannelListLogic) ChannelList(req *types.ChannelListRequest) (*types.ChannelListResponse, error) {
	var total int64

	// 使用原生SQL避免GORM模型问题
	countSQL := "SELECT COUNT(*) FROM channel WHERE tenant_id = ? AND deleted_at IS NULL"
	l.svcCtx.DB.Raw(countSQL, 1).Scan(&total)

	// 构建查询SQL
	querySQL := `
		SELECT id, name, code, description, icon, status, sort
		FROM channel
		WHERE tenant_id = ? AND deleted_at IS NULL
	`
	args := []interface{}{1}

	// 条件过滤
	if req.Name != "" {
		querySQL += " AND name LIKE ?"
		args = append(args, "%"+req.Name+"%")
	}
	if req.Status != nil {
		querySQL += " AND status = ?"
		args = append(args, *req.Status)
	}

	// 排序和分页
	querySQL += " ORDER BY sort ASC, id ASC LIMIT ? OFFSET ?"
	offset := (req.Page - 1) * req.PageSize
	args = append(args, req.PageSize, offset)

	// 执行查询
	type ChannelResult struct {
		ID          uint
		Name        string
		Code        string
		Description string
		Icon        string
		Status      int8
		Sort        int
	}
	var results []ChannelResult
	err := l.svcCtx.DB.Raw(querySQL, args...).Scan(&results).Error

	if err != nil {
		return nil, err
	}

	// 转换响应
	list := make([]types.Channel, 0, len(results))
	for _, ch := range results {
		list = append(list, types.Channel{
			ID:          ch.ID,
			Name:        ch.Name,
			Code:        ch.Code,
			Description: ch.Description,
			Icon:        ch.Icon,
			Status:      ch.Status,
			Sort:        ch.Sort,
		})
	}

	return &types.ChannelListResponse{
		Total: total,
		List:  list,
	}, nil
}
