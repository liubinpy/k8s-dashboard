package services

import (
	"context"
	"server/internal/common/response"
	"server/internal/k8sclient"

	"server/internal/svc"
	"server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetServiceDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetServiceDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetServiceDetailLogic {
	return &GetServiceDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetServiceDetailLogic) GetServiceDetail(req *types.GetServiceDetailRequest) (resp *types.GetServiceDetailResponse, err error) {
	client, err := l.svcCtx.K8sClient.GetClientByClusterName(req.Cluster)
	if err != nil {
		logx.Errorf("获取集群client失败: %s", err)
		return &types.GetServiceDetailResponse{Code: response.Failed, Message: err.Error()}, nil
	}

	service := k8sclient.Service{}

	serviceDetail, err := service.GetServiceDetail(client, req.Namespace, req.ServiceName)
	if err != nil {
		return &types.GetServiceDetailResponse{Code: response.Failed, Message: err.Error()}, nil
	}
	return &types.GetServiceDetailResponse{
		Code: response.Success,
		Data: serviceDetail,
	}, nil
}
