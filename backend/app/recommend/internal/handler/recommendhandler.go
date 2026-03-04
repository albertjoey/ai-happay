package handler

import (
	"net/http"
	"strconv"

	"happy/app/recommend/internal/logic"
	"happy/app/recommend/internal/svc"
	"happy/app/recommend/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func RecommendHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RecommendRequest
		if err := httpx.ParseJsonBody(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewRecommendLogic(r.Context(), svcCtx)
		resp, err := l.Recommend(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

// AlgorithmRecommendHandler 算法推荐
func AlgorithmRecommendHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RecommendRequest
		if err := httpx.ParseJsonBody(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		req.Strategy = types.StrategyAlgorithm
		l := logic.NewRecommendLogic(r.Context(), svcCtx)
		resp, err := l.Recommend(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

// ManualRecommendHandler 人工推荐
func ManualRecommendHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RecommendRequest
		if err := httpx.ParseJsonBody(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		req.Strategy = types.StrategyManual
		l := logic.NewRecommendLogic(r.Context(), svcCtx)
		resp, err := l.Recommend(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

// RandomRecommendHandler 随机推荐
func RandomRecommendHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RecommendRequest
		if err := httpx.ParseJsonBody(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		req.Strategy = types.StrategyRandom
		l := logic.NewRecommendLogic(r.Context(), svcCtx)
		resp, err := l.Recommend(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

// FilterRecommendHandler 条件筛选
func FilterRecommendHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RecommendRequest
		if err := httpx.ParseJsonBody(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		req.Strategy = types.StrategyFilter
		l := logic.NewRecommendLogic(r.Context(), svcCtx)
		resp, err := l.Recommend(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func RankingHotHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RecommendRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		req.Strategy = types.StrategyAlgorithm
		req.AlgorithmType = types.AlgorithmHot
		req.Limit = 100

		l := logic.NewRecommendLogic(r.Context(), svcCtx)
		resp, err := l.Recommend(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, &types.RankingResponse{List: resp.Content})
		}
	}
}

func RankingNewHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RecommendRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		req.Strategy = types.StrategyAlgorithm
		req.AlgorithmType = types.AlgorithmTime
		req.Limit = 100

		l := logic.NewRecommendLogic(r.Context(), svcCtx)
		resp, err := l.Recommend(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, &types.RankingResponse{List: resp.Content})
		}
	}
}

// HotScoreUpdateHandler 手动更新所有热度分数
func HotScoreUpdateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := svcCtx.UpdateAllHotScores()
		if err != nil {
			httpx.Error(w, err)
			return
		}
		httpx.OkJson(w, &types.HotScoreUpdateResponse{
			Success: true,
			Message: "热度分数更新成功",
		})
	}
}

// HotScoreUpdateOneHandler 更新单个内容的热度分数
func HotScoreUpdateOneHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := r.PathValue("id")
		id, err := strconv.ParseUint(idStr, 10, 64)
		if err != nil {
			httpx.Error(w, err)
			return
		}

		err = svcCtx.UpdateHotScore(uint(id))
		if err != nil {
			httpx.Error(w, err)
			return
		}
		httpx.OkJson(w, &types.HotScoreUpdateResponse{
			Success: true,
			Message: "热度分数更新成功",
		})
	}
}
