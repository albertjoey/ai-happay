package handler

import (
	"net/http"

	"happy/app/search/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/api/v1/search",
				Handler: SearchHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/v1/search/suggest",
				Handler: SuggestHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/v1/search/sync",
				Handler: SyncContentHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/v1/search/sync-all",
				Handler: SyncAllHandler(serverCtx),
			},
		},
	)
}
