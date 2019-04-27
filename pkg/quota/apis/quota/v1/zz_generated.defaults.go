package v1

import (
	v1 "github.com/openshift/api/quota/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	corev1 "k8s.io/kubernetes/pkg/apis/core/v1"
)

func RegisterDefaults(scheme *runtime.Scheme) error {
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
	scheme.AddTypeDefaultingFunc(&v1.AppliedClusterResourceQuota{}, func(obj interface{}) {
		SetObjectDefaults_AppliedClusterResourceQuota(obj.(*v1.AppliedClusterResourceQuota))
	})
	scheme.AddTypeDefaultingFunc(&v1.AppliedClusterResourceQuotaList{}, func(obj interface{}) {
		SetObjectDefaults_AppliedClusterResourceQuotaList(obj.(*v1.AppliedClusterResourceQuotaList))
	})
	scheme.AddTypeDefaultingFunc(&v1.ClusterResourceQuota{}, func(obj interface{}) {
		SetObjectDefaults_ClusterResourceQuota(obj.(*v1.ClusterResourceQuota))
	})
	scheme.AddTypeDefaultingFunc(&v1.ClusterResourceQuotaList{}, func(obj interface{}) {
		SetObjectDefaults_ClusterResourceQuotaList(obj.(*v1.ClusterResourceQuotaList))
	})
	return nil
}
func SetObjectDefaults_AppliedClusterResourceQuota(in *v1.AppliedClusterResourceQuota) {
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
	corev1.SetDefaults_ResourceList(&in.Spec.Quota.Hard)
	corev1.SetDefaults_ResourceList(&in.Status.Total.Hard)
	corev1.SetDefaults_ResourceList(&in.Status.Total.Used)
	for i := range in.Status.Namespaces {
		a := &in.Status.Namespaces[i]
		corev1.SetDefaults_ResourceList(&a.Status.Hard)
		corev1.SetDefaults_ResourceList(&a.Status.Used)
	}
}
func SetObjectDefaults_AppliedClusterResourceQuotaList(in *v1.AppliedClusterResourceQuotaList) {
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
	for i := range in.Items {
		a := &in.Items[i]
		SetObjectDefaults_AppliedClusterResourceQuota(a)
	}
}
func SetObjectDefaults_ClusterResourceQuota(in *v1.ClusterResourceQuota) {
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
	corev1.SetDefaults_ResourceList(&in.Spec.Quota.Hard)
	corev1.SetDefaults_ResourceList(&in.Status.Total.Hard)
	corev1.SetDefaults_ResourceList(&in.Status.Total.Used)
	for i := range in.Status.Namespaces {
		a := &in.Status.Namespaces[i]
		corev1.SetDefaults_ResourceList(&a.Status.Hard)
		corev1.SetDefaults_ResourceList(&a.Status.Used)
	}
}
func SetObjectDefaults_ClusterResourceQuotaList(in *v1.ClusterResourceQuotaList) {
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
	for i := range in.Items {
		a := &in.Items[i]
		SetObjectDefaults_ClusterResourceQuota(a)
	}
}
