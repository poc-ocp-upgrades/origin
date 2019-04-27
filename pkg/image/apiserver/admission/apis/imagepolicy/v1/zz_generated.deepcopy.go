package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

func (in *ImageCondition) DeepCopyInto(out *ImageCondition) {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	if in.OnResources != nil {
		in, out := &in.OnResources, &out.OnResources
		*out = make([]metav1.GroupResource, len(*in))
		copy(*out, *in)
	}
	if in.MatchRegistries != nil {
		in, out := &in.MatchRegistries, &out.MatchRegistries
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.MatchDockerImageLabels != nil {
		in, out := &in.MatchDockerImageLabels, &out.MatchDockerImageLabels
		*out = make([]ValueCondition, len(*in))
		copy(*out, *in)
	}
	if in.MatchImageLabels != nil {
		in, out := &in.MatchImageLabels, &out.MatchImageLabels
		*out = make([]metav1.LabelSelector, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.MatchImageLabelSelectors != nil {
		in, out := &in.MatchImageLabelSelectors, &out.MatchImageLabelSelectors
		*out = make([]labels.Selector, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				(*out)[i] = (*in)[i].DeepCopySelector()
			}
		}
	}
	if in.MatchImageAnnotations != nil {
		in, out := &in.MatchImageAnnotations, &out.MatchImageAnnotations
		*out = make([]ValueCondition, len(*in))
		copy(*out, *in)
	}
	return
}
func (in *ImageCondition) DeepCopy() *ImageCondition {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(ImageCondition)
	in.DeepCopyInto(out)
	return out
}
func (in *ImageExecutionPolicyRule) DeepCopyInto(out *ImageExecutionPolicyRule) {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	in.ImageCondition.DeepCopyInto(&out.ImageCondition)
	return
}
func (in *ImageExecutionPolicyRule) DeepCopy() *ImageExecutionPolicyRule {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(ImageExecutionPolicyRule)
	in.DeepCopyInto(out)
	return out
}
func (in *ImagePolicyConfig) DeepCopyInto(out *ImagePolicyConfig) {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.ResolutionRules != nil {
		in, out := &in.ResolutionRules, &out.ResolutionRules
		*out = make([]ImageResolutionPolicyRule, len(*in))
		copy(*out, *in)
	}
	if in.ExecutionRules != nil {
		in, out := &in.ExecutionRules, &out.ExecutionRules
		*out = make([]ImageExecutionPolicyRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}
func (in *ImagePolicyConfig) DeepCopy() *ImagePolicyConfig {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(ImagePolicyConfig)
	in.DeepCopyInto(out)
	return out
}
func (in *ImagePolicyConfig) DeepCopyObject() runtime.Object {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
func (in *ImageResolutionPolicyRule) DeepCopyInto(out *ImageResolutionPolicyRule) {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	out.TargetResource = in.TargetResource
	return
}
func (in *ImageResolutionPolicyRule) DeepCopy() *ImageResolutionPolicyRule {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(ImageResolutionPolicyRule)
	in.DeepCopyInto(out)
	return out
}
func (in *ValueCondition) DeepCopyInto(out *ValueCondition) {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	return
}
func (in *ValueCondition) DeepCopy() *ValueCondition {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(ValueCondition)
	in.DeepCopyInto(out)
	return out
}
