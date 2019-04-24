package build

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
	core "k8s.io/kubernetes/pkg/apis/core"
)

func (in *BinaryBuildRequestOptions) DeepCopyInto(out *BinaryBuildRequestOptions) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	return
}
func (in *BinaryBuildRequestOptions) DeepCopy() *BinaryBuildRequestOptions {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(BinaryBuildRequestOptions)
	in.DeepCopyInto(out)
	return out
}
func (in *BinaryBuildRequestOptions) DeepCopyObject() runtime.Object {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
func (in *BinaryBuildSource) DeepCopyInto(out *BinaryBuildSource) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	return
}
func (in *BinaryBuildSource) DeepCopy() *BinaryBuildSource {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(BinaryBuildSource)
	in.DeepCopyInto(out)
	return out
}
func (in *BitbucketWebHookCause) DeepCopyInto(out *BitbucketWebHookCause) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	in.CommonWebHookCause.DeepCopyInto(&out.CommonWebHookCause)
	return
}
func (in *BitbucketWebHookCause) DeepCopy() *BitbucketWebHookCause {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(BitbucketWebHookCause)
	in.DeepCopyInto(out)
	return out
}
func (in *Build) DeepCopyInto(out *Build) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}
func (in *Build) DeepCopy() *Build {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(Build)
	in.DeepCopyInto(out)
	return out
}
func (in *Build) DeepCopyObject() runtime.Object {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
func (in *BuildConfig) DeepCopyInto(out *BuildConfig) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}
func (in *BuildConfig) DeepCopy() *BuildConfig {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(BuildConfig)
	in.DeepCopyInto(out)
	return out
}
func (in *BuildConfig) DeepCopyObject() runtime.Object {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
func (in *BuildConfigList) DeepCopyInto(out *BuildConfigList) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]BuildConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}
func (in *BuildConfigList) DeepCopy() *BuildConfigList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(BuildConfigList)
	in.DeepCopyInto(out)
	return out
}
func (in *BuildConfigList) DeepCopyObject() runtime.Object {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
func (in *BuildConfigSpec) DeepCopyInto(out *BuildConfigSpec) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	if in.Triggers != nil {
		in, out := &in.Triggers, &out.Triggers
		*out = make([]BuildTriggerPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.CommonSpec.DeepCopyInto(&out.CommonSpec)
	if in.SuccessfulBuildsHistoryLimit != nil {
		in, out := &in.SuccessfulBuildsHistoryLimit, &out.SuccessfulBuildsHistoryLimit
		*out = new(int32)
		**out = **in
	}
	if in.FailedBuildsHistoryLimit != nil {
		in, out := &in.FailedBuildsHistoryLimit, &out.FailedBuildsHistoryLimit
		*out = new(int32)
		**out = **in
	}
	return
}
func (in *BuildConfigSpec) DeepCopy() *BuildConfigSpec {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(BuildConfigSpec)
	in.DeepCopyInto(out)
	return out
}
func (in *BuildConfigStatus) DeepCopyInto(out *BuildConfigStatus) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	return
}
func (in *BuildConfigStatus) DeepCopy() *BuildConfigStatus {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(BuildConfigStatus)
	in.DeepCopyInto(out)
	return out
}
func (in *BuildList) DeepCopyInto(out *BuildList) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Build, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}
func (in *BuildList) DeepCopy() *BuildList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(BuildList)
	in.DeepCopyInto(out)
	return out
}
func (in *BuildList) DeepCopyObject() runtime.Object {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
func (in *BuildLog) DeepCopyInto(out *BuildLog) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	out.TypeMeta = in.TypeMeta
	return
}
func (in *BuildLog) DeepCopy() *BuildLog {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(BuildLog)
	in.DeepCopyInto(out)
	return out
}
func (in *BuildLog) DeepCopyObject() runtime.Object {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
func (in *BuildLogOptions) DeepCopyInto(out *BuildLogOptions) {
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
func (in *BuildLogOptions) DeepCopy() *BuildLogOptions {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(BuildLogOptions)
	in.DeepCopyInto(out)
	return out
}
func (in *BuildLogOptions) DeepCopyObject() runtime.Object {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
func (in *BuildOutput) DeepCopyInto(out *BuildOutput) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	if in.To != nil {
		in, out := &in.To, &out.To
		*out = new(core.ObjectReference)
		**out = **in
	}
	if in.PushSecret != nil {
		in, out := &in.PushSecret, &out.PushSecret
		*out = new(core.LocalObjectReference)
		**out = **in
	}
	if in.ImageLabels != nil {
		in, out := &in.ImageLabels, &out.ImageLabels
		*out = make([]ImageLabel, len(*in))
		copy(*out, *in)
	}
	return
}
func (in *BuildOutput) DeepCopy() *BuildOutput {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(BuildOutput)
	in.DeepCopyInto(out)
	return out
}
func (in *BuildPostCommitSpec) DeepCopyInto(out *BuildPostCommitSpec) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	if in.Command != nil {
		in, out := &in.Command, &out.Command
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Args != nil {
		in, out := &in.Args, &out.Args
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}
func (in *BuildPostCommitSpec) DeepCopy() *BuildPostCommitSpec {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(BuildPostCommitSpec)
	in.DeepCopyInto(out)
	return out
}
func (in *BuildRequest) DeepCopyInto(out *BuildRequest) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.Revision != nil {
		in, out := &in.Revision, &out.Revision
		*out = new(SourceRevision)
		(*in).DeepCopyInto(*out)
	}
	if in.TriggeredByImage != nil {
		in, out := &in.TriggeredByImage, &out.TriggeredByImage
		*out = new(core.ObjectReference)
		**out = **in
	}
	if in.From != nil {
		in, out := &in.From, &out.From
		*out = new(core.ObjectReference)
		**out = **in
	}
	if in.Binary != nil {
		in, out := &in.Binary, &out.Binary
		*out = new(BinaryBuildSource)
		**out = **in
	}
	if in.LastVersion != nil {
		in, out := &in.LastVersion, &out.LastVersion
		*out = new(int64)
		**out = **in
	}
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]core.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.TriggeredBy != nil {
		in, out := &in.TriggeredBy, &out.TriggeredBy
		*out = make([]BuildTriggerCause, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.DockerStrategyOptions != nil {
		in, out := &in.DockerStrategyOptions, &out.DockerStrategyOptions
		*out = new(DockerStrategyOptions)
		(*in).DeepCopyInto(*out)
	}
	if in.SourceStrategyOptions != nil {
		in, out := &in.SourceStrategyOptions, &out.SourceStrategyOptions
		*out = new(SourceStrategyOptions)
		(*in).DeepCopyInto(*out)
	}
	return
}
func (in *BuildRequest) DeepCopy() *BuildRequest {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(BuildRequest)
	in.DeepCopyInto(out)
	return out
}
func (in *BuildRequest) DeepCopyObject() runtime.Object {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
func (in *BuildSource) DeepCopyInto(out *BuildSource) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	if in.Binary != nil {
		in, out := &in.Binary, &out.Binary
		*out = new(BinaryBuildSource)
		**out = **in
	}
	if in.Dockerfile != nil {
		in, out := &in.Dockerfile, &out.Dockerfile
		*out = new(string)
		**out = **in
	}
	if in.Git != nil {
		in, out := &in.Git, &out.Git
		*out = new(GitBuildSource)
		(*in).DeepCopyInto(*out)
	}
	if in.Images != nil {
		in, out := &in.Images, &out.Images
		*out = make([]ImageSource, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SourceSecret != nil {
		in, out := &in.SourceSecret, &out.SourceSecret
		*out = new(core.LocalObjectReference)
		**out = **in
	}
	if in.Secrets != nil {
		in, out := &in.Secrets, &out.Secrets
		*out = make([]SecretBuildSource, len(*in))
		copy(*out, *in)
	}
	if in.ConfigMaps != nil {
		in, out := &in.ConfigMaps, &out.ConfigMaps
		*out = make([]ConfigMapBuildSource, len(*in))
		copy(*out, *in)
	}
	return
}
func (in *BuildSource) DeepCopy() *BuildSource {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(BuildSource)
	in.DeepCopyInto(out)
	return out
}
func (in *BuildSpec) DeepCopyInto(out *BuildSpec) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	in.CommonSpec.DeepCopyInto(&out.CommonSpec)
	if in.TriggeredBy != nil {
		in, out := &in.TriggeredBy, &out.TriggeredBy
		*out = make([]BuildTriggerCause, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}
func (in *BuildSpec) DeepCopy() *BuildSpec {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(BuildSpec)
	in.DeepCopyInto(out)
	return out
}
func (in *BuildStatus) DeepCopyInto(out *BuildStatus) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	if in.StartTimestamp != nil {
		in, out := &in.StartTimestamp, &out.StartTimestamp
		*out = (*in).DeepCopy()
	}
	if in.CompletionTimestamp != nil {
		in, out := &in.CompletionTimestamp, &out.CompletionTimestamp
		*out = (*in).DeepCopy()
	}
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = new(core.ObjectReference)
		**out = **in
	}
	in.Output.DeepCopyInto(&out.Output)
	if in.Stages != nil {
		in, out := &in.Stages, &out.Stages
		*out = make([]StageInfo, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}
func (in *BuildStatus) DeepCopy() *BuildStatus {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(BuildStatus)
	in.DeepCopyInto(out)
	return out
}
func (in *BuildStatusOutput) DeepCopyInto(out *BuildStatusOutput) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	if in.To != nil {
		in, out := &in.To, &out.To
		*out = new(BuildStatusOutputTo)
		**out = **in
	}
	return
}
func (in *BuildStatusOutput) DeepCopy() *BuildStatusOutput {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(BuildStatusOutput)
	in.DeepCopyInto(out)
	return out
}
func (in *BuildStatusOutputTo) DeepCopyInto(out *BuildStatusOutputTo) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	return
}
func (in *BuildStatusOutputTo) DeepCopy() *BuildStatusOutputTo {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(BuildStatusOutputTo)
	in.DeepCopyInto(out)
	return out
}
func (in *BuildStrategy) DeepCopyInto(out *BuildStrategy) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	if in.DockerStrategy != nil {
		in, out := &in.DockerStrategy, &out.DockerStrategy
		*out = new(DockerBuildStrategy)
		(*in).DeepCopyInto(*out)
	}
	if in.SourceStrategy != nil {
		in, out := &in.SourceStrategy, &out.SourceStrategy
		*out = new(SourceBuildStrategy)
		(*in).DeepCopyInto(*out)
	}
	if in.CustomStrategy != nil {
		in, out := &in.CustomStrategy, &out.CustomStrategy
		*out = new(CustomBuildStrategy)
		(*in).DeepCopyInto(*out)
	}
	if in.JenkinsPipelineStrategy != nil {
		in, out := &in.JenkinsPipelineStrategy, &out.JenkinsPipelineStrategy
		*out = new(JenkinsPipelineBuildStrategy)
		(*in).DeepCopyInto(*out)
	}
	return
}
func (in *BuildStrategy) DeepCopy() *BuildStrategy {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(BuildStrategy)
	in.DeepCopyInto(out)
	return out
}
func (in *BuildTriggerCause) DeepCopyInto(out *BuildTriggerCause) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	if in.GenericWebHook != nil {
		in, out := &in.GenericWebHook, &out.GenericWebHook
		*out = new(GenericWebHookCause)
		(*in).DeepCopyInto(*out)
	}
	if in.GitHubWebHook != nil {
		in, out := &in.GitHubWebHook, &out.GitHubWebHook
		*out = new(GitHubWebHookCause)
		(*in).DeepCopyInto(*out)
	}
	if in.ImageChangeBuild != nil {
		in, out := &in.ImageChangeBuild, &out.ImageChangeBuild
		*out = new(ImageChangeCause)
		(*in).DeepCopyInto(*out)
	}
	if in.GitLabWebHook != nil {
		in, out := &in.GitLabWebHook, &out.GitLabWebHook
		*out = new(GitLabWebHookCause)
		(*in).DeepCopyInto(*out)
	}
	if in.BitbucketWebHook != nil {
		in, out := &in.BitbucketWebHook, &out.BitbucketWebHook
		*out = new(BitbucketWebHookCause)
		(*in).DeepCopyInto(*out)
	}
	return
}
func (in *BuildTriggerCause) DeepCopy() *BuildTriggerCause {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(BuildTriggerCause)
	in.DeepCopyInto(out)
	return out
}
func (in *BuildTriggerPolicy) DeepCopyInto(out *BuildTriggerPolicy) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	if in.GitHubWebHook != nil {
		in, out := &in.GitHubWebHook, &out.GitHubWebHook
		*out = new(WebHookTrigger)
		(*in).DeepCopyInto(*out)
	}
	if in.GenericWebHook != nil {
		in, out := &in.GenericWebHook, &out.GenericWebHook
		*out = new(WebHookTrigger)
		(*in).DeepCopyInto(*out)
	}
	if in.ImageChange != nil {
		in, out := &in.ImageChange, &out.ImageChange
		*out = new(ImageChangeTrigger)
		(*in).DeepCopyInto(*out)
	}
	if in.GitLabWebHook != nil {
		in, out := &in.GitLabWebHook, &out.GitLabWebHook
		*out = new(WebHookTrigger)
		(*in).DeepCopyInto(*out)
	}
	if in.BitbucketWebHook != nil {
		in, out := &in.BitbucketWebHook, &out.BitbucketWebHook
		*out = new(WebHookTrigger)
		(*in).DeepCopyInto(*out)
	}
	return
}
func (in *BuildTriggerPolicy) DeepCopy() *BuildTriggerPolicy {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(BuildTriggerPolicy)
	in.DeepCopyInto(out)
	return out
}
func (in *CommonSpec) DeepCopyInto(out *CommonSpec) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	in.Source.DeepCopyInto(&out.Source)
	if in.Revision != nil {
		in, out := &in.Revision, &out.Revision
		*out = new(SourceRevision)
		(*in).DeepCopyInto(*out)
	}
	in.Strategy.DeepCopyInto(&out.Strategy)
	in.Output.DeepCopyInto(&out.Output)
	in.Resources.DeepCopyInto(&out.Resources)
	in.PostCommit.DeepCopyInto(&out.PostCommit)
	if in.CompletionDeadlineSeconds != nil {
		in, out := &in.CompletionDeadlineSeconds, &out.CompletionDeadlineSeconds
		*out = new(int64)
		**out = **in
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}
func (in *CommonSpec) DeepCopy() *CommonSpec {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(CommonSpec)
	in.DeepCopyInto(out)
	return out
}
func (in *CommonWebHookCause) DeepCopyInto(out *CommonWebHookCause) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	if in.Revision != nil {
		in, out := &in.Revision, &out.Revision
		*out = new(SourceRevision)
		(*in).DeepCopyInto(*out)
	}
	return
}
func (in *CommonWebHookCause) DeepCopy() *CommonWebHookCause {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(CommonWebHookCause)
	in.DeepCopyInto(out)
	return out
}
func (in *ConfigMapBuildSource) DeepCopyInto(out *ConfigMapBuildSource) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	out.ConfigMap = in.ConfigMap
	return
}
func (in *ConfigMapBuildSource) DeepCopy() *ConfigMapBuildSource {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(ConfigMapBuildSource)
	in.DeepCopyInto(out)
	return out
}
func (in *CustomBuildStrategy) DeepCopyInto(out *CustomBuildStrategy) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	out.From = in.From
	if in.PullSecret != nil {
		in, out := &in.PullSecret, &out.PullSecret
		*out = new(core.LocalObjectReference)
		**out = **in
	}
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]core.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Secrets != nil {
		in, out := &in.Secrets, &out.Secrets
		*out = make([]SecretSpec, len(*in))
		copy(*out, *in)
	}
	return
}
func (in *CustomBuildStrategy) DeepCopy() *CustomBuildStrategy {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(CustomBuildStrategy)
	in.DeepCopyInto(out)
	return out
}
func (in *DockerBuildStrategy) DeepCopyInto(out *DockerBuildStrategy) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	if in.From != nil {
		in, out := &in.From, &out.From
		*out = new(core.ObjectReference)
		**out = **in
	}
	if in.PullSecret != nil {
		in, out := &in.PullSecret, &out.PullSecret
		*out = new(core.LocalObjectReference)
		**out = **in
	}
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]core.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.BuildArgs != nil {
		in, out := &in.BuildArgs, &out.BuildArgs
		*out = make([]core.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ImageOptimizationPolicy != nil {
		in, out := &in.ImageOptimizationPolicy, &out.ImageOptimizationPolicy
		*out = new(ImageOptimizationPolicy)
		**out = **in
	}
	return
}
func (in *DockerBuildStrategy) DeepCopy() *DockerBuildStrategy {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(DockerBuildStrategy)
	in.DeepCopyInto(out)
	return out
}
func (in *DockerStrategyOptions) DeepCopyInto(out *DockerStrategyOptions) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	if in.BuildArgs != nil {
		in, out := &in.BuildArgs, &out.BuildArgs
		*out = make([]core.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.NoCache != nil {
		in, out := &in.NoCache, &out.NoCache
		*out = new(bool)
		**out = **in
	}
	return
}
func (in *DockerStrategyOptions) DeepCopy() *DockerStrategyOptions {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(DockerStrategyOptions)
	in.DeepCopyInto(out)
	return out
}
func (in *GenericWebHookCause) DeepCopyInto(out *GenericWebHookCause) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	if in.Revision != nil {
		in, out := &in.Revision, &out.Revision
		*out = new(SourceRevision)
		(*in).DeepCopyInto(*out)
	}
	return
}
func (in *GenericWebHookCause) DeepCopy() *GenericWebHookCause {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(GenericWebHookCause)
	in.DeepCopyInto(out)
	return out
}
func (in *GenericWebHookEvent) DeepCopyInto(out *GenericWebHookEvent) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	if in.Git != nil {
		in, out := &in.Git, &out.Git
		*out = new(GitInfo)
		(*in).DeepCopyInto(*out)
	}
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]core.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.DockerStrategyOptions != nil {
		in, out := &in.DockerStrategyOptions, &out.DockerStrategyOptions
		*out = new(DockerStrategyOptions)
		(*in).DeepCopyInto(*out)
	}
	return
}
func (in *GenericWebHookEvent) DeepCopy() *GenericWebHookEvent {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(GenericWebHookEvent)
	in.DeepCopyInto(out)
	return out
}
func (in *GitBuildSource) DeepCopyInto(out *GitBuildSource) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	in.ProxyConfig.DeepCopyInto(&out.ProxyConfig)
	return
}
func (in *GitBuildSource) DeepCopy() *GitBuildSource {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(GitBuildSource)
	in.DeepCopyInto(out)
	return out
}
func (in *GitHubWebHookCause) DeepCopyInto(out *GitHubWebHookCause) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	if in.Revision != nil {
		in, out := &in.Revision, &out.Revision
		*out = new(SourceRevision)
		(*in).DeepCopyInto(*out)
	}
	return
}
func (in *GitHubWebHookCause) DeepCopy() *GitHubWebHookCause {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(GitHubWebHookCause)
	in.DeepCopyInto(out)
	return out
}
func (in *GitInfo) DeepCopyInto(out *GitInfo) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	in.GitBuildSource.DeepCopyInto(&out.GitBuildSource)
	out.GitSourceRevision = in.GitSourceRevision
	if in.Refs != nil {
		in, out := &in.Refs, &out.Refs
		*out = make([]GitRefInfo, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}
func (in *GitInfo) DeepCopy() *GitInfo {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(GitInfo)
	in.DeepCopyInto(out)
	return out
}
func (in *GitLabWebHookCause) DeepCopyInto(out *GitLabWebHookCause) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	in.CommonWebHookCause.DeepCopyInto(&out.CommonWebHookCause)
	return
}
func (in *GitLabWebHookCause) DeepCopy() *GitLabWebHookCause {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(GitLabWebHookCause)
	in.DeepCopyInto(out)
	return out
}
func (in *GitRefInfo) DeepCopyInto(out *GitRefInfo) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	in.GitBuildSource.DeepCopyInto(&out.GitBuildSource)
	out.GitSourceRevision = in.GitSourceRevision
	return
}
func (in *GitRefInfo) DeepCopy() *GitRefInfo {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(GitRefInfo)
	in.DeepCopyInto(out)
	return out
}
func (in *GitSourceRevision) DeepCopyInto(out *GitSourceRevision) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	out.Author = in.Author
	out.Committer = in.Committer
	return
}
func (in *GitSourceRevision) DeepCopy() *GitSourceRevision {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(GitSourceRevision)
	in.DeepCopyInto(out)
	return out
}
func (in *ImageChangeCause) DeepCopyInto(out *ImageChangeCause) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	if in.FromRef != nil {
		in, out := &in.FromRef, &out.FromRef
		*out = new(core.ObjectReference)
		**out = **in
	}
	return
}
func (in *ImageChangeCause) DeepCopy() *ImageChangeCause {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(ImageChangeCause)
	in.DeepCopyInto(out)
	return out
}
func (in *ImageChangeTrigger) DeepCopyInto(out *ImageChangeTrigger) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	if in.From != nil {
		in, out := &in.From, &out.From
		*out = new(core.ObjectReference)
		**out = **in
	}
	return
}
func (in *ImageChangeTrigger) DeepCopy() *ImageChangeTrigger {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(ImageChangeTrigger)
	in.DeepCopyInto(out)
	return out
}
func (in *ImageLabel) DeepCopyInto(out *ImageLabel) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	return
}
func (in *ImageLabel) DeepCopy() *ImageLabel {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(ImageLabel)
	in.DeepCopyInto(out)
	return out
}
func (in *ImageSource) DeepCopyInto(out *ImageSource) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	out.From = in.From
	if in.As != nil {
		in, out := &in.As, &out.As
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Paths != nil {
		in, out := &in.Paths, &out.Paths
		*out = make([]ImageSourcePath, len(*in))
		copy(*out, *in)
	}
	if in.PullSecret != nil {
		in, out := &in.PullSecret, &out.PullSecret
		*out = new(core.LocalObjectReference)
		**out = **in
	}
	return
}
func (in *ImageSource) DeepCopy() *ImageSource {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(ImageSource)
	in.DeepCopyInto(out)
	return out
}
func (in *ImageSourcePath) DeepCopyInto(out *ImageSourcePath) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	return
}
func (in *ImageSourcePath) DeepCopy() *ImageSourcePath {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(ImageSourcePath)
	in.DeepCopyInto(out)
	return out
}
func (in *JenkinsPipelineBuildStrategy) DeepCopyInto(out *JenkinsPipelineBuildStrategy) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]core.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}
func (in *JenkinsPipelineBuildStrategy) DeepCopy() *JenkinsPipelineBuildStrategy {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(JenkinsPipelineBuildStrategy)
	in.DeepCopyInto(out)
	return out
}
func (in *ProxyConfig) DeepCopyInto(out *ProxyConfig) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	if in.HTTPProxy != nil {
		in, out := &in.HTTPProxy, &out.HTTPProxy
		*out = new(string)
		**out = **in
	}
	if in.HTTPSProxy != nil {
		in, out := &in.HTTPSProxy, &out.HTTPSProxy
		*out = new(string)
		**out = **in
	}
	if in.NoProxy != nil {
		in, out := &in.NoProxy, &out.NoProxy
		*out = new(string)
		**out = **in
	}
	return
}
func (in *ProxyConfig) DeepCopy() *ProxyConfig {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(ProxyConfig)
	in.DeepCopyInto(out)
	return out
}
func (in *SecretBuildSource) DeepCopyInto(out *SecretBuildSource) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	out.Secret = in.Secret
	return
}
func (in *SecretBuildSource) DeepCopy() *SecretBuildSource {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(SecretBuildSource)
	in.DeepCopyInto(out)
	return out
}
func (in *SecretLocalReference) DeepCopyInto(out *SecretLocalReference) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	return
}
func (in *SecretLocalReference) DeepCopy() *SecretLocalReference {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(SecretLocalReference)
	in.DeepCopyInto(out)
	return out
}
func (in *SecretSpec) DeepCopyInto(out *SecretSpec) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	out.SecretSource = in.SecretSource
	return
}
func (in *SecretSpec) DeepCopy() *SecretSpec {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(SecretSpec)
	in.DeepCopyInto(out)
	return out
}
func (in *SourceBuildStrategy) DeepCopyInto(out *SourceBuildStrategy) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	out.From = in.From
	if in.PullSecret != nil {
		in, out := &in.PullSecret, &out.PullSecret
		*out = new(core.LocalObjectReference)
		**out = **in
	}
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]core.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Incremental != nil {
		in, out := &in.Incremental, &out.Incremental
		*out = new(bool)
		**out = **in
	}
	return
}
func (in *SourceBuildStrategy) DeepCopy() *SourceBuildStrategy {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(SourceBuildStrategy)
	in.DeepCopyInto(out)
	return out
}
func (in *SourceControlUser) DeepCopyInto(out *SourceControlUser) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	return
}
func (in *SourceControlUser) DeepCopy() *SourceControlUser {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(SourceControlUser)
	in.DeepCopyInto(out)
	return out
}
func (in *SourceRevision) DeepCopyInto(out *SourceRevision) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	if in.Git != nil {
		in, out := &in.Git, &out.Git
		*out = new(GitSourceRevision)
		**out = **in
	}
	return
}
func (in *SourceRevision) DeepCopy() *SourceRevision {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(SourceRevision)
	in.DeepCopyInto(out)
	return out
}
func (in *SourceStrategyOptions) DeepCopyInto(out *SourceStrategyOptions) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	if in.Incremental != nil {
		in, out := &in.Incremental, &out.Incremental
		*out = new(bool)
		**out = **in
	}
	return
}
func (in *SourceStrategyOptions) DeepCopy() *SourceStrategyOptions {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(SourceStrategyOptions)
	in.DeepCopyInto(out)
	return out
}
func (in *StageInfo) DeepCopyInto(out *StageInfo) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	in.StartTime.DeepCopyInto(&out.StartTime)
	if in.Steps != nil {
		in, out := &in.Steps, &out.Steps
		*out = make([]StepInfo, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}
func (in *StageInfo) DeepCopy() *StageInfo {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(StageInfo)
	in.DeepCopyInto(out)
	return out
}
func (in *StepInfo) DeepCopyInto(out *StepInfo) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	in.StartTime.DeepCopyInto(&out.StartTime)
	return
}
func (in *StepInfo) DeepCopy() *StepInfo {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(StepInfo)
	in.DeepCopyInto(out)
	return out
}
func (in *WebHookTrigger) DeepCopyInto(out *WebHookTrigger) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	if in.SecretReference != nil {
		in, out := &in.SecretReference, &out.SecretReference
		*out = new(SecretLocalReference)
		**out = **in
	}
	return
}
func (in *WebHookTrigger) DeepCopy() *WebHookTrigger {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(WebHookTrigger)
	in.DeepCopyInto(out)
	return out
}
