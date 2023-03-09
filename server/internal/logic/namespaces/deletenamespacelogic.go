package namespaces

import (
	"context"
	"server/internal/common/response"
	"server/internal/k8sclient"

	"server/internal/svc"
	"server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteNamespaceLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteNamespaceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteNamespaceLogic {
	return &DeleteNamespaceLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteNamespaceLogic) DeleteNamespace(req *types.DeleteNamespaceRequest) (resp *types.DeleteNamespaceResponse, err error) {
	client, err := l.svcCtx.K8sClient.GetClientByClusterName(req.Cluster)
	if err != nil {
		logx.Errorf("获取集群client失败: %s", err)
		return &types.DeleteNamespaceResponse{Code: response.Failed, Message: err.Error()}, nil
	}

	namespace := k8sclient.Namespace{}
	err = namespace.DeleteNamespace(client, req.NamespaceName)
	if err != nil {
		return &types.DeleteNamespaceResponse{Code: response.Failed, Message: err.Error()}, nil
	}

	return &types.DeleteNamespaceResponse{
		Code: response.Success,
	}, nil
}
