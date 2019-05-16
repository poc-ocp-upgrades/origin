package v1alpha1

import (
	goformat "fmt"
	imagepolicyv1alpha1 "k8s.io/api/imagepolicy/v1alpha1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

const GroupName = "imagepolicy.k8s.io"

var SchemeGroupVersion = schema.GroupVersion{Group: GroupName, Version: "v1alpha1"}

func Resource(resource string) schema.GroupResource {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return SchemeGroupVersion.WithResource(resource).GroupResource()
}

var (
	localSchemeBuilder = &imagepolicyv1alpha1.SchemeBuilder
	AddToScheme        = localSchemeBuilder.AddToScheme
)

func init() {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	localSchemeBuilder.Register(RegisterDefaults)
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
