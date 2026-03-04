package handler

import (
	"happy/app/recommend/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  "POST",
				Path:    "/api/v1/recommend",
				Handler: RecommendHandler(serverCtx),
			},
			{
				Method:  "POST",
				Path:    "/api/v1/recommend/algorithm",
				Handler: AlgorithmRecommendHandler(serverCtx),
			},
			{
				Method:  "POST",
				Path:    "/api/v1/recommend/manual",
				Handler: ManualRecommendHandler(serverCtx),
			},
			{
				Method:  "POST",
				Path:    "/api/v1/recommend/random",
				Handler: RandomRecommendHandler(serverCtx),
			},
			{
				Method:  "POST",
				Path:    "/api/v1/recommend/filter",
				Handler: FilterRecommendHandler(serverCtx),
			},
			{
				Method:  "GET",
				Path:    "/api/v1/ranking/hot",
				Handler: RankingHotHandler(serverCtx),
			},
			{
				Method:  "GET",
				Path:    "/api/v1/ranking/new",
				Handler: RankingNewHandler(serverCtx),
			},
			{
				Method:  "POST",
				Path:    "/api/v1/hot-score/update",
				Handler: HotScoreUpdateHandler(serverCtx),
			},
			{
				Method:  "POST",
				Path:    "/api/v1/hot-score/update/:id",
				Handler: HotScoreUpdateOneHandler(serverCtx),
			},
		},
	)
}
