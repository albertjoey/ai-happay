package logic

import (
	"context"
	"happy/app/channel/internal/svc"
	"happy/app/channel/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChapterListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChapterListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChapterListLogic {
	return &ChapterListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

type ChapterItem struct {
	ID          uint     `json:"id"`
	MaterialID  uint     `json:"material_id"`
	ChapterType string   `json:"chapter_type"`
	Title       string   `json:"title"`
	Content     string   `json:"content"`
	Images      []string `json:"images"`
	VideoURL    string   `json:"video_url"`
	WordCount   int      `json:"word_count"`
	Duration    int      `json:"duration"`
	Sort        int      `json:"sort"`
	IsFree      int      `json:"is_free"`
	Price       int      `json:"price"`
}

type ChapterListResponse struct {
	List  []ChapterItem `json:"list"`
	Total int64         `json:"total"`
}

func (l *ChapterListLogic) ChapterList(materialID uint) (*ChapterListResponse, error) {
	// 使用Repository接口
	list, total, err := l.svcCtx.ChapterRepo.List(l.ctx, materialID, 1, 1000)
	if err != nil {
		return nil, err
	}

	// 转换响应格式
	items := make([]ChapterItem, 0, len(list))
	for _, ch := range list {
		items = append(items, ChapterItem{
			ID:          ch.ID,
			MaterialID:  ch.MaterialID,
			ChapterType: ch.ChapterType,
			Title:       ch.Title,
			Content:     ch.Content,
			VideoURL:    ch.VideoURL,
			WordCount:   int(ch.WordCount),
			Duration:    int(ch.Duration),
			Sort:        ch.Sort,
			IsFree:      int(ch.IsFree),
			Price:       int(ch.Price),
		})
	}

	return &ChapterListResponse{
		List:  items,
		Total: total,
	}, nil
}

type ChapterDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChapterDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChapterDetailLogic {
	return &ChapterDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

type ChapterDetailResponse struct {
	ID          uint     `json:"id"`
	MaterialID  uint     `json:"material_id"`
	ChapterType string   `json:"chapter_type"`
	Title       string   `json:"title"`
	Content     string   `json:"content"`
	Images      []string `json:"images"`
	VideoURL    string   `json:"video_url"`
	WordCount   int      `json:"word_count"`
	Duration    int      `json:"duration"`
	Sort        int      `json:"sort"`
	IsFree      int      `json:"is_free"`
	Price       int      `json:"price"`
	PrevID      uint     `json:"prev_id"`
	NextID      uint     `json:"next_id"`
}

func (l *ChapterDetailLogic) ChapterDetail(chapterID uint) (*ChapterDetailResponse, error) {
	// 使用Repository接口
	chapter, err := l.svcCtx.ChapterRepo.FindByID(l.ctx, chapterID)
	if err != nil {
		return nil, err
	}

	// 获取上一章/下一章
	prevID, nextID, err := l.svcCtx.ChapterRepo.FindPrevNext(l.ctx, chapterID)
	if err != nil {
		return nil, err
	}

	return &ChapterDetailResponse{
		ID:          chapter.ID,
		MaterialID:  chapter.MaterialID,
		ChapterType: chapter.ChapterType,
		Title:       chapter.Title,
		Content:     chapter.Content,
		VideoURL:    chapter.VideoURL,
		WordCount:   int(chapter.WordCount),
		Duration:    int(chapter.Duration),
		Sort:        chapter.Sort,
		IsFree:      int(chapter.IsFree),
		Price:       int(chapter.Price),
		PrevID:      prevID,
		NextID:      nextID,
	}, nil
}

// ==================== 章节管理CRUD ====================

type ChapterCreateRequest struct {
	MaterialID  uint     `json:"material_id"`
	ChapterType string   `json:"chapter_type"`
	Title       string   `json:"title"`
	Content     string   `json:"content"`
	Images      []string `json:"images"`
	VideoURL    string   `json:"video_url"`
	WordCount   int      `json:"word_count"`
	Duration    int      `json:"duration"`
	Sort        int      `json:"sort"`
	IsFree      int      `json:"is_free"`
	Price       int      `json:"price"`
}

type ChapterUpdateRequest struct {
	ChapterType string   `json:"chapter_type"`
	Title       string   `json:"title"`
	Content     string   `json:"content"`
	Images      []string `json:"images"`
	VideoURL    string   `json:"video_url"`
	WordCount   int      `json:"word_count"`
	Duration    int      `json:"duration"`
	Sort        int      `json:"sort"`
	IsFree      int      `json:"is_free"`
	Price       int      `json:"price"`
}

type ChapterCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChapterCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChapterCreateLogic {
	return &ChapterCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChapterCreateLogic) ChapterCreate(req *ChapterCreateRequest) (interface{}, error) {
	// 获取最大排序值
	maxSort, _ := l.svcCtx.ChapterRepo.GetMaxSort(l.ctx, req.MaterialID)
	sort := req.Sort
	if sort == 0 {
		sort = maxSort + 1
	}

	// 计算字数
	wordCount := req.WordCount
	if wordCount == 0 && req.Content != "" {
		wordCount = len([]rune(req.Content))
	}

	chapter := &types.Chapter{
		MaterialID:  req.MaterialID,
		ChapterType: req.ChapterType,
		Title:       req.Title,
		Content:     req.Content,
		VideoURL:    req.VideoURL,
		WordCount:   uint(wordCount),
		Duration:    uint(req.Duration),
		Sort:        sort,
		IsFree:      int8(req.IsFree),
		Price:       uint(req.Price),
	}

	err := l.svcCtx.ChapterRepo.Create(l.ctx, chapter)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{"id": chapter.ID, "success": true}, nil
}

type ChapterUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChapterUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChapterUpdateLogic {
	return &ChapterUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChapterUpdateLogic) ChapterUpdate(id uint, req *ChapterUpdateRequest) (interface{}, error) {
	// 计算字数
	wordCount := req.WordCount
	if wordCount == 0 && req.Content != "" {
		wordCount = len([]rune(req.Content))
	}

	chapter := &types.Chapter{
		ID:          id,
		ChapterType: req.ChapterType,
		Title:       req.Title,
		Content:     req.Content,
		VideoURL:    req.VideoURL,
		WordCount:   uint(wordCount),
		Duration:    uint(req.Duration),
		Sort:        req.Sort,
		IsFree:      int8(req.IsFree),
		Price:       uint(req.Price),
	}

	err := l.svcCtx.ChapterRepo.Update(l.ctx, chapter)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{"success": true}, nil
}

type ChapterDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChapterDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChapterDeleteLogic {
	return &ChapterDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChapterDeleteLogic) ChapterDelete(id uint) (interface{}, error) {
	err := l.svcCtx.ChapterRepo.Delete(l.ctx, id)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{"success": true}, nil
}
