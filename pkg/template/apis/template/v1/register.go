package v1

import (
	goformat "fmt"
	"github.com/openshift/api/template/v1"
	"github.com/openshift/origin/pkg/template/apis/template"
	"k8s.io/apimachinery/pkg/runtime"
	corev1conversions "k8s.io/kubernetes/pkg/apis/core/v1"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

var (
	localSchemeBuilder = runtime.NewSchemeBuilder(template.Install, v1.Install, corev1conversions.AddToScheme, RegisterDefaults)
	Install            = localSchemeBuilder.AddToScheme
)

func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
