package repository

import (
	"context"
	"happy/app/channel/internal/types"

	"gorm.io/gorm"
)

// ==================== Banner仓储实现 ====================

type bannerRepository struct {
	db *gorm.DB
}

// NewBannerRepository 创建Banner仓储
func NewBannerRepository(db *gorm.DB) BannerRepository {
	return &bannerRepository{db: db}
}

func (r *bannerRepository) List(ctx context.Context, channelID uint) ([]types.Banner, error) {
	var banners []types.Banner
	err := r.db.WithContext(ctx).Raw(`
		SELECT id, channel_id, title, image, link_url, link_type, sort, status
		FROM banner
		WHERE channel_id = ? AND deleted_at IS NULL
		ORDER BY sort ASC
	`, channelID).Scan(&banners).Error
	return banners, err
}

// ==================== 金刚位仓储实现 ====================

type diamondRepository struct {
	db *gorm.DB
}

// NewDiamondRepository 创建金刚位仓储
func NewDiamondRepository(db *gorm.DB) DiamondRepository {
	return &diamondRepository{db: db}
}

func (r *diamondRepository) List(ctx context.Context, channelID uint) ([]types.Diamond, error) {
	var diamonds []types.Diamond
	err := r.db.WithContext(ctx).Raw(`
		SELECT id, channel_id, group_id, sort, title, icon, link_type, link_value, status, description, material_id
		FROM diamond
		WHERE channel_id = ? AND deleted_at IS NULL
		ORDER BY group_id ASC, sort ASC
	`, channelID).Scan(&diamonds).Error
	return diamonds, err
}

func (r *diamondRepository) Create(ctx context.Context, diamond *types.Diamond) error {
	return r.db.WithContext(ctx).Exec(`
		INSERT INTO diamond (channel_id, group_id, title, icon, link_type, link_value, sort, status)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)
	`, diamond.ChannelID, diamond.GroupID, diamond.Title, diamond.Icon, diamond.LinkType, diamond.LinkValue, diamond.Sort, diamond.Status).Error
}

// ==================== 推荐位仓储实现 ====================

type recommendRepository struct {
	db *gorm.DB
}

// NewRecommendRepository 创建推荐位仓储
func NewRecommendRepository(db *gorm.DB) RecommendRepository {
	return &recommendRepository{db: db}
}

func (r *recommendRepository) GetByChannelID(ctx context.Context, channelID uint) (*types.RecommendConfig, error) {
	var config types.RecommendConfig
	err := r.db.WithContext(ctx).Raw(`
		SELECT id, channel_id, title, type, material_ids, sort, status
		FROM recommend_config
		WHERE channel_id = ? AND deleted_at IS NULL
	`, channelID).Scan(&config).Error
	if err != nil {
		return nil, err
	}
	return &config, nil
}

func (r *recommendRepository) Preview(ctx context.Context, recommendID uint) ([]types.Material, error) {
	var materials []types.Material
	
	// 先获取推荐配置中的物料ID
	var materialIDs string
	r.db.WithContext(ctx).Raw(`
		SELECT material_ids FROM recommend_config WHERE id = ?
	`, recommendID).Scan(&materialIDs)

	if materialIDs == "" {
		return materials, nil
	}

	// 查询物料详情
	err := r.db.WithContext(ctx).Raw(`
		SELECT id, title, subtitle, type, cover_url, description,
			   author, category, view_count, like_count, comment_count, 
			   share_count, collect_count, duration, word_count, chapter_count, status
		FROM material
		WHERE id IN (?) AND deleted_at IS NULL
	`, materialIDs).Scan(&materials).Error

	return materials, err
}

// ==================== 广告位仓储实现 ====================

type adSlotRepository struct {
	db *gorm.DB
}

// NewAdSlotRepository 创建广告位仓储
func NewAdSlotRepository(db *gorm.DB) AdSlotRepository {
	return &adSlotRepository{db: db}
}

func (r *adSlotRepository) List(ctx context.Context, channelID uint) ([]types.AdSlot, error) {
	var adSlots []types.AdSlot
	err := r.db.WithContext(ctx).Raw(`
		SELECT id, channel_id, name, insert_type, insert_rule, ad_type, ad_content, link_url, sort, description
		FROM ad_slot
		WHERE channel_id = ? AND deleted_at IS NULL
		ORDER BY sort ASC
	`, channelID).Scan(&adSlots).Error
	return adSlots, err
}

func (r *adSlotRepository) Create(ctx context.Context, adSlot *types.AdSlot) error {
	return r.db.WithContext(ctx).Exec(`
		INSERT INTO ad_slot (channel_id, name, insert_type, insert_rule, ad_type, ad_content, link_url, sort, description)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
	`, adSlot.ChannelID, adSlot.Name, adSlot.InsertType, adSlot.InsertRule, adSlot.AdType, adSlot.AdContent, adSlot.LinkURL, adSlot.Sort, adSlot.Description).Error
}

