package deployments

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"server/internal/logic/deployments"
	"server/internal/svc"
	"server/internal/types"
)

func GetDeploymentDetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetDeploymentDetailRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := deployments.NewGetDeploymentDetailLogic(r.Context(), svcCtx)
		resp, err := l.GetDeploymentDetail(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
