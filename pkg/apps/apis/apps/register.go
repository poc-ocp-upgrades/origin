package apps

import (
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	core "k8s.io/kubernetes/pkg/apis/core"
	"k8s.io/kubernetes/pkg/apis/extensions"
)

const (
	GroupName = "apps.openshift.io"
)

var (
	schemeBuilder      = runtime.NewSchemeBuilder(addKnownTypes, core.AddToScheme, extensions.AddToScheme)
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
	scheme.AddKnownTypes(SchemeGroupVersion, &DeploymentConfig{}, &DeploymentConfigList{}, &DeploymentConfigRollback{}, &DeploymentRequest{}, &DeploymentLog{}, &DeploymentLogOptions{})
	return nil
}
