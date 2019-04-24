package login

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"
	"github.com/MakeNowJust/heredoc"
	"github.com/openshift/origin/pkg/client/config"
	"github.com/openshift/origin/pkg/oauth/util"
	kapierrs "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	restclient "k8s.io/client-go/rest"
	kclientcmdapi "k8s.io/client-go/tools/clientcmd/api"
)

const (
	oauthMetadataEndpoint = "/.well-known/oauth-authorization-server"
)

func TestNormalizeServerURL(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	testCases := []struct {
		originalServerURL	string
		normalizedServerURL	string
	}{{originalServerURL: "localhost", normalizedServerURL: "https://localhost:443"}, {originalServerURL: "https://localhost", normalizedServerURL: "https://localhost:443"}, {originalServerURL: "localhost:443", normalizedServerURL: "https://localhost:443"}, {originalServerURL: "https://localhost:443", normalizedServerURL: "https://localhost:443"}, {originalServerURL: "http://localhost", normalizedServerURL: "http://localhost:80"}, {originalServerURL: "localhost:8443", normalizedServerURL: "https://localhost:8443"}}
	for _, test := range testCases {
		t.Logf("evaluating test: normalize %s -> %s", test.originalServerURL, test.normalizedServerURL)
		normalized, err := config.NormalizeServerURL(test.originalServerURL)
		if err != nil {
			t.Errorf("unexpected error normalizing %s: %s", test.originalServerURL, err)
		}
		if normalized != test.normalizedServerURL {
			t.Errorf("unexpected server URL normalization result for %s: expected %s, got %s", test.originalServerURL, test.normalizedServerURL, normalized)
		}
	}
}
func TestTLSWithCertificateNotMatchingHostname(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	invalidHostCert := heredoc.Doc(`
		-----BEGIN CERTIFICATE-----
		MIICBjCCAW+gAwIBAgIRALOIWXyeLzunaiVkP2itHAEwDQYJKoZIhvcNAQELBQAw
		EjEQMA4GA1UEChMHQWNtZSBDbzAgFw03MDAxMDEwMDAwMDBaGA8yMDg0MDEyOTE2
		MDAwMFowEjEQMA4GA1UEChMHQWNtZSBDbzCBnzANBgkqhkiG9w0BAQEFAAOBjQAw
		gYkCgYEAuKDlC4aMBbHaXgS+MFud5h3zeE4boSqKgFI6HceySF/a+qg0v+ID6EwQ
		DpJ2W5AdJGEBfixo+tym6q3oKWHJUX0hInkJ6dXIdUbVOeO5dIsGG0fZmRD7DDDx
		snkXrDB/E0JglHNckRbIh/jvznbDfbddIcdgZ7JVIfnNpigtHZECAwEAAaNaMFgw
		DgYDVR0PAQH/BAQDAgKkMBMGA1UdJQQMMAoGCCsGAQUFBwMBMA8GA1UdEwEB/wQF
		MAMBAf8wIAYDVR0RBBkwF4IPaW52YWxpZGhvc3QuY29thwQICAgIMA0GCSqGSIb3
		DQEBCwUAA4GBAAkPU044aFkBl4f/muwSh/oPGinnA4fp8ei0KMnLk+0/CjNb3Waa
		GtuRVIudRTK2M/RzdpUrwfWlVmkezV4BR1K/aOH9a29zqDTkEjnkIbWwe+piAs+w
		VxIxrTqM8rqq8qxeWS54AyF/OaLJgXzDpCFnCb7kY3iyHv6lcmCjluLW
		-----END CERTIFICATE-----`)
	invalidHostKey := heredoc.Doc(`
		-----BEGIN RSA PRIVATE KEY-----
		MIICXQIBAAKBgQC4oOULhowFsdpeBL4wW53mHfN4ThuhKoqAUjodx7JIX9r6qDS/
		4gPoTBAOknZbkB0kYQF+LGj63KbqregpYclRfSEieQnp1ch1RtU547l0iwYbR9mZ
		EPsMMPGyeResMH8TQmCUc1yRFsiH+O/OdsN9t10hx2BnslUh+c2mKC0dkQIDAQAB
		AoGAZ0ZAuNC7NFhHEL5QcJZe3aC1Vv9B/0XfkWXtckkJFejggcNjNk5D50Xc2Xnd
		0NvtITNN9Xj8BA83IyDCM5uqUwDbOLIc6qYgAGWzxZZSDAQg1iOAAZoXmMTNS6Zf
		hQhNUIwB68ELGvbcq7cxQL7L9n4GfISz7PKOOUKTZp0Q8G0CQQD07K7NES340c3I
		QVkCW5/ygNK0GuQ8nTcG5yC8R5SDS47N8YzPp17Pajah8+wawYiemY1fUmD7P/bq
		Cjl2RtIHAkEAwPo1GzJubN7PSYgPir3TxUGtMJoyc3jfdjblXyGJHwTu2YxeRjd2
		YUPVRpu9JvNjZc+GONvTbTZeNWCvy0JNpwJBAKEsi49JCd6eefOZBTDnCKd1nLKG
		q8Ezl/2D5WfhFtsbwrrFhOs1cc++Tnte3/VvfC8aTwz2UfmkyyCSX+P0kMsCQCIL
		glb7/LNEU7mbQXKurq+8OHu8mG36wyGt6aVw2yoXyrOiqfclTcM3HmdIjoRSqBSM
		Ghfp4FECKHiuSBVJ6z0CQQDF37CRpdQRDPnAedhyApLcIxSbYo1oUm7FxBLyVb7V
		HQjFvsOylsSCABXz0FyC7zXQxkEo6CiSahVI/PHz6Zta
		-----END RSA PRIVATE KEY-----`)
	server, err := newTLSServer(invalidHostCert, invalidHostKey)
	if err != nil {
		t.Errorf(err.Error())
	}
	server.StartTLS()
	defer server.Close()
	testCases := map[string]struct {
		serverURL	string
		skipTLSVerify	bool
		expectedErrMsg	*regexp.Regexp
	}{"succeed skipping tls": {serverURL: server.URL, skipTLSVerify: true}, "certificate hostname doesn't match": {serverURL: server.URL, expectedErrMsg: regexp.MustCompile(`The server is using a certificate that does not match its hostname(.*)is valid for 8\.8\.8\.8`)}}
	for name, test := range testCases {
		t.Logf("evaluating test: %s", name)
		options := &LoginOptions{Server: test.serverURL, InsecureTLS: test.skipTLSVerify, StartingKubeConfig: &kclientcmdapi.Config{}}
		if _, err = options.getClientConfig(); err != nil {
			if !test.expectedErrMsg.MatchString(err.Error()) {
				t.Errorf("%s: expected error %q but got %q", name, test.expectedErrMsg, err)
			}
			if test.expectedErrMsg == nil {
				t.Errorf("%s: unexpected error: %v", name, err)
			}
		} else {
			if test.expectedErrMsg != nil {
				t.Errorf("%s: expected error but got nothing", name)
			}
		}
	}
}
func TestTLSWithExpiredCertificate(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	expiredCert := heredoc.Doc(`
		-----BEGIN CERTIFICATE-----
		MIICEjCCAXugAwIBAgIRALf82bYpro/jQS8fP74dG5EwDQYJKoZIhvcNAQELBQAw
		EjEQMA4GA1UEChMHQWNtZSBDbzAeFw03MDAxMDEwMDAwMDBaFw03MDAxMDEwMTAw
		MDBaMBIxEDAOBgNVBAoTB0FjbWUgQ28wgZ8wDQYJKoZIhvcNAQEBBQADgY0AMIGJ
		AoGBAONNgDXBk2Q1i/aJjTwt03KpQ3nQblMS3IX/H9JWw6ta6UublKBOaD/2o5Xt
		FM+Q7XDEnzYw88CK5KHdyejkJo5IBpUjQYJZFzUJ1BC8Lw7yy6dXWYBJboRR1S+1
		JhkMJOtpPecv+4cTaynplYj0WMBjcQthg2RM7tdpyUYpsp2rAgMBAAGjaDBmMA4G
		A1UdDwEB/wQEAwICpDATBgNVHSUEDDAKBggrBgEFBQcDATAPBgNVHRMBAf8EBTAD
		AQH/MC4GA1UdEQQnMCWCC2V4YW1wbGUuY29thwR/AAABhxAAAAAAAAAAAAAAAAAA
		AAABMA0GCSqGSIb3DQEBCwUAA4GBAFpdiiM5YAQQN0H5ZMNuHWGlprjp7qVilO8/
		WFePZRWY2vQF8g7/c1cX4bPqG+qFJd+9j2UZNjhadNfMCxvu6BY7NCupOHVHmnRQ
		ocvkPoSqobE7qDPfiUuU1J+61Libu6b2IjV3/K9pvZkLiBrqn0YhoXXa0PG+rG1L
		9X7+mb5z
		-----END CERTIFICATE-----`)
	expiredKey := heredoc.Doc(`
		-----BEGIN RSA PRIVATE KEY-----
		MIICXQIBAAKBgQDjTYA1wZNkNYv2iY08LdNyqUN50G5TEtyF/x/SVsOrWulLm5Sg
		Tmg/9qOV7RTPkO1wxJ82MPPAiuSh3cno5CaOSAaVI0GCWRc1CdQQvC8O8sunV1mA
		SW6EUdUvtSYZDCTraT3nL/uHE2sp6ZWI9FjAY3ELYYNkTO7XaclGKbKdqwIDAQAB
		AoGBAJPFWKqZ9CZboWhfuE/9Qs/yNonE9VRQmMkMOTXXblHCQpUCyjcFgkTDJUpc
		3QCsKZD8Yr0qSe1M3qJUu+UKHf18LqwiL/ynnalYggxIFS5/SidWCngKvIuEfkLK
		VsnCK3jt5qx21iljGHU6bQZHnHB9IGEiBYcnQlvvw/WdvRDBAkEA8/KMpJVwnI1W
		7fzcZ1+mbMeSJoAVIa9u7MgI+LIRZMokDRYeAMvEjm3GYpZDqA5l1dp7KochMep/
		0vSSTHt7ewJBAO6IbcUIDhXuh2qdxR/Xk5DdDCoxaD1o4ivyj9JsSlGa9JWD7kKN
		6ZFFrn8i7uQuniC1Rwc/4yHhs6OqbiF695ECQQCBwVKzvFUwwDEr1yK4zXStSZ3g
		YqJaz4CV63RyK+z6ilaQq2H8FGaRR6yNBdYozre1/0ciAMxUS6H/6Fzk141/AkBe
		SguqIP8AaGObH3Z2mc65KsfOPe2IqNcOrDlx4mCWVXxtRdN+933mcPcDRpnMFSlo
		oH/NO9Ha6M8L2SjjjyohAkBJHU61+OWz/TAy1nxsMbFsISLn/JrdEZIf2uFORlDN
		Z3/XIQ+yeg4Jk1VbTMZ0/fHf9xMFR8acC/7n7jxnzQau
		-----END RSA PRIVATE KEY-----`)
	server, err := newTLSServer(expiredCert, expiredKey)
	if err != nil {
		t.Errorf(err.Error())
	}
	server.StartTLS()
	defer server.Close()
	testCases := map[string]struct {
		serverURL	string
		skipTLSVerify	bool
		expectedErrMsg	*regexp.Regexp
	}{"succeed skipping tls": {serverURL: server.URL, skipTLSVerify: true}, "certificate expired": {serverURL: server.URL, expectedErrMsg: regexp.MustCompile(`The server is using an invalid certificate(.*)has expired`)}}
	for name, test := range testCases {
		t.Logf("evaluating test: %s", name)
		options := &LoginOptions{Server: test.serverURL, InsecureTLS: test.skipTLSVerify, StartingKubeConfig: &kclientcmdapi.Config{}}
		if _, err = options.getClientConfig(); err != nil {
			if !test.expectedErrMsg.MatchString(err.Error()) {
				t.Errorf("%s: expected error %q but got %q", name, test.expectedErrMsg, err)
			}
			if test.expectedErrMsg == nil {
				t.Errorf("%s: unexpected error: %v", name, err)
			}
		} else {
			if test.expectedErrMsg != nil {
				t.Errorf("%s: expected error but got nothing", name)
			}
		}
	}
}
func TestDialToHTTPServer(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	invoked := make(chan struct{}, 1)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		invoked <- struct{}{}
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()
	testCases := map[string]struct {
		serverURL	string
		evalExpectedErr	func(error) bool
	}{"succeed dialing": {serverURL: server.URL}}
	for name, test := range testCases {
		t.Logf("evaluating test: %s", name)
		clientConfig := &restclient.Config{Host: test.serverURL}
		if err := dialToServer(*clientConfig); err != nil {
			if test.evalExpectedErr == nil || !test.evalExpectedErr(err) {
				t.Errorf("%s: unexpected error: %v", name, err)
			}
		} else {
			if test.evalExpectedErr != nil {
				t.Errorf("%s: expected error but got nothing", name)
			}
		}
	}
}

