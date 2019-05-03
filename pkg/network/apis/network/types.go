package network

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	ClusterNetworkDefault       = "default"
	EgressNetworkPolicyMaxRules = 50
)

type ClusterNetwork struct {
	metav1.TypeMeta
	metav1.ObjectMeta
	ClusterNetworks  []ClusterNetworkEntry
	Network          string
	HostSubnetLength uint32
	ServiceNetwork   string
	PluginName       string
	VXLANPort        *uint32
}
type ClusterNetworkEntry struct {
	CIDR             string
	HostSubnetLength uint32
}
type ClusterNetworkList struct {
	metav1.TypeMeta
	metav1.ListMeta
	Items []ClusterNetwork
}
type HostSubnet struct {
	metav1.TypeMeta
	metav1.ObjectMeta
	Host        string
	HostIP      string
	Subnet      string
	EgressIPs   []string
	EgressCIDRs []string
}
type HostSubnetList struct {
	metav1.TypeMeta
	metav1.ListMeta
	Items []HostSubnet
}
type NetNamespace struct {
	metav1.TypeMeta
	metav1.ObjectMeta
	NetName   string
	NetID     uint32
	EgressIPs []string
}
type NetNamespaceList struct {
	metav1.TypeMeta
	metav1.ListMeta
	Items []NetNamespace
}
type EgressNetworkPolicyRuleType string

const (
	EgressNetworkPolicyRuleAllow EgressNetworkPolicyRuleType = "Allow"
	EgressNetworkPolicyRuleDeny  EgressNetworkPolicyRuleType = "Deny"
)

type EgressNetworkPolicyPeer struct {
	CIDRSelector string
	DNSName      string
}
type EgressNetworkPolicyRule struct {
	Type EgressNetworkPolicyRuleType
	To   EgressNetworkPolicyPeer
}
type EgressNetworkPolicySpec struct{ Egress []EgressNetworkPolicyRule }
type EgressNetworkPolicy struct {
	metav1.TypeMeta
	metav1.ObjectMeta
	Spec EgressNetworkPolicySpec
}
type EgressNetworkPolicyList struct {
	metav1.TypeMeta
	metav1.ListMeta
	Items []EgressNetworkPolicy
}
