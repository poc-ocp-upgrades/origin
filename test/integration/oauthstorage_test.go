package integration

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/RangelReale/osin"
	"github.com/RangelReale/osincli"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kuser "k8s.io/apiserver/pkg/authentication/user"
	oauthapi "github.com/openshift/api/oauth/v1"
	oauthclient "github.com/openshift/client-go/oauth/clientset/versioned/typed/oauth/v1"
	"github.com/openshift/origin/pkg/oauthserver/osinserver"
	"github.com/openshift/origin/pkg/oauthserver/osinserver/registrystorage"
	testutil "github.com/openshift/origin/test/util"
	testserver "github.com/openshift/origin/test/util/server"
)

func TestOAuthStorage(t *testing.T) {
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
	masterOptions, clusterAdminKubeConfig, err := testserver.StartTestMasterAPI()
	if err != nil {
		t.Fatal(err)
	}
	defer testserver.CleanupMasterEtcd(t, masterOptions)
	clusterAdminClientConfig := testutil.GetClusterAdminClientConfigOrDie(clusterAdminKubeConfig)
	oauthClient := oauthclient.NewForConfigOrDie(clusterAdminClientConfig)
	storage := registrystorage.New(oauthClient.OAuthAccessTokens(), oauthClient.OAuthAuthorizeTokens(), oauthClient.OAuthClients(), 0)
	const (
		testUser		= "test-uid"
		testUID			= "007"
		testClient		= "test-client"
		testClientSecret0	= "secret0"
		testClientSecret1	= "secret1"
		testClientRedirect	= "/assert"
		authorizePath		= "/authorize"
		tokenPath		= "/token"
	)
	oauthServer := osinserver.New(osinserver.NewDefaultServerConfig(), storage, osinserver.AuthorizeHandlerFunc(func(ar *osin.AuthorizeRequest, resp *osin.Response, w http.ResponseWriter) (bool, error) {
		ar.UserData = &kuser.DefaultInfo{Name: testUser, UID: testUID}
		ar.Authorized = true
		return false, nil
	}), osinserver.AccessHandlerFunc(func(ar *osin.AccessRequest, w http.ResponseWriter) error {
		ar.UserData = ar.AuthorizeData.UserData
		ar.Authorized = true
		return nil
	}), osinserver.NewDefaultErrorHandler())
	mux := http.NewServeMux()
	oauthServer.Install(mux, "")
	server := httptest.NewServer(mux)
	defer server.Close()
	tokenChannel := make(chan string, 1)
	var oaclient *osincli.Client
	var authReq *osincli.AuthorizeRequest
	assertServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		data, err := authReq.HandleRequest(r)
		if err != nil {
			t.Fatal(err)
		}
		tokenReq := oaclient.NewAccessRequest(osincli.AUTHORIZATION_CODE, data)
		token, err := tokenReq.GetToken()
		if err != nil {
			t.Fatal(err)
		}
		tokenChannel <- token.AccessToken
	}))
	oaclientConfig := &osincli.ClientConfig{ClientId: testClient, ClientSecret: testClientSecret0, RedirectUrl: assertServer.URL + testClientRedirect, AuthorizeUrl: server.URL + authorizePath, TokenUrl: server.URL + tokenPath}
	oaclient, err = osincli.NewClient(oaclientConfig)
	if err != nil {
		t.Fatal(err)
	}
	authReq = oaclient.NewAuthorizeRequest(osincli.CODE)
	{
		if _, err := oauthClient.OAuthClients().Create(&oauthapi.OAuthClient{ObjectMeta: metav1.ObjectMeta{Name: testClient}, Secret: testClientSecret0, AdditionalSecrets: []string{testClientSecret1}, RedirectURIs: []string{assertServer.URL + testClientRedirect}}); err != nil {
			t.Fatal(err)
		}
		storedClient, err := storage.GetClient(testClient)
		if err != nil {
			t.Fatalf("unexpected get client error: %v", err)
		}
		if !osin.CheckClientSecret(storedClient, testClientSecret0) {
			t.Fatalf("unexpected stored client secret failure: %#v", storedClient)
		}
		if !osin.CheckClientSecret(storedClient, testClientSecret1) {
			t.Fatalf("unexpected stored client secret failure: %#v", storedClient)
		}
		if osin.CheckClientSecret(storedClient, "secret2") {
			t.Fatalf("unexpected stored client secret success: %#v", storedClient)
		}
	}
	{
		resp, err := http.DefaultClient.Get(authReq.GetAuthorizeUrl().String())
		if err != nil {
			t.Fatal(err)
		}
		if resp.StatusCode != http.StatusOK {
			t.Fatalf("unexpected response: %#v", resp)
		}
		var token string
		select {
		case token = <-tokenChannel:
		default:
			t.Fatal("unable to retrieve access token")
		}
		if len(token) == 0 {
			t.Fatalf("unexpected empty access token: %#v", token)
		}
		actualToken, err := oauthClient.OAuthAccessTokens().Get(token, metav1.GetOptions{})
		if err != nil {
			t.Fatal(err)
		}
		if actualToken.UserUID != testUID || actualToken.UserName != testUser {
			t.Fatalf("unexpected stored token: %#v", actualToken)
		}
	}
}
