package logic

import (
	"context"
	"encoding/json"
	"happy/app/channel/internal/svc"
	"happy/app/channel/internal/types"
	"time"

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
	querySQL := `
		SELECT id, channel_id, name, insert_type, insert_rule, ad_type, ad_content, link_url, status, sort, description
		FROM ad_slot
		WHERE tenant_id = ? AND channel_id = ? AND deleted_at IS NULL
	`
	args := []interface{}{1, req.ChannelID}

	if req.Status != nil {
		querySQL += " AND status = ?"
		args = append(args, *req.Status)
	}

	querySQL += " ORDER BY sort ASC"

	type AdSlotResult struct {
		ID          uint
		ChannelID   uint
		Name        string
		InsertType  string
		InsertRule  []byte
		AdType      string
		AdContent   []byte
		LinkURL     string
		Status      int8
		Sort        int
		Description string
	}

	var results []AdSlotResult
	err := l.svcCtx.DB.Raw(querySQL, args...).Scan(&results).Error
	if err != nil {
		return nil, err
	}

	list := make([]types.AdSlot, 0, len(results))
	for _, a := range results {
		var insertRule map[string]interface{}
		var adContent map[string]interface{}

		if a.InsertRule != nil {
			json.Unmarshal(a.InsertRule, &insertRule)
		}
		if a.AdContent != nil {
			json.Unmarshal(a.AdContent, &adContent)
		}

		list = append(list, types.AdSlot{
			ID:          a.ID,
			ChannelID:   a.ChannelID,
			Name:        a.Name,
			InsertType:  a.InsertType,
			InsertRule:  insertRule,
			AdType:      a.AdType,
			AdContent:   adContent,
			LinkURL:     a.LinkURL,
			Status:      a.Status,
			Sort:        a.Sort,
			Description: a.Description,
		})
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
	now := time.Now()

	insertRuleJSON, _ := json.Marshal(req.InsertRule)
	adContentJSON, _ := json.Marshal(req.AdContent)

	insertSQL := `
		INSERT INTO ad_slot (tenant_id, channel_id, name, insert_type, insert_rule, ad_type, ad_content, link_url, status, sort, description, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, 1, ?, ?, ?, ?)
	`

	result := l.svcCtx.DB.Exec(insertSQL, 1, req.ChannelID, req.Name, req.InsertType, insertRuleJSON, req.AdType, adContentJSON, req.LinkURL, req.Sort, req.Description, now, now)
	if result.Error != nil {
		return nil, result.Error
	}

	return map[string]interface{}{"message": "创建成功"}, nil
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
	updateSQL := "UPDATE ad_slot SET updated_at = ?"
	args := []interface{}{time.Now()}

	if req.Name != "" {
		updateSQL += ", name = ?"
		args = append(args, req.Name)
	}
	if req.InsertType != "" {
		updateSQL += ", insert_type = ?"
		args = append(args, req.InsertType)
	}
	if req.InsertRule != nil {
		insertRuleJSON, _ := json.Marshal(req.InsertRule)
		updateSQL += ", insert_rule = ?"
		args = append(args, insertRuleJSON)
	}
	if req.AdType != "" {
		updateSQL += ", ad_type = ?"
		args = append(args, req.AdType)
	}
	if req.AdContent != nil {
		adContentJSON, _ := json.Marshal(req.AdContent)
		updateSQL += ", ad_content = ?"
		args = append(args, adContentJSON)
	}
	if req.LinkURL != "" {
		updateSQL += ", link_url = ?"
		args = append(args, req.LinkURL)
	}
	if req.Sort != nil {
		updateSQL += ", sort = ?"
		args = append(args, *req.Sort)
	}
	if req.Status != nil {
		updateSQL += ", status = ?"
		args = append(args, *req.Status)
	}
	if req.Description != "" {
		updateSQL += ", description = ?"
		args = append(args, req.Description)
	}

	updateSQL += " WHERE id = ? AND tenant_id = ? AND deleted_at IS NULL"
	args = append(args, req.ID, 1)

	result := l.svcCtx.DB.Exec(updateSQL, args...)
	if result.Error != nil {
		return nil, result.Error
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
	deleteSQL := "UPDATE ad_slot SET deleted_at = ? WHERE id = ? AND tenant_id = ?"
	result := l.svcCtx.DB.Exec(deleteSQL, time.Now(), id, 1)
	if result.Error != nil {
		return nil, result.Error
	}

	return map[string]interface{}{"message": "删除成功"}, nil
}
