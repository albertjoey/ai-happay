package handler

import (
	"net/http"
	"strconv"

	"happy/app/channel/internal/logic"
	"happy/app/channel/internal/svc"
	"happy/app/channel/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func ChannelListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ChannelListRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewChannelListLogic(r.Context(), svcCtx)
		resp, err := l.ChannelList(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func ChannelCreateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ChannelCreateRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewChannelCreateLogic(r.Context(), svcCtx)
		resp, err := l.ChannelCreate(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func ChannelUpdateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ChannelUpdateRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewChannelUpdateLogic(r.Context(), svcCtx)
		resp, err := l.ChannelUpdate(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func ChannelDeleteHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 从URL路径中获取ID
		idStr := r.PathValue("id")
		if idStr == "" {
			// 尝试从查询参数获取
			idStr = r.URL.Query().Get("id")
		}
		id, err := strconv.ParseUint(idStr, 10, 64)
		if err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewChannelDeleteLogic(r.Context(), svcCtx)
		err = l.ChannelDelete(uint(id))
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, map[string]interface{}{"success": true})
		}
	}
}

func ChannelSortHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ChannelSortRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewChannelSortLogic(r.Context(), svcCtx)
		err := l.ChannelSort(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, map[string]interface{}{"success": true})
		}
	}
}

func ChannelConfigGetHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 从URL路径中获取ID
		var req types.ChannelConfigRequest
		if err := httpx.ParsePath(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewChannelConfigLogic(r.Context(), svcCtx)
		resp, err := l.GetChannelConfig(req.ChannelID)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func ChannelConfigUpdateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 从URL路径中获取ID
		var pathReq types.ChannelConfigRequest
		if err := httpx.ParsePath(r, &pathReq); err != nil {
			httpx.Error(w, err)
			return
		}

		var req types.ChannelConfigResponse
		if err := httpx.ParseJsonBody(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewChannelConfigLogic(r.Context(), svcCtx)
		err := l.UpdateChannelConfig(pathReq.ChannelID, &req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, map[string]interface{}{"success": true})
		}
	}
}

// BannerListHandler Banner列表
func BannerListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		channelIDStr := r.URL.Query().Get("channel_id")
		channelID, _ := strconv.ParseUint(channelIDStr, 10, 64)
		statusStr := r.URL.Query().Get("status")
		status, _ := strconv.Atoi(statusStr)

		l := logic.NewBannerListLogic(r.Context(), svcCtx)
		resp, err := l.BannerList(uint(channelID), status)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
