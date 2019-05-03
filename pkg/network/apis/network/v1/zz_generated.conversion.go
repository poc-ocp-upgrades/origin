package v1

import (
	v1 "github.com/openshift/api/network/v1"
	network "github.com/openshift/origin/pkg/network/apis/network"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	unsafe "unsafe"
)

func init() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	localSchemeBuilder.Register(RegisterConversions)
}
func RegisterConversions(s *runtime.Scheme) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := s.AddGeneratedConversionFunc((*v1.ClusterNetwork)(nil), (*network.ClusterNetwork)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_ClusterNetwork_To_network_ClusterNetwork(a.(*v1.ClusterNetwork), b.(*network.ClusterNetwork), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*network.ClusterNetwork)(nil), (*v1.ClusterNetwork)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_network_ClusterNetwork_To_v1_ClusterNetwork(a.(*network.ClusterNetwork), b.(*v1.ClusterNetwork), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.ClusterNetworkEntry)(nil), (*network.ClusterNetworkEntry)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_ClusterNetworkEntry_To_network_ClusterNetworkEntry(a.(*v1.ClusterNetworkEntry), b.(*network.ClusterNetworkEntry), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*network.ClusterNetworkEntry)(nil), (*v1.ClusterNetworkEntry)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_network_ClusterNetworkEntry_To_v1_ClusterNetworkEntry(a.(*network.ClusterNetworkEntry), b.(*v1.ClusterNetworkEntry), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.ClusterNetworkList)(nil), (*network.ClusterNetworkList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_ClusterNetworkList_To_network_ClusterNetworkList(a.(*v1.ClusterNetworkList), b.(*network.ClusterNetworkList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*network.ClusterNetworkList)(nil), (*v1.ClusterNetworkList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_network_ClusterNetworkList_To_v1_ClusterNetworkList(a.(*network.ClusterNetworkList), b.(*v1.ClusterNetworkList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.EgressNetworkPolicy)(nil), (*network.EgressNetworkPolicy)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_EgressNetworkPolicy_To_network_EgressNetworkPolicy(a.(*v1.EgressNetworkPolicy), b.(*network.EgressNetworkPolicy), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*network.EgressNetworkPolicy)(nil), (*v1.EgressNetworkPolicy)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_network_EgressNetworkPolicy_To_v1_EgressNetworkPolicy(a.(*network.EgressNetworkPolicy), b.(*v1.EgressNetworkPolicy), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.EgressNetworkPolicyList)(nil), (*network.EgressNetworkPolicyList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_EgressNetworkPolicyList_To_network_EgressNetworkPolicyList(a.(*v1.EgressNetworkPolicyList), b.(*network.EgressNetworkPolicyList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*network.EgressNetworkPolicyList)(nil), (*v1.EgressNetworkPolicyList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_network_EgressNetworkPolicyList_To_v1_EgressNetworkPolicyList(a.(*network.EgressNetworkPolicyList), b.(*v1.EgressNetworkPolicyList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.EgressNetworkPolicyPeer)(nil), (*network.EgressNetworkPolicyPeer)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_EgressNetworkPolicyPeer_To_network_EgressNetworkPolicyPeer(a.(*v1.EgressNetworkPolicyPeer), b.(*network.EgressNetworkPolicyPeer), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*network.EgressNetworkPolicyPeer)(nil), (*v1.EgressNetworkPolicyPeer)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_network_EgressNetworkPolicyPeer_To_v1_EgressNetworkPolicyPeer(a.(*network.EgressNetworkPolicyPeer), b.(*v1.EgressNetworkPolicyPeer), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.EgressNetworkPolicyRule)(nil), (*network.EgressNetworkPolicyRule)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_EgressNetworkPolicyRule_To_network_EgressNetworkPolicyRule(a.(*v1.EgressNetworkPolicyRule), b.(*network.EgressNetworkPolicyRule), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*network.EgressNetworkPolicyRule)(nil), (*v1.EgressNetworkPolicyRule)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_network_EgressNetworkPolicyRule_To_v1_EgressNetworkPolicyRule(a.(*network.EgressNetworkPolicyRule), b.(*v1.EgressNetworkPolicyRule), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.EgressNetworkPolicySpec)(nil), (*network.EgressNetworkPolicySpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_EgressNetworkPolicySpec_To_network_EgressNetworkPolicySpec(a.(*v1.EgressNetworkPolicySpec), b.(*network.EgressNetworkPolicySpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*network.EgressNetworkPolicySpec)(nil), (*v1.EgressNetworkPolicySpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_network_EgressNetworkPolicySpec_To_v1_EgressNetworkPolicySpec(a.(*network.EgressNetworkPolicySpec), b.(*v1.EgressNetworkPolicySpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.HostSubnet)(nil), (*network.HostSubnet)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_HostSubnet_To_network_HostSubnet(a.(*v1.HostSubnet), b.(*network.HostSubnet), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*network.HostSubnet)(nil), (*v1.HostSubnet)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_network_HostSubnet_To_v1_HostSubnet(a.(*network.HostSubnet), b.(*v1.HostSubnet), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.HostSubnetList)(nil), (*network.HostSubnetList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_HostSubnetList_To_network_HostSubnetList(a.(*v1.HostSubnetList), b.(*network.HostSubnetList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*network.HostSubnetList)(nil), (*v1.HostSubnetList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_network_HostSubnetList_To_v1_HostSubnetList(a.(*network.HostSubnetList), b.(*v1.HostSubnetList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.NetNamespace)(nil), (*network.NetNamespace)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_NetNamespace_To_network_NetNamespace(a.(*v1.NetNamespace), b.(*network.NetNamespace), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*network.NetNamespace)(nil), (*v1.NetNamespace)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_network_NetNamespace_To_v1_NetNamespace(a.(*network.NetNamespace), b.(*v1.NetNamespace), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.NetNamespaceList)(nil), (*network.NetNamespaceList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_NetNamespaceList_To_network_NetNamespaceList(a.(*v1.NetNamespaceList), b.(*network.NetNamespaceList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*network.NetNamespaceList)(nil), (*v1.NetNamespaceList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_network_NetNamespaceList_To_v1_NetNamespaceList(a.(*network.NetNamespaceList), b.(*v1.NetNamespaceList), scope)
	}); err != nil {
		return err
	}
	return nil
}
func autoConvert_v1_ClusterNetwork_To_network_ClusterNetwork(in *v1.ClusterNetwork, out *network.ClusterNetwork, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ObjectMeta = in.ObjectMeta
	out.Network = in.Network
	out.HostSubnetLength = in.HostSubnetLength
	out.ServiceNetwork = in.ServiceNetwork
	out.PluginName = in.PluginName
	out.ClusterNetworks = *(*[]network.ClusterNetworkEntry)(unsafe.Pointer(&in.ClusterNetworks))
	out.VXLANPort = (*uint32)(unsafe.Pointer(in.VXLANPort))
	return nil
}
func Convert_v1_ClusterNetwork_To_network_ClusterNetwork(in *v1.ClusterNetwork, out *network.ClusterNetwork, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_ClusterNetwork_To_network_ClusterNetwork(in, out, s)
}
func autoConvert_network_ClusterNetwork_To_v1_ClusterNetwork(in *network.ClusterNetwork, out *v1.ClusterNetwork, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ObjectMeta = in.ObjectMeta
	out.ClusterNetworks = *(*[]v1.ClusterNetworkEntry)(unsafe.Pointer(&in.ClusterNetworks))
	out.Network = in.Network
	out.HostSubnetLength = in.HostSubnetLength
	out.ServiceNetwork = in.ServiceNetwork
	out.PluginName = in.PluginName
	out.VXLANPort = (*uint32)(unsafe.Pointer(in.VXLANPort))
	return nil
}
func Convert_network_ClusterNetwork_To_v1_ClusterNetwork(in *network.ClusterNetwork, out *v1.ClusterNetwork, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_network_ClusterNetwork_To_v1_ClusterNetwork(in, out, s)
}
func autoConvert_v1_ClusterNetworkEntry_To_network_ClusterNetworkEntry(in *v1.ClusterNetworkEntry, out *network.ClusterNetworkEntry, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.CIDR = in.CIDR
	out.HostSubnetLength = in.HostSubnetLength
	return nil
}
func Convert_v1_ClusterNetworkEntry_To_network_ClusterNetworkEntry(in *v1.ClusterNetworkEntry, out *network.ClusterNetworkEntry, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_ClusterNetworkEntry_To_network_ClusterNetworkEntry(in, out, s)
}
func autoConvert_network_ClusterNetworkEntry_To_v1_ClusterNetworkEntry(in *network.ClusterNetworkEntry, out *v1.ClusterNetworkEntry, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.CIDR = in.CIDR
	out.HostSubnetLength = in.HostSubnetLength
	return nil
}
func Convert_network_ClusterNetworkEntry_To_v1_ClusterNetworkEntry(in *network.ClusterNetworkEntry, out *v1.ClusterNetworkEntry, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_network_ClusterNetworkEntry_To_v1_ClusterNetworkEntry(in, out, s)
}
func autoConvert_v1_ClusterNetworkList_To_network_ClusterNetworkList(in *v1.ClusterNetworkList, out *network.ClusterNetworkList, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]network.ClusterNetwork, len(*in))
		for i := range *in {
			if err := Convert_v1_ClusterNetwork_To_network_ClusterNetwork(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}
func Convert_v1_ClusterNetworkList_To_network_ClusterNetworkList(in *v1.ClusterNetworkList, out *network.ClusterNetworkList, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_ClusterNetworkList_To_network_ClusterNetworkList(in, out, s)
}
func autoConvert_network_ClusterNetworkList_To_v1_ClusterNetworkList(in *network.ClusterNetworkList, out *v1.ClusterNetworkList, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]v1.ClusterNetwork, len(*in))
		for i := range *in {
			if err := Convert_network_ClusterNetwork_To_v1_ClusterNetwork(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}
func Convert_network_ClusterNetworkList_To_v1_ClusterNetworkList(in *network.ClusterNetworkList, out *v1.ClusterNetworkList, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_network_ClusterNetworkList_To_v1_ClusterNetworkList(in, out, s)
}
func autoConvert_v1_EgressNetworkPolicy_To_network_EgressNetworkPolicy(in *v1.EgressNetworkPolicy, out *network.EgressNetworkPolicy, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1_EgressNetworkPolicySpec_To_network_EgressNetworkPolicySpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	return nil
}
func Convert_v1_EgressNetworkPolicy_To_network_EgressNetworkPolicy(in *v1.EgressNetworkPolicy, out *network.EgressNetworkPolicy, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_EgressNetworkPolicy_To_network_EgressNetworkPolicy(in, out, s)
}
func autoConvert_network_EgressNetworkPolicy_To_v1_EgressNetworkPolicy(in *network.EgressNetworkPolicy, out *v1.EgressNetworkPolicy, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_network_EgressNetworkPolicySpec_To_v1_EgressNetworkPolicySpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	return nil
}
func Convert_network_EgressNetworkPolicy_To_v1_EgressNetworkPolicy(in *network.EgressNetworkPolicy, out *v1.EgressNetworkPolicy, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_network_EgressNetworkPolicy_To_v1_EgressNetworkPolicy(in, out, s)
}
func autoConvert_v1_EgressNetworkPolicyList_To_network_EgressNetworkPolicyList(in *v1.EgressNetworkPolicyList, out *network.EgressNetworkPolicyList, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ListMeta = in.ListMeta
	out.Items = *(*[]network.EgressNetworkPolicy)(unsafe.Pointer(&in.Items))
	return nil
}
func Convert_v1_EgressNetworkPolicyList_To_network_EgressNetworkPolicyList(in *v1.EgressNetworkPolicyList, out *network.EgressNetworkPolicyList, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_EgressNetworkPolicyList_To_network_EgressNetworkPolicyList(in, out, s)
}
func autoConvert_network_EgressNetworkPolicyList_To_v1_EgressNetworkPolicyList(in *network.EgressNetworkPolicyList, out *v1.EgressNetworkPolicyList, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ListMeta = in.ListMeta
	out.Items = *(*[]v1.EgressNetworkPolicy)(unsafe.Pointer(&in.Items))
	return nil
}
func Convert_network_EgressNetworkPolicyList_To_v1_EgressNetworkPolicyList(in *network.EgressNetworkPolicyList, out *v1.EgressNetworkPolicyList, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_network_EgressNetworkPolicyList_To_v1_EgressNetworkPolicyList(in, out, s)
}
func autoConvert_v1_EgressNetworkPolicyPeer_To_network_EgressNetworkPolicyPeer(in *v1.EgressNetworkPolicyPeer, out *network.EgressNetworkPolicyPeer, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.CIDRSelector = in.CIDRSelector
	out.DNSName = in.DNSName
	return nil
}
func Convert_v1_EgressNetworkPolicyPeer_To_network_EgressNetworkPolicyPeer(in *v1.EgressNetworkPolicyPeer, out *network.EgressNetworkPolicyPeer, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_EgressNetworkPolicyPeer_To_network_EgressNetworkPolicyPeer(in, out, s)
}
func autoConvert_network_EgressNetworkPolicyPeer_To_v1_EgressNetworkPolicyPeer(in *network.EgressNetworkPolicyPeer, out *v1.EgressNetworkPolicyPeer, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.CIDRSelector = in.CIDRSelector
	out.DNSName = in.DNSName
	return nil
}
func Convert_network_EgressNetworkPolicyPeer_To_v1_EgressNetworkPolicyPeer(in *network.EgressNetworkPolicyPeer, out *v1.EgressNetworkPolicyPeer, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_network_EgressNetworkPolicyPeer_To_v1_EgressNetworkPolicyPeer(in, out, s)
}
func autoConvert_v1_EgressNetworkPolicyRule_To_network_EgressNetworkPolicyRule(in *v1.EgressNetworkPolicyRule, out *network.EgressNetworkPolicyRule, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Type = network.EgressNetworkPolicyRuleType(in.Type)
	if err := Convert_v1_EgressNetworkPolicyPeer_To_network_EgressNetworkPolicyPeer(&in.To, &out.To, s); err != nil {
		return err
	}
	return nil
}
func Convert_v1_EgressNetworkPolicyRule_To_network_EgressNetworkPolicyRule(in *v1.EgressNetworkPolicyRule, out *network.EgressNetworkPolicyRule, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_EgressNetworkPolicyRule_To_network_EgressNetworkPolicyRule(in, out, s)
}
func autoConvert_network_EgressNetworkPolicyRule_To_v1_EgressNetworkPolicyRule(in *network.EgressNetworkPolicyRule, out *v1.EgressNetworkPolicyRule, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Type = v1.EgressNetworkPolicyRuleType(in.Type)
	if err := Convert_network_EgressNetworkPolicyPeer_To_v1_EgressNetworkPolicyPeer(&in.To, &out.To, s); err != nil {
		return err
	}
	return nil
}
func Convert_network_EgressNetworkPolicyRule_To_v1_EgressNetworkPolicyRule(in *network.EgressNetworkPolicyRule, out *v1.EgressNetworkPolicyRule, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_network_EgressNetworkPolicyRule_To_v1_EgressNetworkPolicyRule(in, out, s)
}
func autoConvert_v1_EgressNetworkPolicySpec_To_network_EgressNetworkPolicySpec(in *v1.EgressNetworkPolicySpec, out *network.EgressNetworkPolicySpec, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Egress = *(*[]network.EgressNetworkPolicyRule)(unsafe.Pointer(&in.Egress))
	return nil
}
func Convert_v1_EgressNetworkPolicySpec_To_network_EgressNetworkPolicySpec(in *v1.EgressNetworkPolicySpec, out *network.EgressNetworkPolicySpec, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_EgressNetworkPolicySpec_To_network_EgressNetworkPolicySpec(in, out, s)
}
func autoConvert_network_EgressNetworkPolicySpec_To_v1_EgressNetworkPolicySpec(in *network.EgressNetworkPolicySpec, out *v1.EgressNetworkPolicySpec, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Egress = *(*[]v1.EgressNetworkPolicyRule)(unsafe.Pointer(&in.Egress))
	return nil
}
func Convert_network_EgressNetworkPolicySpec_To_v1_EgressNetworkPolicySpec(in *network.EgressNetworkPolicySpec, out *v1.EgressNetworkPolicySpec, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_network_EgressNetworkPolicySpec_To_v1_EgressNetworkPolicySpec(in, out, s)
}
func autoConvert_v1_HostSubnet_To_network_HostSubnet(in *v1.HostSubnet, out *network.HostSubnet, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ObjectMeta = in.ObjectMeta
	out.Host = in.Host
	out.HostIP = in.HostIP
	out.Subnet = in.Subnet
	out.EgressIPs = *(*[]string)(unsafe.Pointer(&in.EgressIPs))
	out.EgressCIDRs = *(*[]string)(unsafe.Pointer(&in.EgressCIDRs))
	return nil
}
func Convert_v1_HostSubnet_To_network_HostSubnet(in *v1.HostSubnet, out *network.HostSubnet, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_HostSubnet_To_network_HostSubnet(in, out, s)
}
func autoConvert_network_HostSubnet_To_v1_HostSubnet(in *network.HostSubnet, out *v1.HostSubnet, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ObjectMeta = in.ObjectMeta
	out.Host = in.Host
	out.HostIP = in.HostIP
	out.Subnet = in.Subnet
	out.EgressIPs = *(*[]string)(unsafe.Pointer(&in.EgressIPs))
	out.EgressCIDRs = *(*[]string)(unsafe.Pointer(&in.EgressCIDRs))
	return nil
}
func Convert_network_HostSubnet_To_v1_HostSubnet(in *network.HostSubnet, out *v1.HostSubnet, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_network_HostSubnet_To_v1_HostSubnet(in, out, s)
}
func autoConvert_v1_HostSubnetList_To_network_HostSubnetList(in *v1.HostSubnetList, out *network.HostSubnetList, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ListMeta = in.ListMeta
	out.Items = *(*[]network.HostSubnet)(unsafe.Pointer(&in.Items))
	return nil
}
func Convert_v1_HostSubnetList_To_network_HostSubnetList(in *v1.HostSubnetList, out *network.HostSubnetList, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_HostSubnetList_To_network_HostSubnetList(in, out, s)
}
func autoConvert_network_HostSubnetList_To_v1_HostSubnetList(in *network.HostSubnetList, out *v1.HostSubnetList, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ListMeta = in.ListMeta
	out.Items = *(*[]v1.HostSubnet)(unsafe.Pointer(&in.Items))
	return nil
}
func Convert_network_HostSubnetList_To_v1_HostSubnetList(in *network.HostSubnetList, out *v1.HostSubnetList, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_network_HostSubnetList_To_v1_HostSubnetList(in, out, s)
}
func autoConvert_v1_NetNamespace_To_network_NetNamespace(in *v1.NetNamespace, out *network.NetNamespace, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ObjectMeta = in.ObjectMeta
	out.NetName = in.NetName
	out.NetID = in.NetID
	out.EgressIPs = *(*[]string)(unsafe.Pointer(&in.EgressIPs))
	return nil
}
func Convert_v1_NetNamespace_To_network_NetNamespace(in *v1.NetNamespace, out *network.NetNamespace, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_NetNamespace_To_network_NetNamespace(in, out, s)
}
func autoConvert_network_NetNamespace_To_v1_NetNamespace(in *network.NetNamespace, out *v1.NetNamespace, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ObjectMeta = in.ObjectMeta
	out.NetName = in.NetName
	out.NetID = in.NetID
	out.EgressIPs = *(*[]string)(unsafe.Pointer(&in.EgressIPs))
	return nil
}
func Convert_network_NetNamespace_To_v1_NetNamespace(in *network.NetNamespace, out *v1.NetNamespace, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_network_NetNamespace_To_v1_NetNamespace(in, out, s)
}
func autoConvert_v1_NetNamespaceList_To_network_NetNamespaceList(in *v1.NetNamespaceList, out *network.NetNamespaceList, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ListMeta = in.ListMeta
	out.Items = *(*[]network.NetNamespace)(unsafe.Pointer(&in.Items))
	return nil
}
func Convert_v1_NetNamespaceList_To_network_NetNamespaceList(in *v1.NetNamespaceList, out *network.NetNamespaceList, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_NetNamespaceList_To_network_NetNamespaceList(in, out, s)
}
func autoConvert_network_NetNamespaceList_To_v1_NetNamespaceList(in *network.NetNamespaceList, out *v1.NetNamespaceList, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ListMeta = in.ListMeta
	out.Items = *(*[]v1.NetNamespace)(unsafe.Pointer(&in.Items))
	return nil
}
func Convert_network_NetNamespaceList_To_v1_NetNamespaceList(in *network.NetNamespaceList, out *v1.NetNamespaceList, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_network_NetNamespaceList_To_v1_NetNamespaceList(in, out, s)
}
