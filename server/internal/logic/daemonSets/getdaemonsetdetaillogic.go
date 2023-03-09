package daemonSets

import (
	"context"
	"server/internal/common/response"
	"server/internal/k8sclient"

	"server/internal/svc"
	"server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDaemonSetDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetDaemonSetDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDaemonSetDetailLogic {
	return &GetDaemonSetDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetDaemonSetDetailLogic) GetDaemonSetDetail(req *types.GetDaemonSetDetailRequest) (resp *types.GetDaemonSetDetailResponse, err error) {
	client, err := l.svcCtx.K8sClient.GetClientByClusterName(req.Cluster)
	if err != nil {
		logx.Errorf("获取集群client失败: %s", err)
		return &types.GetDaemonSetDetailResponse{Code: response.Failed, Message: err.Error()}, nil
	}
	daemonSetDetail, err := k8sclient.DaemonSetClient.GetDaemonSetDetail(client, req.DaemonSetName, req.Namespace)
	if err != nil {
		return &types.GetDaemonSetDetailResponse{Code: response.Failed, Message: err.Error()}, nil
	}
	return &types.GetDaemonSetDetailResponse{
		Code: response.Success,
		Data: daemonSetDetail,
	}, nil
}
