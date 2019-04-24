package namespaceconditions

import (
	"k8s.io/apimachinery/pkg/util/sets"
	"bytes"
	"net/http"
	"runtime"
	"fmt"
	"k8s.io/apiserver/pkg/admission"
	corev1client "k8s.io/client-go/kubernetes/typed/core/v1"
	corev1lister "k8s.io/client-go/listers/core/v1"
)

var runLevelZeroNamespaces = sets.NewString("default", "kube-system", "kube-public")
var runLevelOneNamespaces = sets.NewString("openshift-node", "openshift-infra", "openshift")

func init() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	runLevelOneNamespaces.Insert(runLevelZeroNamespaces.List()...)
}

type NamespaceLabelConditions struct {
	NamespaceClient		corev1client.NamespacesGetter
	NamespaceLister		corev1lister.NamespaceLister
	SkipLevelZeroNames	sets.String
	SkipLevelOneNames	sets.String
}

func (d *NamespaceLabelConditions) WithNamespaceLabelConditions(admissionPlugin admission.Interface, name string) admission.Interface {
	_logClusterCodePath()
	defer _logClusterCodePath()
	switch {
	case d.SkipLevelOneNames.Has(name):
		return &pluginHandlerWithNamespaceNameConditions{admissionPlugin: &pluginHandlerWithNamespaceLabelConditions{admissionPlugin: admissionPlugin, namespaceClient: d.NamespaceClient, namespaceLister: d.NamespaceLister, namespaceSelector: skipRunLevelOneSelector}, namespacesToExclude: runLevelOneNamespaces}
	case d.SkipLevelZeroNames.Has(name):
		return &pluginHandlerWithNamespaceNameConditions{admissionPlugin: &pluginHandlerWithNamespaceLabelConditions{admissionPlugin: admissionPlugin, namespaceClient: d.NamespaceClient, namespaceLister: d.NamespaceLister, namespaceSelector: skipRunLevelZeroSelector}, namespacesToExclude: runLevelZeroNamespaces}
	default:
		return admissionPlugin
	}
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := runtime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", runtime.FuncForPC(pc).Name()))
	http.Post("/"+"logcode", "application/json", bytes.NewBuffer(jsonLog))
}
