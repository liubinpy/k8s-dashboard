package secrets

import (
	"context"
	"server/internal/common/response"
	"server/internal/k8sclient"

	"server/internal/svc"
	"server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateSecretLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateSecretLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSecretLogic {
	return &UpdateSecretLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateSecretLogic) UpdateSecret(req *types.UpdateSecretRequest) (resp *types.UpdateSecretResponse, err error) {
	client, err := l.svcCtx.K8sClient.GetClientByClusterName(req.Cluster)
	if err != nil {
		logx.Errorf("获取集群client失败: %s", err)
		return &types.UpdateSecretResponse{Code: response.Failed, Message: err.Error()}, nil
	}

	err = k8sclient.SecretClient.UpdateSecret(client, req.Namespace, req.Content)
	if err != nil {
		return &types.UpdateSecretResponse{Code: response.Failed, Message: err.Error()}, nil
	}
	return &types.UpdateSecretResponse{
		Code:    response.Success,
		Message: "",
	}, nil
}
