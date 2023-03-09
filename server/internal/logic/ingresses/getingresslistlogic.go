package ingresses

import (
	"context"
	"server/internal/common/response"
	"server/internal/k8sclient"

	"server/internal/svc"
	"server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetIngressListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetIngressListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetIngressListLogic {
	return &GetIngressListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetIngressListLogic) GetIngressList(req *types.GetIngressListRequest) (resp *types.GetIngressListResponse, err error) {
	client, err := l.svcCtx.K8sClient.GetClientByClusterName(req.Cluster)
	if err != nil {
		logx.Errorf("获取集群client失败: %s", err)
		return &types.GetIngressListResponse{Code: response.Failed, Message: err.Error()}, nil
	}
	total, ingresses, err := k8sclient.IngressClient.GetIngressesList(client, req.FilterName, req.Namespace, req.Limit, req.Page)
	if err != nil {
		return &types.GetIngressListResponse{Code: response.Failed, Message: err.Error()}, nil
	}

	return &types.GetIngressListResponse{
		Code: response.Success,
		Data: types.Ingresses{
			Total:       total,
			IngressList: ingresses,
		},
	}, nil
}
