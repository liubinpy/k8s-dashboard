package k8sclient

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/logx"
	networkingv1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/json"
	"k8s.io/client-go/kubernetes"
)

var IngressClient ingress

type ingress struct {
}

func (i *ingress) toCells(std []networkingv1.Ingress) []DataCell {
	cells := make([]DataCell, len(std))
	for i := range std {
		cells[i] = ingressCell(std[i])
	}
	return cells
}

func (i *ingress) fromCells(cells []DataCell) []networkingv1.Ingress {
	ingresses := make([]networkingv1.Ingress, len(cells))
	for i := range cells {
		ingresses[i] = networkingv1.Ingress(cells[i].(ingressCell))
	}
	return ingresses
}

// GetIngressesList 获取Ingress列表
func (i *ingress) GetIngressesList(client *kubernetes.Clientset, filterName, namespace string, limit, page int) (total int, ingresses []networkingv1.Ingress, err error) {
	IngressList, err := client.NetworkingV1().Ingresses(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		logx.Errorf("获取Ingress列表失败, %s", err.Error())
		return 0, nil, errors.New("获取Ingress列表失败")
	}

	// 实例化
	selectableData := &dataSelector{
		GenericDataList: i.toCells(IngressList.Items),
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
	ingresses = i.fromCells(filtered.Sort().Paginate().GenericDataList)

	return total, ingresses, nil
}

// GetIngressDetail 获取Ingress详情
func (i *ingress) GetIngressDetail(client *kubernetes.Clientset, IngressName, namespace string) (Ingress *networkingv1.Ingress, err error) {
	Ingress, err = client.NetworkingV1().Ingresses(namespace).Get(context.TODO(), IngressName, metav1.GetOptions{})
	if err != nil {
		logx.Errorf("获取Ingress详情失败%s", err.Error())
		return nil, errors.New("获取Ingress详情失败")
	}
	return Ingress, nil
}

// DeleteIngresses 删除Ingress
func (i *ingress) DeleteIngresses(client *kubernetes.Clientset, IngressName, namespace string) (err error) {
	err = client.NetworkingV1().Ingresses(namespace).Delete(context.TODO(), IngressName, metav1.DeleteOptions{})
	if err != nil {
		logx.Errorf("删除Ingress失败%s", err.Error())
		return errors.New("删除Ingress失败")
	}
	return nil
}

// UpdateIngresses 更新Ingress
func (i *ingress) UpdateIngresses(client *kubernetes.Clientset, namespace, content string) (err error) {
	newIngresses := new(networkingv1.Ingress)
	err = json.Unmarshal([]byte(content), newIngresses)
	if err != nil {
		logx.Errorf("序列化失败%s", err.Error())
		return errors.New("更新Ingress序列化失败")
	}
	_, err = client.NetworkingV1().Ingresses(namespace).Update(context.TODO(), newIngresses, metav1.UpdateOptions{})
	if err != nil {
		logx.Errorf("更新Ingress失败%s", err.Error())
		return errors.New("更新Ingress失败")
	}
	return nil
}

// CreateIngresses 创建Ingress
func (i *ingress) CreateIngresses(client *kubernetes.Clientset, namespace, content string) (err error) {
	newIngresses := new(networkingv1.Ingress)
	err = json.Unmarshal([]byte(content), newIngresses)
	if err != nil {
		logx.Errorf("序列化失败%s", err.Error())
		return errors.New("创建Ingress序列化失败")
	}

	_, err = client.NetworkingV1().Ingresses(namespace).Create(context.TODO(), newIngresses, metav1.CreateOptions{})
	if err != nil {
		logx.Errorf("创建Ingress失败%s", err.Error())
		return errors.New("创建Ingress失败")
	}
	return nil
}
