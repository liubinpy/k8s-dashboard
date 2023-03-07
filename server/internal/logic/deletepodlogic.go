package logic

import (
	"context"
	"server/internal/common/response"
	"server/internal/k8sclient"

	"server/internal/svc"
	"server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeletePodLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeletePodLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeletePodLogic {
	return &DeletePodLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeletePodLogic) DeletePod(req *types.DeletePodRequest) (resp *types.DeletePodResponse, err error) {
	// 获取client
	client, err := l.svcCtx.K8sClient.GetClientByClusterName(req.Cluster)
	if err != nil {
		logx.Errorf("获取集群client失败: %s", err)
		return &types.DeletePodResponse{Code: response.Failed, Message: err.Error()}, nil
	}
	pod := k8sclient.Pod{}
	err = pod.DeletePod(client, req.PodName, req.Namespace)
	if err != nil {
		return &types.DeletePodResponse{Code: response.Failed, Message: err.Error()}, nil
	}

	return &types.DeletePodResponse{Code: response.Success}, nil

}
