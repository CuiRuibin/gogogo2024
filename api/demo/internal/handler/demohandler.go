package handler

import (
	"net/http"

	"github.com/CuiRuibin/gogogo2024/api/demo/internal/logic"
	"github.com/CuiRuibin/gogogo2024/api/demo/internal/svc"
	"github.com/CuiRuibin/gogogo2024/api/demo/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func DemoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewDemoLogic(r.Context(), svcCtx)
		resp, err := l.Demo(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
