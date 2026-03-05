package logic

import (
	"context"
	"happy/app/channel/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// ==================== 点赞系统 ====================

type LikeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLikeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LikeLogic {
	return &LikeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// Like 点赞
func (l *LikeLogic) Like(userID uint, targetType string, targetID uint) (interface{}, error) {
	// 使用Repository接口
	err := l.svcCtx.InteractionRepo.Like(l.ctx, userID, targetID)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{"liked": true, "message": "点赞成功"}, nil
}

// GetLikeStatus 获取点赞状态
func (l *LikeLogic) GetLikeStatus(userID uint, targetType string, targetID uint) (bool, error) {
	return l.svcCtx.InteractionRepo.IsLiked(l.ctx, userID, targetID)
}

// GetLikeCount 获取点赞数量
func (l *LikeLogic) GetLikeCount(targetType string, targetID uint) (int64, error) {
	// 暂时返回0，后续可以扩展Repository接口
	return 0, nil
}

// ==================== 收藏系统 ====================

type CollectLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCollectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CollectLogic {
	return &CollectLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// Collect 收藏
func (l *CollectLogic) Collect(userID uint, targetType string, targetID uint) (interface{}, error) {
	// 使用Repository接口
	err := l.svcCtx.InteractionRepo.Collect(l.ctx, userID, targetID)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{"collected": true, "message": "收藏成功"}, nil
}

// GetCollectStatus 获取收藏状态
func (l *CollectLogic) GetCollectStatus(userID uint, targetType string, targetID uint) (bool, error) {
	return l.svcCtx.InteractionRepo.IsCollected(l.ctx, userID, targetID)
}

// GetCollectList 获取用户收藏列表
func (l *CollectLogic) GetCollectList(userID uint, page, pageSize int) (interface{}, error) {
	// 暂时返回空列表，后续可以扩展Repository接口
	return map[string]interface{}{
		"total":     0,
		"list":      []interface{}{},
		"materials": []interface{}{},
	}, nil
}

// ==================== 评论系统 ====================

type CommentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommentLogic {
	return &CommentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// CreateComment 创建评论
func (l *CommentLogic) CreateComment(userID uint, targetType string, targetID uint, content string) (interface{}, error) {
	// 使用Repository接口
	err := l.svcCtx.InteractionRepo.Comment(l.ctx, userID, targetID, content, 0)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{"success": true}, nil
}

// GetCommentList 获取评论列表
func (l *CommentLogic) GetCommentList(targetType string, targetID uint, page, pageSize int) (interface{}, error) {
	// 使用Repository接口
	comments, total, err := l.svcCtx.InteractionRepo.GetComments(l.ctx, targetID, page, pageSize)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"total": total,
		"list":  comments,
	}, nil
}

// DeleteComment 删除评论
func (l *CommentLogic) DeleteComment(userID, commentID uint) (interface{}, error) {
	// 暂时返回成功，后续可以扩展Repository接口
	return map[string]interface{}{"success": true}, nil
}

// LikeComment 点赞评论
func (l *CommentLogic) LikeComment(commentID uint) (interface{}, error) {
	// 暂时返回成功，后续可以扩展Repository接口
	return map[string]interface{}{"success": true}, nil
}
