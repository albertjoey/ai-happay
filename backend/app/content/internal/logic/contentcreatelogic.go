package logic

import (
	"context"
	"encoding/json"
	"happy/app/content/internal/svc"
	"happy/app/content/internal/types"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type ContentCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewContentCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ContentCreateLogic {
	return &ContentCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ContentCreateLogic) ContentCreate(req *types.CreateContentRequest) (*types.ContentResponse, error) {
	// 将媒体资源转为JSON
	mediaJSON, err := json.Marshal(req.Media)
	if err != nil {
		return nil, err
	}

	// 插入内容
	now := time.Now()
	insertSQL := `
		INSERT INTO content (tenant_id, user_id, title, description, cover, type, status, media, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`
	result := l.svcCtx.DB.Exec(insertSQL, 1, 1, req.Title, req.Description, req.Cover, req.Type, 0, string(mediaJSON), now, now)
	if result.Error != nil {
		return nil, result.Error
	}

	// 获取插入的ID
	var contentID uint
	l.svcCtx.DB.Raw("SELECT LAST_INSERT_ID()").Scan(&contentID)

	// 插入话题关联
	if len(req.Topics) > 0 {
		for _, topicID := range req.Topics {
			l.svcCtx.DB.Exec("INSERT INTO content_topic (content_id, topic_id, created_at) VALUES (?, ?, ?)", contentID, topicID, now)
		}
	}

	// 插入标签关联
	if len(req.Tags) > 0 {
		for _, tagID := range req.Tags {
			l.svcCtx.DB.Exec("INSERT INTO content_tag (content_id, tag_id, created_at) VALUES (?, ?, ?)", contentID, tagID, now)
		}
	}

	// 返回创建的内容
	return l.getContentByID(contentID)
}

func (l *ContentCreateLogic) getContentByID(id uint) (*types.ContentResponse, error) {
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
