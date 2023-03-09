package k8sclient

import (
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"server/internal/model"
	"time"
)

func SyncEvent(clientset *kubernetes.Clientset, EventModel model.EventModel, cluster string) {
	factory := informers.NewSharedInformerFactoryWithOptions(clientset, 1000)
	informer := factory.Core().V1().Events().Informer()
	_, err := informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			event := obj.(*corev1.Event)
			// 先去判断有没有 name, kind, namespace, reason, cluster string, eventTime *time.Time
			err := EventModel.HasEvent(event.InvolvedObject.Name, event.InvolvedObject.Kind, event.InvolvedObject.Namespace,
				event.Reason, cluster, &event.CreationTimestamp.Time)
			if err != nil && err == sqlc.ErrNotFound {
				err := EventModel.InsertEvent(&model.Event{
					Kind:       event.InvolvedObject.Kind,
					Namespace:  event.InvolvedObject.Namespace,
					Rtype:      event.Type,
					Reason:     event.Reason,
					Message:    event.Message,
					EventTime:  event.CreationTimestamp.Time,
					Cluster:    cluster,
					CreateTime: time.Now(),
					Name:       event.InvolvedObject.Name,
				})
				if err != nil {
					logx.Errorf("写入event异常%s", err.Error())
				}
			}
		},
	})

	if err != nil {
		panic(err)
	}
	// 启动factor
	stopChan := make(chan struct{})
	factory.Start(stopChan)

	factory.WaitForCacheSync(stopChan)
	logx.Info("同步event任务启动成功.....")
	<-stopChan

}
