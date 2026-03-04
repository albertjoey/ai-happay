package logic

import (
	"context"

	"happy/app/user/internal/svc"
	"happy/app/user/internal/types"
	"happy/common/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserListLogic {
	return &GetUserListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserListLogic) GetUserList(req *types.UserListRequest) (*types.UserListResponse, error) {
	var users []model.User
	var total int64

	db := l.svcCtx.DB.Model(&model.User{})

	// 搜索条件
	if req.Keyword != "" {
		db = db.Where("username LIKE ? OR nickname LIKE ? OR email LIKE ?", "%"+req.Keyword+"%", "%"+req.Keyword+"%", "%"+req.Keyword+"%")
	}
	if req.Role != 0 {
		db = db.Where("role = ?", req.Role)
	}

	// 统计总数
	db.Count(&total)

	// 分页查询
	offset := (req.Page - 1) * req.PageSize
	if err := db.Offset(offset).Limit(req.PageSize).Find(&users).Error; err != nil {
		return nil, err
	}

	// 转换为响应格式
	list := make([]types.UserResponse, len(users))
	for i, user := range users {
		list[i] = types.UserResponse{
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
		}
	}

	return &types.UserListResponse{
		Total: total,
		List:  list,
	}, nil
}
