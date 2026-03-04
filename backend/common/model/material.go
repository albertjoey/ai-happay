package model

import (
	"time"

	"gorm.io/gorm"
)

// Material 物料模型
type Material struct {
	ID            uint           `gorm:"primarykey" json:"id"`
	TenantID      uint           `gorm:"not null;index" json:"tenant_id"`
	Title         string         `gorm:"size:200;not null" json:"title"`
	Subtitle      string         `gorm:"size:500" json:"subtitle"`
	Type          string         `gorm:"size:20;not null;index" json:"type"` // image_text, novel, video, banner, comic, short_drama
	CoverURL      string         `gorm:"size:500" json:"cover_url"`
	ContentURL    string         `gorm:"size:500" json:"content_url"`
	Description   string         `gorm:"type:text" json:"description"`
	Author        string         `gorm:"size:100" json:"author"`
	Tags          StringArray    `gorm:"type:json" json:"tags"`
	Category      string         `gorm:"size:50" json:"category"`
	ViewCount     uint           `gorm:"default:0" json:"view_count"`
	LikeCount     uint           `gorm:"default:0" json:"like_count"`
	CommentCount  uint           `gorm:"default:0" json:"comment_count"`
	ShareCount    uint           `gorm:"default:0" json:"share_count"`
	CollectCount  uint           `gorm:"default:0" json:"collect_count"`
	Duration      uint           `gorm:"default:0" json:"duration"`      // 时长(秒)
	WordCount     uint           `gorm:"default:0" json:"word_count"`    // 字数
	ChapterCount  uint           `gorm:"default:0" json:"chapter_count"` // 章节数
	Status        int8           `gorm:"not null;default:1;index" json:"status"` // 0-草稿, 1-已发布, 2-已下架
	Sort          int            `gorm:"default:0" json:"sort"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

// TableName 指定表名
func (Material) TableName() string {
	return "material"
}
