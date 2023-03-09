package deployments

import (
	"context"
	"server/internal/common/response"
	"server/internal/k8sclient"

	"server/internal/svc"
	"server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteDeploymentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteDeploymentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteDeploymentLogic {
	return &DeleteDeploymentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteDeploymentLogic) DeleteDeployment(req *types.DeleteDeploymentRequest) (resp *types.DeleteDeploymentResponse, err error) {
	// 获取client
	client, err := l.svcCtx.K8sClient.GetClientByClusterName(req.Cluster)
	if err != nil {
		logx.Errorf("获取集群client失败: %s", err)
		return &types.DeleteDeploymentResponse{Code: response.Failed, Message: err.Error()}, nil
	}

	err = k8sclient.DeploymentClient.DeleteDeployment(client, req.DeploymentName, req.Namespace)
	if err != nil {
		return &types.DeleteDeploymentResponse{
			Code:    response.Failed,
			Message: err.Error(),
		}, nil
	}
	return &types.DeleteDeploymentResponse{
		Code: response.Success,
	}, nil
}
