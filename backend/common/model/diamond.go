package model

import (
	"time"

	"gorm.io/gorm"
)

// Diamond 金刚位模型
type Diamond struct {
	ID          uint           `gorm:"primarykey" json:"id"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
	TenantID    uint           `gorm:"not null;index" json:"tenant_id"`
	ChannelID   uint           `gorm:"not null;index" json:"channel_id"`
	GroupID     int            `gorm:"not null;default:1" json:"group_id"`      // 分组ID (1-5组)
	Sort        int            `gorm:"not null;default:0" json:"sort"`          // 组内排序
	Title       string         `gorm:"size:50;not null" json:"title"`           // 标题
	Icon        string         `gorm:"size:255" json:"icon"`                    // 图标
	LinkType    string         `gorm:"size:20;not null" json:"link_type"`       // 链接类型: channel/topic/content/external
	LinkValue   string         `gorm:"size:500" json:"link_value"`              // 链接值
	Status      int8           `gorm:"not null;default:1" json:"status"`        // 状态: 0-禁用 1-启用
	Description string         `gorm:"size:200" json:"description"`             // 描述
}

func (Diamond) TableName() string {
	return "diamond"
}
