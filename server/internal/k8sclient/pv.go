package k8sclient

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/logx"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

var PVClient pv

// PV 持久化存储
type pv struct {
}

func (p *pv) toCells(std []corev1.PersistentVolume) []DataCell {
	cells := make([]DataCell, len(std))
	for i := range std {
		cells[i] = pvCell(std[i])
	}
	return cells
}

func (p *pv) fromCells(cells []DataCell) []corev1.PersistentVolume {
	pvs := make([]corev1.PersistentVolume, len(cells))
	for i := range cells {
		pvs[i] = corev1.PersistentVolume(cells[i].(pvCell))
	}
	return pvs
}

// GetPVList 获取GetPVList列表
func (p *pv) GetPVList(client *kubernetes.Clientset, filterName string, limit, page int) (total int, pvs []corev1.PersistentVolume, err error) {
	pvList, err := client.CoreV1().PersistentVolumes().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		logx.Errorf("获取pv列表失败, %s", err.Error())
		return 0, nil, err
	}

	// 实例化
	selectableData := &dataSelector{
		GenericDataList: p.toCells(pvList.Items),
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
	pvs = p.fromCells(filtered.Sort().Paginate().GenericDataList)

	return total, pvs, nil
}

// GetPVDetail 获取pv详情
func (p *pv) GetPVDetail(client *kubernetes.Clientset, pvName string) (pv *corev1.PersistentVolume, err error) {
	pv, err = client.CoreV1().PersistentVolumes().Get(context.TODO(), pvName, metav1.GetOptions{})

	if err != nil {
		logx.Errorf("获取pv详情失败%s", err.Error())
		return nil, errors.New("获取namespace详情失败")
	}
	return pv, nil
}

// DeletePV 删除pv
func (p *pv) DeletePV(client *kubernetes.Clientset, pvName string) (err error) {
	err = client.CoreV1().PersistentVolumes().Delete(context.TODO(), pvName, metav1.DeleteOptions{})
	if err != nil {
		logx.Errorf("删除pv失败:%s", err.Error())
		return errors.New("删除pv失败")
	}
	return nil
}
