package ingresses

import (
	"context"
	"server/internal/common/response"
	"server/internal/k8sclient"

	"server/internal/svc"
	"server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetIngressDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetIngressDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetIngressDetailLogic {
	return &GetIngressDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetIngressDetailLogic) GetIngressDetail(req *types.GetIngressDetailRequest) (resp *types.GetIngressDetailResponse, err error) {
	client, err := l.svcCtx.K8sClient.GetClientByClusterName(req.Cluster)
	if err != nil {
		logx.Errorf("获取集群client失败: %s", err)
		return &types.GetIngressDetailResponse{Code: response.Failed, Message: err.Error()}, nil
	}

	detail, err := k8sclient.IngressClient.GetIngressDetail(client, req.IngressName, req.Namespace)
	if err != nil {
		return &types.GetIngressDetailResponse{Code: response.Failed, Message: err.Error()}, nil
	}
	return &types.GetIngressDetailResponse{
		Code: response.Success,
		Data: detail,
	}, nil
}
