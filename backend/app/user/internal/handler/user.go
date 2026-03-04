package handler

import (
	"net/http"

	"happy/app/user/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/api/v1/user/register",
				Handler: RegisterHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/v1/user/login",
				Handler: LoginHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/v1/user/info",
				Handler: GetUserInfoHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/api/v1/user/info",
				Handler: UpdateUserHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/v1/user/follow",
				Handler: FollowHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/v1/user/list",
				Handler: GetUserListHandler(serverCtx),
			},
		},
	)
}
