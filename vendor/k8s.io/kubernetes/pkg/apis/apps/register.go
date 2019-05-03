package apps

import (
 "k8s.io/apimachinery/pkg/runtime"
 "k8s.io/apimachinery/pkg/runtime/schema"
 "k8s.io/kubernetes/pkg/apis/autoscaling"
)

var (
 SchemeBuilder = runtime.NewSchemeBuilder(addKnownTypes)
 AddToScheme   = SchemeBuilder.AddToScheme
)

const GroupName = "apps"

var SchemeGroupVersion = schema.GroupVersion{Group: GroupName, Version: runtime.APIVersionInternal}

func Kind(kind string) schema.GroupKind {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return SchemeGroupVersion.WithKind(kind).GroupKind()
}
func Resource(resource string) schema.GroupResource {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return SchemeGroupVersion.WithResource(resource).GroupResource()
}
func addKnownTypes(scheme *runtime.Scheme) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 scheme.AddKnownTypes(SchemeGroupVersion, &DaemonSet{}, &DaemonSetList{}, &Deployment{}, &DeploymentList{}, &DeploymentRollback{}, &autoscaling.Scale{}, &StatefulSet{}, &StatefulSetList{}, &ControllerRevision{}, &ControllerRevisionList{}, &ReplicaSet{}, &ReplicaSetList{})
 return nil
}
