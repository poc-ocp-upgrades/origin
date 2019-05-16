package alwayspullimages

import (
	goformat "fmt"
	"io"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"k8s.io/apiserver/pkg/admission"
	api "k8s.io/kubernetes/pkg/apis/core"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

const PluginName = "AlwaysPullImages"

func Register(plugins *admission.Plugins) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	plugins.Register(PluginName, func(config io.Reader) (admission.Interface, error) {
		return NewAlwaysPullImages(), nil
	})
}

type AlwaysPullImages struct{ *admission.Handler }

var _ admission.MutationInterface = &AlwaysPullImages{}
var _ admission.ValidationInterface = &AlwaysPullImages{}

func (a *AlwaysPullImages) Admit(attributes admission.Attributes) (err error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if shouldIgnore(attributes) {
		return nil
	}
	pod, ok := attributes.GetObject().(*api.Pod)
	if !ok {
		return apierrors.NewBadRequest("Resource was marked with kind Pod but was unable to be converted")
	}
	for i := range pod.Spec.InitContainers {
		pod.Spec.InitContainers[i].ImagePullPolicy = api.PullAlways
	}
	for i := range pod.Spec.Containers {
		pod.Spec.Containers[i].ImagePullPolicy = api.PullAlways
	}
	return nil
}
func (*AlwaysPullImages) Validate(attributes admission.Attributes) (err error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if shouldIgnore(attributes) {
		return nil
	}
	pod, ok := attributes.GetObject().(*api.Pod)
	if !ok {
		return apierrors.NewBadRequest("Resource was marked with kind Pod but was unable to be converted")
	}
	for i := range pod.Spec.InitContainers {
		if pod.Spec.InitContainers[i].ImagePullPolicy != api.PullAlways {
			return admission.NewForbidden(attributes, field.NotSupported(field.NewPath("spec", "initContainers").Index(i).Child("imagePullPolicy"), pod.Spec.InitContainers[i].ImagePullPolicy, []string{string(api.PullAlways)}))
		}
	}
	for i := range pod.Spec.Containers {
		if pod.Spec.Containers[i].ImagePullPolicy != api.PullAlways {
			return admission.NewForbidden(attributes, field.NotSupported(field.NewPath("spec", "containers").Index(i).Child("imagePullPolicy"), pod.Spec.Containers[i].ImagePullPolicy, []string{string(api.PullAlways)}))
		}
	}
	return nil
}
func shouldIgnore(attributes admission.Attributes) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if len(attributes.GetSubresource()) != 0 || attributes.GetResource().GroupResource() != api.Resource("pods") {
		return true
	}
	return false
}
func NewAlwaysPullImages() *AlwaysPullImages {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &AlwaysPullImages{Handler: admission.NewHandler(admission.Create, admission.Update)}
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
