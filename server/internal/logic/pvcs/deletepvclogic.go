package pvcs

import (
	"context"
	"server/internal/common/response"
	"server/internal/k8sclient"

	"server/internal/svc"
	"server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeletePVCLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeletePVCLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeletePVCLogic {
	return &DeletePVCLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeletePVCLogic) DeletePVC(req *types.DeletePVCRequest) (resp *types.DeletePVCResponse, err error) {
	client, err := l.svcCtx.K8sClient.GetClientByClusterName(req.Cluster)
	if err != nil {
		logx.Errorf("获取集群client失败: %s", err)
		return &types.DeletePVCResponse{Code: response.Failed, Message: err.Error()}, nil
	}

	pvc := k8sclient.PVC{}
	err = pvc.DeletePVC(client, req.Namespace, req.PVCName)
	if err != nil {
		return &types.DeletePVCResponse{Code: response.Failed, Message: err.Error()}, nil
	}
	return &types.DeletePVCResponse{
		Code: response.Success,
	}, nil
}
