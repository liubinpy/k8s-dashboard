package deployments

import (
	"context"
	"server/internal/common/response"
	"server/internal/k8sclient"

	"server/internal/svc"
	"server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ScaleDeploymentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewScaleDeploymentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ScaleDeploymentLogic {
	return &ScaleDeploymentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ScaleDeploymentLogic) ScaleDeployment(req *types.ScaleDeploymentRequest) (resp *types.ScaleDeploymentResponse, err error) {
	// 获取client
	client, err := l.svcCtx.K8sClient.GetClientByClusterName(req.Cluster)
	if err != nil {
		logx.Errorf("获取集群client失败: %s", err)
		return &types.ScaleDeploymentResponse{Code: response.Failed, Message: err.Error()}, nil
	}

	err = k8sclient.DeploymentClient.ScaleDeployment(client, req.DeploymentName, req.Namespace, req.Replica)
	if err != nil {
		return &types.ScaleDeploymentResponse{
			Code:    response.Failed,
			Message: err.Error(),
		}, nil
	}

	return &types.ScaleDeploymentResponse{
		Code: response.Success,
	}, nil
}
