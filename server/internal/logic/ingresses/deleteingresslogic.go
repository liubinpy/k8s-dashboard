package ingresses

import (
	"context"
	"server/internal/common/response"
	"server/internal/k8sclient"

	"server/internal/svc"
	"server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteIngressLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteIngressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteIngressLogic {
	return &DeleteIngressLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteIngressLogic) DeleteIngress(req *types.DeleteIngressRequest) (resp *types.DeleteIngressResponse, err error) {
	client, err := l.svcCtx.K8sClient.GetClientByClusterName(req.Cluster)
	if err != nil {
		logx.Errorf("获取集群client失败: %s", err)
		return &types.DeleteIngressResponse{Code: response.Failed, Message: err.Error()}, nil
	}

	ingress := k8sclient.Ingresses{}
	err = ingress.DeleteIngresses(client, req.IngressName, req.Namespace)
	if err != nil {
		return &types.DeleteIngressResponse{Code: response.Failed, Message: err.Error()}, nil
	}

	return &types.DeleteIngressResponse{
		Code: response.Success,
	}, nil
}
