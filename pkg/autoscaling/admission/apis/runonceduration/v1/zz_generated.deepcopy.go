package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

func (in *RunOnceDurationConfig) DeepCopyInto(out *RunOnceDurationConfig) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.ActiveDeadlineSecondsOverride != nil {
		in, out := &in.ActiveDeadlineSecondsOverride, &out.ActiveDeadlineSecondsOverride
		*out = new(int64)
		**out = **in
	}
	return
}
func (in *RunOnceDurationConfig) DeepCopy() *RunOnceDurationConfig {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(RunOnceDurationConfig)
	in.DeepCopyInto(out)
	return out
}
func (in *RunOnceDurationConfig) DeepCopyObject() runtime.Object {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
