package types

// Channel 频道
type Channel struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Code        string `json:"code"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
	Status      int8   `json:"status"`
	Sort        int    `json:"sort"`
}

// ChannelListRequest 频道列表请求
type ChannelListRequest struct {
	Page     int    `form:"page,default=1"`
	PageSize int    `form:"page_size,default=20"`
	Name     string `form:"name,optional"`
	Status   *int8  `form:"status,optional"`
}

// ChannelListResponse 频道列表响应
type ChannelListResponse struct {
	Total int64     `json:"total"`
	List  []Channel `json:"list"`
}

// ChannelCreateRequest 创建频道请求
type ChannelCreateRequest struct {
	Name        string `json:"name"`
	Code        string `json:"code"`
	Description string `json:"description,optional"`
	Icon        string `json:"icon,optional"`
	Sort        int    `json:"sort,optional"`
}

// ChannelUpdateRequest 更新频道请求
type ChannelUpdateRequest struct {
	ID          uint   `path:"id"`
	Name        string `json:"name,optional"`
	Description string `json:"description,optional"`
	Icon        string `json:"icon,optional"`
	Status      *int8  `json:"status,optional"`
	Sort        *int   `json:"sort,optional"`
}

// ChannelSortRequest 频道排序请求
type ChannelSortRequest struct {
	Items []struct {
		ID   uint `json:"id"`
		Sort int  `json:"sort"`
	} `json:"items"`
}

// ChannelConfigRequest 频道配置请求
type ChannelConfigRequest struct {
	ChannelID uint `path:"id"`
}

// ChannelConfigResponse 频道配置响应
type ChannelConfigResponse struct {
	ChannelID   uint                   `json:"channel_id"`
	ContentType map[string]bool        `json:"content_type"`
	DisplayMode string                 `json:"display_mode"`
	CustomData  map[string]interface{} `json:"custom_data"`
	// 新增：频道页面组件配置
	PageConfig  *ChannelPageConfig     `json:"page_config,optional"`
}

// ChannelPageConfig 频道页面配置
type ChannelPageConfig struct {
	Banner      *BannerConfig      `json:"banner,optional"`       // Banner配置
	Diamond     *DiamondConfig     `json:"diamond,optional"`      // 金刚位配置
	Recommends  []RecommendConfig  `json:"recommends,optional"`   // 推荐位列表
	Feed        *FeedConfigItem    `json:"feed,optional"`         // Feed流配置
}

// BannerConfig Banner配置
type BannerConfig struct {
	Enabled  bool     `json:"enabled"`            // 是否启用
	BannerIDs []uint  `json:"banner_ids,optional"` // Banner ID列表
}

// DiamondConfig 金刚位配置
type DiamondConfig struct {
	Enabled    bool  `json:"enabled"`              // 是否启用
	GroupIDs   []int `json:"group_ids,optional"`   // 显示的金刚位分组ID
}

// RecommendConfig 推荐位配置
type RecommendConfig struct {
	ID        uint   `json:"id"`                   // 推荐位ID
	Title     string `json:"title,optional"`       // 标题（覆盖默认）
	Sort      int    `json:"sort,optional"`        // 排序
}

// FeedConfigItem Feed流配置
type FeedConfigItem struct {
	Enabled     bool   `json:"enabled"`               // 是否启用
	FeedID      uint   `json:"feed_id,optional"`      // Feed配置ID
	AutoLoad    bool   `json:"auto_load,optional"`    // 是否自动加载
	ShowTitle   bool   `json:"show_title,optional"`   // 是否显示标题
}

// ChannelConfigUpdateRequest 更新频道配置请求
type ChannelConfigUpdateRequest struct {
	ChannelID   uint                   `path:"id"`
	ContentType map[string]bool        `json:"content_type,optional"`
	DisplayMode string                 `json:"display_mode,optional"`
	CustomData  map[string]interface{} `json:"custom_data,optional"`
}

// ==================== 金刚位相关 ====================

// Diamond 金刚位
type Diamond struct {
	ID          uint                   `json:"id"`
	ChannelID   uint                   `json:"channel_id"`
	GroupID     int                    `json:"group_id"`
	Sort        int                    `json:"sort"`
	Title       string                 `json:"title"`
	Icon        string                 `json:"icon"`
	LinkType    string                 `json:"link_type"`
	LinkValue   string                 `json:"link_value"`
	Status      int8                   `json:"status"`
	Description string                 `json:"description"`
	MaterialID  uint                   `json:"material_id"`
	Material    map[string]interface{} `json:"material,omitempty"`
}

// DiamondListRequest 金刚位列表请求
type DiamondListRequest struct {
	ChannelID uint  `form:"channel_id"`
	GroupID   *int  `form:"group_id,optional"`
	Status    *int8 `form:"status,optional"`
}

// DiamondListResponse 金刚位列表响应
type DiamondListResponse struct {
	List []Diamond `json:"list"`
}

// DiamondCreateRequest 创建金刚位请求
type DiamondCreateRequest struct {
	ChannelID   uint   `json:"channel_id"`
	GroupID     int    `json:"group_id"`
	Sort        int    `json:"sort,optional"`
	Title       string `json:"title"`
	Icon        string `json:"icon,optional"`
	LinkType    string `json:"link_type"`
	LinkValue   string `json:"link_value"`
	Description string `json:"description,optional"`
}

// DiamondUpdateRequest 更新金刚位请求
type DiamondUpdateRequest struct {
	ID          uint   `path:"id"`
	GroupID     *int   `json:"group_id,optional"`
	Sort        *int   `json:"sort,optional"`
	Title       string `json:"title,optional"`
	Icon        string `json:"icon,optional"`
	LinkType    string `json:"link_type,optional"`
	LinkValue   string `json:"link_value,optional"`
	Status      *int8  `json:"status,optional"`
	Description string `json:"description,optional"`
}

// ==================== 推荐位相关 ====================

// Recommend 推荐位
type Recommend struct {
	ID          uint                   `json:"id"`
	ChannelID   uint                   `json:"channel_id"`
	Title       string                 `json:"title"`
	DisplayType string                 `json:"display_type"`
	SourceType  string                 `json:"source_type"`
	ContentIDs  []uint                 `json:"content_ids"`
	FilterRule  map[string]interface{} `json:"filter_rule"`
	Sort        int                    `json:"sort"`
	Status      int8                   `json:"status"`
	Description string                 `json:"description"`
	Materials   []Material             `json:"materials,optional"`
}

// RecommendListRequest 推荐位列表请求
type RecommendListRequest struct {
	ChannelID uint  `form:"channel_id"`
	Status    *int8 `form:"status,optional"`
}

// RecommendListResponse 推荐位列表响应
type RecommendListResponse struct {
	List []Recommend `json:"list"`
}

// RecommendCreateRequest 创建推荐位请求
type RecommendCreateRequest struct {
	ChannelID   uint                   `json:"channel_id"`
	Title       string                 `json:"title"`
	DisplayType string                 `json:"display_type"`
	SourceType  string                 `json:"source_type"`
	ContentIDs  []uint                 `json:"content_ids,optional"`
	FilterRule  map[string]interface{} `json:"filter_rule,optional"`
	Sort        int                    `json:"sort,optional"`
	Description string                 `json:"description,optional"`
}

// RecommendUpdateRequest 更新推荐位请求
type RecommendUpdateRequest struct {
	ID          uint                   `path:"id"`
	Title       string                 `json:"title,optional"`
	DisplayType string                 `json:"display_type,optional"`
	SourceType  string                 `json:"source_type,optional"`
	ContentIDs  []uint                 `json:"content_ids,optional"`
	FilterRule  map[string]interface{} `json:"filter_rule,optional"`
	Sort        *int                   `json:"sort,optional"`
	Status      *int8                  `json:"status,optional"`
	Description string                 `json:"description,optional"`
}

// ==================== Feed流相关 ====================

// FeedConfig Feed流配置
type FeedConfig struct {
	ID              uint                   `json:"id"`
	ChannelID       uint                   `json:"channel_id"`
	Title           string                 `json:"title"`
	LayoutType      string                 `json:"layout_type"`
	ContentStrategy string                 `json:"content_strategy"`
	ContentIDs      []uint                 `json:"content_ids"`
	FilterRule      map[string]interface{} `json:"filter_rule"`
	Sort            int                    `json:"sort"`
	Status          int8                   `json:"status"`
	Description     string                 `json:"description"`
	Materials       []Material             `json:"materials,optional"`
}

// FeedConfigListRequest Feed流配置列表请求
type FeedConfigListRequest struct {
	ChannelID uint  `form:"channel_id"`
	Status    *int8 `form:"status,optional"`
}

// FeedConfigListResponse Feed流配置列表响应
type FeedConfigListResponse struct {
	List []FeedConfig `json:"list"`
}

// FeedConfigCreateRequest 创建Feed流配置请求
type FeedConfigCreateRequest struct {
	ChannelID       uint                   `json:"channel_id"`
	Title           string                 `json:"title"`
	LayoutType      string                 `json:"layout_type"`
	ContentStrategy string                 `json:"content_strategy"`
	ContentIDs      []uint                 `json:"content_ids,optional"`
	FilterRule      map[string]interface{} `json:"filter_rule,optional"`
	Sort            int                    `json:"sort,optional"`
	Description     string                 `json:"description,optional"`
}

// FeedConfigUpdateRequest 更新Feed流配置请求
type FeedConfigUpdateRequest struct {
	ID              uint                   `path:"id"`
	Title           string                 `json:"title,optional"`
	LayoutType      string                 `json:"layout_type,optional"`
	ContentStrategy string                 `json:"content_strategy,optional"`
	ContentIDs      []uint                 `json:"content_ids,optional"`
	FilterRule      map[string]interface{} `json:"filter_rule,optional"`
	Sort            *int                   `json:"sort,optional"`
	Status          *int8                  `json:"status,optional"`
	Description     string                 `json:"description,optional"`
}

// ==================== 广告位相关 ====================

// AdSlot 广告位
type AdSlot struct {
	ID          uint                   `json:"id"`
	ChannelID   uint                   `json:"channel_id"`
	Name        string                 `json:"name"`
	InsertType  string                 `json:"insert_type"`
	InsertRule  map[string]interface{} `json:"insert_rule"`
	AdType      string                 `json:"ad_type"`
	AdContent   map[string]interface{} `json:"ad_content"`
	LinkURL     string                 `json:"link_url"`
	Status      int8                   `json:"status"`
	Sort        int                    `json:"sort"`
	Description string                 `json:"description"`
}

// AdSlotListRequest 广告位列表请求
type AdSlotListRequest struct {
	ChannelID uint  `form:"channel_id"`
	Status    *int8 `form:"status,optional"`
}

// AdSlotListResponse 广告位列表响应
type AdSlotListResponse struct {
	List []AdSlot `json:"list"`
}

// AdSlotCreateRequest 创建广告位请求
type AdSlotCreateRequest struct {
	ChannelID   uint                   `json:"channel_id"`
	Name        string                 `json:"name"`
	InsertType  string                 `json:"insert_type"`
	InsertRule  map[string]interface{} `json:"insert_rule,optional"`
	AdType      string                 `json:"ad_type"`
	AdContent   map[string]interface{} `json:"ad_content,optional"`
	LinkURL     string                 `json:"link_url,optional"`
	Sort        int                    `json:"sort,optional"`
	Description string                 `json:"description,optional"`
}

// AdSlotUpdateRequest 更新广告位请求
type AdSlotUpdateRequest struct {
	ID          uint                   `path:"id"`
	Name        string                 `json:"name,optional"`
	InsertType  string                 `json:"insert_type,optional"`
	InsertRule  map[string]interface{} `json:"insert_rule,optional"`
	AdType      string                 `json:"ad_type,optional"`
	AdContent   map[string]interface{} `json:"ad_content,optional"`
	LinkURL     string                 `json:"link_url,optional"`
	Sort        *int                   `json:"sort,optional"`
	Status      *int8                  `json:"status,optional"`
	Description string                 `json:"description,optional"`
}

// ==================== 物料管理 ====================

// Material 物料
type Material struct {
	ID           uint     `json:"id"`
	Title        string   `json:"title"`
	Subtitle     string   `json:"subtitle"`
	Type         string   `json:"type"`
	CoverURL     string   `json:"cover_url"`
	ContentURL   string   `json:"content_url"`
	Description  string   `json:"description"`
	Author       string   `json:"author"`
	Tags         []string `json:"tags"`
	Category     string   `json:"category"`
	ViewCount    uint     `json:"view_count"`
	LikeCount    uint     `json:"like_count"`
	CommentCount uint     `json:"comment_count"`
	ShareCount   uint     `json:"share_count"`
	CollectCount uint     `json:"collect_count"`
	Duration     uint     `json:"duration"`
	WordCount    uint     `json:"word_count"`
	ChapterCount uint     `json:"chapter_count"`
	Status       int8     `json:"status"`
	Sort         int      `json:"sort"`
}

// MaterialListRequest 物料列表请求
type MaterialListRequest struct {
	Page     int    `form:"page,default=1"`
	PageSize int    `form:"page_size,default=20"`
	Type     string `form:"type,optional"`
	Category string `form:"category,optional"`
	Keyword  string `form:"keyword,optional"`
	Status   *int8  `form:"status,optional"`
}

// MaterialListResponse 物料列表响应
type MaterialListResponse struct {
	Total int64      `json:"total"`
	List  []Material `json:"list"`
}

// MaterialCreateRequest 创建物料请求
type MaterialCreateRequest struct {
	Title        string   `json:"title"`
	Subtitle     string   `json:"subtitle,optional"`
	Type         string   `json:"type"`
	CoverURL     string   `json:"cover_url,optional"`
	ContentURL   string   `json:"content_url,optional"`
	Description  string   `json:"description,optional"`
	Author       string   `json:"author,optional"`
	Tags         []string `json:"tags,optional"`
	Category     string   `json:"category,optional"`
	Duration     uint     `json:"duration,optional"`
	WordCount    uint     `json:"word_count,optional"`
	ChapterCount uint     `json:"chapter_count,optional"`
	Sort         int      `json:"sort,optional"`
}

// MaterialUpdateRequest 更新物料请求
type MaterialUpdateRequest struct {
	ID           uint     `path:"id"`
	Title        string   `json:"title,optional"`
	Subtitle     string   `json:"subtitle,optional"`
	CoverURL     string   `json:"cover_url,optional"`
	ContentURL   string   `json:"content_url,optional"`
	Description  string   `json:"description,optional"`
	Author       string   `json:"author,optional"`
	Tags         []string `json:"tags,optional"`
	Category     string   `json:"category,optional"`
	Duration     *uint    `json:"duration,optional"`
	WordCount    *uint    `json:"word_count,optional"`
	ChapterCount *uint    `json:"chapter_count,optional"`
	Sort         *int     `json:"sort,optional"`
	Status       *int8    `json:"status,optional"`
}

// ==================== 推荐预览相关 ====================

// RecommendPreviewRequest 推荐预览请求
type RecommendPreviewRequest struct {
	Strategy      string                 `json:"strategy"`
	AlgorithmType string                 `json:"algorithm_type,optional"`
	ContentType   []int8                 `json:"content_type,optional"`
	TagIDs        []uint                 `json:"tag_ids,optional"`
	TopicIDs      []uint                 `json:"topic_ids,optional"`
	ContentIDs    []uint                 `json:"content_ids,optional"`
	SortBy        string                 `json:"sort_by,optional"`
	SortOrder     string                 `json:"sort_order,optional"`
	Limit         int                    `json:"limit,default=10"`
	Offset        int                    `json:"offset,default=0"`
}

// RecommendPreviewResponse 推荐预览响应
type RecommendPreviewResponse struct {
	Total   int64                `json:"total"`
	Content []RecommendContentItem `json:"content"`
}

// RecommendContentItem 推荐内容项
type RecommendContentItem struct {
	ID           uint    `json:"id"`
	Title        string  `json:"title"`
	Description  string  `json:"description"`
	Cover        string  `json:"cover"`
	Type         int8    `json:"type"`
	Status       int8    `json:"status"`
	ViewCount    int     `json:"view_count"`
	LikeCount    int     `json:"like_count"`
	CommentCount int     `json:"comment_count"`
	CollectCount int     `json:"collect_count"`
	ShareCount   int     `json:"share_count"`
	HotScore     float64 `json:"hot_score"`
	CreatedAt    string  `json:"created_at"`
}
