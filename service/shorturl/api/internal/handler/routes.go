// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	shorturl "github.com/kiritokun07/go-zero-study/service/shorturl/api/internal/handler/shorturl"
	"github.com/kiritokun07/go-zero-study/service/shorturl/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/shorten",
				Handler: shorturl.ShortenHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/expand",
				Handler: shorturl.ExpandHandler(serverCtx),
			},
		},
		rest.WithPrefix("/shorturl"),
	)
}
