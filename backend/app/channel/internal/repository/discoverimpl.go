package repository

import (
	"context"

	"happy/app/channel/internal/types"

	"gorm.io/gorm"
)

type discoverRepository struct {
	db *gorm.DB
}

func NewDiscoverRepository(db *gorm.DB) DiscoverRepository {
	return &discoverRepository{db: db}
}

// GetModuleList 获取模块配置列表
func (r *discoverRepository) GetModuleList(ctx context.Context) ([]types.DiscoverConfig, error) {
	var configs []types.DiscoverConfig
	err := r.db.WithContext(ctx).Raw(`
		SELECT id, module_type as module, title, sort as sort_order, status as is_enabled,
		       created_at, updated_at
		FROM discover_module
		WHERE tenant_id = 1 AND deleted_at IS NULL
		ORDER BY sort ASC
	`).Scan(&configs).Error
	return configs, err
}

// UpdateModule 更新模块配置
func (r *discoverRepository) UpdateModule(ctx context.Context, req *types.DiscoverConfigUpdateRequest) error {
	status := 1
	if req.IsEnabled != nil && !*req.IsEnabled {
		status = 0
	}

	sortOrder := 0
	if req.SortOrder != nil {
		sortOrder = *req.SortOrder
	}

	return r.db.WithContext(ctx).Exec(`
		UPDATE discover_module
		SET title = ?, sort = ?, status = ?, updated_at = NOW()
		WHERE id = ? AND tenant_id = 1
	`, req.Title, sortOrder, status, req.ID).Error
}

// GetContentList 获取内容列表
func (r *discoverRepository) GetContentList(ctx context.Context, module string, page, pageSize int) ([]types.DiscoverItem, int64, error) {
	var items []types.DiscoverItem
	var total int64

	query := `
		SELECT dc.id, dc.module_id as config_id, dm.module_type as module,
		       dc.content_type as item_type, dc.content_id as item_id,
		       dc.sort as sort_order, dc.status as is_enabled,
		       dc.created_at, dc.updated_at
		FROM discover_content dc
		LEFT JOIN discover_module dm ON dc.module_id = dm.id
		WHERE dc.tenant_id = 1 AND dc.deleted_at IS NULL
	`

	args := []interface{}{}

	if module != "" {
		query += " AND dm.module_type = ?"
		args = append(args, module)
	}

	// 获取总数
	countQuery := "SELECT COUNT(*) FROM (" + query + ") t"
	r.db.WithContext(ctx).Raw(countQuery, args...).Scan(&total)

	// 获取列表
	query += " ORDER BY dc.sort ASC"
	if pageSize > 0 {
		offset := (page - 1) * pageSize
		query += " LIMIT ? OFFSET ?"
		args = append(args, pageSize, offset)
	}

	err := r.db.WithContext(ctx).Raw(query, args...).Scan(&items).Error
	return items, total, err
}

// CreateContent 创建内容
func (r *discoverRepository) CreateContent(ctx context.Context, moduleID uint, contentType string, contentID uint, sort int) error {
	return r.db.WithContext(ctx).Exec(`
		INSERT INTO discover_content (tenant_id, module_id, content_type, content_id, sort, status, created_at)
		VALUES (1, ?, ?, ?, ?, 1, NOW())
	`, moduleID, contentType, contentID, sort).Error
}

// UpdateContent 更新内容
func (r *discoverRepository) UpdateContent(ctx context.Context, id uint, sort int, status int8) error {
	return r.db.WithContext(ctx).Exec(`
		UPDATE discover_content
		SET sort = ?, status = ?, updated_at = NOW()
		WHERE id = ? AND tenant_id = 1
	`, sort, status, id).Error
}

// DeleteContent 删除内容
func (r *discoverRepository) DeleteContent(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Exec(`
		UPDATE discover_content
		SET deleted_at = NOW()
		WHERE id = ? AND tenant_id = 1
	`, id).Error
}

