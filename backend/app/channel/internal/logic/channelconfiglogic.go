package logic

import (
	"context"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"happy/app/channel/internal/svc"
	"happy/app/channel/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChannelConfigLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChannelConfigLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChannelConfigLogic {
	return &ChannelConfigLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// ChannelConfig 频道配置模型
type ChannelConfig struct {
	ID           uint                   `json:"id"`
	ChannelID    uint                   `json:"channel_id"`
	TenantID     uint                   `json:"tenant_id"`
	ContentTypes []string               `json:"content_types"`
	DisplayMode  string                 `json:"display_mode"`
	CustomConfig map[string]interface{} `json:"custom_config"`
}

// ContentTypes JSON类型
type ContentTypes []string

func (c *ContentTypes) Value() (driver.Value, error) {
	return json.Marshal(c)
}

func (c *ContentTypes) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(bytes, c)
}

// CustomConfig JSON类型
type CustomConfig map[string]interface{}

func (c *CustomConfig) Value() (driver.Value, error) {
	return json.Marshal(c)
}

func (c *CustomConfig) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(bytes, c)
}

func (l *ChannelConfigLogic) GetChannelConfig(channelID uint) (*types.ChannelConfigResponse, error) {
	var config ChannelConfig

	query := `
		SELECT id, channel_id, tenant_id, content_types, display_mode, custom_config
		FROM channel_config
		WHERE channel_id = ? AND tenant_id = 1 AND deleted_at IS NULL
	`

	var contentTypesJSON, customConfigJSON string
	err := l.svcCtx.DB.Raw(query, channelID).Row().Scan(
		&config.ID,
		&config.ChannelID,
		&config.TenantID,
		&contentTypesJSON,
		&config.DisplayMode,
		&customConfigJSON,
	)

	if err != nil {
		// 如果没有配置，返回默认配置
		return &types.ChannelConfigResponse{
			ChannelID:   channelID,
			ContentType: map[string]bool{"video": true, "image": true, "article": true},
			DisplayMode: "default",
			CustomData:  map[string]interface{}{},
			PageConfig:  getDefaultPageConfig(),
		}, nil
	}

	// 解析JSON
	var contentTypes []string
	json.Unmarshal([]byte(contentTypesJSON), &contentTypes)

	var customConfig map[string]interface{}
	json.Unmarshal([]byte(customConfigJSON), &customConfig)

	// 转换content_types为map
	contentTypeMap := make(map[string]bool)
	for _, ct := range contentTypes {
		contentTypeMap[ct] = true
	}

	// 转换custom_config为map[string]interface{}
	customData := make(map[string]interface{})
	for k, v := range customConfig {
		customData[k] = v
	}

	// 解析页面配置
	pageConfig := parsePageConfig(customConfig)

	return &types.ChannelConfigResponse{
		ChannelID:   config.ChannelID,
		ContentType: contentTypeMap,
		DisplayMode: config.DisplayMode,
		CustomData:  customData,
		PageConfig:  pageConfig,
	}, nil
}

// getDefaultPageConfig 获取默认页面配置
func getDefaultPageConfig() *types.ChannelPageConfig {
	return &types.ChannelPageConfig{
		Banner: &types.BannerConfig{Enabled: true},
		Diamond: &types.DiamondConfig{Enabled: true},
		Recommends: []types.RecommendConfig{},
		Feed: &types.FeedConfigItem{Enabled: true, AutoLoad: true, ShowTitle: true},
	}
}

