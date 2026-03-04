package types

// 用户注册请求
type RegisterRequest struct {
	Username string `json:"username,optional"`
	Email    string `json:"email,optional"`
	Phone    string `json:"phone,optional"`
	Password string `json:"password"`
	Nickname string `json:"nickname,optional"`
}

// 用户登录请求
type LoginRequest struct {
	Username string `json:"username,optional"`
	Email    string `json:"email,optional"`
	Phone    string `json:"phone,optional"`
	Password string `json:"password"`
}

// 第三方登录请求
type ThirdPartyLoginRequest struct {
	Type        string `json:"type"` // google/wechat/apple/x
	ThirdPartyID string `json:"third_party_id"`
	Nickname    string `json:"nickname,optional"`
	Avatar      string `json:"avatar,optional"`
}

// 用户信息响应
type UserResponse struct {
	ID          uint   `json:"id"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	Nickname    string `json:"nickname"`
	Avatar      string `json:"avatar"`
	Gender      int8   `json:"gender"`
	Bio         string `json:"bio"`
	Role        int8   `json:"role"`
	FollowCount int    `json:"follow_count"`
	FansCount   int    `json:"fans_count"`
	LikeCount   int    `json:"like_count"`
}

// 登录响应
type LoginResponse struct {
	Token string       `json:"token"`
	User  UserResponse `json:"user"`
}

// 更新用户信息请求
type UpdateUserRequest struct {
	Nickname string `json:"nickname,optional"`
	Avatar   string `json:"avatar,optional"`
	Gender   int8   `json:"gender,optional"`
	Bio      string `json:"bio,optional"`
}

// 关注请求
type FollowRequest struct {
	UserID uint `json:"user_id"`
}

// 用户列表请求
type UserListRequest struct {
	Page     int    `form:"page,default=1"`
	PageSize int    `form:"page_size,default=20"`
	Keyword  string `form:"keyword,optional"`
	Role     int8   `form:"role,optional"`
}

// 用户列表响应
type UserListResponse struct {
	Total int64          `json:"total"`
	List  []UserResponse `json:"list"`
}
