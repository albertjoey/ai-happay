package handler

import (
	"net/http"

	"happy/app/channel/internal/logic"
	"happy/app/channel/internal/svc"
	"happy/app/channel/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func DiscoverConfigListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewDiscoverConfigLogic(r.Context(), svcCtx)
		resp, err := l.GetConfigList()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func DiscoverConfigUpdateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DiscoverConfigUpdateRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewDiscoverConfigLogic(r.Context(), svcCtx)
		err := l.UpdateConfig(&req)
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
