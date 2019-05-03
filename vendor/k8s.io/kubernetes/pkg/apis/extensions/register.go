package extensions

import (
 "k8s.io/apimachinery/pkg/runtime"
 "k8s.io/apimachinery/pkg/runtime/schema"
 "k8s.io/kubernetes/pkg/apis/apps"
 "k8s.io/kubernetes/pkg/apis/autoscaling"
 "k8s.io/kubernetes/pkg/apis/networking"
 "k8s.io/kubernetes/pkg/apis/policy"
)

const GroupName = "extensions"

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

var (
 SchemeBuilder = runtime.NewSchemeBuilder(addKnownTypes)
 AddToScheme   = SchemeBuilder.AddToScheme
)

func addKnownTypes(scheme *runtime.Scheme) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 scheme.AddKnownTypes(SchemeGroupVersion, &apps.Deployment{}, &apps.DeploymentList{}, &apps.DeploymentRollback{}, &ReplicationControllerDummy{}, &apps.DaemonSetList{}, &apps.DaemonSet{}, &Ingress{}, &IngressList{}, &apps.ReplicaSet{}, &apps.ReplicaSetList{}, &policy.PodSecurityPolicy{}, &policy.PodSecurityPolicyList{}, &autoscaling.Scale{}, &networking.NetworkPolicy{}, &networking.NetworkPolicyList{})
 return nil
}
