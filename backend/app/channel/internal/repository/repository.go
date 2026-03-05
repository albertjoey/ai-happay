package repository

import (
	"context"
	"happy/app/channel/internal/types"
)

// ==================== 频道相关 ====================

// ChannelRepository 频道仓储接口
type ChannelRepository interface {
	// List 获取频道列表
	List(ctx context.Context, req *types.ChannelListRequest) ([]types.Channel, int64, error)
	//FindByID 根据ID获取频道
	FindByID(ctx context.Context, id uint) (*types.Channel, error)
	// Create 创建频道
	Create(ctx context.Context, channel *types.Channel) error
	// Update 更新频道
	Update(ctx context.Context, channel *types.Channel) error
	// Delete 删除频道
	Delete(ctx context.Context, id uint) error
	// UpdateSort 更新排序
	UpdateSort(ctx context.Context, id uint, sort int) error
}

// ==================== 物料相关 ====================

// MaterialRepository 物料仓储接口
type MaterialRepository interface {
	// List 获取物料列表
	List(ctx context.Context, req *types.MaterialListRequest) ([]types.Material, int64, error)
	// FindByID 根据ID获取物料
	FindByID(ctx context.Context, id uint) (*types.Material, error)
	// Create 创建物料
	Create(ctx context.Context, material *types.Material) error
	// Update 更新物料
	Update(ctx context.Context, material *types.Material) error
	// Delete 删除物料
	Delete(ctx context.Context, id uint) error
}

// ==================== 章节相关 ====================

// ChapterRepository 章节仓储接口
type ChapterRepository interface {
	// List 获取章节列表
	List(ctx context.Context, materialID uint, page, pageSize int) ([]types.Chapter, int64, error)
	// FindByID 根据ID获取章节
	FindByID(ctx context.Context, id uint) (*types.Chapter, error)
	// FindPrevNext 获取上一章/下一章ID
	FindPrevNext(ctx context.Context, chapterID uint) (prevID, nextID uint, err error)
	// Create 创建章节
	Create(ctx context.Context, chapter *types.Chapter) error
	// Update 更新章节
	Update(ctx context.Context, chapter *types.Chapter) error
	// Delete 删除章节
	Delete(ctx context.Context, id uint) error
	// GetMaxSort 获取最大排序值
	GetMaxSort(ctx context.Context, materialID uint) (int, error)
}

// ==================== 角色相关 ====================

// RoleRepository 角色仓储接口
type RoleRepository interface {
	// List 获取角色列表
	List(ctx context.Context, req *types.RoleListRequest) ([]types.Role, int64, error)
	// FindByID 根据ID获取角色
	FindByID(ctx context.Context, id uint) (*types.Role, error)
	// Create 创建角色
	Create(ctx context.Context, role *types.Role) error
	// Update 更新角色
	Update(ctx context.Context, role *types.Role) error
	// Delete 删除角色
	Delete(ctx context.Context, id uint) error
	// GetPermissions 获取角色权限
	GetPermissions(ctx context.Context, roleID uint) ([]uint, error)
	// AssignPermissions 分配角色权限
	AssignPermissions(ctx context.Context, roleID uint, permissionIDs []uint) error
}

// ==================== 权限相关 ====================

// PermissionRepository 权限仓储接口
type PermissionRepository interface {
	// List 获取权限列表
	List(ctx context.Context) ([]types.Permission, error)
	// Create 创建权限
	Create(ctx context.Context, permission *types.Permission) error
	// Update 更新权限
	Update(ctx context.Context, permission *types.Permission) error
	// Delete 删除权限
	Delete(ctx context.Context, id uint) error
}

// ==================== 管理员相关 ====================

// AdminUserRepository 管理员仓储接口
type AdminUserRepository interface {
	// List 获取管理员列表
	List(ctx context.Context, req *types.AdminUserListRequest) ([]types.AdminUser, int64, error)
	// FindByID 根据ID获取管理员
	FindByID(ctx context.Context, id uint) (*types.AdminUser, error)
	// FindByUsername 根据用户名获取管理员
	FindByUsername(ctx context.Context, username string) (*types.AdminUser, error)
	// Create 创建管理员
	Create(ctx context.Context, admin *types.AdminUser) error
	// Update 更新管理员
	Update(ctx context.Context, admin *types.AdminUser) error
	// Delete 删除管理员
	Delete(ctx context.Context, id uint) error
	// GetRoles 获取管理员角色
	GetRoles(ctx context.Context, adminUserID uint) ([]uint, error)
	// AssignRoles 分配管理员角色
	AssignRoles(ctx context.Context, adminUserID uint, roleIDs []uint) error
}

