package externalipranger

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ExternalIPRangerAdmissionConfig struct {
	metav1.TypeMeta
	ExternalIPNetworkCIDRs []string
	AllowIngressIP         bool
}
