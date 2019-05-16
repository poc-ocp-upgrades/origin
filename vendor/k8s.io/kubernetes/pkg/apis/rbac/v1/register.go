package v1

import (
	rbacv1 "k8s.io/api/rbac/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

const GroupName = "rbac.authorization.k8s.io"

var SchemeGroupVersion = schema.GroupVersion{Group: GroupName, Version: "v1"}

func Resource(resource string) schema.GroupResource {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return SchemeGroupVersion.WithResource(resource).GroupResource()
}

var (
	localSchemeBuilder = &rbacv1.SchemeBuilder
	AddToScheme        = localSchemeBuilder.AddToScheme
)

func init() {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	localSchemeBuilder.Register(addDefaultingFuncs)
}