// GetDiscoverModules 获取发现页模块
func (r *discoverRepository) GetDiscoverModules(ctx context.Context) ([]map[string]interface{}, error) {
	var modules []struct {
		ID         uint
		ModuleType string
		Title      string
		Sort       int
		Status     int8
		Config     string
	}

	err := r.db.WithContext(ctx).Raw(`
		SELECT id, module_type, title, sort, status, config
		FROM discover_module
		WHERE tenant_id = 1 AND status = 1 AND deleted_at IS NULL
		ORDER BY sort ASC
	`).Scan(&modules).Error

	if err != nil {
		return nil, err
	}

	result := make([]map[string]interface{}, len(modules))
	for i, module := range modules {
		result[i] = map[string]interface{}{
			"id":     module.ID,
			"module": module.ModuleType,
			"title":  module.Title,
		}
	}

	return result, nil
}

// GetModuleContents 获取模块内容
func (r *discoverRepository) GetModuleContents(ctx context.Context, moduleID uint, moduleType string) ([]map[string]interface{}, error) {
	var contents []struct {
		ContentID uint
		Sort      int
	}

	err := r.db.WithContext(ctx).Raw(`
		SELECT content_id, sort
		FROM discover_content
		WHERE module_id = ? AND status = 1 AND deleted_at IS NULL
		ORDER BY sort ASC
	`, moduleID).Scan(&contents).Error

	if err != nil {
		return nil, err
	}

	// 根据模块类型获取内容详情
	switch moduleType {
	case "hot_topics":
		return r.getTopicContents(ctx, contents)
	case "hot_rank":
		return r.getMaterialContents(ctx, contents, true)
	case "guess_you_like":
		return r.getMaterialContents(ctx, contents, false)
	default:
		return []map[string]interface{}{}, nil
	}
}

// getTopicContents 获取话题内容
func (r *discoverRepository) getTopicContents(ctx context.Context, contents []struct {
	ContentID uint
	Sort      int
}) ([]map[string]interface{}, error) {
	var topics []map[string]interface{}
	for _, c := range contents {
		var topic struct {
			ID    uint
			Name  string
			Cover string
		}

		err := r.db.WithContext(ctx).Raw(`
			SELECT id, name, cover
			FROM topic
			WHERE id = ? AND deleted_at IS NULL
		`, c.ContentID).Scan(&topic).Error

		if err == nil {
			topics = append(topics, map[string]interface{}{
				"id":         topic.ID,
				"title":      topic.Name,
				"cover_url":  topic.Cover,
				"view_count": 0,
				"post_count": 0,
			})
		}
	}
	return topics, nil
}

// getMaterialContents 获取物料内容
func (r *discoverRepository) getMaterialContents(ctx context.Context, contents []struct {
	ContentID uint
	Sort      int
}, withRank bool) ([]map[string]interface{}, error) {
	var materials []map[string]interface{}
	for i, c := range contents {
		var material struct {
			ID        uint
			Title     string
			CoverURL  string
			Author    string
			ViewCount int
			Type      string
		}

		err := r.db.WithContext(ctx).Raw(`
			SELECT id, title, cover_url, author, view_count, type
			FROM material
			WHERE id = ? AND deleted_at IS NULL
		`, c.ContentID).Scan(&material).Error

		if err == nil {
			item := map[string]interface{}{
				"id":         material.ID,
				"title":      material.Title,
				"cover_url":  material.CoverURL,
				"author":     material.Author,
				"view_count": material.ViewCount,
			}

			if withRank {
				item["rank"] = i + 1
				item["hot_score"] = 100.0 - float64(i)*2.5
			} else {
				item["type"] = material.Type
			}

			materials = append(materials, item)
		}
	}
	return materials, nil
}

// GetModuleIDByType 根据模块类型获取模块ID
func (r *discoverRepository) GetModuleIDByType(ctx context.Context, moduleType string) (uint, error) {
	var moduleID uint
	err := r.db.WithContext(ctx).Raw(`
		SELECT id FROM discover_module
		WHERE module_type = ? AND tenant_id = 1 AND deleted_at IS NULL
	`, moduleType).Scan(&moduleID).Error
	return moduleID, err
}
