package openshiftkubeapiserver

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"k8s.io/apimachinery/pkg/util/sets"
	apifilters "k8s.io/apiserver/pkg/endpoints/filters"
	apiserver "k8s.io/apiserver/pkg/server"
	"k8s.io/kubernetes/pkg/api/legacyscheme"
	kubecontrolplanev1 "github.com/openshift/api/kubecontrolplane/v1"
)

var (
	currentOCKubeResources			= "oc/v1.2.0 (linux/amd64) kubernetes/bc4550d"
	currentOCOriginResources		= "oc/v1.1.3 (linux/amd64) openshift/b348c2f"
	currentOpenshiftKubectlKubeResources	= "openshift/v1.2.0 (linux/amd64) kubernetes/bc4550d"
	currentOpenshiftKubectlOriginResources	= "openshift/v1.1.3 (linux/amd64) openshift/b348c2f"
	currentOADMKubeResources		= "oadm/v1.2.0 (linux/amd64) kubernetes/bc4550d"
	currentOADMOriginResources		= "oadm/v1.1.3 (linux/amd64) openshift/b348c2f"
	currentVersionUserAgents		= []string{currentOCKubeResources, currentOCOriginResources, currentOpenshiftKubectlKubeResources, currentOpenshiftKubectlOriginResources, currentOADMKubeResources, currentOADMOriginResources}
	olderOCKubeResources			= "oc/v1.1.10 (linux/amd64) kubernetes/bc4550d"
	olderOCOriginResources			= "oc/v1.1.1 (linux/amd64) openshift/b348c2f"
	oldestOCOriginResources			= "oc/v1.0.1 (linux/amd64) openshift/b348c2f"
	olderOpenshiftKubectlKubeResources	= "openshift/v1.1.10 (linux/amd64) kubernetes/bc4550d"
	olderOpenshiftKubectlOriginResources	= "openshift/v1.1.1 (linux/amd64) openshift/b348c2f"
	olderOADMKubeResources			= "oadm/v1.1.10 (linux/amd64) kubernetes/bc4550d"
	olderOADMOriginResources		= "oadm/v1.1.1 (linux/amd64) openshift/b348c2f"
	olderVersionUserAgents			= []string{olderOCKubeResources, olderOCOriginResources, oldestOCOriginResources, olderOpenshiftKubectlKubeResources, olderOpenshiftKubectlOriginResources, olderOADMKubeResources, olderOADMOriginResources}
	newerOCKubeResources			= "oc/v1.2.1 (linux/amd64) kubernetes/bc4550d"
	newerOCOriginResources			= "oc/v1.1.4 (linux/amd64) openshift/b348c2f"
	newerOpenshiftKubectlKubeResources	= "openshift/v1.2.1 (linux/amd64) kubernetes/bc4550d"
	newerOpenshiftKubectlOriginResources	= "openshift/v1.1.4 (linux/amd64) openshift/b348c2f"
	newerOADMKubeResources			= "oadm/v1.2.1 (linux/amd64) kubernetes/bc4550d"
	newerOADMOriginResources		= "oadm/v1.1.4 (linux/amd64) openshift/b348c2f"
	newerVersionUserAgents			= []string{newerOCKubeResources, newerOCOriginResources, newerOpenshiftKubectlKubeResources, newerOpenshiftKubectlOriginResources, newerOADMKubeResources, newerOADMOriginResources}
	notOCVersion				= "something else"
	openshiftServerVersion			= `v1\.1\.3`
	kubeServerVersion			= `v1\.2\.0`
)

type versionSkewTestCase struct {
	name		string
	userAgents	[]string
	failureMessage	string
	methods		[]string
}

