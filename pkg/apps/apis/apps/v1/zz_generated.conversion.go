package v1

import (
	unsafe "unsafe"
	v1 "github.com/openshift/api/apps/v1"
	apps "github.com/openshift/origin/pkg/apps/apis/apps"
	apicorev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	core "k8s.io/kubernetes/pkg/apis/core"
	corev1 "k8s.io/kubernetes/pkg/apis/core/v1"
)

func init() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	localSchemeBuilder.Register(RegisterConversions)
}
func RegisterConversions(s *runtime.Scheme) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := s.AddGeneratedConversionFunc((*v1.CustomDeploymentStrategyParams)(nil), (*apps.CustomDeploymentStrategyParams)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_CustomDeploymentStrategyParams_To_apps_CustomDeploymentStrategyParams(a.(*v1.CustomDeploymentStrategyParams), b.(*apps.CustomDeploymentStrategyParams), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*apps.CustomDeploymentStrategyParams)(nil), (*v1.CustomDeploymentStrategyParams)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_apps_CustomDeploymentStrategyParams_To_v1_CustomDeploymentStrategyParams(a.(*apps.CustomDeploymentStrategyParams), b.(*v1.CustomDeploymentStrategyParams), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.DeploymentCause)(nil), (*apps.DeploymentCause)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_DeploymentCause_To_apps_DeploymentCause(a.(*v1.DeploymentCause), b.(*apps.DeploymentCause), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*apps.DeploymentCause)(nil), (*v1.DeploymentCause)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_apps_DeploymentCause_To_v1_DeploymentCause(a.(*apps.DeploymentCause), b.(*v1.DeploymentCause), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.DeploymentCauseImageTrigger)(nil), (*apps.DeploymentCauseImageTrigger)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_DeploymentCauseImageTrigger_To_apps_DeploymentCauseImageTrigger(a.(*v1.DeploymentCauseImageTrigger), b.(*apps.DeploymentCauseImageTrigger), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*apps.DeploymentCauseImageTrigger)(nil), (*v1.DeploymentCauseImageTrigger)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_apps_DeploymentCauseImageTrigger_To_v1_DeploymentCauseImageTrigger(a.(*apps.DeploymentCauseImageTrigger), b.(*v1.DeploymentCauseImageTrigger), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.DeploymentCondition)(nil), (*apps.DeploymentCondition)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_DeploymentCondition_To_apps_DeploymentCondition(a.(*v1.DeploymentCondition), b.(*apps.DeploymentCondition), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*apps.DeploymentCondition)(nil), (*v1.DeploymentCondition)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_apps_DeploymentCondition_To_v1_DeploymentCondition(a.(*apps.DeploymentCondition), b.(*v1.DeploymentCondition), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.DeploymentConfig)(nil), (*apps.DeploymentConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_DeploymentConfig_To_apps_DeploymentConfig(a.(*v1.DeploymentConfig), b.(*apps.DeploymentConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*apps.DeploymentConfig)(nil), (*v1.DeploymentConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_apps_DeploymentConfig_To_v1_DeploymentConfig(a.(*apps.DeploymentConfig), b.(*v1.DeploymentConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.DeploymentConfigList)(nil), (*apps.DeploymentConfigList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_DeploymentConfigList_To_apps_DeploymentConfigList(a.(*v1.DeploymentConfigList), b.(*apps.DeploymentConfigList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*apps.DeploymentConfigList)(nil), (*v1.DeploymentConfigList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_apps_DeploymentConfigList_To_v1_DeploymentConfigList(a.(*apps.DeploymentConfigList), b.(*v1.DeploymentConfigList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.DeploymentConfigRollback)(nil), (*apps.DeploymentConfigRollback)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_DeploymentConfigRollback_To_apps_DeploymentConfigRollback(a.(*v1.DeploymentConfigRollback), b.(*apps.DeploymentConfigRollback), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*apps.DeploymentConfigRollback)(nil), (*v1.DeploymentConfigRollback)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_apps_DeploymentConfigRollback_To_v1_DeploymentConfigRollback(a.(*apps.DeploymentConfigRollback), b.(*v1.DeploymentConfigRollback), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.DeploymentConfigRollbackSpec)(nil), (*apps.DeploymentConfigRollbackSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_DeploymentConfigRollbackSpec_To_apps_DeploymentConfigRollbackSpec(a.(*v1.DeploymentConfigRollbackSpec), b.(*apps.DeploymentConfigRollbackSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*apps.DeploymentConfigRollbackSpec)(nil), (*v1.DeploymentConfigRollbackSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_apps_DeploymentConfigRollbackSpec_To_v1_DeploymentConfigRollbackSpec(a.(*apps.DeploymentConfigRollbackSpec), b.(*v1.DeploymentConfigRollbackSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.DeploymentConfigSpec)(nil), (*apps.DeploymentConfigSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_DeploymentConfigSpec_To_apps_DeploymentConfigSpec(a.(*v1.DeploymentConfigSpec), b.(*apps.DeploymentConfigSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*apps.DeploymentConfigSpec)(nil), (*v1.DeploymentConfigSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_apps_DeploymentConfigSpec_To_v1_DeploymentConfigSpec(a.(*apps.DeploymentConfigSpec), b.(*v1.DeploymentConfigSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.DeploymentConfigStatus)(nil), (*apps.DeploymentConfigStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_DeploymentConfigStatus_To_apps_DeploymentConfigStatus(a.(*v1.DeploymentConfigStatus), b.(*apps.DeploymentConfigStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*apps.DeploymentConfigStatus)(nil), (*v1.DeploymentConfigStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_apps_DeploymentConfigStatus_To_v1_DeploymentConfigStatus(a.(*apps.DeploymentConfigStatus), b.(*v1.DeploymentConfigStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.DeploymentDetails)(nil), (*apps.DeploymentDetails)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_DeploymentDetails_To_apps_DeploymentDetails(a.(*v1.DeploymentDetails), b.(*apps.DeploymentDetails), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*apps.DeploymentDetails)(nil), (*v1.DeploymentDetails)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_apps_DeploymentDetails_To_v1_DeploymentDetails(a.(*apps.DeploymentDetails), b.(*v1.DeploymentDetails), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.DeploymentLog)(nil), (*apps.DeploymentLog)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_DeploymentLog_To_apps_DeploymentLog(a.(*v1.DeploymentLog), b.(*apps.DeploymentLog), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*apps.DeploymentLog)(nil), (*v1.DeploymentLog)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_apps_DeploymentLog_To_v1_DeploymentLog(a.(*apps.DeploymentLog), b.(*v1.DeploymentLog), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.DeploymentLogOptions)(nil), (*apps.DeploymentLogOptions)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_DeploymentLogOptions_To_apps_DeploymentLogOptions(a.(*v1.DeploymentLogOptions), b.(*apps.DeploymentLogOptions), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*apps.DeploymentLogOptions)(nil), (*v1.DeploymentLogOptions)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_apps_DeploymentLogOptions_To_v1_DeploymentLogOptions(a.(*apps.DeploymentLogOptions), b.(*v1.DeploymentLogOptions), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.DeploymentRequest)(nil), (*apps.DeploymentRequest)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_DeploymentRequest_To_apps_DeploymentRequest(a.(*v1.DeploymentRequest), b.(*apps.DeploymentRequest), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*apps.DeploymentRequest)(nil), (*v1.DeploymentRequest)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_apps_DeploymentRequest_To_v1_DeploymentRequest(a.(*apps.DeploymentRequest), b.(*v1.DeploymentRequest), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.DeploymentStrategy)(nil), (*apps.DeploymentStrategy)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_DeploymentStrategy_To_apps_DeploymentStrategy(a.(*v1.DeploymentStrategy), b.(*apps.DeploymentStrategy), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*apps.DeploymentStrategy)(nil), (*v1.DeploymentStrategy)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_apps_DeploymentStrategy_To_v1_DeploymentStrategy(a.(*apps.DeploymentStrategy), b.(*v1.DeploymentStrategy), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.DeploymentTriggerImageChangeParams)(nil), (*apps.DeploymentTriggerImageChangeParams)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_DeploymentTriggerImageChangeParams_To_apps_DeploymentTriggerImageChangeParams(a.(*v1.DeploymentTriggerImageChangeParams), b.(*apps.DeploymentTriggerImageChangeParams), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*apps.DeploymentTriggerImageChangeParams)(nil), (*v1.DeploymentTriggerImageChangeParams)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_apps_DeploymentTriggerImageChangeParams_To_v1_DeploymentTriggerImageChangeParams(a.(*apps.DeploymentTriggerImageChangeParams), b.(*v1.DeploymentTriggerImageChangeParams), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.DeploymentTriggerPolicy)(nil), (*apps.DeploymentTriggerPolicy)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_DeploymentTriggerPolicy_To_apps_DeploymentTriggerPolicy(a.(*v1.DeploymentTriggerPolicy), b.(*apps.DeploymentTriggerPolicy), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*apps.DeploymentTriggerPolicy)(nil), (*v1.DeploymentTriggerPolicy)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_apps_DeploymentTriggerPolicy_To_v1_DeploymentTriggerPolicy(a.(*apps.DeploymentTriggerPolicy), b.(*v1.DeploymentTriggerPolicy), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.ExecNewPodHook)(nil), (*apps.ExecNewPodHook)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_ExecNewPodHook_To_apps_ExecNewPodHook(a.(*v1.ExecNewPodHook), b.(*apps.ExecNewPodHook), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*apps.ExecNewPodHook)(nil), (*v1.ExecNewPodHook)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_apps_ExecNewPodHook_To_v1_ExecNewPodHook(a.(*apps.ExecNewPodHook), b.(*v1.ExecNewPodHook), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.LifecycleHook)(nil), (*apps.LifecycleHook)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_LifecycleHook_To_apps_LifecycleHook(a.(*v1.LifecycleHook), b.(*apps.LifecycleHook), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*apps.LifecycleHook)(nil), (*v1.LifecycleHook)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_apps_LifecycleHook_To_v1_LifecycleHook(a.(*apps.LifecycleHook), b.(*v1.LifecycleHook), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.RecreateDeploymentStrategyParams)(nil), (*apps.RecreateDeploymentStrategyParams)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_RecreateDeploymentStrategyParams_To_apps_RecreateDeploymentStrategyParams(a.(*v1.RecreateDeploymentStrategyParams), b.(*apps.RecreateDeploymentStrategyParams), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*apps.RecreateDeploymentStrategyParams)(nil), (*v1.RecreateDeploymentStrategyParams)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_apps_RecreateDeploymentStrategyParams_To_v1_RecreateDeploymentStrategyParams(a.(*apps.RecreateDeploymentStrategyParams), b.(*v1.RecreateDeploymentStrategyParams), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.RollingDeploymentStrategyParams)(nil), (*apps.RollingDeploymentStrategyParams)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_RollingDeploymentStrategyParams_To_apps_RollingDeploymentStrategyParams(a.(*v1.RollingDeploymentStrategyParams), b.(*apps.RollingDeploymentStrategyParams), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*apps.RollingDeploymentStrategyParams)(nil), (*v1.RollingDeploymentStrategyParams)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_apps_RollingDeploymentStrategyParams_To_v1_RollingDeploymentStrategyParams(a.(*apps.RollingDeploymentStrategyParams), b.(*v1.RollingDeploymentStrategyParams), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.TagImageHook)(nil), (*apps.TagImageHook)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_TagImageHook_To_apps_TagImageHook(a.(*v1.TagImageHook), b.(*apps.TagImageHook), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*apps.TagImageHook)(nil), (*v1.TagImageHook)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_apps_TagImageHook_To_v1_TagImageHook(a.(*apps.TagImageHook), b.(*v1.TagImageHook), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*apps.DeploymentTriggerImageChangeParams)(nil), (*v1.DeploymentTriggerImageChangeParams)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_apps_DeploymentTriggerImageChangeParams_To_v1_DeploymentTriggerImageChangeParams(a.(*apps.DeploymentTriggerImageChangeParams), b.(*v1.DeploymentTriggerImageChangeParams), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*apps.RollingDeploymentStrategyParams)(nil), (*v1.RollingDeploymentStrategyParams)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_apps_RollingDeploymentStrategyParams_To_v1_RollingDeploymentStrategyParams(a.(*apps.RollingDeploymentStrategyParams), b.(*v1.RollingDeploymentStrategyParams), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*v1.DeploymentTriggerImageChangeParams)(nil), (*apps.DeploymentTriggerImageChangeParams)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_DeploymentTriggerImageChangeParams_To_apps_DeploymentTriggerImageChangeParams(a.(*v1.DeploymentTriggerImageChangeParams), b.(*apps.DeploymentTriggerImageChangeParams), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*v1.RollingDeploymentStrategyParams)(nil), (*apps.RollingDeploymentStrategyParams)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_RollingDeploymentStrategyParams_To_apps_RollingDeploymentStrategyParams(a.(*v1.RollingDeploymentStrategyParams), b.(*apps.RollingDeploymentStrategyParams), scope)
	}); err != nil {
		return err
	}
	return nil
}
func autoConvert_v1_CustomDeploymentStrategyParams_To_apps_CustomDeploymentStrategyParams(in *v1.CustomDeploymentStrategyParams, out *apps.CustomDeploymentStrategyParams, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Image = in.Image
	if in.Environment != nil {
		in, out := &in.Environment, &out.Environment
		*out = make([]core.EnvVar, len(*in))
		for i := range *in {
			if err := corev1.Convert_v1_EnvVar_To_core_EnvVar(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Environment = nil
	}
	out.Command = *(*[]string)(unsafe.Pointer(&in.Command))
	return nil
}
func Convert_v1_CustomDeploymentStrategyParams_To_apps_CustomDeploymentStrategyParams(in *v1.CustomDeploymentStrategyParams, out *apps.CustomDeploymentStrategyParams, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_CustomDeploymentStrategyParams_To_apps_CustomDeploymentStrategyParams(in, out, s)
}
func autoConvert_apps_CustomDeploymentStrategyParams_To_v1_CustomDeploymentStrategyParams(in *apps.CustomDeploymentStrategyParams, out *v1.CustomDeploymentStrategyParams, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Image = in.Image
	if in.Environment != nil {
		in, out := &in.Environment, &out.Environment
		*out = make([]apicorev1.EnvVar, len(*in))
		for i := range *in {
			if err := corev1.Convert_core_EnvVar_To_v1_EnvVar(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Environment = nil
	}
	out.Command = *(*[]string)(unsafe.Pointer(&in.Command))
	return nil
}
func Convert_apps_CustomDeploymentStrategyParams_To_v1_CustomDeploymentStrategyParams(in *apps.CustomDeploymentStrategyParams, out *v1.CustomDeploymentStrategyParams, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_apps_CustomDeploymentStrategyParams_To_v1_CustomDeploymentStrategyParams(in, out, s)
}
func autoConvert_v1_DeploymentCause_To_apps_DeploymentCause(in *v1.DeploymentCause, out *apps.DeploymentCause, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Type = apps.DeploymentTriggerType(in.Type)
	if in.ImageTrigger != nil {
		in, out := &in.ImageTrigger, &out.ImageTrigger
		*out = new(apps.DeploymentCauseImageTrigger)
		if err := Convert_v1_DeploymentCauseImageTrigger_To_apps_DeploymentCauseImageTrigger(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.ImageTrigger = nil
	}
	return nil
}
func Convert_v1_DeploymentCause_To_apps_DeploymentCause(in *v1.DeploymentCause, out *apps.DeploymentCause, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_DeploymentCause_To_apps_DeploymentCause(in, out, s)
}
func autoConvert_apps_DeploymentCause_To_v1_DeploymentCause(in *apps.DeploymentCause, out *v1.DeploymentCause, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Type = v1.DeploymentTriggerType(in.Type)
	if in.ImageTrigger != nil {
		in, out := &in.ImageTrigger, &out.ImageTrigger
		*out = new(v1.DeploymentCauseImageTrigger)
		if err := Convert_apps_DeploymentCauseImageTrigger_To_v1_DeploymentCauseImageTrigger(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.ImageTrigger = nil
	}
	return nil
}
func Convert_apps_DeploymentCause_To_v1_DeploymentCause(in *apps.DeploymentCause, out *v1.DeploymentCause, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_apps_DeploymentCause_To_v1_DeploymentCause(in, out, s)
}
func autoConvert_v1_DeploymentCauseImageTrigger_To_apps_DeploymentCauseImageTrigger(in *v1.DeploymentCauseImageTrigger, out *apps.DeploymentCauseImageTrigger, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := corev1.Convert_v1_ObjectReference_To_core_ObjectReference(&in.From, &out.From, s); err != nil {
		return err
	}
	return nil
}
func Convert_v1_DeploymentCauseImageTrigger_To_apps_DeploymentCauseImageTrigger(in *v1.DeploymentCauseImageTrigger, out *apps.DeploymentCauseImageTrigger, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_DeploymentCauseImageTrigger_To_apps_DeploymentCauseImageTrigger(in, out, s)
}
func autoConvert_apps_DeploymentCauseImageTrigger_To_v1_DeploymentCauseImageTrigger(in *apps.DeploymentCauseImageTrigger, out *v1.DeploymentCauseImageTrigger, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := corev1.Convert_core_ObjectReference_To_v1_ObjectReference(&in.From, &out.From, s); err != nil {
		return err
	}
	return nil
}
func Convert_apps_DeploymentCauseImageTrigger_To_v1_DeploymentCauseImageTrigger(in *apps.DeploymentCauseImageTrigger, out *v1.DeploymentCauseImageTrigger, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_apps_DeploymentCauseImageTrigger_To_v1_DeploymentCauseImageTrigger(in, out, s)
}
func autoConvert_v1_DeploymentCondition_To_apps_DeploymentCondition(in *v1.DeploymentCondition, out *apps.DeploymentCondition, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Type = apps.DeploymentConditionType(in.Type)
	out.Status = core.ConditionStatus(in.Status)
	out.LastUpdateTime = in.LastUpdateTime
	out.LastTransitionTime = in.LastTransitionTime
	out.Reason = apps.DeploymentConditionReason(in.Reason)
	out.Message = in.Message
	return nil
}
func Convert_v1_DeploymentCondition_To_apps_DeploymentCondition(in *v1.DeploymentCondition, out *apps.DeploymentCondition, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_DeploymentCondition_To_apps_DeploymentCondition(in, out, s)
}
func autoConvert_apps_DeploymentCondition_To_v1_DeploymentCondition(in *apps.DeploymentCondition, out *v1.DeploymentCondition, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Type = v1.DeploymentConditionType(in.Type)
	out.Status = apicorev1.ConditionStatus(in.Status)
	out.LastUpdateTime = in.LastUpdateTime
	out.LastTransitionTime = in.LastTransitionTime
	out.Reason = string(in.Reason)
	out.Message = in.Message
	return nil
}
func Convert_apps_DeploymentCondition_To_v1_DeploymentCondition(in *apps.DeploymentCondition, out *v1.DeploymentCondition, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_apps_DeploymentCondition_To_v1_DeploymentCondition(in, out, s)
}
func autoConvert_v1_DeploymentConfig_To_apps_DeploymentConfig(in *v1.DeploymentConfig, out *apps.DeploymentConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1_DeploymentConfigSpec_To_apps_DeploymentConfigSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1_DeploymentConfigStatus_To_apps_DeploymentConfigStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}
func Convert_v1_DeploymentConfig_To_apps_DeploymentConfig(in *v1.DeploymentConfig, out *apps.DeploymentConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_DeploymentConfig_To_apps_DeploymentConfig(in, out, s)
}
func autoConvert_apps_DeploymentConfig_To_v1_DeploymentConfig(in *apps.DeploymentConfig, out *v1.DeploymentConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_apps_DeploymentConfigSpec_To_v1_DeploymentConfigSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_apps_DeploymentConfigStatus_To_v1_DeploymentConfigStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}
func Convert_apps_DeploymentConfig_To_v1_DeploymentConfig(in *apps.DeploymentConfig, out *v1.DeploymentConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_apps_DeploymentConfig_To_v1_DeploymentConfig(in, out, s)
}
func autoConvert_v1_DeploymentConfigList_To_apps_DeploymentConfigList(in *v1.DeploymentConfigList, out *apps.DeploymentConfigList, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]apps.DeploymentConfig, len(*in))
		for i := range *in {
			if err := Convert_v1_DeploymentConfig_To_apps_DeploymentConfig(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}
func Convert_v1_DeploymentConfigList_To_apps_DeploymentConfigList(in *v1.DeploymentConfigList, out *apps.DeploymentConfigList, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_DeploymentConfigList_To_apps_DeploymentConfigList(in, out, s)
}
func autoConvert_apps_DeploymentConfigList_To_v1_DeploymentConfigList(in *apps.DeploymentConfigList, out *v1.DeploymentConfigList, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]v1.DeploymentConfig, len(*in))
		for i := range *in {
			if err := Convert_apps_DeploymentConfig_To_v1_DeploymentConfig(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}
func Convert_apps_DeploymentConfigList_To_v1_DeploymentConfigList(in *apps.DeploymentConfigList, out *v1.DeploymentConfigList, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_apps_DeploymentConfigList_To_v1_DeploymentConfigList(in, out, s)
}
func autoConvert_v1_DeploymentConfigRollback_To_apps_DeploymentConfigRollback(in *v1.DeploymentConfigRollback, out *apps.DeploymentConfigRollback, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Name = in.Name
	out.UpdatedAnnotations = *(*map[string]string)(unsafe.Pointer(&in.UpdatedAnnotations))
	if err := Convert_v1_DeploymentConfigRollbackSpec_To_apps_DeploymentConfigRollbackSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	return nil
}
func Convert_v1_DeploymentConfigRollback_To_apps_DeploymentConfigRollback(in *v1.DeploymentConfigRollback, out *apps.DeploymentConfigRollback, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_DeploymentConfigRollback_To_apps_DeploymentConfigRollback(in, out, s)
}
func autoConvert_apps_DeploymentConfigRollback_To_v1_DeploymentConfigRollback(in *apps.DeploymentConfigRollback, out *v1.DeploymentConfigRollback, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Name = in.Name
	out.UpdatedAnnotations = *(*map[string]string)(unsafe.Pointer(&in.UpdatedAnnotations))
	if err := Convert_apps_DeploymentConfigRollbackSpec_To_v1_DeploymentConfigRollbackSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	return nil
}
func Convert_apps_DeploymentConfigRollback_To_v1_DeploymentConfigRollback(in *apps.DeploymentConfigRollback, out *v1.DeploymentConfigRollback, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_apps_DeploymentConfigRollback_To_v1_DeploymentConfigRollback(in, out, s)
}
func autoConvert_v1_DeploymentConfigRollbackSpec_To_apps_DeploymentConfigRollbackSpec(in *v1.DeploymentConfigRollbackSpec, out *apps.DeploymentConfigRollbackSpec, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := corev1.Convert_v1_ObjectReference_To_core_ObjectReference(&in.From, &out.From, s); err != nil {
		return err
	}
	out.Revision = in.Revision
	out.IncludeTriggers = in.IncludeTriggers
	out.IncludeTemplate = in.IncludeTemplate
	out.IncludeReplicationMeta = in.IncludeReplicationMeta
	out.IncludeStrategy = in.IncludeStrategy
	return nil
}
func Convert_v1_DeploymentConfigRollbackSpec_To_apps_DeploymentConfigRollbackSpec(in *v1.DeploymentConfigRollbackSpec, out *apps.DeploymentConfigRollbackSpec, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_DeploymentConfigRollbackSpec_To_apps_DeploymentConfigRollbackSpec(in, out, s)
}
func autoConvert_apps_DeploymentConfigRollbackSpec_To_v1_DeploymentConfigRollbackSpec(in *apps.DeploymentConfigRollbackSpec, out *v1.DeploymentConfigRollbackSpec, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := corev1.Convert_core_ObjectReference_To_v1_ObjectReference(&in.From, &out.From, s); err != nil {
		return err
	}
	out.Revision = in.Revision
	out.IncludeTriggers = in.IncludeTriggers
	out.IncludeTemplate = in.IncludeTemplate
	out.IncludeReplicationMeta = in.IncludeReplicationMeta
	out.IncludeStrategy = in.IncludeStrategy
	return nil
}
func Convert_apps_DeploymentConfigRollbackSpec_To_v1_DeploymentConfigRollbackSpec(in *apps.DeploymentConfigRollbackSpec, out *v1.DeploymentConfigRollbackSpec, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_apps_DeploymentConfigRollbackSpec_To_v1_DeploymentConfigRollbackSpec(in, out, s)
}
func autoConvert_v1_DeploymentConfigSpec_To_apps_DeploymentConfigSpec(in *v1.DeploymentConfigSpec, out *apps.DeploymentConfigSpec, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := Convert_v1_DeploymentStrategy_To_apps_DeploymentStrategy(&in.Strategy, &out.Strategy, s); err != nil {
		return err
	}
	out.MinReadySeconds = in.MinReadySeconds
	if in.Triggers != nil {
		in, out := &in.Triggers, &out.Triggers
		*out = make([]apps.DeploymentTriggerPolicy, len(*in))
		for i := range *in {
			if err := Convert_v1_DeploymentTriggerPolicy_To_apps_DeploymentTriggerPolicy(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Triggers = nil
	}
	out.Replicas = in.Replicas
	out.RevisionHistoryLimit = (*int32)(unsafe.Pointer(in.RevisionHistoryLimit))
	out.Test = in.Test
	out.Paused = in.Paused
	out.Selector = *(*map[string]string)(unsafe.Pointer(&in.Selector))
	if in.Template != nil {
		in, out := &in.Template, &out.Template
		*out = new(core.PodTemplateSpec)
		if err := corev1.Convert_v1_PodTemplateSpec_To_core_PodTemplateSpec(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.Template = nil
	}
	return nil
}
func Convert_v1_DeploymentConfigSpec_To_apps_DeploymentConfigSpec(in *v1.DeploymentConfigSpec, out *apps.DeploymentConfigSpec, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_DeploymentConfigSpec_To_apps_DeploymentConfigSpec(in, out, s)
}
func autoConvert_apps_DeploymentConfigSpec_To_v1_DeploymentConfigSpec(in *apps.DeploymentConfigSpec, out *v1.DeploymentConfigSpec, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := Convert_apps_DeploymentStrategy_To_v1_DeploymentStrategy(&in.Strategy, &out.Strategy, s); err != nil {
		return err
	}
	out.MinReadySeconds = in.MinReadySeconds
	if in.Triggers != nil {
		in, out := &in.Triggers, &out.Triggers
		*out = make(v1.DeploymentTriggerPolicies, len(*in))
		for i := range *in {
			if err := Convert_apps_DeploymentTriggerPolicy_To_v1_DeploymentTriggerPolicy(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Triggers = nil
	}
	out.Replicas = in.Replicas
	out.RevisionHistoryLimit = (*int32)(unsafe.Pointer(in.RevisionHistoryLimit))
	out.Test = in.Test
	out.Paused = in.Paused
	out.Selector = *(*map[string]string)(unsafe.Pointer(&in.Selector))
	if in.Template != nil {
		in, out := &in.Template, &out.Template
		*out = new(apicorev1.PodTemplateSpec)
		if err := corev1.Convert_core_PodTemplateSpec_To_v1_PodTemplateSpec(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.Template = nil
	}
	return nil
}
func Convert_apps_DeploymentConfigSpec_To_v1_DeploymentConfigSpec(in *apps.DeploymentConfigSpec, out *v1.DeploymentConfigSpec, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_apps_DeploymentConfigSpec_To_v1_DeploymentConfigSpec(in, out, s)
}
func autoConvert_v1_DeploymentConfigStatus_To_apps_DeploymentConfigStatus(in *v1.DeploymentConfigStatus, out *apps.DeploymentConfigStatus, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.LatestVersion = in.LatestVersion
	out.ObservedGeneration = in.ObservedGeneration
	out.Replicas = in.Replicas
	out.UpdatedReplicas = in.UpdatedReplicas
	out.AvailableReplicas = in.AvailableReplicas
	out.UnavailableReplicas = in.UnavailableReplicas
	if in.Details != nil {
		in, out := &in.Details, &out.Details
		*out = new(apps.DeploymentDetails)
		if err := Convert_v1_DeploymentDetails_To_apps_DeploymentDetails(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.Details = nil
	}
	out.Conditions = *(*[]apps.DeploymentCondition)(unsafe.Pointer(&in.Conditions))
	out.ReadyReplicas = in.ReadyReplicas
	return nil
}
func Convert_v1_DeploymentConfigStatus_To_apps_DeploymentConfigStatus(in *v1.DeploymentConfigStatus, out *apps.DeploymentConfigStatus, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_DeploymentConfigStatus_To_apps_DeploymentConfigStatus(in, out, s)
}
func autoConvert_apps_DeploymentConfigStatus_To_v1_DeploymentConfigStatus(in *apps.DeploymentConfigStatus, out *v1.DeploymentConfigStatus, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.LatestVersion = in.LatestVersion
	out.ObservedGeneration = in.ObservedGeneration
	out.Replicas = in.Replicas
	out.UpdatedReplicas = in.UpdatedReplicas
	out.AvailableReplicas = in.AvailableReplicas
	out.UnavailableReplicas = in.UnavailableReplicas
	if in.Details != nil {
		in, out := &in.Details, &out.Details
		*out = new(v1.DeploymentDetails)
		if err := Convert_apps_DeploymentDetails_To_v1_DeploymentDetails(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.Details = nil
	}
	out.Conditions = *(*[]v1.DeploymentCondition)(unsafe.Pointer(&in.Conditions))
	out.ReadyReplicas = in.ReadyReplicas
	return nil
}
func Convert_apps_DeploymentConfigStatus_To_v1_DeploymentConfigStatus(in *apps.DeploymentConfigStatus, out *v1.DeploymentConfigStatus, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_apps_DeploymentConfigStatus_To_v1_DeploymentConfigStatus(in, out, s)
}
func autoConvert_v1_DeploymentDetails_To_apps_DeploymentDetails(in *v1.DeploymentDetails, out *apps.DeploymentDetails, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Message = in.Message
	if in.Causes != nil {
		in, out := &in.Causes, &out.Causes
		*out = make([]apps.DeploymentCause, len(*in))
		for i := range *in {
			if err := Convert_v1_DeploymentCause_To_apps_DeploymentCause(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Causes = nil
	}
	return nil
}
func Convert_v1_DeploymentDetails_To_apps_DeploymentDetails(in *v1.DeploymentDetails, out *apps.DeploymentDetails, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_DeploymentDetails_To_apps_DeploymentDetails(in, out, s)
}
func autoConvert_apps_DeploymentDetails_To_v1_DeploymentDetails(in *apps.DeploymentDetails, out *v1.DeploymentDetails, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Message = in.Message
	if in.Causes != nil {
		in, out := &in.Causes, &out.Causes
		*out = make([]v1.DeploymentCause, len(*in))
		for i := range *in {
			if err := Convert_apps_DeploymentCause_To_v1_DeploymentCause(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Causes = nil
	}
	return nil
}
func Convert_apps_DeploymentDetails_To_v1_DeploymentDetails(in *apps.DeploymentDetails, out *v1.DeploymentDetails, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_apps_DeploymentDetails_To_v1_DeploymentDetails(in, out, s)
}
func autoConvert_v1_DeploymentLog_To_apps_DeploymentLog(in *v1.DeploymentLog, out *apps.DeploymentLog, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return nil
}
func Convert_v1_DeploymentLog_To_apps_DeploymentLog(in *v1.DeploymentLog, out *apps.DeploymentLog, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_DeploymentLog_To_apps_DeploymentLog(in, out, s)
}
func autoConvert_apps_DeploymentLog_To_v1_DeploymentLog(in *apps.DeploymentLog, out *v1.DeploymentLog, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return nil
}
func Convert_apps_DeploymentLog_To_v1_DeploymentLog(in *apps.DeploymentLog, out *v1.DeploymentLog, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_apps_DeploymentLog_To_v1_DeploymentLog(in, out, s)
}
func autoConvert_v1_DeploymentLogOptions_To_apps_DeploymentLogOptions(in *v1.DeploymentLogOptions, out *apps.DeploymentLogOptions, s conversion.Scope) error {
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
func Convert_v1_DeploymentLogOptions_To_apps_DeploymentLogOptions(in *v1.DeploymentLogOptions, out *apps.DeploymentLogOptions, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_DeploymentLogOptions_To_apps_DeploymentLogOptions(in, out, s)
}
func autoConvert_apps_DeploymentLogOptions_To_v1_DeploymentLogOptions(in *apps.DeploymentLogOptions, out *v1.DeploymentLogOptions, s conversion.Scope) error {
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
func Convert_apps_DeploymentLogOptions_To_v1_DeploymentLogOptions(in *apps.DeploymentLogOptions, out *v1.DeploymentLogOptions, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_apps_DeploymentLogOptions_To_v1_DeploymentLogOptions(in, out, s)
}
func autoConvert_v1_DeploymentRequest_To_apps_DeploymentRequest(in *v1.DeploymentRequest, out *apps.DeploymentRequest, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Name = in.Name
	out.Latest = in.Latest
	out.Force = in.Force
	out.ExcludeTriggers = *(*[]apps.DeploymentTriggerType)(unsafe.Pointer(&in.ExcludeTriggers))
	return nil
}
func Convert_v1_DeploymentRequest_To_apps_DeploymentRequest(in *v1.DeploymentRequest, out *apps.DeploymentRequest, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_DeploymentRequest_To_apps_DeploymentRequest(in, out, s)
}
func autoConvert_apps_DeploymentRequest_To_v1_DeploymentRequest(in *apps.DeploymentRequest, out *v1.DeploymentRequest, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Name = in.Name
	out.Latest = in.Latest
	out.Force = in.Force
	out.ExcludeTriggers = *(*[]v1.DeploymentTriggerType)(unsafe.Pointer(&in.ExcludeTriggers))
	return nil
}
func Convert_apps_DeploymentRequest_To_v1_DeploymentRequest(in *apps.DeploymentRequest, out *v1.DeploymentRequest, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_apps_DeploymentRequest_To_v1_DeploymentRequest(in, out, s)
}
func autoConvert_v1_DeploymentStrategy_To_apps_DeploymentStrategy(in *v1.DeploymentStrategy, out *apps.DeploymentStrategy, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Type = apps.DeploymentStrategyType(in.Type)
	if in.CustomParams != nil {
		in, out := &in.CustomParams, &out.CustomParams
		*out = new(apps.CustomDeploymentStrategyParams)
		if err := Convert_v1_CustomDeploymentStrategyParams_To_apps_CustomDeploymentStrategyParams(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.CustomParams = nil
	}
	if in.RecreateParams != nil {
		in, out := &in.RecreateParams, &out.RecreateParams
		*out = new(apps.RecreateDeploymentStrategyParams)
		if err := Convert_v1_RecreateDeploymentStrategyParams_To_apps_RecreateDeploymentStrategyParams(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.RecreateParams = nil
	}
	if in.RollingParams != nil {
		in, out := &in.RollingParams, &out.RollingParams
		*out = new(apps.RollingDeploymentStrategyParams)
		if err := Convert_v1_RollingDeploymentStrategyParams_To_apps_RollingDeploymentStrategyParams(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.RollingParams = nil
	}
	if err := corev1.Convert_v1_ResourceRequirements_To_core_ResourceRequirements(&in.Resources, &out.Resources, s); err != nil {
		return err
	}
	out.Labels = *(*map[string]string)(unsafe.Pointer(&in.Labels))
	out.Annotations = *(*map[string]string)(unsafe.Pointer(&in.Annotations))
	out.ActiveDeadlineSeconds = (*int64)(unsafe.Pointer(in.ActiveDeadlineSeconds))
	return nil
}
func Convert_v1_DeploymentStrategy_To_apps_DeploymentStrategy(in *v1.DeploymentStrategy, out *apps.DeploymentStrategy, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_DeploymentStrategy_To_apps_DeploymentStrategy(in, out, s)
}
func autoConvert_apps_DeploymentStrategy_To_v1_DeploymentStrategy(in *apps.DeploymentStrategy, out *v1.DeploymentStrategy, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Type = v1.DeploymentStrategyType(in.Type)
	if in.CustomParams != nil {
		in, out := &in.CustomParams, &out.CustomParams
		*out = new(v1.CustomDeploymentStrategyParams)
		if err := Convert_apps_CustomDeploymentStrategyParams_To_v1_CustomDeploymentStrategyParams(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.CustomParams = nil
	}
	if in.RecreateParams != nil {
		in, out := &in.RecreateParams, &out.RecreateParams
		*out = new(v1.RecreateDeploymentStrategyParams)
		if err := Convert_apps_RecreateDeploymentStrategyParams_To_v1_RecreateDeploymentStrategyParams(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.RecreateParams = nil
	}
	if in.RollingParams != nil {
		in, out := &in.RollingParams, &out.RollingParams
		*out = new(v1.RollingDeploymentStrategyParams)
		if err := Convert_apps_RollingDeploymentStrategyParams_To_v1_RollingDeploymentStrategyParams(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.RollingParams = nil
	}
	if err := corev1.Convert_core_ResourceRequirements_To_v1_ResourceRequirements(&in.Resources, &out.Resources, s); err != nil {
		return err
	}
	out.Labels = *(*map[string]string)(unsafe.Pointer(&in.Labels))
	out.Annotations = *(*map[string]string)(unsafe.Pointer(&in.Annotations))
	out.ActiveDeadlineSeconds = (*int64)(unsafe.Pointer(in.ActiveDeadlineSeconds))
	return nil
}
func Convert_apps_DeploymentStrategy_To_v1_DeploymentStrategy(in *apps.DeploymentStrategy, out *v1.DeploymentStrategy, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_apps_DeploymentStrategy_To_v1_DeploymentStrategy(in, out, s)
}
func autoConvert_v1_DeploymentTriggerImageChangeParams_To_apps_DeploymentTriggerImageChangeParams(in *v1.DeploymentTriggerImageChangeParams, out *apps.DeploymentTriggerImageChangeParams, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Automatic = in.Automatic
	out.ContainerNames = *(*[]string)(unsafe.Pointer(&in.ContainerNames))
	if err := corev1.Convert_v1_ObjectReference_To_core_ObjectReference(&in.From, &out.From, s); err != nil {
		return err
	}
	out.LastTriggeredImage = in.LastTriggeredImage
	return nil
}
func autoConvert_apps_DeploymentTriggerImageChangeParams_To_v1_DeploymentTriggerImageChangeParams(in *apps.DeploymentTriggerImageChangeParams, out *v1.DeploymentTriggerImageChangeParams, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Automatic = in.Automatic
	out.ContainerNames = *(*[]string)(unsafe.Pointer(&in.ContainerNames))
	if err := corev1.Convert_core_ObjectReference_To_v1_ObjectReference(&in.From, &out.From, s); err != nil {
		return err
	}
	out.LastTriggeredImage = in.LastTriggeredImage
	return nil
}
func autoConvert_v1_DeploymentTriggerPolicy_To_apps_DeploymentTriggerPolicy(in *v1.DeploymentTriggerPolicy, out *apps.DeploymentTriggerPolicy, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Type = apps.DeploymentTriggerType(in.Type)
	if in.ImageChangeParams != nil {
		in, out := &in.ImageChangeParams, &out.ImageChangeParams
		*out = new(apps.DeploymentTriggerImageChangeParams)
		if err := Convert_v1_DeploymentTriggerImageChangeParams_To_apps_DeploymentTriggerImageChangeParams(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.ImageChangeParams = nil
	}
	return nil
}
func Convert_v1_DeploymentTriggerPolicy_To_apps_DeploymentTriggerPolicy(in *v1.DeploymentTriggerPolicy, out *apps.DeploymentTriggerPolicy, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_DeploymentTriggerPolicy_To_apps_DeploymentTriggerPolicy(in, out, s)
}
func autoConvert_apps_DeploymentTriggerPolicy_To_v1_DeploymentTriggerPolicy(in *apps.DeploymentTriggerPolicy, out *v1.DeploymentTriggerPolicy, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Type = v1.DeploymentTriggerType(in.Type)
	if in.ImageChangeParams != nil {
		in, out := &in.ImageChangeParams, &out.ImageChangeParams
		*out = new(v1.DeploymentTriggerImageChangeParams)
		if err := Convert_apps_DeploymentTriggerImageChangeParams_To_v1_DeploymentTriggerImageChangeParams(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.ImageChangeParams = nil
	}
	return nil
}
func Convert_apps_DeploymentTriggerPolicy_To_v1_DeploymentTriggerPolicy(in *apps.DeploymentTriggerPolicy, out *v1.DeploymentTriggerPolicy, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_apps_DeploymentTriggerPolicy_To_v1_DeploymentTriggerPolicy(in, out, s)
}
func autoConvert_v1_ExecNewPodHook_To_apps_ExecNewPodHook(in *v1.ExecNewPodHook, out *apps.ExecNewPodHook, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Command = *(*[]string)(unsafe.Pointer(&in.Command))
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
	out.ContainerName = in.ContainerName
	out.Volumes = *(*[]string)(unsafe.Pointer(&in.Volumes))
	return nil
}
func Convert_v1_ExecNewPodHook_To_apps_ExecNewPodHook(in *v1.ExecNewPodHook, out *apps.ExecNewPodHook, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_ExecNewPodHook_To_apps_ExecNewPodHook(in, out, s)
}
func autoConvert_apps_ExecNewPodHook_To_v1_ExecNewPodHook(in *apps.ExecNewPodHook, out *v1.ExecNewPodHook, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Command = *(*[]string)(unsafe.Pointer(&in.Command))
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
	out.ContainerName = in.ContainerName
	out.Volumes = *(*[]string)(unsafe.Pointer(&in.Volumes))
	return nil
}
func Convert_apps_ExecNewPodHook_To_v1_ExecNewPodHook(in *apps.ExecNewPodHook, out *v1.ExecNewPodHook, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_apps_ExecNewPodHook_To_v1_ExecNewPodHook(in, out, s)
}
func autoConvert_v1_LifecycleHook_To_apps_LifecycleHook(in *v1.LifecycleHook, out *apps.LifecycleHook, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.FailurePolicy = apps.LifecycleHookFailurePolicy(in.FailurePolicy)
	if in.ExecNewPod != nil {
		in, out := &in.ExecNewPod, &out.ExecNewPod
		*out = new(apps.ExecNewPodHook)
		if err := Convert_v1_ExecNewPodHook_To_apps_ExecNewPodHook(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.ExecNewPod = nil
	}
	if in.TagImages != nil {
		in, out := &in.TagImages, &out.TagImages
		*out = make([]apps.TagImageHook, len(*in))
		for i := range *in {
			if err := Convert_v1_TagImageHook_To_apps_TagImageHook(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.TagImages = nil
	}
	return nil
}
func Convert_v1_LifecycleHook_To_apps_LifecycleHook(in *v1.LifecycleHook, out *apps.LifecycleHook, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_LifecycleHook_To_apps_LifecycleHook(in, out, s)
}
func autoConvert_apps_LifecycleHook_To_v1_LifecycleHook(in *apps.LifecycleHook, out *v1.LifecycleHook, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.FailurePolicy = v1.LifecycleHookFailurePolicy(in.FailurePolicy)
	if in.ExecNewPod != nil {
		in, out := &in.ExecNewPod, &out.ExecNewPod
		*out = new(v1.ExecNewPodHook)
		if err := Convert_apps_ExecNewPodHook_To_v1_ExecNewPodHook(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.ExecNewPod = nil
	}
	if in.TagImages != nil {
		in, out := &in.TagImages, &out.TagImages
		*out = make([]v1.TagImageHook, len(*in))
		for i := range *in {
			if err := Convert_apps_TagImageHook_To_v1_TagImageHook(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.TagImages = nil
	}
	return nil
}
func Convert_apps_LifecycleHook_To_v1_LifecycleHook(in *apps.LifecycleHook, out *v1.LifecycleHook, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_apps_LifecycleHook_To_v1_LifecycleHook(in, out, s)
}
func autoConvert_v1_RecreateDeploymentStrategyParams_To_apps_RecreateDeploymentStrategyParams(in *v1.RecreateDeploymentStrategyParams, out *apps.RecreateDeploymentStrategyParams, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.TimeoutSeconds = (*int64)(unsafe.Pointer(in.TimeoutSeconds))
	if in.Pre != nil {
		in, out := &in.Pre, &out.Pre
		*out = new(apps.LifecycleHook)
		if err := Convert_v1_LifecycleHook_To_apps_LifecycleHook(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.Pre = nil
	}
	if in.Mid != nil {
		in, out := &in.Mid, &out.Mid
		*out = new(apps.LifecycleHook)
		if err := Convert_v1_LifecycleHook_To_apps_LifecycleHook(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.Mid = nil
	}
	if in.Post != nil {
		in, out := &in.Post, &out.Post
		*out = new(apps.LifecycleHook)
		if err := Convert_v1_LifecycleHook_To_apps_LifecycleHook(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.Post = nil
	}
	return nil
}
func Convert_v1_RecreateDeploymentStrategyParams_To_apps_RecreateDeploymentStrategyParams(in *v1.RecreateDeploymentStrategyParams, out *apps.RecreateDeploymentStrategyParams, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_RecreateDeploymentStrategyParams_To_apps_RecreateDeploymentStrategyParams(in, out, s)
}
func autoConvert_apps_RecreateDeploymentStrategyParams_To_v1_RecreateDeploymentStrategyParams(in *apps.RecreateDeploymentStrategyParams, out *v1.RecreateDeploymentStrategyParams, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.TimeoutSeconds = (*int64)(unsafe.Pointer(in.TimeoutSeconds))
	if in.Pre != nil {
		in, out := &in.Pre, &out.Pre
		*out = new(v1.LifecycleHook)
		if err := Convert_apps_LifecycleHook_To_v1_LifecycleHook(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.Pre = nil
	}
	if in.Mid != nil {
		in, out := &in.Mid, &out.Mid
		*out = new(v1.LifecycleHook)
		if err := Convert_apps_LifecycleHook_To_v1_LifecycleHook(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.Mid = nil
	}
	if in.Post != nil {
		in, out := &in.Post, &out.Post
		*out = new(v1.LifecycleHook)
		if err := Convert_apps_LifecycleHook_To_v1_LifecycleHook(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.Post = nil
	}
	return nil
}
func Convert_apps_RecreateDeploymentStrategyParams_To_v1_RecreateDeploymentStrategyParams(in *apps.RecreateDeploymentStrategyParams, out *v1.RecreateDeploymentStrategyParams, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_apps_RecreateDeploymentStrategyParams_To_v1_RecreateDeploymentStrategyParams(in, out, s)
}
func autoConvert_v1_RollingDeploymentStrategyParams_To_apps_RollingDeploymentStrategyParams(in *v1.RollingDeploymentStrategyParams, out *apps.RollingDeploymentStrategyParams, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.UpdatePeriodSeconds = (*int64)(unsafe.Pointer(in.UpdatePeriodSeconds))
	out.IntervalSeconds = (*int64)(unsafe.Pointer(in.IntervalSeconds))
	out.TimeoutSeconds = (*int64)(unsafe.Pointer(in.TimeoutSeconds))
	if in.Pre != nil {
		in, out := &in.Pre, &out.Pre
		*out = new(apps.LifecycleHook)
		if err := Convert_v1_LifecycleHook_To_apps_LifecycleHook(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.Pre = nil
	}
	if in.Post != nil {
		in, out := &in.Post, &out.Post
		*out = new(apps.LifecycleHook)
		if err := Convert_v1_LifecycleHook_To_apps_LifecycleHook(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.Post = nil
	}
	return nil
}
func autoConvert_apps_RollingDeploymentStrategyParams_To_v1_RollingDeploymentStrategyParams(in *apps.RollingDeploymentStrategyParams, out *v1.RollingDeploymentStrategyParams, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.UpdatePeriodSeconds = (*int64)(unsafe.Pointer(in.UpdatePeriodSeconds))
	out.IntervalSeconds = (*int64)(unsafe.Pointer(in.IntervalSeconds))
	out.TimeoutSeconds = (*int64)(unsafe.Pointer(in.TimeoutSeconds))
	if in.Pre != nil {
		in, out := &in.Pre, &out.Pre
		*out = new(v1.LifecycleHook)
		if err := Convert_apps_LifecycleHook_To_v1_LifecycleHook(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.Pre = nil
	}
	if in.Post != nil {
		in, out := &in.Post, &out.Post
		*out = new(v1.LifecycleHook)
		if err := Convert_apps_LifecycleHook_To_v1_LifecycleHook(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.Post = nil
	}
	return nil
}
func autoConvert_v1_TagImageHook_To_apps_TagImageHook(in *v1.TagImageHook, out *apps.TagImageHook, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ContainerName = in.ContainerName
	if err := corev1.Convert_v1_ObjectReference_To_core_ObjectReference(&in.To, &out.To, s); err != nil {
		return err
	}
	return nil
}
func Convert_v1_TagImageHook_To_apps_TagImageHook(in *v1.TagImageHook, out *apps.TagImageHook, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_TagImageHook_To_apps_TagImageHook(in, out, s)
}
func autoConvert_apps_TagImageHook_To_v1_TagImageHook(in *apps.TagImageHook, out *v1.TagImageHook, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ContainerName = in.ContainerName
	if err := corev1.Convert_core_ObjectReference_To_v1_ObjectReference(&in.To, &out.To, s); err != nil {
		return err
	}
	return nil
}
func Convert_apps_TagImageHook_To_v1_TagImageHook(in *apps.TagImageHook, out *v1.TagImageHook, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_apps_TagImageHook_To_v1_TagImageHook(in, out, s)
}
