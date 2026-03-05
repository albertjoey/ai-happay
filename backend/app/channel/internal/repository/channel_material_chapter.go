package repository

import (
	"context"
	"happy/app/channel/internal/types"
	"time"

	"gorm.io/gorm"
)

// ==================== 频道仓储实现 ====================

type channelRepository struct {
	db *gorm.DB
}

// NewChannelRepository 创建频道仓储
func NewChannelRepository(db *gorm.DB) ChannelRepository {
	return &channelRepository{db: db}
}

func (r *channelRepository) List(ctx context.Context, req *types.ChannelListRequest) ([]types.Channel, int64, error) {
	var total int64
	var list []types.Channel

	// 默认租户ID
	tenantID := uint(1)

	// 查询总数
	countSQL := "SELECT COUNT(*) FROM channel WHERE tenant_id = ? AND deleted_at IS NULL"
	r.db.WithContext(ctx).Raw(countSQL, tenantID).Scan(&total)

	// 构建查询
	querySQL := `
		SELECT id, name, code, description, icon, status, sort
		FROM channel
		WHERE tenant_id = ? AND deleted_at IS NULL
	`
	args := []interface{}{tenantID}

	if req.Name != "" {
		querySQL += " AND name LIKE ?"
		args = append(args, "%"+req.Name+"%")
	}
	if req.Status != nil {
		querySQL += " AND status = ?"
		args = append(args, *req.Status)
	}

	querySQL += " ORDER BY sort ASC, id ASC LIMIT ? OFFSET ?"
	offset := (req.Page - 1) * req.PageSize
	args = append(args, req.PageSize, offset)

	err := r.db.WithContext(ctx).Raw(querySQL, args...).Scan(&list).Error
	return list, total, err
}

func (r *channelRepository) FindByID(ctx context.Context, id uint) (*types.Channel, error) {
	var channel types.Channel
	err := r.db.WithContext(ctx).Raw(`
		SELECT id, name, code, description, icon, status, sort
		FROM channel
		WHERE id = ? AND deleted_at IS NULL
	`, id).Scan(&channel).Error
	if err != nil {
		return nil, err
	}
	return &channel, nil
}

func (r *channelRepository) Create(ctx context.Context, channel *types.Channel) error {
	now := time.Now()
	return r.db.WithContext(ctx).Exec(`
		INSERT INTO channel (tenant_id, name, code, description, icon, status, sort, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
	`, 1, channel.Name, channel.Code, channel.Description, channel.Icon, channel.Status, channel.Sort, now, now).Error
}

func (r *channelRepository) Update(ctx context.Context, channel *types.Channel) error {
	return r.db.WithContext(ctx).Exec(`
		UPDATE channel SET name = ?, description = ?, icon = ?, status = ?, sort = ?, updated_at = ?
		WHERE id = ? AND deleted_at IS NULL
	`, channel.Name, channel.Description, channel.Icon, channel.Status, channel.Sort, time.Now(), channel.ID).Error
}

func (r *channelRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Exec("UPDATE channel SET deleted_at = ? WHERE id = ?", time.Now(), id).Error
}

func (r *channelRepository) UpdateSort(ctx context.Context, id uint, sort int) error {
	return r.db.WithContext(ctx).Exec("UPDATE channel SET sort = ?, updated_at = ? WHERE id = ?", sort, time.Now(), id).Error
}

// ==================== 物料仓储实现 ====================

type materialRepository struct {
	db *gorm.DB
}

// NewMaterialRepository 创建物料仓储
func NewMaterialRepository(db *gorm.DB) MaterialRepository {
	return &materialRepository{db: db}
}

func (r *materialRepository) List(ctx context.Context, req *types.MaterialListRequest) ([]types.Material, int64, error) {
	var total int64
	var list []types.Material

	// 查询总数
	countSQL := "SELECT COUNT(*) FROM material WHERE deleted_at IS NULL"
	countArgs := []interface{}{}

	if req.Type != "" {
		countSQL += " AND type = ?"
		countArgs = append(countArgs, req.Type)
	}
	if req.Status != nil {
		countSQL += " AND status = ?"
		countArgs = append(countArgs, *req.Status)
	}

	r.db.WithContext(ctx).Raw(countSQL, countArgs...).Scan(&total)

	// 查询列表
	page := req.Page
	if page < 1 {
		page = 1
	}
	pageSize := req.PageSize
	if pageSize < 1 {
		pageSize = 10
	}
	if pageSize > 100 {
		pageSize = 100
	}
	offset := (page - 1) * pageSize

	listSQL := `
		SELECT id, title, subtitle, type, cover_url, content_url, description,
			   author, category, view_count, like_count, comment_count, 
			   share_count, collect_count, duration, word_count, chapter_count, status, sort
		FROM material
		WHERE deleted_at IS NULL
	`
	listArgs := []interface{}{}

	if req.Type != "" {
		listSQL += " AND type = ?"
		listArgs = append(listArgs, req.Type)
	}
	if req.Status != nil {
		listSQL += " AND status = ?"
		listArgs = append(listArgs, *req.Status)
	}

	listSQL += " ORDER BY id DESC LIMIT ? OFFSET ?"
	listArgs = append(listArgs, pageSize, offset)

	err := r.db.WithContext(ctx).Raw(listSQL, listArgs...).Scan(&list).Error
	return list, total, err
}

func (r *materialRepository) FindByID(ctx context.Context, id uint) (*types.Material, error) {
	var material types.Material
	err := r.db.WithContext(ctx).Raw(`
		SELECT id, title, subtitle, type, cover_url, content_url, description,
			   author, category, view_count, like_count, comment_count, 
			   share_count, collect_count, duration, word_count, chapter_count, status, sort
		FROM material
		WHERE id = ? AND deleted_at IS NULL
	`, id).Scan(&material).Error
	if err != nil {
		return nil, err
	}
	return &material, nil
}

