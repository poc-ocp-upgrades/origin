package originpolymorphichelpers

import (
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/kubernetes/pkg/kubectl/generate"
	"k8s.io/kubernetes/pkg/kubectl/polymorphichelpers"
	appsv1 "github.com/openshift/api/apps/v1"
)

func NewMapBasedSelectorForObjectFn(delegate polymorphichelpers.MapBasedSelectorForObjectFunc) polymorphichelpers.MapBasedSelectorForObjectFunc {
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
	return func(object runtime.Object) (string, error) {
		switch t := object.(type) {
		case *appsv1.DeploymentConfig:
			return generate.MakeLabels(t.Spec.Selector), nil
		default:
			return delegate(object)
		}
	}
}
