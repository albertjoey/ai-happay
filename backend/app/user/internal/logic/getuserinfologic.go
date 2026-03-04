package logic

import (
	"context"

	"happy/app/user/internal/svc"
	"happy/app/user/internal/types"
	"happy/common/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserInfoLogic) GetUserInfo() (*types.UserResponse, error) {
	// 从context获取用户ID（实际应该从JWT中获取）
	userID := l.ctx.Value("user_id")
	if userID == nil {
		userID = uint(1) // 测试用，默认返回第一个用户
	}

	var user model.User
	if err := l.svcCtx.DB.First(&user, userID).Error; err != nil {
		return nil, err
	}

	return &types.UserResponse{
		ID:          user.ID,
		Username:    user.Username,
		Email:       user.Email,
		Phone:       user.Phone,
		Nickname:    user.Nickname,
		Avatar:      user.Avatar,
		Gender:      user.Gender,
		Bio:         user.Bio,
		Role:        user.Role,
		FollowCount: user.FollowCount,
		FansCount:   user.FansCount,
		LikeCount:   user.LikeCount,
	}, nil
}
