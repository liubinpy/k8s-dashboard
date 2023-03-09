package statefulSets

import (
	"context"
	"server/internal/common/response"
	"server/internal/k8sclient"

	"server/internal/svc"
	"server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetStatefulSetDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetStatefulSetDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetStatefulSetDetailLogic {
	return &GetStatefulSetDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetStatefulSetDetailLogic) GetStatefulSetDetail(req *types.GetStatefulSetDetailRequest) (resp *types.GetStatefulSetDetailResponse, err error) {
	client, err := l.svcCtx.K8sClient.GetClientByClusterName(req.Cluster)
	if err != nil {
		logx.Errorf("获取集群client失败: %s", err)
		return &types.GetStatefulSetDetailResponse{Code: response.Failed, Message: err.Error()}, nil
	}
	statefulset := k8sclient.StatefulSet{}
	statefulSet, err := statefulset.GetStatefulSetDetail(client, req.StatefulSetName, req.Namespace)
	if err != nil {
		return &types.GetStatefulSetDetailResponse{Code: response.Failed, Message: err.Error()}, nil
	}
	return &types.GetStatefulSetDetailResponse{
		Code: response.Success,
		Data: statefulSet,
	}, nil
}
