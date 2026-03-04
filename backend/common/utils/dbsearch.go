package utils

import (
	"happy/common/types"

	"gorm.io/gorm"
)

// DBSearch 数据库搜索（临时方案，后续切换到ES）
type DBSearch struct {
	db *gorm.DB
}

// NewDBSearch 创建数据库搜索
func NewDBSearch(db *gorm.DB) *DBSearch {
	return &DBSearch{db: db}
}

// Search 搜索内容
func (s *DBSearch) Search(keyword string, tenantID uint, page, pageSize int) (*types.SearchResponse, error) {
	var total int64
	var results []types.SearchResult

	// 构建查询 - 使用原生SQL避免编码问题
	sql := `
		SELECT 
			c.id, c.title, c.description, c.cover, c.type, 
			c.view_count, c.like_count, 
			u.nickname as author_name, u.avatar as author_avatar
		FROM content c
		LEFT JOIN user u ON u.id = c.user_id
		WHERE c.tenant_id = ? AND c.status = 1
		AND (c.title LIKE ? OR c.description LIKE ?)
		ORDER BY (c.view_count + c.like_count * 2) DESC, c.created_at DESC
		LIMIT ? OFFSET ?
	`

	countSQL := `
		SELECT COUNT(*)
		FROM content
		WHERE tenant_id = ? AND status = 1
		AND (title LIKE ? OR description LIKE ?)
	`

	// 统计总数
	s.db.Raw(countSQL, tenantID, "%"+keyword+"%", "%"+keyword+"%").Scan(&total)

	// 分页查询
	offset := (page - 1) * pageSize
	err := s.db.Raw(sql, tenantID, "%"+keyword+"%", "%"+keyword+"%", pageSize, offset).Scan(&results).Error

	if err != nil {
		return nil, err
	}

	return &types.SearchResponse{
		Total: total,
		List:  results,
	}, nil
}

// Suggest 搜索建议
func (s *DBSearch) Suggest(keyword string, tenantID uint, limit int) ([]string, error) {
	var suggestions []string

	sql := `
		SELECT title
		FROM content
		WHERE tenant_id = ? AND status = 1 AND title LIKE ?
		ORDER BY view_count DESC
		LIMIT ?
	`

	err := s.db.Raw(sql, tenantID, keyword+"%", limit).Scan(&suggestions).Error

	if err != nil {
		return nil, err
	}

	return suggestions, nil
}
