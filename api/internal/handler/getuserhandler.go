package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"go-zero-demo/api/internal/logic"
	"go-zero-demo/api/internal/svc"
	"go-zero-demo/api/internal/types"
)

func getUserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewGetUserLogic(r.Context(), svcCtx)
		resp, err := l.GetUser(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
