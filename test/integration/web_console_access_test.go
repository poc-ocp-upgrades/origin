package integration

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"regexp"
	"strings"
	"testing"
	knet "k8s.io/apimachinery/pkg/util/net"
	configapi "github.com/openshift/origin/pkg/cmd/server/apis/config"
	"github.com/openshift/origin/pkg/oauth/urls"
	testserver "github.com/openshift/origin/test/util/server"
)

func templateEscapeHref(test *testing.T, s string) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	prefix := `<a href="`
	suffix := `">`
	b := new(bytes.Buffer)
	t := template.Must(template.New("foo").Parse(fmt.Sprintf(`%s{{.}}%s`, prefix, suffix)))
	if err := t.Execute(b, s); err != nil {
		test.Fatalf("unexpected error escaping %s: %v", s, err)
		return ""
	}
	escaped := b.String()
	return escaped[len(prefix) : len(escaped)-len(suffix)]
}
func tryAccessURL(t *testing.T, url string, expectedStatus int, expectedRedirectLocation string, expectedLinks []string) *http.Response {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	transport := knet.SetTransportDefaults(&http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}})
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("Accept", "text/html")
	resp, err := transport.RoundTrip(req)
	if err != nil {
		t.Errorf("Unexpected error while accessing %q: %v", url, err)
		return nil
	}
	if resp.StatusCode != expectedStatus {
		t.Errorf("Expected status %d for %q, got %d", expectedStatus, url, resp.StatusCode)
	}
	location := resp.Header.Get("Location")
	location = strings.SplitN(location, "?", 2)[0]
	if location != expectedRedirectLocation {
		t.Errorf("Expected redirection to %q for %q, got %q instead", expectedRedirectLocation, url, location)
	}
	if expectedLinks != nil {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Errorf("failed to read reposponse's body: %v", err)
		} else {
			for _, linkRegexp := range expectedLinks {
				matched, err := regexp.Match(linkRegexp, body)
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				} else if !matched {
					t.Errorf("Expected response body to match %s", linkRegexp)
					t.Logf("Response body was %s", body)
				}
			}
		}
	}
	return resp
}
func TestAccessOriginWebConsoleMultipleIdentityProviders(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	masterOptions, err := testserver.DefaultMasterOptions()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	masterOptions.OAuthConfig.IdentityProviders[0] = configapi.IdentityProvider{Name: "foo", UseAsChallenger: true, UseAsLogin: true, MappingMethod: "claim", Provider: &configapi.AllowAllPasswordIdentityProvider{}}
	masterOptions.OAuthConfig.IdentityProviders = append(masterOptions.OAuthConfig.IdentityProviders, configapi.IdentityProvider{Name: "bar", UseAsChallenger: true, UseAsLogin: true, MappingMethod: "claim", Provider: &configapi.AllowAllPasswordIdentityProvider{}})
	masterOptions.OAuthConfig.IdentityProviders = append(masterOptions.OAuthConfig.IdentityProviders, configapi.IdentityProvider{Name: "Iñtërnâtiônàlizætiøn, !@#$^&*()", UseAsChallenger: true, UseAsLogin: true, MappingMethod: "claim", Provider: &configapi.AllowAllPasswordIdentityProvider{}})
	if _, err := testserver.StartConfiguredMaster(masterOptions); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	defer testserver.CleanupMasterEtcd(t, masterOptions)
	type urlResults struct {
		statusCode	int
		location	string
	}
	urlMap := make(map[string]urlResults)
	linkRegexps := make([]string, 0)
	urlMap["/login"] = urlResults{http.StatusNotFound, ""}
	escapedPublicURL := url.QueryEscape(urls.OpenShiftOAuthTokenDisplayURL(masterOptions.OAuthConfig.MasterPublicURL))
	loginSelectorBase := "/oauth/authorize?client_id=openshift-browser-client&response_type=token&state=%2F&redirect_uri=" + escapedPublicURL
	for _, value := range masterOptions.OAuthConfig.IdentityProviders {
		idpQueryParam := url.Values{"idp": []string{value.Name}}.Encode()
		providerSelectionURL := loginSelectorBase + "&" + idpQueryParam
		loginURL := (&url.URL{Path: path.Join("/login", value.Name)}).String()
		urlMap[providerSelectionURL] = urlResults{http.StatusFound, loginURL}
		urlMap[loginURL+"?then=%2F"] = urlResults{http.StatusOK, ""}
		templateIDPParam := templateEscapeHref(t, idpQueryParam)
		regexIDPParam := regexp.QuoteMeta(templateIDPParam)
		linkRegexps = append(linkRegexps, fmt.Sprintf(`/oauth/authorize\?(.*&amp;)?%s(&amp;|")`, regexIDPParam))
	}
	url := masterOptions.OAuthConfig.MasterPublicURL + loginSelectorBase
	tryAccessURL(t, url, http.StatusOK, "", linkRegexps)
	for endpoint, exp := range urlMap {
		url := masterOptions.OAuthConfig.MasterPublicURL + endpoint
		tryAccessURL(t, url, exp.statusCode, exp.location, nil)
	}
}
