package types

// 搜索请求
type SearchRequest struct {
	Keyword string `form:"keyword"`              // 搜索关键词
	Page    int    `form:"page,default=1"`       // 页码
	PageSize int   `form:"page_size,default=20"` // 每页数量
}

// 搜索响应
type SearchResponse struct {
	Total int64            `json:"total"` // 总数
	List  []SearchResult   `json:"list"`  // 结果列表
}

// 搜索结果
type SearchResult struct {
	ID           uint   `json:"id"`            // 内容ID
	Title        string `json:"title"`         // 标题
	Description  string `json:"description"`   // 描述
	Cover        string `json:"cover"`         // 封面
	Type         int8   `json:"type"`          // 类型
	ViewCount    int    `json:"view_count"`    // 浏览量
	LikeCount    int    `json:"like_count"`    // 点赞数
	AuthorName   string `json:"author_name"`   // 作者名称
	AuthorAvatar string `json:"author_avatar"` // 作者头像
	Highlight    string `json:"highlight"`     // 高亮内容
	Score        float64 `json:"score"`        // 相关度分数
}

// 搜索建议请求
type SearchSuggestRequest struct {
	Keyword string `form:"keyword"` // 搜索关键词
	Limit   int    `form:"limit,default=10"` // 返回数量
}

// 搜索建议响应
type SearchSuggestResponse struct {
	List []string `json:"list"` // 建议列表
}

// 索引内容
type IndexContent struct {
	ID          uint     `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Type        int8     `json:"type"`
	Tags        []string `json:"tags"`
	Topics      []string `json:"topics"`
	AuthorName  string   `json:"author_name"`
	ViewCount   int      `json:"view_count"`
	LikeCount   int      `json:"like_count"`
	Status      int8     `json:"status"`
	TenantID    uint     `json:"tenant_id"`
	CreatedAt   string   `json:"created_at"`
}
