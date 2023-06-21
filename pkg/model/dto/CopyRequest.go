package dto

import (
	"bee-pod-master/pkg/util/rand"
)

type CopyRequest struct {
	FromTopic string `json:"fromTopic"`
	FromId    int    `json:"fromId"`
	ToTopic   string `json:"toTopic"`
	ToId      int    `json:"toId"`
	PodName   string `json:"podName"`
	Key       string `json:"key"`
}

func (r *CopyRequest) UpdateCopyRequest() {
	r.PodName = "bee-copy-" + rand.String(8)
}
