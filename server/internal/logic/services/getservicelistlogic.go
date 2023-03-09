package services

import (
	"context"
	"server/internal/common/response"
	"server/internal/k8sclient"

	"server/internal/svc"
	"server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetServiceListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetServiceListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetServiceListLogic {
	return &GetServiceListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetServiceListLogic) GetServiceList(req *types.GetServiceListRequest) (resp *types.GetServiceListResponse, err error) {
	// 获取client
	client, err := l.svcCtx.K8sClient.GetClientByClusterName(req.Cluster)
	if err != nil {
		logx.Errorf("获取集群client失败: %s", err)
		return &types.GetServiceListResponse{Code: response.Failed, Message: err.Error()}, nil
	}
	service := k8sclient.Service{}
	total, services, err := service.GetServiceList(client, req.FilterName, req.Namespace, req.Limit, req.Page)
	if err != nil {
		return &types.GetServiceListResponse{
			Code:    response.Failed,
			Message: err.Error(),
		}, nil
	}
	return &types.GetServiceListResponse{
		Code: response.Success,
		Data: types.Services{
			Total:       total,
			ServiceList: services,
		},
	}, nil
	return
}
