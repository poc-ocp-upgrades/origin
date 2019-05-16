package authorization

import (
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/kubernetes/pkg/apis/core"
	"k8s.io/kubernetes/pkg/apis/rbac"
)

const (
	GroupName = "authorization.openshift.io"
)

var (
	schemeBuilder      = runtime.NewSchemeBuilder(addKnownTypes, core.AddToScheme, rbac.AddToScheme)
	Install            = schemeBuilder.AddToScheme
	SchemeGroupVersion = schema.GroupVersion{Group: GroupName, Version: runtime.APIVersionInternal}
	AddToScheme        = schemeBuilder.AddToScheme
)

func Resource(resource string) schema.GroupResource {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return SchemeGroupVersion.WithResource(resource).GroupResource()
}
func addKnownTypes(scheme *runtime.Scheme) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	scheme.AddKnownTypes(SchemeGroupVersion, &Role{}, &RoleBinding{}, &RoleBindingList{}, &RoleList{}, &SelfSubjectRulesReview{}, &SubjectRulesReview{}, &ResourceAccessReview{}, &SubjectAccessReview{}, &LocalResourceAccessReview{}, &LocalSubjectAccessReview{}, &ResourceAccessReviewResponse{}, &SubjectAccessReviewResponse{}, &IsPersonalSubjectAccessReview{}, &ClusterRole{}, &ClusterRoleBinding{}, &ClusterRoleBindingList{}, &ClusterRoleList{}, &RoleBindingRestriction{}, &RoleBindingRestrictionList{})
	return nil
}
