package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type IngressAdmissionConfig struct {
	metav1.TypeMeta      `json:",inline"`
	AllowHostnameChanges bool `json:"allowHostnameChanges"`
}
