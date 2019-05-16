package options

import (
	goformat "fmt"
	"github.com/spf13/pflag"
	_ "k8s.io/apiserver/pkg/admission"
	"k8s.io/apiserver/pkg/util/globalflag"
	_ "k8s.io/kubernetes/pkg/cloudprovider/providers"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

func AddCustomGlobalFlags(fs *pflag.FlagSet) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	globalflag.Register(fs, "cloud-provider-gce-lb-src-cidrs")
	globalflag.Register(fs, "default-not-ready-toleration-seconds")
	globalflag.Register(fs, "default-unreachable-toleration-seconds")
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
