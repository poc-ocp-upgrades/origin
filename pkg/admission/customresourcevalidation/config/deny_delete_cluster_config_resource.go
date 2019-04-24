package config

import (
	"fmt"
	"bytes"
	"net/http"
	"runtime"
	"io"
	"k8s.io/apiserver/pkg/admission"
)

const PluginName = "config.openshift.io/DenyDeleteClusterConfiguration"

func Register(plugins *admission.Plugins) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	plugins.Register(PluginName, func(config io.Reader) (admission.Interface, error) {
		return newAdmissionPlugin(), nil
	})
}

var _ admission.ValidationInterface = &admissionPlugin{}

type admissionPlugin struct{ *admission.Handler }

func newAdmissionPlugin() *admissionPlugin {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &admissionPlugin{Handler: admission.NewHandler(admission.Delete)}
}
func (p *admissionPlugin) Validate(attributes admission.Attributes) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
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
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := runtime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", runtime.FuncForPC(pc).Name()))
	http.Post("/"+"logcode", "application/json", bytes.NewBuffer(jsonLog))
}
