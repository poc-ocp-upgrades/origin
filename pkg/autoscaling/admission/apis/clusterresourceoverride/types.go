package clusterresourceoverride

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ClusterResourceOverrideConfig struct {
	metav1.TypeMeta
	LimitCPUToMemoryPercent     int64
	CPURequestToLimitPercent    int64
	MemoryRequestToLimitPercent int64
}
