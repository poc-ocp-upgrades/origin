package oauth

import (
	"errors"
	"time"
	"k8s.io/apimachinery/pkg/util/clock"
	"k8s.io/klog"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/runtime"
	oauthv1 "github.com/openshift/api/oauth/v1"
	userv1 "github.com/openshift/api/user/v1"
	oauthclient "github.com/openshift/client-go/oauth/clientset/versioned/typed/oauth/v1"
	oauthclientlister "github.com/openshift/client-go/oauth/listers/oauth/v1"
	"github.com/openshift/origin/pkg/util/rankedset"
)

var errTimedout = errors.New("token timed out")
var _ = rankedset.Item(&tokenData{})

type tokenData struct {
	token	*oauthv1.OAuthAccessToken
	seen	time.Time
}

func (a *tokenData) timeout() time.Time {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return a.token.CreationTimestamp.Time.Add(time.Duration(a.token.InactivityTimeoutSeconds) * time.Second)
}
func (a *tokenData) Key() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return a.token.Name
}
func (a *tokenData) Rank() int64 {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return a.timeout().Unix()
}
func timeoutAsDuration(timeout int32) time.Duration {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return time.Duration(timeout) * time.Second
}

type TimeoutValidator struct {
	oauthClients	oauthclientlister.OAuthClientLister
	tokens		oauthclient.OAuthAccessTokenInterface
	tokenChannel	chan *tokenData
	data		*rankedset.RankedSet
	defaultTimeout	time.Duration
	tickerInterval	time.Duration
	flushHandler	func(flushHorizon time.Time)
	putTokenHandler	func(td *tokenData)
	clock		clock.Clock
}

func NewTimeoutValidator(tokens oauthclient.OAuthAccessTokenInterface, oauthClients oauthclientlister.OAuthClientLister, defaultTimeout int32, minValidTimeout int32) *TimeoutValidator {
	_logClusterCodePath()
	defer _logClusterCodePath()
	a := &TimeoutValidator{oauthClients: oauthClients, tokens: tokens, tokenChannel: make(chan *tokenData), data: rankedset.New(), defaultTimeout: timeoutAsDuration(defaultTimeout), tickerInterval: timeoutAsDuration(minValidTimeout / 3), clock: clock.RealClock{}}
	a.flushHandler = a.flush
	a.putTokenHandler = a.putToken
	klog.V(5).Infof("Token Timeout Validator primed with defaultTimeout=%s tickerInterval=%s", a.defaultTimeout, a.tickerInterval)
	return a
}
func (a *TimeoutValidator) Validate(token *oauthv1.OAuthAccessToken, _ *userv1.User) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if token.InactivityTimeoutSeconds == 0 {
		return nil
	}
	td := &tokenData{token: token, seen: a.clock.Now()}
	if td.timeout().Before(td.seen) {
		return errTimedout
	}
	if token.ExpiresIn != 0 && token.ExpiresIn <= int64(token.InactivityTimeoutSeconds) {
		return nil
	}
	go a.putTokenHandler(td)
	return nil
}
func (a *TimeoutValidator) putToken(td *tokenData) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	a.tokenChannel <- td
}
func (a *TimeoutValidator) clientTimeout(name string) time.Duration {
	_logClusterCodePath()
	defer _logClusterCodePath()
	oauthClient, err := a.oauthClients.Get(name)
	if err != nil {
		klog.V(5).Infof("Failed to fetch OAuthClient %q for timeout value: %v", name, err)
		return a.defaultTimeout
	}
	if oauthClient.AccessTokenInactivityTimeoutSeconds == nil {
		return a.defaultTimeout
	}
	return timeoutAsDuration(*oauthClient.AccessTokenInactivityTimeoutSeconds)
}
func (a *TimeoutValidator) update(td *tokenData) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	delta := a.clientTimeout(td.token.ClientName)
	newTimeout := int32(0)
	if delta > 0 {
		newTimeout = int32((td.seen.Sub(td.token.CreationTimestamp.Time) + delta) / time.Second)
	}
	token, err := a.tokens.Get(td.token.Name, v1.GetOptions{})
	if err != nil {
		return err
	}
	if newTimeout != 0 && token.InactivityTimeoutSeconds >= newTimeout {
		return nil
	}
	token.InactivityTimeoutSeconds = newTimeout
	_, err = a.tokens.Update(token)
	return err
}
func (a *TimeoutValidator) flush(flushHorizon time.Time) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	klog.V(5).Infof("Flushing tokens timing out before %s", flushHorizon)
	tokenList := a.data.LessThan(flushHorizon.Unix(), true)
	var retryList []*tokenData
	flushedTokens := 0
	for _, item := range tokenList {
		td := item.(*tokenData)
		err := a.update(td)
		switch {
		case err == nil:
			flushedTokens++
		case apierrors.IsConflict(err) || apierrors.IsServerTimeout(err):
			klog.V(5).Infof("Token update deferred for token belonging to %s", td.token.UserName)
			retryList = append(retryList, td)
		default:
			klog.V(5).Infof("Token timeout for user=%q client=%q scopes=%v was not updated", td.token.UserName, td.token.ClientName, td.token.Scopes)
		}
	}
	for _, td := range retryList {
		err := a.update(td)
		if err != nil {
			klog.V(5).Infof("Token timeout for user=%q client=%q scopes=%v was not updated", td.token.UserName, td.token.ClientName, td.token.Scopes)
		} else {
			flushedTokens++
		}
	}
	klog.V(5).Infof("Successfully flushed %d tokens out of %d", flushedTokens, len(tokenList))
}
func (a *TimeoutValidator) nextTick() time.Time {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return a.clock.Now().Add(a.tickerInterval + 10*time.Second)
}
func (a *TimeoutValidator) Run(stopCh <-chan struct{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	defer runtime.HandleCrash()
	klog.V(5).Infof("Started Token Timeout Flush Handling thread!")
	ticker := a.clock.NewTicker(a.tickerInterval)
	defer ticker.Stop()
	nextTick := a.nextTick()
	for {
		select {
		case <-stopCh:
			return
		case td := <-a.tokenChannel:
			a.data.Insert(td)
			tokenTimeout := td.timeout()
			if tokenTimeout.Before(nextTick) {
				klog.V(5).Infof("Timeout for user=%q client=%q scopes=%v falls before next ticker (%s < %s), forcing flush!", td.token.UserName, td.token.ClientName, td.token.Scopes, tokenTimeout, nextTick)
				a.flushHandler(nextTick)
			}
		case <-ticker.C():
			nextTick = a.nextTick()
			a.flushHandler(nextTick)
		}
	}
}
