package network

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

func (in *ClusterNetwork) DeepCopyInto(out *ClusterNetwork) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.ClusterNetworks != nil {
		in, out := &in.ClusterNetworks, &out.ClusterNetworks
		*out = make([]ClusterNetworkEntry, len(*in))
		copy(*out, *in)
	}
	if in.VXLANPort != nil {
		in, out := &in.VXLANPort, &out.VXLANPort
		*out = new(uint32)
		**out = **in
	}
	return
}
func (in *ClusterNetwork) DeepCopy() *ClusterNetwork {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(ClusterNetwork)
	in.DeepCopyInto(out)
	return out
}
func (in *ClusterNetwork) DeepCopyObject() runtime.Object {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
func (in *ClusterNetworkEntry) DeepCopyInto(out *ClusterNetworkEntry) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	return
}
func (in *ClusterNetworkEntry) DeepCopy() *ClusterNetworkEntry {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(ClusterNetworkEntry)
	in.DeepCopyInto(out)
	return out
}
func (in *ClusterNetworkList) DeepCopyInto(out *ClusterNetworkList) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterNetwork, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}
func (in *ClusterNetworkList) DeepCopy() *ClusterNetworkList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(ClusterNetworkList)
	in.DeepCopyInto(out)
	return out
}
func (in *ClusterNetworkList) DeepCopyObject() runtime.Object {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
func (in *EgressNetworkPolicy) DeepCopyInto(out *EgressNetworkPolicy) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}
func (in *EgressNetworkPolicy) DeepCopy() *EgressNetworkPolicy {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(EgressNetworkPolicy)
	in.DeepCopyInto(out)
	return out
}
func (in *EgressNetworkPolicy) DeepCopyObject() runtime.Object {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
func (in *EgressNetworkPolicyList) DeepCopyInto(out *EgressNetworkPolicyList) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]EgressNetworkPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}
func (in *EgressNetworkPolicyList) DeepCopy() *EgressNetworkPolicyList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(EgressNetworkPolicyList)
	in.DeepCopyInto(out)
	return out
}
func (in *EgressNetworkPolicyList) DeepCopyObject() runtime.Object {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
func (in *EgressNetworkPolicyPeer) DeepCopyInto(out *EgressNetworkPolicyPeer) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	return
}
func (in *EgressNetworkPolicyPeer) DeepCopy() *EgressNetworkPolicyPeer {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(EgressNetworkPolicyPeer)
	in.DeepCopyInto(out)
	return out
}
func (in *EgressNetworkPolicyRule) DeepCopyInto(out *EgressNetworkPolicyRule) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	out.To = in.To
	return
}
func (in *EgressNetworkPolicyRule) DeepCopy() *EgressNetworkPolicyRule {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(EgressNetworkPolicyRule)
	in.DeepCopyInto(out)
	return out
}
func (in *EgressNetworkPolicySpec) DeepCopyInto(out *EgressNetworkPolicySpec) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	if in.Egress != nil {
		in, out := &in.Egress, &out.Egress
		*out = make([]EgressNetworkPolicyRule, len(*in))
		copy(*out, *in)
	}
	return
}
func (in *EgressNetworkPolicySpec) DeepCopy() *EgressNetworkPolicySpec {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(EgressNetworkPolicySpec)
	in.DeepCopyInto(out)
	return out
}
func (in *HostSubnet) DeepCopyInto(out *HostSubnet) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.EgressIPs != nil {
		in, out := &in.EgressIPs, &out.EgressIPs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.EgressCIDRs != nil {
		in, out := &in.EgressCIDRs, &out.EgressCIDRs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}
func (in *HostSubnet) DeepCopy() *HostSubnet {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(HostSubnet)
	in.DeepCopyInto(out)
	return out
}
func (in *HostSubnet) DeepCopyObject() runtime.Object {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
func (in *HostSubnetList) DeepCopyInto(out *HostSubnetList) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]HostSubnet, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}
func (in *HostSubnetList) DeepCopy() *HostSubnetList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(HostSubnetList)
	in.DeepCopyInto(out)
	return out
}
func (in *HostSubnetList) DeepCopyObject() runtime.Object {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
func (in *NetNamespace) DeepCopyInto(out *NetNamespace) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.EgressIPs != nil {
		in, out := &in.EgressIPs, &out.EgressIPs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}
func (in *NetNamespace) DeepCopy() *NetNamespace {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(NetNamespace)
	in.DeepCopyInto(out)
	return out
}
func (in *NetNamespace) DeepCopyObject() runtime.Object {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
func (in *NetNamespaceList) DeepCopyInto(out *NetNamespaceList) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]NetNamespace, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}
func (in *NetNamespaceList) DeepCopy() *NetNamespaceList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(NetNamespaceList)
	in.DeepCopyInto(out)
	return out
}
func (in *NetNamespaceList) DeepCopyObject() runtime.Object {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
