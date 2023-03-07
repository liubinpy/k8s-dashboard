package k8sclient

import (
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	networkingv1 "k8s.io/api/networking/v1"
	"sort"
	"strings"
	"time"
)

// dataSelector 用户封装排序、过滤、分页的数据
type dataSelector struct {
	GenericDataList   []DataCell
	dataSelectorQuery *DataSelectorQuery
}

// DataCell 接口，用于各种资源list的类型转换，转换后可以使用dataSelector的排序、过滤和分页
type DataCell interface {
	GetCreation() time.Time
	GetName() string
}

// DataSelectorQuery 定义过滤和分页的属性，过滤使用name，分页：limit和page
type DataSelectorQuery struct {
	FilterQuery   *FilterQuery
	PaginateQuery *PaginateQuery
}

type FilterQuery struct {
	Name string
}

type PaginateQuery struct {
	Limit int
	Page  int
}

// Len 获取数组的长度
func (d *dataSelector) Len() int {
	return len(d.GenericDataList)
}

// Swap 用于数组中的元素在比较大小后交换位置
func (d *dataSelector) Swap(i, j int) {
	d.GenericDataList[i], d.GenericDataList[j] = d.GenericDataList[j], d.GenericDataList[i]
}

// Less 用于定于数组中元素排序的大小的比较方式
func (d *dataSelector) Less(i, j int) bool {
	a := d.GenericDataList[i].GetCreation()
	b := d.GenericDataList[j].GetCreation()

	return b.Before(a)
}

// Sort 排序
func (d *dataSelector) Sort() *dataSelector {
	sort.Sort(d)
	return d
}

// Filter 过滤
func (d *dataSelector) Filter() *dataSelector {
	if d.dataSelectorQuery.FilterQuery.Name == "" {
		return d
	}
	filteredList := make([]DataCell, 0)
	for _, value := range d.GenericDataList {
		objName := value.GetName()
		if !strings.Contains(objName, d.dataSelectorQuery.FilterQuery.Name) {
			continue
		} else {
			filteredList = append(filteredList, value)
		}
	}
	d.GenericDataList = filteredList
	return d
}

// Paginate 分页
func (d *dataSelector) Paginate() *dataSelector {
	limit := d.dataSelectorQuery.PaginateQuery.Limit
	page := d.dataSelectorQuery.PaginateQuery.Page
	if limit <= 0 || page <= 0 {
		return d
	}

	startIndex := limit * (page - 1)
	endIndex := limit * page
	if len(d.GenericDataList) < endIndex {
		endIndex = len(d.GenericDataList)
	}
	if startIndex > len(d.GenericDataList) {
		d.GenericDataList = []DataCell{}
		return d
	}

	d.GenericDataList = d.GenericDataList[startIndex:endIndex]

	return d
}

// podCell 定义podCell类型
type podCell corev1.Pod

func (p podCell) GetCreation() time.Time {
	return p.CreationTimestamp.Time
}

func (p podCell) GetName() string {
	return p.Name
}

// serviceCell 定义serviceCell类型
type serviceCell corev1.Service

func (s serviceCell) GetCreation() time.Time {
	return s.CreationTimestamp.Time
}

func (s serviceCell) GetName() string {
	return s.Name
}

// nodeCell 定义nodeCell类型
type nodeCell corev1.Node

func (n nodeCell) GetCreation() time.Time {
	return n.CreationTimestamp.Time
}

func (n nodeCell) GetName() string {
	return n.Name
}

// namespaceCell 定义namespaceCell类型
type namespaceCell corev1.Namespace

func (n namespaceCell) GetCreation() time.Time {
	return n.CreationTimestamp.Time
}

func (n namespaceCell) GetName() string {
	return n.Name
}

// pvCell 定义pvCell
type pvCell corev1.PersistentVolume

func (p pvCell) GetCreation() time.Time {
	return p.CreationTimestamp.Time
}

func (p pvCell) GetName() string {
	return p.Name
}

// deploymentCell 定义deploymentCell
type deploymentCell appsv1.Deployment

func (d deploymentCell) GetCreation() time.Time {
	return d.CreationTimestamp.Time
}

func (d deploymentCell) GetName() string {
	return d.Name
}

// daemonSetCell 定义daemonSetCell
type daemonSetCell appsv1.DaemonSet

func (d daemonSetCell) GetCreation() time.Time {
	return d.CreationTimestamp.Time
}

func (d daemonSetCell) GetName() string {
	return d.Name
}

// statefulSetCell 定义statefulSetCell
type statefulSetCell appsv1.StatefulSet

func (d statefulSetCell) GetCreation() time.Time {
	return d.CreationTimestamp.Time
}

func (d statefulSetCell) GetName() string {
	return d.Name
}

// ingressCell 定义ingressCell
type ingressCell networkingv1.Ingress

func (i ingressCell) GetCreation() time.Time {
	return i.CreationTimestamp.Time
}

func (i ingressCell) GetName() string {
	return i.Name
}

// configmapCell 定义configmapCell
type configmapCell corev1.ConfigMap

func (c configmapCell) GetCreation() time.Time {
	return c.CreationTimestamp.Time
}

func (c configmapCell) GetName() string {
	return c.Name
}

// secretCell 定义secretCell
type secretCell corev1.Secret

func (s secretCell) GetCreation() time.Time {
	return s.CreationTimestamp.Time
}

func (s secretCell) GetName() string {
	return s.Name
}

// pvcCell 定义pvcCell
type pvcCell corev1.PersistentVolumeClaim

func (p pvcCell) GetCreation() time.Time {
	return p.CreationTimestamp.Time
}

func (p pvcCell) GetName() string {
	return p.Name
}
