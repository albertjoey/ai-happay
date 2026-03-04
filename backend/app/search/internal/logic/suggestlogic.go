package logic

import (
	"context"

	"happy/app/search/internal/svc"
	"happy/app/search/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SuggestLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSuggestLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SuggestLogic {
	return &SuggestLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SuggestLogic) Suggest(req *types.SearchSuggestRequest) (*types.SearchSuggestResponse, error) {
	// 获取租户ID
	tenantID := uint(1)
	if tid := l.ctx.Value("tenant_id"); tid != nil {
		tenantID = tid.(uint)
	}

	var suggestions []string
	var err error

	// 优先使用ES，降级到数据库
	if l.svcCtx.ES != nil {
		suggestions, err = l.svcCtx.ES.Suggest(req.Keyword, tenantID, req.Limit)
	} else {
		suggestions, err = l.svcCtx.DBSearch.Suggest(req.Keyword, tenantID, req.Limit)
	}

	if err != nil {
		return nil, err
	}

	return &types.SearchSuggestResponse{
		List: suggestions,
	}, nil
}
