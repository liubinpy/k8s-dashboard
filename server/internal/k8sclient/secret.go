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

var SecretClient secret

// secret 配置
type secret struct {
}

func (s *secret) toCells(std []corev1.Secret) []DataCell {
	cells := make([]DataCell, len(std))
	for i := range std {
		cells[i] = secretCell(std[i])
	}
	return cells
}

func (s *secret) fromCells(cells []DataCell) []corev1.Secret {
	secrets := make([]corev1.Secret, len(cells))
	for i := range cells {
		secrets[i] = corev1.Secret(cells[i].(secretCell))
	}
	return secrets
}

// GetSecretList 获取secret列表
func (s *secret) GetSecretList(client *kubernetes.Clientset, filterName, namespace string, limit, page int) (total int, secrets []corev1.Secret, err error) {
	secretList, err := client.CoreV1().Secrets(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		logx.Errorf("获取secret列表失败, %s", err.Error())
		return 0, nil, errors.New("获取secret列表失败")
	}

	// 实例化
	selectableData := &dataSelector{
		GenericDataList: s.toCells(secretList.Items),
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
	secrets = s.fromCells(filtered.Sort().Paginate().GenericDataList)

	return total, secrets, nil
}

// CreateSecret 创建secret
func (s *secret) CreateSecret(client *kubernetes.Clientset, namespace, content string) (err error) {
	newSecret := &corev1.Secret{}
	err = json.Unmarshal([]byte(content), newSecret)
	if err != nil {
		logx.Errorf("序列化secret失败:%s", err.Error())
		return errors.New("序列化secret失败")
	}

	_, err = client.CoreV1().Secrets(namespace).Create(context.TODO(), newSecret, metav1.CreateOptions{})
	if err != nil {
		logx.Errorf("创建secret失败:%s", err.Error())
		return errors.New("创建secret失败")
	}
	return nil
}

// GetSecretDetail 获取secret详情
func (s *secret) GetSecretDetail(client *kubernetes.Clientset, namespace, secretName string) (secret *corev1.Secret, err error) {
	secret, err = client.CoreV1().Secrets(namespace).Get(context.TODO(), secretName, metav1.GetOptions{})
	if err != nil {
		logx.Errorf("获取secret失败:%s", err.Error())
		return nil, errors.New("获取secret失败")
	}

	return secret, nil
}

// DeleteSecret 删除secret
func (s *secret) DeleteSecret(client *kubernetes.Clientset, namespace, secretName string) (err error) {
	err = client.CoreV1().Secrets(namespace).Delete(context.TODO(), secretName, metav1.DeleteOptions{})
	if err != nil {
		logx.Errorf("删除secret失败:%s", err.Error())
		return errors.New("删除secret失败")
	}
	return nil
}

// UpdateSecret 更新secret
func (s *secret) UpdateSecret(client *kubernetes.Clientset, namespace, content string) (err error) {
	newSecret := &corev1.Secret{}
	err = json.Unmarshal([]byte(content), newSecret)
	if err != nil {
		logx.Errorf("序列化secret失败:%s", err.Error())
		return errors.New("序列化secret失败")
	}

	_, err = client.CoreV1().Secrets(namespace).Update(context.TODO(), newSecret, metav1.UpdateOptions{})
	if err != nil {
		logx.Errorf("更新secret失败:%s", err.Error())
		return errors.New("更新secret失败")
	}
	return nil
}
