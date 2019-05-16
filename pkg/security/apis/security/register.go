package security

import (
	goformat "fmt"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/kubernetes/pkg/apis/core"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

const (
	GroupName = "security.openshift.io"
)

var (
	schemeBuilder      = runtime.NewSchemeBuilder(addKnownTypes, core.AddToScheme)
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
	scheme.AddKnownTypes(SchemeGroupVersion, &SecurityContextConstraints{}, &SecurityContextConstraintsList{}, &PodSecurityPolicySubjectReview{}, &PodSecurityPolicySelfSubjectReview{}, &PodSecurityPolicyReview{}, &RangeAllocation{}, &RangeAllocationList{})
	return nil
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
