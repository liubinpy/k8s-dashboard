package pvs

import (
	"context"
	"server/internal/common/response"
	"server/internal/k8sclient"

	"server/internal/svc"
	"server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPVDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPVDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPVDetailLogic {
	return &GetPVDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPVDetailLogic) GetPVDetail(req *types.GetPVDetailRequest) (resp *types.GetPVDetailResponse, err error) {
	client, err := l.svcCtx.K8sClient.GetClientByClusterName(req.Cluster)
	if err != nil {
		logx.Errorf("获取集群client失败: %s", err)
		return &types.GetPVDetailResponse{Code: response.Failed, Message: err.Error()}, nil
	}
	pv := k8sclient.PV{}

	detail, err := pv.GetPVDetail(client, req.PVName)
	if err != nil {

		return &types.GetPVDetailResponse{Code: response.Failed, Message: err.Error()}, nil
	}
	return &types.GetPVDetailResponse{
		Code: response.Success,
		Data: detail,
	}, nil
}
