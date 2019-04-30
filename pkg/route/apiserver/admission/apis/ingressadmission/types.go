package ingressadmission

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type IngressAdmissionConfig struct {
	metav1.TypeMeta
	AllowHostnameChanges	bool
}
