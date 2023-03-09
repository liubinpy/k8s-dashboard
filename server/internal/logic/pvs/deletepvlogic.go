package pvs

import (
	"context"
	"server/internal/common/response"
	"server/internal/k8sclient"

	"server/internal/svc"
	"server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeletePVLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeletePVLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeletePVLogic {
	return &DeletePVLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeletePVLogic) DeletePV(req *types.DeletePVRequest) (resp *types.DeletePVResponse, err error) {
	client, err := l.svcCtx.K8sClient.GetClientByClusterName(req.Cluster)
	if err != nil {
		logx.Errorf("获取集群client失败: %s", err)
		return &types.DeletePVResponse{Code: response.Failed, Message: err.Error()}, nil
	}

	err = k8sclient.PVClient.DeletePV(client, req.PVName)
	if err != nil {
		return &types.DeletePVResponse{Code: response.Failed, Message: err.Error()}, nil
	}

	return &types.DeletePVResponse{
		Code: response.Success,
	}, nil
}
