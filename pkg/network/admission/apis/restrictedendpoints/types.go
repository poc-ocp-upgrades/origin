package restrictedendpoints

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type RestrictedEndpointsAdmissionConfig struct {
	metav1.TypeMeta
	RestrictedCIDRs	[]string
}
