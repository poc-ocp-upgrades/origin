package originpolymorphichelpers

import (
	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/kubernetes/pkg/kubectl"
	"k8s.io/kubernetes/pkg/kubectl/polymorphichelpers"
	appsv1 "github.com/openshift/api/apps/v1"
	deploymentcmd "github.com/openshift/origin/pkg/oc/originpolymorphichelpers/deploymentconfigs"
)

func NewStatusViewerFn(delegate polymorphichelpers.StatusViewerFunc) polymorphichelpers.StatusViewerFunc {
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
	return func(mapping *meta.RESTMapping) (kubectl.StatusViewer, error) {
		if appsv1.SchemeGroupVersion.WithKind("DeploymentConfig").GroupKind() == mapping.GroupVersionKind.GroupKind() {
			return deploymentcmd.NewDeploymentConfigStatusViewer(), nil
		}
		return delegate(mapping)
	}
}
