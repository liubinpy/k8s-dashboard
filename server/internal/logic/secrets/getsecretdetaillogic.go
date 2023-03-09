package secrets

import (
	"context"
	"server/internal/common/response"
	"server/internal/k8sclient"

	"server/internal/svc"
	"server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSecretDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSecretDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSecretDetailLogic {
	return &GetSecretDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSecretDetailLogic) GetSecretDetail(req *types.GetSecretDetailRequest) (resp *types.GetSecretDetailResponse, err error) {
	client, err := l.svcCtx.K8sClient.GetClientByClusterName(req.Cluster)
	if err != nil {
		logx.Errorf("获取集群client失败: %s", err)
		return &types.GetSecretDetailResponse{Code: response.Failed, Message: err.Error()}, nil
	}
	secret := k8sclient.Secret{}
	secretDetail, err := secret.GetSecretDetail(client, req.Namespace, req.SecretName)
	if err != nil {
		return &types.GetSecretDetailResponse{Code: response.Failed, Message: err.Error()}, nil
	}
	return &types.GetSecretDetailResponse{
		Code: response.Success,
		Data: secretDetail,
	}, nil
}
