package testing

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

func (in *AdmissionPluginTestConfig) DeepCopyInto(out *AdmissionPluginTestConfig) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	out.TypeMeta = in.TypeMeta
	return
}
func (in *AdmissionPluginTestConfig) DeepCopy() *AdmissionPluginTestConfig {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(AdmissionPluginTestConfig)
	in.DeepCopyInto(out)
	return out
}
func (in *AdmissionPluginTestConfig) DeepCopyObject() runtime.Object {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