func (r *adSlotRepository) Update(ctx context.Context, adSlot *types.AdSlot) error {
	return r.db.WithContext(ctx).Exec(`
		UPDATE ad_slot SET name = ?, insert_type = ?, insert_rule = ?, ad_type = ?, ad_content = ?, link_url = ?, sort = ?, description = ?
		WHERE id = ? AND deleted_at IS NULL
	`, adSlot.Name, adSlot.InsertType, adSlot.InsertRule, adSlot.AdType, adSlot.AdContent, adSlot.LinkURL, adSlot.Sort, adSlot.Description, adSlot.ID).Error
}

func (r *adSlotRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Exec("UPDATE ad_slot SET deleted_at = NOW() WHERE id = ?", id).Error
}

// ==================== Feed配置仓储实现 ====================

type feedConfigRepository struct {
	db *gorm.DB
}

// NewFeedConfigRepository 创建Feed配置仓储
func NewFeedConfigRepository(db *gorm.DB) FeedConfigRepository {
	return &feedConfigRepository{db: db}
}

func (r *feedConfigRepository) GetByChannelID(ctx context.Context, channelID uint) (*types.FeedConfig, error) {
	var config types.FeedConfig
	err := r.db.WithContext(ctx).Raw(`
		SELECT id, channel_id, name, type, materials, sort, status
		FROM feed_config
		WHERE channel_id = ? AND deleted_at IS NULL
	`, channelID).Scan(&config).Error
	if err != nil {
		return nil, err
	}
	return &config, nil
}

// ==================== 互动仓储实现 ====================

type interactionRepository struct {
	db *gorm.DB
}

// NewInteractionRepository 创建互动仓储
func NewInteractionRepository(db *gorm.DB) InteractionRepository {
	return &interactionRepository{db: db}
}

func (r *interactionRepository) Like(ctx context.Context, userID, materialID uint) error {
	return r.db.WithContext(ctx).Exec(`
		INSERT INTO user_like (user_id, material_id, created_at)
		VALUES (?, ?, NOW())
		ON DUPLICATE KEY UPDATE deleted_at = NULL
	`, userID, materialID).Error
}

func (r *interactionRepository) Unlike(ctx context.Context, userID, materialID uint) error {
	return r.db.WithContext(ctx).Exec(`
		UPDATE user_like SET deleted_at = NOW()
		WHERE user_id = ? AND material_id = ?
	`, userID, materialID).Error
}

func (r *interactionRepository) Collect(ctx context.Context, userID, materialID uint) error {
	return r.db.WithContext(ctx).Exec(`
		INSERT INTO user_collect (user_id, material_id, created_at)
		VALUES (?, ?, NOW())
		ON DUPLICATE KEY UPDATE deleted_at = NULL
	`, userID, materialID).Error
}

func (r *interactionRepository) Uncollect(ctx context.Context, userID, materialID uint) error {
	return r.db.WithContext(ctx).Exec(`
		UPDATE user_collect SET deleted_at = NOW()
		WHERE user_id = ? AND material_id = ?
	`, userID, materialID).Error
}

func (r *interactionRepository) Comment(ctx context.Context, userID, materialID uint, content string, parentID uint) error {
	return r.db.WithContext(ctx).Exec(`
		INSERT INTO comment (user_id, material_id, content, parent_id, created_at)
		VALUES (?, ?, ?, ?, NOW())
	`, userID, materialID, content, parentID).Error
}

func (r *interactionRepository) GetComments(ctx context.Context, materialID uint, page, pageSize int) ([]types.Comment, int64, error) {
	var total int64
	var comments []types.Comment

	r.db.WithContext(ctx).Raw(`
		SELECT COUNT(*) FROM comment 
		WHERE material_id = ? AND parent_id = 0 AND deleted_at IS NULL
	`, materialID).Scan(&total)

	offset := (page - 1) * pageSize
	err := r.db.WithContext(ctx).Raw(`
		SELECT id, user_id, material_id, content, parent_id, like_count, created_at
		FROM comment
		WHERE material_id = ? AND parent_id = 0 AND deleted_at IS NULL
		ORDER BY created_at DESC
		LIMIT ? OFFSET ?
	`, materialID, pageSize, offset).Scan(&comments).Error

	return comments, total, err
}

func (r *interactionRepository) IsLiked(ctx context.Context, userID, materialID uint) (bool, error) {
	var count int64
	r.db.WithContext(ctx).Raw(`
		SELECT COUNT(*) FROM user_like
		WHERE user_id = ? AND material_id = ? AND deleted_at IS NULL
	`, userID, materialID).Scan(&count)
	return count > 0, nil
}

func (r *interactionRepository) IsCollected(ctx context.Context, userID, materialID uint) (bool, error) {
	var count int64
	r.db.WithContext(ctx).Raw(`
		SELECT COUNT(*) FROM user_collect
		WHERE user_id = ? AND material_id = ? AND deleted_at IS NULL
	`, userID, materialID).Scan(&count)
	return count > 0, nil
}
