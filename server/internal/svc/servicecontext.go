package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"server/internal/config"
	"server/internal/k8sclient"
	"server/internal/model"
)

type ServiceContext struct {
	Config     config.Config
	K8sClient  *k8sclient.K8sClient
	EventModel model.EventModel
}

func NewServiceContext(c config.Config) *ServiceContext {

	sqlConn := sqlx.NewMysql(c.Mysql.Datasource)
	eventModel := model.NewEventModel(sqlConn)
	kClient := k8sclient.NewK8sClient(c)
	// 开启event同步任务
	StartEventSync(kClient, eventModel)

	return &ServiceContext{
		Config:     c,
		K8sClient:  kClient,
		EventModel: eventModel,
	}
}

func StartEventSync(kClient *k8sclient.K8sClient, eventModel model.EventModel) {
	for cluster, client := range kClient.ClientMap {
		go k8sclient.SyncEvent(client, eventModel, cluster)
	}
}
