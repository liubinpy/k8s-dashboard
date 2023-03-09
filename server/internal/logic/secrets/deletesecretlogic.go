package secrets

import (
	"context"
	"server/internal/common/response"
	"server/internal/k8sclient"

	"server/internal/svc"
	"server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteSecretLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteSecretLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteSecretLogic {
	return &DeleteSecretLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteSecretLogic) DeleteSecret(req *types.DeleteSecretRequest) (resp *types.DeleteSecretResponse, err error) {
	client, err := l.svcCtx.K8sClient.GetClientByClusterName(req.Cluster)
	if err != nil {
		logx.Errorf("获取集群client失败: %s", err)
		return &types.DeleteSecretResponse{Code: response.Failed, Message: err.Error()}, nil
	}

	err = k8sclient.SecretClient.DeleteSecret(client, req.Namespace, req.SecretName)
	if err != nil {
		return &types.DeleteSecretResponse{
			Code: response.Failed,
		}, nil
	}
	return &types.DeleteSecretResponse{
		Code: response.Success,
	}, nil
}
