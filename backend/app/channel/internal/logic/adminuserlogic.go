package logic

import (
	"context"
	"happy/app/channel/internal/svc"
	"happy/app/channel/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AdminUserListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAdminUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdminUserListLogic {
	return &AdminUserListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AdminUserListLogic) AdminUserList(req *types.AdminUserListRequest) (*types.AdminUserListResponse, error) {
	// 使用Repository接口
	list, total, err := l.svcCtx.AdminUserRepo.List(l.ctx, req)
	if err != nil {
		return nil, err
	}

	return &types.AdminUserListResponse{
		Total: total,
		List:  list,
	}, nil
}

type AdminUserCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAdminUserCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdminUserCreateLogic {
	return &AdminUserCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AdminUserCreateLogic) AdminUserCreate(req *types.AdminUserCreateRequest) (interface{}, error) {
	admin := &types.AdminUser{
		Username: req.Username,
		Realname: req.Realname,
		Email:    req.Email,
		Phone:    req.Phone,
		Status:   1,
	}

	err := l.svcCtx.AdminUserRepo.Create(l.ctx, admin)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{"id": admin.ID, "success": true}, nil
}

type AdminUserUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAdminUserUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdminUserUpdateLogic {
	return &AdminUserUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AdminUserUpdateLogic) AdminUserUpdate(req *types.AdminUserUpdateRequest) (interface{}, error) {
	admin := &types.AdminUser{
		ID:       req.ID,
		Realname: req.Realname,
		Email:    req.Email,
		Phone:    req.Phone,
	}
	if req.Status != nil {
		admin.Status = *req.Status
	}

	err := l.svcCtx.AdminUserRepo.Update(l.ctx, admin)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{"success": true}, nil
}

type AdminUserDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAdminUserDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdminUserDeleteLogic {
	return &AdminUserDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AdminUserDeleteLogic) AdminUserDelete(id uint) (interface{}, error) {
	err := l.svcCtx.AdminUserRepo.Delete(l.ctx, id)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{"success": true}, nil
}
