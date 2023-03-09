package configmaps

import (
	"context"
	"server/internal/common/response"
	"server/internal/k8sclient"

	"server/internal/svc"
	"server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetConfigmapsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetConfigmapsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetConfigmapsLogic {
	return &GetConfigmapsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetConfigmapsLogic) GetConfigmaps(req *types.GetConfigmapsRequest) (resp *types.GetConfigmapsResponse, err error) {
	client, err := l.svcCtx.K8sClient.GetClientByClusterName(req.Cluster)
	if err != nil {
		logx.Errorf("获取集群client失败: %s", err)
		return &types.GetConfigmapsResponse{Code: response.Failed, Message: err.Error()}, nil
	}
	configmap := k8sclient.Configmap{}

	total, configmaps, err := configmap.GetConfigmapList(client, req.FilterName, req.Namespace, req.Limit, req.Page)
	if err != nil {
		return &types.GetConfigmapsResponse{Code: response.Failed, Message: err.Error()}, nil
	}
	return &types.GetConfigmapsResponse{
		Code: response.Success,
		Data: types.Configmaps{
			Total:         total,
			ConfigmapList: configmaps,
		},
	}, nil
}
