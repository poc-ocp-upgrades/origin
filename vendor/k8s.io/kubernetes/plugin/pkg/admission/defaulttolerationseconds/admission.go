package defaulttolerationseconds

import (
	"flag"
	"fmt"
	goformat "fmt"
	"io"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apiserver/pkg/admission"
	api "k8s.io/kubernetes/pkg/apis/core"
	schedulerapi "k8s.io/kubernetes/pkg/scheduler/api"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

const PluginName = "DefaultTolerationSeconds"

var (
	defaultNotReadyTolerationSeconds    = flag.Int64("default-not-ready-toleration-seconds", 300, "Indicates the tolerationSeconds of the toleration for notReady:NoExecute"+" that is added by default to every pod that does not already have such a toleration.")
	defaultUnreachableTolerationSeconds = flag.Int64("default-unreachable-toleration-seconds", 300, "Indicates the tolerationSeconds of the toleration for unreachable:NoExecute"+" that is added by default to every pod that does not already have such a toleration.")
	notReadyToleration                  = api.Toleration{Key: schedulerapi.TaintNodeNotReady, Operator: api.TolerationOpExists, Effect: api.TaintEffectNoExecute, TolerationSeconds: defaultNotReadyTolerationSeconds}
	unreachableToleration               = api.Toleration{Key: schedulerapi.TaintNodeUnreachable, Operator: api.TolerationOpExists, Effect: api.TaintEffectNoExecute, TolerationSeconds: defaultUnreachableTolerationSeconds}
)

func Register(plugins *admission.Plugins) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	plugins.Register(PluginName, func(config io.Reader) (admission.Interface, error) {
		return NewDefaultTolerationSeconds(), nil
	})
}

type Plugin struct{ *admission.Handler }

var _ admission.MutationInterface = &Plugin{}

func NewDefaultTolerationSeconds() *Plugin {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &Plugin{Handler: admission.NewHandler(admission.Create, admission.Update)}
}
func (p *Plugin) Admit(attributes admission.Attributes) (err error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if attributes.GetResource().GroupResource() != api.Resource("pods") {
		return nil
	}
	if len(attributes.GetSubresource()) > 0 {
		return nil
	}
	pod, ok := attributes.GetObject().(*api.Pod)
	if !ok {
		return errors.NewBadRequest(fmt.Sprintf("expected *api.Pod but got %T", attributes.GetObject()))
	}
	tolerations := pod.Spec.Tolerations
	toleratesNodeNotReady := false
	toleratesNodeUnreachable := false
	for _, toleration := range tolerations {
		if (toleration.Key == schedulerapi.TaintNodeNotReady || len(toleration.Key) == 0) && (toleration.Effect == api.TaintEffectNoExecute || len(toleration.Effect) == 0) {
			toleratesNodeNotReady = true
		}
		if (toleration.Key == schedulerapi.TaintNodeUnreachable || len(toleration.Key) == 0) && (toleration.Effect == api.TaintEffectNoExecute || len(toleration.Effect) == 0) {
			toleratesNodeUnreachable = true
		}
	}
	if !toleratesNodeNotReady {
		pod.Spec.Tolerations = append(pod.Spec.Tolerations, notReadyToleration)
	}
	if !toleratesNodeUnreachable {
		pod.Spec.Tolerations = append(pod.Spec.Tolerations, unreachableToleration)
	}
	return nil
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
