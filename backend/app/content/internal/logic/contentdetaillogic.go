package logic

import (
	"context"
	"encoding/json"
	"happy/app/content/internal/svc"
	"happy/app/content/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ContentDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewContentDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ContentDetailLogic {
	return &ContentDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ContentDetailLogic) ContentDetail(id uint) (*types.ContentResponse, error) {
	querySQL := `
		SELECT c.id, c.title, c.description, c.cover, c.type, c.status, 
		       c.view_count, c.like_count, c.comment_count, c.collect_count, c.share_count,
		       c.media, c.created_at, u.id as author_id, u.nickname as author_nickname, u.avatar as author_avatar
		FROM content c
		LEFT JOIN user u ON c.user_id = u.id
		WHERE c.id = ? AND c.deleted_at IS NULL
	`

	type ContentResult struct {
		ID            uint
		Title         string
		Description   string
		Cover         string
		Type          int8
		Status        int8
		ViewCount     int
		LikeCount     int
		CommentCount  int
		CollectCount  int
		ShareCount    int
		Media         string
		CreatedAt     string
		AuthorID      uint
		AuthorNickname string
		AuthorAvatar  string
	}

	var result ContentResult
	err := l.svcCtx.DB.Raw(querySQL, id).Scan(&result).Error
	if err != nil {
		return nil, err
	}

	// 解析媒体资源
	var mediaList []types.Media
	if result.Media != "" {
		json.Unmarshal([]byte(result.Media), &mediaList)
	}

	// 查询话题
	topicSQL := `
		SELECT t.id, t.name 
		FROM topic t
		INNER JOIN content_topic ct ON t.id = ct.topic_id
		WHERE ct.content_id = ? AND t.deleted_at IS NULL
	`
	var topics []types.Topic
	l.svcCtx.DB.Raw(topicSQL, id).Scan(&topics)

	// 查询标签
	tagSQL := `
		SELECT t.id, t.name 
		FROM tag t
		INNER JOIN content_tag ct ON t.id = ct.tag_id
		WHERE ct.content_id = ? AND t.deleted_at IS NULL
	`
	var tags []types.Tag
	l.svcCtx.DB.Raw(tagSQL, id).Scan(&tags)

	return &types.ContentResponse{
		ID:           result.ID,
		Title:        result.Title,
		Description:  result.Description,
		Cover:        result.Cover,
		Type:         result.Type,
		Status:       result.Status,
		ViewCount:    result.ViewCount,
		LikeCount:    result.LikeCount,
		CommentCount: result.CommentCount,
		CollectCount: result.CollectCount,
		ShareCount:   result.ShareCount,
		Media:        mediaList,
		Topics:       topics,
		Tags:         tags,
		Author: types.Author{
			ID:       result.AuthorID,
			Nickname: result.AuthorNickname,
			Avatar:   result.AuthorAvatar,
		},
		CreatedAt: result.CreatedAt,
	}, nil
}
