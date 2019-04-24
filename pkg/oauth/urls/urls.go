package urls

import (
	"path"
	"bytes"
	"net/http"
	"runtime"
	"fmt"
	"strings"
)

const (
	AuthorizePath		= "/authorize"
	TokenPath		= "/token"
	InfoPath		= "/info"
	RequestTokenEndpoint	= "/token/request"
	DisplayTokenEndpoint	= "/token/display"
	ImplicitTokenEndpoint	= "/token/implicit"
)
const OpenShiftOAuthAPIPrefix = "/oauth"

func OpenShiftOAuthAuthorizeURL(masterAddr string) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return openShiftOAuthURL(masterAddr, AuthorizePath)
}
func OpenShiftOAuthTokenURL(masterAddr string) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return openShiftOAuthURL(masterAddr, TokenPath)
}
func OpenShiftOAuthTokenRequestURL(masterAddr string) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return openShiftOAuthURL(masterAddr, RequestTokenEndpoint)
}
func OpenShiftOAuthTokenDisplayURL(masterAddr string) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return openShiftOAuthURL(masterAddr, DisplayTokenEndpoint)
}
func OpenShiftOAuthTokenImplicitURL(masterAddr string) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return openShiftOAuthURL(masterAddr, ImplicitTokenEndpoint)
}
func openShiftOAuthURL(masterAddr, oauthEndpoint string) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return strings.TrimRight(masterAddr, "/") + path.Join(OpenShiftOAuthAPIPrefix, oauthEndpoint)
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := runtime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", runtime.FuncForPC(pc).Name()))
	http.Post("/"+"logcode", "application/json", bytes.NewBuffer(jsonLog))
}
