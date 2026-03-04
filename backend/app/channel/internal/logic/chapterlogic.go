package logic

import (
	"context"
	"encoding/json"
	"time"

	"happy/app/channel/internal/svc"

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
	ID          uint                   `json:"id"`
	MaterialID  uint                   `json:"material_id"`
	ChapterType string                 `json:"chapter_type"`
	Title       string                 `json:"title"`
	Content     string                 `json:"content"`
	Images      []string               `json:"images"`
	VideoURL    string                 `json:"video_url"`
	WordCount   int                    `json:"word_count"`
	Duration    int                    `json:"duration"`
	Sort        int                    `json:"sort"`
	IsFree      int                    `json:"is_free"`
	Price       int                    `json:"price"`
}

type ChapterListResponse struct {
	List  []ChapterItem `json:"list"`
	Total int64         `json:"total"`
}

func (l *ChapterListLogic) ChapterList(materialID uint) (*ChapterListResponse, error) {
	query := `SELECT id, material_id, chapter_type, title, content, images, video_url, word_count, duration, sort, is_free, price 
			  FROM material_chapter 
			  WHERE material_id = ? AND deleted_at IS NULL AND status = 1 
			  ORDER BY sort ASC`

	var list []ChapterItem
	err := l.svcCtx.DB.Raw(query, materialID).Scan(&list).Error
	if err != nil {
		return nil, err
	}

	// 获取总数
	var total int64
	countQuery := `SELECT COUNT(*) FROM material_chapter WHERE material_id = ? AND deleted_at IS NULL AND status = 1`
	l.svcCtx.DB.Raw(countQuery, materialID).Scan(&total)

	return &ChapterListResponse{
		List:  list,
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
	query := `SELECT id, material_id, chapter_type, title, content, images, video_url, word_count, duration, sort, is_free, price 
			  FROM material_chapter 
			  WHERE id = ? AND deleted_at IS NULL AND status = 1`

	var resp ChapterDetailResponse
	err := l.svcCtx.DB.Raw(query, chapterID).Scan(&resp).Error
	if err != nil {
		return nil, err
	}

	// 获取上一章ID
	var prevID *uint
	prevQuery := `SELECT id FROM material_chapter WHERE material_id = ? AND sort < ? AND deleted_at IS NULL AND status = 1 ORDER BY sort DESC LIMIT 1`
	l.svcCtx.DB.Raw(prevQuery, resp.MaterialID, resp.Sort).Scan(&prevID)
	if prevID != nil {
		resp.PrevID = *prevID
	}

	// 获取下一章ID
	var nextID *uint
	nextQuery := `SELECT id FROM material_chapter WHERE material_id = ? AND sort > ? AND deleted_at IS NULL AND status = 1 ORDER BY sort ASC LIMIT 1`
	l.svcCtx.DB.Raw(nextQuery, resp.MaterialID, resp.Sort).Scan(&nextID)
	if nextID != nil {
		resp.NextID = *nextID
	}

	return &resp, nil
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
	now := time.Now()
	
	// 如果没有指定排序，获取当前最大排序+1
	if req.Sort == 0 {
		var maxSort int
		l.svcCtx.DB.Raw("SELECT COALESCE(MAX(sort), 0) FROM material_chapter WHERE material_id = ? AND deleted_at IS NULL", req.MaterialID).Scan(&maxSort)
		req.Sort = maxSort + 1
	}
	
	// 如果没有指定字数，计算字数
	if req.WordCount == 0 && req.Content != "" {
		req.WordCount = len([]rune(req.Content))
	}
	
	// 序列化images
	imagesJSON, _ := json.Marshal(req.Images)
	
	insertSQL := `
		INSERT INTO material_chapter (material_id, chapter_type, title, content, images, video_url, word_count, duration, sort, is_free, price, status, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, 1, ?, ?)
	`
	result := l.svcCtx.DB.Exec(insertSQL, req.MaterialID, req.ChapterType, req.Title, req.Content, imagesJSON, req.VideoURL, req.WordCount, req.Duration, req.Sort, req.IsFree, req.Price, now, now)
	if result.Error != nil {
		return nil, result.Error
	}
	
	return map[string]interface{}{"id": result.RowsAffected, "success": true}, nil
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
	now := time.Now()
	
	// 如果没有指定字数，计算字数
	if req.WordCount == 0 && req.Content != "" {
		req.WordCount = len([]rune(req.Content))
	}
	
	// 序列化images
	imagesJSON, _ := json.Marshal(req.Images)
	
	updateSQL := `
		UPDATE material_chapter 
		SET chapter_type = ?, title = ?, content = ?, images = ?, video_url = ?, word_count = ?, duration = ?, sort = ?, is_free = ?, price = ?, updated_at = ?
		WHERE id = ? AND deleted_at IS NULL
	`
	result := l.svcCtx.DB.Exec(updateSQL, req.ChapterType, req.Title, req.Content, imagesJSON, req.VideoURL, req.WordCount, req.Duration, req.Sort, req.IsFree, req.Price, now, id)
	if result.Error != nil {
		return nil, result.Error
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
	now := time.Now()
	
	deleteSQL := `UPDATE material_chapter SET deleted_at = ? WHERE id = ?`
	result := l.svcCtx.DB.Exec(deleteSQL, now, id)
	if result.Error != nil {
		return nil, result.Error
	}
	
	return map[string]interface{}{"success": true}, nil
}
