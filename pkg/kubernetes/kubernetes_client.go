package kubernetes

import (
	"bee-pod-master/pkg/config"
	"bee-pod-master/pkg/logger"
	"bee-pod-master/pkg/model/dto"
	"context"
	"errors"
	"go.uber.org/zap"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"os"
	"strconv"
)

type KubernetesService struct {
	client    *kubernetes.Clientset
	appConfig *config.ApplicationConfig
}

func NewKubernetesService(applicationConfig *config.ApplicationConfig) *KubernetesService {
	clusterConfig, err := rest.InClusterConfig()
	if err != nil {
		logger.Logger().Error("Kubernetes config was not configured", zap.Error(err))
		panic(err)
	}

	client, err := kubernetes.NewForConfig(clusterConfig)
	if err != nil {
		logger.Logger().Error("Kubernetes config was not configured", zap.Error(err))
	}
	return &KubernetesService{client, applicationConfig}
}

func (ks *KubernetesService) PodDelete(podName string) error {
	err := ks.client.CoreV1().Pods(ks.appConfig.Namespace).Delete(context.TODO(), podName, metav1.DeleteOptions{})
	return err
}

func (ks *KubernetesService) CopyPodCreate(req *dto.CopyRequest) error {
	p := &v1.Pod{
		ObjectMeta: metav1.ObjectMeta{Name: req.PodName, Labels: map[string]string{"app": "copy"}},
		Spec: v1.PodSpec{
			ServiceAccountName: ks.appConfig.ServiceAccountName,
			Containers: []v1.Container{
				{
					Name:  "bee-copy",
					Image: ks.appConfig.CopyImage,
					Ports: []v1.ContainerPort{
						{ContainerPort: 8080},
					},
					Env: []v1.EnvVar{
						{
							Name:  "KEY",
							Value: req.Key,
						},
						{
							Name:  "POD_NAME",
							Value: req.PodName,
						},
						{
							Name:  "TO_TOPIC",
							Value: req.ToTopic,
						},
						{
							Name:  "TO_ID",
							Value: strconv.Itoa(req.ToId),
						},
						{
							Name:  "FROM_TOPIC",
							Value: req.FromTopic,
						},
						{
							Name:  "FROM_ID",
							Value: strconv.Itoa(req.FromId),
						},
						{
							Name:  "KAFKA_CONFIGS",
							Value: os.Getenv("KAFKA_CONFIGS"),
						},
						{
							Name:  "ENVIRONMENT",
							Value: os.Getenv("ENVIRONMENT"),
						},
						{
							Name:  "REDIS_HOST",
							Value: os.Getenv("REDIS_HOST"),
						},
						{
							Name:  "REDIS_PASSWORD",
							Value: os.Getenv("REDIS_PASSWORD"),
						},
						{
							Name:  "REDIS_MASTERNAME",
							Value: os.Getenv("REDIS_MASTERNAME"),
						},
						{
							Name:  "REDIS_DB",
							Value: os.Getenv("REDIS_DB"),
						},
						{
							Name:  "POD_MASTER_URL",
							Value: os.Getenv("POD_MASTER_URL"),
						},
					},
				},
			},
		},
	}

	err := ks.CreatePod(p)
	if err != nil {
		return err
	}

	return nil
}

func (ks *KubernetesService) SearchPodCreate(req *dto.SearchRequest) error {
	p := &v1.Pod{
		ObjectMeta: metav1.ObjectMeta{Name: req.PodName, Labels: map[string]string{"app": "search"}},
		Spec: v1.PodSpec{
			ServiceAccountName: ks.appConfig.ServiceAccountName,
			Containers: []v1.Container{
				{
					Name:  "bee-search",
					Image: ks.appConfig.SearchImage,
					Ports: []v1.ContainerPort{
						{ContainerPort: 8080},
					},
					Env: []v1.EnvVar{
						{
							Name:  "KEY",
							Value: req.Key,
						},
						{
							Name:  "METADATA_KEY",
							Value: req.MetadataKey,
						},
						{
							Name:  "POD_NAME",
							Value: req.PodName,
						},
						{
							Name:  "TOPIC",
							Value: req.Topic,
						},
						{
							Name:  "KAFKA_ID",
							Value: strconv.Itoa(req.KafkaId),
						},
						{
							Name:  "VALUE",
							Value: req.Value,
						},
						{
							Name:  "START_DATE",
							Value: strconv.FormatInt(req.StartDate, 10),
						},
						{
							Name:  "END_DATE",
							Value: strconv.FormatInt(req.EndDate, 10),
						},
						{
							Name:  "VALUE_TYPE",
							Value: strconv.Itoa(req.ValueType),
						},
						{
							Name:  "KAFKA_CONFIGS",
							Value: os.Getenv("KAFKA_CONFIGS"),
						},
						{
							Name:  "ENVIRONMENT",
							Value: os.Getenv("ENVIRONMENT"),
						},
						{
							Name:  "REDIS_HOST",
							Value: os.Getenv("REDIS_HOST"),
						},
						{
							Name:  "REDIS_PASSWORD",
							Value: os.Getenv("REDIS_PASSWORD"),
						},
						{
							Name:  "REDIS_MASTERNAME",
							Value: os.Getenv("REDIS_MASTERNAME"),
						},
						{
							Name:  "REDIS_DB",
							Value: os.Getenv("REDIS_DB"),
						},
						{
							Name:  "POD_MASTER_URL",
							Value: os.Getenv("POD_MASTER_URL"),
						},
					},
				},
			},
		},
	}

	err := ks.CreatePod(p)
	if err != nil {
		return err
	}

	return nil
}

func (ks *KubernetesService) CreatePod(p *v1.Pod) error {
	if ks.client == nil {
		return errors.New("couldn't connect to k8s")
	}

	_, err := ks.client.CoreV1().Pods(ks.appConfig.Namespace).Create(context.TODO(), p, metav1.CreateOptions{})
	if err != nil {
		logger.Logger().Error("Error occured when pod was creating ... ", zap.Error(err))
		return err
	}

	return nil
}
