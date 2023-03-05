package k8sclient

import (
	corev1 "k8s.io/api/core/v1"
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
