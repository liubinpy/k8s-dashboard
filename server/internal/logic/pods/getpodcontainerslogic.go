package pods

import (
	"context"
	"server/internal/common/response"
	"server/internal/k8sclient"

	"server/internal/svc"
	"server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPodContainersLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPodContainersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPodContainersLogic {
	return &GetPodContainersLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPodContainersLogic) GetPodContainers(req *types.GetPodContainersRequest) (resp *types.GetPodContainersResponse, err error) {
	// 获取client
	client, err := l.svcCtx.K8sClient.GetClientByClusterName(req.Cluster)
	if err != nil {
		logx.Errorf("获取集群client失败: %s", err)
		return &types.GetPodContainersResponse{Code: response.Failed, Message: err.Error()}, nil
	}
	pod := k8sclient.Pod{}
	containers, err := pod.GetPodContainers(client, req.PodName, req.Namespace)
	if err != nil {
		return &types.GetPodContainersResponse{Code: response.Failed, Message: err.Error()}, nil
	}

	return &types.GetPodContainersResponse{
		Code:       response.Success,
		Containers: containers,
	}, nil
}
