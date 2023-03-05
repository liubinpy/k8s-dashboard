package svc

import (
	"server/internal/config"
	"server/internal/k8sclient"
)

type ServiceContext struct {
	Config    config.Config
	K8sClient *k8sclient.K8sClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		K8sClient: k8sclient.NewK8sClient(c),
	}
}
