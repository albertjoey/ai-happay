package logic

import (
	"context"
	"encoding/json"
	"happy/app/channel/internal/svc"
	"happy/app/channel/internal/types"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type DiamondCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDiamondCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DiamondCreateLogic {
	return &DiamondCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DiamondCreateLogic) DiamondCreate(req *types.DiamondCreateRequest) (interface{}, error) {
	now := time.Now()
	insertSQL := `
		INSERT INTO diamond (tenant_id, channel_id, group_id, sort, title, icon, link_type, link_value, status, description, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, 1, ?, ?, ?)
	`

	result := l.svcCtx.DB.Exec(insertSQL, 1, req.ChannelID, req.GroupID, req.Sort, req.Title, req.Icon, req.LinkType, req.LinkValue, req.Description, now, now)
	if result.Error != nil {
		return nil, result.Error
	}

	return map[string]interface{}{"message": "创建成功"}, nil
}

type DiamondUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDiamondUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DiamondUpdateLogic {
	return &DiamondUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DiamondUpdateLogic) DiamondUpdate(req *types.DiamondUpdateRequest) (interface{}, error) {
	updateSQL := "UPDATE diamond SET updated_at = ?"
	args := []interface{}{time.Now()}

	if req.GroupID != nil {
		updateSQL += ", group_id = ?"
		args = append(args, *req.GroupID)
	}
	if req.Sort != nil {
		updateSQL += ", sort = ?"
		args = append(args, *req.Sort)
	}
	if req.Title != "" {
		updateSQL += ", title = ?"
		args = append(args, req.Title)
	}
	if req.Icon != "" {
		updateSQL += ", icon = ?"
		args = append(args, req.Icon)
	}
	if req.LinkType != "" {
		updateSQL += ", link_type = ?"
		args = append(args, req.LinkType)
	}
	if req.LinkValue != "" {
		updateSQL += ", link_value = ?"
		args = append(args, req.LinkValue)
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

type DiamondDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDiamondDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DiamondDeleteLogic {
	return &DiamondDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DiamondDeleteLogic) DiamondDelete(id uint) (interface{}, error) {
	deleteSQL := "UPDATE diamond SET deleted_at = ? WHERE id = ? AND tenant_id = ?"
	result := l.svcCtx.DB.Exec(deleteSQL, time.Now(), id, 1)
	if result.Error != nil {
		return nil, result.Error
	}

	return map[string]interface{}{"message": "删除成功"}, nil
}

// ==================== 推荐位相关 ====================

type RecommendListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRecommendListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RecommendListLogic {
	return &RecommendListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RecommendListLogic) RecommendList(req *types.RecommendListRequest) (*types.RecommendListResponse, error) {
	querySQL := `
		SELECT id, channel_id, title, display_type, source_type, content_ids, filter_rule, sort, status, description
		FROM recommend
		WHERE tenant_id = ? AND channel_id = ? AND deleted_at IS NULL
	`
	args := []interface{}{1, req.ChannelID}

	if req.Status != nil {
		querySQL += " AND status = ?"
		args = append(args, *req.Status)
	}

	querySQL += " ORDER BY sort ASC"

	type RecommendResult struct {
		ID          uint
		ChannelID   uint
		Title       string
		DisplayType string
		SourceType  string
		ContentIDs  []byte
		FilterRule  []byte
		Sort        int
		Status      int8
		Description string
	}

	var results []RecommendResult
	err := l.svcCtx.DB.Raw(querySQL, args...).Scan(&results).Error
	if err != nil {
		return nil, err
	}

	list := make([]types.Recommend, 0, len(results))
	for _, r := range results {
		var contentIDs []uint
		var filterRule map[string]interface{}

		if r.ContentIDs != nil {
			json.Unmarshal(r.ContentIDs, &contentIDs)
		}
		if r.FilterRule != nil {
			json.Unmarshal(r.FilterRule, &filterRule)
		}

		// 查询物料详情
		var materials []types.Material
		
		// 优先使用 content_ids
		if len(contentIDs) > 0 {
			materialSQL := `
				SELECT id, title, subtitle, type, cover_url, content_url, description,
					   author, category, view_count, like_count, comment_count, 
					   share_count, collect_count, duration, word_count, chapter_count, status, sort
				FROM material
				WHERE id IN (?) AND deleted_at IS NULL
			`
			l.svcCtx.DB.Raw(materialSQL, contentIDs).Scan(&materials)
		} else if len(filterRule) > 0 {
			// 根据 filter_rule 动态查询
			materialSQL := `
				SELECT id, title, subtitle, type, cover_url, content_url, description,
					   author, category, view_count, like_count, comment_count, 
					   share_count, collect_count, duration, word_count, chapter_count, status, sort
				FROM material
				WHERE deleted_at IS NULL AND status = 1
			`
			materialArgs := []interface{}{}
			
			// 按类型筛选
			if types, ok := filterRule["types"].([]interface{}); ok && len(types) > 0 {
				materialSQL += " AND type IN (?)"
				typeStrs := make([]string, len(types))
				for i, t := range types {
					typeStrs[i] = t.(string)
				}
				materialArgs = append(materialArgs, typeStrs)
			}
			
			// 按分类筛选
			if category, ok := filterRule["category"].(string); ok && category != "" {
				materialSQL += " AND category = ?"
				materialArgs = append(materialArgs, category)
			}
			
			// 排序
			if orderBy, ok := filterRule["order_by"].(string); ok && orderBy != "" {
				switch orderBy {
				case "view_count":
					materialSQL += " ORDER BY view_count DESC"
				case "like_count":
					materialSQL += " ORDER BY like_count DESC"
				case "created_at":
					materialSQL += " ORDER BY created_at DESC"
				default:
					materialSQL += " ORDER BY id DESC"
				}
			} else {
				materialSQL += " ORDER BY id DESC"
			}
			
			// 限制数量
			if limit, ok := filterRule["limit"].(float64); ok && limit > 0 {
				materialSQL += " LIMIT ?"
				materialArgs = append(materialArgs, int(limit))
			} else {
				materialSQL += " LIMIT 10"
			}
			
			l.svcCtx.DB.Raw(materialSQL, materialArgs...).Scan(&materials)
		}

		list = append(list, types.Recommend{
			ID:          r.ID,
			ChannelID:   r.ChannelID,
			Title:       r.Title,
			DisplayType: r.DisplayType,
			SourceType:  r.SourceType,
			ContentIDs:  contentIDs,
			FilterRule:  filterRule,
			Sort:        r.Sort,
			Status:      r.Status,
			Description: r.Description,
			Materials:   materials,
		})
	}

	return &types.RecommendListResponse{List: list}, nil
}

type RecommendCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRecommendCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RecommendCreateLogic {
	return &RecommendCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RecommendCreateLogic) RecommendCreate(req *types.RecommendCreateRequest) (interface{}, error) {
	now := time.Now()

	contentIDsJSON, _ := json.Marshal(req.ContentIDs)
	filterRuleJSON, _ := json.Marshal(req.FilterRule)

	insertSQL := `
		INSERT INTO recommend (tenant_id, channel_id, title, display_type, source_type, content_ids, filter_rule, sort, status, description, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, 1, ?, ?, ?)
	`

	result := l.svcCtx.DB.Exec(insertSQL, 1, req.ChannelID, req.Title, req.DisplayType, req.SourceType, contentIDsJSON, filterRuleJSON, req.Sort, req.Description, now, now)
	if result.Error != nil {
		return nil, result.Error
	}

	return map[string]interface{}{"message": "创建成功"}, nil
}

type RecommendUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRecommendUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RecommendUpdateLogic {
	return &RecommendUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RecommendUpdateLogic) RecommendUpdate(req *types.RecommendUpdateRequest) (interface{}, error) {
	updateSQL := "UPDATE recommend SET updated_at = ?"
	args := []interface{}{time.Now()}

	if req.Title != "" {
		updateSQL += ", title = ?"
		args = append(args, req.Title)
	}
	if req.DisplayType != "" {
		updateSQL += ", display_type = ?"
		args = append(args, req.DisplayType)
	}
	if req.SourceType != "" {
		updateSQL += ", source_type = ?"
		args = append(args, req.SourceType)
	}
	if req.ContentIDs != nil {
		contentIDsJSON, _ := json.Marshal(req.ContentIDs)
		updateSQL += ", content_ids = ?"
		args = append(args, contentIDsJSON)
	}
	if req.FilterRule != nil {
		filterRuleJSON, _ := json.Marshal(req.FilterRule)
		updateSQL += ", filter_rule = ?"
		args = append(args, filterRuleJSON)
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

type RecommendDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRecommendDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RecommendDeleteLogic {
	return &RecommendDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RecommendDeleteLogic) RecommendDelete(id uint) (interface{}, error) {
	deleteSQL := "UPDATE recommend SET deleted_at = ? WHERE id = ? AND tenant_id = ?"
	result := l.svcCtx.DB.Exec(deleteSQL, time.Now(), id, 1)
	if result.Error != nil {
		return nil, result.Error
	}

	return map[string]interface{}{"message": "删除成功"}, nil
}

// ==================== Feed流相关 ====================

type FeedConfigListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFeedConfigListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FeedConfigListLogic {
	return &FeedConfigListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FeedConfigListLogic) FeedConfigList(req *types.FeedConfigListRequest) (*types.FeedConfigListResponse, error) {
	querySQL := `
		SELECT id, channel_id, title, layout_type, content_strategy, content_ids, filter_rule, sort, status, description
		FROM feed_config
		WHERE tenant_id = ? AND channel_id = ? AND deleted_at IS NULL
	`
	args := []interface{}{1, req.ChannelID}

	if req.Status != nil {
		querySQL += " AND status = ?"
		args = append(args, *req.Status)
	}

	querySQL += " ORDER BY sort ASC"

	type FeedConfigResult struct {
		ID              uint
		ChannelID       uint
		Title           string
		LayoutType      string
		ContentStrategy string
		ContentIDs      []byte
		FilterRule      []byte
		Sort            int
		Status          int8
		Description     string
	}

	var results []FeedConfigResult
	err := l.svcCtx.DB.Raw(querySQL, args...).Scan(&results).Error
	if err != nil {
		return nil, err
	}

	list := make([]types.FeedConfig, 0, len(results))
	for _, f := range results {
		var contentIDs []uint
		var filterRule map[string]interface{}

		if f.ContentIDs != nil {
			json.Unmarshal(f.ContentIDs, &contentIDs)
		}
		if f.FilterRule != nil {
			json.Unmarshal(f.FilterRule, &filterRule)
		}

		// 查询物料详情
		var materials []types.Material
		
		// 优先使用 content_ids
		if len(contentIDs) > 0 {
			materialSQL := `
				SELECT id, title, subtitle, type, cover_url, content_url, description,
					   author, category, view_count, like_count, comment_count, 
					   share_count, collect_count, duration, word_count, chapter_count, status, sort
				FROM material
				WHERE id IN (?) AND deleted_at IS NULL
			`
			l.svcCtx.DB.Raw(materialSQL, contentIDs).Scan(&materials)
		} else if len(filterRule) > 0 {
			// 根据 filter_rule 动态查询
			materialSQL := `
				SELECT id, title, subtitle, type, cover_url, content_url, description,
					   author, category, view_count, like_count, comment_count, 
					   share_count, collect_count, duration, word_count, chapter_count, status, sort
				FROM material
				WHERE deleted_at IS NULL AND status = 1
			`
			materialArgs := []interface{}{}
			
			// 按类型筛选
			if types, ok := filterRule["types"].([]interface{}); ok && len(types) > 0 {
				materialSQL += " AND type IN (?)"
				typeStrs := make([]string, len(types))
				for i, t := range types {
					typeStrs[i] = t.(string)
				}
				materialArgs = append(materialArgs, typeStrs)
			}
			
			// 排序
			if orderBy, ok := filterRule["order_by"].(string); ok && orderBy != "" {
				switch orderBy {
				case "view_count":
					materialSQL += " ORDER BY view_count DESC"
				case "like_count":
					materialSQL += " ORDER BY like_count DESC"
				case "created_at":
					materialSQL += " ORDER BY created_at DESC"
				default:
					materialSQL += " ORDER BY id DESC"
				}
			} else {
				materialSQL += " ORDER BY id DESC"
			}
			
			// 限制数量
			if limit, ok := filterRule["limit"].(float64); ok && limit > 0 {
				materialSQL += " LIMIT ?"
				materialArgs = append(materialArgs, int(limit))
			} else {
				materialSQL += " LIMIT 20"
			}
			
			l.svcCtx.DB.Raw(materialSQL, materialArgs...).Scan(&materials)
		}

		list = append(list, types.FeedConfig{
			ID:              f.ID,
			ChannelID:       f.ChannelID,
			Title:           f.Title,
			LayoutType:      f.LayoutType,
			ContentStrategy: f.ContentStrategy,
			ContentIDs:      contentIDs,
			FilterRule:      filterRule,
			Sort:            f.Sort,
			Status:          f.Status,
			Description:     f.Description,
			Materials:       materials,
		})
	}

	return &types.FeedConfigListResponse{List: list}, nil
}

type FeedConfigCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFeedConfigCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FeedConfigCreateLogic {
	return &FeedConfigCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FeedConfigCreateLogic) FeedConfigCreate(req *types.FeedConfigCreateRequest) (interface{}, error) {
	now := time.Now()

	contentIDsJSON, _ := json.Marshal(req.ContentIDs)
	filterRuleJSON, _ := json.Marshal(req.FilterRule)

	insertSQL := `
		INSERT INTO feed_config (tenant_id, channel_id, title, layout_type, content_strategy, content_ids, filter_rule, sort, status, description, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, 1, ?, ?, ?)
	`

	result := l.svcCtx.DB.Exec(insertSQL, 1, req.ChannelID, req.Title, req.LayoutType, req.ContentStrategy, contentIDsJSON, filterRuleJSON, req.Sort, req.Description, now, now)
	if result.Error != nil {
		return nil, result.Error
	}

	return map[string]interface{}{"message": "创建成功"}, nil
}

type FeedConfigUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFeedConfigUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FeedConfigUpdateLogic {
	return &FeedConfigUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FeedConfigUpdateLogic) FeedConfigUpdate(req *types.FeedConfigUpdateRequest) (interface{}, error) {
	updateSQL := "UPDATE feed_config SET updated_at = ?"
	args := []interface{}{time.Now()}

	if req.Title != "" {
		updateSQL += ", title = ?"
		args = append(args, req.Title)
	}
	if req.LayoutType != "" {
		updateSQL += ", layout_type = ?"
		args = append(args, req.LayoutType)
	}
	if req.ContentStrategy != "" {
		updateSQL += ", content_strategy = ?"
		args = append(args, req.ContentStrategy)
	}
	if req.ContentIDs != nil {
		contentIDsJSON, _ := json.Marshal(req.ContentIDs)
		updateSQL += ", content_ids = ?"
		args = append(args, contentIDsJSON)
	}
	if req.FilterRule != nil {
		filterRuleJSON, _ := json.Marshal(req.FilterRule)
		updateSQL += ", filter_rule = ?"
		args = append(args, filterRuleJSON)
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

type FeedConfigDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFeedConfigDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FeedConfigDeleteLogic {
	return &FeedConfigDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FeedConfigDeleteLogic) FeedConfigDelete(id uint) (interface{}, error) {
	deleteSQL := "UPDATE feed_config SET deleted_at = ? WHERE id = ? AND tenant_id = ?"
	result := l.svcCtx.DB.Exec(deleteSQL, time.Now(), id, 1)
	if result.Error != nil {
		return nil, result.Error
	}

	return map[string]interface{}{"message": "删除成功"}, nil
}
