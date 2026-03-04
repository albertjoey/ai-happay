package handler

import (
	"net/http"

	"happy/app/search/internal/logic"
	"happy/app/search/internal/svc"
	"happy/app/search/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func SearchHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SearchRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewSearchLogic(r.Context(), svcCtx)
		resp, err := l.Search(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

func SuggestHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SearchSuggestRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewSuggestLogic(r.Context(), svcCtx)
		resp, err := l.Suggest(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

func SyncContentHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			ContentID uint `json:"content_id"`
		}
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewSyncLogic(r.Context(), svcCtx)
		err := l.SyncContent(req.ContentID)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, map[string]string{"message": "同步成功"})
		}
	}
}

func SyncAllHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewSyncLogic(r.Context(), svcCtx)
		err := l.SyncAll(1) // 默认同步租户1
		if err != nil {
			httpx.OkJsonCtx(r.Context(), w, map[string]string{"message": err.Error()})
		} else {
			httpx.OkJsonCtx(r.Context(), w, map[string]string{"message": "同步完成"})
		}
	}
}
