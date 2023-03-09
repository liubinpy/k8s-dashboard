package nodes

import (
	"context"
	"server/internal/common/response"
	"server/internal/k8sclient"

	"server/internal/svc"
	"server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetNodesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetNodesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetNodesLogic {
	return &GetNodesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetNodesLogic) GetNodes(req *types.GetNodesRequest) (resp *types.GetNodesResponse, err error) {
	client, err := l.svcCtx.K8sClient.GetClientByClusterName(req.Cluster)
	if err != nil {
		logx.Errorf("获取集群client失败: %s", err)
		return &types.GetNodesResponse{Code: response.Failed, Message: err.Error()}, nil
	}

	total, nodes, err := k8sclient.NodeClient.GetNodeList(client, req.FilterName, req.Limit, req.Page)
	if err != nil {
		return &types.GetNodesResponse{Code: response.Failed, Message: err.Error()}, nil
	}
	return &types.GetNodesResponse{
		Code: response.Success,
		Data: types.Nodes{
			Total:     total,
			NodesList: nodes,
		},
	}, nil
}
