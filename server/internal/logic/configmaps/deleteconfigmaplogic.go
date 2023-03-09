package configmaps

import (
	"context"
	"server/internal/common/response"
	"server/internal/k8sclient"

	"server/internal/svc"
	"server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteConfigmapLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteConfigmapLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteConfigmapLogic {
	return &DeleteConfigmapLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteConfigmapLogic) DeleteConfigmap(req *types.DeleteConfigmapRequest) (resp *types.DeleteConfigmapResponse, err error) {
	client, err := l.svcCtx.K8sClient.GetClientByClusterName(req.Cluster)
	if err != nil {
		logx.Errorf("获取集群client失败: %s", err)
		return &types.DeleteConfigmapResponse{Code: response.Failed, Message: err.Error()}, nil
	}

	configmap := k8sclient.Configmap{}
	err = configmap.DeleteConfigmap(client, req.Namespace, req.ConfigmapName)
	if err != nil {
		return &types.DeleteConfigmapResponse{Code: response.Failed, Message: err.Error()}, nil
	}
	return &types.DeleteConfigmapResponse{
		Code: response.Success,
	}, nil
}
