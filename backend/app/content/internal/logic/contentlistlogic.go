package logic

import (
	"context"
	"encoding/json"
	"happy/app/content/internal/svc"
	"happy/app/content/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ContentListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewContentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ContentListLogic {
	return &ContentListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ContentListLogic) ContentList(req *types.ContentListRequest) (*types.ContentListResponse, error) {
	var total int64

	// 统计总数
	countSQL := "SELECT COUNT(*) FROM content WHERE tenant_id = ? AND deleted_at IS NULL"
	l.svcCtx.DB.Raw(countSQL, 1).Scan(&total)

	// 构建查询SQL
	querySQL := `
		SELECT c.id, c.title, c.description, c.cover, c.type, c.status, 
		       c.view_count, c.like_count, c.comment_count, c.collect_count, c.share_count,
		       c.created_at, u.id as author_id, u.nickname as author_nickname, u.avatar as author_avatar
		FROM content c
		LEFT JOIN user u ON c.user_id = u.id
		WHERE c.tenant_id = ? AND c.deleted_at IS NULL
	`
	args := []interface{}{1}

	// 条件过滤
	if req.Type != 0 {
		querySQL += " AND c.type = ?"
		args = append(args, req.Type)
	}
	if req.Status != 0 {
		querySQL += " AND c.status = ?"
		args = append(args, req.Status)
	}
	if req.UserID != 0 {
		querySQL += " AND c.user_id = ?"
		args = append(args, req.UserID)
	}

	// 排序和分页
	querySQL += " ORDER BY c.id DESC LIMIT ? OFFSET ?"
	offset := (req.Page - 1) * req.PageSize
	args = append(args, req.PageSize, offset)

	// 执行查询
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
		CreatedAt     string
		AuthorID      uint
		AuthorNickname string
		AuthorAvatar  string
	}
	var results []ContentResult
	err := l.svcCtx.DB.Raw(querySQL, args...).Scan(&results).Error

	if err != nil {
		return nil, err
	}

	// 转换响应
	list := make([]types.ContentResponse, 0, len(results))
	for _, item := range results {
		// 查询媒体资源
		var mediaJSON string
		l.svcCtx.DB.Raw("SELECT media FROM content WHERE id = ?", item.ID).Scan(&mediaJSON)
		
		var mediaList []types.Media
		if mediaJSON != "" {
			json.Unmarshal([]byte(mediaJSON), &mediaList)
		}

		// 查询话题
		topicSQL := `
			SELECT t.id, t.name 
			FROM topic t
			INNER JOIN content_topic ct ON t.id = ct.topic_id
			WHERE ct.content_id = ? AND t.deleted_at IS NULL
		`
		var topics []types.Topic
		l.svcCtx.DB.Raw(topicSQL, item.ID).Scan(&topics)

		// 查询标签
		tagSQL := `
			SELECT t.id, t.name 
			FROM tag t
			INNER JOIN content_tag ct ON t.id = ct.tag_id
			WHERE ct.content_id = ? AND t.deleted_at IS NULL
		`
		var tags []types.Tag
		l.svcCtx.DB.Raw(tagSQL, item.ID).Scan(&tags)

		list = append(list, types.ContentResponse{
			ID:           item.ID,
			Title:        item.Title,
			Description:  item.Description,
			Cover:        item.Cover,
			Type:         item.Type,
			Status:       item.Status,
			ViewCount:    item.ViewCount,
			LikeCount:    item.LikeCount,
			CommentCount: item.CommentCount,
			CollectCount: item.CollectCount,
			ShareCount:   item.ShareCount,
			Media:        mediaList,
			Topics:       topics,
			Tags:         tags,
			Author: types.Author{
				ID:       item.AuthorID,
				Nickname: item.AuthorNickname,
				Avatar:   item.AuthorAvatar,
			},
			CreatedAt: item.CreatedAt,
		})
	}

	return &types.ContentListResponse{
		Total: total,
		List:  list,
	}, nil
}
