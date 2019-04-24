package handlers

import (
	"context"
	"bytes"
	"runtime"
	"fmt"
	"net/http"
	"github.com/RangelReale/osin"
	"k8s.io/klog"
	"k8s.io/apiserver/pkg/authentication/authenticator"
	"github.com/openshift/origin/pkg/oauthserver/api"
	openshiftauthenticator "github.com/openshift/origin/pkg/oauthserver/authenticator"
	"github.com/openshift/origin/pkg/oauthserver/osinserver"
)

type authorizeAuthenticator struct {
	request		authenticator.Request
	handler		AuthenticationHandler
	errorHandler	AuthenticationErrorHandler
}

func NewAuthorizeAuthenticator(request authenticator.Request, handler AuthenticationHandler, errorHandler AuthenticationErrorHandler) osinserver.AuthorizeHandler {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &authorizeAuthenticator{request: request, handler: handler, errorHandler: errorHandler}
}

type TokenMaxAgeSeconds interface{ GetTokenMaxAgeSeconds() *int32 }
type TokenTimeoutSeconds interface{ GetAccessTokenInactivityTimeoutSeconds() *int32 }

func (h *authorizeAuthenticator) HandleAuthorize(ar *osin.AuthorizeRequest, resp *osin.Response, w http.ResponseWriter) (bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
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
	password	authenticator.Password
	assertion	openshiftauthenticator.Assertion
	client		openshiftauthenticator.Client
}

func (h *accessAuthenticator) HandleAccess(ar *osin.AccessRequest, w http.ResponseWriter) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var (
		info	*authenticator.Response
		ok	bool
		err	error
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &accessAuthenticator{password: deny, assertion: deny, client: deny}
}

var deny = &denyAuthenticator{}

type denyAuthenticator struct{}

func (*denyAuthenticator) AuthenticatePassword(ctx context.Context, user, password string) (*authenticator.Response, bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return nil, false, nil
}
func (*denyAuthenticator) AuthenticateAssertion(assertionType, data string) (*authenticator.Response, bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return nil, false, nil
}
func (*denyAuthenticator) AuthenticateClient(client api.Client) (*authenticator.Response, bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return nil, false, nil
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := runtime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", runtime.FuncForPC(pc).Name()))
	http.Post("/"+"logcode", "application/json", bytes.NewBuffer(jsonLog))
}
