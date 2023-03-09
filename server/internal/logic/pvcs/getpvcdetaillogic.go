package pvcs

import (
	"context"
	"server/internal/common/response"
	"server/internal/k8sclient"

	"server/internal/svc"
	"server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPVCDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPVCDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPVCDetailLogic {
	return &GetPVCDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPVCDetailLogic) GetPVCDetail(req *types.GetPVCDetailRequest) (resp *types.GetPVCDetailResponse, err error) {
	client, err := l.svcCtx.K8sClient.GetClientByClusterName(req.Cluster)
	if err != nil {
		logx.Errorf("获取集群client失败: %s", err)
		return &types.GetPVCDetailResponse{Code: response.Failed, Message: err.Error()}, nil
	}

	pvcDetail, err := k8sclient.PVCClient.GetPVCDetail(client, req.Namespace, req.PVCName)
	if err != nil {
		return &types.GetPVCDetailResponse{Code: response.Failed, Message: err.Error()}, nil
	}
	return &types.GetPVCDetailResponse{
		Code: response.Success,
		Data: pvcDetail,
	}, nil
}
