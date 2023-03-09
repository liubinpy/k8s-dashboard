package daemonSets

import (
	"context"
	"server/internal/common/response"
	"server/internal/k8sclient"

	"server/internal/svc"
	"server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDaemonSetListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetDaemonSetListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDaemonSetListLogic {
	return &GetDaemonSetListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetDaemonSetListLogic) GetDaemonSetList(req *types.GetDaemonSetListRequest) (resp *types.GetDaemonSetListResponse, err error) {
	client, err := l.svcCtx.K8sClient.GetClientByClusterName(req.Cluster)
	if err != nil {
		logx.Errorf("获取集群client失败: %s", err)
		return &types.GetDaemonSetListResponse{Code: response.Failed, Message: err.Error()}, nil
	}

	daemonset := k8sclient.DaemonSet{}
	total, daemonSets, err := daemonset.GetDaemonSetList(client, req.FilterName, req.Namespace, req.Limit, req.Page)
	if err != nil {
		return &types.GetDaemonSetListResponse{Code: response.Failed, Message: err.Error()}, nil
	}
	return &types.GetDaemonSetListResponse{
		Code: response.Success,
		Data: types.DaemonSets{
			Total:         total,
			DaemonSetList: daemonSets,
		},
	}, nil
}
