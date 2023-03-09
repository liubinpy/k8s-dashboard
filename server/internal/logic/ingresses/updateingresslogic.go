package ingresses

import (
	"context"
	"server/internal/common/response"
	"server/internal/k8sclient"

	"server/internal/svc"
	"server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateIngressLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateIngressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateIngressLogic {
	return &UpdateIngressLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateIngressLogic) UpdateIngress(req *types.UpdateIngressRequest) (resp *types.UpdateIngressResponse, err error) {
	client, err := l.svcCtx.K8sClient.GetClientByClusterName(req.Cluster)
	if err != nil {
		logx.Errorf("获取集群client失败: %s", err)
		return &types.UpdateIngressResponse{Code: response.Failed, Message: err.Error()}, nil
	}
	ingress := k8sclient.Ingresses{}
	err = ingress.UpdateIngresses(client, req.Namespace, req.Content)
	if err != nil {
		return &types.UpdateIngressResponse{Code: response.Failed, Message: err.Error()}, nil
	}
	return &types.UpdateIngressResponse{
		Code: response.Failed,
	}, nil
}
