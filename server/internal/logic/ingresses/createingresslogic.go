package ingresses

import (
	"context"
	"server/internal/common/response"
	"server/internal/k8sclient"

	"server/internal/svc"
	"server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateIngressLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateIngressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateIngressLogic {
	return &CreateIngressLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateIngressLogic) CreateIngress(req *types.CreateIngressRequest) (resp *types.CreateIngressResponse, err error) {
	client, err := l.svcCtx.K8sClient.GetClientByClusterName(req.Cluster)
	if err != nil {
		logx.Errorf("获取集群client失败: %s", err)
		return &types.CreateIngressResponse{Code: response.Failed, Message: err.Error()}, nil
	}

	ingress := k8sclient.Ingresses{}
	err = ingress.CreateIngresses(client, req.Namespace, req.Content)
	if err != nil {
		if err != nil {
			return &types.CreateIngressResponse{Code: response.Failed, Message: err.Error()}, nil
		}
	}

	return &types.CreateIngressResponse{
		Code: response.Success,
	}, nil
}
