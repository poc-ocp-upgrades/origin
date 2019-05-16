package v1

import (
	goformat "fmt"
	"github.com/openshift/origin/pkg/network/admission/apis/restrictedendpoints"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

var GroupVersion = schema.GroupVersion{Group: "network.openshift.io", Version: "v1"}
var (
	schemeBuilder = runtime.NewSchemeBuilder(addKnownTypes, restrictedendpoints.Install)
	Install       = schemeBuilder.AddToScheme
)

func addKnownTypes(scheme *runtime.Scheme) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	scheme.AddKnownTypes(GroupVersion, &ExternalIPRangerAdmissionConfig{})
	return nil
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
