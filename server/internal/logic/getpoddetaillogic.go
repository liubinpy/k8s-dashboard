package logic

import (
	"context"
	"server/internal/common/response"
	"server/internal/k8sclient"

	"server/internal/svc"
	"server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPodDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPodDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPodDetailLogic {
	return &GetPodDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPodDetailLogic) GetPodDetail(req *types.GetPodDetailRequest) (resp *types.GetPodDetailResponse, err error) {
	// 获取client
	client, err := l.svcCtx.K8sClient.GetClientByClusterName(req.Cluster)
	if err != nil {
		logx.Errorf("获取集群client失败: %s", err)
		return &types.GetPodDetailResponse{Code: response.Failed, Message: err.Error()}, nil
	}
	pod := k8sclient.Pod{}

	podDetail, err := pod.GetPodDetail(client, req.PodName, req.Namespace)
	if err != nil {
		return &types.GetPodDetailResponse{Code: response.Failed, Message: err.Error()}, nil
	}
	return &types.GetPodDetailResponse{
		Code: response.Success,
		Data: podDetail,
	}, nil
}
