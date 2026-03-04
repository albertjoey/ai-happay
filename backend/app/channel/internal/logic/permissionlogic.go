package logic

import (
	"context"
	"happy/app/channel/internal/svc"
	"happy/app/channel/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PermissionListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPermissionListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PermissionListLogic {
	return &PermissionListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PermissionListLogic) PermissionTree() ([]types.Permission, error) {
	// 查询所有权限
	var permissions []types.Permission
	l.svcCtx.DB.Raw(`
		SELECT id, name, code, type, parent_id, path, icon, status
		FROM permission
		WHERE deleted_at IS NULL
		ORDER BY id ASC
	`).Scan(&permissions)

	// 构建树形结构
	return buildPermissionTree(permissions, 0), nil
}

func buildPermissionTree(permissions []types.Permission, parentID uint) []types.Permission {
	var tree []types.Permission
	for _, p := range permissions {
		if p.ParentID == parentID {
			children := buildPermissionTree(permissions, p.ID)
			if len(children) > 0 {
				p.Children = children
			}
			tree = append(tree, p)
		}
	}
	return tree
}

type PermissionCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPermissionCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PermissionCreateLogic {
	return &PermissionCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PermissionCreateLogic) PermissionCreate(req *types.PermissionCreateRequest) (interface{}, error) {
	result := l.svcCtx.DB.Exec(`
		INSERT INTO permission (name, code, type, parent_id, path, icon, status)
		VALUES (?, ?, ?, ?, ?, ?, 1)
	`, req.Name, req.Code, req.Type, req.ParentID, req.Path, req.Icon)

	if result.Error != nil {
		return nil, result.Error
	}

	return map[string]interface{}{"id": 1, "success": true}, nil
}

type PermissionUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPermissionUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PermissionUpdateLogic {
	return &PermissionUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PermissionUpdateLogic) PermissionUpdate(req *types.PermissionUpdateRequest) (interface{}, error) {
	result := l.svcCtx.DB.Exec(`
		UPDATE permission SET name = ?, type = ?, parent_id = ?, path = ?, icon = ?, status = ?
		WHERE id = ? AND deleted_at IS NULL
	`, req.Name, req.Type, req.ParentID, req.Path, req.Icon, req.Status, req.ID)

	if result.Error != nil {
		return nil, result.Error
	}

	return map[string]interface{}{"success": true}, nil
}

type PermissionDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPermissionDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PermissionDeleteLogic {
	return &PermissionDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PermissionDeleteLogic) PermissionDelete(id uint) (interface{}, error) {
	result := l.svcCtx.DB.Exec("UPDATE permission SET deleted_at = NOW() WHERE id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}

	return map[string]interface{}{"success": true}, nil
}
