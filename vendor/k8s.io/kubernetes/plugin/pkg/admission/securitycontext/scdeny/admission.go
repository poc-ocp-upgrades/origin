package scdeny

import (
	"fmt"
	goformat "fmt"
	"io"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apiserver/pkg/admission"
	api "k8s.io/kubernetes/pkg/apis/core"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

const PluginName = "SecurityContextDeny"

func Register(plugins *admission.Plugins) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	plugins.Register(PluginName, func(config io.Reader) (admission.Interface, error) {
		return NewSecurityContextDeny(), nil
	})
}

type Plugin struct{ *admission.Handler }

var _ admission.ValidationInterface = &Plugin{}

func NewSecurityContextDeny() *Plugin {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &Plugin{Handler: admission.NewHandler(admission.Create, admission.Update)}
}
func (p *Plugin) Validate(a admission.Attributes) (err error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if a.GetSubresource() != "" || a.GetResource().GroupResource() != api.Resource("pods") {
		return nil
	}
	pod, ok := a.GetObject().(*api.Pod)
	if !ok {
		return apierrors.NewBadRequest("Resource was marked with kind Pod but was unable to be converted")
	}
	if pod.Spec.SecurityContext != nil {
		if pod.Spec.SecurityContext.SupplementalGroups != nil {
			return apierrors.NewForbidden(a.GetResource().GroupResource(), pod.Name, fmt.Errorf("pod.Spec.SecurityContext.SupplementalGroups is forbidden"))
		}
		if pod.Spec.SecurityContext.SELinuxOptions != nil {
			return apierrors.NewForbidden(a.GetResource().GroupResource(), pod.Name, fmt.Errorf("pod.Spec.SecurityContext.SELinuxOptions is forbidden"))
		}
		if pod.Spec.SecurityContext.RunAsUser != nil {
			return apierrors.NewForbidden(a.GetResource().GroupResource(), pod.Name, fmt.Errorf("pod.Spec.SecurityContext.RunAsUser is forbidden"))
		}
		if pod.Spec.SecurityContext.FSGroup != nil {
			return apierrors.NewForbidden(a.GetResource().GroupResource(), pod.Name, fmt.Errorf("pod.Spec.SecurityContext.FSGroup is forbidden"))
		}
	}
	for _, v := range pod.Spec.InitContainers {
		if v.SecurityContext != nil {
			if v.SecurityContext.SELinuxOptions != nil {
				return apierrors.NewForbidden(a.GetResource().GroupResource(), pod.Name, fmt.Errorf("SecurityContext.SELinuxOptions is forbidden"))
			}
			if v.SecurityContext.RunAsUser != nil {
				return apierrors.NewForbidden(a.GetResource().GroupResource(), pod.Name, fmt.Errorf("SecurityContext.RunAsUser is forbidden"))
			}
		}
	}
	for _, v := range pod.Spec.Containers {
		if v.SecurityContext != nil {
			if v.SecurityContext.SELinuxOptions != nil {
				return apierrors.NewForbidden(a.GetResource().GroupResource(), pod.Name, fmt.Errorf("SecurityContext.SELinuxOptions is forbidden"))
			}
			if v.SecurityContext.RunAsUser != nil {
				return apierrors.NewForbidden(a.GetResource().GroupResource(), pod.Name, fmt.Errorf("SecurityContext.RunAsUser is forbidden"))
			}
		}
	}
	return nil
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