func (tc versionSkewTestCase) Run(url string, t *testing.T) {
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
	for _, method := range tc.methods {
		for _, userAgent := range tc.userAgents {
			req, err := http.NewRequest(method, url, nil)
			if err != nil {
				t.Errorf("%s: unexpected error: %v", tc.name, err)
				return
			}
			req.Header.Add("User-Agent", userAgent)
			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				t.Errorf("%s: unexpected error: %v", tc.name, err)
				return
			}
			if len(tc.failureMessage) == 0 {
				if resp.StatusCode != http.StatusOK {
					t.Errorf("%s: %s: unexpected status: %v", tc.name, userAgent, resp.StatusCode)
					return
				}
			} else {
				if resp.StatusCode != http.StatusForbidden {
					t.Errorf("%s: %s: unexpected status: %v", tc.name, userAgent, resp.StatusCode)
					return
				}
				body, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					t.Errorf("%s: unexpected error: %v", tc.name, err)
					return
				}
				if !strings.Contains(string(body), tc.failureMessage) {
					t.Errorf("%s: expected %v, got %v", tc.name, tc.failureMessage, string(body))
					return
				}
			}
		}
	}
}
func TestVersionSkewFilterDenyOld(t *testing.T) {
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
	verbs := []string{"PATCH", "POST"}
	doNothingHandler := http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
	})
	userAgentMatchingConfig := kubecontrolplanev1.UserAgentMatchingConfig{}
	userAgentMatchingConfig.DeniedClients = []kubecontrolplanev1.UserAgentDenyRule{{UserAgentMatchRule: kubecontrolplanev1.UserAgentMatchRule{Regex: `\w+/v1\.1\.10 \(.+/.+\) kubernetes/\w{7}`, HTTPVerbs: verbs}, RejectionMessage: "rejected for reasons!"}, {UserAgentMatchRule: kubecontrolplanev1.UserAgentMatchRule{Regex: `\w+/v(?:(?:1\.1\.1)|(?:1\.0\.1)) \(.+/.+\) openshift/\w{7}`, HTTPVerbs: verbs}, RejectionMessage: "rejected for reasons!"}}
	handler := versionSkewFilter(doNothingHandler, userAgentMatchingConfig)
	server := httptest.NewServer(testHandlerChain(handler))
	defer server.Close()
	testCases := []versionSkewTestCase{{name: "missing", userAgents: []string{""}, methods: verbs}, {name: "not oc", userAgents: []string{notOCVersion}, methods: verbs}, {name: "older", userAgents: olderVersionUserAgents, failureMessage: "rejected for reasons!", methods: verbs}, {name: "newer", userAgents: newerVersionUserAgents, methods: verbs}, {name: "exact", userAgents: currentVersionUserAgents, methods: verbs}}
	for _, tc := range testCases {
		tc.Run(server.URL+"/api/v1/namespaces", t)
	}
}
func TestVersionSkewFilterDenySkewed(t *testing.T) {
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
	verbs := []string{"PUT", "DELETE"}
	doNothingHandler := http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
	})
	userAgentMatchingConfig := kubecontrolplanev1.UserAgentMatchingConfig{}
	userAgentMatchingConfig.RequiredClients = []kubecontrolplanev1.UserAgentMatchRule{{Regex: `\w+/` + kubeServerVersion + ` \(.+/.+\) kubernetes/\w{7}`, HTTPVerbs: verbs}, {Regex: `\w+/` + openshiftServerVersion + ` \(.+/.+\) openshift/\w{7}`, HTTPVerbs: verbs}}
	userAgentMatchingConfig.DefaultRejectionMessage = "rejected for reasons!"
	handler := versionSkewFilter(doNothingHandler, userAgentMatchingConfig)
	server := httptest.NewServer(testHandlerChain(handler))
	defer server.Close()
	testCases := []versionSkewTestCase{{name: "missing", userAgents: []string{""}, failureMessage: "rejected for reasons!", methods: verbs}, {name: "not oc", userAgents: []string{notOCVersion}, failureMessage: "rejected for reasons!", methods: verbs}, {name: "older", userAgents: olderVersionUserAgents, failureMessage: "rejected for reasons!", methods: verbs}, {name: "newer", userAgents: newerVersionUserAgents, failureMessage: "rejected for reasons!", methods: verbs}, {name: "current", userAgents: currentVersionUserAgents, methods: verbs}}
	for _, tc := range testCases {
		tc.Run(server.URL+"/api/v1/namespaces", t)
	}
}
func TestVersionSkewFilterSkippedOnNonAPIRequest(t *testing.T) {
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
	verbs := []string{"PUT", "DELETE"}
	doNothingHandler := http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
	})
	userAgentMatchingConfig := kubecontrolplanev1.UserAgentMatchingConfig{}
	userAgentMatchingConfig.RequiredClients = []kubecontrolplanev1.UserAgentMatchRule{{Regex: `\w+/` + kubeServerVersion + ` \(.+/.+\) kubernetes/\w{7}`, HTTPVerbs: verbs}, {Regex: `\w+/` + openshiftServerVersion + ` \(.+/.+\) openshift/\w{7}`, HTTPVerbs: verbs}}
	userAgentMatchingConfig.DefaultRejectionMessage = "rejected for reasons!"
	handler := versionSkewFilter(doNothingHandler, userAgentMatchingConfig)
	server := httptest.NewServer(testHandlerChain(handler))
	defer server.Close()
	testCases := []versionSkewTestCase{{name: "missing", userAgents: []string{""}, methods: verbs}, {name: "not oc", userAgents: []string{notOCVersion}, methods: verbs}, {name: "older", userAgents: olderVersionUserAgents, methods: verbs}, {name: "newer", userAgents: newerVersionUserAgents, methods: verbs}, {name: "current", userAgents: currentVersionUserAgents, methods: verbs}}
	for _, tc := range testCases {
		tc.Run(server.URL+"/api/v1", t)
	}
}
func testHandlerChain(handler http.Handler) http.Handler {
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
	kgenericconfig := apiserver.NewConfig(legacyscheme.Codecs)
	kgenericconfig.LegacyAPIGroupPrefixes = sets.NewString(apiserver.DefaultLegacyAPIPrefix)
	handler = apifilters.WithRequestInfo(handler, apiserver.NewRequestInfoResolver(kgenericconfig))
	return handler
}
