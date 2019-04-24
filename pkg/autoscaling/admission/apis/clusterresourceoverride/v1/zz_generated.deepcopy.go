package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

func (in *ClusterResourceOverrideConfig) DeepCopyInto(out *ClusterResourceOverrideConfig) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	out.TypeMeta = in.TypeMeta
	return
}
func (in *ClusterResourceOverrideConfig) DeepCopy() *ClusterResourceOverrideConfig {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(ClusterResourceOverrideConfig)
	in.DeepCopyInto(out)
	return out
}
func (in *ClusterResourceOverrideConfig) DeepCopyObject() runtime.Object {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
