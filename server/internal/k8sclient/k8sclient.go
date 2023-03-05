package k8sclient

import (
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"server/internal/config"
)

type K8sClient struct {
	// 提供多集群client,key为集群名称
	ClientMap map[string]*kubernetes.Clientset
	Config    config.Config
}

func NewK8sClient(config config.Config) *K8sClient {
	kc := &K8sClient{
		Config: config,
	}
	err := kc.InitClientSet()
	if err != nil {
		logx.Errorf("初始化多集群client失败", err)
		return nil
	}
	return kc
}

func (k *K8sClient) GetClientByClusterName(cluster string) (*kubernetes.Clientset, error) {
	client, ok := k.ClientMap[cluster]
	if !ok {
		return nil, errors.New(fmt.Sprintf("集群：%s不存在，无法获取client", cluster))
	}
	return client, nil
}

func (k *K8sClient) InitClientSet() error {
	cm := make(map[string]*kubernetes.Clientset)
	// 获取配置
	for _, cluster := range k.Config.Clusters {
		// 初始化client
		clientset, err := getClientSet(cluster.File)
		if err != nil {
			logx.Errorf("初始化集群%d的clientset失败", cluster)
			return err
		}
		cm[cluster.Name] = clientset
	}
	k.ClientMap = cm

	return nil
}

func getClientSet(file string) (*kubernetes.Clientset, error) {
	cf, err := clientcmd.BuildConfigFromFlags("", file)
	if err != nil {
		return nil, err
	}
	clientset, err := kubernetes.NewForConfig(cf)
	return clientset, err
}
