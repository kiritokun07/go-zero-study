// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	user "github.com/kiritokun07/go-zero-study/service/demo/api/internal/handler/user"
	"github.com/kiritokun07/go-zero-study/service/demo/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/",
				Handler: user.LoginHandler(serverCtx),
			},
		},
		rest.WithPrefix("/demo"),
	)
}