package deny

import (
	"errors"
	goformat "fmt"
	"io"
	"k8s.io/apiserver/pkg/admission"
	"k8s.io/klog"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

const PluginName = "AlwaysDeny"

func Register(plugins *admission.Plugins) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	plugins.Register(PluginName, func(config io.Reader) (admission.Interface, error) {
		return NewAlwaysDeny(), nil
	})
}

type alwaysDeny struct{}

var _ admission.MutationInterface = alwaysDeny{}
var _ admission.ValidationInterface = alwaysDeny{}

func (alwaysDeny) Admit(a admission.Attributes) (err error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return admission.NewForbidden(a, errors.New("admission control is denying all modifications"))
}
func (alwaysDeny) Validate(a admission.Attributes) (err error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return admission.NewForbidden(a, errors.New("admission control is denying all modifications"))
}
func (alwaysDeny) Handles(operation admission.Operation) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return true
}
func NewAlwaysDeny() admission.Interface {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	klog.Warningf("%s admission controller is deprecated. "+"Please remove this controller from your configuration files and scripts.", PluginName)
	return new(alwaysDeny)
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
