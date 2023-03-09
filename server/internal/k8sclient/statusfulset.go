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

type StatefulSet struct {
}

func (d *StatefulSet) toCells(std []appsv1.StatefulSet) []DataCell {
	cells := make([]DataCell, len(std))
	for i := range std {
		cells[i] = statefulSetCell(std[i])
	}
	return cells
}

func (d *StatefulSet) fromCells(cells []DataCell) []appsv1.StatefulSet {
	statefulSets := make([]appsv1.StatefulSet, len(cells))
	for i := range cells {
		statefulSets[i] = appsv1.StatefulSet(cells[i].(statefulSetCell))
	}
	return statefulSets
}

// GetStatefulSetList 获取statefulSet列表
func (d *StatefulSet) GetStatefulSetList(client *kubernetes.Clientset, filterName, namespace string, limit, page int) (total int, statefulSets []appsv1.StatefulSet, err error) {
	statefulSetList, err := client.AppsV1().StatefulSets(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		logx.Errorf("获取statefulSet列表失败, %s", err.Error())
		return 0, nil, errors.New("获取statefulSet列表失败")
	}

	// 实例化
	selectableData := &dataSelector{
		GenericDataList: d.toCells(statefulSetList.Items),
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
	statefulSets = d.fromCells(filtered.Sort().Paginate().GenericDataList)

	return total, statefulSets, nil
}

// GetStatefulSetDetail 获取StatefulSet详情
func (d *StatefulSet) GetStatefulSetDetail(client *kubernetes.Clientset, statefulSetName, namespace string) (statefulSet *appsv1.StatefulSet, err error) {
	statefulSet, err = client.AppsV1().StatefulSets(namespace).Get(context.TODO(), statefulSetName, metav1.GetOptions{})
	if err != nil {
		logx.Errorf("获取StatefulSet详情失败%s", err.Error())
		return nil, errors.New("获取StatefulSet详情失败")
	}
	return statefulSet, nil
}

// DeleteStatefulSet 删除StatefulSet
func (d *StatefulSet) DeleteStatefulSet(client *kubernetes.Clientset, statefulSetName, namespace string) (err error) {
	err = client.AppsV1().StatefulSets(namespace).Delete(context.TODO(), statefulSetName, metav1.DeleteOptions{})
	if err != nil {
		logx.Errorf("删除StatefulSet失败%s", err.Error())
		return errors.New("删除StatefulSet失败")
	}
	return nil
}

// UpdateStatefulSet 更新StatefulSet
func (d *StatefulSet) UpdateStatefulSet(client *kubernetes.Clientset, namespace, content string) (err error) {
	newStatefulSet := new(appsv1.StatefulSet)
	err = json.Unmarshal([]byte(content), newStatefulSet)
	if err != nil {
		logx.Errorf("序列化失败%s", err.Error())
		return errors.New("更新StatefulSet序列化失败")
	}
	_, err = client.AppsV1().StatefulSets(namespace).Update(context.TODO(), newStatefulSet, metav1.UpdateOptions{})
	if err != nil {
		logx.Errorf("更新StatefulSet失败%s", err.Error())
		return errors.New("更新StatefulSet失败")
	}
	return nil
}
