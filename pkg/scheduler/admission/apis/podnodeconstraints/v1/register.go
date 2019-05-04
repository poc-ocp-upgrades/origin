package v1

import (
    "github.com/openshift/origin/pkg/scheduler/admission/apis/podnodeconstraints"
    "k8s.io/apimachinery/pkg/runtime"
    "k8s.io/apimachinery/pkg/runtime/schema"
)

func (obj *PodNodeConstraintsConfig) GetObjectKind() schema.ObjectKind {
    _logClusterCodePath()
    defer _logClusterCodePath()
    return &obj.TypeMeta
}

var GroupVersion = schema.GroupVersion{Group: "scheduling.openshift.io", Version: "v1"}
var (
    schemeBuilder = runtime.NewSchemeBuilder(addKnownTypes, podnodeconstraints.Install, addDefaultingFuncs)
    Install       = schemeBuilder.AddToScheme
)

func addKnownTypes(scheme *runtime.Scheme) error {
    _logClusterCodePath()
    defer _logClusterCodePath()
    scheme.AddKnownTypes(GroupVersion, &PodNodeConstraintsConfig{})
    return nil
}