// ==================== Banner相关 ====================

// BannerRepository Banner仓储接口
type BannerRepository interface {
	// List 获取Banner列表
	List(ctx context.Context, channelID uint) ([]types.Banner, error)
}

// ==================== 金刚位相关 ====================

// DiamondRepository 金刚位仓储接口
type DiamondRepository interface {
	// List 获取金刚位列表
	List(ctx context.Context, channelID uint) ([]types.Diamond, error)
	// Create 创建金刚位
	Create(ctx context.Context, diamond *types.Diamond) error
}

// ==================== 推荐位相关 ====================

// RecommendRepository 推荐位仓储接口
type RecommendRepository interface {
	// GetByChannelID 根据频道ID获取推荐配置
	GetByChannelID(ctx context.Context, channelID uint) (*types.RecommendConfig, error)
	// Preview 预览推荐内容
	Preview(ctx context.Context, recommendID uint) ([]types.Material, error)
}

// ==================== 广告位相关 ====================

// AdSlotRepository 广告位仓储接口
type AdSlotRepository interface {
	// List 获取广告位列表
	List(ctx context.Context, channelID uint) ([]types.AdSlot, error)
	// Create 创建广告位
	Create(ctx context.Context, adSlot *types.AdSlot) error
	// Update 更新广告位
	Update(ctx context.Context, adSlot *types.AdSlot) error
	// Delete 删除广告位
	Delete(ctx context.Context, id uint) error
}

// ==================== Feed流相关 ====================

// FeedConfigRepository Feed配置仓储接口
type FeedConfigRepository interface {
	// GetByChannelID 根据频道ID获取Feed配置
	GetByChannelID(ctx context.Context, channelID uint) (*types.FeedConfig, error)
}

// ==================== 互动相关 ====================

// InteractionRepository 互动仓储接口
type InteractionRepository interface {
	// Like 点赞
	Like(ctx context.Context, userID, materialID uint) error
	// Unlike 取消点赞
	Unlike(ctx context.Context, userID, materialID uint) error
	// Collect 收藏
	Collect(ctx context.Context, userID, materialID uint) error
	// Uncollect 取消收藏
	Uncollect(ctx context.Context, userID, materialID uint) error
	// Comment 评论
	Comment(ctx context.Context, userID, materialID uint, content string, parentID uint) error
	// GetComments 获取评论列表
	GetComments(ctx context.Context, materialID uint, page, pageSize int) ([]types.Comment, int64, error)
	// IsLiked 是否已点赞
	IsLiked(ctx context.Context, userID, materialID uint) (bool, error)
	// IsCollected 是否已收藏
	IsCollected(ctx context.Context, userID, materialID uint) (bool, error)
}

// ==================== 话题相关 ====================

// TopicRepository 话题仓储接口
type TopicRepository interface {
	// List 获取话题列表
	List(ctx context.Context, req *types.TopicListRequest) ([]types.Topic, int64, error)
	// FindByID 根据ID获取话题
	FindByID(ctx context.Context, id uint) (*types.Topic, error)
	// Create 创建话题
	Create(ctx context.Context, topic *types.Topic) error
	// Update 更新话题
	Update(ctx context.Context, topic *types.Topic) error
	// Delete 删除话题
	Delete(ctx context.Context, id uint) error
}

// ==================== 标签相关 ====================

// TagRepository 标签仓储接口
type TagRepository interface {
	// List 获取标签列表
	List(ctx context.Context, req *types.TagListRequest) ([]types.Tag, int64, error)
	// FindByID 根据ID获取标签
	FindByID(ctx context.Context, id uint) (*types.Tag, error)
	// Create 创建标签
	Create(ctx context.Context, tag *types.Tag) error
	// Update 更新标签
	Update(ctx context.Context, tag *types.Tag) error
	// Delete 删除标签
	Delete(ctx context.Context, id uint) error
}
