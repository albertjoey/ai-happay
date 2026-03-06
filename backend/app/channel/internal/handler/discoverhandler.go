package handler

import (
	"net/http"

	"happy/app/channel/internal/logic"
	"happy/app/channel/internal/svc"
	"happy/app/channel/internal/types"
	"happy/common/result"

	"github.com/zeromicro/go-zero/rest/httpx"
)

// GetDiscoverPageHandler 获取发现页完整数据
func GetDiscoverPageHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DiscoverPageRequest
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := logic.NewDiscoverLogic(r.Context(), svcCtx)
		resp, err := l.GetDiscoverPage(&req)
		result.HttpResult(r, w, resp, err)
	}
}

// GetDiscoverConfigListHandler 获取发现页配置列表
func GetDiscoverConfigListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DiscoverConfigListRequest
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := logic.NewDiscoverLogic(r.Context(), svcCtx)
		resp, err := l.GetDiscoverConfigList(&req)
		result.HttpResult(r, w, resp, err)
	}
}

// UpdateDiscoverConfigHandler 更新发现页配置
func UpdateDiscoverConfigHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DiscoverConfigUpdateRequest
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := logic.NewDiscoverLogic(r.Context(), svcCtx)
		err := l.UpdateDiscoverConfig(&req)
		result.HttpResult(r, w, nil, err)
	}
}

// GetDiscoverItemListHandler 获取发现页内容列表
func GetDiscoverItemListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DiscoverItemListRequest
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := logic.NewDiscoverLogic(r.Context(), svcCtx)
		resp, err := l.GetDiscoverItemList(&req)
		result.HttpResult(r, w, resp, err)
	}
}

// CreateDiscoverItemHandler 创建发现页内容
func CreateDiscoverItemHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DiscoverItemCreateRequest
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := logic.NewDiscoverLogic(r.Context(), svcCtx)
		err := l.CreateDiscoverItem(&req)
		result.HttpResult(r, w, nil, err)
	}
}

// UpdateDiscoverItemHandler 更新发现页内容
func UpdateDiscoverItemHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DiscoverItemUpdateRequest
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := logic.NewDiscoverLogic(r.Context(), svcCtx)
		err := l.UpdateDiscoverItem(&req)
		result.HttpResult(r, w, nil, err)
	}
}

// DeleteDiscoverItemHandler 删除发现页内容
func DeleteDiscoverItemHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			ID uint `path:"id"`
		}
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := logic.NewDiscoverLogic(r.Context(), svcCtx)
		err := l.DeleteDiscoverItem(req.ID)
		result.HttpResult(r, w, nil, err)
	}
}
