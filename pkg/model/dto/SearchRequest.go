package dto

import (
	"bee-pod-master/pkg/util/rand"
)

type SearchRequest struct {
	Topic       string `json:"topic"`
	KafkaId     int    `json:"kafkaId"`
	Value       string `json:"value"`
	PodName     string `json:"podName"`
	Key         string `json:"key"`
	MetadataKey string `json:"metadataKey"`
	StartDate   int64  `json:"startDate"`
	EndDate     int64  `json:"endDate"`
	ValueType   int    `json:"valueType"`
}

func (r *SearchRequest) UpdateSearchRequest() {
	r.PodName = "bee-search-" + rand.String(8)
}
