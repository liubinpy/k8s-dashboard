package pvs

import (
	"context"
	"server/internal/common/response"
	"server/internal/k8sclient"

	"server/internal/svc"
	"server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPVsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPVsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPVsLogic {
	return &GetPVsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPVsLogic) GetPVs(req *types.GetPVsRequest) (resp *types.GetPVsResponse, err error) {
	client, err := l.svcCtx.K8sClient.GetClientByClusterName(req.Cluster)
	if err != nil {
		logx.Errorf("获取集群client失败: %s", err)
		return &types.GetPVsResponse{Code: response.Failed, Message: err.Error()}, nil
	}
	pv := k8sclient.PV{}
	total, pvs, err := pv.GetPVList(client, req.FilterName, req.Limit, req.Page)
	if err != nil {
		return &types.GetPVsResponse{Code: response.Failed, Message: err.Error()}, nil
	}
	return &types.GetPVsResponse{
		Code: response.Success,
		Data: types.PVs{
			Total:   total,
			PVsList: pvs,
		},
	}, nil
}
