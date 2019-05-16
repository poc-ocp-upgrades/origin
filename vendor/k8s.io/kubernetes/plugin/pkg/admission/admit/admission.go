package admit

import (
	goformat "fmt"
	"io"
	"k8s.io/apiserver/pkg/admission"
	"k8s.io/klog"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

const PluginName = "AlwaysAdmit"

func Register(plugins *admission.Plugins) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	plugins.Register(PluginName, func(config io.Reader) (admission.Interface, error) {
		return NewAlwaysAdmit(), nil
	})
}

type alwaysAdmit struct{}

var _ admission.MutationInterface = alwaysAdmit{}
var _ admission.ValidationInterface = alwaysAdmit{}

func (alwaysAdmit) Admit(a admission.Attributes) (err error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return nil
}
func (alwaysAdmit) Validate(a admission.Attributes) (err error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return nil
}
func (alwaysAdmit) Handles(operation admission.Operation) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return true
}
func NewAlwaysAdmit() admission.Interface {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	klog.Warningf("%s admission controller is deprecated. "+"Please remove this controller from your configuration files and scripts.", PluginName)
	return new(alwaysAdmit)
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
