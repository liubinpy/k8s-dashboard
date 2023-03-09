package deployments

import (
	"context"
	"server/internal/common/response"
	"server/internal/k8sclient"

	"server/internal/svc"
	"server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RestartDeploymentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRestartDeploymentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RestartDeploymentLogic {
	return &RestartDeploymentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RestartDeploymentLogic) RestartDeployment(req *types.RestartDeploymentRequest) (resp *types.RestartDeploymentResponse, err error) {
	// 获取client
	client, err := l.svcCtx.K8sClient.GetClientByClusterName(req.Cluster)
	if err != nil {
		logx.Errorf("获取集群client失败: %s", err)
		return &types.RestartDeploymentResponse{Code: response.Failed, Message: err.Error()}, nil
	}

	err = k8sclient.DeploymentClient.RestartDeployment(client, req.DeploymentName, req.Namespace)
	if err != nil {
		return &types.RestartDeploymentResponse{
			Code:    response.Failed,
			Message: err.Error(),
		}, nil
	}
	return &types.RestartDeploymentResponse{
		Code: response.Success,
	}, nil
}
