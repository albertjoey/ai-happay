package model

import (
	"time"

	"gorm.io/gorm"
)

// AdSlot 广告位模型
type AdSlot struct {
	ID           uint           `gorm:"primarykey" json:"id"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
	TenantID     uint           `gorm:"not null;index" json:"tenant_id"`
	ChannelID    uint           `gorm:"not null;index" json:"channel_id"`
	Name         string         `gorm:"size:100;not null" json:"name"`              // 广告位名称
	InsertType   string         `gorm:"size:20;not null" json:"insert_type"`        // 插入方式: fixed/random/position
	InsertRule   string         `gorm:"type:text" json:"insert_rule"`               // 插入规则(JSON)
	AdType       string         `gorm:"size:20;not null" json:"ad_type"`            // 广告类型: image/video
	AdContent    string         `gorm:"type:text" json:"ad_content"`                // 广告内容(JSON)
	LinkURL      string         `gorm:"size:500" json:"link_url"`                   // 跳转链接
	Status       int8           `gorm:"not null;default:1" json:"status"`           // 状态: 0-禁用 1-启用
	Sort         int            `gorm:"not null;default:0" json:"sort"`             // 排序
	Description  string         `gorm:"size:200" json:"description"`                // 描述
}

func (AdSlot) TableName() string {
	return "ad_slot"
}

// AdContent 广告内容结构
type AdContent struct {
	ImageURL string `json:"image_url"` // 图片地址
	VideoURL string `json:"video_url"` // 视频地址
	Title    string `json:"title"`     // 标题
	Duration int    `json:"duration"`  // 时长(秒)
}

// InsertRule 插入规则结构
type InsertRule struct {
	FixedPosition int `json:"fixed_position"` // 固定位置(第N条后插入)
	Interval      int `json:"interval"`       // 间隔(每N条插入一个)
	MaxCount      int `json:"max_count"`      // 最大插入数量
}
