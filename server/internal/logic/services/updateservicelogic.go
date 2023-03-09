package services

import (
	"context"
	"server/internal/common/response"
	"server/internal/k8sclient"

	"server/internal/svc"
	"server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateServiceLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateServiceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateServiceLogic {
	return &UpdateServiceLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateServiceLogic) UpdateService(req *types.UpdateServiceRequest) (resp *types.UpdateServiceResponse, err error) {
	client, err := l.svcCtx.K8sClient.GetClientByClusterName(req.Cluster)
	if err != nil {
		logx.Errorf("获取集群client失败: %s", err)
		return &types.UpdateServiceResponse{Code: response.Failed, Message: err.Error()}, nil
	}
	err = k8sclient.ServiceClient.UpdateService(client, req.Namespace, req.Content)
	if err != nil {
		return &types.UpdateServiceResponse{Code: response.Failed, Message: err.Error()}, nil
	}

	return &types.UpdateServiceResponse{
		Code: response.Failed,
	}, nil
}
