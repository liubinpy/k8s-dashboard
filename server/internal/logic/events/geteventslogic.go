package events

import (
	"context"
	"server/internal/common/response"

	"server/internal/svc"
	"server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetEventsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetEventsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetEventsLogic {
	return &GetEventsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetEventsLogic) GetEvents(req *types.GetEventsRequest) (resp *types.GetEventsResponse, err error) {
	total, events, err := l.svcCtx.EventModel.FindAll(req.Cluster, req.Page, req.Limit)
	if err != nil {
		logx.Errorf("查询event列表失败%s", err.Error())
		return &types.GetEventsResponse{
			Code: response.Failed,
		}, nil
	}
	return &types.GetEventsResponse{
		Code:  response.Success,
		Total: total,
		Data:  events,
	}, nil
}
