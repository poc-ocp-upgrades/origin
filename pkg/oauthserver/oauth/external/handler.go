package external

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	goformat "fmt"
	"github.com/RangelReale/osincli"
	authapi "github.com/openshift/origin/pkg/oauthserver/api"
	"github.com/openshift/origin/pkg/oauthserver/authenticator/identitymapper"
	"github.com/openshift/origin/pkg/oauthserver/oauth/handlers"
	"github.com/openshift/origin/pkg/oauthserver/server/csrf"
	"k8s.io/apiserver/pkg/authentication/authenticator"
	"k8s.io/apiserver/pkg/authentication/user"
	"k8s.io/klog"
	"net/http"
	"net/url"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

type Handler struct {
	provider     Provider
	state        State
	clientConfig *osincli.ClientConfig
	client       *osincli.Client
	success      handlers.AuthenticationSuccessHandler
	errorHandler handlers.AuthenticationErrorHandler
	mapper       authapi.UserIdentityMapper
}

func NewExternalOAuthRedirector(provider Provider, state State, redirectURL string, success handlers.AuthenticationSuccessHandler, errorHandler handlers.AuthenticationErrorHandler, mapper authapi.UserIdentityMapper) (handlers.AuthenticationRedirector, http.Handler, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	clientConfig, err := provider.NewConfig()
	if err != nil {
		return nil, nil, err
	}
	clientConfig.RedirectUrl = redirectURL
	client, err := osincli.NewClient(clientConfig)
	if err != nil {
		return nil, nil, err
	}
	transport, err := provider.GetTransport()
	if err != nil {
		return nil, nil, err
	}
	client.Transport = transport
	handler := &Handler{provider: provider, state: state, clientConfig: clientConfig, client: client, success: success, errorHandler: errorHandler, mapper: mapper}
	return handler, handler, nil
}
func (h *Handler) AuthenticationRedirect(w http.ResponseWriter, req *http.Request) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	klog.V(4).Infof("Authentication needed for %v", h.provider)
	authReq := h.client.NewAuthorizeRequest(osincli.CODE)
	h.provider.AddCustomParameters(authReq)
	state, err := h.state.Generate(w, req)
	if err != nil {
		klog.V(4).Infof("Error generating state: %v", err)
		return err
	}
	oauthURL := authReq.GetAuthorizeUrlWithParams(state)
	klog.V(4).Infof("redirect to %v", oauthURL)
	http.Redirect(w, req, oauthURL.String(), http.StatusFound)
	return nil
}
func NewOAuthPasswordAuthenticator(provider Provider, mapper authapi.UserIdentityMapper) (authenticator.Password, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	clientConfig, err := provider.NewConfig()
	if err != nil {
		return nil, err
	}
	clientConfig.RedirectUrl = "/"
	client, err := osincli.NewClient(clientConfig)
	if err != nil {
		return nil, err
	}
	transport, err := provider.GetTransport()
	if err != nil {
		return nil, err
	}
	client.Transport = transport
	return &Handler{provider: provider, clientConfig: clientConfig, client: client, mapper: mapper}, nil
}
func (h *Handler) AuthenticatePassword(ctx context.Context, username, password string) (*authenticator.Response, bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	accessReq := h.client.NewAccessRequest(osincli.PASSWORD, &osincli.AuthorizeData{Username: username, Password: password})
	accessData, err := accessReq.GetToken()
	if err != nil {
		if oauthErr, ok := err.(*osincli.Error); ok && oauthErr.Id == "invalid_grant" {
			return nil, false, nil
		}
		klog.V(4).Infof("Error getting access token using resource owner password grant: %v", err)
		return nil, false, err
	}
	klog.V(5).Infof("Got access data for %s", username)
	identity, ok, err := h.provider.GetUserIdentity(accessData)
	if err != nil {
		klog.V(4).Infof("Error getting userIdentityInfo info: %v", err)
		return nil, false, err
	}
	if !ok {
		klog.V(4).Infof("Could not get userIdentityInfo info from access token")
		err := errors.New("Could not get userIdentityInfo info from access token")
		return nil, false, err
	}
	return identitymapper.ResponseFor(h.mapper, identity)
}
func (h *Handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	authReq := h.client.NewAuthorizeRequest(osincli.CODE)
	authData, err := authReq.HandleRequest(req)
	if err != nil {
		klog.V(4).Infof("Error handling request: %v", err)
		h.handleError(err, w, req)
		return
	}
	klog.V(4).Infof("Got auth data")
	ok, err := h.state.Check(authData.State, req)
	if err != nil {
		klog.V(4).Infof("Error verifying state: %v", err)
		h.handleError(err, w, req)
		return
	}
	if !ok {
		klog.V(4).Infof("State is invalid")
		err := errors.New("State is invalid")
		h.handleError(err, w, req)
		return
	}
	accessReq := h.client.NewAccessRequest(osincli.AUTHORIZATION_CODE, authData)
	accessData, err := accessReq.GetToken()
	if err != nil {
		klog.V(4).Infof("Error getting access token: %v", err)
		h.handleError(err, w, req)
		return
	}
	klog.V(5).Infof("Got access data")
	identity, ok, err := h.provider.GetUserIdentity(accessData)
	if err != nil {
		klog.V(4).Infof("Error getting userIdentityInfo info: %v", err)
		h.handleError(err, w, req)
		return
	}
	if !ok {
		klog.V(4).Infof("Could not get userIdentityInfo info from access token")
		err := errors.New("Could not get userIdentityInfo info from access token")
		h.handleError(err, w, req)
		return
	}
	user, err := h.mapper.UserFor(identity)
	if err != nil {
		klog.V(4).Infof("Error creating or updating mapping for: %#v due to %v", identity, err)
		h.handleError(err, w, req)
		return
	}
	klog.V(4).Infof("Got userIdentityMapping: %#v", user)
	_, err = h.success.AuthenticationSucceeded(user, authData.State, w, req)
	if err != nil {
		klog.V(4).Infof("Error calling success handler: %v", err)
		h.handleError(err, w, req)
		return
	}
}
func (h *Handler) handleError(err error, w http.ResponseWriter, req *http.Request) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	handled, err := h.errorHandler.AuthenticationError(err, w, req)
	if handled {
		return
	}
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(`An error occurred`))
}

