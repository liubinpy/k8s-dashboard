package configmaps

import (
	"context"
	"server/internal/common/response"
	"server/internal/k8sclient"

	"server/internal/svc"
	"server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateConfigmapLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateConfigmapLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateConfigmapLogic {
	return &UpdateConfigmapLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateConfigmapLogic) UpdateConfigmap(req *types.UpdateConfigmapRequest) (resp *types.UpdateConfigmapResponse, err error) {
	client, err := l.svcCtx.K8sClient.GetClientByClusterName(req.Cluster)
	if err != nil {
		logx.Errorf("获取集群client失败: %s", err)
		return &types.UpdateConfigmapResponse{Code: response.Failed, Message: err.Error()}, nil
	}

	err = k8sclient.ConfigmapClient.UpdateConfigmap(client, req.Namespace, req.Content)
	if err != nil {
		return &types.UpdateConfigmapResponse{Code: response.Failed, Message: err.Error()}, nil
	}
	return &types.UpdateConfigmapResponse{
		Code: response.Success,
	}, nil
}
