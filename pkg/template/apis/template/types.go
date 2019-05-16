package template

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/kubernetes/pkg/apis/core"
)

type Template struct {
	metav1.TypeMeta
	metav1.ObjectMeta
	Message      string
	Parameters   []Parameter
	Objects      []runtime.Object
	ObjectLabels map[string]string
}
type TemplateList struct {
	metav1.TypeMeta
	metav1.ListMeta
	Items []Template
}
type Parameter struct {
	Name        string
	DisplayName string
	Description string
	Value       string
	Generate    string
	From        string
	Required    bool
}
type TemplateInstance struct {
	metav1.TypeMeta
	metav1.ObjectMeta
	Spec   TemplateInstanceSpec
	Status TemplateInstanceStatus
}
type TemplateInstanceSpec struct {
	Template  Template
	Secret    *core.LocalObjectReference
	Requester *TemplateInstanceRequester
}
type TemplateInstanceRequester struct {
	Username string
	UID      string
	Groups   []string
	Extra    map[string]ExtraValue
}
type ExtraValue []string
type TemplateInstanceStatus struct {
	Conditions []TemplateInstanceCondition
	Objects    []TemplateInstanceObject
}
type TemplateInstanceCondition struct {
	Type               TemplateInstanceConditionType
	Status             core.ConditionStatus
	LastTransitionTime metav1.Time
	Reason             string
	Message            string
}
type TemplateInstanceConditionType string

const (
	TemplateInstanceReady              TemplateInstanceConditionType = "Ready"
	TemplateInstanceInstantiateFailure TemplateInstanceConditionType = "InstantiateFailure"
)

type TemplateInstanceObject struct{ Ref core.ObjectReference }
type TemplateInstanceList struct {
	metav1.TypeMeta
	metav1.ListMeta
	Items []TemplateInstance
}
type BrokerTemplateInstance struct {
	metav1.TypeMeta
	metav1.ObjectMeta
	Spec BrokerTemplateInstanceSpec
}
type BrokerTemplateInstanceSpec struct {
	TemplateInstance core.ObjectReference
	Secret           core.ObjectReference
	BindingIDs       []string
}
type BrokerTemplateInstanceList struct {
	metav1.TypeMeta
	metav1.ListMeta
	Items []BrokerTemplateInstance
}
