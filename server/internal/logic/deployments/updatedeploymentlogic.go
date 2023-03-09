package deployments

import (
	"context"
	"server/internal/common/response"
	"server/internal/k8sclient"

	"server/internal/svc"
	"server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateDeploymentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateDeploymentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateDeploymentLogic {
	return &UpdateDeploymentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateDeploymentLogic) UpdateDeployment(req *types.UpdateDeploymentRequest) (resp *types.UpdateDeploymentResponse, err error) {
	// 获取client
	client, err := l.svcCtx.K8sClient.GetClientByClusterName(req.Cluster)
	if err != nil {
		logx.Errorf("获取集群client失败: %s", err)
		return &types.UpdateDeploymentResponse{Code: response.Failed, Message: err.Error()}, nil
	}

	err = k8sclient.DeploymentClient.UpdateDeployment(client, req.Namespace, req.Content)
	if err != nil {
		return &types.UpdateDeploymentResponse{Code: response.Failed, Message: err.Error()}, nil
	}

	return &types.UpdateDeploymentResponse{Code: response.Success}, nil
}
