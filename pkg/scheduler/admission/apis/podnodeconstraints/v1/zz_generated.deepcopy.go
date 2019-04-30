package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

func (in *PodNodeConstraintsConfig) DeepCopyInto(out *PodNodeConstraintsConfig) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.NodeSelectorLabelBlacklist != nil {
		in, out := &in.NodeSelectorLabelBlacklist, &out.NodeSelectorLabelBlacklist
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}
func (in *PodNodeConstraintsConfig) DeepCopy() *PodNodeConstraintsConfig {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(PodNodeConstraintsConfig)
	in.DeepCopyInto(out)
	return out
}
func (in *PodNodeConstraintsConfig) DeepCopyObject() runtime.Object {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
