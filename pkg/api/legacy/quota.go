package legacy

import (
	quotav1 "github.com/openshift/api/quota/v1"
	"github.com/openshift/origin/pkg/quota/apis/quota"
	quotav1helpers "github.com/openshift/origin/pkg/quota/apis/quota/v1"
	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
)

func InstallInternalLegacyQuota(scheme *runtime.Scheme) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	InstallExternalLegacyQuota(scheme)
	schemeBuilder := runtime.NewSchemeBuilder(addUngroupifiedInternalQuotaTypes, quotav1helpers.RegisterDefaults, quotav1helpers.RegisterConversions)
	utilruntime.Must(schemeBuilder.AddToScheme(scheme))
}
func InstallExternalLegacyQuota(scheme *runtime.Scheme) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	schemeBuilder := runtime.NewSchemeBuilder(addUngroupifiedQuotaTypes)
	utilruntime.Must(schemeBuilder.AddToScheme(scheme))
}
func addUngroupifiedQuotaTypes(scheme *runtime.Scheme) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	types := []runtime.Object{&quotav1.ClusterResourceQuota{}, &quotav1.ClusterResourceQuotaList{}, &quotav1.AppliedClusterResourceQuota{}, &quotav1.AppliedClusterResourceQuotaList{}}
	scheme.AddKnownTypes(GroupVersion, types...)
	return nil
}
func addUngroupifiedInternalQuotaTypes(scheme *runtime.Scheme) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	scheme.AddKnownTypes(InternalGroupVersion, &quota.ClusterResourceQuota{}, &quota.ClusterResourceQuotaList{}, &quota.AppliedClusterResourceQuota{}, &quota.AppliedClusterResourceQuotaList{})
	return nil
}
