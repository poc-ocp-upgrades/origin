package namespaceconditions

import (
	goformat "fmt"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/apiserver/pkg/admission"
	corev1client "k8s.io/client-go/kubernetes/typed/core/v1"
	corev1lister "k8s.io/client-go/listers/core/v1"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

var runLevelZeroNamespaces = sets.NewString("default", "kube-system", "kube-public")
var runLevelOneNamespaces = sets.NewString("openshift-node", "openshift-infra", "openshift")

func init() {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	runLevelOneNamespaces.Insert(runLevelZeroNamespaces.List()...)
}

type NamespaceLabelConditions struct {
	NamespaceClient    corev1client.NamespacesGetter
	NamespaceLister    corev1lister.NamespaceLister
	SkipLevelZeroNames sets.String
	SkipLevelOneNames  sets.String
}

func (d *NamespaceLabelConditions) WithNamespaceLabelConditions(admissionPlugin admission.Interface, name string) admission.Interface {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	switch {
	case d.SkipLevelOneNames.Has(name):
		return &pluginHandlerWithNamespaceNameConditions{admissionPlugin: &pluginHandlerWithNamespaceLabelConditions{admissionPlugin: admissionPlugin, namespaceClient: d.NamespaceClient, namespaceLister: d.NamespaceLister, namespaceSelector: skipRunLevelOneSelector}, namespacesToExclude: runLevelOneNamespaces}
	case d.SkipLevelZeroNames.Has(name):
		return &pluginHandlerWithNamespaceNameConditions{admissionPlugin: &pluginHandlerWithNamespaceLabelConditions{admissionPlugin: admissionPlugin, namespaceClient: d.NamespaceClient, namespaceLister: d.NamespaceLister, namespaceSelector: skipRunLevelZeroSelector}, namespacesToExclude: runLevelZeroNamespaces}
	default:
		return admissionPlugin
	}
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
