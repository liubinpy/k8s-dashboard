package namespaces

import (
	"context"
	"server/internal/common/response"
	"server/internal/k8sclient"

	"server/internal/svc"
	"server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetNamespaceDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetNamespaceDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetNamespaceDetailLogic {
	return &GetNamespaceDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetNamespaceDetailLogic) GetNamespaceDetail(req *types.GetNamespaceDetailRequest) (resp *types.GetNamespaceDetailResponse, err error) {
	client, err := l.svcCtx.K8sClient.GetClientByClusterName(req.Cluster)
	if err != nil {
		logx.Errorf("获取集群client失败: %s", err)
		return &types.GetNamespaceDetailResponse{Code: response.Failed, Message: err.Error()}, nil
	}

	namespace := k8sclient.Namespace{}
	namespaceDetail, err := namespace.GetNamespaceDetail(client, req.NamespaceName)
	if err != nil {
		return &types.GetNamespaceDetailResponse{Code: response.Failed, Message: err.Error()}, nil
	}

	return &types.GetNamespaceDetailResponse{
		Code: response.Success,
		Data: namespaceDetail,
	}, nil
}
