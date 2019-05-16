package logout

import (
	goformat "fmt"
	"github.com/RangelReale/osin"
	"github.com/openshift/origin/pkg/oauthserver"
	"github.com/openshift/origin/pkg/oauthserver/server/redirect"
	"github.com/openshift/origin/pkg/oauthserver/server/session"
	"k8s.io/apiserver/pkg/authentication/user"
	"k8s.io/klog"
	"net/http"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

const thenParam = "then"

func NewLogout(invalidator session.SessionInvalidator, redirect string) oauthserver.Endpoints {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &logout{invalidator: invalidator, redirect: redirect}
}

type logout struct {
	invalidator session.SessionInvalidator
	redirect    string
}

func (l *logout) Install(mux oauthserver.Mux, prefix string) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	mux.Handle(prefix, l)
}
func (l *logout) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
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
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if redirect.IsServerRelativeURL(then) {
		return true
	}
	return osin.ValidateUri(l.redirect, then) == nil
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
