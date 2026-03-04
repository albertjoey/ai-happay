package model

import (
	"time"

	"gorm.io/gorm"
)

// Tenant 租户/站点模型
type Tenant struct {
	ID          uint           `gorm:"primarykey" json:"id"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
	Name        string         `gorm:"type:varchar(100);uniqueIndex;not null" json:"name"`           // 站点名称
	Code        string         `gorm:"type:varchar(50);uniqueIndex;not null" json:"code"`            // 站点代码
	Domain      string         `gorm:"type:varchar(255);uniqueIndex;not null" json:"domain"`         // 站点域名
	Description string         `gorm:"type:text" json:"description"`                                 // 站点描述
	Logo        string         `gorm:"type:varchar(255)" json:"logo"`                                // 站点Logo
	Status      int8           `gorm:"type:tinyint;default:1;index" json:"status"`                   // 状态: 0-禁用 1-启用
	Config      string         `gorm:"type:json" json:"config"`                                      // 站点配置JSON
}

// User 用户模型
type User struct {
	ID           uint           `gorm:"primarykey" json:"id"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
	TenantID     uint           `gorm:"index;not null" json:"tenant_id"`                              // 租户ID
	Username     string         `gorm:"type:varchar(50);index" json:"username"`                       // 用户名
	Email        string         `gorm:"type:varchar(100);index" json:"email"`                         // 邮箱
	Phone        string         `gorm:"type:varchar(20);index" json:"phone"`                          // 手机号
	Password     string         `gorm:"type:varchar(255)" json:"-"`                                   // 密码
	Nickname     string         `gorm:"type:varchar(50)" json:"nickname"`                             // 昵称
	Avatar       string         `gorm:"type:varchar(255)" json:"avatar"`                              // 头像
	Gender       int8           `gorm:"type:tinyint;default:0" json:"gender"`                         // 性别: 0-未知 1-男 2-女
	Birthday     *time.Time     `json:"birthday"`                                                     // 生日
	Bio          string         `gorm:"type:text" json:"bio"`                                         // 个人简介
	Status       int8           `gorm:"type:tinyint;default:1;index" json:"status"`                   // 状态: 0-禁用 1-启用
	Role         int8           `gorm:"type:tinyint;default:0;index" json:"role"`                     // 角色: 0-普通用户 1-博主 2-管理员
	FollowCount  int            `gorm:"type:int;default:0" json:"follow_count"`                       // 关注数
	FansCount    int            `gorm:"type:int;default:0" json:"fans_count"`                         // 粉丝数
	LikeCount    int            `gorm:"type:int;default:0" json:"like_count"`                        // 获赞数
	ThirdPartyID string         `gorm:"type:varchar(100);index" json:"third_party_id"`                // 第三方登录ID
	ThirdType    string         `gorm:"type:varchar(20);index" json:"third_type"`                     // 第三方登录类型: google/wechat/apple/x
}

// Content 内容模型
type Content struct {
	ID          uint           `gorm:"primarykey" json:"id"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
	TenantID    uint           `gorm:"index;not null" json:"tenant_id"`                              // 租户ID
	UserID      uint           `gorm:"index;not null" json:"user_id"`                                // 创建者ID
	Title       string         `gorm:"type:varchar(255);index" json:"title"`                         // 标题
	Description string         `gorm:"type:text" json:"description"`                                 // 描述
	Cover       string         `gorm:"type:varchar(255)" json:"cover"`                               // 封面
	Type        int8           `gorm:"type:tinyint;index;not null" json:"type"`                      // 类型: 1-长视频 2-短视频 3-短剧 4-漫剧 5-小说 6-图文
	Status      int8           `gorm:"type:tinyint;default:0;index" json:"status"`                   // 状态: 0-草稿 1-已发布 2-已下架 3-审核中
	ViewCount   int            `gorm:"type:int;default:0" json:"view_count"`                         // 浏览数
	LikeCount   int            `gorm:"type:int;default:0" json:"like_count"`                         // 点赞数
	CommentCount int           `gorm:"type:int;default:0" json:"comment_count"`                      // 评论数
	CollectCount int           `gorm:"type:int;default:0" json:"collect_count"`                      // 收藏数
	ShareCount  int            `gorm:"type:int;default:0" json:"share_count"`                        // 分享数
	Media       string         `gorm:"type:json" json:"media"`                                       // 媒体资源JSON
	Extra       string         `gorm:"type:json" json:"extra"`                                       // 扩展信息JSON
	PublishAt   *time.Time     `json:"publish_at"`                                                   // 发布时间
}

// ContentMedia 内容媒体资源
type ContentMedia struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	ContentID uint      `gorm:"index;not null" json:"content_id"`                             // 内容ID
	Type      int8      `gorm:"type:tinyint;not null" json:"type"`                            // 类型: 1-图片 2-视频 3-音频
	URL       string    `gorm:"type:varchar(255);not null" json:"url"`                        // 资源URL
	Thumbnail string    `gorm:"type:varchar(255)" json:"thumbnail"`                           // 缩略图
	Duration  int       `gorm:"type:int;default:0" json:"duration"`                           // 时长(秒)
	Width     int       `gorm:"type:int;default:0" json:"width"`                             // 宽度
	Height    int       `gorm:"type:int;default:0" json:"height"`                            // 高度
	Size      int64     `gorm:"type:bigint;default:0" json:"size"`                           // 文件大小
	Sort      int       `gorm:"type:int;default:0" json:"sort"`                              // 排序
}

// Topic 话题模型
type Topic struct {
	ID          uint           `gorm:"primarykey" json:"id"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
	TenantID    uint           `gorm:"index;not null" json:"tenant_id"`                              // 租户ID
	Name        string         `gorm:"type:varchar(100);index;not null" json:"name"`                 // 话题名称
	Description string         `gorm:"type:text" json:"description"`                                 // 话题描述
	Cover       string         `gorm:"type:varchar(255)" json:"cover"`                               // 话题封面
	Status      int8           `gorm:"type:tinyint;default:1;index" json:"status"`                   // 状态: 0-禁用 1-启用
	Sort        int            `gorm:"type:int;default:0" json:"sort"`                               // 排序
}

