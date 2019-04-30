package testing

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

type AdmissionPluginTestConfig struct {
	metav1.TypeMeta
	Data	string	`json:"data"`
}

func (obj *AdmissionPluginTestConfig) GetObjectKind() schema.ObjectKind {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &obj.TypeMeta
}
