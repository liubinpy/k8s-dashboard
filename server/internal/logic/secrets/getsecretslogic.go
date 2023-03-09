package secrets

import (
	"context"
	"server/internal/common/response"
	"server/internal/k8sclient"

	"server/internal/svc"
	"server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSecretsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSecretsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSecretsLogic {
	return &GetSecretsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSecretsLogic) GetSecrets(req *types.GetSecretsRequest) (resp *types.GetSecretsResponse, err error) {
	client, err := l.svcCtx.K8sClient.GetClientByClusterName(req.Cluster)
	if err != nil {
		logx.Errorf("获取集群client失败: %s", err)
		return &types.GetSecretsResponse{Code: response.Failed, Message: err.Error()}, nil
	}

	total, secrets, err := k8sclient.SecretClient.GetSecretList(client, req.FilterName, req.Namespace, req.Limit, req.Page)
	if err != nil {
		return &types.GetSecretsResponse{Code: response.Failed, Message: err.Error()}, nil
	}
	return &types.GetSecretsResponse{
		Code: response.Success,
		Data: types.Secrets{
			Total:      total,
			SecretList: secrets,
		},
	}, nil
}
