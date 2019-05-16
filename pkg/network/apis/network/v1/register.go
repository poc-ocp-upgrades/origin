package v1

import (
	goformat "fmt"
	"github.com/openshift/api/network/v1"
	"github.com/openshift/origin/pkg/network/apis/network"
	"k8s.io/apimachinery/pkg/runtime"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

var (
	localSchemeBuilder = runtime.NewSchemeBuilder(network.Install, v1.Install, RegisterDefaults)
	Install            = localSchemeBuilder.AddToScheme
)

func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
