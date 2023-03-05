package logic

import (
	"context"
	"server/internal/k8sclient"
	"server/internal/svc"
	"server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPodListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPodListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPodListLogic {
	return &GetPodListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPodListLogic) GetPodList(req *types.GetPodListRequest) (resp *types.GetPodListResponse, err error) {
	// 获取client
	client, err := l.svcCtx.K8sClient.GetClientByClusterName(req.Cluster)
	if err != nil {
		logx.Errorf("获取集群client失败: %s", err)
		return nil, err
	}
	pod := k8sclient.Pod{}
	total, pods, err := pod.GetPodList(client, req.FilterName, req.Namespace, req.Limit, req.Page)
	if err != nil {
		logx.Errorf("获取pod列表失败%s", err)
		return nil, err
	}
	return &types.GetPodListResponse{
		Code: "10000",
		Data: types.Pods{
			Total:   total,
			PodList: pods,
		},
	}, nil
}
