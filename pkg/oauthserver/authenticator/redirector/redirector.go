package redirector

import (
	"net/http"
	"bytes"
	"runtime"
	"fmt"
	"net/url"
	"strings"
	"github.com/openshift/origin/pkg/oauthserver/authenticator/tokens"
	oauthhandlers "github.com/openshift/origin/pkg/oauthserver/oauth/handlers"
)

func NewRedirector(baseRequestURL *url.URL, redirectURL string) oauthhandlers.AuthenticationRedirector {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &redirector{BaseRequestURL: baseRequestURL, RedirectURL: redirectURL}
}
func NewChallenger(baseRequestURL *url.URL, redirectURL string) oauthhandlers.AuthenticationChallenger {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &redirector{BaseRequestURL: baseRequestURL, RedirectURL: redirectURL}
}

type redirector struct {
	BaseRequestURL	*url.URL
	RedirectURL	string
}

func (r *redirector) AuthenticationChallenge(req *http.Request) (http.Header, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	redirectURL, err := buildRedirectURL(r.RedirectURL, r.BaseRequestURL, req.URL)
	if err != nil {
		return nil, err
	}
	headers := http.Header{}
	headers.Add("Location", redirectURL.String())
	return headers, nil
}
func (r *redirector) AuthenticationRedirect(w http.ResponseWriter, req *http.Request) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	redirectURL, err := buildRedirectURL(r.RedirectURL, r.BaseRequestURL, req.URL)
	if err != nil {
		return nil
	}
	http.Redirect(w, req, redirectURL.String(), http.StatusFound)
	return nil
}
func buildRedirectURL(redirectTemplate string, baseRequestURL, requestURL *url.URL) (*url.URL, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if baseRequestURL != nil {
		requestURL = baseRequestURL.ResolveReference(requestURL)
	}
	redirectURL, err := url.Parse(redirectTemplate)
	if err != nil {
		return nil, err
	}
	serverRelativeRequestURL := &url.URL{Path: requestURL.Path, RawQuery: requestURL.RawQuery}
	redirectURL.RawQuery = strings.Replace(redirectURL.RawQuery, tokens.QueryToken, requestURL.RawQuery, -1)
	redirectURL.RawQuery = strings.Replace(redirectURL.RawQuery, tokens.URLToken, url.QueryEscape(requestURL.String()), -1)
	redirectURL.RawQuery = strings.Replace(redirectURL.RawQuery, tokens.ServerRelativeURLToken, url.QueryEscape(serverRelativeRequestURL.String()), -1)
	return redirectURL, nil
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := runtime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", runtime.FuncForPC(pc).Name()))
	http.Post("/"+"logcode", "application/json", bytes.NewBuffer(jsonLog))
}
