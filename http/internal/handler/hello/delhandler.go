package hello

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go_zero_demo/http/internal/logic/hello"
	"go_zero_demo/http/internal/svc"
	"go_zero_demo/http/internal/types"
)

func DelHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DelHelloReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := hello.NewDelLogic(r.Context(), svcCtx)
		resp, err := l.Del(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
