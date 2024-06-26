package hello

import (
	"go_zero_demo/http/consts_code"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go_zero_demo/http/internal/logic/hello"
	"go_zero_demo/http/internal/svc"
	"go_zero_demo/http/internal/types"
)

func AddHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddHelloReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJsonCtx(r.Context(), w, map[string]interface{}{
				"code": consts_code.ReqParamsErr,
				"msg":  consts_code.ReqParamsErrDoc + err.Error(),
			})
			return
		}
		l := hello.NewAddLogic(r.Context(), svcCtx)
		resp, err := l.Add(&req)
		if err != nil {
			httpx.OkJsonCtx(r.Context(), w, map[string]interface{}{
				"code": consts_code.BusinessErr,
				"msg":  consts_code.BusinessErrDoc + err.Error(),
			})
		} else {
			httpx.OkJsonCtx(r.Context(), w, map[string]interface{}{
				"code": 0,
				"data": resp,
			})
		}
	}
}
