package logic

import (
	"context"

	"happy/app/channel/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type MaterialDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMaterialDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MaterialDetailLogic {
	return &MaterialDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

type MaterialDetailResponse struct {
	ID           uint   `json:"id"`
	Title        string `json:"title"`
	Subtitle     string `json:"subtitle"`
	Type         string `json:"type"`
	CoverURL     string `json:"cover_url"`
	ContentURL   string `json:"content_url"`
	Description  string `json:"description"`
	Author       string `json:"author"`
	Tags         string `json:"tags"`
	Category     string `json:"category"`
	ViewCount    int    `json:"view_count"`
	LikeCount    int    `json:"like_count"`
	CommentCount int    `json:"comment_count"`
	ShareCount   int    `json:"share_count"`
	CollectCount int    `json:"collect_count"`
	Duration     int    `json:"duration"`
	WordCount    int    `json:"word_count"`
	ChapterCount int    `json:"chapter_count"`
	Status       int    `json:"status"`
}

func (l *MaterialDetailLogic) MaterialDetail(id uint) (*MaterialDetailResponse, error) {
	query := `
		SELECT id, title, subtitle, type, cover_url, content_url, description,
		       author, tags, category, view_count, like_count, comment_count,
		       share_count, collect_count, duration, word_count, chapter_count, status
		FROM material
		WHERE id = ? AND deleted_at IS NULL
	`

	var resp MaterialDetailResponse
	err := l.svcCtx.DB.Raw(query, id).Scan(&resp).Error
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
