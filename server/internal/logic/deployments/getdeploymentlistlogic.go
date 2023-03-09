package deployments

import (
	"context"
	"server/internal/common/response"
	"server/internal/k8sclient"

	"server/internal/svc"
	"server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDeploymentListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetDeploymentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDeploymentListLogic {
	return &GetDeploymentListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetDeploymentListLogic) GetDeploymentList(req *types.GetDeploymentListRequest) (resp *types.GetDeploymentListResponse, err error) {
	// 获取client
	client, err := l.svcCtx.K8sClient.GetClientByClusterName(req.Cluster)
	if err != nil {
		logx.Errorf("获取集群client失败: %s", err)
		return &types.GetDeploymentListResponse{Code: response.Failed, Message: err.Error()}, nil
	}

	deployment := k8sclient.Deployment{}
	total, deployments, err := deployment.GetDeploymentList(client, req.FilterName, req.Namespace, req.Limit, req.Page)
	if err != nil {
		return &types.GetDeploymentListResponse{
			Code:    response.Failed,
			Message: err.Error(),
		}, nil
	}
	return &types.GetDeploymentListResponse{
		Code: response.Success,
		Data: types.Deployments{
			Total:          total,
			DeploymentList: deployments,
		},
	}, nil
}
