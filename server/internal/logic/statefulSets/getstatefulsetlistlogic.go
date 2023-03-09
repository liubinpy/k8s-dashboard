package statefulSets

import (
	"context"
	"server/internal/common/response"
	"server/internal/k8sclient"

	"server/internal/svc"
	"server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetStatefulSetListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetStatefulSetListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetStatefulSetListLogic {
	return &GetStatefulSetListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetStatefulSetListLogic) GetStatefulSetList(req *types.GetStatefulSetListRequest) (resp *types.GetStatefulSetListResponse, err error) {
	client, err := l.svcCtx.K8sClient.GetClientByClusterName(req.Cluster)
	if err != nil {
		logx.Errorf("获取集群client失败: %s", err)
		return &types.GetStatefulSetListResponse{Code: response.Failed, Message: err.Error()}, nil
	}
	statefulset := k8sclient.StatefulSet{}

	total, statefulsetlist, err := statefulset.GetStatefulSetList(client, req.FilterName, req.Namespace, req.Limit, req.Page)
	if err != nil {
		return &types.GetStatefulSetListResponse{Code: response.Failed, Message: err.Error()}, nil
	}
	return &types.GetStatefulSetListResponse{
		Code: response.Success,
		Data: types.StatefulSets{
			Total:           total,
			StatefulSetList: statefulsetlist,
		},
	}, nil
}
