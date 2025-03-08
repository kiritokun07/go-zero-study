package handler

import (
	"github.com/kiritokun07/go-zero-study/service/xxljob/api/internal/logic"
	"github.com/kiritokun07/go-zero-study/service/xxljob/api/internal/svc"
	"github.com/kiritokun07/go-zero-study/service/xxljob/api/internal/types"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func XxljobHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewXxljobLogic(r.Context(), svcCtx)
		resp, err := l.Xxljob(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
