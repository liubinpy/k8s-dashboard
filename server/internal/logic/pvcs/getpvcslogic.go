package pvcs

import (
	"context"
	"server/internal/common/response"
	"server/internal/k8sclient"

	"server/internal/svc"
	"server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPVCsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPVCsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPVCsLogic {
	return &GetPVCsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPVCsLogic) GetPVCs(req *types.GetPVCsRequest) (resp *types.GetPVCsResponse, err error) {
	client, err := l.svcCtx.K8sClient.GetClientByClusterName(req.Cluster)
	if err != nil {
		logx.Errorf("获取集群client失败: %s", err)
		return &types.GetPVCsResponse{Code: response.Failed, Message: err.Error()}, nil
	}
	pvc := k8sclient.PVC{}
	total, pvcs, err := pvc.GetPVCList(client, req.FilterName, req.Namespace, req.Limit, req.Page)
	if err != nil {
		return &types.GetPVCsResponse{Code: response.Failed, Message: err.Error()}, nil
	}
	return &types.GetPVCsResponse{
		Code: response.Success,
		Data: types.PVCs{
			Total:    total,
			PVCsList: pvcs,
		},
	}, nil
}
