package k8sclient

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// Namespace 名称空间
type Namespace struct {
}

func (n *Namespace) toCells(std []corev1.Namespace) []DataCell {
	cells := make([]DataCell, len(std))
	for i := range std {
		cells[i] = namespaceCell(std[i])
	}
	return cells
}

func (n *Namespace) fromCells(cells []DataCell) []corev1.Namespace {
	namespaces := make([]corev1.Namespace, len(cells))
	for i := range cells {
		namespaces[i] = corev1.Namespace(cells[i].(namespaceCell))
	}
	return namespaces
}

// GetNamespaceList 获取GetNamespaceList列表
func (n *Namespace) GetNamespaceList(client *kubernetes.Clientset, filterName string, limit, page int) (total int, namespaces []corev1.Namespace, err error) {
	namespaceList, err := client.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		logx.Errorf("获取namespace列表失败, %s", err.Error())
		return 0, nil, err
	}

	// 实例化
	selectableData := &dataSelector{
		GenericDataList: n.toCells(namespaceList.Items),
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
	namespaces = n.fromCells(filtered.Sort().Paginate().GenericDataList)

	return total, namespaces, nil
}
