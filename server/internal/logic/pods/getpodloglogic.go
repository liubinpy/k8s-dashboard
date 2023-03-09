package pods

import (
	"context"
	"server/internal/common/response"
	"server/internal/k8sclient"

	"server/internal/svc"
	"server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPodLogLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPodLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPodLogLogic {
	return &GetPodLogLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPodLogLogic) GetPodLog(req *types.GetPodLogRequest) (resp *types.GetPodLogResponse, err error) {
	// 获取client
	client, err := l.svcCtx.K8sClient.GetClientByClusterName(req.Cluster)
	if err != nil {
		logx.Errorf("获取集群client失败: %s", err)
		return &types.GetPodLogResponse{Code: response.Failed, Message: err.Error()}, nil
	}
	pod := k8sclient.Pod{}
	podLog, err := pod.GetPodLog(client, l.svcCtx.Config.PodLogTailLine, req.ContainerName, req.PodName, req.Namespace)
	if err != nil {
		return &types.GetPodLogResponse{
			Code:    response.Failed,
			Message: err.Error(),
		}, nil
	}

	return &types.GetPodLogResponse{
		Code: response.Success,
		Log:  podLog,
	}, nil
}
