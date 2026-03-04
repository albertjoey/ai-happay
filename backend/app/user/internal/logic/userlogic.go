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

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterRequest) (*types.LoginResponse, error) {
	// 验证参数
	if req.Username == "" && req.Email == "" && req.Phone == "" {
		return nil, errors.New("用户名、邮箱或手机号至少填写一项")
	}

	// 检查用户是否已存在
	var count int64
	if req.Username != "" {
		l.svcCtx.DB.Model(&model.User{}).Where("username = ?", req.Username).Count(&count)
		if count > 0 {
			return nil, errors.New("用户名已存在")
		}
	}
	if req.Email != "" {
		l.svcCtx.DB.Model(&model.User{}).Where("email = ?", req.Email).Count(&count)
		if count > 0 {
			return nil, errors.New("邮箱已存在")
		}
	}
	if req.Phone != "" {
		l.svcCtx.DB.Model(&model.User{}).Where("phone = ?", req.Phone).Count(&count)
		if count > 0 {
			return nil, errors.New("手机号已存在")
		}
	}

	// 加密密码
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	// 创建用户
	user := &model.User{
		TenantID: 1, // 默认租户
		Username: req.Username,
		Email:    req.Email,
		Phone:    req.Phone,
		Password: hashedPassword,
		Nickname: req.Nickname,
		Status:   1,
		Role:     0,
	}

	if user.Nickname == "" {
		if req.Username != "" {
			user.Nickname = req.Username
		} else if req.Email != "" {
			user.Nickname = req.Email
		} else {
			user.Nickname = req.Phone
		}
	}

	if err := l.svcCtx.DB.Create(user).Error; err != nil {
		return nil, err
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
