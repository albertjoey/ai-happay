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
	// 使用Repository接口
	list, total, err := l.svcCtx.RoleRepo.List(l.ctx, req)
	if err != nil {
		return nil, err
	}

	return &types.RoleListResponse{
		Total: total,
		List:  list,
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
	role := &types.Role{
		Name:        req.Name,
		Code:        req.Code,
		Description: req.Description,
	}

	err := l.svcCtx.RoleRepo.Create(l.ctx, role)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{"id": role.ID, "success": true}, nil
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
	role := &types.Role{
		ID:          req.ID,
		Name:        req.Name,
		Description: req.Description,
	}
	if req.Status != nil {
		role.Status = *req.Status
	}

	err := l.svcCtx.RoleRepo.Update(l.ctx, role)
	if err != nil {
		return nil, err
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
	err := l.svcCtx.RoleRepo.Delete(l.ctx, id)
	if err != nil {
		return nil, err
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

func (l *RolePermissionsLogic) GetRolePermissions(roleID uint) ([]uint, error) {
	return l.svcCtx.RoleRepo.GetPermissions(l.ctx, roleID)
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

func (l *AssignRolePermissionsLogic) AssignRolePermissions(roleID uint, permissionIDs []uint) (interface{}, error) {
	err := l.svcCtx.RoleRepo.AssignPermissions(l.ctx, roleID, permissionIDs)
	if err != nil {
		return nil, err
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

func (l *AdminRolesLogic) GetAdminRoles(adminUserID uint) ([]uint, error) {
	return l.svcCtx.AdminUserRepo.GetRoles(l.ctx, adminUserID)
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

func (l *AssignAdminRolesLogic) AssignAdminRoles(adminUserID uint, roleIDs []uint) (interface{}, error) {
	err := l.svcCtx.AdminUserRepo.AssignRoles(l.ctx, adminUserID, roleIDs)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{"success": true}, nil
}
