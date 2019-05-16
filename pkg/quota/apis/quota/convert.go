package quota

import (
	goformat "fmt"
	quotav1 "github.com/openshift/api/quota/v1"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

func ConvertAppliedClusterResourceQuotaToClusterResourceQuota(in *AppliedClusterResourceQuota) *ClusterResourceQuota {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &ClusterResourceQuota{ObjectMeta: in.ObjectMeta, Spec: in.Spec, Status: in.Status}
}
func ConvertV1ClusterResourceQuotaToV1AppliedClusterResourceQuota(in *quotav1.ClusterResourceQuota) *quotav1.AppliedClusterResourceQuota {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &quotav1.AppliedClusterResourceQuota{ObjectMeta: in.ObjectMeta, Spec: in.Spec, Status: in.Status}
}
func ConvertV1AppliedClusterResourceQuotaToV1ClusterResourceQuota(in *quotav1.AppliedClusterResourceQuota) *quotav1.ClusterResourceQuota {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &quotav1.ClusterResourceQuota{ObjectMeta: in.ObjectMeta, Spec: in.Spec, Status: in.Status}
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
