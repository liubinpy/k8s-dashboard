package statefulSets

import (
	"context"
	"server/internal/common/response"
	"server/internal/k8sclient"

	"server/internal/svc"
	"server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteStatefulSetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteStatefulSetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteStatefulSetLogic {
	return &DeleteStatefulSetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteStatefulSetLogic) DeleteStatefulSet(req *types.DeleteStatefulSetRequest) (resp *types.DeleteStatefulSetResponse, err error) {
	client, err := l.svcCtx.K8sClient.GetClientByClusterName(req.Cluster)
	if err != nil {
		logx.Errorf("获取集群client失败: %s", err)
		return &types.DeleteStatefulSetResponse{Code: response.Failed, Message: err.Error()}, nil
	}

	statefulset := k8sclient.StatefulSet{}
	err = statefulset.DeleteStatefulSet(client, req.StatefulSetName, req.Namespace)
	if err != nil {
		return &types.DeleteStatefulSetResponse{Code: response.Failed, Message: err.Error()}, nil
	}
	return &types.DeleteStatefulSetResponse{
		Code: response.Success,
	}, nil
}
