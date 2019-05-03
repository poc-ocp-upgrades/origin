package v1

import (
	"github.com/openshift/origin/pkg/autoscaling/admission/apis/clusterresourceoverride"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func (obj *ClusterResourceOverrideConfig) GetObjectKind() schema.ObjectKind {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &obj.TypeMeta
}

var GroupVersion = schema.GroupVersion{Group: "autoscaling.openshift.io", Version: "v1"}
var (
	schemeBuilder = runtime.NewSchemeBuilder(addKnownTypes, clusterresourceoverride.Install)
	Install       = schemeBuilder.AddToScheme
)

func addKnownTypes(scheme *runtime.Scheme) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	scheme.AddKnownTypes(GroupVersion, &ClusterResourceOverrideConfig{})
	return nil
}
