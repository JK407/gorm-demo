package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"test_demo/api/internal/logic"
	"test_demo/api/internal/svc"
	"test_demo/api/internal/types"
)

func MultiOrderHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.MultiOrderRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewMultiOrderLogic(r.Context(), svcCtx)
		resp, err := l.MultiOrder(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
