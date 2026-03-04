package types

// ==================== 角色管理 ====================

type Role struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Code        string `json:"code"`
	Description string `json:"description"`
	Status      int8   `json:"status"`
	CreatedAt   string `json:"created_at"`
}

type RoleListRequest struct {
	Page     int `form:"page,default=1"`
	PageSize int `form:"page_size,default=20"`
}

type RoleListResponse struct {
	Total int64  `json:"total"`
	List  []Role `json:"list"`
}

type RoleCreateRequest struct {
	Name        string `json:"name"`
	Code        string `json:"code"`
	Description string `json:"description,optional"`
}

type RoleUpdateRequest struct {
	ID          uint   `path:"id"`
	Name        string `json:"name,optional"`
	Description string `json:"description,optional"`
	Status      *int8  `json:"status,optional"`
}

// ==================== 权限管理 ====================

type Permission struct {
	ID       uint         `json:"id"`
	Name     string       `json:"name"`
	Code     string       `json:"code"`
	Type     string       `json:"type"`
	ParentID uint         `json:"parent_id"`
	Path     string       `json:"path"`
	Icon     string       `json:"icon"`
	Status   int8         `json:"status"`
	Children []Permission `json:"children,omitempty"`
}

type PermissionCreateRequest struct {
	Name     string `json:"name"`
	Code     string `json:"code"`
	Type     string `json:"type"`
	ParentID uint   `json:"parent_id,optional"`
	Path     string `json:"path,optional"`
	Icon     string `json:"icon,optional"`
}

type PermissionUpdateRequest struct {
	ID       uint   `path:"id"`
	Name     string `json:"name,optional"`
	Type     string `json:"type,optional"`
	ParentID *uint  `json:"parent_id,optional"`
	Path     string `json:"path,optional"`
	Icon     string `json:"icon,optional"`
	Status   *int8  `json:"status,optional"`
}

// ==================== 管理员管理 ====================

type AdminUser struct {
	ID          uint   `json:"id"`
	Username    string `json:"username"`
	Realname    string `json:"realname"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	Avatar      string `json:"avatar"`
	Status      int8   `json:"status"`
	LastLoginAt string `json:"last_login_at"`
	CreatedAt   string `json:"created_at"`
}

type AdminUserListRequest struct {
	Page     int    `form:"page,default=1"`
	PageSize int    `form:"page_size,default=20"`
	Keyword  string `form:"keyword,optional"`
}

type AdminUserListResponse struct {
	Total int64       `json:"total"`
	List  []AdminUser `json:"list"`
}

type AdminUserCreateRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Realname string `json:"realname"`
	Email    string `json:"email,optional"`
	Phone    string `json:"phone,optional"`
}

type AdminUserUpdateRequest struct {
	ID       uint   `path:"id"`
	Realname string `json:"realname,optional"`
	Email    string `json:"email,optional"`
	Phone    string `json:"phone,optional"`
	Status   *int8  `json:"status,optional"`
}

// ==================== 权限分配 ====================

type AssignRolePermissionsRequest struct {
	RoleID        uint   `path:"role_id"`
	PermissionIDs []uint `json:"permission_ids"`
}

type AssignAdminRolesRequest struct {
	AdminID uint   `path:"admin_id"`
	RoleIDs []uint `json:"role_ids"`
}
