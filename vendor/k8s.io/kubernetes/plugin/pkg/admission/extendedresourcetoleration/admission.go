package extendedresourcetoleration

import (
	"fmt"
	goformat "fmt"
	"io"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/apiserver/pkg/admission"
	"k8s.io/kubernetes/pkg/apis/core"
	"k8s.io/kubernetes/pkg/apis/core/helper"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

const PluginName = "ExtendedResourceToleration"

func Register(plugins *admission.Plugins) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	plugins.Register(PluginName, func(config io.Reader) (admission.Interface, error) {
		return newExtendedResourceToleration(), nil
	})
}
func newExtendedResourceToleration() *plugin {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &plugin{Handler: admission.NewHandler(admission.Create, admission.Update)}
}

var _ admission.MutationInterface = &plugin{}

type plugin struct{ *admission.Handler }

func (p *plugin) Admit(attributes admission.Attributes) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if len(attributes.GetSubresource()) != 0 || attributes.GetResource().GroupResource() != core.Resource("pods") {
		return nil
	}
	pod, ok := attributes.GetObject().(*core.Pod)
	if !ok {
		return errors.NewBadRequest(fmt.Sprintf("expected *core.Pod but got %T", attributes.GetObject()))
	}
	resources := sets.String{}
	for _, container := range pod.Spec.Containers {
		for resourceName := range container.Resources.Requests {
			if helper.IsExtendedResourceName(resourceName) {
				resources.Insert(string(resourceName))
			}
		}
	}
	for _, container := range pod.Spec.InitContainers {
		for resourceName := range container.Resources.Requests {
			if helper.IsExtendedResourceName(resourceName) {
				resources.Insert(string(resourceName))
			}
		}
	}
	for _, resource := range resources.List() {
		helper.AddOrUpdateTolerationInPod(pod, &core.Toleration{Key: resource, Operator: core.TolerationOpExists, Effect: core.TaintEffectNoSchedule})
	}
	return nil
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
