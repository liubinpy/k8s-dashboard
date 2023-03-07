package k8sclient

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// Node 节点
type Node struct {
}

func (n *Node) toCells(std []corev1.Node) []DataCell {
	cells := make([]DataCell, len(std))
	for i := range std {
		cells[i] = nodeCell(std[i])
	}
	return cells
}

func (n *Node) fromCells(cells []DataCell) []corev1.Node {
	nodes := make([]corev1.Node, len(cells))
	for i := range cells {
		nodes[i] = corev1.Node(cells[i].(nodeCell))
	}
	return nodes
}

// GetNodeList 获取pod列表
func (n *Node) GetNodeList(client *kubernetes.Clientset, filterName string, limit, page int) (total int, nodes []corev1.Node, err error) {
	nodeList, err := client.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		logx.Errorf("获取node列表失败, %s", err.Error())
		return 0, nil, err
	}

	// 实例化
	selectableData := &dataSelector{
		GenericDataList: n.toCells(nodeList.Items),
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
	nodes = n.fromCells(filtered.Sort().Paginate().GenericDataList)

	return total, nodes, nil
}
