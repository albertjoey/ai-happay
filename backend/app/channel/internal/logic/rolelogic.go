package logic

import (
	"context"
	"happy/app/channel/internal/svc"
	"happy/app/channel/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoleListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRoleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleListLogic {
	return &RoleListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RoleListLogic) RoleList(req *types.RoleListRequest) (*types.RoleListResponse, error) {
	// 查询总数
	var total int64
	l.svcCtx.DB.Raw("SELECT COUNT(*) FROM role WHERE deleted_at IS NULL").Scan(&total)

	// 查询列表
	offset := (req.Page - 1) * req.PageSize
	var roles []types.Role
	l.svcCtx.DB.Raw(`
		SELECT id, name, code, description, status, DATE_FORMAT(created_at, '%Y-%m-%d %H:%i:%s') as created_at
		FROM role
		WHERE deleted_at IS NULL
		ORDER BY id ASC
		LIMIT ? OFFSET ?
	`, req.PageSize, offset).Scan(&roles)

	return &types.RoleListResponse{
		Total: total,
		List:  roles,
	}, nil
}

type RoleCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRoleCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleCreateLogic {
	return &RoleCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RoleCreateLogic) RoleCreate(req *types.RoleCreateRequest) (interface{}, error) {
	result := l.svcCtx.DB.Exec(`
		INSERT INTO role (name, code, description, status)
		VALUES (?, ?, ?, 1)
	`, req.Name, req.Code, req.Description)

	if result.Error != nil {
		return nil, result.Error
	}

	return map[string]interface{}{"id": 1, "success": true}, nil
}

type RoleUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRoleUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleUpdateLogic {
	return &RoleUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RoleUpdateLogic) RoleUpdate(req *types.RoleUpdateRequest) (interface{}, error) {
	result := l.svcCtx.DB.Exec(`
		UPDATE role SET name = ?, description = ?, status = ?
		WHERE id = ? AND deleted_at IS NULL
	`, req.Name, req.Description, req.Status, req.ID)

	if result.Error != nil {
		return nil, result.Error
	}

	return map[string]interface{}{"success": true}, nil
}

type RoleDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRoleDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleDeleteLogic {
	return &RoleDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RoleDeleteLogic) RoleDelete(id uint) (interface{}, error) {
	result := l.svcCtx.DB.Exec("UPDATE role SET deleted_at = NOW() WHERE id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}

	return map[string]interface{}{"success": true}, nil
}

// ==================== 角色权限分配 ====================

type RolePermissionsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRolePermissionsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RolePermissionsLogic {
	return &RolePermissionsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// GetRolePermissions 获取角色的权限ID列表
func (l *RolePermissionsLogic) GetRolePermissions(roleID uint) ([]uint, error) {
	var permissionIDs []uint
	l.svcCtx.DB.Raw(`
		SELECT permission_id FROM role_permission 
		WHERE role_id = ?
	`, roleID).Scan(&permissionIDs)
	
	return permissionIDs, nil
}

type AssignRolePermissionsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAssignRolePermissionsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AssignRolePermissionsLogic {
	return &AssignRolePermissionsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// AssignRolePermissions 分配角色权限
func (l *AssignRolePermissionsLogic) AssignRolePermissions(roleID uint, permissionIDs []uint) (interface{}, error) {
	// 先删除旧的关联
	l.svcCtx.DB.Exec("DELETE FROM role_permission WHERE role_id = ?", roleID)
	
	// 插入新的关联
	for _, permID := range permissionIDs {
		l.svcCtx.DB.Exec(`
			INSERT INTO role_permission (role_id, permission_id)
			VALUES (?, ?)
		`, roleID, permID)
	}
	
	return map[string]interface{}{"success": true}, nil
}

// ==================== 管理员角色分配 ====================

type AdminRolesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAdminRolesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdminRolesLogic {
	return &AdminRolesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// GetAdminRoles 获取管理员的角色ID列表
func (l *AdminRolesLogic) GetAdminRoles(adminUserID uint) ([]uint, error) {
	var roleIDs []uint
	l.svcCtx.DB.Raw(`
		SELECT role_id FROM admin_user_roles 
		WHERE admin_user_id = ?
	`, adminUserID).Scan(&roleIDs)
	
	return roleIDs, nil
}

type AssignAdminRolesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAssignAdminRolesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AssignAdminRolesLogic {
	return &AssignAdminRolesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// AssignAdminRoles 分配管理员角色
func (l *AssignAdminRolesLogic) AssignAdminRoles(adminUserID uint, roleIDs []uint) (interface{}, error) {
	// 先删除旧的关联
	l.svcCtx.DB.Exec("DELETE FROM admin_user_roles WHERE admin_user_id = ?", adminUserID)
	
	// 插入新的关联
	for _, roleID := range roleIDs {
		l.svcCtx.DB.Exec(`
			INSERT INTO admin_user_roles (admin_user_id, role_id)
			VALUES (?, ?)
		`, adminUserID, roleID)
	}
	
	return map[string]interface{}{"success": true}, nil
}
