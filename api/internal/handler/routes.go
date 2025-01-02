// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3

package handler

import (
	"net/http"

	"zero-online-conf/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/ping",
				Handler: PingHandler(serverCtx),
			},
			{
				// 用户登录
				Method:  http.MethodPost,
				Path:    "/user/v1/userLogin",
				Handler: UserLoginHandler(serverCtx),
			},
		},
	)
}
