package apps

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
	core "k8s.io/kubernetes/pkg/apis/core"
)

func (in *CustomDeploymentStrategyParams) DeepCopyInto(out *CustomDeploymentStrategyParams) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	if in.Environment != nil {
		in, out := &in.Environment, &out.Environment
		*out = make([]core.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Command != nil {
		in, out := &in.Command, &out.Command
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}
func (in *CustomDeploymentStrategyParams) DeepCopy() *CustomDeploymentStrategyParams {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(CustomDeploymentStrategyParams)
	in.DeepCopyInto(out)
	return out
}
func (in *DeploymentCause) DeepCopyInto(out *DeploymentCause) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	if in.ImageTrigger != nil {
		in, out := &in.ImageTrigger, &out.ImageTrigger
		*out = new(DeploymentCauseImageTrigger)
		**out = **in
	}
	return
}
func (in *DeploymentCause) DeepCopy() *DeploymentCause {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(DeploymentCause)
	in.DeepCopyInto(out)
	return out
}
func (in *DeploymentCauseImageTrigger) DeepCopyInto(out *DeploymentCauseImageTrigger) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	out.From = in.From
	return
}
func (in *DeploymentCauseImageTrigger) DeepCopy() *DeploymentCauseImageTrigger {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(DeploymentCauseImageTrigger)
	in.DeepCopyInto(out)
	return out
}
func (in *DeploymentCondition) DeepCopyInto(out *DeploymentCondition) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	in.LastUpdateTime.DeepCopyInto(&out.LastUpdateTime)
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}
func (in *DeploymentCondition) DeepCopy() *DeploymentCondition {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(DeploymentCondition)
	in.DeepCopyInto(out)
	return out
}
func (in *DeploymentConfig) DeepCopyInto(out *DeploymentConfig) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}
func (in *DeploymentConfig) DeepCopy() *DeploymentConfig {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(DeploymentConfig)
	in.DeepCopyInto(out)
	return out
}
func (in *DeploymentConfig) DeepCopyObject() runtime.Object {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
func (in *DeploymentConfigList) DeepCopyInto(out *DeploymentConfigList) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DeploymentConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}
func (in *DeploymentConfigList) DeepCopy() *DeploymentConfigList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(DeploymentConfigList)
	in.DeepCopyInto(out)
	return out
}
func (in *DeploymentConfigList) DeepCopyObject() runtime.Object {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
func (in *DeploymentConfigRollback) DeepCopyInto(out *DeploymentConfigRollback) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.UpdatedAnnotations != nil {
		in, out := &in.UpdatedAnnotations, &out.UpdatedAnnotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	out.Spec = in.Spec
	return
}
func (in *DeploymentConfigRollback) DeepCopy() *DeploymentConfigRollback {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(DeploymentConfigRollback)
	in.DeepCopyInto(out)
	return out
}
func (in *DeploymentConfigRollback) DeepCopyObject() runtime.Object {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
func (in *DeploymentConfigRollbackSpec) DeepCopyInto(out *DeploymentConfigRollbackSpec) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	out.From = in.From
	return
}
func (in *DeploymentConfigRollbackSpec) DeepCopy() *DeploymentConfigRollbackSpec {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(DeploymentConfigRollbackSpec)
	in.DeepCopyInto(out)
	return out
}
func (in *DeploymentConfigSpec) DeepCopyInto(out *DeploymentConfigSpec) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	in.Strategy.DeepCopyInto(&out.Strategy)
	if in.Triggers != nil {
		in, out := &in.Triggers, &out.Triggers
		*out = make([]DeploymentTriggerPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.RevisionHistoryLimit != nil {
		in, out := &in.RevisionHistoryLimit, &out.RevisionHistoryLimit
		*out = new(int32)
		**out = **in
	}
	if in.Selector != nil {
		in, out := &in.Selector, &out.Selector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Template != nil {
		in, out := &in.Template, &out.Template
		*out = new(core.PodTemplateSpec)
		(*in).DeepCopyInto(*out)
	}
	return
}
func (in *DeploymentConfigSpec) DeepCopy() *DeploymentConfigSpec {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(DeploymentConfigSpec)
	in.DeepCopyInto(out)
	return out
}
func (in *DeploymentConfigStatus) DeepCopyInto(out *DeploymentConfigStatus) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	if in.Details != nil {
		in, out := &in.Details, &out.Details
		*out = new(DeploymentDetails)
		(*in).DeepCopyInto(*out)
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]DeploymentCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}
func (in *DeploymentConfigStatus) DeepCopy() *DeploymentConfigStatus {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(DeploymentConfigStatus)
	in.DeepCopyInto(out)
	return out
}
func (in *DeploymentDetails) DeepCopyInto(out *DeploymentDetails) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	if in.Causes != nil {
		in, out := &in.Causes, &out.Causes
		*out = make([]DeploymentCause, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}
func (in *DeploymentDetails) DeepCopy() *DeploymentDetails {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(DeploymentDetails)
	in.DeepCopyInto(out)
	return out
}
func (in *DeploymentLog) DeepCopyInto(out *DeploymentLog) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	out.TypeMeta = in.TypeMeta
	return
}
func (in *DeploymentLog) DeepCopy() *DeploymentLog {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(DeploymentLog)
	in.DeepCopyInto(out)
	return out
}
func (in *DeploymentLog) DeepCopyObject() runtime.Object {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
func (in *DeploymentLogOptions) DeepCopyInto(out *DeploymentLogOptions) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.SinceSeconds != nil {
		in, out := &in.SinceSeconds, &out.SinceSeconds
		*out = new(int64)
		**out = **in
	}
	if in.SinceTime != nil {
		in, out := &in.SinceTime, &out.SinceTime
		*out = (*in).DeepCopy()
	}
	if in.TailLines != nil {
		in, out := &in.TailLines, &out.TailLines
		*out = new(int64)
		**out = **in
	}
	if in.LimitBytes != nil {
		in, out := &in.LimitBytes, &out.LimitBytes
		*out = new(int64)
		**out = **in
	}
	if in.Version != nil {
		in, out := &in.Version, &out.Version
		*out = new(int64)
		**out = **in
	}
	return
}
func (in *DeploymentLogOptions) DeepCopy() *DeploymentLogOptions {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(DeploymentLogOptions)
	in.DeepCopyInto(out)
	return out
}
func (in *DeploymentLogOptions) DeepCopyObject() runtime.Object {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
func (in *DeploymentRequest) DeepCopyInto(out *DeploymentRequest) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.ExcludeTriggers != nil {
		in, out := &in.ExcludeTriggers, &out.ExcludeTriggers
		*out = make([]DeploymentTriggerType, len(*in))
		copy(*out, *in)
	}
	return
}
func (in *DeploymentRequest) DeepCopy() *DeploymentRequest {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(DeploymentRequest)
	in.DeepCopyInto(out)
	return out
}
func (in *DeploymentRequest) DeepCopyObject() runtime.Object {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
func (in *DeploymentStrategy) DeepCopyInto(out *DeploymentStrategy) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	if in.CustomParams != nil {
		in, out := &in.CustomParams, &out.CustomParams
		*out = new(CustomDeploymentStrategyParams)
		(*in).DeepCopyInto(*out)
	}
	if in.RecreateParams != nil {
		in, out := &in.RecreateParams, &out.RecreateParams
		*out = new(RecreateDeploymentStrategyParams)
		(*in).DeepCopyInto(*out)
	}
	if in.RollingParams != nil {
		in, out := &in.RollingParams, &out.RollingParams
		*out = new(RollingDeploymentStrategyParams)
		(*in).DeepCopyInto(*out)
	}
	in.Resources.DeepCopyInto(&out.Resources)
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.ActiveDeadlineSeconds != nil {
		in, out := &in.ActiveDeadlineSeconds, &out.ActiveDeadlineSeconds
		*out = new(int64)
		**out = **in
	}
	return
}
func (in *DeploymentStrategy) DeepCopy() *DeploymentStrategy {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(DeploymentStrategy)
	in.DeepCopyInto(out)
	return out
}
func (in *DeploymentTriggerImageChangeParams) DeepCopyInto(out *DeploymentTriggerImageChangeParams) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	if in.ContainerNames != nil {
		in, out := &in.ContainerNames, &out.ContainerNames
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	out.From = in.From
	return
}
func (in *DeploymentTriggerImageChangeParams) DeepCopy() *DeploymentTriggerImageChangeParams {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(DeploymentTriggerImageChangeParams)
	in.DeepCopyInto(out)
	return out
}
func (in *DeploymentTriggerPolicy) DeepCopyInto(out *DeploymentTriggerPolicy) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	if in.ImageChangeParams != nil {
		in, out := &in.ImageChangeParams, &out.ImageChangeParams
		*out = new(DeploymentTriggerImageChangeParams)
		(*in).DeepCopyInto(*out)
	}
	return
}
func (in *DeploymentTriggerPolicy) DeepCopy() *DeploymentTriggerPolicy {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(DeploymentTriggerPolicy)
	in.DeepCopyInto(out)
	return out
}
func (in *ExecNewPodHook) DeepCopyInto(out *ExecNewPodHook) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	if in.Command != nil {
		in, out := &in.Command, &out.Command
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]core.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Volumes != nil {
		in, out := &in.Volumes, &out.Volumes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}
func (in *ExecNewPodHook) DeepCopy() *ExecNewPodHook {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(ExecNewPodHook)
	in.DeepCopyInto(out)
	return out
}
func (in *LifecycleHook) DeepCopyInto(out *LifecycleHook) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	if in.ExecNewPod != nil {
		in, out := &in.ExecNewPod, &out.ExecNewPod
		*out = new(ExecNewPodHook)
		(*in).DeepCopyInto(*out)
	}
	if in.TagImages != nil {
		in, out := &in.TagImages, &out.TagImages
		*out = make([]TagImageHook, len(*in))
		copy(*out, *in)
	}
	return
}
func (in *LifecycleHook) DeepCopy() *LifecycleHook {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(LifecycleHook)
	in.DeepCopyInto(out)
	return out
}
func (in *RecreateDeploymentStrategyParams) DeepCopyInto(out *RecreateDeploymentStrategyParams) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	if in.TimeoutSeconds != nil {
		in, out := &in.TimeoutSeconds, &out.TimeoutSeconds
		*out = new(int64)
		**out = **in
	}
	if in.Pre != nil {
		in, out := &in.Pre, &out.Pre
		*out = new(LifecycleHook)
		(*in).DeepCopyInto(*out)
	}
	if in.Mid != nil {
		in, out := &in.Mid, &out.Mid
		*out = new(LifecycleHook)
		(*in).DeepCopyInto(*out)
	}
	if in.Post != nil {
		in, out := &in.Post, &out.Post
		*out = new(LifecycleHook)
		(*in).DeepCopyInto(*out)
	}
	return
}
func (in *RecreateDeploymentStrategyParams) DeepCopy() *RecreateDeploymentStrategyParams {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(RecreateDeploymentStrategyParams)
	in.DeepCopyInto(out)
	return out
}
func (in *RollingDeploymentStrategyParams) DeepCopyInto(out *RollingDeploymentStrategyParams) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	if in.UpdatePeriodSeconds != nil {
		in, out := &in.UpdatePeriodSeconds, &out.UpdatePeriodSeconds
		*out = new(int64)
		**out = **in
	}
	if in.IntervalSeconds != nil {
		in, out := &in.IntervalSeconds, &out.IntervalSeconds
		*out = new(int64)
		**out = **in
	}
	if in.TimeoutSeconds != nil {
		in, out := &in.TimeoutSeconds, &out.TimeoutSeconds
		*out = new(int64)
		**out = **in
	}
	out.MaxUnavailable = in.MaxUnavailable
	out.MaxSurge = in.MaxSurge
	if in.Pre != nil {
		in, out := &in.Pre, &out.Pre
		*out = new(LifecycleHook)
		(*in).DeepCopyInto(*out)
	}
	if in.Post != nil {
		in, out := &in.Post, &out.Post
		*out = new(LifecycleHook)
		(*in).DeepCopyInto(*out)
	}
	return
}
func (in *RollingDeploymentStrategyParams) DeepCopy() *RollingDeploymentStrategyParams {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(RollingDeploymentStrategyParams)
	in.DeepCopyInto(out)
	return out
}
func (in *TagImageHook) DeepCopyInto(out *TagImageHook) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	out.To = in.To
	return
}
func (in *TagImageHook) DeepCopy() *TagImageHook {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(TagImageHook)
	in.DeepCopyInto(out)
	return out
}
