package services

import (
	"context"
	"server/internal/common/response"
	"server/internal/k8sclient"

	"server/internal/svc"
	"server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteServiceLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteServiceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteServiceLogic {
	return &DeleteServiceLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteServiceLogic) DeleteService(req *types.DeleteServiceRequest) (resp *types.DeleteServiceResponse, err error) {
	client, err := l.svcCtx.K8sClient.GetClientByClusterName(req.Cluster)
	if err != nil {
		logx.Errorf("获取集群client失败: %s", err)
		return &types.DeleteServiceResponse{Code: response.Failed, Message: err.Error()}, nil
	}
	service := k8sclient.Service{}

	err = service.DeleteService(client, req.Namespace, req.ServiceName)
	if err != nil {
		return &types.DeleteServiceResponse{Code: response.Failed, Message: err.Error()}, nil
	}
	return &types.DeleteServiceResponse{
		Code: response.Success,
	}, nil
}
