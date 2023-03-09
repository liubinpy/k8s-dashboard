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

var PVCClient pvc

// pvc 配置
type pvc struct {
}

func (p *pvc) toCells(std []corev1.PersistentVolumeClaim) []DataCell {
	cells := make([]DataCell, len(std))
	for i := range std {
		cells[i] = pvcCell(std[i])
	}
	return cells
}

func (p *pvc) fromCells(cells []DataCell) []corev1.PersistentVolumeClaim {
	pvcs := make([]corev1.PersistentVolumeClaim, len(cells))
	for i := range cells {
		pvcs[i] = corev1.PersistentVolumeClaim(cells[i].(pvcCell))
	}
	return pvcs
}

// GetPVCList 获取secret列表
func (p *pvc) GetPVCList(client *kubernetes.Clientset, filterName, namespace string, limit, page int) (total int, pvcs []corev1.PersistentVolumeClaim, err error) {
	pvcList, err := client.CoreV1().PersistentVolumeClaims(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		logx.Errorf("获取pvc列表失败, %s", err.Error())
		return 0, nil, errors.New("获取pvc列表失败")
	}

	// 实例化
	selectableData := &dataSelector{
		GenericDataList: p.toCells(pvcList.Items),
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
	pvcs = p.fromCells(filtered.Sort().Paginate().GenericDataList)

	return total, pvcs, nil
}

// CreatePVC 创建pvc
func (p *pvc) CreatePVC(client *kubernetes.Clientset, namespace, content string) (err error) {
	newPVC := &corev1.PersistentVolumeClaim{}
	err = json.Unmarshal([]byte(content), newPVC)
	if err != nil {
		logx.Errorf("序列化pvc失败:%s", err.Error())
		return errors.New("序列化pvc失败")
	}

	_, err = client.CoreV1().PersistentVolumeClaims(namespace).Create(context.TODO(), newPVC, metav1.CreateOptions{})
	if err != nil {
		logx.Errorf("创建pvc失败:%s", err.Error())
		return errors.New("创建pvc失败")
	}
	return nil
}

// GetPVCDetail 获取pvc详情
func (p *pvc) GetPVCDetail(client *kubernetes.Clientset, namespace, pvcName string) (pvc *corev1.PersistentVolumeClaim, err error) {
	pvc, err = client.CoreV1().PersistentVolumeClaims(namespace).Get(context.TODO(), pvcName, metav1.GetOptions{})
	if err != nil {
		logx.Errorf("获取pvc失败:%s", err.Error())
		return nil, errors.New("获取pvc失败")
	}

	return pvc, nil
}

// DeletePVC 删除pvc
func (p *pvc) DeletePVC(client *kubernetes.Clientset, namespace, pvcName string) (err error) {
	err = client.CoreV1().PersistentVolumeClaims(namespace).Delete(context.TODO(), pvcName, metav1.DeleteOptions{})
	if err != nil {
		logx.Errorf("删除pvc失败:%s", err.Error())
		return errors.New("删除pvc失败")
	}
	return nil
}

// UpdatePVC 更新pvc
func (p *pvc) UpdatePVC(client *kubernetes.Clientset, namespace, content string) (err error) {
	newPVC := &corev1.PersistentVolumeClaim{}
	err = json.Unmarshal([]byte(content), newPVC)
	if err != nil {
		logx.Errorf("序列化pvc失败:%s", err.Error())
		return errors.New("序列化pvc失败")
	}

	_, err = client.CoreV1().PersistentVolumeClaims(namespace).Update(context.TODO(), newPVC, metav1.UpdateOptions{})
	if err != nil {
		logx.Errorf("更新pvc失败:%s", err.Error())
		return errors.New("更新pvc失败")
	}
	return nil
}
