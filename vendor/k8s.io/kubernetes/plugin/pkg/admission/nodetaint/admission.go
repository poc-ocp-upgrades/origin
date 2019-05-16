package nodetaint

import (
	"fmt"
	goformat "fmt"
	"io"
	"k8s.io/apiserver/pkg/admission"
	utilfeature "k8s.io/apiserver/pkg/util/feature"
	api "k8s.io/kubernetes/pkg/apis/core"
	"k8s.io/kubernetes/pkg/features"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

const (
	PluginName        = "TaintNodesByCondition"
	TaintNodeNotReady = "node.kubernetes.io/not-ready"
)

func Register(plugins *admission.Plugins) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	plugins.Register(PluginName, func(config io.Reader) (admission.Interface, error) {
		return NewPlugin(), nil
	})
}
func NewPlugin() *Plugin {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &Plugin{Handler: admission.NewHandler(admission.Create), features: utilfeature.DefaultFeatureGate}
}

type Plugin struct {
	*admission.Handler
	features utilfeature.FeatureGate
}

var (
	_ = admission.Interface(&Plugin{})
)
var (
	nodeResource = api.Resource("nodes")
)

func (p *Plugin) Admit(a admission.Attributes) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if !p.features.Enabled(features.TaintNodesByCondition) {
		return nil
	}
	if a.GetResource().GroupResource() != nodeResource || a.GetSubresource() != "" {
		return nil
	}
	node, ok := a.GetObject().(*api.Node)
	if !ok {
		return admission.NewForbidden(a, fmt.Errorf("unexpected type %T", a.GetObject()))
	}
	addNotReadyTaint(node)
	return nil
}
func addNotReadyTaint(node *api.Node) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	notReadyTaint := api.Taint{Key: TaintNodeNotReady, Effect: api.TaintEffectNoSchedule}
	for _, taint := range node.Spec.Taints {
		if taint.MatchTaint(notReadyTaint) {
			return
		}
	}
	node.Spec.Taints = append(node.Spec.Taints, notReadyTaint)
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
