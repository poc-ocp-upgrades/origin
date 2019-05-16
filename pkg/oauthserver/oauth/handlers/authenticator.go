package handlers

import (
	"context"
	goformat "fmt"
	"github.com/RangelReale/osin"
	"github.com/openshift/origin/pkg/oauthserver/api"
	openshiftauthenticator "github.com/openshift/origin/pkg/oauthserver/authenticator"
	"github.com/openshift/origin/pkg/oauthserver/osinserver"
	"k8s.io/apiserver/pkg/authentication/authenticator"
	"k8s.io/klog"
	"net/http"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

type authorizeAuthenticator struct {
	request      authenticator.Request
	handler      AuthenticationHandler
	errorHandler AuthenticationErrorHandler
}

func NewAuthorizeAuthenticator(request authenticator.Request, handler AuthenticationHandler, errorHandler AuthenticationErrorHandler) osinserver.AuthorizeHandler {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &authorizeAuthenticator{request: request, handler: handler, errorHandler: errorHandler}
}

type TokenMaxAgeSeconds interface{ GetTokenMaxAgeSeconds() *int32 }
type TokenTimeoutSeconds interface{ GetAccessTokenInactivityTimeoutSeconds() *int32 }

func (h *authorizeAuthenticator) HandleAuthorize(ar *osin.AuthorizeRequest, resp *osin.Response, w http.ResponseWriter) (bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	info, ok, err := h.request.AuthenticateRequest(ar.HttpRequest)
	if err != nil {
		klog.V(4).Infof("OAuth authentication error: %v", err)
		return h.errorHandler.AuthenticationError(err, w, ar.HttpRequest)
	}
	if !ok {
		return h.handler.AuthenticationNeeded(ar.Client, w, ar.HttpRequest)
	}
	klog.V(4).Infof("OAuth authentication succeeded: %#v", info.User)
	ar.UserData = info.User
	ar.Authorized = true
	if ar.Type == osin.TOKEN {
		if e, ok := ar.Client.(TokenMaxAgeSeconds); ok {
			if maxAge := e.GetTokenMaxAgeSeconds(); maxAge != nil {
				ar.Expiration = *maxAge
			}
		}
	}
	return false, nil
}

type accessAuthenticator struct {
	password  authenticator.Password
	assertion openshiftauthenticator.Assertion
	client    openshiftauthenticator.Client
}

func (h *accessAuthenticator) HandleAccess(ar *osin.AccessRequest, w http.ResponseWriter) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var (
		info *authenticator.Response
		ok   bool
		err  error
	)
	switch ar.Type {
	case osin.AUTHORIZATION_CODE, osin.REFRESH_TOKEN:
		ok = true
	case osin.PASSWORD:
		ctx := context.TODO()
		if ar.HttpRequest != nil {
			ctx = ar.HttpRequest.Context()
		}
		info, ok, err = h.password.AuthenticatePassword(ctx, ar.Username, ar.Password)
	case osin.ASSERTION:
		info, ok, err = h.assertion.AuthenticateAssertion(ar.AssertionType, ar.Assertion)
	case osin.CLIENT_CREDENTIALS:
		info, ok, err = h.client.AuthenticateClient(ar.Client)
	default:
		klog.Warningf("Received unknown access token type: %s", ar.Type)
	}
	if err != nil {
		klog.V(4).Infof("Unable to authenticate %s: %v", ar.Type, err)
		return err
	}
	if ok {
		ar.GenerateRefresh = false
		ar.Authorized = true
		if info != nil {
			ar.AccessData.UserData = info.User
		}
		if e, ok := ar.Client.(TokenMaxAgeSeconds); ok {
			if maxAge := e.GetTokenMaxAgeSeconds(); maxAge != nil {
				ar.Expiration = *maxAge
			}
		}
	}
	return nil
}
func NewDenyAccessAuthenticator() osinserver.AccessHandler {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &accessAuthenticator{password: deny, assertion: deny, client: deny}
}

var deny = &denyAuthenticator{}

type denyAuthenticator struct{}

func (*denyAuthenticator) AuthenticatePassword(ctx context.Context, user, password string) (*authenticator.Response, bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return nil, false, nil
}
func (*denyAuthenticator) AuthenticateAssertion(assertionType, data string) (*authenticator.Response, bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return nil, false, nil
}
func (*denyAuthenticator) AuthenticateClient(client api.Client) (*authenticator.Response, bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return nil, false, nil
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
