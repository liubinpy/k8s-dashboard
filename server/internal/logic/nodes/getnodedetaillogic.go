package nodes

import (
	"context"
	"server/internal/common/response"
	"server/internal/k8sclient"

	"server/internal/svc"
	"server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetNodeDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetNodeDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetNodeDetailLogic {
	return &GetNodeDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetNodeDetailLogic) GetNodeDetail(req *types.GetNodeDetailRequest) (resp *types.GetNodeDetailResponse, err error) {
	client, err := l.svcCtx.K8sClient.GetClientByClusterName(req.Cluster)
	if err != nil {
		logx.Errorf("获取集群client失败: %s", err)
		return &types.GetNodeDetailResponse{Code: response.Failed, Message: err.Error()}, nil
	}

	node := k8sclient.Node{}
	detail, err := node.GetNodeDetail(client, req.NodeName)
	if err != nil {
		return &types.GetNodeDetailResponse{Code: response.Failed, Message: err.Error()}, nil
	}

	return &types.GetNodeDetailResponse{
		Code: response.Success,
		Data: detail,
	}, nil
}
