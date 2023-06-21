package client

import (
	"bee-pod-master/pkg/config"
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

type RedisService struct {
	client *redis.Client
	ctx    context.Context
}

func NewRedisService(rc config.RedisConfig) *RedisService {
	rdb := redis.NewFailoverClient(&redis.FailoverOptions{
		MasterName:    rc.MasterName,
		SentinelAddrs: rc.Host,
		Password:      rc.Password,
		DB:            rc.Database,
		MaxRetries:    5,
	})

	return &RedisService{rdb, context.Background()}
}

func (r *RedisService) Delete(key string) {
	r.client.Del(r.ctx, key)
}

func (r *RedisService) RequestedCopyEvent(key, podName string) error {
	err := r.client.HMSet(r.ctx, key, map[string]interface{}{"status": "Requested", "podName": podName, "createdDate": time.Now().UTC().UnixMilli()}).Err()
	r.SearchExpire(key, time.Hour*24)
	return err
}

func (r *RedisService) RequestedSearchEvent(key, podName, dataKey string) error {
	err := r.client.HMSet(r.ctx, key, map[string]interface{}{"status": "In Queue", "podName": podName, "createdDate": time.Now().UTC().UnixMilli(), "dataKey": dataKey}).Err()
	r.SearchExpire(key, time.Hour*1)
	return err
}

func (r *RedisService) SearchExpire(key string, duration time.Duration) {
	_ = r.client.Expire(r.ctx, key, duration).Err()
}
