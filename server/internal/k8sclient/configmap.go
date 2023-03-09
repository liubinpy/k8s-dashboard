package k8sclient

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/logx"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/json"
	"k8s.io/client-go/kubernetes"
)

var ConfigmapClient configmap

// configmap 配置
type configmap struct {
}

func (s *configmap) toCells(std []corev1.ConfigMap) []DataCell {
	cells := make([]DataCell, len(std))
	for i := range std {
		cells[i] = configmapCell(std[i])
	}
	return cells
}

func (s *configmap) fromCells(cells []DataCell) []corev1.ConfigMap {
	configmaps := make([]corev1.ConfigMap, len(cells))
	for i := range cells {
		configmaps[i] = corev1.ConfigMap(cells[i].(configmapCell))
	}
	return configmaps
}

// GetConfigmapList 获取configmap列表
func (s *configmap) GetConfigmapList(client *kubernetes.Clientset, filterName, namespace string, limit, page int) (total int, configmaps []corev1.ConfigMap, err error) {
	configmapList, err := client.CoreV1().ConfigMaps(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		logx.Errorf("获取configmap列表失败, %s", err.Error())
		return 0, nil, errors.New("获取configmap列表失败")
	}

	// 实例化
	selectableData := &dataSelector{
		GenericDataList: s.toCells(configmapList.Items),
		dataSelectorQuery: &DataSelectorQuery{
			FilterQuery: &FilterQuery{Name: filterName},
			PaginateQuery: &PaginateQuery{
				Limit: limit,
				Page:  page,
			},
		},
	}

	// 过滤
	filtered := selectableData.Filter()
	total = len(filtered.GenericDataList)

	// 排序筛选后转换
	configmaps = s.fromCells(filtered.Sort().Paginate().GenericDataList)

	return total, configmaps, nil
}

// CreateConfigmap 创建configmap
func (s *configmap) CreateConfigmap(client *kubernetes.Clientset, namespace, content string) (err error) {
	newConfigmap := &corev1.ConfigMap{}
	err = json.Unmarshal([]byte(content), newConfigmap)
	if err != nil {
		logx.Errorf("序列化configmap失败:%s", err.Error())
		return errors.New("序列化configmap失败")
	}

	_, err = client.CoreV1().ConfigMaps(namespace).Create(context.TODO(), newConfigmap, metav1.CreateOptions{})
	if err != nil {
		logx.Errorf("创建configmap失败:%s", err.Error())
		return errors.New("创建configmap失败")
	}
	return nil
}

// GetConfigmapDetail 获取configmap详情
func (s *configmap) GetConfigmapDetail(client *kubernetes.Clientset, namespace, configmapName string) (configmap *corev1.ConfigMap, err error) {
	configmap, err = client.CoreV1().ConfigMaps(namespace).Get(context.TODO(), configmapName, metav1.GetOptions{})
	if err != nil {
		logx.Errorf("获取configmap失败:%s", err.Error())
		return nil, errors.New("获取configmap失败")
	}

	return configmap, nil
}

// DeleteConfigmap 删除configmap
func (s *configmap) DeleteConfigmap(client *kubernetes.Clientset, namespace, configmapName string) (err error) {
	err = client.CoreV1().ConfigMaps(namespace).Delete(context.TODO(), configmapName, metav1.DeleteOptions{})
	if err != nil {
		logx.Errorf("删除configmap失败:%s", err.Error())
		return errors.New("删除configmap失败")
	}
	return nil
}

// UpdateConfigmap 更新configmap
func (s *configmap) UpdateConfigmap(client *kubernetes.Clientset, namespace, content string) (err error) {
	newConfigmap := &corev1.ConfigMap{}
	err = json.Unmarshal([]byte(content), newConfigmap)
	if err != nil {
		logx.Errorf("序列化configmap失败:%s", err.Error())
		return errors.New("序列化configmap失败")
	}

	_, err = client.CoreV1().ConfigMaps(namespace).Update(context.TODO(), newConfigmap, metav1.UpdateOptions{})
	if err != nil {
		logx.Errorf("更新configmap失败:%s", err.Error())
		return errors.New("更新configmap失败")
	}
	return nil
}
