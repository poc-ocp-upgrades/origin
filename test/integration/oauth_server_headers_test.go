package integration

import (
	"net/http"
	"net/url"
	"reflect"
	"testing"
	"k8s.io/apimachinery/pkg/util/diff"
	"k8s.io/apimachinery/pkg/util/sets"
	restclient "k8s.io/client-go/rest"
	testutil "github.com/openshift/origin/test/util"
	testserver "github.com/openshift/origin/test/util/server"
)

func TestOAuthServerHeaders(t *testing.T) {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	masterOptions, err := testserver.DefaultMasterOptions()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	defer testserver.CleanupMasterEtcd(t, masterOptions)
	clusterAdminKubeConfig, err := testserver.StartConfiguredMaster(masterOptions)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	clientConfig, err := testutil.GetClusterAdminClientConfig(clusterAdminKubeConfig)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	anonConfig := restclient.AnonymousClientConfig(clientConfig)
	transport, err := restclient.TransportFor(anonConfig)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	baseURL, err := url.Parse(clientConfig.Host)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	for _, path := range []string{"/login", "/login/kube:admin", "/login/anypassword", "/logout", "/oauth/token", "/oauth/authorize", "/oauth/authorize/approve", "/oauth/token/request"} {
		t.Run(path, func(t *testing.T) {
			urlCopy := *baseURL
			urlCopy.Path = path
			checkNewReqHeaders(t, transport, urlCopy.String())
		})
	}
}
func checkNewReqHeaders(t *testing.T, rt http.RoundTripper, checkUrl string) {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	req, err := http.NewRequest("GET", checkUrl, nil)
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}
	req.Header.Set("Accept", "text/html; charset=utf-8")
	resp, err := rt.RoundTrip(req)
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}
	allHeaders := http.Header{}
	for key, val := range map[string]string{"Cache-Control": "no-cache, no-store, max-age=0, must-revalidate", "Pragma": "no-cache", "Expires": "0", "Referrer-Policy": "strict-origin-when-cross-origin", "X-Frame-Options": "DENY", "X-Content-Type-Options": "nosniff", "X-DNS-Prefetch-Control": "off", "X-XSS-Protection": "1; mode=block"} {
		allHeaders.Set(key, val)
	}
	ignoredHeaders := []string{"Audit-Id", "Date", "Content-Type", "Content-Length", "Location"}
	for _, h := range ignoredHeaders {
		resp.Header.Del(h)
	}
	expires := resp.Header["Expires"]
	if len(expires) == 2 && expires[1] == "Fri, 01 Jan 1990 00:00:00 GMT" {
		resp.Header["Expires"] = expires[:1]
	}
	for k, vv := range resp.Header {
		resp.Header[k] = sets.NewString(vv...).List()
	}
	if !reflect.DeepEqual(allHeaders, resp.Header) {
		t.Errorf("Header for %s does not match: expected: %#v got: %#v diff: %s", checkUrl, allHeaders, resp.Header, diff.ObjectDiff(allHeaders, resp.Header))
	}
}
