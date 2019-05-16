package v1alpha1

import (
	goformat "fmt"
	kruntime "k8s.io/apimachinery/pkg/runtime"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

func addDefaultingFuncs(scheme *kruntime.Scheme) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return RegisterDefaults(scheme)
}
func SetDefaults_Configuration(obj *Configuration) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
