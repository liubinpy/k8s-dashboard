package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"server/internal/logic"
	"server/internal/svc"
	"server/internal/types"
)

func GetPodLogHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetPodLogRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewGetPodLogLogic(r.Context(), svcCtx)
		resp, err := l.GetPodLog(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
