package daemonSets

import (
	"context"
	"server/internal/common/response"
	"server/internal/k8sclient"

	"server/internal/svc"
	"server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateDaemonSetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateDaemonSetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateDaemonSetLogic {
	return &UpdateDaemonSetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateDaemonSetLogic) UpdateDaemonSet(req *types.UpdateDaemonSetRequest) (resp *types.UpdateDaemonSetResponse, err error) {
	client, err := l.svcCtx.K8sClient.GetClientByClusterName(req.Cluster)
	if err != nil {
		logx.Errorf("获取集群client失败: %s", err)
		return &types.UpdateDaemonSetResponse{Code: response.Failed, Message: err.Error()}, nil
	}
	daemonset := k8sclient.DaemonSet{}

	err = daemonset.UpdateDaemonSet(client, req.Namespace, req.Content)
	if err != nil {
		return &types.UpdateDaemonSetResponse{Code: response.Failed, Message: err.Error()}, nil
	}
	return &types.UpdateDaemonSetResponse{
		Code: response.Success,
	}, nil
}
