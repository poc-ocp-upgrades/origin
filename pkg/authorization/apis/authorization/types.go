package authorization

import (
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kruntime "k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/sets"
	kapi "k8s.io/kubernetes/pkg/apis/core"
	rbacapi "k8s.io/kubernetes/pkg/apis/rbac"
)

const (
	PolicyName		= "default"
	APIGroupAll		= "*"
	ResourceAll		= "*"
	VerbAll			= "*"
	NonResourceAll		= "*"
	ScopesKey		= "scopes.authorization.openshift.io"
	ScopesAllNamespaces	= "*"
	UserKind		= "User"
	GroupKind		= "Group"
	ServiceAccountKind	= "ServiceAccount"
	SystemUserKind		= "SystemUser"
	SystemGroupKind		= "SystemGroup"
	UserResource		= "users"
	GroupResource		= "groups"
	ServiceAccountResource	= "serviceaccounts"
	SystemUserResource	= "systemusers"
	SystemGroupResource	= "systemgroups"
)

var DiscoveryRule = rbacv1.PolicyRule{Verbs: []string{"get"}, NonResourceURLs: []string{"/version", "/version/*", "/api", "/api/*", "/apis", "/apis/*", "/oapi", "/oapi/*", "/openapi/v2", "/swaggerapi", "/swaggerapi/*", "/swagger.json", "/swagger-2.0.0.pb-v1", "/osapi", "/osapi/", "/.well-known", "/.well-known/*", "/"}}

type PolicyRule struct {
	Verbs			sets.String
	AttributeRestrictions	kruntime.Object
	APIGroups		[]string
	Resources		sets.String
	ResourceNames		sets.String
	NonResourceURLs		sets.String
}
type IsPersonalSubjectAccessReview struct{ metav1.TypeMeta }
type Role struct {
	metav1.TypeMeta
	metav1.ObjectMeta
	Rules	[]PolicyRule
}
type RoleBinding struct {
	metav1.TypeMeta
	metav1.ObjectMeta
	Subjects	[]kapi.ObjectReference
	RoleRef		kapi.ObjectReference
}
type SelfSubjectRulesReview struct {
	metav1.TypeMeta
	Spec	SelfSubjectRulesReviewSpec
	Status	SubjectRulesReviewStatus
}
type SelfSubjectRulesReviewSpec struct{ Scopes []string }
type SubjectRulesReview struct {
	metav1.TypeMeta
	Spec	SubjectRulesReviewSpec
	Status	SubjectRulesReviewStatus
}
type SubjectRulesReviewSpec struct {
	User	string
	Groups	[]string
	Scopes	[]string
}
type SubjectRulesReviewStatus struct {
	Rules		[]PolicyRule
	EvaluationError	string
}
type ResourceAccessReviewResponse struct {
	metav1.TypeMeta
	Namespace	string
	Users		sets.String
	Groups		sets.String
	EvaluationError	string
}
type ResourceAccessReview struct {
	metav1.TypeMeta
	Action
}
type SubjectAccessReviewResponse struct {
	metav1.TypeMeta
	Namespace	string
	Allowed		bool
	Reason		string
	EvaluationError	string
}
type SubjectAccessReview struct {
	metav1.TypeMeta
	Action
	User	string
	Groups	sets.String
	Scopes	[]string
}
type LocalResourceAccessReview struct {
	metav1.TypeMeta
	Action
}
type LocalSubjectAccessReview struct {
	metav1.TypeMeta
	Action
	User	string
	Groups	sets.String
	Scopes	[]string
}
type Action struct {
	Namespace		string
	Verb			string
	Group			string
	Version			string
	Resource		string
	ResourceName		string
	Path			string
	IsNonResourceURL	bool
	Content			kruntime.Object
}
type RoleBindingList struct {
	metav1.TypeMeta
	metav1.ListMeta
	Items	[]RoleBinding
}
type RoleList struct {
	metav1.TypeMeta
	metav1.ListMeta
	Items	[]Role
}
type ClusterRole struct {
	metav1.TypeMeta
	metav1.ObjectMeta
	Rules		[]PolicyRule
	AggregationRule	*rbacapi.AggregationRule
}
type ClusterRoleBinding struct {
	metav1.TypeMeta
	metav1.ObjectMeta
	Subjects	[]kapi.ObjectReference
	RoleRef		kapi.ObjectReference
}
type ClusterRoleBindingList struct {
	metav1.TypeMeta
	metav1.ListMeta
	Items	[]ClusterRoleBinding
}
type ClusterRoleList struct {
	metav1.TypeMeta
	metav1.ListMeta
	Items	[]ClusterRole
}
type RoleBindingRestriction struct {
	metav1.TypeMeta
	metav1.ObjectMeta
	Spec	RoleBindingRestrictionSpec
}
type RoleBindingRestrictionSpec struct {
	UserRestriction			*UserRestriction
	GroupRestriction		*GroupRestriction
	ServiceAccountRestriction	*ServiceAccountRestriction
}
type RoleBindingRestrictionList struct {
	metav1.TypeMeta
	metav1.ListMeta
	Items	[]RoleBindingRestriction
}
type UserRestriction struct {
	Users		[]string
	Groups		[]string
	Selectors	[]metav1.LabelSelector
}
type GroupRestriction struct {
	Groups		[]string
	Selectors	[]metav1.LabelSelector
}
type ServiceAccountRestriction struct {
	ServiceAccounts	[]ServiceAccountReference
	Namespaces	[]string
}
type ServiceAccountReference struct {
	Name		string
	Namespace	string
}
