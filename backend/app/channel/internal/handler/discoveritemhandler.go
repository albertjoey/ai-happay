package handler

import (
	"net/http"
	"strconv"

	"happy/app/channel/internal/logic"
	"happy/app/channel/internal/svc"
	"happy/app/channel/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func DiscoverItemListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DiscoverItemListRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewDiscoverItemLogic(r.Context(), svcCtx)
		resp, err := l.GetItemList(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func DiscoverItemCreateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DiscoverItemCreateRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewDiscoverItemLogic(r.Context(), svcCtx)
		err := l.CreateItem(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, map[string]interface{}{
				"success": true,
				"message": "创建成功",
			})
		}
	}
}

func DiscoverItemUpdateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DiscoverItemUpdateRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewDiscoverItemLogic(r.Context(), svcCtx)
		err := l.UpdateItem(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, map[string]interface{}{
				"success": true,
				"message": "更新成功",
			})
		}
	}
}

func DiscoverItemDeleteHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := r.PathValue("id")
		id, _ := strconv.ParseUint(idStr, 10, 32)

		l := logic.NewDiscoverItemLogic(r.Context(), svcCtx)
		err := l.DeleteItem(uint(id))
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, map[string]interface{}{
				"success": true,
				"message": "删除成功",
			})
		}
	}
}
