package logic

import (
	"context"

	"server/internal/svc"
	"server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetServiceistLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetServiceistLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetServiceistLogic {
	return &GetServiceistLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetServiceistLogic) GetServiceist(req *types.GetServiceistRequest) (resp *types.GetServiceistResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
