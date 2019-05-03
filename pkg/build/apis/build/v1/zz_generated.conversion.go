package v1

import (
	v1 "github.com/openshift/api/build/v1"
	build "github.com/openshift/origin/pkg/build/apis/build"
	apicorev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	core "k8s.io/kubernetes/pkg/apis/core"
	corev1 "k8s.io/kubernetes/pkg/apis/core/v1"
	time "time"
	unsafe "unsafe"
)

func init() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	localSchemeBuilder.Register(RegisterConversions)
}
func RegisterConversions(s *runtime.Scheme) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := s.AddGeneratedConversionFunc((*v1.BinaryBuildRequestOptions)(nil), (*build.BinaryBuildRequestOptions)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_BinaryBuildRequestOptions_To_build_BinaryBuildRequestOptions(a.(*v1.BinaryBuildRequestOptions), b.(*build.BinaryBuildRequestOptions), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*build.BinaryBuildRequestOptions)(nil), (*v1.BinaryBuildRequestOptions)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_build_BinaryBuildRequestOptions_To_v1_BinaryBuildRequestOptions(a.(*build.BinaryBuildRequestOptions), b.(*v1.BinaryBuildRequestOptions), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.BinaryBuildSource)(nil), (*build.BinaryBuildSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_BinaryBuildSource_To_build_BinaryBuildSource(a.(*v1.BinaryBuildSource), b.(*build.BinaryBuildSource), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*build.BinaryBuildSource)(nil), (*v1.BinaryBuildSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_build_BinaryBuildSource_To_v1_BinaryBuildSource(a.(*build.BinaryBuildSource), b.(*v1.BinaryBuildSource), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.BitbucketWebHookCause)(nil), (*build.BitbucketWebHookCause)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_BitbucketWebHookCause_To_build_BitbucketWebHookCause(a.(*v1.BitbucketWebHookCause), b.(*build.BitbucketWebHookCause), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*build.BitbucketWebHookCause)(nil), (*v1.BitbucketWebHookCause)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_build_BitbucketWebHookCause_To_v1_BitbucketWebHookCause(a.(*build.BitbucketWebHookCause), b.(*v1.BitbucketWebHookCause), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.Build)(nil), (*build.Build)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_Build_To_build_Build(a.(*v1.Build), b.(*build.Build), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*build.Build)(nil), (*v1.Build)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_build_Build_To_v1_Build(a.(*build.Build), b.(*v1.Build), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.BuildConfig)(nil), (*build.BuildConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_BuildConfig_To_build_BuildConfig(a.(*v1.BuildConfig), b.(*build.BuildConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*build.BuildConfig)(nil), (*v1.BuildConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_build_BuildConfig_To_v1_BuildConfig(a.(*build.BuildConfig), b.(*v1.BuildConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.BuildConfigList)(nil), (*build.BuildConfigList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_BuildConfigList_To_build_BuildConfigList(a.(*v1.BuildConfigList), b.(*build.BuildConfigList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*build.BuildConfigList)(nil), (*v1.BuildConfigList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_build_BuildConfigList_To_v1_BuildConfigList(a.(*build.BuildConfigList), b.(*v1.BuildConfigList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.BuildConfigSpec)(nil), (*build.BuildConfigSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_BuildConfigSpec_To_build_BuildConfigSpec(a.(*v1.BuildConfigSpec), b.(*build.BuildConfigSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*build.BuildConfigSpec)(nil), (*v1.BuildConfigSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_build_BuildConfigSpec_To_v1_BuildConfigSpec(a.(*build.BuildConfigSpec), b.(*v1.BuildConfigSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.BuildConfigStatus)(nil), (*build.BuildConfigStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_BuildConfigStatus_To_build_BuildConfigStatus(a.(*v1.BuildConfigStatus), b.(*build.BuildConfigStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*build.BuildConfigStatus)(nil), (*v1.BuildConfigStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_build_BuildConfigStatus_To_v1_BuildConfigStatus(a.(*build.BuildConfigStatus), b.(*v1.BuildConfigStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.BuildList)(nil), (*build.BuildList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_BuildList_To_build_BuildList(a.(*v1.BuildList), b.(*build.BuildList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*build.BuildList)(nil), (*v1.BuildList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_build_BuildList_To_v1_BuildList(a.(*build.BuildList), b.(*v1.BuildList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.BuildLog)(nil), (*build.BuildLog)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_BuildLog_To_build_BuildLog(a.(*v1.BuildLog), b.(*build.BuildLog), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*build.BuildLog)(nil), (*v1.BuildLog)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_build_BuildLog_To_v1_BuildLog(a.(*build.BuildLog), b.(*v1.BuildLog), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.BuildLogOptions)(nil), (*build.BuildLogOptions)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_BuildLogOptions_To_build_BuildLogOptions(a.(*v1.BuildLogOptions), b.(*build.BuildLogOptions), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*build.BuildLogOptions)(nil), (*v1.BuildLogOptions)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_build_BuildLogOptions_To_v1_BuildLogOptions(a.(*build.BuildLogOptions), b.(*v1.BuildLogOptions), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.BuildOutput)(nil), (*build.BuildOutput)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_BuildOutput_To_build_BuildOutput(a.(*v1.BuildOutput), b.(*build.BuildOutput), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*build.BuildOutput)(nil), (*v1.BuildOutput)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_build_BuildOutput_To_v1_BuildOutput(a.(*build.BuildOutput), b.(*v1.BuildOutput), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.BuildPostCommitSpec)(nil), (*build.BuildPostCommitSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_BuildPostCommitSpec_To_build_BuildPostCommitSpec(a.(*v1.BuildPostCommitSpec), b.(*build.BuildPostCommitSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*build.BuildPostCommitSpec)(nil), (*v1.BuildPostCommitSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_build_BuildPostCommitSpec_To_v1_BuildPostCommitSpec(a.(*build.BuildPostCommitSpec), b.(*v1.BuildPostCommitSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.BuildRequest)(nil), (*build.BuildRequest)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_BuildRequest_To_build_BuildRequest(a.(*v1.BuildRequest), b.(*build.BuildRequest), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*build.BuildRequest)(nil), (*v1.BuildRequest)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_build_BuildRequest_To_v1_BuildRequest(a.(*build.BuildRequest), b.(*v1.BuildRequest), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.BuildSource)(nil), (*build.BuildSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_BuildSource_To_build_BuildSource(a.(*v1.BuildSource), b.(*build.BuildSource), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*build.BuildSource)(nil), (*v1.BuildSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_build_BuildSource_To_v1_BuildSource(a.(*build.BuildSource), b.(*v1.BuildSource), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.BuildSpec)(nil), (*build.BuildSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_BuildSpec_To_build_BuildSpec(a.(*v1.BuildSpec), b.(*build.BuildSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*build.BuildSpec)(nil), (*v1.BuildSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_build_BuildSpec_To_v1_BuildSpec(a.(*build.BuildSpec), b.(*v1.BuildSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.BuildStatus)(nil), (*build.BuildStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_BuildStatus_To_build_BuildStatus(a.(*v1.BuildStatus), b.(*build.BuildStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*build.BuildStatus)(nil), (*v1.BuildStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_build_BuildStatus_To_v1_BuildStatus(a.(*build.BuildStatus), b.(*v1.BuildStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.BuildStatusOutput)(nil), (*build.BuildStatusOutput)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_BuildStatusOutput_To_build_BuildStatusOutput(a.(*v1.BuildStatusOutput), b.(*build.BuildStatusOutput), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*build.BuildStatusOutput)(nil), (*v1.BuildStatusOutput)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_build_BuildStatusOutput_To_v1_BuildStatusOutput(a.(*build.BuildStatusOutput), b.(*v1.BuildStatusOutput), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.BuildStatusOutputTo)(nil), (*build.BuildStatusOutputTo)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_BuildStatusOutputTo_To_build_BuildStatusOutputTo(a.(*v1.BuildStatusOutputTo), b.(*build.BuildStatusOutputTo), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*build.BuildStatusOutputTo)(nil), (*v1.BuildStatusOutputTo)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_build_BuildStatusOutputTo_To_v1_BuildStatusOutputTo(a.(*build.BuildStatusOutputTo), b.(*v1.BuildStatusOutputTo), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.BuildStrategy)(nil), (*build.BuildStrategy)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_BuildStrategy_To_build_BuildStrategy(a.(*v1.BuildStrategy), b.(*build.BuildStrategy), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*build.BuildStrategy)(nil), (*v1.BuildStrategy)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_build_BuildStrategy_To_v1_BuildStrategy(a.(*build.BuildStrategy), b.(*v1.BuildStrategy), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.BuildTriggerCause)(nil), (*build.BuildTriggerCause)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_BuildTriggerCause_To_build_BuildTriggerCause(a.(*v1.BuildTriggerCause), b.(*build.BuildTriggerCause), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*build.BuildTriggerCause)(nil), (*v1.BuildTriggerCause)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_build_BuildTriggerCause_To_v1_BuildTriggerCause(a.(*build.BuildTriggerCause), b.(*v1.BuildTriggerCause), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.BuildTriggerPolicy)(nil), (*build.BuildTriggerPolicy)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_BuildTriggerPolicy_To_build_BuildTriggerPolicy(a.(*v1.BuildTriggerPolicy), b.(*build.BuildTriggerPolicy), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*build.BuildTriggerPolicy)(nil), (*v1.BuildTriggerPolicy)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_build_BuildTriggerPolicy_To_v1_BuildTriggerPolicy(a.(*build.BuildTriggerPolicy), b.(*v1.BuildTriggerPolicy), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.CommonSpec)(nil), (*build.CommonSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_CommonSpec_To_build_CommonSpec(a.(*v1.CommonSpec), b.(*build.CommonSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*build.CommonSpec)(nil), (*v1.CommonSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_build_CommonSpec_To_v1_CommonSpec(a.(*build.CommonSpec), b.(*v1.CommonSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.CommonWebHookCause)(nil), (*build.CommonWebHookCause)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_CommonWebHookCause_To_build_CommonWebHookCause(a.(*v1.CommonWebHookCause), b.(*build.CommonWebHookCause), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*build.CommonWebHookCause)(nil), (*v1.CommonWebHookCause)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_build_CommonWebHookCause_To_v1_CommonWebHookCause(a.(*build.CommonWebHookCause), b.(*v1.CommonWebHookCause), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.ConfigMapBuildSource)(nil), (*build.ConfigMapBuildSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_ConfigMapBuildSource_To_build_ConfigMapBuildSource(a.(*v1.ConfigMapBuildSource), b.(*build.ConfigMapBuildSource), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*build.ConfigMapBuildSource)(nil), (*v1.ConfigMapBuildSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_build_ConfigMapBuildSource_To_v1_ConfigMapBuildSource(a.(*build.ConfigMapBuildSource), b.(*v1.ConfigMapBuildSource), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.CustomBuildStrategy)(nil), (*build.CustomBuildStrategy)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_CustomBuildStrategy_To_build_CustomBuildStrategy(a.(*v1.CustomBuildStrategy), b.(*build.CustomBuildStrategy), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*build.CustomBuildStrategy)(nil), (*v1.CustomBuildStrategy)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_build_CustomBuildStrategy_To_v1_CustomBuildStrategy(a.(*build.CustomBuildStrategy), b.(*v1.CustomBuildStrategy), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.DockerBuildStrategy)(nil), (*build.DockerBuildStrategy)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_DockerBuildStrategy_To_build_DockerBuildStrategy(a.(*v1.DockerBuildStrategy), b.(*build.DockerBuildStrategy), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*build.DockerBuildStrategy)(nil), (*v1.DockerBuildStrategy)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_build_DockerBuildStrategy_To_v1_DockerBuildStrategy(a.(*build.DockerBuildStrategy), b.(*v1.DockerBuildStrategy), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.DockerStrategyOptions)(nil), (*build.DockerStrategyOptions)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_DockerStrategyOptions_To_build_DockerStrategyOptions(a.(*v1.DockerStrategyOptions), b.(*build.DockerStrategyOptions), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*build.DockerStrategyOptions)(nil), (*v1.DockerStrategyOptions)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_build_DockerStrategyOptions_To_v1_DockerStrategyOptions(a.(*build.DockerStrategyOptions), b.(*v1.DockerStrategyOptions), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.GenericWebHookCause)(nil), (*build.GenericWebHookCause)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_GenericWebHookCause_To_build_GenericWebHookCause(a.(*v1.GenericWebHookCause), b.(*build.GenericWebHookCause), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*build.GenericWebHookCause)(nil), (*v1.GenericWebHookCause)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_build_GenericWebHookCause_To_v1_GenericWebHookCause(a.(*build.GenericWebHookCause), b.(*v1.GenericWebHookCause), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.GenericWebHookEvent)(nil), (*build.GenericWebHookEvent)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_GenericWebHookEvent_To_build_GenericWebHookEvent(a.(*v1.GenericWebHookEvent), b.(*build.GenericWebHookEvent), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*build.GenericWebHookEvent)(nil), (*v1.GenericWebHookEvent)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_build_GenericWebHookEvent_To_v1_GenericWebHookEvent(a.(*build.GenericWebHookEvent), b.(*v1.GenericWebHookEvent), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.GitBuildSource)(nil), (*build.GitBuildSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_GitBuildSource_To_build_GitBuildSource(a.(*v1.GitBuildSource), b.(*build.GitBuildSource), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*build.GitBuildSource)(nil), (*v1.GitBuildSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_build_GitBuildSource_To_v1_GitBuildSource(a.(*build.GitBuildSource), b.(*v1.GitBuildSource), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.GitHubWebHookCause)(nil), (*build.GitHubWebHookCause)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_GitHubWebHookCause_To_build_GitHubWebHookCause(a.(*v1.GitHubWebHookCause), b.(*build.GitHubWebHookCause), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*build.GitHubWebHookCause)(nil), (*v1.GitHubWebHookCause)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_build_GitHubWebHookCause_To_v1_GitHubWebHookCause(a.(*build.GitHubWebHookCause), b.(*v1.GitHubWebHookCause), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.GitInfo)(nil), (*build.GitInfo)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_GitInfo_To_build_GitInfo(a.(*v1.GitInfo), b.(*build.GitInfo), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*build.GitInfo)(nil), (*v1.GitInfo)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_build_GitInfo_To_v1_GitInfo(a.(*build.GitInfo), b.(*v1.GitInfo), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.GitLabWebHookCause)(nil), (*build.GitLabWebHookCause)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_GitLabWebHookCause_To_build_GitLabWebHookCause(a.(*v1.GitLabWebHookCause), b.(*build.GitLabWebHookCause), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*build.GitLabWebHookCause)(nil), (*v1.GitLabWebHookCause)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_build_GitLabWebHookCause_To_v1_GitLabWebHookCause(a.(*build.GitLabWebHookCause), b.(*v1.GitLabWebHookCause), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.GitRefInfo)(nil), (*build.GitRefInfo)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_GitRefInfo_To_build_GitRefInfo(a.(*v1.GitRefInfo), b.(*build.GitRefInfo), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*build.GitRefInfo)(nil), (*v1.GitRefInfo)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_build_GitRefInfo_To_v1_GitRefInfo(a.(*build.GitRefInfo), b.(*v1.GitRefInfo), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.GitSourceRevision)(nil), (*build.GitSourceRevision)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_GitSourceRevision_To_build_GitSourceRevision(a.(*v1.GitSourceRevision), b.(*build.GitSourceRevision), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*build.GitSourceRevision)(nil), (*v1.GitSourceRevision)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_build_GitSourceRevision_To_v1_GitSourceRevision(a.(*build.GitSourceRevision), b.(*v1.GitSourceRevision), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.ImageChangeCause)(nil), (*build.ImageChangeCause)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_ImageChangeCause_To_build_ImageChangeCause(a.(*v1.ImageChangeCause), b.(*build.ImageChangeCause), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*build.ImageChangeCause)(nil), (*v1.ImageChangeCause)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_build_ImageChangeCause_To_v1_ImageChangeCause(a.(*build.ImageChangeCause), b.(*v1.ImageChangeCause), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.ImageChangeTrigger)(nil), (*build.ImageChangeTrigger)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_ImageChangeTrigger_To_build_ImageChangeTrigger(a.(*v1.ImageChangeTrigger), b.(*build.ImageChangeTrigger), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*build.ImageChangeTrigger)(nil), (*v1.ImageChangeTrigger)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_build_ImageChangeTrigger_To_v1_ImageChangeTrigger(a.(*build.ImageChangeTrigger), b.(*v1.ImageChangeTrigger), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.ImageLabel)(nil), (*build.ImageLabel)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_ImageLabel_To_build_ImageLabel(a.(*v1.ImageLabel), b.(*build.ImageLabel), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*build.ImageLabel)(nil), (*v1.ImageLabel)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_build_ImageLabel_To_v1_ImageLabel(a.(*build.ImageLabel), b.(*v1.ImageLabel), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.ImageSource)(nil), (*build.ImageSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_ImageSource_To_build_ImageSource(a.(*v1.ImageSource), b.(*build.ImageSource), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*build.ImageSource)(nil), (*v1.ImageSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_build_ImageSource_To_v1_ImageSource(a.(*build.ImageSource), b.(*v1.ImageSource), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.ImageSourcePath)(nil), (*build.ImageSourcePath)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_ImageSourcePath_To_build_ImageSourcePath(a.(*v1.ImageSourcePath), b.(*build.ImageSourcePath), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*build.ImageSourcePath)(nil), (*v1.ImageSourcePath)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_build_ImageSourcePath_To_v1_ImageSourcePath(a.(*build.ImageSourcePath), b.(*v1.ImageSourcePath), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.JenkinsPipelineBuildStrategy)(nil), (*build.JenkinsPipelineBuildStrategy)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_JenkinsPipelineBuildStrategy_To_build_JenkinsPipelineBuildStrategy(a.(*v1.JenkinsPipelineBuildStrategy), b.(*build.JenkinsPipelineBuildStrategy), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*build.JenkinsPipelineBuildStrategy)(nil), (*v1.JenkinsPipelineBuildStrategy)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_build_JenkinsPipelineBuildStrategy_To_v1_JenkinsPipelineBuildStrategy(a.(*build.JenkinsPipelineBuildStrategy), b.(*v1.JenkinsPipelineBuildStrategy), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.ProxyConfig)(nil), (*build.ProxyConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_ProxyConfig_To_build_ProxyConfig(a.(*v1.ProxyConfig), b.(*build.ProxyConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*build.ProxyConfig)(nil), (*v1.ProxyConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_build_ProxyConfig_To_v1_ProxyConfig(a.(*build.ProxyConfig), b.(*v1.ProxyConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.SecretBuildSource)(nil), (*build.SecretBuildSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_SecretBuildSource_To_build_SecretBuildSource(a.(*v1.SecretBuildSource), b.(*build.SecretBuildSource), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*build.SecretBuildSource)(nil), (*v1.SecretBuildSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_build_SecretBuildSource_To_v1_SecretBuildSource(a.(*build.SecretBuildSource), b.(*v1.SecretBuildSource), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.SecretLocalReference)(nil), (*build.SecretLocalReference)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_SecretLocalReference_To_build_SecretLocalReference(a.(*v1.SecretLocalReference), b.(*build.SecretLocalReference), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*build.SecretLocalReference)(nil), (*v1.SecretLocalReference)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_build_SecretLocalReference_To_v1_SecretLocalReference(a.(*build.SecretLocalReference), b.(*v1.SecretLocalReference), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.SecretSpec)(nil), (*build.SecretSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_SecretSpec_To_build_SecretSpec(a.(*v1.SecretSpec), b.(*build.SecretSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*build.SecretSpec)(nil), (*v1.SecretSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_build_SecretSpec_To_v1_SecretSpec(a.(*build.SecretSpec), b.(*v1.SecretSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.SourceBuildStrategy)(nil), (*build.SourceBuildStrategy)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_SourceBuildStrategy_To_build_SourceBuildStrategy(a.(*v1.SourceBuildStrategy), b.(*build.SourceBuildStrategy), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*build.SourceBuildStrategy)(nil), (*v1.SourceBuildStrategy)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_build_SourceBuildStrategy_To_v1_SourceBuildStrategy(a.(*build.SourceBuildStrategy), b.(*v1.SourceBuildStrategy), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.SourceControlUser)(nil), (*build.SourceControlUser)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_SourceControlUser_To_build_SourceControlUser(a.(*v1.SourceControlUser), b.(*build.SourceControlUser), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*build.SourceControlUser)(nil), (*v1.SourceControlUser)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_build_SourceControlUser_To_v1_SourceControlUser(a.(*build.SourceControlUser), b.(*v1.SourceControlUser), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.SourceRevision)(nil), (*build.SourceRevision)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_SourceRevision_To_build_SourceRevision(a.(*v1.SourceRevision), b.(*build.SourceRevision), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*build.SourceRevision)(nil), (*v1.SourceRevision)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_build_SourceRevision_To_v1_SourceRevision(a.(*build.SourceRevision), b.(*v1.SourceRevision), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.SourceStrategyOptions)(nil), (*build.SourceStrategyOptions)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_SourceStrategyOptions_To_build_SourceStrategyOptions(a.(*v1.SourceStrategyOptions), b.(*build.SourceStrategyOptions), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*build.SourceStrategyOptions)(nil), (*v1.SourceStrategyOptions)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_build_SourceStrategyOptions_To_v1_SourceStrategyOptions(a.(*build.SourceStrategyOptions), b.(*v1.SourceStrategyOptions), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.StageInfo)(nil), (*build.StageInfo)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_StageInfo_To_build_StageInfo(a.(*v1.StageInfo), b.(*build.StageInfo), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*build.StageInfo)(nil), (*v1.StageInfo)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_build_StageInfo_To_v1_StageInfo(a.(*build.StageInfo), b.(*v1.StageInfo), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.StepInfo)(nil), (*build.StepInfo)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_StepInfo_To_build_StepInfo(a.(*v1.StepInfo), b.(*build.StepInfo), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*build.StepInfo)(nil), (*v1.StepInfo)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_build_StepInfo_To_v1_StepInfo(a.(*build.StepInfo), b.(*v1.StepInfo), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.WebHookTrigger)(nil), (*build.WebHookTrigger)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_WebHookTrigger_To_build_WebHookTrigger(a.(*v1.WebHookTrigger), b.(*build.WebHookTrigger), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*build.WebHookTrigger)(nil), (*v1.WebHookTrigger)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_build_WebHookTrigger_To_v1_WebHookTrigger(a.(*build.WebHookTrigger), b.(*v1.WebHookTrigger), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*build.BuildSource)(nil), (*v1.BuildSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_build_BuildSource_To_v1_BuildSource(a.(*build.BuildSource), b.(*v1.BuildSource), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*build.BuildStrategy)(nil), (*v1.BuildStrategy)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_build_BuildStrategy_To_v1_BuildStrategy(a.(*build.BuildStrategy), b.(*v1.BuildStrategy), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*build.SourceRevision)(nil), (*v1.SourceRevision)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_build_SourceRevision_To_v1_SourceRevision(a.(*build.SourceRevision), b.(*v1.SourceRevision), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*v1.BuildConfig)(nil), (*build.BuildConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_BuildConfig_To_build_BuildConfig(a.(*v1.BuildConfig), b.(*build.BuildConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*v1.BuildOutput)(nil), (*build.BuildOutput)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_BuildOutput_To_build_BuildOutput(a.(*v1.BuildOutput), b.(*build.BuildOutput), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*v1.BuildTriggerPolicy)(nil), (*build.BuildTriggerPolicy)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_BuildTriggerPolicy_To_build_BuildTriggerPolicy(a.(*v1.BuildTriggerPolicy), b.(*build.BuildTriggerPolicy), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*v1.CustomBuildStrategy)(nil), (*build.CustomBuildStrategy)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_CustomBuildStrategy_To_build_CustomBuildStrategy(a.(*v1.CustomBuildStrategy), b.(*build.CustomBuildStrategy), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*v1.DockerBuildStrategy)(nil), (*build.DockerBuildStrategy)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_DockerBuildStrategy_To_build_DockerBuildStrategy(a.(*v1.DockerBuildStrategy), b.(*build.DockerBuildStrategy), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*v1.SourceBuildStrategy)(nil), (*build.SourceBuildStrategy)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_SourceBuildStrategy_To_build_SourceBuildStrategy(a.(*v1.SourceBuildStrategy), b.(*build.SourceBuildStrategy), scope)
	}); err != nil {
		return err
	}
	return nil
}
func autoConvert_v1_BinaryBuildRequestOptions_To_build_BinaryBuildRequestOptions(in *v1.BinaryBuildRequestOptions, out *build.BinaryBuildRequestOptions, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ObjectMeta = in.ObjectMeta
	out.AsFile = in.AsFile
	out.Commit = in.Commit
	out.Message = in.Message
	out.AuthorName = in.AuthorName
	out.AuthorEmail = in.AuthorEmail
	out.CommitterName = in.CommitterName
	out.CommitterEmail = in.CommitterEmail
	return nil
}
func Convert_v1_BinaryBuildRequestOptions_To_build_BinaryBuildRequestOptions(in *v1.BinaryBuildRequestOptions, out *build.BinaryBuildRequestOptions, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_BinaryBuildRequestOptions_To_build_BinaryBuildRequestOptions(in, out, s)
}
func autoConvert_build_BinaryBuildRequestOptions_To_v1_BinaryBuildRequestOptions(in *build.BinaryBuildRequestOptions, out *v1.BinaryBuildRequestOptions, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ObjectMeta = in.ObjectMeta
	out.AsFile = in.AsFile
	out.Commit = in.Commit
	out.Message = in.Message
	out.AuthorName = in.AuthorName
	out.AuthorEmail = in.AuthorEmail
	out.CommitterName = in.CommitterName
	out.CommitterEmail = in.CommitterEmail
	return nil
}
func Convert_build_BinaryBuildRequestOptions_To_v1_BinaryBuildRequestOptions(in *build.BinaryBuildRequestOptions, out *v1.BinaryBuildRequestOptions, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_build_BinaryBuildRequestOptions_To_v1_BinaryBuildRequestOptions(in, out, s)
}
func autoConvert_v1_BinaryBuildSource_To_build_BinaryBuildSource(in *v1.BinaryBuildSource, out *build.BinaryBuildSource, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.AsFile = in.AsFile
	return nil
}
func Convert_v1_BinaryBuildSource_To_build_BinaryBuildSource(in *v1.BinaryBuildSource, out *build.BinaryBuildSource, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_BinaryBuildSource_To_build_BinaryBuildSource(in, out, s)
}
func autoConvert_build_BinaryBuildSource_To_v1_BinaryBuildSource(in *build.BinaryBuildSource, out *v1.BinaryBuildSource, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.AsFile = in.AsFile
	return nil
}
func Convert_build_BinaryBuildSource_To_v1_BinaryBuildSource(in *build.BinaryBuildSource, out *v1.BinaryBuildSource, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_build_BinaryBuildSource_To_v1_BinaryBuildSource(in, out, s)
}
func autoConvert_v1_BitbucketWebHookCause_To_build_BitbucketWebHookCause(in *v1.BitbucketWebHookCause, out *build.BitbucketWebHookCause, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := Convert_v1_CommonWebHookCause_To_build_CommonWebHookCause(&in.CommonWebHookCause, &out.CommonWebHookCause, s); err != nil {
		return err
	}
	return nil
}
func Convert_v1_BitbucketWebHookCause_To_build_BitbucketWebHookCause(in *v1.BitbucketWebHookCause, out *build.BitbucketWebHookCause, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_BitbucketWebHookCause_To_build_BitbucketWebHookCause(in, out, s)
}
func autoConvert_build_BitbucketWebHookCause_To_v1_BitbucketWebHookCause(in *build.BitbucketWebHookCause, out *v1.BitbucketWebHookCause, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := Convert_build_CommonWebHookCause_To_v1_CommonWebHookCause(&in.CommonWebHookCause, &out.CommonWebHookCause, s); err != nil {
		return err
	}
	return nil
}
func Convert_build_BitbucketWebHookCause_To_v1_BitbucketWebHookCause(in *build.BitbucketWebHookCause, out *v1.BitbucketWebHookCause, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_build_BitbucketWebHookCause_To_v1_BitbucketWebHookCause(in, out, s)
}
func autoConvert_v1_Build_To_build_Build(in *v1.Build, out *build.Build, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1_BuildSpec_To_build_BuildSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1_BuildStatus_To_build_BuildStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}
func Convert_v1_Build_To_build_Build(in *v1.Build, out *build.Build, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_Build_To_build_Build(in, out, s)
}
func autoConvert_build_Build_To_v1_Build(in *build.Build, out *v1.Build, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_build_BuildSpec_To_v1_BuildSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_build_BuildStatus_To_v1_BuildStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}
func Convert_build_Build_To_v1_Build(in *build.Build, out *v1.Build, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_build_Build_To_v1_Build(in, out, s)
}
func autoConvert_v1_BuildConfig_To_build_BuildConfig(in *v1.BuildConfig, out *build.BuildConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1_BuildConfigSpec_To_build_BuildConfigSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1_BuildConfigStatus_To_build_BuildConfigStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}
func autoConvert_build_BuildConfig_To_v1_BuildConfig(in *build.BuildConfig, out *v1.BuildConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_build_BuildConfigSpec_To_v1_BuildConfigSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_build_BuildConfigStatus_To_v1_BuildConfigStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}
func Convert_build_BuildConfig_To_v1_BuildConfig(in *build.BuildConfig, out *v1.BuildConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_build_BuildConfig_To_v1_BuildConfig(in, out, s)
}
func autoConvert_v1_BuildConfigList_To_build_BuildConfigList(in *v1.BuildConfigList, out *build.BuildConfigList, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]build.BuildConfig, len(*in))
		for i := range *in {
			if err := Convert_v1_BuildConfig_To_build_BuildConfig(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}
func Convert_v1_BuildConfigList_To_build_BuildConfigList(in *v1.BuildConfigList, out *build.BuildConfigList, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_BuildConfigList_To_build_BuildConfigList(in, out, s)
}
func autoConvert_build_BuildConfigList_To_v1_BuildConfigList(in *build.BuildConfigList, out *v1.BuildConfigList, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]v1.BuildConfig, len(*in))
		for i := range *in {
			if err := Convert_build_BuildConfig_To_v1_BuildConfig(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}
func Convert_build_BuildConfigList_To_v1_BuildConfigList(in *build.BuildConfigList, out *v1.BuildConfigList, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_build_BuildConfigList_To_v1_BuildConfigList(in, out, s)
}
func autoConvert_v1_BuildConfigSpec_To_build_BuildConfigSpec(in *v1.BuildConfigSpec, out *build.BuildConfigSpec, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in.Triggers != nil {
		in, out := &in.Triggers, &out.Triggers
		*out = make([]build.BuildTriggerPolicy, len(*in))
		for i := range *in {
			if err := Convert_v1_BuildTriggerPolicy_To_build_BuildTriggerPolicy(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Triggers = nil
	}
	out.RunPolicy = build.BuildRunPolicy(in.RunPolicy)
	if err := Convert_v1_CommonSpec_To_build_CommonSpec(&in.CommonSpec, &out.CommonSpec, s); err != nil {
		return err
	}
	out.SuccessfulBuildsHistoryLimit = (*int32)(unsafe.Pointer(in.SuccessfulBuildsHistoryLimit))
	out.FailedBuildsHistoryLimit = (*int32)(unsafe.Pointer(in.FailedBuildsHistoryLimit))
	return nil
}
func Convert_v1_BuildConfigSpec_To_build_BuildConfigSpec(in *v1.BuildConfigSpec, out *build.BuildConfigSpec, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_BuildConfigSpec_To_build_BuildConfigSpec(in, out, s)
}
func autoConvert_build_BuildConfigSpec_To_v1_BuildConfigSpec(in *build.BuildConfigSpec, out *v1.BuildConfigSpec, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in.Triggers != nil {
		in, out := &in.Triggers, &out.Triggers
		*out = make([]v1.BuildTriggerPolicy, len(*in))
		for i := range *in {
			if err := Convert_build_BuildTriggerPolicy_To_v1_BuildTriggerPolicy(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Triggers = nil
	}
	out.RunPolicy = v1.BuildRunPolicy(in.RunPolicy)
	if err := Convert_build_CommonSpec_To_v1_CommonSpec(&in.CommonSpec, &out.CommonSpec, s); err != nil {
		return err
	}
	out.SuccessfulBuildsHistoryLimit = (*int32)(unsafe.Pointer(in.SuccessfulBuildsHistoryLimit))
	out.FailedBuildsHistoryLimit = (*int32)(unsafe.Pointer(in.FailedBuildsHistoryLimit))
	return nil
}
func Convert_build_BuildConfigSpec_To_v1_BuildConfigSpec(in *build.BuildConfigSpec, out *v1.BuildConfigSpec, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_build_BuildConfigSpec_To_v1_BuildConfigSpec(in, out, s)
}
func autoConvert_v1_BuildConfigStatus_To_build_BuildConfigStatus(in *v1.BuildConfigStatus, out *build.BuildConfigStatus, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.LastVersion = in.LastVersion
	return nil
}
func Convert_v1_BuildConfigStatus_To_build_BuildConfigStatus(in *v1.BuildConfigStatus, out *build.BuildConfigStatus, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_BuildConfigStatus_To_build_BuildConfigStatus(in, out, s)
}
func autoConvert_build_BuildConfigStatus_To_v1_BuildConfigStatus(in *build.BuildConfigStatus, out *v1.BuildConfigStatus, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.LastVersion = in.LastVersion
	return nil
}
func Convert_build_BuildConfigStatus_To_v1_BuildConfigStatus(in *build.BuildConfigStatus, out *v1.BuildConfigStatus, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_build_BuildConfigStatus_To_v1_BuildConfigStatus(in, out, s)
}
func autoConvert_v1_BuildList_To_build_BuildList(in *v1.BuildList, out *build.BuildList, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]build.Build, len(*in))
		for i := range *in {
			if err := Convert_v1_Build_To_build_Build(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}
func Convert_v1_BuildList_To_build_BuildList(in *v1.BuildList, out *build.BuildList, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_BuildList_To_build_BuildList(in, out, s)
}
func autoConvert_build_BuildList_To_v1_BuildList(in *build.BuildList, out *v1.BuildList, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]v1.Build, len(*in))
		for i := range *in {
			if err := Convert_build_Build_To_v1_Build(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}
func Convert_build_BuildList_To_v1_BuildList(in *build.BuildList, out *v1.BuildList, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_build_BuildList_To_v1_BuildList(in, out, s)
}
func autoConvert_v1_BuildLog_To_build_BuildLog(in *v1.BuildLog, out *build.BuildLog, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return nil
}
func Convert_v1_BuildLog_To_build_BuildLog(in *v1.BuildLog, out *build.BuildLog, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_BuildLog_To_build_BuildLog(in, out, s)
}
func autoConvert_build_BuildLog_To_v1_BuildLog(in *build.BuildLog, out *v1.BuildLog, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return nil
}
func Convert_build_BuildLog_To_v1_BuildLog(in *build.BuildLog, out *v1.BuildLog, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_build_BuildLog_To_v1_BuildLog(in, out, s)
}
func autoConvert_v1_BuildLogOptions_To_build_BuildLogOptions(in *v1.BuildLogOptions, out *build.BuildLogOptions, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Container = in.Container
	out.Follow = in.Follow
	out.Previous = in.Previous
	out.SinceSeconds = (*int64)(unsafe.Pointer(in.SinceSeconds))
	out.SinceTime = (*metav1.Time)(unsafe.Pointer(in.SinceTime))
	out.Timestamps = in.Timestamps
	out.TailLines = (*int64)(unsafe.Pointer(in.TailLines))
	out.LimitBytes = (*int64)(unsafe.Pointer(in.LimitBytes))
	out.NoWait = in.NoWait
	out.Version = (*int64)(unsafe.Pointer(in.Version))
	return nil
}
func Convert_v1_BuildLogOptions_To_build_BuildLogOptions(in *v1.BuildLogOptions, out *build.BuildLogOptions, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_BuildLogOptions_To_build_BuildLogOptions(in, out, s)
}
func autoConvert_build_BuildLogOptions_To_v1_BuildLogOptions(in *build.BuildLogOptions, out *v1.BuildLogOptions, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Container = in.Container
	out.Follow = in.Follow
	out.Previous = in.Previous
	out.SinceSeconds = (*int64)(unsafe.Pointer(in.SinceSeconds))
	out.SinceTime = (*metav1.Time)(unsafe.Pointer(in.SinceTime))
	out.Timestamps = in.Timestamps
	out.TailLines = (*int64)(unsafe.Pointer(in.TailLines))
	out.LimitBytes = (*int64)(unsafe.Pointer(in.LimitBytes))
	out.NoWait = in.NoWait
	out.Version = (*int64)(unsafe.Pointer(in.Version))
	return nil
}
func Convert_build_BuildLogOptions_To_v1_BuildLogOptions(in *build.BuildLogOptions, out *v1.BuildLogOptions, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_build_BuildLogOptions_To_v1_BuildLogOptions(in, out, s)
}
func autoConvert_v1_BuildOutput_To_build_BuildOutput(in *v1.BuildOutput, out *build.BuildOutput, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in.To != nil {
		in, out := &in.To, &out.To
		*out = new(core.ObjectReference)
		if err := corev1.Convert_v1_ObjectReference_To_core_ObjectReference(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.To = nil
	}
	if in.PushSecret != nil {
		in, out := &in.PushSecret, &out.PushSecret
		*out = new(core.LocalObjectReference)
		if err := corev1.Convert_v1_LocalObjectReference_To_core_LocalObjectReference(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.PushSecret = nil
	}
	out.ImageLabels = *(*[]build.ImageLabel)(unsafe.Pointer(&in.ImageLabels))
	return nil
}
func autoConvert_build_BuildOutput_To_v1_BuildOutput(in *build.BuildOutput, out *v1.BuildOutput, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in.To != nil {
		in, out := &in.To, &out.To
		*out = new(apicorev1.ObjectReference)
		if err := corev1.Convert_core_ObjectReference_To_v1_ObjectReference(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.To = nil
	}
	if in.PushSecret != nil {
		in, out := &in.PushSecret, &out.PushSecret
		*out = new(apicorev1.LocalObjectReference)
		if err := corev1.Convert_core_LocalObjectReference_To_v1_LocalObjectReference(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.PushSecret = nil
	}
	out.ImageLabels = *(*[]v1.ImageLabel)(unsafe.Pointer(&in.ImageLabels))
	return nil
}
func Convert_build_BuildOutput_To_v1_BuildOutput(in *build.BuildOutput, out *v1.BuildOutput, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_build_BuildOutput_To_v1_BuildOutput(in, out, s)
}
func autoConvert_v1_BuildPostCommitSpec_To_build_BuildPostCommitSpec(in *v1.BuildPostCommitSpec, out *build.BuildPostCommitSpec, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Command = *(*[]string)(unsafe.Pointer(&in.Command))
	out.Args = *(*[]string)(unsafe.Pointer(&in.Args))
	out.Script = in.Script
	return nil
}
func Convert_v1_BuildPostCommitSpec_To_build_BuildPostCommitSpec(in *v1.BuildPostCommitSpec, out *build.BuildPostCommitSpec, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_BuildPostCommitSpec_To_build_BuildPostCommitSpec(in, out, s)
}
func autoConvert_build_BuildPostCommitSpec_To_v1_BuildPostCommitSpec(in *build.BuildPostCommitSpec, out *v1.BuildPostCommitSpec, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Command = *(*[]string)(unsafe.Pointer(&in.Command))
	out.Args = *(*[]string)(unsafe.Pointer(&in.Args))
	out.Script = in.Script
	return nil
}
func Convert_build_BuildPostCommitSpec_To_v1_BuildPostCommitSpec(in *build.BuildPostCommitSpec, out *v1.BuildPostCommitSpec, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_build_BuildPostCommitSpec_To_v1_BuildPostCommitSpec(in, out, s)
}
func autoConvert_v1_BuildRequest_To_build_BuildRequest(in *v1.BuildRequest, out *build.BuildRequest, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ObjectMeta = in.ObjectMeta
	if in.Revision != nil {
		in, out := &in.Revision, &out.Revision
		*out = new(build.SourceRevision)
		if err := Convert_v1_SourceRevision_To_build_SourceRevision(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.Revision = nil
	}
	if in.TriggeredByImage != nil {
		in, out := &in.TriggeredByImage, &out.TriggeredByImage
		*out = new(core.ObjectReference)
		if err := corev1.Convert_v1_ObjectReference_To_core_ObjectReference(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.TriggeredByImage = nil
	}
	if in.From != nil {
		in, out := &in.From, &out.From
		*out = new(core.ObjectReference)
		if err := corev1.Convert_v1_ObjectReference_To_core_ObjectReference(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.From = nil
	}
	out.Binary = (*build.BinaryBuildSource)(unsafe.Pointer(in.Binary))
	out.LastVersion = (*int64)(unsafe.Pointer(in.LastVersion))
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]core.EnvVar, len(*in))
		for i := range *in {
			if err := corev1.Convert_v1_EnvVar_To_core_EnvVar(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Env = nil
	}
	if in.TriggeredBy != nil {
		in, out := &in.TriggeredBy, &out.TriggeredBy
		*out = make([]build.BuildTriggerCause, len(*in))
		for i := range *in {
			if err := Convert_v1_BuildTriggerCause_To_build_BuildTriggerCause(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.TriggeredBy = nil
	}
	if in.DockerStrategyOptions != nil {
		in, out := &in.DockerStrategyOptions, &out.DockerStrategyOptions
		*out = new(build.DockerStrategyOptions)
		if err := Convert_v1_DockerStrategyOptions_To_build_DockerStrategyOptions(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.DockerStrategyOptions = nil
	}
	out.SourceStrategyOptions = (*build.SourceStrategyOptions)(unsafe.Pointer(in.SourceStrategyOptions))
	return nil
}
func Convert_v1_BuildRequest_To_build_BuildRequest(in *v1.BuildRequest, out *build.BuildRequest, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_BuildRequest_To_build_BuildRequest(in, out, s)
}
func autoConvert_build_BuildRequest_To_v1_BuildRequest(in *build.BuildRequest, out *v1.BuildRequest, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ObjectMeta = in.ObjectMeta
	if in.Revision != nil {
		in, out := &in.Revision, &out.Revision
		*out = new(v1.SourceRevision)
		if err := Convert_build_SourceRevision_To_v1_SourceRevision(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.Revision = nil
	}
	if in.TriggeredByImage != nil {
		in, out := &in.TriggeredByImage, &out.TriggeredByImage
		*out = new(apicorev1.ObjectReference)
		if err := corev1.Convert_core_ObjectReference_To_v1_ObjectReference(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.TriggeredByImage = nil
	}
	if in.From != nil {
		in, out := &in.From, &out.From
		*out = new(apicorev1.ObjectReference)
		if err := corev1.Convert_core_ObjectReference_To_v1_ObjectReference(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.From = nil
	}
	out.Binary = (*v1.BinaryBuildSource)(unsafe.Pointer(in.Binary))
	out.LastVersion = (*int64)(unsafe.Pointer(in.LastVersion))
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]apicorev1.EnvVar, len(*in))
		for i := range *in {
			if err := corev1.Convert_core_EnvVar_To_v1_EnvVar(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Env = nil
	}
	if in.TriggeredBy != nil {
		in, out := &in.TriggeredBy, &out.TriggeredBy
		*out = make([]v1.BuildTriggerCause, len(*in))
		for i := range *in {
			if err := Convert_build_BuildTriggerCause_To_v1_BuildTriggerCause(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.TriggeredBy = nil
	}
	if in.DockerStrategyOptions != nil {
		in, out := &in.DockerStrategyOptions, &out.DockerStrategyOptions
		*out = new(v1.DockerStrategyOptions)
		if err := Convert_build_DockerStrategyOptions_To_v1_DockerStrategyOptions(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.DockerStrategyOptions = nil
	}
	out.SourceStrategyOptions = (*v1.SourceStrategyOptions)(unsafe.Pointer(in.SourceStrategyOptions))
	return nil
}
func Convert_build_BuildRequest_To_v1_BuildRequest(in *build.BuildRequest, out *v1.BuildRequest, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_build_BuildRequest_To_v1_BuildRequest(in, out, s)
}
func autoConvert_v1_BuildSource_To_build_BuildSource(in *v1.BuildSource, out *build.BuildSource, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Binary = (*build.BinaryBuildSource)(unsafe.Pointer(in.Binary))
	out.Dockerfile = (*string)(unsafe.Pointer(in.Dockerfile))
	out.Git = (*build.GitBuildSource)(unsafe.Pointer(in.Git))
	if in.Images != nil {
		in, out := &in.Images, &out.Images
		*out = make([]build.ImageSource, len(*in))
		for i := range *in {
			if err := Convert_v1_ImageSource_To_build_ImageSource(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Images = nil
	}
	out.ContextDir = in.ContextDir
	if in.SourceSecret != nil {
		in, out := &in.SourceSecret, &out.SourceSecret
		*out = new(core.LocalObjectReference)
		if err := corev1.Convert_v1_LocalObjectReference_To_core_LocalObjectReference(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.SourceSecret = nil
	}
	if in.Secrets != nil {
		in, out := &in.Secrets, &out.Secrets
		*out = make([]build.SecretBuildSource, len(*in))
		for i := range *in {
			if err := Convert_v1_SecretBuildSource_To_build_SecretBuildSource(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Secrets = nil
	}
	if in.ConfigMaps != nil {
		in, out := &in.ConfigMaps, &out.ConfigMaps
		*out = make([]build.ConfigMapBuildSource, len(*in))
		for i := range *in {
			if err := Convert_v1_ConfigMapBuildSource_To_build_ConfigMapBuildSource(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.ConfigMaps = nil
	}
	return nil
}
func Convert_v1_BuildSource_To_build_BuildSource(in *v1.BuildSource, out *build.BuildSource, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_BuildSource_To_build_BuildSource(in, out, s)
}
func autoConvert_build_BuildSource_To_v1_BuildSource(in *build.BuildSource, out *v1.BuildSource, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Binary = (*v1.BinaryBuildSource)(unsafe.Pointer(in.Binary))
	out.Dockerfile = (*string)(unsafe.Pointer(in.Dockerfile))
	out.Git = (*v1.GitBuildSource)(unsafe.Pointer(in.Git))
	if in.Images != nil {
		in, out := &in.Images, &out.Images
		*out = make([]v1.ImageSource, len(*in))
		for i := range *in {
			if err := Convert_build_ImageSource_To_v1_ImageSource(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Images = nil
	}
	out.ContextDir = in.ContextDir
	if in.SourceSecret != nil {
		in, out := &in.SourceSecret, &out.SourceSecret
		*out = new(apicorev1.LocalObjectReference)
		if err := corev1.Convert_core_LocalObjectReference_To_v1_LocalObjectReference(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.SourceSecret = nil
	}
	if in.Secrets != nil {
		in, out := &in.Secrets, &out.Secrets
		*out = make([]v1.SecretBuildSource, len(*in))
		for i := range *in {
			if err := Convert_build_SecretBuildSource_To_v1_SecretBuildSource(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Secrets = nil
	}
	if in.ConfigMaps != nil {
		in, out := &in.ConfigMaps, &out.ConfigMaps
		*out = make([]v1.ConfigMapBuildSource, len(*in))
		for i := range *in {
			if err := Convert_build_ConfigMapBuildSource_To_v1_ConfigMapBuildSource(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.ConfigMaps = nil
	}
	return nil
}
func autoConvert_v1_BuildSpec_To_build_BuildSpec(in *v1.BuildSpec, out *build.BuildSpec, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := Convert_v1_CommonSpec_To_build_CommonSpec(&in.CommonSpec, &out.CommonSpec, s); err != nil {
		return err
	}
	if in.TriggeredBy != nil {
		in, out := &in.TriggeredBy, &out.TriggeredBy
		*out = make([]build.BuildTriggerCause, len(*in))
		for i := range *in {
			if err := Convert_v1_BuildTriggerCause_To_build_BuildTriggerCause(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.TriggeredBy = nil
	}
	return nil
}
func Convert_v1_BuildSpec_To_build_BuildSpec(in *v1.BuildSpec, out *build.BuildSpec, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_BuildSpec_To_build_BuildSpec(in, out, s)
}
func autoConvert_build_BuildSpec_To_v1_BuildSpec(in *build.BuildSpec, out *v1.BuildSpec, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := Convert_build_CommonSpec_To_v1_CommonSpec(&in.CommonSpec, &out.CommonSpec, s); err != nil {
		return err
	}
	if in.TriggeredBy != nil {
		in, out := &in.TriggeredBy, &out.TriggeredBy
		*out = make([]v1.BuildTriggerCause, len(*in))
		for i := range *in {
			if err := Convert_build_BuildTriggerCause_To_v1_BuildTriggerCause(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.TriggeredBy = nil
	}
	return nil
}
func Convert_build_BuildSpec_To_v1_BuildSpec(in *build.BuildSpec, out *v1.BuildSpec, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_build_BuildSpec_To_v1_BuildSpec(in, out, s)
}
func autoConvert_v1_BuildStatus_To_build_BuildStatus(in *v1.BuildStatus, out *build.BuildStatus, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Phase = build.BuildPhase(in.Phase)
	out.Cancelled = in.Cancelled
	out.Reason = build.StatusReason(in.Reason)
	out.Message = in.Message
	out.StartTimestamp = (*metav1.Time)(unsafe.Pointer(in.StartTimestamp))
	out.CompletionTimestamp = (*metav1.Time)(unsafe.Pointer(in.CompletionTimestamp))
	out.Duration = time.Duration(in.Duration)
	out.OutputDockerImageReference = in.OutputDockerImageReference
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = new(core.ObjectReference)
		if err := corev1.Convert_v1_ObjectReference_To_core_ObjectReference(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.Config = nil
	}
	if err := Convert_v1_BuildStatusOutput_To_build_BuildStatusOutput(&in.Output, &out.Output, s); err != nil {
		return err
	}
	out.Stages = *(*[]build.StageInfo)(unsafe.Pointer(&in.Stages))
	out.LogSnippet = in.LogSnippet
	return nil
}
func Convert_v1_BuildStatus_To_build_BuildStatus(in *v1.BuildStatus, out *build.BuildStatus, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_BuildStatus_To_build_BuildStatus(in, out, s)
}
func autoConvert_build_BuildStatus_To_v1_BuildStatus(in *build.BuildStatus, out *v1.BuildStatus, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Phase = v1.BuildPhase(in.Phase)
	out.Cancelled = in.Cancelled
	out.Reason = v1.StatusReason(in.Reason)
	out.Message = in.Message
	out.StartTimestamp = (*metav1.Time)(unsafe.Pointer(in.StartTimestamp))
	out.CompletionTimestamp = (*metav1.Time)(unsafe.Pointer(in.CompletionTimestamp))
	out.Duration = time.Duration(in.Duration)
	out.OutputDockerImageReference = in.OutputDockerImageReference
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = new(apicorev1.ObjectReference)
		if err := corev1.Convert_core_ObjectReference_To_v1_ObjectReference(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.Config = nil
	}
	if err := Convert_build_BuildStatusOutput_To_v1_BuildStatusOutput(&in.Output, &out.Output, s); err != nil {
		return err
	}
	out.Stages = *(*[]v1.StageInfo)(unsafe.Pointer(&in.Stages))
	out.LogSnippet = in.LogSnippet
	return nil
}
func Convert_build_BuildStatus_To_v1_BuildStatus(in *build.BuildStatus, out *v1.BuildStatus, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_build_BuildStatus_To_v1_BuildStatus(in, out, s)
}
func autoConvert_v1_BuildStatusOutput_To_build_BuildStatusOutput(in *v1.BuildStatusOutput, out *build.BuildStatusOutput, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.To = (*build.BuildStatusOutputTo)(unsafe.Pointer(in.To))
	return nil
}
func Convert_v1_BuildStatusOutput_To_build_BuildStatusOutput(in *v1.BuildStatusOutput, out *build.BuildStatusOutput, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_BuildStatusOutput_To_build_BuildStatusOutput(in, out, s)
}
func autoConvert_build_BuildStatusOutput_To_v1_BuildStatusOutput(in *build.BuildStatusOutput, out *v1.BuildStatusOutput, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.To = (*v1.BuildStatusOutputTo)(unsafe.Pointer(in.To))
	return nil
}
func Convert_build_BuildStatusOutput_To_v1_BuildStatusOutput(in *build.BuildStatusOutput, out *v1.BuildStatusOutput, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_build_BuildStatusOutput_To_v1_BuildStatusOutput(in, out, s)
}
func autoConvert_v1_BuildStatusOutputTo_To_build_BuildStatusOutputTo(in *v1.BuildStatusOutputTo, out *build.BuildStatusOutputTo, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ImageDigest = in.ImageDigest
	return nil
}
func Convert_v1_BuildStatusOutputTo_To_build_BuildStatusOutputTo(in *v1.BuildStatusOutputTo, out *build.BuildStatusOutputTo, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_BuildStatusOutputTo_To_build_BuildStatusOutputTo(in, out, s)
}
func autoConvert_build_BuildStatusOutputTo_To_v1_BuildStatusOutputTo(in *build.BuildStatusOutputTo, out *v1.BuildStatusOutputTo, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ImageDigest = in.ImageDigest
	return nil
}
func Convert_build_BuildStatusOutputTo_To_v1_BuildStatusOutputTo(in *build.BuildStatusOutputTo, out *v1.BuildStatusOutputTo, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_build_BuildStatusOutputTo_To_v1_BuildStatusOutputTo(in, out, s)
}
func autoConvert_v1_BuildStrategy_To_build_BuildStrategy(in *v1.BuildStrategy, out *build.BuildStrategy, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in.DockerStrategy != nil {
		in, out := &in.DockerStrategy, &out.DockerStrategy
		*out = new(build.DockerBuildStrategy)
		if err := Convert_v1_DockerBuildStrategy_To_build_DockerBuildStrategy(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.DockerStrategy = nil
	}
	if in.SourceStrategy != nil {
		in, out := &in.SourceStrategy, &out.SourceStrategy
		*out = new(build.SourceBuildStrategy)
		if err := Convert_v1_SourceBuildStrategy_To_build_SourceBuildStrategy(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.SourceStrategy = nil
	}
	if in.CustomStrategy != nil {
		in, out := &in.CustomStrategy, &out.CustomStrategy
		*out = new(build.CustomBuildStrategy)
		if err := Convert_v1_CustomBuildStrategy_To_build_CustomBuildStrategy(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.CustomStrategy = nil
	}
	if in.JenkinsPipelineStrategy != nil {
		in, out := &in.JenkinsPipelineStrategy, &out.JenkinsPipelineStrategy
		*out = new(build.JenkinsPipelineBuildStrategy)
		if err := Convert_v1_JenkinsPipelineBuildStrategy_To_build_JenkinsPipelineBuildStrategy(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.JenkinsPipelineStrategy = nil
	}
	return nil
}
func Convert_v1_BuildStrategy_To_build_BuildStrategy(in *v1.BuildStrategy, out *build.BuildStrategy, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_BuildStrategy_To_build_BuildStrategy(in, out, s)
}
func autoConvert_build_BuildStrategy_To_v1_BuildStrategy(in *build.BuildStrategy, out *v1.BuildStrategy, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in.DockerStrategy != nil {
		in, out := &in.DockerStrategy, &out.DockerStrategy
		*out = new(v1.DockerBuildStrategy)
		if err := Convert_build_DockerBuildStrategy_To_v1_DockerBuildStrategy(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.DockerStrategy = nil
	}
	if in.SourceStrategy != nil {
		in, out := &in.SourceStrategy, &out.SourceStrategy
		*out = new(v1.SourceBuildStrategy)
		if err := Convert_build_SourceBuildStrategy_To_v1_SourceBuildStrategy(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.SourceStrategy = nil
	}
	if in.CustomStrategy != nil {
		in, out := &in.CustomStrategy, &out.CustomStrategy
		*out = new(v1.CustomBuildStrategy)
		if err := Convert_build_CustomBuildStrategy_To_v1_CustomBuildStrategy(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.CustomStrategy = nil
	}
	if in.JenkinsPipelineStrategy != nil {
		in, out := &in.JenkinsPipelineStrategy, &out.JenkinsPipelineStrategy
		*out = new(v1.JenkinsPipelineBuildStrategy)
		if err := Convert_build_JenkinsPipelineBuildStrategy_To_v1_JenkinsPipelineBuildStrategy(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.JenkinsPipelineStrategy = nil
	}
	return nil
}
func autoConvert_v1_BuildTriggerCause_To_build_BuildTriggerCause(in *v1.BuildTriggerCause, out *build.BuildTriggerCause, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Message = in.Message
	if in.GenericWebHook != nil {
		in, out := &in.GenericWebHook, &out.GenericWebHook
		*out = new(build.GenericWebHookCause)
		if err := Convert_v1_GenericWebHookCause_To_build_GenericWebHookCause(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.GenericWebHook = nil
	}
	if in.GitHubWebHook != nil {
		in, out := &in.GitHubWebHook, &out.GitHubWebHook
		*out = new(build.GitHubWebHookCause)
		if err := Convert_v1_GitHubWebHookCause_To_build_GitHubWebHookCause(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.GitHubWebHook = nil
	}
	if in.ImageChangeBuild != nil {
		in, out := &in.ImageChangeBuild, &out.ImageChangeBuild
		*out = new(build.ImageChangeCause)
		if err := Convert_v1_ImageChangeCause_To_build_ImageChangeCause(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.ImageChangeBuild = nil
	}
	if in.GitLabWebHook != nil {
		in, out := &in.GitLabWebHook, &out.GitLabWebHook
		*out = new(build.GitLabWebHookCause)
		if err := Convert_v1_GitLabWebHookCause_To_build_GitLabWebHookCause(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.GitLabWebHook = nil
	}
	if in.BitbucketWebHook != nil {
		in, out := &in.BitbucketWebHook, &out.BitbucketWebHook
		*out = new(build.BitbucketWebHookCause)
		if err := Convert_v1_BitbucketWebHookCause_To_build_BitbucketWebHookCause(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.BitbucketWebHook = nil
	}
	return nil
}
func Convert_v1_BuildTriggerCause_To_build_BuildTriggerCause(in *v1.BuildTriggerCause, out *build.BuildTriggerCause, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_BuildTriggerCause_To_build_BuildTriggerCause(in, out, s)
}
func autoConvert_build_BuildTriggerCause_To_v1_BuildTriggerCause(in *build.BuildTriggerCause, out *v1.BuildTriggerCause, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Message = in.Message
	if in.GenericWebHook != nil {
		in, out := &in.GenericWebHook, &out.GenericWebHook
		*out = new(v1.GenericWebHookCause)
		if err := Convert_build_GenericWebHookCause_To_v1_GenericWebHookCause(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.GenericWebHook = nil
	}
	if in.GitHubWebHook != nil {
		in, out := &in.GitHubWebHook, &out.GitHubWebHook
		*out = new(v1.GitHubWebHookCause)
		if err := Convert_build_GitHubWebHookCause_To_v1_GitHubWebHookCause(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.GitHubWebHook = nil
	}
	if in.ImageChangeBuild != nil {
		in, out := &in.ImageChangeBuild, &out.ImageChangeBuild
		*out = new(v1.ImageChangeCause)
		if err := Convert_build_ImageChangeCause_To_v1_ImageChangeCause(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.ImageChangeBuild = nil
	}
	if in.GitLabWebHook != nil {
		in, out := &in.GitLabWebHook, &out.GitLabWebHook
		*out = new(v1.GitLabWebHookCause)
		if err := Convert_build_GitLabWebHookCause_To_v1_GitLabWebHookCause(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.GitLabWebHook = nil
	}
	if in.BitbucketWebHook != nil {
		in, out := &in.BitbucketWebHook, &out.BitbucketWebHook
		*out = new(v1.BitbucketWebHookCause)
		if err := Convert_build_BitbucketWebHookCause_To_v1_BitbucketWebHookCause(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.BitbucketWebHook = nil
	}
	return nil
}
func Convert_build_BuildTriggerCause_To_v1_BuildTriggerCause(in *build.BuildTriggerCause, out *v1.BuildTriggerCause, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_build_BuildTriggerCause_To_v1_BuildTriggerCause(in, out, s)
}
func autoConvert_v1_BuildTriggerPolicy_To_build_BuildTriggerPolicy(in *v1.BuildTriggerPolicy, out *build.BuildTriggerPolicy, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Type = build.BuildTriggerType(in.Type)
	out.GitHubWebHook = (*build.WebHookTrigger)(unsafe.Pointer(in.GitHubWebHook))
	out.GenericWebHook = (*build.WebHookTrigger)(unsafe.Pointer(in.GenericWebHook))
	if in.ImageChange != nil {
		in, out := &in.ImageChange, &out.ImageChange
		*out = new(build.ImageChangeTrigger)
		if err := Convert_v1_ImageChangeTrigger_To_build_ImageChangeTrigger(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.ImageChange = nil
	}
	out.GitLabWebHook = (*build.WebHookTrigger)(unsafe.Pointer(in.GitLabWebHook))
	out.BitbucketWebHook = (*build.WebHookTrigger)(unsafe.Pointer(in.BitbucketWebHook))
	return nil
}
func autoConvert_build_BuildTriggerPolicy_To_v1_BuildTriggerPolicy(in *build.BuildTriggerPolicy, out *v1.BuildTriggerPolicy, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Type = v1.BuildTriggerType(in.Type)
	out.GitHubWebHook = (*v1.WebHookTrigger)(unsafe.Pointer(in.GitHubWebHook))
	out.GenericWebHook = (*v1.WebHookTrigger)(unsafe.Pointer(in.GenericWebHook))
	if in.ImageChange != nil {
		in, out := &in.ImageChange, &out.ImageChange
		*out = new(v1.ImageChangeTrigger)
		if err := Convert_build_ImageChangeTrigger_To_v1_ImageChangeTrigger(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.ImageChange = nil
	}
	out.GitLabWebHook = (*v1.WebHookTrigger)(unsafe.Pointer(in.GitLabWebHook))
	out.BitbucketWebHook = (*v1.WebHookTrigger)(unsafe.Pointer(in.BitbucketWebHook))
	return nil
}
func Convert_build_BuildTriggerPolicy_To_v1_BuildTriggerPolicy(in *build.BuildTriggerPolicy, out *v1.BuildTriggerPolicy, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_build_BuildTriggerPolicy_To_v1_BuildTriggerPolicy(in, out, s)
}
func autoConvert_v1_CommonSpec_To_build_CommonSpec(in *v1.CommonSpec, out *build.CommonSpec, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ServiceAccount = in.ServiceAccount
	if err := Convert_v1_BuildSource_To_build_BuildSource(&in.Source, &out.Source, s); err != nil {
		return err
	}
	if in.Revision != nil {
		in, out := &in.Revision, &out.Revision
		*out = new(build.SourceRevision)
		if err := Convert_v1_SourceRevision_To_build_SourceRevision(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.Revision = nil
	}
	if err := Convert_v1_BuildStrategy_To_build_BuildStrategy(&in.Strategy, &out.Strategy, s); err != nil {
		return err
	}
	if err := Convert_v1_BuildOutput_To_build_BuildOutput(&in.Output, &out.Output, s); err != nil {
		return err
	}
	if err := corev1.Convert_v1_ResourceRequirements_To_core_ResourceRequirements(&in.Resources, &out.Resources, s); err != nil {
		return err
	}
	if err := Convert_v1_BuildPostCommitSpec_To_build_BuildPostCommitSpec(&in.PostCommit, &out.PostCommit, s); err != nil {
		return err
	}
	out.CompletionDeadlineSeconds = (*int64)(unsafe.Pointer(in.CompletionDeadlineSeconds))
	out.NodeSelector = *(*map[string]string)(unsafe.Pointer(&in.NodeSelector))
	return nil
}
func Convert_v1_CommonSpec_To_build_CommonSpec(in *v1.CommonSpec, out *build.CommonSpec, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_CommonSpec_To_build_CommonSpec(in, out, s)
}
func autoConvert_build_CommonSpec_To_v1_CommonSpec(in *build.CommonSpec, out *v1.CommonSpec, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ServiceAccount = in.ServiceAccount
	if err := Convert_build_BuildSource_To_v1_BuildSource(&in.Source, &out.Source, s); err != nil {
		return err
	}
	if in.Revision != nil {
		in, out := &in.Revision, &out.Revision
		*out = new(v1.SourceRevision)
		if err := Convert_build_SourceRevision_To_v1_SourceRevision(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.Revision = nil
	}
	if err := Convert_build_BuildStrategy_To_v1_BuildStrategy(&in.Strategy, &out.Strategy, s); err != nil {
		return err
	}
	if err := Convert_build_BuildOutput_To_v1_BuildOutput(&in.Output, &out.Output, s); err != nil {
		return err
	}
	if err := corev1.Convert_core_ResourceRequirements_To_v1_ResourceRequirements(&in.Resources, &out.Resources, s); err != nil {
		return err
	}
	if err := Convert_build_BuildPostCommitSpec_To_v1_BuildPostCommitSpec(&in.PostCommit, &out.PostCommit, s); err != nil {
		return err
	}
	out.CompletionDeadlineSeconds = (*int64)(unsafe.Pointer(in.CompletionDeadlineSeconds))
	out.NodeSelector = *(*v1.OptionalNodeSelector)(unsafe.Pointer(&in.NodeSelector))
	return nil
}
func Convert_build_CommonSpec_To_v1_CommonSpec(in *build.CommonSpec, out *v1.CommonSpec, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_build_CommonSpec_To_v1_CommonSpec(in, out, s)
}
func autoConvert_v1_CommonWebHookCause_To_build_CommonWebHookCause(in *v1.CommonWebHookCause, out *build.CommonWebHookCause, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in.Revision != nil {
		in, out := &in.Revision, &out.Revision
		*out = new(build.SourceRevision)
		if err := Convert_v1_SourceRevision_To_build_SourceRevision(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.Revision = nil
	}
	out.Secret = in.Secret
	return nil
}
func Convert_v1_CommonWebHookCause_To_build_CommonWebHookCause(in *v1.CommonWebHookCause, out *build.CommonWebHookCause, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_CommonWebHookCause_To_build_CommonWebHookCause(in, out, s)
}
func autoConvert_build_CommonWebHookCause_To_v1_CommonWebHookCause(in *build.CommonWebHookCause, out *v1.CommonWebHookCause, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in.Revision != nil {
		in, out := &in.Revision, &out.Revision
		*out = new(v1.SourceRevision)
		if err := Convert_build_SourceRevision_To_v1_SourceRevision(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.Revision = nil
	}
	out.Secret = in.Secret
	return nil
}
func Convert_build_CommonWebHookCause_To_v1_CommonWebHookCause(in *build.CommonWebHookCause, out *v1.CommonWebHookCause, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_build_CommonWebHookCause_To_v1_CommonWebHookCause(in, out, s)
}
func autoConvert_v1_ConfigMapBuildSource_To_build_ConfigMapBuildSource(in *v1.ConfigMapBuildSource, out *build.ConfigMapBuildSource, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := corev1.Convert_v1_LocalObjectReference_To_core_LocalObjectReference(&in.ConfigMap, &out.ConfigMap, s); err != nil {
		return err
	}
	out.DestinationDir = in.DestinationDir
	return nil
}
func Convert_v1_ConfigMapBuildSource_To_build_ConfigMapBuildSource(in *v1.ConfigMapBuildSource, out *build.ConfigMapBuildSource, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_ConfigMapBuildSource_To_build_ConfigMapBuildSource(in, out, s)
}
func autoConvert_build_ConfigMapBuildSource_To_v1_ConfigMapBuildSource(in *build.ConfigMapBuildSource, out *v1.ConfigMapBuildSource, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := corev1.Convert_core_LocalObjectReference_To_v1_LocalObjectReference(&in.ConfigMap, &out.ConfigMap, s); err != nil {
		return err
	}
	out.DestinationDir = in.DestinationDir
	return nil
}
func Convert_build_ConfigMapBuildSource_To_v1_ConfigMapBuildSource(in *build.ConfigMapBuildSource, out *v1.ConfigMapBuildSource, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_build_ConfigMapBuildSource_To_v1_ConfigMapBuildSource(in, out, s)
}
func autoConvert_v1_CustomBuildStrategy_To_build_CustomBuildStrategy(in *v1.CustomBuildStrategy, out *build.CustomBuildStrategy, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := corev1.Convert_v1_ObjectReference_To_core_ObjectReference(&in.From, &out.From, s); err != nil {
		return err
	}
	if in.PullSecret != nil {
		in, out := &in.PullSecret, &out.PullSecret
		*out = new(core.LocalObjectReference)
		if err := corev1.Convert_v1_LocalObjectReference_To_core_LocalObjectReference(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.PullSecret = nil
	}
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]core.EnvVar, len(*in))
		for i := range *in {
			if err := corev1.Convert_v1_EnvVar_To_core_EnvVar(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Env = nil
	}
	out.ExposeDockerSocket = in.ExposeDockerSocket
	out.ForcePull = in.ForcePull
	if in.Secrets != nil {
		in, out := &in.Secrets, &out.Secrets
		*out = make([]build.SecretSpec, len(*in))
		for i := range *in {
			if err := Convert_v1_SecretSpec_To_build_SecretSpec(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Secrets = nil
	}
	out.BuildAPIVersion = in.BuildAPIVersion
	return nil
}
func autoConvert_build_CustomBuildStrategy_To_v1_CustomBuildStrategy(in *build.CustomBuildStrategy, out *v1.CustomBuildStrategy, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := corev1.Convert_core_ObjectReference_To_v1_ObjectReference(&in.From, &out.From, s); err != nil {
		return err
	}
	if in.PullSecret != nil {
		in, out := &in.PullSecret, &out.PullSecret
		*out = new(apicorev1.LocalObjectReference)
		if err := corev1.Convert_core_LocalObjectReference_To_v1_LocalObjectReference(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.PullSecret = nil
	}
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]apicorev1.EnvVar, len(*in))
		for i := range *in {
			if err := corev1.Convert_core_EnvVar_To_v1_EnvVar(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Env = nil
	}
	out.ExposeDockerSocket = in.ExposeDockerSocket
	out.ForcePull = in.ForcePull
	if in.Secrets != nil {
		in, out := &in.Secrets, &out.Secrets
		*out = make([]v1.SecretSpec, len(*in))
		for i := range *in {
			if err := Convert_build_SecretSpec_To_v1_SecretSpec(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Secrets = nil
	}
	out.BuildAPIVersion = in.BuildAPIVersion
	return nil
}
func Convert_build_CustomBuildStrategy_To_v1_CustomBuildStrategy(in *build.CustomBuildStrategy, out *v1.CustomBuildStrategy, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_build_CustomBuildStrategy_To_v1_CustomBuildStrategy(in, out, s)
}
func autoConvert_v1_DockerBuildStrategy_To_build_DockerBuildStrategy(in *v1.DockerBuildStrategy, out *build.DockerBuildStrategy, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in.From != nil {
		in, out := &in.From, &out.From
		*out = new(core.ObjectReference)
		if err := corev1.Convert_v1_ObjectReference_To_core_ObjectReference(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.From = nil
	}
	if in.PullSecret != nil {
		in, out := &in.PullSecret, &out.PullSecret
		*out = new(core.LocalObjectReference)
		if err := corev1.Convert_v1_LocalObjectReference_To_core_LocalObjectReference(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.PullSecret = nil
	}
	out.NoCache = in.NoCache
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]core.EnvVar, len(*in))
		for i := range *in {
			if err := corev1.Convert_v1_EnvVar_To_core_EnvVar(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Env = nil
	}
	out.ForcePull = in.ForcePull
	out.DockerfilePath = in.DockerfilePath
	if in.BuildArgs != nil {
		in, out := &in.BuildArgs, &out.BuildArgs
		*out = make([]core.EnvVar, len(*in))
		for i := range *in {
			if err := corev1.Convert_v1_EnvVar_To_core_EnvVar(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.BuildArgs = nil
	}
	out.ImageOptimizationPolicy = (*build.ImageOptimizationPolicy)(unsafe.Pointer(in.ImageOptimizationPolicy))
	return nil
}
func autoConvert_build_DockerBuildStrategy_To_v1_DockerBuildStrategy(in *build.DockerBuildStrategy, out *v1.DockerBuildStrategy, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in.From != nil {
		in, out := &in.From, &out.From
		*out = new(apicorev1.ObjectReference)
		if err := corev1.Convert_core_ObjectReference_To_v1_ObjectReference(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.From = nil
	}
	if in.PullSecret != nil {
		in, out := &in.PullSecret, &out.PullSecret
		*out = new(apicorev1.LocalObjectReference)
		if err := corev1.Convert_core_LocalObjectReference_To_v1_LocalObjectReference(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.PullSecret = nil
	}
	out.NoCache = in.NoCache
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]apicorev1.EnvVar, len(*in))
		for i := range *in {
			if err := corev1.Convert_core_EnvVar_To_v1_EnvVar(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Env = nil
	}
	if in.BuildArgs != nil {
		in, out := &in.BuildArgs, &out.BuildArgs
		*out = make([]apicorev1.EnvVar, len(*in))
		for i := range *in {
			if err := corev1.Convert_core_EnvVar_To_v1_EnvVar(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.BuildArgs = nil
	}
	out.ForcePull = in.ForcePull
	out.DockerfilePath = in.DockerfilePath
	out.ImageOptimizationPolicy = (*v1.ImageOptimizationPolicy)(unsafe.Pointer(in.ImageOptimizationPolicy))
	return nil
}
func Convert_build_DockerBuildStrategy_To_v1_DockerBuildStrategy(in *build.DockerBuildStrategy, out *v1.DockerBuildStrategy, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_build_DockerBuildStrategy_To_v1_DockerBuildStrategy(in, out, s)
}
func autoConvert_v1_DockerStrategyOptions_To_build_DockerStrategyOptions(in *v1.DockerStrategyOptions, out *build.DockerStrategyOptions, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in.BuildArgs != nil {
		in, out := &in.BuildArgs, &out.BuildArgs
		*out = make([]core.EnvVar, len(*in))
		for i := range *in {
			if err := corev1.Convert_v1_EnvVar_To_core_EnvVar(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.BuildArgs = nil
	}
	out.NoCache = (*bool)(unsafe.Pointer(in.NoCache))
	return nil
}
func Convert_v1_DockerStrategyOptions_To_build_DockerStrategyOptions(in *v1.DockerStrategyOptions, out *build.DockerStrategyOptions, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_DockerStrategyOptions_To_build_DockerStrategyOptions(in, out, s)
}
func autoConvert_build_DockerStrategyOptions_To_v1_DockerStrategyOptions(in *build.DockerStrategyOptions, out *v1.DockerStrategyOptions, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in.BuildArgs != nil {
		in, out := &in.BuildArgs, &out.BuildArgs
		*out = make([]apicorev1.EnvVar, len(*in))
		for i := range *in {
			if err := corev1.Convert_core_EnvVar_To_v1_EnvVar(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.BuildArgs = nil
	}
	out.NoCache = (*bool)(unsafe.Pointer(in.NoCache))
	return nil
}
func Convert_build_DockerStrategyOptions_To_v1_DockerStrategyOptions(in *build.DockerStrategyOptions, out *v1.DockerStrategyOptions, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_build_DockerStrategyOptions_To_v1_DockerStrategyOptions(in, out, s)
}
func autoConvert_v1_GenericWebHookCause_To_build_GenericWebHookCause(in *v1.GenericWebHookCause, out *build.GenericWebHookCause, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in.Revision != nil {
		in, out := &in.Revision, &out.Revision
		*out = new(build.SourceRevision)
		if err := Convert_v1_SourceRevision_To_build_SourceRevision(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.Revision = nil
	}
	out.Secret = in.Secret
	return nil
}
func Convert_v1_GenericWebHookCause_To_build_GenericWebHookCause(in *v1.GenericWebHookCause, out *build.GenericWebHookCause, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_GenericWebHookCause_To_build_GenericWebHookCause(in, out, s)
}
func autoConvert_build_GenericWebHookCause_To_v1_GenericWebHookCause(in *build.GenericWebHookCause, out *v1.GenericWebHookCause, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in.Revision != nil {
		in, out := &in.Revision, &out.Revision
		*out = new(v1.SourceRevision)
		if err := Convert_build_SourceRevision_To_v1_SourceRevision(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.Revision = nil
	}
	out.Secret = in.Secret
	return nil
}
func Convert_build_GenericWebHookCause_To_v1_GenericWebHookCause(in *build.GenericWebHookCause, out *v1.GenericWebHookCause, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_build_GenericWebHookCause_To_v1_GenericWebHookCause(in, out, s)
}
func autoConvert_v1_GenericWebHookEvent_To_build_GenericWebHookEvent(in *v1.GenericWebHookEvent, out *build.GenericWebHookEvent, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Git = (*build.GitInfo)(unsafe.Pointer(in.Git))
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]core.EnvVar, len(*in))
		for i := range *in {
			if err := corev1.Convert_v1_EnvVar_To_core_EnvVar(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Env = nil
	}
	if in.DockerStrategyOptions != nil {
		in, out := &in.DockerStrategyOptions, &out.DockerStrategyOptions
		*out = new(build.DockerStrategyOptions)
		if err := Convert_v1_DockerStrategyOptions_To_build_DockerStrategyOptions(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.DockerStrategyOptions = nil
	}
	return nil
}
func Convert_v1_GenericWebHookEvent_To_build_GenericWebHookEvent(in *v1.GenericWebHookEvent, out *build.GenericWebHookEvent, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_GenericWebHookEvent_To_build_GenericWebHookEvent(in, out, s)
}
func autoConvert_build_GenericWebHookEvent_To_v1_GenericWebHookEvent(in *build.GenericWebHookEvent, out *v1.GenericWebHookEvent, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Git = (*v1.GitInfo)(unsafe.Pointer(in.Git))
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]apicorev1.EnvVar, len(*in))
		for i := range *in {
			if err := corev1.Convert_core_EnvVar_To_v1_EnvVar(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Env = nil
	}
	if in.DockerStrategyOptions != nil {
		in, out := &in.DockerStrategyOptions, &out.DockerStrategyOptions
		*out = new(v1.DockerStrategyOptions)
		if err := Convert_build_DockerStrategyOptions_To_v1_DockerStrategyOptions(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.DockerStrategyOptions = nil
	}
	return nil
}
func Convert_build_GenericWebHookEvent_To_v1_GenericWebHookEvent(in *build.GenericWebHookEvent, out *v1.GenericWebHookEvent, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_build_GenericWebHookEvent_To_v1_GenericWebHookEvent(in, out, s)
}
func autoConvert_v1_GitBuildSource_To_build_GitBuildSource(in *v1.GitBuildSource, out *build.GitBuildSource, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.URI = in.URI
	out.Ref = in.Ref
	if err := Convert_v1_ProxyConfig_To_build_ProxyConfig(&in.ProxyConfig, &out.ProxyConfig, s); err != nil {
		return err
	}
	return nil
}
func Convert_v1_GitBuildSource_To_build_GitBuildSource(in *v1.GitBuildSource, out *build.GitBuildSource, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_GitBuildSource_To_build_GitBuildSource(in, out, s)
}
func autoConvert_build_GitBuildSource_To_v1_GitBuildSource(in *build.GitBuildSource, out *v1.GitBuildSource, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.URI = in.URI
	out.Ref = in.Ref
	if err := Convert_build_ProxyConfig_To_v1_ProxyConfig(&in.ProxyConfig, &out.ProxyConfig, s); err != nil {
		return err
	}
	return nil
}
func Convert_build_GitBuildSource_To_v1_GitBuildSource(in *build.GitBuildSource, out *v1.GitBuildSource, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_build_GitBuildSource_To_v1_GitBuildSource(in, out, s)
}
func autoConvert_v1_GitHubWebHookCause_To_build_GitHubWebHookCause(in *v1.GitHubWebHookCause, out *build.GitHubWebHookCause, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in.Revision != nil {
		in, out := &in.Revision, &out.Revision
		*out = new(build.SourceRevision)
		if err := Convert_v1_SourceRevision_To_build_SourceRevision(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.Revision = nil
	}
	out.Secret = in.Secret
	return nil
}
func Convert_v1_GitHubWebHookCause_To_build_GitHubWebHookCause(in *v1.GitHubWebHookCause, out *build.GitHubWebHookCause, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_GitHubWebHookCause_To_build_GitHubWebHookCause(in, out, s)
}
func autoConvert_build_GitHubWebHookCause_To_v1_GitHubWebHookCause(in *build.GitHubWebHookCause, out *v1.GitHubWebHookCause, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in.Revision != nil {
		in, out := &in.Revision, &out.Revision
		*out = new(v1.SourceRevision)
		if err := Convert_build_SourceRevision_To_v1_SourceRevision(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.Revision = nil
	}
	out.Secret = in.Secret
	return nil
}
func Convert_build_GitHubWebHookCause_To_v1_GitHubWebHookCause(in *build.GitHubWebHookCause, out *v1.GitHubWebHookCause, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_build_GitHubWebHookCause_To_v1_GitHubWebHookCause(in, out, s)
}
func autoConvert_v1_GitInfo_To_build_GitInfo(in *v1.GitInfo, out *build.GitInfo, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := Convert_v1_GitBuildSource_To_build_GitBuildSource(&in.GitBuildSource, &out.GitBuildSource, s); err != nil {
		return err
	}
	if err := Convert_v1_GitSourceRevision_To_build_GitSourceRevision(&in.GitSourceRevision, &out.GitSourceRevision, s); err != nil {
		return err
	}
	out.Refs = *(*[]build.GitRefInfo)(unsafe.Pointer(&in.Refs))
	return nil
}
func Convert_v1_GitInfo_To_build_GitInfo(in *v1.GitInfo, out *build.GitInfo, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_GitInfo_To_build_GitInfo(in, out, s)
}
func autoConvert_build_GitInfo_To_v1_GitInfo(in *build.GitInfo, out *v1.GitInfo, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := Convert_build_GitBuildSource_To_v1_GitBuildSource(&in.GitBuildSource, &out.GitBuildSource, s); err != nil {
		return err
	}
	if err := Convert_build_GitSourceRevision_To_v1_GitSourceRevision(&in.GitSourceRevision, &out.GitSourceRevision, s); err != nil {
		return err
	}
	out.Refs = *(*[]v1.GitRefInfo)(unsafe.Pointer(&in.Refs))
	return nil
}
func Convert_build_GitInfo_To_v1_GitInfo(in *build.GitInfo, out *v1.GitInfo, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_build_GitInfo_To_v1_GitInfo(in, out, s)
}
func autoConvert_v1_GitLabWebHookCause_To_build_GitLabWebHookCause(in *v1.GitLabWebHookCause, out *build.GitLabWebHookCause, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := Convert_v1_CommonWebHookCause_To_build_CommonWebHookCause(&in.CommonWebHookCause, &out.CommonWebHookCause, s); err != nil {
		return err
	}
	return nil
}
func Convert_v1_GitLabWebHookCause_To_build_GitLabWebHookCause(in *v1.GitLabWebHookCause, out *build.GitLabWebHookCause, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_GitLabWebHookCause_To_build_GitLabWebHookCause(in, out, s)
}
func autoConvert_build_GitLabWebHookCause_To_v1_GitLabWebHookCause(in *build.GitLabWebHookCause, out *v1.GitLabWebHookCause, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := Convert_build_CommonWebHookCause_To_v1_CommonWebHookCause(&in.CommonWebHookCause, &out.CommonWebHookCause, s); err != nil {
		return err
	}
	return nil
}
func Convert_build_GitLabWebHookCause_To_v1_GitLabWebHookCause(in *build.GitLabWebHookCause, out *v1.GitLabWebHookCause, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_build_GitLabWebHookCause_To_v1_GitLabWebHookCause(in, out, s)
}
func autoConvert_v1_GitRefInfo_To_build_GitRefInfo(in *v1.GitRefInfo, out *build.GitRefInfo, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := Convert_v1_GitBuildSource_To_build_GitBuildSource(&in.GitBuildSource, &out.GitBuildSource, s); err != nil {
		return err
	}
	if err := Convert_v1_GitSourceRevision_To_build_GitSourceRevision(&in.GitSourceRevision, &out.GitSourceRevision, s); err != nil {
		return err
	}
	return nil
}
func Convert_v1_GitRefInfo_To_build_GitRefInfo(in *v1.GitRefInfo, out *build.GitRefInfo, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_GitRefInfo_To_build_GitRefInfo(in, out, s)
}
func autoConvert_build_GitRefInfo_To_v1_GitRefInfo(in *build.GitRefInfo, out *v1.GitRefInfo, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := Convert_build_GitBuildSource_To_v1_GitBuildSource(&in.GitBuildSource, &out.GitBuildSource, s); err != nil {
		return err
	}
	if err := Convert_build_GitSourceRevision_To_v1_GitSourceRevision(&in.GitSourceRevision, &out.GitSourceRevision, s); err != nil {
		return err
	}
	return nil
}
func Convert_build_GitRefInfo_To_v1_GitRefInfo(in *build.GitRefInfo, out *v1.GitRefInfo, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_build_GitRefInfo_To_v1_GitRefInfo(in, out, s)
}
func autoConvert_v1_GitSourceRevision_To_build_GitSourceRevision(in *v1.GitSourceRevision, out *build.GitSourceRevision, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Commit = in.Commit
	if err := Convert_v1_SourceControlUser_To_build_SourceControlUser(&in.Author, &out.Author, s); err != nil {
		return err
	}
	if err := Convert_v1_SourceControlUser_To_build_SourceControlUser(&in.Committer, &out.Committer, s); err != nil {
		return err
	}
	out.Message = in.Message
	return nil
}
func Convert_v1_GitSourceRevision_To_build_GitSourceRevision(in *v1.GitSourceRevision, out *build.GitSourceRevision, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_GitSourceRevision_To_build_GitSourceRevision(in, out, s)
}
func autoConvert_build_GitSourceRevision_To_v1_GitSourceRevision(in *build.GitSourceRevision, out *v1.GitSourceRevision, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Commit = in.Commit
	if err := Convert_build_SourceControlUser_To_v1_SourceControlUser(&in.Author, &out.Author, s); err != nil {
		return err
	}
	if err := Convert_build_SourceControlUser_To_v1_SourceControlUser(&in.Committer, &out.Committer, s); err != nil {
		return err
	}
	out.Message = in.Message
	return nil
}
func Convert_build_GitSourceRevision_To_v1_GitSourceRevision(in *build.GitSourceRevision, out *v1.GitSourceRevision, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_build_GitSourceRevision_To_v1_GitSourceRevision(in, out, s)
}
func autoConvert_v1_ImageChangeCause_To_build_ImageChangeCause(in *v1.ImageChangeCause, out *build.ImageChangeCause, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ImageID = in.ImageID
	if in.FromRef != nil {
		in, out := &in.FromRef, &out.FromRef
		*out = new(core.ObjectReference)
		if err := corev1.Convert_v1_ObjectReference_To_core_ObjectReference(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.FromRef = nil
	}
	return nil
}
func Convert_v1_ImageChangeCause_To_build_ImageChangeCause(in *v1.ImageChangeCause, out *build.ImageChangeCause, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_ImageChangeCause_To_build_ImageChangeCause(in, out, s)
}
func autoConvert_build_ImageChangeCause_To_v1_ImageChangeCause(in *build.ImageChangeCause, out *v1.ImageChangeCause, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ImageID = in.ImageID
	if in.FromRef != nil {
		in, out := &in.FromRef, &out.FromRef
		*out = new(apicorev1.ObjectReference)
		if err := corev1.Convert_core_ObjectReference_To_v1_ObjectReference(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.FromRef = nil
	}
	return nil
}
func Convert_build_ImageChangeCause_To_v1_ImageChangeCause(in *build.ImageChangeCause, out *v1.ImageChangeCause, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_build_ImageChangeCause_To_v1_ImageChangeCause(in, out, s)
}
func autoConvert_v1_ImageChangeTrigger_To_build_ImageChangeTrigger(in *v1.ImageChangeTrigger, out *build.ImageChangeTrigger, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.LastTriggeredImageID = in.LastTriggeredImageID
	if in.From != nil {
		in, out := &in.From, &out.From
		*out = new(core.ObjectReference)
		if err := corev1.Convert_v1_ObjectReference_To_core_ObjectReference(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.From = nil
	}
	out.Paused = in.Paused
	return nil
}
func Convert_v1_ImageChangeTrigger_To_build_ImageChangeTrigger(in *v1.ImageChangeTrigger, out *build.ImageChangeTrigger, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_ImageChangeTrigger_To_build_ImageChangeTrigger(in, out, s)
}
func autoConvert_build_ImageChangeTrigger_To_v1_ImageChangeTrigger(in *build.ImageChangeTrigger, out *v1.ImageChangeTrigger, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.LastTriggeredImageID = in.LastTriggeredImageID
	if in.From != nil {
		in, out := &in.From, &out.From
		*out = new(apicorev1.ObjectReference)
		if err := corev1.Convert_core_ObjectReference_To_v1_ObjectReference(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.From = nil
	}
	out.Paused = in.Paused
	return nil
}
func Convert_build_ImageChangeTrigger_To_v1_ImageChangeTrigger(in *build.ImageChangeTrigger, out *v1.ImageChangeTrigger, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_build_ImageChangeTrigger_To_v1_ImageChangeTrigger(in, out, s)
}
func autoConvert_v1_ImageLabel_To_build_ImageLabel(in *v1.ImageLabel, out *build.ImageLabel, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Name = in.Name
	out.Value = in.Value
	return nil
}
func Convert_v1_ImageLabel_To_build_ImageLabel(in *v1.ImageLabel, out *build.ImageLabel, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_ImageLabel_To_build_ImageLabel(in, out, s)
}
func autoConvert_build_ImageLabel_To_v1_ImageLabel(in *build.ImageLabel, out *v1.ImageLabel, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Name = in.Name
	out.Value = in.Value
	return nil
}
func Convert_build_ImageLabel_To_v1_ImageLabel(in *build.ImageLabel, out *v1.ImageLabel, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_build_ImageLabel_To_v1_ImageLabel(in, out, s)
}
func autoConvert_v1_ImageSource_To_build_ImageSource(in *v1.ImageSource, out *build.ImageSource, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := corev1.Convert_v1_ObjectReference_To_core_ObjectReference(&in.From, &out.From, s); err != nil {
		return err
	}
	out.As = *(*[]string)(unsafe.Pointer(&in.As))
	out.Paths = *(*[]build.ImageSourcePath)(unsafe.Pointer(&in.Paths))
	if in.PullSecret != nil {
		in, out := &in.PullSecret, &out.PullSecret
		*out = new(core.LocalObjectReference)
		if err := corev1.Convert_v1_LocalObjectReference_To_core_LocalObjectReference(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.PullSecret = nil
	}
	return nil
}
func Convert_v1_ImageSource_To_build_ImageSource(in *v1.ImageSource, out *build.ImageSource, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_ImageSource_To_build_ImageSource(in, out, s)
}
func autoConvert_build_ImageSource_To_v1_ImageSource(in *build.ImageSource, out *v1.ImageSource, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := corev1.Convert_core_ObjectReference_To_v1_ObjectReference(&in.From, &out.From, s); err != nil {
		return err
	}
	out.As = *(*[]string)(unsafe.Pointer(&in.As))
	out.Paths = *(*[]v1.ImageSourcePath)(unsafe.Pointer(&in.Paths))
	if in.PullSecret != nil {
		in, out := &in.PullSecret, &out.PullSecret
		*out = new(apicorev1.LocalObjectReference)
		if err := corev1.Convert_core_LocalObjectReference_To_v1_LocalObjectReference(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.PullSecret = nil
	}
	return nil
}
func Convert_build_ImageSource_To_v1_ImageSource(in *build.ImageSource, out *v1.ImageSource, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_build_ImageSource_To_v1_ImageSource(in, out, s)
}
func autoConvert_v1_ImageSourcePath_To_build_ImageSourcePath(in *v1.ImageSourcePath, out *build.ImageSourcePath, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.SourcePath = in.SourcePath
	out.DestinationDir = in.DestinationDir
	return nil
}
func Convert_v1_ImageSourcePath_To_build_ImageSourcePath(in *v1.ImageSourcePath, out *build.ImageSourcePath, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_ImageSourcePath_To_build_ImageSourcePath(in, out, s)
}
func autoConvert_build_ImageSourcePath_To_v1_ImageSourcePath(in *build.ImageSourcePath, out *v1.ImageSourcePath, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.SourcePath = in.SourcePath
	out.DestinationDir = in.DestinationDir
	return nil
}
func Convert_build_ImageSourcePath_To_v1_ImageSourcePath(in *build.ImageSourcePath, out *v1.ImageSourcePath, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_build_ImageSourcePath_To_v1_ImageSourcePath(in, out, s)
}
func autoConvert_v1_JenkinsPipelineBuildStrategy_To_build_JenkinsPipelineBuildStrategy(in *v1.JenkinsPipelineBuildStrategy, out *build.JenkinsPipelineBuildStrategy, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.JenkinsfilePath = in.JenkinsfilePath
	out.Jenkinsfile = in.Jenkinsfile
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]core.EnvVar, len(*in))
		for i := range *in {
			if err := corev1.Convert_v1_EnvVar_To_core_EnvVar(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Env = nil
	}
	return nil
}
func Convert_v1_JenkinsPipelineBuildStrategy_To_build_JenkinsPipelineBuildStrategy(in *v1.JenkinsPipelineBuildStrategy, out *build.JenkinsPipelineBuildStrategy, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_JenkinsPipelineBuildStrategy_To_build_JenkinsPipelineBuildStrategy(in, out, s)
}
func autoConvert_build_JenkinsPipelineBuildStrategy_To_v1_JenkinsPipelineBuildStrategy(in *build.JenkinsPipelineBuildStrategy, out *v1.JenkinsPipelineBuildStrategy, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.JenkinsfilePath = in.JenkinsfilePath
	out.Jenkinsfile = in.Jenkinsfile
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]apicorev1.EnvVar, len(*in))
		for i := range *in {
			if err := corev1.Convert_core_EnvVar_To_v1_EnvVar(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Env = nil
	}
	return nil
}
func Convert_build_JenkinsPipelineBuildStrategy_To_v1_JenkinsPipelineBuildStrategy(in *build.JenkinsPipelineBuildStrategy, out *v1.JenkinsPipelineBuildStrategy, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_build_JenkinsPipelineBuildStrategy_To_v1_JenkinsPipelineBuildStrategy(in, out, s)
}
func autoConvert_v1_ProxyConfig_To_build_ProxyConfig(in *v1.ProxyConfig, out *build.ProxyConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.HTTPProxy = (*string)(unsafe.Pointer(in.HTTPProxy))
	out.HTTPSProxy = (*string)(unsafe.Pointer(in.HTTPSProxy))
	out.NoProxy = (*string)(unsafe.Pointer(in.NoProxy))
	return nil
}
func Convert_v1_ProxyConfig_To_build_ProxyConfig(in *v1.ProxyConfig, out *build.ProxyConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_ProxyConfig_To_build_ProxyConfig(in, out, s)
}
func autoConvert_build_ProxyConfig_To_v1_ProxyConfig(in *build.ProxyConfig, out *v1.ProxyConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.HTTPProxy = (*string)(unsafe.Pointer(in.HTTPProxy))
	out.HTTPSProxy = (*string)(unsafe.Pointer(in.HTTPSProxy))
	out.NoProxy = (*string)(unsafe.Pointer(in.NoProxy))
	return nil
}
func Convert_build_ProxyConfig_To_v1_ProxyConfig(in *build.ProxyConfig, out *v1.ProxyConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_build_ProxyConfig_To_v1_ProxyConfig(in, out, s)
}
func autoConvert_v1_SecretBuildSource_To_build_SecretBuildSource(in *v1.SecretBuildSource, out *build.SecretBuildSource, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := corev1.Convert_v1_LocalObjectReference_To_core_LocalObjectReference(&in.Secret, &out.Secret, s); err != nil {
		return err
	}
	out.DestinationDir = in.DestinationDir
	return nil
}
func Convert_v1_SecretBuildSource_To_build_SecretBuildSource(in *v1.SecretBuildSource, out *build.SecretBuildSource, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_SecretBuildSource_To_build_SecretBuildSource(in, out, s)
}
func autoConvert_build_SecretBuildSource_To_v1_SecretBuildSource(in *build.SecretBuildSource, out *v1.SecretBuildSource, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := corev1.Convert_core_LocalObjectReference_To_v1_LocalObjectReference(&in.Secret, &out.Secret, s); err != nil {
		return err
	}
	out.DestinationDir = in.DestinationDir
	return nil
}
func Convert_build_SecretBuildSource_To_v1_SecretBuildSource(in *build.SecretBuildSource, out *v1.SecretBuildSource, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_build_SecretBuildSource_To_v1_SecretBuildSource(in, out, s)
}
func autoConvert_v1_SecretLocalReference_To_build_SecretLocalReference(in *v1.SecretLocalReference, out *build.SecretLocalReference, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Name = in.Name
	return nil
}
func Convert_v1_SecretLocalReference_To_build_SecretLocalReference(in *v1.SecretLocalReference, out *build.SecretLocalReference, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_SecretLocalReference_To_build_SecretLocalReference(in, out, s)
}
func autoConvert_build_SecretLocalReference_To_v1_SecretLocalReference(in *build.SecretLocalReference, out *v1.SecretLocalReference, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Name = in.Name
	return nil
}
func Convert_build_SecretLocalReference_To_v1_SecretLocalReference(in *build.SecretLocalReference, out *v1.SecretLocalReference, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_build_SecretLocalReference_To_v1_SecretLocalReference(in, out, s)
}
func autoConvert_v1_SecretSpec_To_build_SecretSpec(in *v1.SecretSpec, out *build.SecretSpec, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := corev1.Convert_v1_LocalObjectReference_To_core_LocalObjectReference(&in.SecretSource, &out.SecretSource, s); err != nil {
		return err
	}
	out.MountPath = in.MountPath
	return nil
}
func Convert_v1_SecretSpec_To_build_SecretSpec(in *v1.SecretSpec, out *build.SecretSpec, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_SecretSpec_To_build_SecretSpec(in, out, s)
}
func autoConvert_build_SecretSpec_To_v1_SecretSpec(in *build.SecretSpec, out *v1.SecretSpec, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := corev1.Convert_core_LocalObjectReference_To_v1_LocalObjectReference(&in.SecretSource, &out.SecretSource, s); err != nil {
		return err
	}
	out.MountPath = in.MountPath
	return nil
}
func Convert_build_SecretSpec_To_v1_SecretSpec(in *build.SecretSpec, out *v1.SecretSpec, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_build_SecretSpec_To_v1_SecretSpec(in, out, s)
}
func autoConvert_v1_SourceBuildStrategy_To_build_SourceBuildStrategy(in *v1.SourceBuildStrategy, out *build.SourceBuildStrategy, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := corev1.Convert_v1_ObjectReference_To_core_ObjectReference(&in.From, &out.From, s); err != nil {
		return err
	}
	if in.PullSecret != nil {
		in, out := &in.PullSecret, &out.PullSecret
		*out = new(core.LocalObjectReference)
		if err := corev1.Convert_v1_LocalObjectReference_To_core_LocalObjectReference(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.PullSecret = nil
	}
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]core.EnvVar, len(*in))
		for i := range *in {
			if err := corev1.Convert_v1_EnvVar_To_core_EnvVar(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Env = nil
	}
	out.Scripts = in.Scripts
	out.Incremental = (*bool)(unsafe.Pointer(in.Incremental))
	out.ForcePull = in.ForcePull
	return nil
}
func autoConvert_build_SourceBuildStrategy_To_v1_SourceBuildStrategy(in *build.SourceBuildStrategy, out *v1.SourceBuildStrategy, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := corev1.Convert_core_ObjectReference_To_v1_ObjectReference(&in.From, &out.From, s); err != nil {
		return err
	}
	if in.PullSecret != nil {
		in, out := &in.PullSecret, &out.PullSecret
		*out = new(apicorev1.LocalObjectReference)
		if err := corev1.Convert_core_LocalObjectReference_To_v1_LocalObjectReference(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.PullSecret = nil
	}
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]apicorev1.EnvVar, len(*in))
		for i := range *in {
			if err := corev1.Convert_core_EnvVar_To_v1_EnvVar(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Env = nil
	}
	out.Scripts = in.Scripts
	out.Incremental = (*bool)(unsafe.Pointer(in.Incremental))
	out.ForcePull = in.ForcePull
	return nil
}
func Convert_build_SourceBuildStrategy_To_v1_SourceBuildStrategy(in *build.SourceBuildStrategy, out *v1.SourceBuildStrategy, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_build_SourceBuildStrategy_To_v1_SourceBuildStrategy(in, out, s)
}
func autoConvert_v1_SourceControlUser_To_build_SourceControlUser(in *v1.SourceControlUser, out *build.SourceControlUser, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Name = in.Name
	out.Email = in.Email
	return nil
}
func Convert_v1_SourceControlUser_To_build_SourceControlUser(in *v1.SourceControlUser, out *build.SourceControlUser, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_SourceControlUser_To_build_SourceControlUser(in, out, s)
}
func autoConvert_build_SourceControlUser_To_v1_SourceControlUser(in *build.SourceControlUser, out *v1.SourceControlUser, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Name = in.Name
	out.Email = in.Email
	return nil
}
func Convert_build_SourceControlUser_To_v1_SourceControlUser(in *build.SourceControlUser, out *v1.SourceControlUser, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_build_SourceControlUser_To_v1_SourceControlUser(in, out, s)
}
func autoConvert_v1_SourceRevision_To_build_SourceRevision(in *v1.SourceRevision, out *build.SourceRevision, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Git = (*build.GitSourceRevision)(unsafe.Pointer(in.Git))
	return nil
}
func Convert_v1_SourceRevision_To_build_SourceRevision(in *v1.SourceRevision, out *build.SourceRevision, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_SourceRevision_To_build_SourceRevision(in, out, s)
}
func autoConvert_build_SourceRevision_To_v1_SourceRevision(in *build.SourceRevision, out *v1.SourceRevision, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Git = (*v1.GitSourceRevision)(unsafe.Pointer(in.Git))
	return nil
}
func autoConvert_v1_SourceStrategyOptions_To_build_SourceStrategyOptions(in *v1.SourceStrategyOptions, out *build.SourceStrategyOptions, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Incremental = (*bool)(unsafe.Pointer(in.Incremental))
	return nil
}
func Convert_v1_SourceStrategyOptions_To_build_SourceStrategyOptions(in *v1.SourceStrategyOptions, out *build.SourceStrategyOptions, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_SourceStrategyOptions_To_build_SourceStrategyOptions(in, out, s)
}
func autoConvert_build_SourceStrategyOptions_To_v1_SourceStrategyOptions(in *build.SourceStrategyOptions, out *v1.SourceStrategyOptions, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Incremental = (*bool)(unsafe.Pointer(in.Incremental))
	return nil
}
func Convert_build_SourceStrategyOptions_To_v1_SourceStrategyOptions(in *build.SourceStrategyOptions, out *v1.SourceStrategyOptions, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_build_SourceStrategyOptions_To_v1_SourceStrategyOptions(in, out, s)
}
func autoConvert_v1_StageInfo_To_build_StageInfo(in *v1.StageInfo, out *build.StageInfo, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Name = build.StageName(in.Name)
	out.StartTime = in.StartTime
	out.DurationMilliseconds = in.DurationMilliseconds
	out.Steps = *(*[]build.StepInfo)(unsafe.Pointer(&in.Steps))
	return nil
}
func Convert_v1_StageInfo_To_build_StageInfo(in *v1.StageInfo, out *build.StageInfo, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_StageInfo_To_build_StageInfo(in, out, s)
}
func autoConvert_build_StageInfo_To_v1_StageInfo(in *build.StageInfo, out *v1.StageInfo, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Name = v1.StageName(in.Name)
	out.StartTime = in.StartTime
	out.DurationMilliseconds = in.DurationMilliseconds
	out.Steps = *(*[]v1.StepInfo)(unsafe.Pointer(&in.Steps))
	return nil
}
func Convert_build_StageInfo_To_v1_StageInfo(in *build.StageInfo, out *v1.StageInfo, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_build_StageInfo_To_v1_StageInfo(in, out, s)
}
func autoConvert_v1_StepInfo_To_build_StepInfo(in *v1.StepInfo, out *build.StepInfo, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Name = build.StepName(in.Name)
	out.StartTime = in.StartTime
	out.DurationMilliseconds = in.DurationMilliseconds
	return nil
}
func Convert_v1_StepInfo_To_build_StepInfo(in *v1.StepInfo, out *build.StepInfo, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_StepInfo_To_build_StepInfo(in, out, s)
}
func autoConvert_build_StepInfo_To_v1_StepInfo(in *build.StepInfo, out *v1.StepInfo, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Name = v1.StepName(in.Name)
	out.StartTime = in.StartTime
	out.DurationMilliseconds = in.DurationMilliseconds
	return nil
}
func Convert_build_StepInfo_To_v1_StepInfo(in *build.StepInfo, out *v1.StepInfo, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_build_StepInfo_To_v1_StepInfo(in, out, s)
}
func autoConvert_v1_WebHookTrigger_To_build_WebHookTrigger(in *v1.WebHookTrigger, out *build.WebHookTrigger, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Secret = in.Secret
	out.AllowEnv = in.AllowEnv
	out.SecretReference = (*build.SecretLocalReference)(unsafe.Pointer(in.SecretReference))
	return nil
}
func Convert_v1_WebHookTrigger_To_build_WebHookTrigger(in *v1.WebHookTrigger, out *build.WebHookTrigger, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_WebHookTrigger_To_build_WebHookTrigger(in, out, s)
}
func autoConvert_build_WebHookTrigger_To_v1_WebHookTrigger(in *build.WebHookTrigger, out *v1.WebHookTrigger, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Secret = in.Secret
	out.AllowEnv = in.AllowEnv
	out.SecretReference = (*v1.SecretLocalReference)(unsafe.Pointer(in.SecretReference))
	return nil
}
func Convert_build_WebHookTrigger_To_v1_WebHookTrigger(in *build.WebHookTrigger, out *v1.WebHookTrigger, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_build_WebHookTrigger_To_v1_WebHookTrigger(in, out, s)
}
