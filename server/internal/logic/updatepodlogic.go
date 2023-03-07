package logic

import (
	"context"
	"server/internal/common/response"
	"server/internal/k8sclient"

	"server/internal/svc"
	"server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdatePodLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdatePodLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatePodLogic {
	return &UpdatePodLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdatePodLogic) UpdatePod(req *types.UpdatePodRequest) (resp *types.UpdatePodResponse, err error) {
	// 获取client
	client, err := l.svcCtx.K8sClient.GetClientByClusterName(req.Cluster)
	if err != nil {
		logx.Errorf("获取集群client失败: %s", err)
		return &types.UpdatePodResponse{Code: response.Failed, Message: err.Error()}, nil
	}
	pod := k8sclient.Pod{}
	err = pod.UpdatePod(client, req.Namespace, req.Content)
	if err != nil {
		return &types.UpdatePodResponse{Code: response.Failed, Message: err.Error()}, nil
	}

	return &types.UpdatePodResponse{
		Code: response.Success,
	}, nil
}
