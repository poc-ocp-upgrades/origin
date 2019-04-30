package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ClusterResourceOverrideConfig struct {
	metav1.TypeMeta			`json:",inline"`
	LimitCPUToMemoryPercent		int64	`json:"limitCPUToMemoryPercent"`
	CPURequestToLimitPercent	int64	`json:"cpuRequestToLimitPercent"`
	MemoryRequestToLimitPercent	int64	`json:"memoryRequestToLimitPercent"`
}
