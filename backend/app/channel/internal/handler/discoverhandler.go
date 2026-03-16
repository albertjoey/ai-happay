package handler

import (
	"net/http"

	"happy/app/channel/internal/logic"
	"happy/app/channel/internal/svc"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func DiscoverHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewDiscoverLogic(r.Context(), svcCtx)
		resp, err := l.Discover()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
