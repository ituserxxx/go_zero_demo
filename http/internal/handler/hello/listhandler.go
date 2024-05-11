package hello

import (
	"http/consts_code"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"http/internal/logic/hello"
	"http/internal/svc"
	"http/internal/types"
)

func ListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.HelloListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJsonCtx(r.Context(), w, map[string]interface{}{
				"code": consts_code.ReqParamsErr,
				"msg":  consts_code.ReqParamsErrDoc + err.Error(),
			})
			return
		}

		l := hello.NewListLogic(r.Context(), svcCtx)
		resp, err := l.List(&req)
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
