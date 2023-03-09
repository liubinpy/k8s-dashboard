package services

import (
	"context"
	"server/internal/common/response"
	"server/internal/k8sclient"

	"server/internal/svc"
	"server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateServiceLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateServiceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateServiceLogic {
	return &CreateServiceLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateServiceLogic) CreateService(req *types.CreateServiceRequest) (resp *types.CreateServiceResponse, err error) {
	client, err := l.svcCtx.K8sClient.GetClientByClusterName(req.Cluster)
	if err != nil {
		logx.Errorf("获取集群client失败: %s", err)
		return &types.CreateServiceResponse{Code: response.Failed, Message: err.Error()}, nil

	}

	err = k8sclient.ServiceClient.CreateService(client, req.Namespace, req.Content)
	if err != nil {
		return &types.CreateServiceResponse{Code: response.Failed, Message: err.Error()}, nil
	}
	return &types.CreateServiceResponse{
		Code: response.Success,
	}, nil
}