// Tag 标签模型
type Tag struct {
	ID          uint           `gorm:"primarykey" json:"id"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
	TenantID    uint           `gorm:"index;not null" json:"tenant_id"`                              // 租户ID
	Name        string         `gorm:"type:varchar(50);index;not null" json:"name"`                  // 标签名称
	Type        int8           `gorm:"type:tinyint;default:0;index" json:"type"`                     // 类型: 0-通用 1-长视频 2-短视频 3-短剧 4-漫剧 5-小说 6-图文
	Status      int8           `gorm:"type:tinyint;default:1;index" json:"status"`                   // 状态: 0-禁用 1-启用
	Sort        int            `gorm:"type:int;default:0" json:"sort"`                               // 排序
}

// ContentTopic 内容话题关联
type ContentTopic struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	ContentID uint      `gorm:"index;not null" json:"content_id"`                             // 内容ID
	TopicID   uint      `gorm:"index;not null" json:"topic_id"`                               // 话题ID
}

// ContentTag 内容标签关联
type ContentTag struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	ContentID uint      `gorm:"index;not null" json:"content_id"`                             // 内容ID
	TagID     uint      `gorm:"index;not null" json:"tag_id"`                                 // 标签ID
}

// Channel 频道模型
type Channel struct {
	ID          uint           `gorm:"primarykey" json:"id"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
	TenantID    uint           `gorm:"index;not null" json:"tenant_id"`                              // 租户ID
	Name        string         `gorm:"type:varchar(50);index;not null" json:"name"`                  // 频道名称
	Code        string         `gorm:"type:varchar(50);index;not null" json:"code"`                  // 频道代码
	Description string         `gorm:"type:text" json:"description"`                                 // 频道描述
	Icon        string         `gorm:"type:varchar(255)" json:"icon"`                                // 频道图标
	Status      int8           `gorm:"type:tinyint;default:1;index" json:"status"`                   // 状态: 0-禁用 1-启用
	Sort        int            `gorm:"type:int;default:0" json:"sort"`                               // 排序
}

