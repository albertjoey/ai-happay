package logic

import (
	"context"

	"happy/app/search/internal/svc"
	"happy/app/search/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchLogic {
	return &SearchLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchLogic) Search(req *types.SearchRequest) (*types.SearchResponse, error) {
	// 获取租户ID（从context中获取，默认为1）
	tenantID := uint(1)
	if tid := l.ctx.Value("tenant_id"); tid != nil {
		tenantID = tid.(uint)
	}

	// 优先使用ES搜索，如果ES不可用则使用数据库搜索
	if l.svcCtx.ES != nil {
		return l.svcCtx.ES.Search(req.Keyword, tenantID, req.Page, req.PageSize)
	}

	return l.svcCtx.DBSearch.Search(req.Keyword, tenantID, req.Page, req.PageSize)
}
