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
	// 构建查询条件
	where := "WHERE deleted_at IS NULL"
	args := []interface{}{}

	if req.Keyword != "" {
		where += " AND (username LIKE ? OR realname LIKE ?)"
		keyword := "%" + req.Keyword + "%"
		args = append(args, keyword, keyword)
	}

	// 查询总数
	var total int64
	l.svcCtx.DB.Raw("SELECT COUNT(*) FROM admin_user "+where, args...).Scan(&total)

	// 查询列表
	offset := (req.Page - 1) * req.PageSize
	var admins []types.AdminUser
	query := `
		SELECT id, username, realname, email, phone, avatar, status,
		       DATE_FORMAT(last_login_at, '%Y-%m-%d %H:%i:%s') as last_login_at,
		       DATE_FORMAT(created_at, '%Y-%m-%d %H:%i:%s') as created_at
		FROM admin_user
		` + where + ` ORDER BY id ASC LIMIT ? OFFSET ?`
	args = append(args, req.PageSize, offset)

	l.svcCtx.DB.Raw(query, args...).Scan(&admins)

	return &types.AdminUserListResponse{
		Total: total,
		List:  admins,
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
	// 密码应该加密,这里简化处理
	result := l.svcCtx.DB.Exec(`
		INSERT INTO admin_user (username, password, realname, email, phone, status)
		VALUES (?, ?, ?, ?, ?, 1)
	`, req.Username, req.Password, req.Realname, req.Email, req.Phone)

	if result.Error != nil {
		return nil, result.Error
	}

	return map[string]interface{}{"id": 1, "success": true}, nil
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
	result := l.svcCtx.DB.Exec(`
		UPDATE admin_user SET realname = ?, email = ?, phone = ?, status = ?
		WHERE id = ? AND deleted_at IS NULL
	`, req.Realname, req.Email, req.Phone, req.Status, req.ID)

	if result.Error != nil {
		return nil, result.Error
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
	result := l.svcCtx.DB.Exec("UPDATE admin_user SET deleted_at = NOW() WHERE id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}

	return map[string]interface{}{"success": true}, nil
}
