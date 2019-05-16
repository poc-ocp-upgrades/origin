package identitymapper

import (
	"fmt"
	goformat "fmt"
	"github.com/openshift/origin/pkg/oauthserver/api"
	"k8s.io/apiserver/pkg/authentication/authenticator"
	"k8s.io/klog"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

func ResponseFor(mapper api.UserIdentityMapper, identity api.UserIdentityInfo) (*authenticator.Response, bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	user, err := mapper.UserFor(identity)
	if err != nil {
		logf("error creating or updating mapping for: %#v due to %v", identity, err)
		return nil, false, err
	}
	logf("got userIdentityMapping: %#v", user)
	return &authenticator.Response{User: user}, true, nil
}
func logf(format string, args ...interface{}) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if klog.V(4) {
		klog.InfoDepth(2, fmt.Sprintf("identitymapper: "+format, args...))
	}
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
