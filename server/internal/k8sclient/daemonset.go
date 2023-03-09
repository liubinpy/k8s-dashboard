package k8sclient

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/logx"
	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/json"
	"k8s.io/client-go/kubernetes"
)

var DaemonSetClient daemonSet

type daemonSet struct {
}

func (d *daemonSet) toCells(std []appsv1.DaemonSet) []DataCell {
	cells := make([]DataCell, len(std))
	for i := range std {
		cells[i] = daemonSetCell(std[i])
	}
	return cells
}

func (d *daemonSet) fromCells(cells []DataCell) []appsv1.DaemonSet {
	daemonSets := make([]appsv1.DaemonSet, len(cells))
	for i := range cells {
		daemonSets[i] = appsv1.DaemonSet(cells[i].(daemonSetCell))
	}
	return daemonSets
}

// GetDaemonSetList 获取daemonSet列表
func (d *daemonSet) GetDaemonSetList(client *kubernetes.Clientset, filterName, namespace string, limit, page int) (total int, daemonSets []appsv1.DaemonSet, err error) {
	daemonSetList, err := client.AppsV1().DaemonSets(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		logx.Errorf("获取daemonSet列表失败, %s", err.Error())
		return 0, nil, errors.New("获取daemonSet列表失败")
	}

	// 实例化
	selectableData := &dataSelector{
		GenericDataList: d.toCells(daemonSetList.Items),
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
	daemonSets = d.fromCells(filtered.Sort().Paginate().GenericDataList)

	return total, daemonSets, nil
}

// GetDaemonSetDetail 获取DaemonSet详情
func (d *daemonSet) GetDaemonSetDetail(client *kubernetes.Clientset, daemonSetName, namespace string) (daemonSet *appsv1.DaemonSet, err error) {
	daemonSet, err = client.AppsV1().DaemonSets(namespace).Get(context.TODO(), daemonSetName, metav1.GetOptions{})
	if err != nil {
		logx.Errorf("获取DaemonSet详情失败%s", err.Error())
		return nil, errors.New("获取DaemonSet详情失败")
	}
	return daemonSet, nil
}

// DeleteDaemonSet 删除DaemonSet
func (d *daemonSet) DeleteDaemonSet(client *kubernetes.Clientset, daemonSetName, namespace string) (err error) {
	err = client.AppsV1().DaemonSets(namespace).Delete(context.TODO(), daemonSetName, metav1.DeleteOptions{})
	if err != nil {
		logx.Errorf("删除DaemonSet失败%s", err.Error())
		return errors.New("删除DaemonSet失败")
	}
	return nil
}

// UpdateDaemonSet 更新DaemonSet
func (d *daemonSet) UpdateDaemonSet(client *kubernetes.Clientset, namespace, content string) (err error) {
	newDaemonSet := new(appsv1.DaemonSet)
	err = json.Unmarshal([]byte(content), newDaemonSet)
	if err != nil {
		logx.Errorf("序列化失败%s", err.Error())
		return errors.New("更新DaemonSet序列化失败")
	}
	_, err = client.AppsV1().DaemonSets(namespace).Update(context.TODO(), newDaemonSet, metav1.UpdateOptions{})
	if err != nil {
		logx.Errorf("更新DaemonSet失败%s", err.Error())
		return errors.New("更新DaemonSet失败")
	}
	return nil
}