// Banner Banner模型
type Banner struct {
	ID          uint           `gorm:"primarykey" json:"id"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
	TenantID    uint           `gorm:"index;not null" json:"tenant_id"`                              // 租户ID
	ChannelID   uint           `gorm:"index" json:"channel_id"`                                      // 频道ID
	Title       string         `gorm:"type:varchar(100)" json:"title"`                               // 标题
	Image       string         `gorm:"type:varchar(255);not null" json:"image"`                      // 图片URL
	LinkType    int8           `gorm:"type:tinyint;default:1" json:"link_type"`                      // 链接类型: 1-内容 2-外链
	LinkURL     string         `gorm:"type:varchar(255)" json:"link_url"`                            // 链接地址
	ContentID   uint           `json:"content_id"`                                                   // 内容ID
	Status      int8           `gorm:"type:tinyint;default:1;index" json:"status"`                   // 状态: 0-禁用 1-启用
	Sort        int            `gorm:"type:int;default:0" json:"sort"`                               // 排序
	StartTime   *time.Time     `json:"start_time"`                                                   // 开始时间
	EndTime     *time.Time     `json:"end_time"`                                                     // 结束时间
}

// DiamondPosition 金刚位模型
type DiamondPosition struct {
	ID          uint           `gorm:"primarykey" json:"id"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
	TenantID    uint           `gorm:"index;not null" json:"tenant_id"`                              // 租户ID
	ChannelID   uint           `gorm:"index" json:"channel_id"`                                      // 频道ID
	Name        string         `gorm:"type:varchar(50);not null" json:"name"`                        // 名称
	Icon        string         `gorm:"type:varchar(255);not null" json:"icon"`                       // 图标
	LinkType    int8           `gorm:"type:tinyint;default:1" json:"link_type"`                      // 链接类型: 1-频道 2-话题 3-外链
	LinkURL     string         `gorm:"type:varchar(255)" json:"link_url"`                            // 链接地址
	ChannelLink uint           `json:"channel_link"`                                                 // 频道链接ID
	TopicLink   uint           `json:"topic_link"`                                                   // 话题链接ID
	Status      int8           `gorm:"type:tinyint;default:1;index" json:"status"`                   // 状态: 0-禁用 1-启用
	Sort        int            `gorm:"type:int;default:0" json:"sort"`                               // 排序
}

// FeedConfig Feed流配置
type FeedConfig struct {
	ID          uint           `gorm:"primarykey" json:"id"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
	TenantID    uint           `gorm:"index;not null" json:"tenant_id"`                              // 租户ID
	ChannelID   uint           `gorm:"index;not null" json:"channel_id"`                             // 频道ID
	Name        string         `gorm:"type:varchar(50);not null" json:"name"`                        // 配置名称
	Layout      int8           `gorm:"type:tinyint;default:1" json:"layout"`                         // 布局: 1-一行两列 2-一行三列
	Strategy    int8           `gorm:"type:tinyint;default:1" json:"strategy"`                       // 策略: 1-算法推荐 2-人工推荐 3-随机
	Status      int8           `gorm:"type:tinyint;default:1;index" json:"status"`                   // 状态: 0-禁用 1-启用
	Sort        int            `gorm:"type:int;default:0" json:"sort"`                               // 排序
}

// AdPosition 广告位模型
type AdPosition struct {
	ID          uint           `gorm:"primarykey" json:"id"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
	TenantID    uint           `gorm:"index;not null" json:"tenant_id"`                              // 租户ID
	ChannelID   uint           `gorm:"index" json:"channel_id"`                                      // 频道ID
	Name        string         `gorm:"type:varchar(50);not null" json:"name"`                        // 广告位名称
	Code        string         `gorm:"type:varchar(50);index;not null" json:"code"`                  // 广告位代码
	Type        int8           `gorm:"type:tinyint;default:1" json:"type"`                           // 类型: 1-图片 2-视频
	ImageURL    string         `gorm:"type:varchar(255)" json:"image_url"`                           // 图片URL
	VideoURL    string         `gorm:"type:varchar(255)" json:"video_url"`                           // 视频URL
	LinkURL     string         `gorm:"type:varchar(255)" json:"link_url"`                            // 跳转链接
	Status      int8           `gorm:"type:tinyint;default:1;index" json:"status"`                   // 状态: 0-禁用 1-启用
	Sort        int            `gorm:"type:int;default:0" json:"sort"`                               // 排序
	StartTime   *time.Time     `json:"start_time"`                                                   // 开始时间
	EndTime     *time.Time     `json:"end_time"`                                                     // 结束时间
}

// Interaction 互动模型
type Interaction struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UserID    uint      `gorm:"index;not null" json:"user_id"`                                // 用户ID
	TargetID  uint      `gorm:"index;not null" json:"target_id"`                              // 目标ID
	Type      int8      `gorm:"type:tinyint;index;not null" json:"type"`                      // 类型: 1-点赞 2-收藏 3-分享
	TargetType int8     `gorm:"type:tinyint;index;not null" json:"target_type"`               // 目标类型: 1-内容 2-评论
}

