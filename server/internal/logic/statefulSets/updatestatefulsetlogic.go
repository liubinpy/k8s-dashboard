package statefulSets

import (
	"context"
	"server/internal/common/response"
	"server/internal/k8sclient"

	"server/internal/svc"
	"server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateStatefulSetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateStatefulSetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateStatefulSetLogic {
	return &UpdateStatefulSetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateStatefulSetLogic) UpdateStatefulSet(req *types.UpdateStatefulSetRequest) (resp *types.UpdateStatefulSetResponse, err error) {
	client, err := l.svcCtx.K8sClient.GetClientByClusterName(req.Cluster)
	if err != nil {
		logx.Errorf("获取集群client失败: %s", err)
		return &types.UpdateStatefulSetResponse{Code: response.Failed, Message: err.Error()}, nil
	}
	statefulset := k8sclient.StatefulSet{}
	err = statefulset.UpdateStatefulSet(client, req.Namespace, req.Content)
	if err != nil {
		return &types.UpdateStatefulSetResponse{Code: response.Failed, Message: err.Error()}, nil
	}
	return &types.UpdateStatefulSetResponse{
		Code: response.Success,
	}, nil
}
