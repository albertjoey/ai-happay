package handler

import (
	"happy/app/content/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	// 内容管理API
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  "GET",
				Path:    "/api/v1/content/list",
				Handler: ContentListHandler(serverCtx),
			},
			{
				Method:  "POST",
				Path:    "/api/v1/content/create",
				Handler: ContentCreateHandler(serverCtx),
			},
			{
				Method:  "PUT",
				Path:    "/api/v1/content/update",
				Handler: ContentUpdateHandler(serverCtx),
			},
			{
				Method:  "DELETE",
				Path:    "/api/v1/content/{id}",
				Handler: ContentDeleteHandler(serverCtx),
			},
			{
				Method:  "GET",
				Path:    "/api/v1/content/{id}",
				Handler: ContentDetailHandler(serverCtx),
			},
			{
				Method:  "POST",
				Path:    "/api/v1/content/{id}/publish",
				Handler: ContentPublishHandler(serverCtx),
			},
			{
				Method:  "POST",
				Path:    "/api/v1/content/{id}/unpublish",
				Handler: ContentUnpublishHandler(serverCtx),
			},
		},
	)
}