func (r *materialRepository) Create(ctx context.Context, material *types.Material) error {
	now := time.Now()
	return r.db.WithContext(ctx).Exec(`
		INSERT INTO material (title, subtitle, type, cover_url, content_url, description,
			author, category, duration, word_count, status, sort, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`, material.Title, material.Subtitle, material.Type, material.CoverURL, material.ContentURL, material.Description,
		material.Author, material.Category, material.Duration, material.WordCount, material.Status, material.Sort, now, now).Error
}

func (r *materialRepository) Update(ctx context.Context, material *types.Material) error {
	return r.db.WithContext(ctx).Exec(`
		UPDATE material SET title = ?, subtitle = ?, description = ?, status = ?, updated_at = ?
		WHERE id = ? AND deleted_at IS NULL
	`, material.Title, material.Subtitle, material.Description, material.Status, time.Now(), material.ID).Error
}

func (r *materialRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Exec("UPDATE material SET deleted_at = ? WHERE id = ?", time.Now(), id).Error
}

// ==================== 章节仓储实现 ====================

type chapterRepository struct {
	db *gorm.DB
}

// NewChapterRepository 创建章节仓储
func NewChapterRepository(db *gorm.DB) ChapterRepository {
	return &chapterRepository{db: db}
}

func (r *chapterRepository) List(ctx context.Context, materialID uint, page, pageSize int) ([]types.Chapter, int64, error) {
	var total int64
	var list []types.Chapter

	query := "SELECT COUNT(*) FROM material_chapter WHERE material_id = ? AND deleted_at IS NULL"
	r.db.WithContext(ctx).Raw(query, materialID).Scan(&total)

	offset := (page - 1) * pageSize
	listQuery := `
		SELECT id, material_id, chapter_type, title, content, images, video_url, 
			   word_count, duration, sort, is_free, price
		FROM material_chapter
		WHERE material_id = ? AND deleted_at IS NULL
		ORDER BY sort ASC
		LIMIT ? OFFSET ?
	`
	err := r.db.WithContext(ctx).Raw(listQuery, materialID, pageSize, offset).Scan(&list).Error
	return list, total, err
}

func (r *chapterRepository) FindByID(ctx context.Context, id uint) (*types.Chapter, error) {
	var chapter types.Chapter
	err := r.db.WithContext(ctx).Raw(`
		SELECT id, material_id, chapter_type, title, content, images, video_url, 
			   word_count, duration, sort, is_free, price
		FROM material_chapter
		WHERE id = ? AND deleted_at IS NULL
	`, id).Scan(&chapter).Error
	if err != nil {
		return nil, err
	}
	return &chapter, nil
}

func (r *chapterRepository) FindPrevNext(ctx context.Context, chapterID uint) (prevID, nextID uint, err error) {
	// 先获取当前章节
	var chapter types.Chapter
	err = r.db.WithContext(ctx).Raw(`
		SELECT id, material_id, sort FROM material_chapter WHERE id = ?
	`, chapterID).Scan(&chapter).Error
	if err != nil {
		return 0, 0, err
	}

	// 获取上一章
	r.db.WithContext(ctx).Raw(`
		SELECT id FROM material_chapter 
		WHERE material_id = ? AND sort < ? AND deleted_at IS NULL 
		ORDER BY sort DESC LIMIT 1
	`, chapter.MaterialID, chapter.Sort).Scan(&prevID)

	// 获取下一章
	r.db.WithContext(ctx).Raw(`
		SELECT id FROM material_chapter 
		WHERE material_id = ? AND sort > ? AND deleted_at IS NULL 
		ORDER BY sort ASC LIMIT 1
	`, chapter.MaterialID, chapter.Sort).Scan(&nextID)

	return prevID, nextID, nil
}

func (r *chapterRepository) Create(ctx context.Context, chapter *types.Chapter) error {
	now := time.Now()
	return r.db.WithContext(ctx).Exec(`
		INSERT INTO material_chapter (material_id, chapter_type, title, content, images, video_url, 
			word_count, duration, sort, is_free, price, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`, chapter.MaterialID, chapter.ChapterType, chapter.Title, chapter.Content, chapter.Images, chapter.VideoURL,
		chapter.WordCount, chapter.Duration, chapter.Sort, chapter.IsFree, chapter.Price, now, now).Error
}

func (r *chapterRepository) Update(ctx context.Context, chapter *types.Chapter) error {
	return r.db.WithContext(ctx).Exec(`
		UPDATE material_chapter SET chapter_type = ?, title = ?, content = ?, images = ?, video_url = ?,
			word_count = ?, duration = ?, sort = ?, is_free = ?, price = ?, updated_at = ?
		WHERE id = ? AND deleted_at IS NULL
	`, chapter.ChapterType, chapter.Title, chapter.Content, chapter.Images, chapter.VideoURL,
		chapter.WordCount, chapter.Duration, chapter.Sort, chapter.IsFree, chapter.Price, time.Now(), chapter.ID).Error
}

func (r *chapterRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Exec("UPDATE material_chapter SET deleted_at = ? WHERE id = ?", time.Now(), id).Error
}

func (r *chapterRepository) GetMaxSort(ctx context.Context, materialID uint) (int, error) {
	var maxSort int
	r.db.WithContext(ctx).Raw(`
		SELECT COALESCE(MAX(sort), 0) FROM material_chapter 
		WHERE material_id = ? AND deleted_at IS NULL
	`, materialID).Scan(&maxSort)
	return maxSort, nil
}
