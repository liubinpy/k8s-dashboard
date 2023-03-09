package configmaps

import (
	"context"
	"server/internal/common/response"
	"server/internal/k8sclient"

	"server/internal/svc"
	"server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetConfigmapDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetConfigmapDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetConfigmapDetailLogic {
	return &GetConfigmapDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetConfigmapDetailLogic) GetConfigmapDetail(req *types.GetConfigmapDetailRequest) (resp *types.GetConfigmapDetailResponse, err error) {
	client, err := l.svcCtx.K8sClient.GetClientByClusterName(req.Cluster)
	if err != nil {
		logx.Errorf("获取集群client失败: %s", err)
		return &types.GetConfigmapDetailResponse{Code: response.Failed, Message: err.Error()}, nil
	}

	configmapDetail, err := k8sclient.ConfigmapClient.GetConfigmapDetail(client, req.Namespace, req.ConfigmapName)
	if err != nil {
		return &types.GetConfigmapDetailResponse{Code: response.Failed, Message: err.Error()}, nil
	}
	return &types.GetConfigmapDetailResponse{
		Code: response.Success,
		Data: configmapDetail,
	}, nil
}
