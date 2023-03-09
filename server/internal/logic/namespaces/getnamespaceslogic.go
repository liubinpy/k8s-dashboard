package namespaces

import (
	"context"
	"server/internal/common/response"
	"server/internal/k8sclient"

	"server/internal/svc"
	"server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetNamespacesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetNamespacesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetNamespacesLogic {
	return &GetNamespacesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetNamespacesLogic) GetNamespaces(req *types.GetNamespacesRequest) (resp *types.GetNamespacesResponse, err error) {
	client, err := l.svcCtx.K8sClient.GetClientByClusterName(req.Cluster)
	if err != nil {
		logx.Errorf("获取集群client失败: %s", err)
		return &types.GetNamespacesResponse{Code: response.Failed, Message: err.Error()}, nil
	}

	namespace := k8sclient.Namespace{}
	total, namespaces, err := namespace.GetNamespaceList(client, req.FilterName, req.Limit, req.Page)
	if err != nil {
		return &types.GetNamespacesResponse{Code: response.Failed, Message: err.Error()}, nil
	}

	return &types.GetNamespacesResponse{
		Code: response.Success,
		Data: types.Namespaces{
			Total:          total,
			NamespacesList: namespaces,
		},
	}, nil
}
