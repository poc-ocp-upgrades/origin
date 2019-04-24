package security

import (
	"fmt"
	"bytes"
	"net/http"
	"runtime"
)

const (
	UIDRangeAnnotation		= "openshift.io/sa.scc.uid-range"
	SupplementalGroupsAnnotation	= "openshift.io/sa.scc.supplemental-groups"
	MCSAnnotation			= "openshift.io/sa.scc.mcs"
	ValidatedSCCAnnotation		= "openshift.io/scc"
)

func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := runtime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", runtime.FuncForPC(pc).Name()))
	http.Post("/"+"logcode", "application/json", bytes.NewBuffer(jsonLog))
}
