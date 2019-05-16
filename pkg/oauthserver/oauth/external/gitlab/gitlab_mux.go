package gitlab

import (
	goformat "fmt"
	"github.com/openshift/origin/pkg/oauthserver/oauth/external"
	"k8s.io/klog"
	"net/http"
	"net/url"
	goos "os"
	godefaultruntime "runtime"
	"strings"
	gotime "time"
)

const gitlabHostedDomain = "gitlab.com"

func NewProvider(providerName, URL, clientID, clientSecret string, transport http.RoundTripper, legacy *bool) (external.Provider, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if isLegacy(legacy, URL) {
		klog.Infof("Using legacy OAuth2 for GitLab identity provider %s url=%s clientID=%s", providerName, URL, clientID)
		return NewOAuthProvider(providerName, URL, clientID, clientSecret, transport)
	}
	klog.Infof("Using OIDC for GitLab identity provider %s url=%s clientID=%s", providerName, URL, clientID)
	return NewOIDCProvider(providerName, URL, clientID, clientSecret, transport)
}
func isLegacy(legacy *bool, URL string) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if legacy != nil {
		return *legacy
	}
	if u, err := url.Parse(URL); err == nil && strings.EqualFold(u.Hostname(), gitlabHostedDomain) {
		return false
	}
	return true
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
