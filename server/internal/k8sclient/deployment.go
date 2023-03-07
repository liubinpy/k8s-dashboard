package k8sclient

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/logx"
	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type Deployment struct {
}

func (p *Deployment) toCells(std []appsv1.Deployment) []DataCell {
	cells := make([]DataCell, len(std))
	for i := range std {
		cells[i] = deploymentCell(std[i])
	}
	return cells
}

func (p *Deployment) fromCells(cells []DataCell) []appsv1.Deployment {
	deployments := make([]appsv1.Deployment, len(cells))
	for i := range cells {
		deployments[i] = appsv1.Deployment(cells[i].(deploymentCell))
	}
	return deployments
}

// GetDeploymentList 获取deployment列表
func (p *Deployment) GetDeploymentList(client *kubernetes.Clientset, filterName, namespace string, limit, page int) (total int, deployments []appsv1.Deployment, err error) {
	deploymentList, err := client.AppsV1().Deployments(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		logx.Errorf("获取deployment列表失败, %s", err.Error())
		return 0, nil, errors.New("获取deployment列表失败")
	}

	// 实例化
	selectableData := &dataSelector{
		GenericDataList: p.toCells(deploymentList.Items),
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
	deployments = p.fromCells(filtered.Sort().Paginate().GenericDataList)

	return total, deployments, nil
}
