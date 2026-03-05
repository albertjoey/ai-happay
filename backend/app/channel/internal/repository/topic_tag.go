package repository

import (
	"context"
	"happy/app/channel/internal/types"
	"time"

	"gorm.io/gorm"
)

// ==================== 话题仓储实现 ====================

type topicRepository struct {
	db *gorm.DB
}

// NewTopicRepository 创建话题仓储
func NewTopicRepository(db *gorm.DB) TopicRepository {
	return &topicRepository{db: db}
}

func (r *topicRepository) List(ctx context.Context, req *types.TopicListRequest) ([]types.Topic, int64, error) {
	var topics []types.Topic
	var total int64

	query := r.db.WithContext(ctx).Model(&types.Topic{}).Where("deleted_at IS NULL")

	if req.Name != "" {
		query = query.Where("name LIKE ?", "%"+req.Name+"%")
	}
	if req.Status != nil {
		query = query.Where("status = ?", *req.Status)
	}

	// 获取总数
	query.Count(&total)

	// 分页查询
	offset := (req.Page - 1) * req.PageSize
	err := query.Order("id DESC").Offset(offset).Limit(req.PageSize).Find(&topics).Error

	return topics, total, err
}

func (r *topicRepository) FindByID(ctx context.Context, id uint) (*types.Topic, error) {
	var topic types.Topic
	err := r.db.WithContext(ctx).Where("id = ? AND deleted_at IS NULL", id).First(&topic).Error
	if err != nil {
		return nil, err
	}
	return &topic, nil
}

func (r *topicRepository) Create(ctx context.Context, topic *types.Topic) error {
	now := time.Now()
	return r.db.WithContext(ctx).Exec(`
		INSERT INTO topic (tenant_id, name, description, cover, status, sort, created_at, updated_at)
		VALUES (1, ?, ?, ?, ?, ?, ?, ?)
	`, topic.Name, topic.Description, topic.Cover, topic.Status, topic.Sort, now, now).Error
}

func (r *topicRepository) Update(ctx context.Context, topic *types.Topic) error {
	now := time.Now()
	return r.db.WithContext(ctx).Exec(`
		UPDATE topic SET updated_at = ?, name = COALESCE(NULLIF(?, ''), name), 
		description = COALESCE(NULLIF(?, ''), description), cover = COALESCE(NULLIF(?, ''), cover),
		status = COALESCE(?, status), sort = COALESCE(?, sort)
		WHERE id = ? AND deleted_at IS NULL
	`, now, topic.Name, topic.Description, topic.Cover, topic.Status, topic.Sort, topic.ID).Error
}

func (r *topicRepository) Delete(ctx context.Context, id uint) error {
	now := time.Now()
	return r.db.WithContext(ctx).Exec(`
		UPDATE topic SET deleted_at = ? WHERE id = ? AND deleted_at IS NULL
	`, now, id).Error
}

// ==================== 标签仓储实现 ====================

type tagRepository struct {
	db *gorm.DB
}

// NewTagRepository 创建标签仓储
func NewTagRepository(db *gorm.DB) TagRepository {
	return &tagRepository{db: db}
}

func (r *tagRepository) List(ctx context.Context, req *types.TagListRequest) ([]types.Tag, int64, error) {
	var tags []types.Tag
	var total int64

	query := r.db.WithContext(ctx).Model(&types.Tag{}).Where("deleted_at IS NULL")

	if req.Name != "" {
		query = query.Where("name LIKE ?", "%"+req.Name+"%")
	}
	if req.Type != nil {
		query = query.Where("type = ?", *req.Type)
	}
	if req.Status != nil {
		query = query.Where("status = ?", *req.Status)
	}

	// 获取总数
	query.Count(&total)

	// 分页查询
	offset := (req.Page - 1) * req.PageSize
	err := query.Order("id DESC").Offset(offset).Limit(req.PageSize).Find(&tags).Error

	return tags, total, err
}

func (r *tagRepository) FindByID(ctx context.Context, id uint) (*types.Tag, error) {
	var tag types.Tag
	err := r.db.WithContext(ctx).Where("id = ? AND deleted_at IS NULL", id).First(&tag).Error
	if err != nil {
		return nil, err
	}
	return &tag, nil
}

func (r *tagRepository) Create(ctx context.Context, tag *types.Tag) error {
	now := time.Now()
	return r.db.WithContext(ctx).Exec(`
		INSERT INTO tag (tenant_id, name, type, status, sort, created_at, updated_at)
		VALUES (1, ?, ?, ?, ?, ?, ?)
	`, tag.Name, tag.Type, tag.Status, tag.Sort, now, now).Error
}

func (r *tagRepository) Update(ctx context.Context, tag *types.Tag) error {
	now := time.Now()
	return r.db.WithContext(ctx).Exec(`
		UPDATE tag SET updated_at = ?, name = COALESCE(NULLIF(?, ''), name), 
		type = COALESCE(?, type), status = COALESCE(?, status), sort = COALESCE(?, sort)
		WHERE id = ? AND deleted_at IS NULL
	`, now, tag.Name, tag.Type, tag.Status, tag.Sort, tag.ID).Error
}

func (r *tagRepository) Delete(ctx context.Context, id uint) error {
	now := time.Now()
	return r.db.WithContext(ctx).Exec(`
		UPDATE tag SET deleted_at = ? WHERE id = ? AND deleted_at IS NULL
	`, now, id).Error
}
