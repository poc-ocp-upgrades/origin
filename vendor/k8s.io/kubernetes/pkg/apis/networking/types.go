package networking

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	api "k8s.io/kubernetes/pkg/apis/core"
)

type NetworkPolicy struct {
	metav1.TypeMeta
	metav1.ObjectMeta
	Spec NetworkPolicySpec
}
type PolicyType string

const (
	PolicyTypeIngress PolicyType = "Ingress"
	PolicyTypeEgress  PolicyType = "Egress"
)

type NetworkPolicySpec struct {
	PodSelector metav1.LabelSelector
	Ingress     []NetworkPolicyIngressRule
	Egress      []NetworkPolicyEgressRule
	PolicyTypes []PolicyType
}
type NetworkPolicyIngressRule struct {
	Ports []NetworkPolicyPort
	From  []NetworkPolicyPeer
}
type NetworkPolicyEgressRule struct {
	Ports []NetworkPolicyPort
	To    []NetworkPolicyPeer
}
type NetworkPolicyPort struct {
	Protocol *api.Protocol
	Port     *intstr.IntOrString
}
type IPBlock struct {
	CIDR   string
	Except []string
}
type NetworkPolicyPeer struct {
	PodSelector       *metav1.LabelSelector
	NamespaceSelector *metav1.LabelSelector
	IPBlock           *IPBlock
}
type NetworkPolicyList struct {
	metav1.TypeMeta
	metav1.ListMeta
	Items []NetworkPolicy
}
