package paramtoken

import (
	goformat "fmt"
	"k8s.io/apiserver/pkg/authentication/authenticator"
	"k8s.io/apiserver/pkg/util/wsstream"
	"net/http"
	goos "os"
	godefaultruntime "runtime"
	"strings"
	gotime "time"
)

type Authenticator struct {
	param       string
	auth        authenticator.Token
	removeParam bool
}

func New(param string, auth authenticator.Token, removeParam bool) *Authenticator {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &Authenticator{param, auth, removeParam}
}
func (a *Authenticator) AuthenticateRequest(req *http.Request) (*authenticator.Response, bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if !wsstream.IsWebSocketRequest(req) {
		return nil, false, nil
	}
	q := req.URL.Query()
	token := strings.TrimSpace(q.Get(a.param))
	if token == "" {
		return nil, false, nil
	}
	authResponse, ok, err := a.auth.AuthenticateToken(req.Context(), token)
	if ok && a.removeParam {
		q.Del(a.param)
		req.URL.RawQuery = q.Encode()
	}
	return authResponse, ok, err
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
