package logic

import (
	"context"
	"errors"

	"happy/app/user/internal/svc"
	"happy/app/user/internal/types"
	"happy/common/model"
	"happy/common/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginRequest) (*types.LoginResponse, error) {
	// 查找用户
	var user model.User
	if req.Username != "" {
		l.svcCtx.DB.Where("username = ?", req.Username).First(&user)
	} else if req.Email != "" {
		l.svcCtx.DB.Where("email = ?", req.Email).First(&user)
	} else if req.Phone != "" {
		l.svcCtx.DB.Where("phone = ?", req.Phone).First(&user)
	} else {
		return nil, errors.New("请输入用户名、邮箱或手机号")
	}

	if user.ID == 0 {
		return nil, errors.New("用户不存在")
	}

	// 验证密码
	if !utils.CheckPassword(req.Password, user.Password) {
		return nil, errors.New("密码错误")
	}

	// 检查用户状态
	if user.Status == 0 {
		return nil, errors.New("用户已被禁用")
	}

	// 生成token
	token, err := utils.GenerateToken(user.ID, user.Username, user.Role, user.TenantID, l.svcCtx.Config.Auth.AccessSecret, l.svcCtx.Config.Auth.AccessExpire)
	if err != nil {
		return nil, err
	}

	return &types.LoginResponse{
		Token: token,
		User: types.UserResponse{
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
		},
	}, nil
}
