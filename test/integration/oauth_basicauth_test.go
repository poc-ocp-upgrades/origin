package integration

import (
	"crypto/tls"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"testing"
	"time"
	"github.com/openshift/library-go/pkg/crypto"
	configapi "github.com/openshift/origin/pkg/cmd/server/apis/config"
	"github.com/openshift/origin/pkg/cmd/util"
	"github.com/openshift/origin/pkg/oc/lib/tokencmd"
	userclient "github.com/openshift/origin/pkg/user/generated/internalclientset/typed/user/internalversion"
	testutil "github.com/openshift/origin/test/util"
	testserver "github.com/openshift/origin/test/util/server"
	apierrs "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	restclient "k8s.io/client-go/rest"
)

const (
	basicAuthRemoteCACert		= "remote-ca.crt"
	basicAuthRemoteServer		= "remote-server"
	basicAuthRemoteServerCert	= "remote-server.crt"
	basicAuthRemoteServerKey	= "remote-server.key"
	basicAuthClient			= "client"
	basicAuthClientCert		= "client.crt"
	basicAuthClientKey		= "client.key"
)

func TestOAuthBasicAuthPassword(t *testing.T) {
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
	expectedLogin := "username"
	expectedPassword := "password"
	expectedAuthHeader := "Basic " + base64.StdEncoding.EncodeToString([]byte(expectedLogin+":"+expectedPassword))
	testcases := map[string]struct {
		RemoteStatus	int
		RemoteHeaders	http.Header
		RemoteBody	[]byte
		ExpectUsername	string
		ExpectSuccess	bool
		ExpectErrStatus	int32
	}{"success": {RemoteStatus: 200, RemoteHeaders: http.Header{"Content-Type": []string{"application/json"}}, RemoteBody: []byte(`{"sub":"remoteusername"}`), ExpectSuccess: true, ExpectUsername: "remoteusername"}, "401": {RemoteStatus: 401, RemoteHeaders: http.Header{"Content-Type": []string{"application/json"}}, RemoteBody: []byte(`{"error":"bad-user"}`), ExpectSuccess: false, ExpectUsername: "", ExpectErrStatus: 401}, "301": {RemoteStatus: 301, RemoteHeaders: http.Header{"Location": []string{"http://www.example.com"}}, ExpectSuccess: false, ExpectUsername: "", ExpectErrStatus: 500}, "302": {RemoteStatus: 302, RemoteHeaders: http.Header{"Location": []string{"http://www.example.com"}}, ExpectSuccess: false, ExpectUsername: "", ExpectErrStatus: 500}, "303": {RemoteStatus: 303, RemoteHeaders: http.Header{"Location": []string{"http://www.example.com"}}, ExpectSuccess: false, ExpectUsername: "", ExpectErrStatus: 500}, "304": {RemoteStatus: 304, RemoteHeaders: http.Header{"Location": []string{"http://www.example.com"}}, ExpectSuccess: false, ExpectUsername: "", ExpectErrStatus: 500}, "305": {RemoteStatus: 305, RemoteHeaders: http.Header{"Location": []string{"http://www.example.com"}}, ExpectSuccess: false, ExpectUsername: "", ExpectErrStatus: 500}, "404": {RemoteStatus: 404, ExpectSuccess: false, ExpectUsername: "", ExpectErrStatus: 500}, "500": {RemoteStatus: 500, ExpectSuccess: false, ExpectUsername: "", ExpectErrStatus: 500}}
	certNames := map[string]string{}
	basicAuthCAPrefix := "basicauthtest"
	certDir, err := ioutil.TempDir("", "oauthbasic")
	caCert, err := createCA(certDir, basicAuthCAPrefix)
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(certDir)
	t.Logf("cert dir is %s\n", certDir)
	certNames[basicAuthRemoteCACert] = caCert
	ip, err := util.DefaultLocalIP4()
	if err != nil {
		t.Fatal(err)
	}
	if _, err := createServerCert([]string{ip.String()}, basicAuthRemoteServer, certDir, basicAuthCAPrefix); err != nil {
		t.Fatal(err)
	}
	certNames[basicAuthRemoteServerCert] = filepath.Join(certDir, basicAuthRemoteServerCert)
	certNames[basicAuthRemoteServerKey] = filepath.Join(certDir, basicAuthRemoteServerKey)
	if _, err := createClientCert(basicAuthClient, certDir, basicAuthCAPrefix); err != nil {
		t.Fatal(err)
	}
	certNames[basicAuthClientCert] = filepath.Join(certDir, basicAuthClientCert)
	certNames[basicAuthClientKey] = filepath.Join(certDir, basicAuthClientKey)
	clientCAs, err := util.CertPoolFromFile(certNames[basicAuthRemoteCACert])
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var (
		remoteStatus	int
		remoteHeaders	http.Header
		remoteBody	[]byte
	)
	remoteHandler := http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if req.TLS == nil {
			w.WriteHeader(http.StatusUnauthorized)
			t.Fatalf("Expected TLS")
		}
		if len(req.TLS.VerifiedChains) != 1 {
			w.WriteHeader(http.StatusUnauthorized)
			t.Fatalf("Expected peer cert verified by server")
		}
		if req.Header.Get("Authorization") != expectedAuthHeader {
			w.WriteHeader(http.StatusUnauthorized)
			t.Fatalf("Expected auth header %s got %s", expectedAuthHeader, req.Header.Get("Authorization"))
		}
		for k, values := range remoteHeaders {
			for _, v := range values {
				w.Header().Add(k, v)
			}
		}
		w.WriteHeader(remoteStatus)
		w.Write(remoteBody)
	})
	remoteAddr, err := testserver.FindAvailableBindAddress(9443, 9999)
	if err != nil {
		t.Fatalf("Couldn't get free address for test server: %v", err)
	}
	remoteServer := &http.Server{Addr: remoteAddr, Handler: remoteHandler, ReadTimeout: 10 * time.Second, WriteTimeout: 10 * time.Second, MaxHeaderBytes: 1 << 20, TLSConfig: crypto.SecureTLSConfig(&tls.Config{ClientAuth: tls.RequireAndVerifyClientCert, ClientCAs: clientCAs})}
	go func() {
		if err := remoteServer.ListenAndServeTLS(certNames[basicAuthRemoteServerCert], certNames[basicAuthRemoteServerKey]); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	}()
	masterOptions, err := testserver.DefaultMasterOptions()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	defer testserver.CleanupMasterEtcd(t, masterOptions)
	masterOptions.OAuthConfig.IdentityProviders[0] = configapi.IdentityProvider{Name: "basicauth", UseAsChallenger: true, UseAsLogin: true, MappingMethod: "claim", Provider: &configapi.BasicAuthPasswordIdentityProvider{RemoteConnectionInfo: configapi.RemoteConnectionInfo{URL: fmt.Sprintf("https://%s", remoteAddr), CA: certNames[basicAuthRemoteCACert], ClientCert: configapi.CertInfo{CertFile: certNames[basicAuthClientCert], KeyFile: certNames[basicAuthClientKey]}}}}
	clusterAdminKubeConfig, err := testserver.StartConfiguredMaster(masterOptions)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	clientConfig, err := testutil.GetClusterAdminClientConfig(clusterAdminKubeConfig)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	anonConfig := restclient.Config{}
	anonConfig.Host = clientConfig.Host
	anonConfig.CAFile = clientConfig.CAFile
	anonConfig.CAData = clientConfig.CAData
	for k, tc := range testcases {
		remoteStatus = tc.RemoteStatus
		remoteHeaders = tc.RemoteHeaders
		remoteBody = tc.RemoteBody
		accessToken, err := tokencmd.RequestToken(&anonConfig, nil, expectedLogin, expectedPassword)
		if !tc.ExpectSuccess {
			if err == nil {
				t.Errorf("%s: Expected error, got token=%v", k, accessToken)
			} else if statusErr, ok := err.(*apierrs.StatusError); !ok {
				t.Errorf("%s: expected status error, got %#v", k, err)
			} else if statusErr.ErrStatus.Code != tc.ExpectErrStatus {
				t.Errorf("%s: expected error status %d, got %#v", k, tc.ExpectErrStatus, statusErr)
			}
			continue
		}
		if err != nil {
			t.Errorf("%s: Unexpected error: %v", k, err)
			continue
		}
		userConfig := anonConfig
		userConfig.BearerToken = accessToken
		userClient, err := userclient.NewForConfig(&userConfig)
		if err != nil {
			t.Fatalf("%s: Unexpected error: %v", k, err)
		}
		user, err := userClient.Users().Get("~", metav1.GetOptions{})
		if err != nil {
			t.Fatalf("%s: Unexpected error: %v", k, err)
		}
		if user.Name != tc.ExpectUsername {
			t.Fatalf("%s: Expected %v as the user, got %v", k, tc.ExpectUsername, user)
		}
	}
}
