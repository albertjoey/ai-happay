package types

// 推荐策略
const (
	StrategyAlgorithm = "algorithm" // 算法推荐
	StrategyManual    = "manual"    // 人工推荐
	StrategyRandom    = "random"    // 随机推荐
	StrategyFilter    = "filter"    // 条件筛选
)

// 算法类型
const (
	AlgorithmHot      = "hot"      // 热度排序
	AlgorithmTime     = "time"     // 时间排序
	AlgorithmPersonal = "personal" // 个性化推荐
)

// 推荐请求
type RecommendRequest struct {
	Strategy      string `json:"strategy"`                // 推荐策略
	AlgorithmType string `json:"algorithm_type,optional"` // 算法类型
	ContentType   []int8 `json:"content_type,optional"`   // 内容类型筛选
	TagIDs        []uint `json:"tag_ids,optional"`        // 标签筛选
	TopicIDs      []uint `json:"topic_ids,optional"`      // 话题筛选
	ChannelID     uint   `json:"channel_id,optional"`     // 频道筛选
	Status        int8   `json:"status,optional"`         // 状态筛选
	SortBy        string `json:"sort_by,optional"`        // 排序字段
	SortOrder     string `json:"sort_order,optional"`     // 排序方向
	ContentIDs    []uint `json:"content_ids,optional"`    // 内容ID列表（人工推荐）
	UserID        uint   `json:"user_id,optional"`        // 用户ID（个性化推荐）
	Limit         int    `json:"limit,default=20"`        // 数量限制
	Offset        int    `json:"offset,default=0"`        // 偏移量
}

// 内容响应
type ContentResponse struct {
	ID           uint      `json:"id"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	Cover        string    `json:"cover"`
	Type         int8      `json:"type"`
	Status       int8      `json:"status"`
	ViewCount    int       `json:"view_count"`
	LikeCount    int       `json:"like_count"`
	CommentCount int       `json:"comment_count"`
	CollectCount int       `json:"collect_count"`
	ShareCount   int       `json:"share_count"`
	HotScore     float64   `json:"hot_score"`
	CreatedAt    string    `json:"created_at"`
}

// 推荐响应
type RecommendResponse struct {
	Total   int64              `json:"total"`
	Content []ContentResponse `json:"content"`
}

// 排行榜响应
type RankingResponse struct {
	List []ContentResponse `json:"list"`
}

// 热度更新响应
type HotScoreUpdateResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