// Comment 评论模型
type Comment struct {
	ID          uint           `gorm:"primarykey" json:"id"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
	ContentID   uint           `gorm:"index;not null" json:"content_id"`                             // 内容ID
	UserID      uint           `gorm:"index;not null" json:"user_id"`                                // 用户ID
	ParentID    uint           `gorm:"index;default:0" json:"parent_id"`                             // 父评论ID
	ReplyUserID uint           `json:"reply_user_id"`                                                // 回复用户ID
	Content     string         `gorm:"type:text;not null" json:"content"`                            // 评论内容
	LikeCount   int            `gorm:"type:int;default:0" json:"like_count"`                         // 点赞数
	Status      int8           `gorm:"type:tinyint;default:1;index" json:"status"`                   // 状态: 0-删除 1-正常
}

// Follow 关注关系
type Follow struct {
	ID         uint      `gorm:"primarykey" json:"id"`
	CreatedAt  time.Time `json:"created_at"`
	FollowerID uint      `gorm:"index;not null" json:"follower_id"`                            // 粉丝ID
	FollowingID uint     `gorm:"index;not null" json:"following_id"`                           // 关注者ID
}

// ViewHistory 观看历史
type ViewHistory struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UserID    uint      `gorm:"index;not null" json:"user_id"`                                // 用户ID
	ContentID uint      `gorm:"index;not null" json:"content_id"`                             // 内容ID
	Duration  int       `gorm:"type:int;default:0" json:"duration"`                           // 观看时长
	Progress  int       `gorm:"type:int;default:0" json:"progress"`                           // 观看进度
}

// Message 消息通知
type Message struct {
	ID          uint           `gorm:"primarykey" json:"id"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
	UserID      uint           `gorm:"index;not null" json:"user_id"`                                // 接收用户ID
	FromUserID  uint           `json:"from_user_id"`                                                 // 发送用户ID
	Type        int8           `gorm:"type:tinyint;index;not null" json:"type"`                      // 类型: 1-点赞 2-评论 3-收藏 4-关注 5-系统
	Title       string         `gorm:"type:varchar(100)" json:"title"`                               // 标题
	Content     string         `gorm:"type:text" json:"content"`                                     // 内容
	TargetID    uint           `json:"target_id"`                                                    // 目标ID
	TargetType  int8           `json:"target_type"`                                                  // 目标类型
	IsRead      int8           `gorm:"type:tinyint;default:0;index" json:"is_read"`                  // 是否已读: 0-未读 1-已读
}

// AdminUser 管理员用户
type AdminUser struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	Username  string         `gorm:"type:varchar(50);uniqueIndex;not null" json:"username"`        // 用户名
	Password  string         `gorm:"type:varchar(255);not null" json:"-"`                          // 密码
	Nickname  string         `gorm:"type:varchar(50)" json:"nickname"`                             // 昵称
	Avatar    string         `gorm:"type:varchar(255)" json:"avatar"`                              // 头像
	Email     string         `gorm:"type:varchar(100)" json:"email"`                               // 邮箱
	Phone     string         `gorm:"type:varchar(20)" json:"phone"`                                // 手机号
	Status    int8           `gorm:"type:tinyint;default:1;index" json:"status"`                   // 状态: 0-禁用 1-启用
	RoleID    uint           `gorm:"index" json:"role_id"`                                         // 角色ID
}

// Role 角色
type Role struct {
	ID          uint           `gorm:"primarykey" json:"id"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
	Name        string         `gorm:"type:varchar(50);not null" json:"name"`                        // 角色名称
	Code        string         `gorm:"type:varchar(50);uniqueIndex;not null" json:"code"`            // 角色代码
	Description string         `gorm:"type:text" json:"description"`                                 // 描述
	Status      int8           `gorm:"type:tinyint;default:1;index" json:"status"`                   // 状态: 0-禁用 1-启用
}

// Permission 权限
type Permission struct {
	ID          uint           `gorm:"primarykey" json:"id"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
	Name        string         `gorm:"type:varchar(50);not null" json:"name"`                        // 权限名称
	Code        string         `gorm:"type:varchar(100);uniqueIndex;not null" json:"code"`           // 权限代码
	Type        int8           `gorm:"type:tinyint;not null" json:"type"`                            // 类型: 1-菜单 2-按钮
	ParentID    uint           `gorm:"index;default:0" json:"parent_id"`                             // 父权限ID
	Path        string         `gorm:"type:varchar(255)" json:"path"`                                // 路由路径
	Icon        string         `gorm:"type:varchar(50)" json:"icon"`                                 // 图标
	Sort        int            `gorm:"type:int;default:0" json:"sort"`                               // 排序
	Status      int8           `gorm:"type:tinyint;default:1;index" json:"status"`                   // 状态: 0-禁用 1-启用
}

// RolePermission 角色权限关联
type RolePermission struct {
	ID           uint      `gorm:"primarykey" json:"id"`
	CreatedAt    time.Time `json:"created_at"`
	RoleID       uint      `gorm:"index;not null" json:"role_id"`                                // 角色ID
	PermissionID uint      `gorm:"index;not null" json:"permission_id"`                          // 权限ID
}
