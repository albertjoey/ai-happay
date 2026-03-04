package logic

import (
	"context"
	"time"

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
	// 检查是否已点赞
	var count int64
	l.svcCtx.DB.Raw(`SELECT COUNT(*) FROM likes WHERE user_id = ? AND target_type = ? AND target_id = ?`, userID, targetType, targetID).Scan(&count)
	
	if count > 0 {
		// 已点赞，取消点赞
		l.svcCtx.DB.Exec(`DELETE FROM likes WHERE user_id = ? AND target_type = ? AND target_id = ?`, userID, targetType, targetID)
		// 更新物料点赞数-1
		l.svcCtx.DB.Exec(`UPDATE material SET like_count = GREATEST(0, like_count - 1) WHERE id = ?`, targetID)
		return map[string]interface{}{"liked": false, "message": "取消点赞成功"}, nil
	}
	
	// 未点赞，添加点赞
	l.svcCtx.DB.Exec(`INSERT INTO likes (site_id, user_id, target_type, target_id) VALUES (1, ?, ?, ?)`, userID, targetType, targetID)
	// 更新物料点赞数+1
	l.svcCtx.DB.Exec(`UPDATE material SET like_count = like_count + 1 WHERE id = ?`, targetID)
	
	return map[string]interface{}{"liked": true, "message": "点赞成功"}, nil
}

// GetLikeStatus 获取点赞状态
func (l *LikeLogic) GetLikeStatus(userID uint, targetType string, targetID uint) (bool, error) {
	var count int64
	l.svcCtx.DB.Raw(`SELECT COUNT(*) FROM likes WHERE user_id = ? AND target_type = ? AND target_id = ?`, userID, targetType, targetID).Scan(&count)
	return count > 0, nil
}

// GetLikeCount 获取点赞数量
func (l *LikeLogic) GetLikeCount(targetType string, targetID uint) (int64, error) {
	var count int64
	l.svcCtx.DB.Raw(`SELECT COUNT(*) FROM likes WHERE target_type = ? AND target_id = ?`, targetType, targetID).Scan(&count)
	return count, nil
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
	// 检查是否已收藏
	var count int64
	l.svcCtx.DB.Raw(`SELECT COUNT(*) FROM collects WHERE user_id = ? AND target_type = ? AND target_id = ?`, userID, targetType, targetID).Scan(&count)
	
	if count > 0 {
		// 已收藏，取消收藏
		l.svcCtx.DB.Exec(`DELETE FROM collects WHERE user_id = ? AND target_type = ? AND target_id = ?`, userID, targetType, targetID)
		// 更新物料收藏数-1
		l.svcCtx.DB.Exec(`UPDATE material SET collect_count = GREATEST(0, collect_count - 1) WHERE id = ?`, targetID)
		return map[string]interface{}{"collected": false, "message": "取消收藏成功"}, nil
	}
	
	// 未收藏，添加收藏
	l.svcCtx.DB.Exec(`INSERT INTO collects (site_id, user_id, target_type, target_id) VALUES (1, ?, ?, ?)`, userID, targetType, targetID)
	// 更新物料收藏数+1
	l.svcCtx.DB.Exec(`UPDATE material SET collect_count = collect_count + 1 WHERE id = ?`, targetID)
	
	return map[string]interface{}{"collected": true, "message": "收藏成功"}, nil
}

// GetCollectStatus 获取收藏状态
func (l *CollectLogic) GetCollectStatus(userID uint, targetType string, targetID uint) (bool, error) {
	var count int64
	l.svcCtx.DB.Raw(`SELECT COUNT(*) FROM collects WHERE user_id = ? AND target_type = ? AND target_id = ?`, userID, targetType, targetID).Scan(&count)
	return count > 0, nil
}

