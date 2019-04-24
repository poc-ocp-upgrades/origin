package oauth

import (
	"context"
	"errors"
	"testing"
	"time"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/clock"
	"k8s.io/apiserver/pkg/authentication/authenticator"
	clienttesting "k8s.io/client-go/testing"
	oauthv1 "github.com/openshift/api/oauth/v1"
	userapi "github.com/openshift/api/user/v1"
	oauthfake "github.com/openshift/client-go/oauth/clientset/versioned/fake"
	oauthclient "github.com/openshift/client-go/oauth/clientset/versioned/typed/oauth/v1"
	userfake "github.com/openshift/client-go/user/clientset/versioned/fake"
)

func TestAuthenticateTokenInvalidUID(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fakeOAuthClient := oauthfake.NewSimpleClientset(&oauthv1.OAuthAccessToken{ObjectMeta: metav1.ObjectMeta{Name: "token", CreationTimestamp: metav1.Time{Time: time.Now()}}, ExpiresIn: 600, UserName: "foo", UserUID: string("bar1")})
	fakeUserClient := userfake.NewSimpleClientset(&userapi.User{ObjectMeta: metav1.ObjectMeta{Name: "foo", UID: "bar2"}})
	tokenAuthenticator := NewTokenAuthenticator(fakeOAuthClient.OauthV1().OAuthAccessTokens(), fakeUserClient.UserV1().Users(), NoopGroupMapper{}, NewUIDValidator())
	userInfo, found, err := tokenAuthenticator.AuthenticateToken(context.TODO(), "token")
	if found {
		t.Error("Found token, but it should be missing!")
	}
	if err.Error() != "user.UID (bar2) does not match token.userUID (bar1)" {
		t.Errorf("Unexpected error: %v", err)
	}
	if userInfo != nil {
		t.Errorf("Unexpected user: %v", userInfo)
	}
}
func TestAuthenticateTokenNotFoundSuppressed(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fakeOAuthClient := oauthfake.NewSimpleClientset()
	fakeUserClient := userfake.NewSimpleClientset()
	tokenAuthenticator := NewTokenAuthenticator(fakeOAuthClient.OauthV1().OAuthAccessTokens(), fakeUserClient.UserV1().Users(), NoopGroupMapper{})
	userInfo, found, err := tokenAuthenticator.AuthenticateToken(context.TODO(), "token")
	if found {
		t.Error("Found token, but it should be missing!")
	}
	if err != errLookup {
		t.Error("Expected not found error to be suppressed with lookup error")
	}
	if userInfo != nil {
		t.Errorf("Unexpected user: %v", userInfo)
	}
}
func TestAuthenticateTokenOtherGetErrorSuppressed(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fakeOAuthClient := oauthfake.NewSimpleClientset()
	fakeOAuthClient.PrependReactor("get", "oauthaccesstokens", func(action clienttesting.Action) (handled bool, ret runtime.Object, err error) {
		return true, nil, errors.New("get error")
	})
	fakeUserClient := userfake.NewSimpleClientset()
	tokenAuthenticator := NewTokenAuthenticator(fakeOAuthClient.OauthV1().OAuthAccessTokens(), fakeUserClient.UserV1().Users(), NoopGroupMapper{})
	userInfo, found, err := tokenAuthenticator.AuthenticateToken(context.TODO(), "token")
	if found {
		t.Error("Found token, but it should be missing!")
	}
	if err != errLookup {
		t.Error("Expected custom get error to be suppressed with lookup error")
	}
	if userInfo != nil {
		t.Errorf("Unexpected user: %v", userInfo)
	}
}
func TestAuthenticateTokenTimeout(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	stopCh := make(chan struct{})
	defer close(stopCh)
	testClock := clock.NewFakeClock(time.Time{})
	defaultTimeout := int32(30)
	clientTimeout := int32(15)
	minTimeout := int32(10)
	testClient := oauthv1.OAuthClient{ObjectMeta: metav1.ObjectMeta{Name: "testClient"}, AccessTokenInactivityTimeoutSeconds: &clientTimeout}
	quickClient := oauthv1.OAuthClient{ObjectMeta: metav1.ObjectMeta{Name: "quickClient"}, AccessTokenInactivityTimeoutSeconds: &minTimeout}
	slowClient := oauthv1.OAuthClient{ObjectMeta: metav1.ObjectMeta{Name: "slowClient"}}
	testToken := oauthv1.OAuthAccessToken{ObjectMeta: metav1.ObjectMeta{Name: "testToken", CreationTimestamp: metav1.Time{Time: testClock.Now()}}, ClientName: "testClient", ExpiresIn: 600, UserName: "foo", UserUID: string("bar"), InactivityTimeoutSeconds: clientTimeout}
	quickToken := oauthv1.OAuthAccessToken{ObjectMeta: metav1.ObjectMeta{Name: "quickToken", CreationTimestamp: metav1.Time{Time: testClock.Now()}}, ClientName: "quickClient", ExpiresIn: 600, UserName: "foo", UserUID: string("bar"), InactivityTimeoutSeconds: minTimeout}
	slowToken := oauthv1.OAuthAccessToken{ObjectMeta: metav1.ObjectMeta{Name: "slowToken", CreationTimestamp: metav1.Time{Time: testClock.Now()}}, ClientName: "slowClient", ExpiresIn: 600, UserName: "foo", UserUID: string("bar"), InactivityTimeoutSeconds: defaultTimeout}
	emergToken := oauthv1.OAuthAccessToken{ObjectMeta: metav1.ObjectMeta{Name: "emergToken", CreationTimestamp: metav1.Time{Time: testClock.Now()}}, ClientName: "quickClient", ExpiresIn: 600, UserName: "foo", UserUID: string("bar"), InactivityTimeoutSeconds: 5}
	fakeOAuthClient := oauthfake.NewSimpleClientset(&testToken, &quickToken, &slowToken, &emergToken, &testClient, &quickClient, &slowClient)
	fakeUserClient := userfake.NewSimpleClientset(&userapi.User{ObjectMeta: metav1.ObjectMeta{Name: "foo", UID: "bar"}})
	accessTokenGetter := fakeOAuthClient.OauthV1().OAuthAccessTokens()
	oauthClients := fakeOAuthClient.OauthV1().OAuthClients()
	lister := &fakeOAuthClientLister{clients: oauthClients}
	timeouts := NewTimeoutValidator(accessTokenGetter, lister, defaultTimeout, minTimeout)
	timeouts.clock = testClock
	originalFlush := timeouts.flushHandler
	timeoutsSync := make(chan struct{})
	timeouts.flushHandler = func(flushHorizon time.Time) {
		originalFlush(flushHorizon)
		timeoutsSync <- struct{}{}
	}
	originalPutToken := timeouts.putTokenHandler
	putTokenSync := make(chan struct{})
	timeouts.putTokenHandler = func(td *tokenData) {
		originalPutToken(td)
		putTokenSync <- struct{}{}
	}
	buffer := time.Nanosecond
	tokenAuthenticator := NewTokenAuthenticator(accessTokenGetter, fakeUserClient.UserV1().Users(), NoopGroupMapper{}, timeouts)
	go timeouts.Run(stopCh)
	checkToken(t, "testToken", tokenAuthenticator, accessTokenGetter, testClock, true)
	wait(t, putTokenSync)
	checkToken(t, "quickToken", tokenAuthenticator, accessTokenGetter, testClock, true)
	wait(t, putTokenSync)
	wait(t, timeoutsSync)
	checkToken(t, "slowToken", tokenAuthenticator, accessTokenGetter, testClock, true)
	wait(t, putTokenSync)
	checkToken(t, "emergToken", tokenAuthenticator, accessTokenGetter, testClock, true)
	wait(t, putTokenSync)
	wait(t, timeoutsSync)
	testClock.Sleep(5*time.Second + buffer)
	wait(t, timeoutsSync)
	checkToken(t, "emergToken", tokenAuthenticator, accessTokenGetter, testClock, true)
	wait(t, putTokenSync)
	wait(t, timeoutsSync)
	testClock.Sleep(time.Duration(minTimeout-5)*time.Second + buffer)
	wait(t, timeoutsSync)
	changeClient, ret := oauthClients.Get("testClient", metav1.GetOptions{})
	if ret != nil {
		t.Error("Failed to get testClient")
	} else {
		longTimeout := int32(20)
		changeClient.AccessTokenInactivityTimeoutSeconds = &longTimeout
		_, ret = oauthClients.Update(changeClient)
		if ret != nil {
			t.Error("Failed to update testClient")
		}
	}
	checkToken(t, "quickToken", tokenAuthenticator, accessTokenGetter, testClock, false)
	checkToken(t, "testToken", tokenAuthenticator, accessTokenGetter, testClock, true)
	wait(t, putTokenSync)
	wait(t, timeoutsSync)
	testClock.Sleep(time.Duration(clientTimeout+1)*time.Second + buffer)
	wait(t, timeoutsSync)
	checkToken(t, "slowToken", tokenAuthenticator, accessTokenGetter, testClock, true)
	wait(t, putTokenSync)
	wait(t, timeoutsSync)
	checkToken(t, "testToken", tokenAuthenticator, accessTokenGetter, testClock, true)
	wait(t, putTokenSync)
	wait(t, timeoutsSync)
	token, err := accessTokenGetter.Get("testToken", metav1.GetOptions{})
	if err != nil {
		t.Error("Failed to get testToken")
	} else {
		if token.InactivityTimeoutSeconds < 31 {
			t.Errorf("Expected timeout in more than 31 seconds, found: %d", token.InactivityTimeoutSeconds)
		}
	}
	changeclient, ret := oauthClients.Get("testClient", metav1.GetOptions{})
	if ret != nil {
		t.Error("Failed to get testClient")
	} else {
		changeclient.AccessTokenInactivityTimeoutSeconds = new(int32)
		_, ret = oauthClients.Update(changeclient)
		if ret != nil {
			t.Error("Failed to update testClient")
		}
	}
	testClock.Sleep(time.Duration(minTimeout)*time.Second + buffer)
	wait(t, timeoutsSync)
	checkToken(t, "testToken", tokenAuthenticator, accessTokenGetter, testClock, true)
	wait(t, putTokenSync)
	wait(t, timeoutsSync)
	token, err = accessTokenGetter.Get("testToken", metav1.GetOptions{})
	if err != nil {
		t.Error("Failed to get testToken")
	} else {
		if token.InactivityTimeoutSeconds != 0 {
			t.Errorf("Expected timeout of 0 seconds, found: %d", token.InactivityTimeoutSeconds)
		}
	}
}

