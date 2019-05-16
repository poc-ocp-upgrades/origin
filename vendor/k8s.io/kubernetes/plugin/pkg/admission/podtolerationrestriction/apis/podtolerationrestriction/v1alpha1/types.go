package v1alpha1

import (
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Configuration struct {
	metav1.TypeMeta `json:",inline"`
	Default         []v1.Toleration `json:"default,omitempty"`
	Whitelist       []v1.Toleration `json:"whitelist,omitempty"`
}
