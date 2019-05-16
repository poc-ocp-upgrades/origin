package v1

import (
	goformat "fmt"
	"github.com/openshift/api/quota/v1"
	internal "github.com/openshift/origin/pkg/quota/apis/quota"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/conversion"
	"k8s.io/apimachinery/pkg/runtime"
	kapi "k8s.io/kubernetes/pkg/apis/core"
	kapiv1 "k8s.io/kubernetes/pkg/apis/core/v1"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

func Convert_v1_ResourceQuotasStatusByNamespace_To_quota_ResourceQuotasStatusByNamespace(in *v1.ResourceQuotasStatusByNamespace, out *internal.ResourceQuotasStatusByNamespace, s conversion.Scope) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if in == nil {
		return nil
	}
	for _, curr := range *in {
		internalStatus := &kapi.ResourceQuotaStatus{}
		kapiv1.Convert_v1_ResourceQuotaStatus_To_core_ResourceQuotaStatus(&curr.Status, internalStatus, s)
		out.Insert(curr.Namespace, *internalStatus)
	}
	return nil
}
func Convert_quota_ResourceQuotasStatusByNamespace_To_v1_ResourceQuotasStatusByNamespace(in *internal.ResourceQuotasStatusByNamespace, out *v1.ResourceQuotasStatusByNamespace, s conversion.Scope) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	for e := in.OrderedKeys().Front(); e != nil; e = e.Next() {
		namespace := e.Value.(string)
		status, _ := in.Get(namespace)
		versionedStatus := &corev1.ResourceQuotaStatus{}
		kapiv1.Convert_core_ResourceQuotaStatus_To_v1_ResourceQuotaStatus(&status, versionedStatus, s)
		if out == nil {
			out = &v1.ResourceQuotasStatusByNamespace{}
		}
		*out = append(*out, v1.ResourceQuotaStatusByNamespace{Namespace: namespace, Status: *versionedStatus})
	}
	return nil
}
func AddConversionFuncs(scheme *runtime.Scheme) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return scheme.AddConversionFuncs(Convert_quota_ResourceQuotasStatusByNamespace_To_v1_ResourceQuotasStatusByNamespace, Convert_v1_ResourceQuotasStatusByNamespace_To_quota_ResourceQuotasStatusByNamespace)
}
func ConvertV1ClusterResourceQuotaToInternalAppliedClusterResourceQuota(in *v1.ClusterResourceQuota) (*internal.AppliedClusterResourceQuota, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	appliedExternal := internal.ConvertV1ClusterResourceQuotaToV1AppliedClusterResourceQuota(in)
	appliedInternal := &internal.AppliedClusterResourceQuota{}
	err := Convert_v1_AppliedClusterResourceQuota_To_quota_AppliedClusterResourceQuota(appliedExternal, appliedInternal, nil)
	return appliedInternal, err
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
