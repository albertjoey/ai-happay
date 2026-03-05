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
	// 使用Repository接口
	permissions, err := l.svcCtx.PermissionRepo.List(l.ctx)
	if err != nil {
		return nil, err
	}

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
	permission := &types.Permission{
		Name:     req.Name,
		Code:     req.Code,
		Type:     req.Type,
		ParentID: req.ParentID,
		Path:     req.Path,
		Icon:     req.Icon,
	}

	err := l.svcCtx.PermissionRepo.Create(l.ctx, permission)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{"id": permission.ID, "success": true}, nil
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
	permission := &types.Permission{
		ID:       req.ID,
		Name:     req.Name,
		Type:     req.Type,
		Path:     req.Path,
		Icon:     req.Icon,
	}
	if req.ParentID != nil {
		permission.ParentID = *req.ParentID
	}
	if req.Status != nil {
		permission.Status = *req.Status
	}

	err := l.svcCtx.PermissionRepo.Update(l.ctx, permission)
	if err != nil {
		return nil, err
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
	err := l.svcCtx.PermissionRepo.Delete(l.ctx, id)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{"success": true}, nil
}
