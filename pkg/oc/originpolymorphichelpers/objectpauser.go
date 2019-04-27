package originpolymorphichelpers

import (
	"errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/kubernetes/pkg/kubectl/polymorphichelpers"
	"k8s.io/kubernetes/pkg/kubectl/scheme"
	appsv1 "github.com/openshift/api/apps/v1"
)

func NewObjectPauserFn(delegate polymorphichelpers.ObjectPauserFunc) polymorphichelpers.ObjectPauserFunc {
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
	return func(obj runtime.Object) ([]byte, error) {
		switch t := obj.(type) {
		case *appsv1.DeploymentConfig:
			if t.Spec.Paused {
				return nil, errors.New("is already paused")
			}
			t.Spec.Paused = true
			return runtime.Encode(scheme.DefaultJSONEncoder(), obj)
		default:
			return delegate(obj)
		}
	}
}
