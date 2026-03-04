package handler

import (
	"net/http"
	"strconv"

	"happy/app/content/internal/logic"
	"happy/app/content/internal/svc"
	"happy/app/content/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func ContentListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ContentListRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewContentListLogic(r.Context(), svcCtx)
		resp, err := l.ContentList(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func ContentCreateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateContentRequest
		if err := httpx.ParseJsonBody(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewContentCreateLogic(r.Context(), svcCtx)
		resp, err := l.ContentCreate(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func ContentUpdateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateContentRequest
		if err := httpx.ParseJsonBody(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewContentUpdateLogic(r.Context(), svcCtx)
		resp, err := l.ContentUpdate(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func ContentDeleteHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := r.PathValue("id")
		if idStr == "" {
			idStr = r.URL.Query().Get("id")
		}
		id, err := strconv.ParseUint(idStr, 10, 64)
		if err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewContentDeleteLogic(r.Context(), svcCtx)
		err = l.ContentDelete(uint(id))
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, map[string]interface{}{"success": true})
		}
	}
}

func ContentDetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ContentDetailRequest
		if err := httpx.ParsePath(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewContentDetailLogic(r.Context(), svcCtx)
		resp, err := l.ContentDetail(req.ID)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func ContentPublishHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := r.PathValue("id")
		if idStr == "" {
			idStr = r.URL.Query().Get("id")
		}
		id, err := strconv.ParseUint(idStr, 10, 64)
		if err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewContentPublishLogic(r.Context(), svcCtx)
		err = l.ContentPublish(uint(id))
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, map[string]interface{}{"success": true})
		}
	}
}

func ContentUnpublishHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := r.PathValue("id")
		if idStr == "" {
			idStr = r.URL.Query().Get("id")
		}
		id, err := strconv.ParseUint(idStr, 10, 64)
		if err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewContentUnpublishLogic(r.Context(), svcCtx)
		err = l.ContentUnpublish(uint(id))
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, map[string]interface{}{"success": true})
		}
	}
}
