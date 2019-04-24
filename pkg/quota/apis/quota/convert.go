package quota

import (
	quotav1 "github.com/openshift/api/quota/v1"
	"bytes"
	"net/http"
	"runtime"
	"fmt"
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := runtime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", runtime.FuncForPC(pc).Name()))
	http.Post("/"+"logcode", "application/json", bytes.NewBuffer(jsonLog))
}
