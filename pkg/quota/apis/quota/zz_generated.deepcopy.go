package quota

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

func (in *AppliedClusterResourceQuota) DeepCopyInto(out *AppliedClusterResourceQuota) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}
func (in *AppliedClusterResourceQuota) DeepCopy() *AppliedClusterResourceQuota {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(AppliedClusterResourceQuota)
	in.DeepCopyInto(out)
	return out
}
func (in *AppliedClusterResourceQuota) DeepCopyObject() runtime.Object {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
func (in *AppliedClusterResourceQuotaList) DeepCopyInto(out *AppliedClusterResourceQuotaList) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AppliedClusterResourceQuota, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}
func (in *AppliedClusterResourceQuotaList) DeepCopy() *AppliedClusterResourceQuotaList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(AppliedClusterResourceQuotaList)
	in.DeepCopyInto(out)
	return out
}
func (in *AppliedClusterResourceQuotaList) DeepCopyObject() runtime.Object {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
func (in *ClusterResourceQuota) DeepCopyInto(out *ClusterResourceQuota) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}
func (in *ClusterResourceQuota) DeepCopy() *ClusterResourceQuota {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(ClusterResourceQuota)
	in.DeepCopyInto(out)
	return out
}
func (in *ClusterResourceQuota) DeepCopyObject() runtime.Object {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
func (in *ClusterResourceQuotaList) DeepCopyInto(out *ClusterResourceQuotaList) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterResourceQuota, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}
func (in *ClusterResourceQuotaList) DeepCopy() *ClusterResourceQuotaList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(ClusterResourceQuotaList)
	in.DeepCopyInto(out)
	return out
}
func (in *ClusterResourceQuotaList) DeepCopyObject() runtime.Object {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
func (in *ClusterResourceQuotaSelector) DeepCopyInto(out *ClusterResourceQuotaSelector) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	if in.LabelSelector != nil {
		in, out := &in.LabelSelector, &out.LabelSelector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.AnnotationSelector != nil {
		in, out := &in.AnnotationSelector, &out.AnnotationSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}
func (in *ClusterResourceQuotaSelector) DeepCopy() *ClusterResourceQuotaSelector {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(ClusterResourceQuotaSelector)
	in.DeepCopyInto(out)
	return out
}
func (in *ClusterResourceQuotaSpec) DeepCopyInto(out *ClusterResourceQuotaSpec) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	in.Selector.DeepCopyInto(&out.Selector)
	in.Quota.DeepCopyInto(&out.Quota)
	return
}
func (in *ClusterResourceQuotaSpec) DeepCopy() *ClusterResourceQuotaSpec {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(ClusterResourceQuotaSpec)
	in.DeepCopyInto(out)
	return out
}
func (in *ClusterResourceQuotaStatus) DeepCopyInto(out *ClusterResourceQuotaStatus) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	in.Total.DeepCopyInto(&out.Total)
	out.Namespaces = in.Namespaces.DeepCopy()
	return
}
func (in *ClusterResourceQuotaStatus) DeepCopy() *ClusterResourceQuotaStatus {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(ClusterResourceQuotaStatus)
	in.DeepCopyInto(out)
	return out
}
func (in *ResourceQuotasStatusByNamespace) DeepCopyInto(out *ResourceQuotasStatusByNamespace) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = in.DeepCopy()
	return
}
