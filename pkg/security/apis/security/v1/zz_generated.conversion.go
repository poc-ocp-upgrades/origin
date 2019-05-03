package v1

import (
	v1 "github.com/openshift/api/security/v1"
	security "github.com/openshift/origin/pkg/security/apis/security"
	apicorev1 "k8s.io/api/core/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	core "k8s.io/kubernetes/pkg/apis/core"
	corev1 "k8s.io/kubernetes/pkg/apis/core/v1"
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
	if err := s.AddGeneratedConversionFunc((*v1.AllowedFlexVolume)(nil), (*security.AllowedFlexVolume)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_AllowedFlexVolume_To_security_AllowedFlexVolume(a.(*v1.AllowedFlexVolume), b.(*security.AllowedFlexVolume), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*security.AllowedFlexVolume)(nil), (*v1.AllowedFlexVolume)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_security_AllowedFlexVolume_To_v1_AllowedFlexVolume(a.(*security.AllowedFlexVolume), b.(*v1.AllowedFlexVolume), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.FSGroupStrategyOptions)(nil), (*security.FSGroupStrategyOptions)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_FSGroupStrategyOptions_To_security_FSGroupStrategyOptions(a.(*v1.FSGroupStrategyOptions), b.(*security.FSGroupStrategyOptions), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*security.FSGroupStrategyOptions)(nil), (*v1.FSGroupStrategyOptions)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_security_FSGroupStrategyOptions_To_v1_FSGroupStrategyOptions(a.(*security.FSGroupStrategyOptions), b.(*v1.FSGroupStrategyOptions), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.IDRange)(nil), (*security.IDRange)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_IDRange_To_security_IDRange(a.(*v1.IDRange), b.(*security.IDRange), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*security.IDRange)(nil), (*v1.IDRange)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_security_IDRange_To_v1_IDRange(a.(*security.IDRange), b.(*v1.IDRange), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.PodSecurityPolicyReview)(nil), (*security.PodSecurityPolicyReview)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_PodSecurityPolicyReview_To_security_PodSecurityPolicyReview(a.(*v1.PodSecurityPolicyReview), b.(*security.PodSecurityPolicyReview), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*security.PodSecurityPolicyReview)(nil), (*v1.PodSecurityPolicyReview)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_security_PodSecurityPolicyReview_To_v1_PodSecurityPolicyReview(a.(*security.PodSecurityPolicyReview), b.(*v1.PodSecurityPolicyReview), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.PodSecurityPolicyReviewSpec)(nil), (*security.PodSecurityPolicyReviewSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_PodSecurityPolicyReviewSpec_To_security_PodSecurityPolicyReviewSpec(a.(*v1.PodSecurityPolicyReviewSpec), b.(*security.PodSecurityPolicyReviewSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*security.PodSecurityPolicyReviewSpec)(nil), (*v1.PodSecurityPolicyReviewSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_security_PodSecurityPolicyReviewSpec_To_v1_PodSecurityPolicyReviewSpec(a.(*security.PodSecurityPolicyReviewSpec), b.(*v1.PodSecurityPolicyReviewSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.PodSecurityPolicyReviewStatus)(nil), (*security.PodSecurityPolicyReviewStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_PodSecurityPolicyReviewStatus_To_security_PodSecurityPolicyReviewStatus(a.(*v1.PodSecurityPolicyReviewStatus), b.(*security.PodSecurityPolicyReviewStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*security.PodSecurityPolicyReviewStatus)(nil), (*v1.PodSecurityPolicyReviewStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_security_PodSecurityPolicyReviewStatus_To_v1_PodSecurityPolicyReviewStatus(a.(*security.PodSecurityPolicyReviewStatus), b.(*v1.PodSecurityPolicyReviewStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.PodSecurityPolicySelfSubjectReview)(nil), (*security.PodSecurityPolicySelfSubjectReview)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_PodSecurityPolicySelfSubjectReview_To_security_PodSecurityPolicySelfSubjectReview(a.(*v1.PodSecurityPolicySelfSubjectReview), b.(*security.PodSecurityPolicySelfSubjectReview), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*security.PodSecurityPolicySelfSubjectReview)(nil), (*v1.PodSecurityPolicySelfSubjectReview)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_security_PodSecurityPolicySelfSubjectReview_To_v1_PodSecurityPolicySelfSubjectReview(a.(*security.PodSecurityPolicySelfSubjectReview), b.(*v1.PodSecurityPolicySelfSubjectReview), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.PodSecurityPolicySelfSubjectReviewSpec)(nil), (*security.PodSecurityPolicySelfSubjectReviewSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_PodSecurityPolicySelfSubjectReviewSpec_To_security_PodSecurityPolicySelfSubjectReviewSpec(a.(*v1.PodSecurityPolicySelfSubjectReviewSpec), b.(*security.PodSecurityPolicySelfSubjectReviewSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*security.PodSecurityPolicySelfSubjectReviewSpec)(nil), (*v1.PodSecurityPolicySelfSubjectReviewSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_security_PodSecurityPolicySelfSubjectReviewSpec_To_v1_PodSecurityPolicySelfSubjectReviewSpec(a.(*security.PodSecurityPolicySelfSubjectReviewSpec), b.(*v1.PodSecurityPolicySelfSubjectReviewSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.PodSecurityPolicySubjectReview)(nil), (*security.PodSecurityPolicySubjectReview)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_PodSecurityPolicySubjectReview_To_security_PodSecurityPolicySubjectReview(a.(*v1.PodSecurityPolicySubjectReview), b.(*security.PodSecurityPolicySubjectReview), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*security.PodSecurityPolicySubjectReview)(nil), (*v1.PodSecurityPolicySubjectReview)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_security_PodSecurityPolicySubjectReview_To_v1_PodSecurityPolicySubjectReview(a.(*security.PodSecurityPolicySubjectReview), b.(*v1.PodSecurityPolicySubjectReview), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.PodSecurityPolicySubjectReviewSpec)(nil), (*security.PodSecurityPolicySubjectReviewSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_PodSecurityPolicySubjectReviewSpec_To_security_PodSecurityPolicySubjectReviewSpec(a.(*v1.PodSecurityPolicySubjectReviewSpec), b.(*security.PodSecurityPolicySubjectReviewSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*security.PodSecurityPolicySubjectReviewSpec)(nil), (*v1.PodSecurityPolicySubjectReviewSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_security_PodSecurityPolicySubjectReviewSpec_To_v1_PodSecurityPolicySubjectReviewSpec(a.(*security.PodSecurityPolicySubjectReviewSpec), b.(*v1.PodSecurityPolicySubjectReviewSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.PodSecurityPolicySubjectReviewStatus)(nil), (*security.PodSecurityPolicySubjectReviewStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_PodSecurityPolicySubjectReviewStatus_To_security_PodSecurityPolicySubjectReviewStatus(a.(*v1.PodSecurityPolicySubjectReviewStatus), b.(*security.PodSecurityPolicySubjectReviewStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*security.PodSecurityPolicySubjectReviewStatus)(nil), (*v1.PodSecurityPolicySubjectReviewStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_security_PodSecurityPolicySubjectReviewStatus_To_v1_PodSecurityPolicySubjectReviewStatus(a.(*security.PodSecurityPolicySubjectReviewStatus), b.(*v1.PodSecurityPolicySubjectReviewStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.RangeAllocation)(nil), (*security.RangeAllocation)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_RangeAllocation_To_security_RangeAllocation(a.(*v1.RangeAllocation), b.(*security.RangeAllocation), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*security.RangeAllocation)(nil), (*v1.RangeAllocation)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_security_RangeAllocation_To_v1_RangeAllocation(a.(*security.RangeAllocation), b.(*v1.RangeAllocation), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.RangeAllocationList)(nil), (*security.RangeAllocationList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_RangeAllocationList_To_security_RangeAllocationList(a.(*v1.RangeAllocationList), b.(*security.RangeAllocationList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*security.RangeAllocationList)(nil), (*v1.RangeAllocationList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_security_RangeAllocationList_To_v1_RangeAllocationList(a.(*security.RangeAllocationList), b.(*v1.RangeAllocationList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.RunAsUserStrategyOptions)(nil), (*security.RunAsUserStrategyOptions)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_RunAsUserStrategyOptions_To_security_RunAsUserStrategyOptions(a.(*v1.RunAsUserStrategyOptions), b.(*security.RunAsUserStrategyOptions), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*security.RunAsUserStrategyOptions)(nil), (*v1.RunAsUserStrategyOptions)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_security_RunAsUserStrategyOptions_To_v1_RunAsUserStrategyOptions(a.(*security.RunAsUserStrategyOptions), b.(*v1.RunAsUserStrategyOptions), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.SELinuxContextStrategyOptions)(nil), (*security.SELinuxContextStrategyOptions)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_SELinuxContextStrategyOptions_To_security_SELinuxContextStrategyOptions(a.(*v1.SELinuxContextStrategyOptions), b.(*security.SELinuxContextStrategyOptions), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*security.SELinuxContextStrategyOptions)(nil), (*v1.SELinuxContextStrategyOptions)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_security_SELinuxContextStrategyOptions_To_v1_SELinuxContextStrategyOptions(a.(*security.SELinuxContextStrategyOptions), b.(*v1.SELinuxContextStrategyOptions), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.SecurityContextConstraints)(nil), (*security.SecurityContextConstraints)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_SecurityContextConstraints_To_security_SecurityContextConstraints(a.(*v1.SecurityContextConstraints), b.(*security.SecurityContextConstraints), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*security.SecurityContextConstraints)(nil), (*v1.SecurityContextConstraints)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_security_SecurityContextConstraints_To_v1_SecurityContextConstraints(a.(*security.SecurityContextConstraints), b.(*v1.SecurityContextConstraints), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.SecurityContextConstraintsList)(nil), (*security.SecurityContextConstraintsList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_SecurityContextConstraintsList_To_security_SecurityContextConstraintsList(a.(*v1.SecurityContextConstraintsList), b.(*security.SecurityContextConstraintsList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*security.SecurityContextConstraintsList)(nil), (*v1.SecurityContextConstraintsList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_security_SecurityContextConstraintsList_To_v1_SecurityContextConstraintsList(a.(*security.SecurityContextConstraintsList), b.(*v1.SecurityContextConstraintsList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.ServiceAccountPodSecurityPolicyReviewStatus)(nil), (*security.ServiceAccountPodSecurityPolicyReviewStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_ServiceAccountPodSecurityPolicyReviewStatus_To_security_ServiceAccountPodSecurityPolicyReviewStatus(a.(*v1.ServiceAccountPodSecurityPolicyReviewStatus), b.(*security.ServiceAccountPodSecurityPolicyReviewStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*security.ServiceAccountPodSecurityPolicyReviewStatus)(nil), (*v1.ServiceAccountPodSecurityPolicyReviewStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_security_ServiceAccountPodSecurityPolicyReviewStatus_To_v1_ServiceAccountPodSecurityPolicyReviewStatus(a.(*security.ServiceAccountPodSecurityPolicyReviewStatus), b.(*v1.ServiceAccountPodSecurityPolicyReviewStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.SupplementalGroupsStrategyOptions)(nil), (*security.SupplementalGroupsStrategyOptions)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_SupplementalGroupsStrategyOptions_To_security_SupplementalGroupsStrategyOptions(a.(*v1.SupplementalGroupsStrategyOptions), b.(*security.SupplementalGroupsStrategyOptions), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*security.SupplementalGroupsStrategyOptions)(nil), (*v1.SupplementalGroupsStrategyOptions)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_security_SupplementalGroupsStrategyOptions_To_v1_SupplementalGroupsStrategyOptions(a.(*security.SupplementalGroupsStrategyOptions), b.(*v1.SupplementalGroupsStrategyOptions), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*security.SecurityContextConstraints)(nil), (*v1.SecurityContextConstraints)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_security_SecurityContextConstraints_To_v1_SecurityContextConstraints(a.(*security.SecurityContextConstraints), b.(*v1.SecurityContextConstraints), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*v1.SecurityContextConstraints)(nil), (*security.SecurityContextConstraints)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_SecurityContextConstraints_To_security_SecurityContextConstraints(a.(*v1.SecurityContextConstraints), b.(*security.SecurityContextConstraints), scope)
	}); err != nil {
		return err
	}
	return nil
}
func autoConvert_v1_AllowedFlexVolume_To_security_AllowedFlexVolume(in *v1.AllowedFlexVolume, out *security.AllowedFlexVolume, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Driver = in.Driver
	return nil
}
func Convert_v1_AllowedFlexVolume_To_security_AllowedFlexVolume(in *v1.AllowedFlexVolume, out *security.AllowedFlexVolume, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_AllowedFlexVolume_To_security_AllowedFlexVolume(in, out, s)
}
func autoConvert_security_AllowedFlexVolume_To_v1_AllowedFlexVolume(in *security.AllowedFlexVolume, out *v1.AllowedFlexVolume, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Driver = in.Driver
	return nil
}
func Convert_security_AllowedFlexVolume_To_v1_AllowedFlexVolume(in *security.AllowedFlexVolume, out *v1.AllowedFlexVolume, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_security_AllowedFlexVolume_To_v1_AllowedFlexVolume(in, out, s)
}
func autoConvert_v1_FSGroupStrategyOptions_To_security_FSGroupStrategyOptions(in *v1.FSGroupStrategyOptions, out *security.FSGroupStrategyOptions, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Type = security.FSGroupStrategyType(in.Type)
	out.Ranges = *(*[]security.IDRange)(unsafe.Pointer(&in.Ranges))
	return nil
}
func Convert_v1_FSGroupStrategyOptions_To_security_FSGroupStrategyOptions(in *v1.FSGroupStrategyOptions, out *security.FSGroupStrategyOptions, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_FSGroupStrategyOptions_To_security_FSGroupStrategyOptions(in, out, s)
}
func autoConvert_security_FSGroupStrategyOptions_To_v1_FSGroupStrategyOptions(in *security.FSGroupStrategyOptions, out *v1.FSGroupStrategyOptions, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Type = v1.FSGroupStrategyType(in.Type)
	out.Ranges = *(*[]v1.IDRange)(unsafe.Pointer(&in.Ranges))
	return nil
}
func Convert_security_FSGroupStrategyOptions_To_v1_FSGroupStrategyOptions(in *security.FSGroupStrategyOptions, out *v1.FSGroupStrategyOptions, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_security_FSGroupStrategyOptions_To_v1_FSGroupStrategyOptions(in, out, s)
}
func autoConvert_v1_IDRange_To_security_IDRange(in *v1.IDRange, out *security.IDRange, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Min = in.Min
	out.Max = in.Max
	return nil
}
func Convert_v1_IDRange_To_security_IDRange(in *v1.IDRange, out *security.IDRange, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_IDRange_To_security_IDRange(in, out, s)
}
func autoConvert_security_IDRange_To_v1_IDRange(in *security.IDRange, out *v1.IDRange, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Min = in.Min
	out.Max = in.Max
	return nil
}
func Convert_security_IDRange_To_v1_IDRange(in *security.IDRange, out *v1.IDRange, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_security_IDRange_To_v1_IDRange(in, out, s)
}
func autoConvert_v1_PodSecurityPolicyReview_To_security_PodSecurityPolicyReview(in *v1.PodSecurityPolicyReview, out *security.PodSecurityPolicyReview, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := Convert_v1_PodSecurityPolicyReviewSpec_To_security_PodSecurityPolicyReviewSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1_PodSecurityPolicyReviewStatus_To_security_PodSecurityPolicyReviewStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}
func Convert_v1_PodSecurityPolicyReview_To_security_PodSecurityPolicyReview(in *v1.PodSecurityPolicyReview, out *security.PodSecurityPolicyReview, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_PodSecurityPolicyReview_To_security_PodSecurityPolicyReview(in, out, s)
}
func autoConvert_security_PodSecurityPolicyReview_To_v1_PodSecurityPolicyReview(in *security.PodSecurityPolicyReview, out *v1.PodSecurityPolicyReview, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := Convert_security_PodSecurityPolicyReviewSpec_To_v1_PodSecurityPolicyReviewSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_security_PodSecurityPolicyReviewStatus_To_v1_PodSecurityPolicyReviewStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}
func Convert_security_PodSecurityPolicyReview_To_v1_PodSecurityPolicyReview(in *security.PodSecurityPolicyReview, out *v1.PodSecurityPolicyReview, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_security_PodSecurityPolicyReview_To_v1_PodSecurityPolicyReview(in, out, s)
}
func autoConvert_v1_PodSecurityPolicyReviewSpec_To_security_PodSecurityPolicyReviewSpec(in *v1.PodSecurityPolicyReviewSpec, out *security.PodSecurityPolicyReviewSpec, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := corev1.Convert_v1_PodTemplateSpec_To_core_PodTemplateSpec(&in.Template, &out.Template, s); err != nil {
		return err
	}
	out.ServiceAccountNames = *(*[]string)(unsafe.Pointer(&in.ServiceAccountNames))
	return nil
}
func Convert_v1_PodSecurityPolicyReviewSpec_To_security_PodSecurityPolicyReviewSpec(in *v1.PodSecurityPolicyReviewSpec, out *security.PodSecurityPolicyReviewSpec, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_PodSecurityPolicyReviewSpec_To_security_PodSecurityPolicyReviewSpec(in, out, s)
}
func autoConvert_security_PodSecurityPolicyReviewSpec_To_v1_PodSecurityPolicyReviewSpec(in *security.PodSecurityPolicyReviewSpec, out *v1.PodSecurityPolicyReviewSpec, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := corev1.Convert_core_PodTemplateSpec_To_v1_PodTemplateSpec(&in.Template, &out.Template, s); err != nil {
		return err
	}
	out.ServiceAccountNames = *(*[]string)(unsafe.Pointer(&in.ServiceAccountNames))
	return nil
}
func Convert_security_PodSecurityPolicyReviewSpec_To_v1_PodSecurityPolicyReviewSpec(in *security.PodSecurityPolicyReviewSpec, out *v1.PodSecurityPolicyReviewSpec, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_security_PodSecurityPolicyReviewSpec_To_v1_PodSecurityPolicyReviewSpec(in, out, s)
}
func autoConvert_v1_PodSecurityPolicyReviewStatus_To_security_PodSecurityPolicyReviewStatus(in *v1.PodSecurityPolicyReviewStatus, out *security.PodSecurityPolicyReviewStatus, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in.AllowedServiceAccounts != nil {
		in, out := &in.AllowedServiceAccounts, &out.AllowedServiceAccounts
		*out = make([]security.ServiceAccountPodSecurityPolicyReviewStatus, len(*in))
		for i := range *in {
			if err := Convert_v1_ServiceAccountPodSecurityPolicyReviewStatus_To_security_ServiceAccountPodSecurityPolicyReviewStatus(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.AllowedServiceAccounts = nil
	}
	return nil
}
func Convert_v1_PodSecurityPolicyReviewStatus_To_security_PodSecurityPolicyReviewStatus(in *v1.PodSecurityPolicyReviewStatus, out *security.PodSecurityPolicyReviewStatus, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_PodSecurityPolicyReviewStatus_To_security_PodSecurityPolicyReviewStatus(in, out, s)
}
func autoConvert_security_PodSecurityPolicyReviewStatus_To_v1_PodSecurityPolicyReviewStatus(in *security.PodSecurityPolicyReviewStatus, out *v1.PodSecurityPolicyReviewStatus, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in.AllowedServiceAccounts != nil {
		in, out := &in.AllowedServiceAccounts, &out.AllowedServiceAccounts
		*out = make([]v1.ServiceAccountPodSecurityPolicyReviewStatus, len(*in))
		for i := range *in {
			if err := Convert_security_ServiceAccountPodSecurityPolicyReviewStatus_To_v1_ServiceAccountPodSecurityPolicyReviewStatus(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.AllowedServiceAccounts = nil
	}
	return nil
}
func Convert_security_PodSecurityPolicyReviewStatus_To_v1_PodSecurityPolicyReviewStatus(in *security.PodSecurityPolicyReviewStatus, out *v1.PodSecurityPolicyReviewStatus, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_security_PodSecurityPolicyReviewStatus_To_v1_PodSecurityPolicyReviewStatus(in, out, s)
}
func autoConvert_v1_PodSecurityPolicySelfSubjectReview_To_security_PodSecurityPolicySelfSubjectReview(in *v1.PodSecurityPolicySelfSubjectReview, out *security.PodSecurityPolicySelfSubjectReview, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := Convert_v1_PodSecurityPolicySelfSubjectReviewSpec_To_security_PodSecurityPolicySelfSubjectReviewSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1_PodSecurityPolicySubjectReviewStatus_To_security_PodSecurityPolicySubjectReviewStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}
func Convert_v1_PodSecurityPolicySelfSubjectReview_To_security_PodSecurityPolicySelfSubjectReview(in *v1.PodSecurityPolicySelfSubjectReview, out *security.PodSecurityPolicySelfSubjectReview, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_PodSecurityPolicySelfSubjectReview_To_security_PodSecurityPolicySelfSubjectReview(in, out, s)
}
func autoConvert_security_PodSecurityPolicySelfSubjectReview_To_v1_PodSecurityPolicySelfSubjectReview(in *security.PodSecurityPolicySelfSubjectReview, out *v1.PodSecurityPolicySelfSubjectReview, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := Convert_security_PodSecurityPolicySelfSubjectReviewSpec_To_v1_PodSecurityPolicySelfSubjectReviewSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_security_PodSecurityPolicySubjectReviewStatus_To_v1_PodSecurityPolicySubjectReviewStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}
func Convert_security_PodSecurityPolicySelfSubjectReview_To_v1_PodSecurityPolicySelfSubjectReview(in *security.PodSecurityPolicySelfSubjectReview, out *v1.PodSecurityPolicySelfSubjectReview, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_security_PodSecurityPolicySelfSubjectReview_To_v1_PodSecurityPolicySelfSubjectReview(in, out, s)
}
func autoConvert_v1_PodSecurityPolicySelfSubjectReviewSpec_To_security_PodSecurityPolicySelfSubjectReviewSpec(in *v1.PodSecurityPolicySelfSubjectReviewSpec, out *security.PodSecurityPolicySelfSubjectReviewSpec, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := corev1.Convert_v1_PodTemplateSpec_To_core_PodTemplateSpec(&in.Template, &out.Template, s); err != nil {
		return err
	}
	return nil
}
func Convert_v1_PodSecurityPolicySelfSubjectReviewSpec_To_security_PodSecurityPolicySelfSubjectReviewSpec(in *v1.PodSecurityPolicySelfSubjectReviewSpec, out *security.PodSecurityPolicySelfSubjectReviewSpec, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_PodSecurityPolicySelfSubjectReviewSpec_To_security_PodSecurityPolicySelfSubjectReviewSpec(in, out, s)
}
func autoConvert_security_PodSecurityPolicySelfSubjectReviewSpec_To_v1_PodSecurityPolicySelfSubjectReviewSpec(in *security.PodSecurityPolicySelfSubjectReviewSpec, out *v1.PodSecurityPolicySelfSubjectReviewSpec, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := corev1.Convert_core_PodTemplateSpec_To_v1_PodTemplateSpec(&in.Template, &out.Template, s); err != nil {
		return err
	}
	return nil
}
func Convert_security_PodSecurityPolicySelfSubjectReviewSpec_To_v1_PodSecurityPolicySelfSubjectReviewSpec(in *security.PodSecurityPolicySelfSubjectReviewSpec, out *v1.PodSecurityPolicySelfSubjectReviewSpec, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_security_PodSecurityPolicySelfSubjectReviewSpec_To_v1_PodSecurityPolicySelfSubjectReviewSpec(in, out, s)
}
func autoConvert_v1_PodSecurityPolicySubjectReview_To_security_PodSecurityPolicySubjectReview(in *v1.PodSecurityPolicySubjectReview, out *security.PodSecurityPolicySubjectReview, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := Convert_v1_PodSecurityPolicySubjectReviewSpec_To_security_PodSecurityPolicySubjectReviewSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1_PodSecurityPolicySubjectReviewStatus_To_security_PodSecurityPolicySubjectReviewStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}
func Convert_v1_PodSecurityPolicySubjectReview_To_security_PodSecurityPolicySubjectReview(in *v1.PodSecurityPolicySubjectReview, out *security.PodSecurityPolicySubjectReview, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_PodSecurityPolicySubjectReview_To_security_PodSecurityPolicySubjectReview(in, out, s)
}
func autoConvert_security_PodSecurityPolicySubjectReview_To_v1_PodSecurityPolicySubjectReview(in *security.PodSecurityPolicySubjectReview, out *v1.PodSecurityPolicySubjectReview, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := Convert_security_PodSecurityPolicySubjectReviewSpec_To_v1_PodSecurityPolicySubjectReviewSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_security_PodSecurityPolicySubjectReviewStatus_To_v1_PodSecurityPolicySubjectReviewStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}
func Convert_security_PodSecurityPolicySubjectReview_To_v1_PodSecurityPolicySubjectReview(in *security.PodSecurityPolicySubjectReview, out *v1.PodSecurityPolicySubjectReview, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_security_PodSecurityPolicySubjectReview_To_v1_PodSecurityPolicySubjectReview(in, out, s)
}
func autoConvert_v1_PodSecurityPolicySubjectReviewSpec_To_security_PodSecurityPolicySubjectReviewSpec(in *v1.PodSecurityPolicySubjectReviewSpec, out *security.PodSecurityPolicySubjectReviewSpec, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := corev1.Convert_v1_PodTemplateSpec_To_core_PodTemplateSpec(&in.Template, &out.Template, s); err != nil {
		return err
	}
	out.User = in.User
	out.Groups = *(*[]string)(unsafe.Pointer(&in.Groups))
	return nil
}
func Convert_v1_PodSecurityPolicySubjectReviewSpec_To_security_PodSecurityPolicySubjectReviewSpec(in *v1.PodSecurityPolicySubjectReviewSpec, out *security.PodSecurityPolicySubjectReviewSpec, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_PodSecurityPolicySubjectReviewSpec_To_security_PodSecurityPolicySubjectReviewSpec(in, out, s)
}
func autoConvert_security_PodSecurityPolicySubjectReviewSpec_To_v1_PodSecurityPolicySubjectReviewSpec(in *security.PodSecurityPolicySubjectReviewSpec, out *v1.PodSecurityPolicySubjectReviewSpec, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := corev1.Convert_core_PodTemplateSpec_To_v1_PodTemplateSpec(&in.Template, &out.Template, s); err != nil {
		return err
	}
	out.User = in.User
	out.Groups = *(*[]string)(unsafe.Pointer(&in.Groups))
	return nil
}
func Convert_security_PodSecurityPolicySubjectReviewSpec_To_v1_PodSecurityPolicySubjectReviewSpec(in *security.PodSecurityPolicySubjectReviewSpec, out *v1.PodSecurityPolicySubjectReviewSpec, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_security_PodSecurityPolicySubjectReviewSpec_To_v1_PodSecurityPolicySubjectReviewSpec(in, out, s)
}
func autoConvert_v1_PodSecurityPolicySubjectReviewStatus_To_security_PodSecurityPolicySubjectReviewStatus(in *v1.PodSecurityPolicySubjectReviewStatus, out *security.PodSecurityPolicySubjectReviewStatus, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in.AllowedBy != nil {
		in, out := &in.AllowedBy, &out.AllowedBy
		*out = new(core.ObjectReference)
		if err := corev1.Convert_v1_ObjectReference_To_core_ObjectReference(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.AllowedBy = nil
	}
	out.Reason = in.Reason
	if err := corev1.Convert_v1_PodTemplateSpec_To_core_PodTemplateSpec(&in.Template, &out.Template, s); err != nil {
		return err
	}
	return nil
}
func Convert_v1_PodSecurityPolicySubjectReviewStatus_To_security_PodSecurityPolicySubjectReviewStatus(in *v1.PodSecurityPolicySubjectReviewStatus, out *security.PodSecurityPolicySubjectReviewStatus, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_PodSecurityPolicySubjectReviewStatus_To_security_PodSecurityPolicySubjectReviewStatus(in, out, s)
}
func autoConvert_security_PodSecurityPolicySubjectReviewStatus_To_v1_PodSecurityPolicySubjectReviewStatus(in *security.PodSecurityPolicySubjectReviewStatus, out *v1.PodSecurityPolicySubjectReviewStatus, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in.AllowedBy != nil {
		in, out := &in.AllowedBy, &out.AllowedBy
		*out = new(apicorev1.ObjectReference)
		if err := corev1.Convert_core_ObjectReference_To_v1_ObjectReference(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.AllowedBy = nil
	}
	out.Reason = in.Reason
	if err := corev1.Convert_core_PodTemplateSpec_To_v1_PodTemplateSpec(&in.Template, &out.Template, s); err != nil {
		return err
	}
	return nil
}
func Convert_security_PodSecurityPolicySubjectReviewStatus_To_v1_PodSecurityPolicySubjectReviewStatus(in *security.PodSecurityPolicySubjectReviewStatus, out *v1.PodSecurityPolicySubjectReviewStatus, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_security_PodSecurityPolicySubjectReviewStatus_To_v1_PodSecurityPolicySubjectReviewStatus(in, out, s)
}
func autoConvert_v1_RangeAllocation_To_security_RangeAllocation(in *v1.RangeAllocation, out *security.RangeAllocation, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ObjectMeta = in.ObjectMeta
	out.Range = in.Range
	out.Data = *(*[]byte)(unsafe.Pointer(&in.Data))
	return nil
}
func Convert_v1_RangeAllocation_To_security_RangeAllocation(in *v1.RangeAllocation, out *security.RangeAllocation, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_RangeAllocation_To_security_RangeAllocation(in, out, s)
}
func autoConvert_security_RangeAllocation_To_v1_RangeAllocation(in *security.RangeAllocation, out *v1.RangeAllocation, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ObjectMeta = in.ObjectMeta
	out.Range = in.Range
	out.Data = *(*[]byte)(unsafe.Pointer(&in.Data))
	return nil
}
func Convert_security_RangeAllocation_To_v1_RangeAllocation(in *security.RangeAllocation, out *v1.RangeAllocation, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_security_RangeAllocation_To_v1_RangeAllocation(in, out, s)
}
func autoConvert_v1_RangeAllocationList_To_security_RangeAllocationList(in *v1.RangeAllocationList, out *security.RangeAllocationList, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ListMeta = in.ListMeta
	out.Items = *(*[]security.RangeAllocation)(unsafe.Pointer(&in.Items))
	return nil
}
func Convert_v1_RangeAllocationList_To_security_RangeAllocationList(in *v1.RangeAllocationList, out *security.RangeAllocationList, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_RangeAllocationList_To_security_RangeAllocationList(in, out, s)
}
func autoConvert_security_RangeAllocationList_To_v1_RangeAllocationList(in *security.RangeAllocationList, out *v1.RangeAllocationList, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ListMeta = in.ListMeta
	out.Items = *(*[]v1.RangeAllocation)(unsafe.Pointer(&in.Items))
	return nil
}
func Convert_security_RangeAllocationList_To_v1_RangeAllocationList(in *security.RangeAllocationList, out *v1.RangeAllocationList, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_security_RangeAllocationList_To_v1_RangeAllocationList(in, out, s)
}
func autoConvert_v1_RunAsUserStrategyOptions_To_security_RunAsUserStrategyOptions(in *v1.RunAsUserStrategyOptions, out *security.RunAsUserStrategyOptions, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Type = security.RunAsUserStrategyType(in.Type)
	out.UID = (*int64)(unsafe.Pointer(in.UID))
	out.UIDRangeMin = (*int64)(unsafe.Pointer(in.UIDRangeMin))
	out.UIDRangeMax = (*int64)(unsafe.Pointer(in.UIDRangeMax))
	return nil
}
func Convert_v1_RunAsUserStrategyOptions_To_security_RunAsUserStrategyOptions(in *v1.RunAsUserStrategyOptions, out *security.RunAsUserStrategyOptions, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_RunAsUserStrategyOptions_To_security_RunAsUserStrategyOptions(in, out, s)
}
func autoConvert_security_RunAsUserStrategyOptions_To_v1_RunAsUserStrategyOptions(in *security.RunAsUserStrategyOptions, out *v1.RunAsUserStrategyOptions, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Type = v1.RunAsUserStrategyType(in.Type)
	out.UID = (*int64)(unsafe.Pointer(in.UID))
	out.UIDRangeMin = (*int64)(unsafe.Pointer(in.UIDRangeMin))
	out.UIDRangeMax = (*int64)(unsafe.Pointer(in.UIDRangeMax))
	return nil
}
func Convert_security_RunAsUserStrategyOptions_To_v1_RunAsUserStrategyOptions(in *security.RunAsUserStrategyOptions, out *v1.RunAsUserStrategyOptions, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_security_RunAsUserStrategyOptions_To_v1_RunAsUserStrategyOptions(in, out, s)
}
func autoConvert_v1_SELinuxContextStrategyOptions_To_security_SELinuxContextStrategyOptions(in *v1.SELinuxContextStrategyOptions, out *security.SELinuxContextStrategyOptions, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Type = security.SELinuxContextStrategyType(in.Type)
	if in.SELinuxOptions != nil {
		in, out := &in.SELinuxOptions, &out.SELinuxOptions
		*out = new(core.SELinuxOptions)
		if err := corev1.Convert_v1_SELinuxOptions_To_core_SELinuxOptions(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.SELinuxOptions = nil
	}
	return nil
}
func Convert_v1_SELinuxContextStrategyOptions_To_security_SELinuxContextStrategyOptions(in *v1.SELinuxContextStrategyOptions, out *security.SELinuxContextStrategyOptions, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_SELinuxContextStrategyOptions_To_security_SELinuxContextStrategyOptions(in, out, s)
}
func autoConvert_security_SELinuxContextStrategyOptions_To_v1_SELinuxContextStrategyOptions(in *security.SELinuxContextStrategyOptions, out *v1.SELinuxContextStrategyOptions, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Type = v1.SELinuxContextStrategyType(in.Type)
	if in.SELinuxOptions != nil {
		in, out := &in.SELinuxOptions, &out.SELinuxOptions
		*out = new(apicorev1.SELinuxOptions)
		if err := corev1.Convert_core_SELinuxOptions_To_v1_SELinuxOptions(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.SELinuxOptions = nil
	}
	return nil
}
func Convert_security_SELinuxContextStrategyOptions_To_v1_SELinuxContextStrategyOptions(in *security.SELinuxContextStrategyOptions, out *v1.SELinuxContextStrategyOptions, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_security_SELinuxContextStrategyOptions_To_v1_SELinuxContextStrategyOptions(in, out, s)
}
func autoConvert_v1_SecurityContextConstraints_To_security_SecurityContextConstraints(in *v1.SecurityContextConstraints, out *security.SecurityContextConstraints, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ObjectMeta = in.ObjectMeta
	out.Priority = (*int32)(unsafe.Pointer(in.Priority))
	out.AllowPrivilegedContainer = in.AllowPrivilegedContainer
	out.DefaultAddCapabilities = *(*[]core.Capability)(unsafe.Pointer(&in.DefaultAddCapabilities))
	out.RequiredDropCapabilities = *(*[]core.Capability)(unsafe.Pointer(&in.RequiredDropCapabilities))
	out.AllowedCapabilities = *(*[]core.Capability)(unsafe.Pointer(&in.AllowedCapabilities))
	out.Volumes = *(*[]security.FSType)(unsafe.Pointer(&in.Volumes))
	out.AllowedFlexVolumes = *(*[]security.AllowedFlexVolume)(unsafe.Pointer(&in.AllowedFlexVolumes))
	out.AllowHostNetwork = in.AllowHostNetwork
	out.AllowHostPorts = in.AllowHostPorts
	out.AllowHostPID = in.AllowHostPID
	out.AllowHostIPC = in.AllowHostIPC
	out.DefaultAllowPrivilegeEscalation = (*bool)(unsafe.Pointer(in.DefaultAllowPrivilegeEscalation))
	out.AllowPrivilegeEscalation = (*bool)(unsafe.Pointer(in.AllowPrivilegeEscalation))
	if err := Convert_v1_SELinuxContextStrategyOptions_To_security_SELinuxContextStrategyOptions(&in.SELinuxContext, &out.SELinuxContext, s); err != nil {
		return err
	}
	if err := Convert_v1_RunAsUserStrategyOptions_To_security_RunAsUserStrategyOptions(&in.RunAsUser, &out.RunAsUser, s); err != nil {
		return err
	}
	if err := Convert_v1_SupplementalGroupsStrategyOptions_To_security_SupplementalGroupsStrategyOptions(&in.SupplementalGroups, &out.SupplementalGroups, s); err != nil {
		return err
	}
	if err := Convert_v1_FSGroupStrategyOptions_To_security_FSGroupStrategyOptions(&in.FSGroup, &out.FSGroup, s); err != nil {
		return err
	}
	out.ReadOnlyRootFilesystem = in.ReadOnlyRootFilesystem
	out.Users = *(*[]string)(unsafe.Pointer(&in.Users))
	out.Groups = *(*[]string)(unsafe.Pointer(&in.Groups))
	out.SeccompProfiles = *(*[]string)(unsafe.Pointer(&in.SeccompProfiles))
	out.AllowedUnsafeSysctls = *(*[]string)(unsafe.Pointer(&in.AllowedUnsafeSysctls))
	out.ForbiddenSysctls = *(*[]string)(unsafe.Pointer(&in.ForbiddenSysctls))
	return nil
}
func autoConvert_security_SecurityContextConstraints_To_v1_SecurityContextConstraints(in *security.SecurityContextConstraints, out *v1.SecurityContextConstraints, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ObjectMeta = in.ObjectMeta
	out.Priority = (*int32)(unsafe.Pointer(in.Priority))
	out.AllowPrivilegedContainer = in.AllowPrivilegedContainer
	out.DefaultAddCapabilities = *(*[]apicorev1.Capability)(unsafe.Pointer(&in.DefaultAddCapabilities))
	out.RequiredDropCapabilities = *(*[]apicorev1.Capability)(unsafe.Pointer(&in.RequiredDropCapabilities))
	out.AllowedCapabilities = *(*[]apicorev1.Capability)(unsafe.Pointer(&in.AllowedCapabilities))
	out.Volumes = *(*[]v1.FSType)(unsafe.Pointer(&in.Volumes))
	out.AllowedFlexVolumes = *(*[]v1.AllowedFlexVolume)(unsafe.Pointer(&in.AllowedFlexVolumes))
	out.AllowHostNetwork = in.AllowHostNetwork
	out.AllowHostPorts = in.AllowHostPorts
	out.AllowHostPID = in.AllowHostPID
	out.AllowHostIPC = in.AllowHostIPC
	out.DefaultAllowPrivilegeEscalation = (*bool)(unsafe.Pointer(in.DefaultAllowPrivilegeEscalation))
	out.AllowPrivilegeEscalation = (*bool)(unsafe.Pointer(in.AllowPrivilegeEscalation))
	if err := Convert_security_SELinuxContextStrategyOptions_To_v1_SELinuxContextStrategyOptions(&in.SELinuxContext, &out.SELinuxContext, s); err != nil {
		return err
	}
	if err := Convert_security_RunAsUserStrategyOptions_To_v1_RunAsUserStrategyOptions(&in.RunAsUser, &out.RunAsUser, s); err != nil {
		return err
	}
	if err := Convert_security_SupplementalGroupsStrategyOptions_To_v1_SupplementalGroupsStrategyOptions(&in.SupplementalGroups, &out.SupplementalGroups, s); err != nil {
		return err
	}
	if err := Convert_security_FSGroupStrategyOptions_To_v1_FSGroupStrategyOptions(&in.FSGroup, &out.FSGroup, s); err != nil {
		return err
	}
	out.ReadOnlyRootFilesystem = in.ReadOnlyRootFilesystem
	out.SeccompProfiles = *(*[]string)(unsafe.Pointer(&in.SeccompProfiles))
	out.Users = *(*[]string)(unsafe.Pointer(&in.Users))
	out.Groups = *(*[]string)(unsafe.Pointer(&in.Groups))
	out.AllowedUnsafeSysctls = *(*[]string)(unsafe.Pointer(&in.AllowedUnsafeSysctls))
	out.ForbiddenSysctls = *(*[]string)(unsafe.Pointer(&in.ForbiddenSysctls))
	return nil
}
func autoConvert_v1_SecurityContextConstraintsList_To_security_SecurityContextConstraintsList(in *v1.SecurityContextConstraintsList, out *security.SecurityContextConstraintsList, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]security.SecurityContextConstraints, len(*in))
		for i := range *in {
			if err := Convert_v1_SecurityContextConstraints_To_security_SecurityContextConstraints(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}
func Convert_v1_SecurityContextConstraintsList_To_security_SecurityContextConstraintsList(in *v1.SecurityContextConstraintsList, out *security.SecurityContextConstraintsList, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_SecurityContextConstraintsList_To_security_SecurityContextConstraintsList(in, out, s)
}
func autoConvert_security_SecurityContextConstraintsList_To_v1_SecurityContextConstraintsList(in *security.SecurityContextConstraintsList, out *v1.SecurityContextConstraintsList, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]v1.SecurityContextConstraints, len(*in))
		for i := range *in {
			if err := Convert_security_SecurityContextConstraints_To_v1_SecurityContextConstraints(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}
func Convert_security_SecurityContextConstraintsList_To_v1_SecurityContextConstraintsList(in *security.SecurityContextConstraintsList, out *v1.SecurityContextConstraintsList, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_security_SecurityContextConstraintsList_To_v1_SecurityContextConstraintsList(in, out, s)
}
func autoConvert_v1_ServiceAccountPodSecurityPolicyReviewStatus_To_security_ServiceAccountPodSecurityPolicyReviewStatus(in *v1.ServiceAccountPodSecurityPolicyReviewStatus, out *security.ServiceAccountPodSecurityPolicyReviewStatus, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := Convert_v1_PodSecurityPolicySubjectReviewStatus_To_security_PodSecurityPolicySubjectReviewStatus(&in.PodSecurityPolicySubjectReviewStatus, &out.PodSecurityPolicySubjectReviewStatus, s); err != nil {
		return err
	}
	out.Name = in.Name
	return nil
}
func Convert_v1_ServiceAccountPodSecurityPolicyReviewStatus_To_security_ServiceAccountPodSecurityPolicyReviewStatus(in *v1.ServiceAccountPodSecurityPolicyReviewStatus, out *security.ServiceAccountPodSecurityPolicyReviewStatus, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_ServiceAccountPodSecurityPolicyReviewStatus_To_security_ServiceAccountPodSecurityPolicyReviewStatus(in, out, s)
}
func autoConvert_security_ServiceAccountPodSecurityPolicyReviewStatus_To_v1_ServiceAccountPodSecurityPolicyReviewStatus(in *security.ServiceAccountPodSecurityPolicyReviewStatus, out *v1.ServiceAccountPodSecurityPolicyReviewStatus, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := Convert_security_PodSecurityPolicySubjectReviewStatus_To_v1_PodSecurityPolicySubjectReviewStatus(&in.PodSecurityPolicySubjectReviewStatus, &out.PodSecurityPolicySubjectReviewStatus, s); err != nil {
		return err
	}
	out.Name = in.Name
	return nil
}
func Convert_security_ServiceAccountPodSecurityPolicyReviewStatus_To_v1_ServiceAccountPodSecurityPolicyReviewStatus(in *security.ServiceAccountPodSecurityPolicyReviewStatus, out *v1.ServiceAccountPodSecurityPolicyReviewStatus, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_security_ServiceAccountPodSecurityPolicyReviewStatus_To_v1_ServiceAccountPodSecurityPolicyReviewStatus(in, out, s)
}
func autoConvert_v1_SupplementalGroupsStrategyOptions_To_security_SupplementalGroupsStrategyOptions(in *v1.SupplementalGroupsStrategyOptions, out *security.SupplementalGroupsStrategyOptions, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Type = security.SupplementalGroupsStrategyType(in.Type)
	out.Ranges = *(*[]security.IDRange)(unsafe.Pointer(&in.Ranges))
	return nil
}
func Convert_v1_SupplementalGroupsStrategyOptions_To_security_SupplementalGroupsStrategyOptions(in *v1.SupplementalGroupsStrategyOptions, out *security.SupplementalGroupsStrategyOptions, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_SupplementalGroupsStrategyOptions_To_security_SupplementalGroupsStrategyOptions(in, out, s)
}
func autoConvert_security_SupplementalGroupsStrategyOptions_To_v1_SupplementalGroupsStrategyOptions(in *security.SupplementalGroupsStrategyOptions, out *v1.SupplementalGroupsStrategyOptions, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Type = v1.SupplementalGroupsStrategyType(in.Type)
	out.Ranges = *(*[]v1.IDRange)(unsafe.Pointer(&in.Ranges))
	return nil
}
func Convert_security_SupplementalGroupsStrategyOptions_To_v1_SupplementalGroupsStrategyOptions(in *security.SupplementalGroupsStrategyOptions, out *v1.SupplementalGroupsStrategyOptions, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_security_SupplementalGroupsStrategyOptions_To_v1_SupplementalGroupsStrategyOptions(in, out, s)
}
