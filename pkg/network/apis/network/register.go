package network

import (
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

const (
	GroupName = "network.openshift.io"
)

var (
	schemeBuilder      = runtime.NewSchemeBuilder(addKnownTypes)
	Install            = schemeBuilder.AddToScheme
	SchemeGroupVersion = schema.GroupVersion{Group: GroupName, Version: runtime.APIVersionInternal}
	AddToScheme        = schemeBuilder.AddToScheme
)

func Resource(resource string) schema.GroupResource {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return SchemeGroupVersion.WithResource(resource).GroupResource()
}
func addKnownTypes(scheme *runtime.Scheme) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	scheme.AddKnownTypes(SchemeGroupVersion, &ClusterNetwork{}, &ClusterNetworkList{}, &HostSubnet{}, &HostSubnetList{}, &NetNamespace{}, &NetNamespaceList{}, &EgressNetworkPolicy{}, &EgressNetworkPolicyList{})
	return nil
}