type defaultState struct{ csrf csrf.CSRF }
type RedirectorState interface {
	State
	handlers.AuthenticationSuccessHandler
	handlers.AuthenticationErrorHandler
}

func CSRFRedirectingState(csrf csrf.CSRF) RedirectorState {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &defaultState{csrf: csrf}
}
func (d *defaultState) Generate(w http.ResponseWriter, req *http.Request) (string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	then := req.URL.String()
	if len(then) == 0 {
		return "", errors.New("cannot generate state: request has no URL")
	}
	state := url.Values{"csrf": {d.csrf.Generate(w, req)}, "then": {then}}
	return encodeState(state), nil
}
func (d *defaultState) Check(state string, req *http.Request) (bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	values, err := decodeState(state)
	if err != nil {
		return false, err
	}
	if ok := d.csrf.Check(req, values.Get("csrf")); !ok {
		return false, fmt.Errorf("state did not contain a valid CSRF token")
	}
	if then := values.Get("then"); len(then) == 0 {
		return false, errors.New("state did not contain a redirect")
	}
	return true, nil
}
func (d *defaultState) AuthenticationSucceeded(user user.Info, state string, w http.ResponseWriter, req *http.Request) (bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	values, err := decodeState(state)
	if err != nil {
		return false, err
	}
	then := values.Get("then")
	if len(then) == 0 {
		return false, errors.New("no redirect given")
	}
	http.Redirect(w, req, then, http.StatusFound)
	return true, nil
}
func (d *defaultState) AuthenticationError(err error, w http.ResponseWriter, req *http.Request) (bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	osinErr, ok := err.(*osincli.Error)
	if !ok {
		return false, err
	}
	if len(osinErr.Id) == 0 {
		return false, err
	}
	ok, stateErr := d.Check(osinErr.State, req)
	if !ok || stateErr != nil {
		return false, err
	}
	values, err := decodeState(osinErr.State)
	if err != nil {
		return false, err
	}
	then := values.Get("then")
	if len(then) == 0 {
		return false, err
	}
	thenURL, urlErr := url.Parse(then)
	if urlErr != nil {
		return false, err
	}
	q := thenURL.Query()
	q.Set("error", osinErr.Id)
	if len(osinErr.Description) > 0 {
		q.Set("error_description", osinErr.Description)
	}
	if len(osinErr.URI) > 0 {
		q.Set("error_uri", osinErr.URI)
	}
	thenURL.RawQuery = q.Encode()
	http.Redirect(w, req, thenURL.String(), http.StatusFound)
	return true, nil
}
func encodeState(values url.Values) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return base64.URLEncoding.EncodeToString([]byte(values.Encode()))
}
func decodeState(state string) (url.Values, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	decodedState, err := base64.URLEncoding.DecodeString(state)
	if err != nil {
		return nil, err
	}
	return url.ParseQuery(string(decodedState))
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
