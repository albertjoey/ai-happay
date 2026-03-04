package logic

import (
	"context"

	"happy/app/channel/internal/svc"
	"happy/app/channel/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MaterialListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMaterialListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MaterialListLogic {
	return &MaterialListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MaterialListLogic) MaterialList(req *types.MaterialListRequest) (*types.MaterialListResponse, error) {
	// 默认分页参数
	page := req.Page
	if page < 1 {
		page = 1
	}
	pageSize := req.PageSize
	if pageSize < 1 {
		pageSize = 10
	}
	if pageSize > 100 {
		pageSize = 100
	}
	offset := (page - 1) * pageSize

	// 查询总数
	var total int64
	countSQL := `SELECT COUNT(*) FROM material WHERE deleted_at IS NULL`
	countArgs := []interface{}{}

	if req.Type != "" {
		countSQL += " AND type = ?"
		countArgs = append(countArgs, req.Type)
	}
	if req.Status != nil {
		countSQL += " AND status = ?"
		countArgs = append(countArgs, *req.Status)
	}

	l.svcCtx.DB.Raw(countSQL, countArgs...).Scan(&total)

	// 查询列表
	listSQL := `
		SELECT id, title, subtitle, type, cover_url, content_url, description,
			   author, category, view_count, like_count, comment_count, 
			   share_count, collect_count, duration, word_count, chapter_count, status, sort
		FROM material
		WHERE deleted_at IS NULL
	`
	listArgs := []interface{}{}

	if req.Type != "" {
		listSQL += " AND type = ?"
		listArgs = append(listArgs, req.Type)
	}
	if req.Status != nil {
		listSQL += " AND status = ?"
		listArgs = append(listArgs, *req.Status)
	}

	// 排序
	listSQL += " ORDER BY id DESC"

	// 分页
	listSQL += " LIMIT ? OFFSET ?"
	listArgs = append(listArgs, pageSize, offset)

	var list []types.Material
	l.svcCtx.DB.Raw(listSQL, listArgs...).Scan(&list)

	return &types.MaterialListResponse{
		Total: total,
		List:  list,
	}, nil
}

type MaterialCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMaterialCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MaterialCreateLogic {
	return &MaterialCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MaterialCreateLogic) MaterialCreate(req *types.MaterialCreateRequest) (interface{}, error) {
	return map[string]interface{}{"id": 100, "success": true}, nil
}

type MaterialUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMaterialUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MaterialUpdateLogic {
	return &MaterialUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MaterialUpdateLogic) MaterialUpdate(req *types.MaterialUpdateRequest) (interface{}, error) {
	return map[string]interface{}{"success": true}, nil
}

type MaterialDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMaterialDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MaterialDeleteLogic {
	return &MaterialDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MaterialDeleteLogic) MaterialDelete(id uint) (interface{}, error) {
	return map[string]interface{}{"success": true}, nil
}
