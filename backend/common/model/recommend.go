package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"

	"gorm.io/gorm"
)

// Recommend 推荐位模型
type Recommend struct {
	ID          uint           `gorm:"primarykey" json:"id"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
	TenantID    uint           `gorm:"not null;index" json:"tenant_id"`
	ChannelID   uint           `gorm:"not null;index" json:"channel_id"`
	Title       string         `gorm:"size:100;not null" json:"title"`          // 标题
	DisplayType string         `gorm:"size:20;not null" json:"display_type"`    // 展示类型: single/scroll/grid
	SourceType  string         `gorm:"size:20;not null" json:"source_type"`     // 内容来源: manual/algorithm/filter
	ContentIDs  JSONArray      `gorm:"type:json" json:"content_ids"`            // 内容ID列表(人工选择)
	FilterRule  JSONB          `gorm:"type:json" json:"filter_rule"`            // 筛选规则(条件筛选)
	Sort        int            `gorm:"not null;default:0" json:"sort"`          // 排序
	Status      int8           `gorm:"not null;default:1" json:"status"`        // 状态: 0-禁用 1-启用
	Description string         `gorm:"size:200" json:"description"`             // 描述
}

func (Recommend) TableName() string {
	return "recommend"
}

// JSONArray 用于存储JSON数组
type JSONArray []uint

func (j JSONArray) Value() (driver.Value, error) {
	return json.Marshal(j)
}

func (j *JSONArray) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(bytes, j)
}

// JSONB 用于存储JSON对象
type JSONB map[string]interface{}

func (j JSONB) Value() (driver.Value, error) {
	return json.Marshal(j)
}

func (j *JSONB) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(bytes, j)
}

// StringArray 用于存储字符串数组
type StringArray []string

func (j StringArray) Value() (driver.Value, error) {
	return json.Marshal(j)
}

func (j *StringArray) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(bytes, j)
}
