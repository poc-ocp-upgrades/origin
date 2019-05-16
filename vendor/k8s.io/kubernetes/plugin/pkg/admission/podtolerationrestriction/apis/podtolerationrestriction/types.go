package podtolerationrestriction

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	api "k8s.io/kubernetes/pkg/apis/core"
)

type Configuration struct {
	metav1.TypeMeta
	Default   []api.Toleration
	Whitelist []api.Toleration
}
