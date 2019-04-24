package restrictedendpoints

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

func (in *RestrictedEndpointsAdmissionConfig) DeepCopyInto(out *RestrictedEndpointsAdmissionConfig) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.RestrictedCIDRs != nil {
		in, out := &in.RestrictedCIDRs, &out.RestrictedCIDRs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}
func (in *RestrictedEndpointsAdmissionConfig) DeepCopy() *RestrictedEndpointsAdmissionConfig {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(RestrictedEndpointsAdmissionConfig)
	in.DeepCopyInto(out)
	return out
}
func (in *RestrictedEndpointsAdmissionConfig) DeepCopyObject() runtime.Object {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
