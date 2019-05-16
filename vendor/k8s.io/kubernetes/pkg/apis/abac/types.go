package abac

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Policy struct {
	metav1.TypeMeta
	Spec PolicySpec
}
type PolicySpec struct {
	User            string
	Group           string
	Readonly        bool
	APIGroup        string
	Resource        string
	Namespace       string
	NonResourcePath string
}
