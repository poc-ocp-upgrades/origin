package route

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/kubernetes/pkg/apis/core"
)

type Route struct {
	metav1.TypeMeta
	metav1.ObjectMeta
	Spec	RouteSpec
	Status	RouteStatus
}
type RouteSpec struct {
	Host			string
	Subdomain		string
	Path			string
	To			RouteTargetReference
	AlternateBackends	[]RouteTargetReference
	Port			*RoutePort
	TLS			*TLSConfig
	WildcardPolicy		WildcardPolicyType
}
type RouteTargetReference struct {
	Kind	string
	Name	string
	Weight	*int32
}
type RoutePort struct{ TargetPort intstr.IntOrString }
type RouteStatus struct{ Ingress []RouteIngress }
type RouteIngress struct {
	Host			string
	RouterName		string
	Conditions		[]RouteIngressCondition
	WildcardPolicy		WildcardPolicyType
	RouterCanonicalHostname	string
}
type RouteIngressConditionType string

const (
	RouteAdmitted			RouteIngressConditionType	= "Admitted"
	RouteExtendedValidationFailed	RouteIngressConditionType	= "ExtendedValidationFailed"
)

type RouteIngressCondition struct {
	Type			RouteIngressConditionType
	Status			core.ConditionStatus
	Reason			string
	Message			string
	LastTransitionTime	*metav1.Time
}
type RouteList struct {
	metav1.TypeMeta
	metav1.ListMeta
	Items	[]Route
}
type RouterShard struct {
	ShardName	string
	DNSSuffix	string
}
type TLSConfig struct {
	Termination			TLSTerminationType
	Certificate			string
	Key				string
	CACertificate			string
	DestinationCACertificate	string
	InsecureEdgeTerminationPolicy	InsecureEdgeTerminationPolicyType
}
type TLSTerminationType string
type InsecureEdgeTerminationPolicyType string

const (
	TLSTerminationEdge			TLSTerminationType			= "edge"
	TLSTerminationPassthrough		TLSTerminationType			= "passthrough"
	TLSTerminationReencrypt			TLSTerminationType			= "reencrypt"
	InsecureEdgeTerminationPolicyNone	InsecureEdgeTerminationPolicyType	= "None"
	InsecureEdgeTerminationPolicyAllow	InsecureEdgeTerminationPolicyType	= "Allow"
	InsecureEdgeTerminationPolicyRedirect	InsecureEdgeTerminationPolicyType	= "Redirect"
)

type WildcardPolicyType string

const (
	WildcardPolicyNone	WildcardPolicyType	= "None"
	WildcardPolicySubdomain	WildcardPolicyType	= "Subdomain"
)
