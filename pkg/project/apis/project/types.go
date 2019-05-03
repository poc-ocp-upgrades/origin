package project

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/kubernetes/pkg/apis/core"
)

type ProjectList struct {
	metav1.TypeMeta
	metav1.ListMeta
	Items []Project
}

const (
	FinalizerOrigin core.FinalizerName = "openshift.io/origin"
)

type ProjectSpec struct{ Finalizers []core.FinalizerName }
type ProjectStatus struct{ Phase core.NamespacePhase }
type Project struct {
	metav1.TypeMeta
	metav1.ObjectMeta
	Spec   ProjectSpec
	Status ProjectStatus
}
type ProjectRequest struct {
	metav1.TypeMeta
	metav1.ObjectMeta
	DisplayName string
	Description string
}

const (
	ProjectNodeSelector = "openshift.io/node-selector"
	ProjectRequester    = "openshift.io/requester"
)
