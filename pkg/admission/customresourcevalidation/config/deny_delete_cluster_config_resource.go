package config

import (
	"fmt"
	goformat "fmt"
	"io"
	"k8s.io/apiserver/pkg/admission"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

const PluginName = "config.openshift.io/DenyDeleteClusterConfiguration"

func Register(plugins *admission.Plugins) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	plugins.Register(PluginName, func(config io.Reader) (admission.Interface, error) {
		return newAdmissionPlugin(), nil
	})
}

var _ admission.ValidationInterface = &admissionPlugin{}

type admissionPlugin struct{ *admission.Handler }

func newAdmissionPlugin() *admissionPlugin {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &admissionPlugin{Handler: admission.NewHandler(admission.Delete)}
}
func (p *admissionPlugin) Validate(attributes admission.Attributes) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if len(attributes.GetSubresource()) > 0 {
		return nil
	}
	if attributes.GetResource().Group != "config.openshift.io" {
		return nil
	}
	switch attributes.GetResource().Resource {
	case "clusteroperators":
		return nil
	case "clusterversions":
		if attributes.GetName() != "version" {
			return nil
		}
	default:
		if attributes.GetName() != "cluster" {
			return nil
		}
	}
	return admission.NewForbidden(attributes, fmt.Errorf("deleting required %s.%s resource, named %s, is not allowed", attributes.GetResource().Resource, attributes.GetResource().Group, attributes.GetName()))
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
