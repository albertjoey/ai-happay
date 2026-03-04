package types

// 内容类型
const (
	ContentTypeLongVideo  int8 = 1 // 长视频
	ContentTypeShortVideo int8 = 2 // 短视频
	ContentTypeDrama      int8 = 3 // 短剧
	ContentTypeComic      int8 = 4 // 漫剧
	ContentTypeNovel      int8 = 5 // 小说
	ContentTypeImageText  int8 = 6 // 图文
)

// 创建内容请求
type CreateContentRequest struct {
	Title       string   `json:"title"`
	Description string   `json:"description,optional"`
	Cover       string   `json:"cover,optional"`
	Type        int8     `json:"type"`
	Media       []Media  `json:"media,optional"`
	Topics      []uint   `json:"topics,optional"`
	Tags        []uint   `json:"tags,optional"`
}

// 媒体资源
type Media struct {
	Type      int8   `json:"type"`                // 1-图片 2-视频 3-音频
	URL       string `json:"url"`
	Thumbnail string `json:"thumbnail,optional"`
	Duration  int    `json:"duration,optional"`
	Width     int    `json:"width,optional"`
	Height    int    `json:"height,optional"`
	Size      int64  `json:"size,optional"`
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
	Media        []Media   `json:"media"`
	Topics       []Topic   `json:"topics"`
	Tags         []Tag     `json:"tags"`
	Author       Author    `json:"author"`
	CreatedAt    string    `json:"created_at"`
}

// 作者信息
type Author struct {
	ID       uint   `json:"id"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
}

// 话题
type Topic struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

// 标签
type Tag struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

// 内容列表请求
type ContentListRequest struct {
	Page     int   `form:"page,default=1"`
	PageSize int   `form:"page_size,default=20"`
	Type     int8  `form:"type,optional"`
	Channel  uint  `form:"channel,optional"`
	UserID   uint  `form:"user_id,optional"`
	Status   int8  `form:"status,optional"`
}

// 内容列表响应
type ContentListResponse struct {
	Total int64             `json:"total"`
	List  []ContentResponse `json:"list"`
}

// 更新内容请求
type UpdateContentRequest struct {
	ID          uint     `json:"id"`
	Title       string   `json:"title,optional"`
	Description string   `json:"description,optional"`
	Cover       string   `json:"cover,optional"`
	Media       []Media  `json:"media,optional"`
	Topics      []uint   `json:"topics,optional"`
	Tags        []uint   `json:"tags,optional"`
}

// 内容详情请求
type ContentDetailRequest struct {
	ID uint `path:"id"`
}

// 频道响应
type ChannelResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Code        string `json:"code"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

// 频道列表响应
type ChannelListResponse struct {
	List []ChannelResponse `json:"list"`
}

// Banner响应
type BannerResponse struct {
	ID     uint   `json:"id"`
	Title  string `json:"title"`
	Image  string `json:"image"`
	Link   string `json:"link"`
}

// Banner列表响应
type BannerListResponse struct {
	List []BannerResponse `json:"list"`
}
