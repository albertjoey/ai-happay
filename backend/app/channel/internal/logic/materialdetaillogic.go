package logic

import (
	"context"
	"happy/app/channel/internal/svc"
	"happy/app/channel/internal/types"

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
	// 使用Repository接口
	material, err := l.svcCtx.MaterialRepo.FindByID(l.ctx, id)
	if err != nil {
		return nil, err
	}

	return &MaterialDetailResponse{
		ID:           material.ID,
		Title:        material.Title,
		Subtitle:     material.Subtitle,
		Type:         material.Type,
		CoverURL:     material.CoverURL,
		ContentURL:   material.ContentURL,
		Description:  material.Description,
		Author:       material.Author,
		Category:     material.Category,
		ViewCount:    int(material.ViewCount),
		LikeCount:    int(material.LikeCount),
		CommentCount: int(material.CommentCount),
		ShareCount:   int(material.ShareCount),
		CollectCount: int(material.CollectCount),
		Duration:     int(material.Duration),
		WordCount:    int(material.WordCount),
		ChapterCount: int(material.ChapterCount),
		Status:       int(material.Status),
	}, nil
}

// MaterialDetailResponse2 返回types.Material的响应
type MaterialDetailResponse2 struct {
	Material types.Material `json:"material"`
}

func (l *MaterialDetailLogic) MaterialDetail2(id uint) (*MaterialDetailResponse2, error) {
	// 使用Repository接口
	material, err := l.svcCtx.MaterialRepo.FindByID(l.ctx, id)
	if err != nil {
		return nil, err
	}

	return &MaterialDetailResponse2{
		Material: *material,
	}, nil
}
