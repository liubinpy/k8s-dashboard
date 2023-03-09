package daemonSets

import (
	"context"
	"server/internal/common/response"
	"server/internal/k8sclient"

	"server/internal/svc"
	"server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteDaemonSetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteDaemonSetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteDaemonSetLogic {
	return &DeleteDaemonSetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteDaemonSetLogic) DeleteDaemonSet(req *types.DeleteDaemonSetRequest) (resp *types.DeleteDaemonSetResponse, err error) {
	client, err := l.svcCtx.K8sClient.GetClientByClusterName(req.Cluster)
	if err != nil {
		logx.Errorf("获取集群client失败: %s", err)
		return &types.DeleteDaemonSetResponse{Code: response.Failed, Message: err.Error()}, nil
	}

	daemonset := k8sclient.DaemonSet{}
	err = daemonset.DeleteDaemonSet(client, req.DaemonSetName, req.Namespace)
	if err != nil {
		return &types.DeleteDaemonSetResponse{Code: response.Failed, Message: err.Error()}, nil
	}

	return &types.DeleteDaemonSetResponse{
		Code: response.Success,
	}, nil
}
