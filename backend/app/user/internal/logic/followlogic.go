package logic

import (
	"context"
	"errors"

	"happy/app/user/internal/svc"
	"happy/app/user/internal/types"
	"happy/common/model"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type FollowLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFollowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FollowLogic {
	return &FollowLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FollowLogic) Follow(req *types.FollowRequest) error {
	// 从context获取用户ID
	userID := l.ctx.Value("user_id")
	if userID == nil {
		userID = uint(1) // 测试用
	}

	followerID := userID.(uint)
	followingID := req.UserID

	// 不能关注自己
	if followerID == followingID {
		return errors.New("不能关注自己")
	}

	// 检查是否已经关注
	var count int64
	l.svcCtx.DB.Model(&model.Follow{}).Where("follower_id = ? AND following_id = ?", followerID, followingID).Count(&count)
	if count > 0 {
		return errors.New("已经关注过了")
	}

	// 创建关注关系
	follow := &model.Follow{
		FollowerID:  followerID,
		FollowingID: followingID,
	}

	return l.svcCtx.DB.Transaction(func(tx *gorm.DB) error {
		// 创建关注记录
		if err := tx.Create(follow).Error; err != nil {
			return err
		}

		// 更新关注数
		if err := tx.Model(&model.User{}).Where("id = ?", followerID).UpdateColumn("follow_count", gorm.Expr("follow_count + ?", 1)).Error; err != nil {
			return err
		}

		// 更新粉丝数
		if err := tx.Model(&model.User{}).Where("id = ?", followingID).UpdateColumn("fans_count", gorm.Expr("fans_count + ?", 1)).Error; err != nil {
			return err
		}

		return nil
	})
}
