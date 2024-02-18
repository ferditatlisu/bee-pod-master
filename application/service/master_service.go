package service

import (
	"bee-pod-master/pkg/client"
	"bee-pod-master/pkg/kubernetes"
	"bee-pod-master/pkg/logger"
	"bee-pod-master/pkg/model/dto"
	"go.uber.org/zap"
)

type masterService struct {
	redis *client.RedisService
	ks    *kubernetes.KubernetesService
}

type MasterService interface {
	Copy(r *dto.CopyRequest) error
	Search(r *dto.SearchRequest) error
	Partition(r *dto.PartitionRequest) error
	Delete(podName string) error
}

func NewMasterService(r *client.RedisService, ks *kubernetes.KubernetesService) MasterService {
	return &masterService{redis: r, ks: ks}
}

func (ms *masterService) Copy(req *dto.CopyRequest) error {
	err := ms.redis.RequestedCopyEvent(req.Key, req.PodName)
	if err != nil {
		logger.Logger().Error("Redis throw error. ", zap.Error(err))
		return err
	}

	err = ms.ks.CopyPodCreate(req)
	if err != nil {
		ms.redis.Delete(req.Key)
	}

	return err
}

func (ms *masterService) Delete(podName string) error {
	err := ms.ks.PodDelete(podName)
	return err
}

func (ms *masterService) Search(req *dto.SearchRequest) error {
	err := ms.redis.RequestedSearchEvent(req.MetadataKey, req.PodName, req.Key)
	if err != nil {
		logger.Logger().Error("Redis throw error. ", zap.Error(err))
		return err
	}

	err = ms.ks.SearchPodCreate(req)
	if err != nil {
		ms.redis.Delete(req.MetadataKey)
	}

	return err
}

func (ms *masterService) Partition(req *dto.PartitionRequest) error {
	err := ms.redis.RequestedCopyEvent(req.Key, req.PodName)
	if err != nil {
		logger.Logger().Error("Redis throw error. ", zap.Error(err))
		return err
	}

	err = ms.ks.PartitionPodCreate(req)
	if err != nil {
		ms.redis.Delete(req.Key)
	}

	return err
}
