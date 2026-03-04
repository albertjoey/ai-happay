package logic

import (
	"context"
	"fmt"
	"happy/app/recommend/internal/svc"
	"happy/app/recommend/internal/types"
	"strings"

	"github.com/zeromicro/go-zero/core/logx"
)

type RecommendLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRecommendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RecommendLogic {
	return &RecommendLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RecommendLogic) Recommend(req *types.RecommendRequest) (*types.RecommendResponse, error) {
	switch req.Strategy {
	case types.StrategyAlgorithm:
		return l.algorithmRecommend(req)
	case types.StrategyManual:
		return l.manualRecommend(req)
	case types.StrategyRandom:
		return l.randomRecommend(req)
	case types.StrategyFilter:
		return l.filterRecommend(req)
	default:
		return l.algorithmRecommend(req)
	}
}

// algorithmRecommend 算法推荐
func (l *RecommendLogic) algorithmRecommend(req *types.RecommendRequest) (*types.RecommendResponse, error) {
	// 个性化推荐
	if req.AlgorithmType == types.AlgorithmPersonal && req.UserID > 0 {
		return l.personalRecommend(req)
	}

	// 构建查询 - 实时计算热度分数
	query := `
		SELECT id, title, description, cover, type, status,
		       view_count, like_count, comment_count, collect_count, share_count,
		       (view_count * 1 + like_count * 5 + comment_count * 10 + collect_count * 8 + share_count * 15) / 
		       POWER(TIMESTAMPDIFF(HOUR, COALESCE(publish_at, created_at), NOW()) + 1, 0.5) as hot_score,
		       created_at
		FROM content
		WHERE deleted_at IS NULL AND status = 1
	`
	args := []interface{}{}

	// 内容类型筛选
	if len(req.ContentType) > 0 {
		placeholders := make([]string, len(req.ContentType))
		for i, t := range req.ContentType {
			placeholders[i] = "?"
			args = append(args, t)
		}
		query += fmt.Sprintf(" AND type IN (%s)", strings.Join(placeholders, ","))
	}

	// 标签筛选
	if len(req.TagIDs) > 0 {
		query += ` AND id IN (
			SELECT content_id FROM content_tag WHERE tag_id IN (
		`
		placeholders := make([]string, len(req.TagIDs))
		for i, tagID := range req.TagIDs {
			placeholders[i] = "?"
			args = append(args, tagID)
		}
		query += strings.Join(placeholders, ",") + `))`
	}

	// 话题筛选
	if len(req.TopicIDs) > 0 {
		query += ` AND id IN (
			SELECT content_id FROM content_topic WHERE topic_id IN (
		`
		placeholders := make([]string, len(req.TopicIDs))
		for i, topicID := range req.TopicIDs {
			placeholders[i] = "?"
			args = append(args, topicID)
		}
		query += strings.Join(placeholders, ",") + `))`
	}

	// 排序
	orderBy := "hot_score"
	if req.SortBy != "" {
		orderBy = req.SortBy
	}
	orderDir := "DESC"
	if req.SortOrder != "" {
		orderDir = req.SortOrder
	}

	// 根据算法类型调整排序
	switch req.AlgorithmType {
	case types.AlgorithmHot:
		orderBy = "hot_score"
	case types.AlgorithmTime:
		orderBy = "created_at"
	}

	query += fmt.Sprintf(" ORDER BY %s %s LIMIT ? OFFSET ?", orderBy, orderDir)
	args = append(args, req.Limit, req.Offset)

	// 执行查询
	var contents []types.ContentResponse
	err := l.svcCtx.DB.Raw(query, args...).Scan(&contents).Error
	if err != nil {
		return nil, err
	}

	// 统计总数
	countQuery := "SELECT COUNT(*) FROM content WHERE deleted_at IS NULL AND status = 1"
	var total int64
	l.svcCtx.DB.Raw(countQuery).Scan(&total)

	return &types.RecommendResponse{
		Total:   total,
		Content: contents,
	}, nil
}