// parsePageConfig 从custom_config解析页面配置
func parsePageConfig(customConfig map[string]interface{}) *types.ChannelPageConfig {
	pageConfig := getDefaultPageConfig()
	
	if pageConfigRaw, ok := customConfig["page_config"]; ok {
		if pageConfigMap, ok := pageConfigRaw.(map[string]interface{}); ok {
			// Banner配置
			if bannerRaw, ok := pageConfigMap["banner"]; ok {
				if bannerMap, ok := bannerRaw.(map[string]interface{}); ok {
					if enabled, ok := bannerMap["enabled"].(bool); ok {
						pageConfig.Banner.Enabled = enabled
					}
					if bannerIDsRaw, ok := bannerMap["banner_ids"].([]interface{}); ok {
						for _, id := range bannerIDsRaw {
							if idFloat, ok := id.(float64); ok {
								pageConfig.Banner.BannerIDs = append(pageConfig.Banner.BannerIDs, uint(idFloat))
							}
						}
					}
				}
			}
			
			// Diamond配置
			if diamondRaw, ok := pageConfigMap["diamond"]; ok {
				if diamondMap, ok := diamondRaw.(map[string]interface{}); ok {
					if enabled, ok := diamondMap["enabled"].(bool); ok {
						pageConfig.Diamond.Enabled = enabled
					}
					if groupIDsRaw, ok := diamondMap["group_ids"].([]interface{}); ok {
						for _, id := range groupIDsRaw {
							if idFloat, ok := id.(float64); ok {
								pageConfig.Diamond.GroupIDs = append(pageConfig.Diamond.GroupIDs, int(idFloat))
							}
						}
					}
				}
			}
			
			// Recommends配置
			if recommendsRaw, ok := pageConfigMap["recommends"]; ok {
				if recommendsSlice, ok := recommendsRaw.([]interface{}); ok {
					pageConfig.Recommends = []types.RecommendConfig{}
					for _, recRaw := range recommendsSlice {
						if recMap, ok := recRaw.(map[string]interface{}); ok {
							rec := types.RecommendConfig{}
							if idFloat, ok := recMap["id"].(float64); ok {
								rec.ID = uint(idFloat)
							}
							if title, ok := recMap["title"].(string); ok {
								rec.Title = title
							}
							if sortFloat, ok := recMap["sort"].(float64); ok {
								rec.Sort = int(sortFloat)
							}
							pageConfig.Recommends = append(pageConfig.Recommends, rec)
						}
					}
				}
			}
			
			// Feed配置
			if feedRaw, ok := pageConfigMap["feed"]; ok {
				if feedMap, ok := feedRaw.(map[string]interface{}); ok {
					if enabled, ok := feedMap["enabled"].(bool); ok {
						pageConfig.Feed.Enabled = enabled
					}
					if feedIDFloat, ok := feedMap["feed_id"].(float64); ok {
						pageConfig.Feed.FeedID = uint(feedIDFloat)
					}
					if autoLoad, ok := feedMap["auto_load"].(bool); ok {
						pageConfig.Feed.AutoLoad = autoLoad
					}
					if showTitle, ok := feedMap["show_title"].(bool); ok {
						pageConfig.Feed.ShowTitle = showTitle
					}
				}
			}
		}
	}
	
	return pageConfig
}

func (l *ChannelConfigLogic) UpdateChannelConfig(channelID uint, req *types.ChannelConfigResponse) error {
	// 转换content_type为数组
	var contentTypes []string
	for k, v := range req.ContentType {
		if v {
			contentTypes = append(contentTypes, k)
		}
	}

	// 转换custom_data为interface map
	customConfig := make(map[string]interface{})
	for k, v := range req.CustomData {
		customConfig[k] = v
	}

	// 序列化JSON
	contentTypesJSON, _ := json.Marshal(contentTypes)
	customConfigJSON, _ := json.Marshal(customConfig)

	// 检查是否存在配置
	var count int64
	l.svcCtx.DB.Raw(
		"SELECT COUNT(*) FROM channel_config WHERE channel_id = ? AND tenant_id = 1 AND deleted_at IS NULL",
		channelID,
	).Scan(&count)

	if count > 0 {
		// 更新
		updateSQL := `
			UPDATE channel_config
			SET content_types = ?, display_mode = ?, custom_config = ?, updated_at = NOW()
			WHERE channel_id = ? AND tenant_id = 1 AND deleted_at IS NULL
		`
		return l.svcCtx.DB.Exec(updateSQL, contentTypesJSON, req.DisplayMode, customConfigJSON, channelID).Error
	} else {
		// 创建
		insertSQL := `
			INSERT INTO channel_config (channel_id, tenant_id, content_types, display_mode, custom_config, created_at, updated_at)
			VALUES (?, 1, ?, ?, ?, NOW(), NOW())
		`
		return l.svcCtx.DB.Exec(insertSQL, channelID, contentTypesJSON, req.DisplayMode, customConfigJSON).Error
	}
}
