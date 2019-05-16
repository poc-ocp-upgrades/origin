package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ExternalIPRangerAdmissionConfig struct {
	metav1.TypeMeta        `json:",inline"`
	ExternalIPNetworkCIDRs []string `json:"externalIPNetworkCIDRs"`
	AllowIngressIP         bool     `json:"allowIngressIP"`
}
