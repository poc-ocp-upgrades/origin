package legacy

import (
	networkv1 "github.com/openshift/api/network/v1"
	"github.com/openshift/origin/pkg/network/apis/network"
	networkv1helpers "github.com/openshift/origin/pkg/network/apis/network/v1"
	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
)

func InstallInternalLegacyNetwork(scheme *runtime.Scheme) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	InstallExternalLegacyNetwork(scheme)
	schemeBuilder := runtime.NewSchemeBuilder(addUngroupifiedInternalNetworkTypes, networkv1helpers.RegisterDefaults, networkv1helpers.RegisterConversions)
	utilruntime.Must(schemeBuilder.AddToScheme(scheme))
}
func InstallExternalLegacyNetwork(scheme *runtime.Scheme) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	schemeBuilder := runtime.NewSchemeBuilder(addUngroupifiedNetworkTypes)
	utilruntime.Must(schemeBuilder.AddToScheme(scheme))
}
func addUngroupifiedNetworkTypes(scheme *runtime.Scheme) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	types := []runtime.Object{&networkv1.ClusterNetwork{}, &networkv1.ClusterNetworkList{}, &networkv1.HostSubnet{}, &networkv1.HostSubnetList{}, &networkv1.NetNamespace{}, &networkv1.NetNamespaceList{}, &networkv1.EgressNetworkPolicy{}, &networkv1.EgressNetworkPolicyList{}}
	scheme.AddKnownTypes(GroupVersion, types...)
	return nil
}
func addUngroupifiedInternalNetworkTypes(scheme *runtime.Scheme) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	scheme.AddKnownTypes(InternalGroupVersion, &network.ClusterNetwork{}, &network.ClusterNetworkList{}, &network.HostSubnet{}, &network.HostSubnetList{}, &network.NetNamespace{}, &network.NetNamespaceList{}, &network.EgressNetworkPolicy{}, &network.EgressNetworkPolicyList{})
	return nil
}
