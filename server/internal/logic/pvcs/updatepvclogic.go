package pvcs

import (
	"context"
	"server/internal/common/response"
	"server/internal/k8sclient"

	"server/internal/svc"
	"server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdatePVCLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdatePVCLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatePVCLogic {
	return &UpdatePVCLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdatePVCLogic) UpdatePVC(req *types.UpdatePVCRequest) (resp *types.UpdatePVCResponse, err error) {
	client, err := l.svcCtx.K8sClient.GetClientByClusterName(req.Cluster)
	if err != nil {
		logx.Errorf("获取集群client失败: %s", err)
		return &types.UpdatePVCResponse{Code: response.Failed, Message: err.Error()}, nil
	}
	pvc := k8sclient.PVC{}
	err = pvc.UpdatePVC(client, req.Namespace, req.Content)
	if err != nil {
		return &types.UpdatePVCResponse{Code: response.Failed, Message: err.Error()}, nil
	}
	return &types.UpdatePVCResponse{
		Code: response.Success,
	}, nil
}
