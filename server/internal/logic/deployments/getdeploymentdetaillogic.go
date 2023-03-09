package deployments

import (
	"context"
	"server/internal/common/response"
	"server/internal/k8sclient"

	"server/internal/svc"
	"server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDeploymentDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetDeploymentDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDeploymentDetailLogic {
	return &GetDeploymentDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetDeploymentDetailLogic) GetDeploymentDetail(req *types.GetDeploymentDetailRequest) (resp *types.GetDeploymentDetailResponse, err error) {
	// 获取client
	client, err := l.svcCtx.K8sClient.GetClientByClusterName(req.Cluster)
	if err != nil {
		logx.Errorf("获取集群client失败: %s", err)
		return &types.GetDeploymentDetailResponse{Code: response.Failed, Message: err.Error()}, nil
	}

	deploymentDetail, err := k8sclient.DeploymentClient.GetDeploymentDetail(client, req.DeploymentName, req.Namespace)
	if err != nil {
		return &types.GetDeploymentDetailResponse{
			Code:    response.Failed,
			Message: err.Error(),
		}, nil
	}
	return &types.GetDeploymentDetailResponse{
		Code: response.Success,
		Data: deploymentDetail,
	}, nil
}
