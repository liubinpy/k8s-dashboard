package k8sclient

import (
	"bytes"
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/logx"
	"io"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/json"
	"k8s.io/client-go/kubernetes"
)

type Pod struct {
}

func (p *Pod) toCells(std []corev1.Pod) []DataCell {
	cells := make([]DataCell, len(std))
	for i := range std {
		cells[i] = podCell(std[i])
	}
	return cells
}

func (p *Pod) fromCells(cells []DataCell) []corev1.Pod {
	pods := make([]corev1.Pod, len(cells))
	for i := range cells {
		pods[i] = corev1.Pod(cells[i].(podCell))
	}
	return pods
}

// GetPodList 获取pod列表
func (p *Pod) GetPodList(client *kubernetes.Clientset, filterName, namespace string, limit, page int) (total int, pods []corev1.Pod, err error) {
	podList, err := client.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		logx.Errorf("获取pods列表失败, %s", err.Error())
		return 0, nil, errors.New("获取pods列表失败")
	}

	// 实例化
	selectableData := &dataSelector{
		GenericDataList: p.toCells(podList.Items),
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

	// 排序筛选后转换为 []corev1.Pod
	pods = p.fromCells(filtered.Sort().Paginate().GenericDataList)

	return total, pods, nil
}

// GetPodDetail 获取pod详情
func (p *Pod) GetPodDetail(client *kubernetes.Clientset, podName, namespace string) (pod *corev1.Pod, err error) {
	pod, err = client.CoreV1().Pods(namespace).Get(context.TODO(), podName, metav1.GetOptions{})
	if err != nil {
		logx.Errorf("获取pod %s详情失败: %s", podName, err.Error())
		return nil, errors.New("获取pod详情失败")
	}
	return pod, err
}

// DeletePod 删除pod
func (p *Pod) DeletePod(client *kubernetes.Clientset, podName, namespace string) (err error) {
	err = client.CoreV1().Pods(namespace).Delete(context.TODO(), podName, metav1.DeleteOptions{})
	if err != nil {
		logx.Errorf("删除pod %s详情失败: %s", podName, err.Error())
		return errors.New("删除pod失败")
	}
	return nil
}

// UpdatePod 更新pod, content为资源的json
func (p *Pod) UpdatePod(client *kubernetes.Clientset, namespace string, content string) (err error) {
	var pod = &corev1.Pod{}
	err = json.Unmarshal([]byte(content), pod)
	if err != nil {
		logx.Errorf("反序列化pod失败", err)
		return errors.New("系统内部错误")
	}
	_, err = client.CoreV1().Pods(namespace).Update(context.TODO(), pod, metav1.UpdateOptions{})
	if err != nil {
		logx.Errorf("更新pod失败: %s", err.Error())
		return errors.New("更新pod失败")
	}
	return nil
}

// GetPodContainers 获取pods中的container
func (p *Pod) GetPodContainers(client *kubernetes.Clientset, podName, namespace string) (containers []string, err error) {
	podDetail, err := p.GetPodDetail(client, podName, namespace)
	if err != nil {
		return nil, err
	}
	for _, container := range podDetail.Spec.Containers {
		containers = append(containers, container.Name)
	}
	return containers, nil
}

// GetPodLog 获取pod中的日志
func (p *Pod) GetPodLog(client *kubernetes.Clientset, podTailLine int, containerName, podName, namespace string) (log string, err error) {
	lineLimit := int64(podTailLine)
	option := &corev1.PodLogOptions{
		Container: containerName,
		TailLines: &lineLimit,
	}
	req := client.CoreV1().Pods(namespace).GetLogs(podName, option)
	podLogs, err := req.Stream(context.TODO())
	if err != nil {
		logx.Errorf("获取日志失败%s", err.Error())
		return "", errors.New("获取pod日志失败")
	}
	defer podLogs.Close()

	buf := new(bytes.Buffer)
	_, err = io.Copy(buf, podLogs)
	if err != nil {
		logx.Errorf("io copy error: %s", err)
		return "", errors.New("服务器异常")
	}
	return buf.String(), nil

}

// PV 持久化存储
type PV struct {
}

func (p *PV) toCells(std []corev1.PersistentVolume) []DataCell {
	cells := make([]DataCell, len(std))
	for i := range std {
		cells[i] = pvCell(std[i])
	}
	return cells
}

func (p *PV) fromCells(cells []DataCell) []corev1.PersistentVolume {
	pvs := make([]corev1.PersistentVolume, len(cells))
	for i := range cells {
		pvs[i] = corev1.PersistentVolume(cells[i].(pvCell))
	}
	return pvs
}

// GetPVList 获取GetPVList列表
func (p *PV) GetPVList(client *kubernetes.Clientset, filterName string, limit, page int) (total int, pvs []corev1.PersistentVolume, err error) {
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
