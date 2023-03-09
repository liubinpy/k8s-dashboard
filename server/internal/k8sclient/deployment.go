package k8sclient

import (
	"context"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	appsv1 "k8s.io/api/apps/v1"
	autoscalingv1 "k8s.io/api/autoscaling/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/json"
	"k8s.io/client-go/kubernetes"
	"time"
)

type Deployment struct {
}

func (d *Deployment) toCells(std []appsv1.Deployment) []DataCell {
	cells := make([]DataCell, len(std))
	for i := range std {
		cells[i] = deploymentCell(std[i])
	}
	return cells
}

func (d *Deployment) fromCells(cells []DataCell) []appsv1.Deployment {
	deployments := make([]appsv1.Deployment, len(cells))
	for i := range cells {
		deployments[i] = appsv1.Deployment(cells[i].(deploymentCell))
	}
	return deployments
}

// GetDeploymentList 获取deployment列表
func (d *Deployment) GetDeploymentList(client *kubernetes.Clientset, filterName, namespace string, limit, page int) (total int, deployments []appsv1.Deployment, err error) {
	deploymentList, err := client.AppsV1().Deployments(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		logx.Errorf("获取deployment列表失败, %s", err.Error())
		return 0, nil, errors.New("获取deployment列表失败")
	}

	// 实例化
	selectableData := &dataSelector{
		GenericDataList: d.toCells(deploymentList.Items),
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
	deployments = d.fromCells(filtered.Sort().Paginate().GenericDataList)

	return total, deployments, nil
}

// GetDeploymentDetail 获取deployment详情
func (d *Deployment) GetDeploymentDetail(client *kubernetes.Clientset, deploymentName, namespace string) (deployment *appsv1.Deployment, err error) {
	deployment, err = client.AppsV1().Deployments(namespace).Get(context.TODO(), deploymentName, metav1.GetOptions{})
	if err != nil {
		logx.Errorf("获取deployment详情失败%s", err.Error())
		return nil, errors.New("获取deployment详情失败")
	}
	return deployment, nil
}

// DeleteDeployment 删除deployment
func (d *Deployment) DeleteDeployment(client *kubernetes.Clientset, deploymentName, namespace string) (err error) {
	err = client.AppsV1().Deployments(namespace).Delete(context.TODO(), deploymentName, metav1.DeleteOptions{})
	if err != nil {
		logx.Errorf("删除deployment失败%s", err.Error())
		return errors.New("删除deployment失败")
	}
	return nil
}

// UpdateDeployment 更新deployment
func (d *Deployment) UpdateDeployment(client *kubernetes.Clientset, namespace, content string) (err error) {
	newDeployment := new(appsv1.Deployment)
	err = json.Unmarshal([]byte(content), newDeployment)
	if err != nil {
		logx.Errorf("序列化失败%s", err.Error())
		return errors.New("更新deployment序列化失败")
	}
	_, err = client.AppsV1().Deployments(namespace).Update(context.TODO(), newDeployment, metav1.UpdateOptions{})
	if err != nil {
		logx.Errorf("更新deployment失败%s", err.Error())
		return errors.New("更新deployment失败")
	}
	return nil
}

// ScaleDeployment 修改deployment副本数
func (d *Deployment) ScaleDeployment(client *kubernetes.Clientset, deploymentName, namespace string, replica int) (err error) {
	scale := &autoscalingv1.Scale{
		Spec: autoscalingv1.ScaleSpec{Replicas: int32(replica)},
	}
	_, err = client.AppsV1().Deployments(namespace).UpdateScale(context.TODO(), deploymentName, scale, metav1.UpdateOptions{})
	if err != nil {
		logx.Errorf("修改deployment副本数失败%s", err.Error())
		return errors.New("修改deployment副本数失败")
	}
	return nil
}

// RestartDeployment 重启deployment
func (d *Deployment) RestartDeployment(client *kubernetes.Clientset, deploymentName, namespace string) (err error) {
	data := fmt.Sprintf(`{"spec": {"template": {"metadata": {"annotations": {"kubectl.kubernetes.io/restartedAt": "%s"}}}}}`, time.Now().Format("20060102150405"))
	_, err = client.AppsV1().Deployments(namespace).Patch(context.TODO(), deploymentName, types.StrategicMergePatchType, []byte(data), metav1.PatchOptions{})
	if err != nil {
		logx.Errorf("重启deployment失败%s", err.Error())
		return errors.New("重启deployment失败")
	}
	return nil
}

// CreateDeployment 创建deployment
func (d *Deployment) CreateDeployment(client *kubernetes.Clientset, namespace, content string) (err error) {
	newDeployment := new(appsv1.Deployment)
	err = json.Unmarshal([]byte(content), newDeployment)
	if err != nil {
		logx.Errorf("序列化失败%s", err.Error())
		return errors.New("创建deployment序列化失败")
	}

	_, err = client.AppsV1().Deployments(namespace).Create(context.TODO(), newDeployment, metav1.CreateOptions{})
	if err != nil {
		logx.Errorf("创建deployment失败%s", err.Error())
		return errors.New("创建deployment失败")
	}
	return nil
}
