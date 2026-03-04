package logic

import (
	"context"
	"happy/app/channel/internal/svc"
	"happy/app/channel/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DiamondListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDiamondListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DiamondListLogic {
	return &DiamondListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DiamondListLogic) DiamondList(req *types.DiamondListRequest) (*types.DiamondListResponse, error) {
	querySQL := `
		SELECT d.id, d.channel_id, d.group_id, d.sort, d.title, d.icon, d.link_type, d.link_value, 
		       d.status, d.description, d.material_id,
		       m.title as material_title, m.type as material_type, m.cover_url as material_cover_url,
		       m.author as material_author, m.view_count as material_view_count
		FROM diamond d
		LEFT JOIN material m ON d.material_id = m.id
		WHERE d.tenant_id = ? AND d.channel_id = ? AND d.deleted_at IS NULL
	`
	args := []interface{}{1, req.ChannelID}

	if req.GroupID != nil {
		querySQL += " AND d.group_id = ?"
		args = append(args, *req.GroupID)
	}

	if req.Status != nil {
		querySQL += " AND d.status = ?"
		args = append(args, *req.Status)
	}

	querySQL += " ORDER BY d.group_id ASC, d.sort ASC"

	type DiamondResult struct {
		ID                uint
		ChannelID         uint
		GroupID           int
		Sort              int
		Title             string
		Icon              string
		LinkType          string
		LinkValue         string
		Status            int8
		Description       string
		MaterialID        uint
		MaterialTitle     *string
		MaterialType      *string
		MaterialCoverURL  *string
		MaterialAuthor    *string
		MaterialViewCount *uint
	}

	var results []DiamondResult
	err := l.svcCtx.DB.Raw(querySQL, args...).Scan(&results).Error
	if err != nil {
		return nil, err
	}

	list := make([]types.Diamond, 0, len(results))
	for _, d := range results {
		diamond := types.Diamond{
			ID:          d.ID,
			ChannelID:   d.ChannelID,
			GroupID:     d.GroupID,
			Sort:        d.Sort,
			Title:       d.Title,
			Icon:        d.Icon,
			LinkType:    d.LinkType,
			LinkValue:   d.LinkValue,
			Status:      d.Status,
			Description: d.Description,
			MaterialID:  d.MaterialID,
		}

		// 如果有关联物料,添加物料信息
		if d.MaterialID > 0 && d.MaterialTitle != nil {
			diamond.Material = map[string]interface{}{
				"id":         d.MaterialID,
				"title":      *d.MaterialTitle,
				"type":       *d.MaterialType,
				"cover_url":  *d.MaterialCoverURL,
				"author":     *d.MaterialAuthor,
				"view_count": *d.MaterialViewCount,
			}
		}

		list = append(list, diamond)
	}

	return &types.DiamondListResponse{List: list}, nil
}
