package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

func (in *ProjectLimitBySelector) DeepCopyInto(out *ProjectLimitBySelector) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	if in.Selector != nil {
		in, out := &in.Selector, &out.Selector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.MaxProjects != nil {
		in, out := &in.MaxProjects, &out.MaxProjects
		*out = new(int)
		**out = **in
	}
	return
}
func (in *ProjectLimitBySelector) DeepCopy() *ProjectLimitBySelector {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(ProjectLimitBySelector)
	in.DeepCopyInto(out)
	return out
}
func (in *ProjectRequestLimitConfig) DeepCopyInto(out *ProjectRequestLimitConfig) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.Limits != nil {
		in, out := &in.Limits, &out.Limits
		*out = make([]ProjectLimitBySelector, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.MaxProjectsForSystemUsers != nil {
		in, out := &in.MaxProjectsForSystemUsers, &out.MaxProjectsForSystemUsers
		*out = new(int)
		**out = **in
	}
	if in.MaxProjectsForServiceAccounts != nil {
		in, out := &in.MaxProjectsForServiceAccounts, &out.MaxProjectsForServiceAccounts
		*out = new(int)
		**out = **in
	}
	return
}
func (in *ProjectRequestLimitConfig) DeepCopy() *ProjectRequestLimitConfig {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(ProjectRequestLimitConfig)
	in.DeepCopyInto(out)
	return out
}
func (in *ProjectRequestLimitConfig) DeepCopyObject() runtime.Object {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
