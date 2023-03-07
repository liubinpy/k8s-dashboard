package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"server/internal/logic"
	"server/internal/svc"
	"server/internal/types"
)

func GetPodContainersHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetPodContainersRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewGetPodContainersLogic(r.Context(), svcCtx)
		resp, err := l.GetPodContainers(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
