package deployments

import (
	"context"
	"server/internal/common/response"
	"server/internal/k8sclient"

	"server/internal/svc"
	"server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateDeploymentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateDeploymentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateDeploymentLogic {
	return &CreateDeploymentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateDeploymentLogic) CreateDeployment(req *types.CreateDeploymentRequest) (resp *types.CreateDeploymentResponse, err error) {
	// 获取client
	client, err := l.svcCtx.K8sClient.GetClientByClusterName(req.Cluster)
	if err != nil {
		logx.Errorf("获取集群client失败: %s", err)
		return &types.CreateDeploymentResponse{Code: response.Failed, Message: err.Error()}, nil
	}

	// 创建deployment
	deployment := k8sclient.Deployment{}
	err = deployment.CreateDeployment(client, req.Namespace, req.Content)
	if err != nil {
		return &types.CreateDeploymentResponse{Code: response.Failed, Message: err.Error()}, nil
	}
	return &types.CreateDeploymentResponse{Code: response.Success}, nil
}
