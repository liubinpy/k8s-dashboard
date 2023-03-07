package k8sclient

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/logx"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// Service 服务
type Service struct {
}

func (s *Service) toCells(std []corev1.Service) []DataCell {
	cells := make([]DataCell, len(std))
	for i := range std {
		cells[i] = serviceCell(std[i])
	}
	return cells
}

func (s *Service) fromCells(cells []DataCell) []corev1.Service {
	services := make([]corev1.Service, len(cells))
	for i := range cells {
		services[i] = corev1.Service(cells[i].(serviceCell))
	}
	return services
}

// GetServiceList 获取pod列表
func (s *Service) GetServiceList(client *kubernetes.Clientset, filterName, namespace string, limit, page int) (total int, services []corev1.Service, err error) {
	serviceList, err := client.CoreV1().Services(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		logx.Errorf("获取services列表失败, %s", err.Error())
		return 0, nil, errors.New("获取服务列表失败")
	}

	// 实例化
	selectableData := &dataSelector{
		GenericDataList: s.toCells(serviceList.Items),
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

	// 排序筛选后转换为 []corev1.Pod
	services = s.fromCells(filtered.Sort().Paginate().GenericDataList)

	return total, services, nil
}
