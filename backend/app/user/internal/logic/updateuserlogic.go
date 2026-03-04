package logic

import (
	"context"

	"happy/app/user/internal/svc"
	"happy/app/user/internal/types"
	"happy/common/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserLogic {
	return &UpdateUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateUserLogic) UpdateUser(req *types.UpdateUserRequest) error {
	// 从context获取用户ID
	userID := l.ctx.Value("user_id")
	if userID == nil {
		userID = uint(1) // 测试用
	}

	updates := make(map[string]interface{})
	if req.Nickname != "" {
		updates["nickname"] = req.Nickname
	}
	if req.Avatar != "" {
		updates["avatar"] = req.Avatar
	}
	if req.Gender != 0 {
		updates["gender"] = req.Gender
	}
	if req.Bio != "" {
		updates["bio"] = req.Bio
	}

	return l.svcCtx.DB.Model(&model.User{}).Where("id = ?", userID).Updates(updates).Error
}
