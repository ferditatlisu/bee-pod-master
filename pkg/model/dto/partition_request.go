package dto

import (
	"bee-pod-master/pkg/util/rand"
	"fmt"
	"strings"
)

type PartitionRequest struct {
	IgnoredPartitions []int  `json:"ignoredPartitions"`
	GroupId           string `json:"groupId"`
	Topic             string `json:"topic"`
	ClusterId         int    `json:"clusterId"`
	PodName           string `json:"podName"`
	Key               string `json:"key"`
}

func (r *PartitionRequest) GetPartitions() string {
	p := ""
	for _, partition := range r.IgnoredPartitions {
		p = p + fmt.Sprint(partition) + ","
	}

	p = strings.TrimSuffix(p, ",")
	return p
}

func (r *PartitionRequest) UpdateRequest() {
	r.PodName = "bee-partition-" + rand.String(8)
}