type oauthMetadataResponse struct {
	metadata *util.OauthAuthorizationServerMetadata
}

func (r *oauthMetadataResponse) Serialize() ([]byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	b, err := json.Marshal(r.metadata)
	if err != nil {
		return []byte{}, err
	}
	return b, nil
}
func TestPreserveErrTypeAuthInfo(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	invoked := make(chan struct{}, 3)
	oauthResponse := []byte{}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		select {
		case invoked <- struct{}{}:
			t.Logf("saw %s request for path: %s", r.Method, r.URL.String())
		default:
			t.Fatalf("unexpected request handled by test server: %v: %v", r.Method, r.URL)
		}
		if r.URL.Path == oauthMetadataEndpoint {
			w.WriteHeader(http.StatusOK)
			w.Write(oauthResponse)
			return
		}
		w.WriteHeader(http.StatusUnauthorized)
	}))
	defer server.Close()
	metadataResponse := &oauthMetadataResponse{}
	metadataResponse.metadata = &util.OauthAuthorizationServerMetadata{Issuer: server.URL, AuthorizationEndpoint: server.URL + "/oauth/authorize", TokenEndpoint: server.URL + "/oauth/token", CodeChallengeMethodsSupported: []string{"plain", "S256"}}
	oauthResponse, err := metadataResponse.Serialize()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	options := &LoginOptions{Server: server.URL, StartingKubeConfig: &kclientcmdapi.Config{}, Username: "test", Password: "test", Config: &restclient.Config{Host: server.URL}, IOStreams: genericclioptions.NewTestIOStreamsDiscard()}
	err = options.gatherAuthInfo()
	if err == nil {
		t.Fatalf("expecting unauthorized error when gathering authinfo")
	}
	if !kapierrs.IsUnauthorized(err) {
		t.Fatalf("expecting error of type metav1.StatusReasonUnauthorized, but got type %T: %v", err, err)
	}
}
func TestDialToHTTPSServer(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	invoked := make(chan struct{}, 1)
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		invoked <- struct{}{}
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()
	testCases := map[string]struct {
		serverURL	string
		skipTLSVerify	bool
		evalExpectedErr	func(error) bool
	}{"succeed dialing": {serverURL: server.URL, skipTLSVerify: true}}
	for name, test := range testCases {
		t.Logf("evaluating test: %s", name)
		clientConfig := &restclient.Config{Host: test.serverURL, TLSClientConfig: restclient.TLSClientConfig{Insecure: test.skipTLSVerify}}
		if err := dialToServer(*clientConfig); err != nil {
			if test.evalExpectedErr == nil || !test.evalExpectedErr(err) {
				t.Errorf("%s: unexpected error: %v", name, err)
			}
		} else {
			if test.evalExpectedErr != nil {
				t.Errorf("%s: expected error but got nothing", name)
			}
		}
	}
}
func newTLSServer(certString, keyString string) (*httptest.Server, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	invoked := make(chan struct{}, 1)
	server := httptest.NewUnstartedServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		invoked <- struct{}{}
		w.WriteHeader(http.StatusOK)
	}))
	cert, err := tls.X509KeyPair([]byte(certString), []byte(keyString))
	if err != nil {
		return nil, fmt.Errorf("error configuring server cert: %s", err)
	}
	server.TLS = &tls.Config{Certificates: []tls.Certificate{cert}}
	return server, nil
}
