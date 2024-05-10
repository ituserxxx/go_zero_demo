package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"http/internal/logic"
	"http/internal/svc"
	"http/internal/types"
)

func HttpHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewHttpLogic(r.Context(), svcCtx)
		resp, err := l.Http(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
