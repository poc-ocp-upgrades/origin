package v1

import (
	unsafe "unsafe"
	v1 "github.com/openshift/api/quota/v1"
	quota "github.com/openshift/origin/pkg/quota/apis/quota"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
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
	if err := s.AddGeneratedConversionFunc((*v1.AppliedClusterResourceQuota)(nil), (*quota.AppliedClusterResourceQuota)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_AppliedClusterResourceQuota_To_quota_AppliedClusterResourceQuota(a.(*v1.AppliedClusterResourceQuota), b.(*quota.AppliedClusterResourceQuota), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*quota.AppliedClusterResourceQuota)(nil), (*v1.AppliedClusterResourceQuota)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_quota_AppliedClusterResourceQuota_To_v1_AppliedClusterResourceQuota(a.(*quota.AppliedClusterResourceQuota), b.(*v1.AppliedClusterResourceQuota), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.AppliedClusterResourceQuotaList)(nil), (*quota.AppliedClusterResourceQuotaList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_AppliedClusterResourceQuotaList_To_quota_AppliedClusterResourceQuotaList(a.(*v1.AppliedClusterResourceQuotaList), b.(*quota.AppliedClusterResourceQuotaList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*quota.AppliedClusterResourceQuotaList)(nil), (*v1.AppliedClusterResourceQuotaList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_quota_AppliedClusterResourceQuotaList_To_v1_AppliedClusterResourceQuotaList(a.(*quota.AppliedClusterResourceQuotaList), b.(*v1.AppliedClusterResourceQuotaList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.ClusterResourceQuota)(nil), (*quota.ClusterResourceQuota)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_ClusterResourceQuota_To_quota_ClusterResourceQuota(a.(*v1.ClusterResourceQuota), b.(*quota.ClusterResourceQuota), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*quota.ClusterResourceQuota)(nil), (*v1.ClusterResourceQuota)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_quota_ClusterResourceQuota_To_v1_ClusterResourceQuota(a.(*quota.ClusterResourceQuota), b.(*v1.ClusterResourceQuota), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.ClusterResourceQuotaList)(nil), (*quota.ClusterResourceQuotaList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_ClusterResourceQuotaList_To_quota_ClusterResourceQuotaList(a.(*v1.ClusterResourceQuotaList), b.(*quota.ClusterResourceQuotaList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*quota.ClusterResourceQuotaList)(nil), (*v1.ClusterResourceQuotaList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_quota_ClusterResourceQuotaList_To_v1_ClusterResourceQuotaList(a.(*quota.ClusterResourceQuotaList), b.(*v1.ClusterResourceQuotaList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.ClusterResourceQuotaSelector)(nil), (*quota.ClusterResourceQuotaSelector)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_ClusterResourceQuotaSelector_To_quota_ClusterResourceQuotaSelector(a.(*v1.ClusterResourceQuotaSelector), b.(*quota.ClusterResourceQuotaSelector), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*quota.ClusterResourceQuotaSelector)(nil), (*v1.ClusterResourceQuotaSelector)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_quota_ClusterResourceQuotaSelector_To_v1_ClusterResourceQuotaSelector(a.(*quota.ClusterResourceQuotaSelector), b.(*v1.ClusterResourceQuotaSelector), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.ClusterResourceQuotaSpec)(nil), (*quota.ClusterResourceQuotaSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_ClusterResourceQuotaSpec_To_quota_ClusterResourceQuotaSpec(a.(*v1.ClusterResourceQuotaSpec), b.(*quota.ClusterResourceQuotaSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*quota.ClusterResourceQuotaSpec)(nil), (*v1.ClusterResourceQuotaSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_quota_ClusterResourceQuotaSpec_To_v1_ClusterResourceQuotaSpec(a.(*quota.ClusterResourceQuotaSpec), b.(*v1.ClusterResourceQuotaSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.ClusterResourceQuotaStatus)(nil), (*quota.ClusterResourceQuotaStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_ClusterResourceQuotaStatus_To_quota_ClusterResourceQuotaStatus(a.(*v1.ClusterResourceQuotaStatus), b.(*quota.ClusterResourceQuotaStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*quota.ClusterResourceQuotaStatus)(nil), (*v1.ClusterResourceQuotaStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_quota_ClusterResourceQuotaStatus_To_v1_ClusterResourceQuotaStatus(a.(*quota.ClusterResourceQuotaStatus), b.(*v1.ClusterResourceQuotaStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*quota.ResourceQuotasStatusByNamespace)(nil), (*v1.ResourceQuotasStatusByNamespace)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_quota_ResourceQuotasStatusByNamespace_To_v1_ResourceQuotasStatusByNamespace(a.(*quota.ResourceQuotasStatusByNamespace), b.(*v1.ResourceQuotasStatusByNamespace), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*v1.ResourceQuotasStatusByNamespace)(nil), (*quota.ResourceQuotasStatusByNamespace)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_ResourceQuotasStatusByNamespace_To_quota_ResourceQuotasStatusByNamespace(a.(*v1.ResourceQuotasStatusByNamespace), b.(*quota.ResourceQuotasStatusByNamespace), scope)
	}); err != nil {
		return err
	}
	return nil
}
func autoConvert_v1_AppliedClusterResourceQuota_To_quota_AppliedClusterResourceQuota(in *v1.AppliedClusterResourceQuota, out *quota.AppliedClusterResourceQuota, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1_ClusterResourceQuotaSpec_To_quota_ClusterResourceQuotaSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1_ClusterResourceQuotaStatus_To_quota_ClusterResourceQuotaStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}
func Convert_v1_AppliedClusterResourceQuota_To_quota_AppliedClusterResourceQuota(in *v1.AppliedClusterResourceQuota, out *quota.AppliedClusterResourceQuota, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_AppliedClusterResourceQuota_To_quota_AppliedClusterResourceQuota(in, out, s)
}
func autoConvert_quota_AppliedClusterResourceQuota_To_v1_AppliedClusterResourceQuota(in *quota.AppliedClusterResourceQuota, out *v1.AppliedClusterResourceQuota, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_quota_ClusterResourceQuotaSpec_To_v1_ClusterResourceQuotaSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_quota_ClusterResourceQuotaStatus_To_v1_ClusterResourceQuotaStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}
func Convert_quota_AppliedClusterResourceQuota_To_v1_AppliedClusterResourceQuota(in *quota.AppliedClusterResourceQuota, out *v1.AppliedClusterResourceQuota, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_quota_AppliedClusterResourceQuota_To_v1_AppliedClusterResourceQuota(in, out, s)
}
func autoConvert_v1_AppliedClusterResourceQuotaList_To_quota_AppliedClusterResourceQuotaList(in *v1.AppliedClusterResourceQuotaList, out *quota.AppliedClusterResourceQuotaList, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]quota.AppliedClusterResourceQuota, len(*in))
		for i := range *in {
			if err := Convert_v1_AppliedClusterResourceQuota_To_quota_AppliedClusterResourceQuota(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}
func Convert_v1_AppliedClusterResourceQuotaList_To_quota_AppliedClusterResourceQuotaList(in *v1.AppliedClusterResourceQuotaList, out *quota.AppliedClusterResourceQuotaList, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_AppliedClusterResourceQuotaList_To_quota_AppliedClusterResourceQuotaList(in, out, s)
}
func autoConvert_quota_AppliedClusterResourceQuotaList_To_v1_AppliedClusterResourceQuotaList(in *quota.AppliedClusterResourceQuotaList, out *v1.AppliedClusterResourceQuotaList, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]v1.AppliedClusterResourceQuota, len(*in))
		for i := range *in {
			if err := Convert_quota_AppliedClusterResourceQuota_To_v1_AppliedClusterResourceQuota(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}
func Convert_quota_AppliedClusterResourceQuotaList_To_v1_AppliedClusterResourceQuotaList(in *quota.AppliedClusterResourceQuotaList, out *v1.AppliedClusterResourceQuotaList, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_quota_AppliedClusterResourceQuotaList_To_v1_AppliedClusterResourceQuotaList(in, out, s)
}
func autoConvert_v1_ClusterResourceQuota_To_quota_ClusterResourceQuota(in *v1.ClusterResourceQuota, out *quota.ClusterResourceQuota, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1_ClusterResourceQuotaSpec_To_quota_ClusterResourceQuotaSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1_ClusterResourceQuotaStatus_To_quota_ClusterResourceQuotaStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}
func Convert_v1_ClusterResourceQuota_To_quota_ClusterResourceQuota(in *v1.ClusterResourceQuota, out *quota.ClusterResourceQuota, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_ClusterResourceQuota_To_quota_ClusterResourceQuota(in, out, s)
}
func autoConvert_quota_ClusterResourceQuota_To_v1_ClusterResourceQuota(in *quota.ClusterResourceQuota, out *v1.ClusterResourceQuota, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_quota_ClusterResourceQuotaSpec_To_v1_ClusterResourceQuotaSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_quota_ClusterResourceQuotaStatus_To_v1_ClusterResourceQuotaStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}
func Convert_quota_ClusterResourceQuota_To_v1_ClusterResourceQuota(in *quota.ClusterResourceQuota, out *v1.ClusterResourceQuota, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_quota_ClusterResourceQuota_To_v1_ClusterResourceQuota(in, out, s)
}
func autoConvert_v1_ClusterResourceQuotaList_To_quota_ClusterResourceQuotaList(in *v1.ClusterResourceQuotaList, out *quota.ClusterResourceQuotaList, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]quota.ClusterResourceQuota, len(*in))
		for i := range *in {
			if err := Convert_v1_ClusterResourceQuota_To_quota_ClusterResourceQuota(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}
func Convert_v1_ClusterResourceQuotaList_To_quota_ClusterResourceQuotaList(in *v1.ClusterResourceQuotaList, out *quota.ClusterResourceQuotaList, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_ClusterResourceQuotaList_To_quota_ClusterResourceQuotaList(in, out, s)
}
func autoConvert_quota_ClusterResourceQuotaList_To_v1_ClusterResourceQuotaList(in *quota.ClusterResourceQuotaList, out *v1.ClusterResourceQuotaList, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]v1.ClusterResourceQuota, len(*in))
		for i := range *in {
			if err := Convert_quota_ClusterResourceQuota_To_v1_ClusterResourceQuota(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}
func Convert_quota_ClusterResourceQuotaList_To_v1_ClusterResourceQuotaList(in *quota.ClusterResourceQuotaList, out *v1.ClusterResourceQuotaList, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_quota_ClusterResourceQuotaList_To_v1_ClusterResourceQuotaList(in, out, s)
}
func autoConvert_v1_ClusterResourceQuotaSelector_To_quota_ClusterResourceQuotaSelector(in *v1.ClusterResourceQuotaSelector, out *quota.ClusterResourceQuotaSelector, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.LabelSelector = (*metav1.LabelSelector)(unsafe.Pointer(in.LabelSelector))
	out.AnnotationSelector = *(*map[string]string)(unsafe.Pointer(&in.AnnotationSelector))
	return nil
}
func Convert_v1_ClusterResourceQuotaSelector_To_quota_ClusterResourceQuotaSelector(in *v1.ClusterResourceQuotaSelector, out *quota.ClusterResourceQuotaSelector, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_ClusterResourceQuotaSelector_To_quota_ClusterResourceQuotaSelector(in, out, s)
}
func autoConvert_quota_ClusterResourceQuotaSelector_To_v1_ClusterResourceQuotaSelector(in *quota.ClusterResourceQuotaSelector, out *v1.ClusterResourceQuotaSelector, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.LabelSelector = (*metav1.LabelSelector)(unsafe.Pointer(in.LabelSelector))
	out.AnnotationSelector = *(*map[string]string)(unsafe.Pointer(&in.AnnotationSelector))
	return nil
}
func Convert_quota_ClusterResourceQuotaSelector_To_v1_ClusterResourceQuotaSelector(in *quota.ClusterResourceQuotaSelector, out *v1.ClusterResourceQuotaSelector, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_quota_ClusterResourceQuotaSelector_To_v1_ClusterResourceQuotaSelector(in, out, s)
}
func autoConvert_v1_ClusterResourceQuotaSpec_To_quota_ClusterResourceQuotaSpec(in *v1.ClusterResourceQuotaSpec, out *quota.ClusterResourceQuotaSpec, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := Convert_v1_ClusterResourceQuotaSelector_To_quota_ClusterResourceQuotaSelector(&in.Selector, &out.Selector, s); err != nil {
		return err
	}
	if err := corev1.Convert_v1_ResourceQuotaSpec_To_core_ResourceQuotaSpec(&in.Quota, &out.Quota, s); err != nil {
		return err
	}
	return nil
}
func Convert_v1_ClusterResourceQuotaSpec_To_quota_ClusterResourceQuotaSpec(in *v1.ClusterResourceQuotaSpec, out *quota.ClusterResourceQuotaSpec, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_ClusterResourceQuotaSpec_To_quota_ClusterResourceQuotaSpec(in, out, s)
}
func autoConvert_quota_ClusterResourceQuotaSpec_To_v1_ClusterResourceQuotaSpec(in *quota.ClusterResourceQuotaSpec, out *v1.ClusterResourceQuotaSpec, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := Convert_quota_ClusterResourceQuotaSelector_To_v1_ClusterResourceQuotaSelector(&in.Selector, &out.Selector, s); err != nil {
		return err
	}
	if err := corev1.Convert_core_ResourceQuotaSpec_To_v1_ResourceQuotaSpec(&in.Quota, &out.Quota, s); err != nil {
		return err
	}
	return nil
}
func Convert_quota_ClusterResourceQuotaSpec_To_v1_ClusterResourceQuotaSpec(in *quota.ClusterResourceQuotaSpec, out *v1.ClusterResourceQuotaSpec, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_quota_ClusterResourceQuotaSpec_To_v1_ClusterResourceQuotaSpec(in, out, s)
}
func autoConvert_v1_ClusterResourceQuotaStatus_To_quota_ClusterResourceQuotaStatus(in *v1.ClusterResourceQuotaStatus, out *quota.ClusterResourceQuotaStatus, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := corev1.Convert_v1_ResourceQuotaStatus_To_core_ResourceQuotaStatus(&in.Total, &out.Total, s); err != nil {
		return err
	}
	if err := Convert_v1_ResourceQuotasStatusByNamespace_To_quota_ResourceQuotasStatusByNamespace(&in.Namespaces, &out.Namespaces, s); err != nil {
		return err
	}
	return nil
}
func Convert_v1_ClusterResourceQuotaStatus_To_quota_ClusterResourceQuotaStatus(in *v1.ClusterResourceQuotaStatus, out *quota.ClusterResourceQuotaStatus, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_ClusterResourceQuotaStatus_To_quota_ClusterResourceQuotaStatus(in, out, s)
}
func autoConvert_quota_ClusterResourceQuotaStatus_To_v1_ClusterResourceQuotaStatus(in *quota.ClusterResourceQuotaStatus, out *v1.ClusterResourceQuotaStatus, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := corev1.Convert_core_ResourceQuotaStatus_To_v1_ResourceQuotaStatus(&in.Total, &out.Total, s); err != nil {
		return err
	}
	if err := Convert_quota_ResourceQuotasStatusByNamespace_To_v1_ResourceQuotasStatusByNamespace(&in.Namespaces, &out.Namespaces, s); err != nil {
		return err
	}
	return nil
}
func Convert_quota_ClusterResourceQuotaStatus_To_v1_ClusterResourceQuotaStatus(in *quota.ClusterResourceQuotaStatus, out *v1.ClusterResourceQuotaStatus, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_quota_ClusterResourceQuotaStatus_To_v1_ClusterResourceQuotaStatus(in, out, s)
}