// GetCollectList 获取用户收藏列表
func (l *CollectLogic) GetCollectList(userID uint, page, pageSize int) (interface{}, error) {
	offset := (page - 1) * pageSize
	
	// 获取总数
	var total int64
	l.svcCtx.DB.Raw(`SELECT COUNT(*) FROM collects WHERE user_id = ?`, userID).Scan(&total)
	
	// 获取收藏列表
	type CollectItem struct {
		ID         uint      `json:"id"`
		TargetType string    `json:"target_type"`
		TargetID   uint      `json:"target_id"`
		CreatedAt  time.Time `json:"created_at"`
	}
	var items []CollectItem
	l.svcCtx.DB.Raw(`
		SELECT id, target_type, target_id, created_at 
		FROM collects 
		WHERE user_id = ? 
		ORDER BY created_at DESC 
		LIMIT ? OFFSET ?
	`, userID, pageSize, offset).Scan(&items)
	
	// 获取物料详情
	type MaterialInfo struct {
		ID       uint   `json:"id"`
		Title    string `json:"title"`
		Type     string `json:"type"`
		CoverURL string `json:"cover_url"`
		Author   string `json:"author"`
	}
	var materials []MaterialInfo
	for _, item := range items {
		var m MaterialInfo
		l.svcCtx.DB.Raw(`SELECT id, title, type, cover_url, author FROM material WHERE id = ?`, item.TargetID).Scan(&m)
		materials = append(materials, m)
	}
	
	return map[string]interface{}{
		"total":     total,
		"list":      items,
		"materials": materials,
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
	result := l.svcCtx.DB.Exec(`
		INSERT INTO comments (video_id, user_id, target_type, target_id, content, status)
		VALUES (?, ?, ?, ?, ?, 1)
	`, targetID, userID, targetType, targetID, content)
	
	if result.Error != nil {
		return nil, result.Error
	}
	
	// 更新物料评论数+1
	l.svcCtx.DB.Exec(`UPDATE material SET comment_count = comment_count + 1 WHERE id = ?`, targetID)
	
	return map[string]interface{}{"id": result.RowsAffected, "success": true}, nil
}

// GetCommentList 获取评论列表
func (l *CommentLogic) GetCommentList(targetType string, targetID uint, page, pageSize int) (interface{}, error) {
	offset := (page - 1) * pageSize
	
	// 获取总数
	var total int64
	l.svcCtx.DB.Raw(`SELECT COUNT(*) FROM comments WHERE target_type = ? AND target_id = ? AND status = 1`, targetType, targetID).Scan(&total)
	
	// 获取评论列表
	type Comment struct {
		ID        uint   `json:"id"`
		UserID    uint   `json:"user_id"`
		Content   string `json:"content"`
		LikeCount int    `json:"like_count"`
		CreatedAt string `json:"created_at"`
	}
	var comments []Comment
	l.svcCtx.DB.Raw(`
		SELECT id, user_id, content, like_count, DATE_FORMAT(created_at, '%Y-%m-%d %H:%i:%s') as created_at
		FROM comments 
		WHERE target_type = ? AND target_id = ? AND status = 1 
		ORDER BY created_at DESC 
		LIMIT ? OFFSET ?
	`, targetType, targetID, pageSize, offset).Scan(&comments)
	
	return map[string]interface{}{
		"total": total,
		"list":  comments,
	}, nil
}

// DeleteComment 删除评论
func (l *CommentLogic) DeleteComment(userID, commentID uint) (interface{}, error) {
	// 获取评论信息
	type CommentInfo struct {
		TargetID   uint
		TargetType string
	}
	var info CommentInfo
	l.svcCtx.DB.Raw(`SELECT target_id, target_type FROM comments WHERE id = ? AND user_id = ?`, commentID, userID).Scan(&info)
	
	if info.TargetID == 0 {
		return map[string]interface{}{"success": false, "message": "评论不存在或无权删除"}, nil
	}
	
	// 删除评论
	l.svcCtx.DB.Exec(`UPDATE comments SET status = 0 WHERE id = ?`, commentID)
	
	// 更新物料评论数-1
	l.svcCtx.DB.Exec(`UPDATE material SET comment_count = GREATEST(0, comment_count - 1) WHERE id = ?`, info.TargetID)
	
	return map[string]interface{}{"success": true}, nil
}

// LikeComment 点赞评论
func (l *CommentLogic) LikeComment(commentID uint) (interface{}, error) {
	l.svcCtx.DB.Exec(`UPDATE comments SET like_count = like_count + 1 WHERE id = ?`, commentID)
	return map[string]interface{}{"success": true}, nil
}