type fakeOAuthClientLister struct {
	clients oauthclient.OAuthClientInterface
}

func (f fakeOAuthClientLister) Get(name string) (*oauthv1.OAuthClient, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return f.clients.Get(name, metav1.GetOptions{})
}
func (f fakeOAuthClientLister) List(selector labels.Selector) ([]*oauthv1.OAuthClient, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	panic("not used")
}
func checkToken(t *testing.T, name string, authf authenticator.Token, tokens oauthclient.OAuthAccessTokenInterface, current clock.Clock, present bool) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	t.Helper()
	userInfo, found, err := authf.AuthenticateToken(context.TODO(), name)
	if present {
		if !found {
			t.Errorf("Did not find token %s!", name)
		}
		if err != nil {
			t.Errorf("Unexpected error checking for token %s: %v", name, err)
		}
		if userInfo == nil {
			t.Errorf("Did not get a user for token %s!", name)
		}
	} else {
		if found {
			token, tokenErr := tokens.Get(name, metav1.GetOptions{})
			if tokenErr != nil {
				t.Fatal(tokenErr)
			}
			t.Errorf("Found token (created=%s, timeout=%di, now=%s), but it should be gone!", token.CreationTimestamp, token.InactivityTimeoutSeconds, current.Now())
		}
		if err != errTimedout {
			t.Errorf("Unexpected error checking absence of token %s: %v", name, err)
		}
		if userInfo != nil {
			t.Errorf("Unexpected user checking absence of token %s: %v", name, userInfo)
		}
	}
}
func wait(t *testing.T, c chan struct{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	t.Helper()
	select {
	case <-c:
	case <-time.After(30 * time.Second):
		t.Fatal("failed to see channel event")
	}
}