// personalRecommend 个性化推荐（基于用户兴趣标签）
func (l *RecommendLogic) personalRecommend(req *types.RecommendRequest) (*types.RecommendResponse, error) {
	// 1. 获取用户兴趣标签（基于用户行为统计）
	userTags, err := l.getUserInterestTags(req.UserID)
	if err != nil {
		logx.Errorf("获取用户兴趣标签失败: %v", err)
		// 降级为热度推荐
		req.AlgorithmType = types.AlgorithmHot
		return l.algorithmRecommend(req)
	}

	if len(userTags) == 0 {
		// 无兴趣标签，降级为热度推荐
		req.AlgorithmType = types.AlgorithmHot
		return l.algorithmRecommend(req)
	}

	// 2. 基于用户兴趣标签推荐内容
	query := `
		SELECT c.id, c.title, c.description, c.cover, c.type, c.status,
		       c.view_count, c.like_count, c.comment_count, c.collect_count, c.share_count,
		       (c.view_count * 1 + c.like_count * 5 + c.comment_count * 10 + c.collect_count * 8 + c.share_count * 15) / 
		       POWER(TIMESTAMPDIFF(HOUR, COALESCE(c.publish_at, c.created_at), NOW()) + 1, 0.5) as hot_score,
		       c.created_at, COUNT(ct.tag_id) as tag_match_count
		FROM content c
		INNER JOIN content_tag ct ON c.id = ct.content_id
		WHERE c.deleted_at IS NULL AND c.status = 1
		AND ct.tag_id IN (
	`
	placeholders := make([]string, len(userTags))
	args := []interface{}{}
	for i, tagID := range userTags {
		placeholders[i] = "?"
		args = append(args, tagID)
	}
	query += strings.Join(placeholders, ",") + `)`

	// 内容类型筛选
	if len(req.ContentType) > 0 {
		placeholders := make([]string, len(req.ContentType))
		for i, t := range req.ContentType {
			placeholders[i] = "?"
			args = append(args, t)
		}
		query += fmt.Sprintf(" AND c.type IN (%s)", strings.Join(placeholders, ","))
	}

	query += ` GROUP BY c.id ORDER BY tag_match_count DESC, c.hot_score DESC LIMIT ? OFFSET ?`
	args = append(args, req.Limit, req.Offset)

	var contents []types.ContentResponse
	err = l.svcCtx.DB.Raw(query, args...).Scan(&contents).Error
	if err != nil {
		return nil, err
	}

	return &types.RecommendResponse{
		Total:   int64(len(contents)),
		Content: contents,
	}, nil
}

// getUserInterestTags 获取用户兴趣标签
func (l *RecommendLogic) getUserInterestTags(userID uint) ([]uint, error) {
	// 基于用户行为（浏览、点赞、收藏）统计兴趣标签
	query := `
		SELECT ct.tag_id, COUNT(*) as weight
		FROM content_tag ct
		WHERE ct.content_id IN (
			-- 用户浏览过的内容
			SELECT content_id FROM view_history WHERE user_id = ?
			UNION
			-- 用户点赞的内容
			SELECT target_id FROM interaction WHERE user_id = ? AND type = 1 AND target_type = 1
			UNION
			-- 用户收藏的内容
			SELECT target_id FROM interaction WHERE user_id = ? AND type = 2 AND target_type = 1
		)
		GROUP BY ct.tag_id
		ORDER BY weight DESC
		LIMIT 10
	`

	var tagResults []struct {
		TagID  uint
		Weight int
	}
	err := l.svcCtx.DB.Raw(query, userID, userID, userID).Scan(&tagResults).Error
	if err != nil {
		return nil, err
	}

	tagIDs := make([]uint, len(tagResults))
	for i, t := range tagResults {
		tagIDs[i] = t.TagID
	}

	return tagIDs, nil
}

// manualRecommend 人工推荐
func (l *RecommendLogic) manualRecommend(req *types.RecommendRequest) (*types.RecommendResponse, error) {
	if len(req.ContentIDs) == 0 {
		return &types.RecommendResponse{Total: 0, Content: []types.ContentResponse{}}, nil
	}

	// 查询指定ID的内容
	query := `
		SELECT id, title, description, cover, type, status,
		       view_count, like_count, comment_count, collect_count, share_count,
		       (view_count * 1 + like_count * 5 + comment_count * 10 + collect_count * 8 + share_count * 15) / 
		       POWER(TIMESTAMPDIFF(HOUR, COALESCE(publish_at, created_at), NOW()) + 1, 0.5) as hot_score,
		       created_at
		FROM content
		WHERE id IN (?) AND deleted_at IS NULL
	`

	var contents []types.ContentResponse
	err := l.svcCtx.DB.Raw(query, req.ContentIDs).Scan(&contents).Error
	if err != nil {
		return nil, err
	}

	return &types.RecommendResponse{
		Total:   int64(len(contents)),
		Content: contents,
	}, nil
}

// randomRecommend 随机推荐
func (l *RecommendLogic) randomRecommend(req *types.RecommendRequest) (*types.RecommendResponse, error) {
	query := `
		SELECT id, title, description, cover, type, status,
		       view_count, like_count, comment_count, collect_count, share_count,
		       (view_count * 1 + like_count * 5 + comment_count * 10 + collect_count * 8 + share_count * 15) / 
		       POWER(TIMESTAMPDIFF(HOUR, COALESCE(publish_at, created_at), NOW()) + 1, 0.5) as hot_score,
		       created_at
		FROM content
		WHERE deleted_at IS NULL AND status = 1
		ORDER BY RAND()
		LIMIT ?
	`

	var contents []types.ContentResponse
	err := l.svcCtx.DB.Raw(query, req.Limit).Scan(&contents).Error
	if err != nil {
		return nil, err
	}

	return &types.RecommendResponse{
		Total:   int64(len(contents)),
		Content: contents,
	}, nil
}

// filterRecommend 条件筛选
func (l *RecommendLogic) filterRecommend(req *types.RecommendRequest) (*types.RecommendResponse, error) {
	// 类似算法推荐，但支持更多筛选条件
	return l.algorithmRecommend(req)
}
