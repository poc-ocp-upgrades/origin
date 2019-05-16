package v1

import (
	goformat "fmt"
	"github.com/openshift/origin/pkg/route/apiserver/admission/apis/ingressadmission"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

func (obj *IngressAdmissionConfig) GetObjectKind() schema.ObjectKind {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &obj.TypeMeta
}

var GroupVersion = schema.GroupVersion{Group: "route.openshift.io", Version: "v1"}
var (
	schemeBuilder = runtime.NewSchemeBuilder(addKnownTypes, ingressadmission.Install)
	Install       = schemeBuilder.AddToScheme
)

func addKnownTypes(scheme *runtime.Scheme) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	scheme.AddKnownTypes(GroupVersion, &IngressAdmissionConfig{})
	return nil
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
