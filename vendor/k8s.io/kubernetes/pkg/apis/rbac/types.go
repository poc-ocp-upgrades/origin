package rbac

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	APIGroupAll             = "*"
	ResourceAll             = "*"
	VerbAll                 = "*"
	NonResourceAll          = "*"
	GroupKind               = "Group"
	ServiceAccountKind      = "ServiceAccount"
	UserKind                = "User"
	AutoUpdateAnnotationKey = "rbac.authorization.kubernetes.io/autoupdate"
)

type PolicyRule struct {
	Verbs           []string
	APIGroups       []string
	Resources       []string
	ResourceNames   []string
	NonResourceURLs []string
}
type Subject struct {
	Kind      string
	APIGroup  string
	Name      string
	Namespace string
}
type RoleRef struct {
	APIGroup string
	Kind     string
	Name     string
}
type Role struct {
	metav1.TypeMeta
	metav1.ObjectMeta
	Rules []PolicyRule
}
type RoleBinding struct {
	metav1.TypeMeta
	metav1.ObjectMeta
	Subjects []Subject
	RoleRef  RoleRef
}
type RoleBindingList struct {
	metav1.TypeMeta
	metav1.ListMeta
	Items []RoleBinding
}
type RoleList struct {
	metav1.TypeMeta
	metav1.ListMeta
	Items []Role
}
type ClusterRole struct {
	metav1.TypeMeta
	metav1.ObjectMeta
	Rules           []PolicyRule
	AggregationRule *AggregationRule
}
type AggregationRule struct{ ClusterRoleSelectors []metav1.LabelSelector }
type ClusterRoleBinding struct {
	metav1.TypeMeta
	metav1.ObjectMeta
	Subjects []Subject
	RoleRef  RoleRef
}
type ClusterRoleBindingList struct {
	metav1.TypeMeta
	metav1.ListMeta
	Items []ClusterRoleBinding
}
type ClusterRoleList struct {
	metav1.TypeMeta
	metav1.ListMeta
	Items []ClusterRole
}
