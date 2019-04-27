package legacy

import (
	"testing"
	"k8s.io/apimachinery/pkg/runtime"
	"github.com/openshift/origin/pkg/api/apihelpers/apitesting"
	"github.com/openshift/origin/pkg/route/apis/route"
)

func TestRouteFieldSelectorConversions(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	install := func(scheme *runtime.Scheme) error {
		InstallInternalLegacyRoute(scheme)
		return nil
	}
	apitesting.FieldKeyCheck{SchemeBuilder: []func(*runtime.Scheme) error{install}, Kind: GroupVersion.WithKind("Route"), AllowedExternalFieldKeys: []string{"spec.host", "spec.path", "spec.to.name"}, FieldKeyEvaluatorFn: route.RouteFieldSelector}.Check(t)
}
