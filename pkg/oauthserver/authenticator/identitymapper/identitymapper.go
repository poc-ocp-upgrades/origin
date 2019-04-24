package identitymapper

import (
	"fmt"
	"bytes"
	"net/http"
	"runtime"
	"k8s.io/klog"
	"k8s.io/apiserver/pkg/authentication/authenticator"
	"github.com/openshift/origin/pkg/oauthserver/api"
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := runtime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", runtime.FuncForPC(pc).Name()))
	http.Post("/"+"logcode", "application/json", bytes.NewBuffer(jsonLog))
}
