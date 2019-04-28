package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"
	"github.com/RangelReale/osin"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/apiserver/pkg/authentication/user"
	oauthapi "github.com/openshift/api/oauth/v1"
	scopeauthorizer "github.com/openshift/origin/pkg/authorization/authorizer/scope"
	"github.com/openshift/origin/pkg/oauth/apis/oauth/validation"
	"github.com/openshift/origin/pkg/oauth/scope"
	"github.com/openshift/origin/pkg/oauthserver/api"
	"github.com/openshift/origin/pkg/oauthserver/osinserver"
)

type GrantCheck struct {
	check		GrantChecker
	handler		GrantHandler
	errorHandler	GrantErrorHandler
}

func NewGrantCheck(check GrantChecker, handler GrantHandler, errorHandler GrantErrorHandler) osinserver.AuthorizeHandler {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &GrantCheck{check, handler, errorHandler}
}
func (h *GrantCheck) HandleAuthorize(ar *osin.AuthorizeRequest, resp *osin.Response, w http.ResponseWriter) (bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if !ar.Authorized {
		return false, nil
	}
	ar.Authorized = false
	user, ok := ar.UserData.(user.Info)
	if !ok || user == nil {
		utilruntime.HandleError(fmt.Errorf("the provided user data is not a user.Info object: %#v", user))
		resp.SetError("server_error", "")
		return false, nil
	}
	client, ok := ar.Client.GetUserData().(*oauthapi.OAuthClient)
	if !ok || client == nil {
		utilruntime.HandleError(fmt.Errorf("the provided client is not an *api.OAuthClient object: %#v", client))
		resp.SetError("server_error", "")
		return false, nil
	}
	scopes := scope.Split(ar.Scope)
	if len(scopes) == 0 {
		scopes = append(scopes, scopeauthorizer.UserFull)
	}
	ar.Scope = scope.Join(scopes)
	if scopeErrors := validation.ValidateScopes(scopes, nil); len(scopeErrors) > 0 {
		resp.SetError("invalid_scope", scopeErrors.ToAggregate().Error())
		return false, nil
	}
	invalidScopes := sets.NewString()
	for _, scope := range scopes {
		if err := scopeauthorizer.ValidateScopeRestrictions(client, scope); err != nil {
			invalidScopes.Insert(scope)
		}
	}
	if len(invalidScopes) > 0 {
		resp.SetError("access_denied", fmt.Sprintf("scope denied: %s", strings.Join(invalidScopes.List(), " ")))
		return false, nil
	}
	grant := &api.Grant{Client: ar.Client, Scope: ar.Scope, Expiration: int64(ar.Expiration), RedirectURI: ar.RedirectUri}
	authorized, err := h.check.HasAuthorizedClient(user, grant)
	if err != nil {
		utilruntime.HandleError(err)
		resp.SetError("server_error", "")
		return false, nil
	}
	if authorized {
		ar.Authorized = true
		return false, nil
	}
	authorized, handled, err := h.handler.GrantNeeded(user, grant, w, ar.HttpRequest)
	if authorized {
		ar.Authorized = true
	}
	return handled, err
}

type emptyGrant struct{}

func NewEmptyGrant() GrantHandler {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return emptyGrant{}
}
func (emptyGrant) GrantNeeded(user user.Info, grant *api.Grant, w http.ResponseWriter, req *http.Request) (bool, bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return false, false, nil
}

type autoGrant struct{}

func NewAutoGrant() GrantHandler {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &autoGrant{}
}
func (g *autoGrant) GrantNeeded(user user.Info, grant *api.Grant, w http.ResponseWriter, req *http.Request) (bool, bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return true, false, nil
}

type redirectGrant struct{ subpath string }

func NewRedirectGrant(subpath string) GrantHandler {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &redirectGrant{subpath}
}
func (g *redirectGrant) GrantNeeded(user user.Info, grant *api.Grant, w http.ResponseWriter, req *http.Request) (bool, bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_, lastSegment := path.Split(req.URL.Path)
	reqURL := &(*req.URL)
	reqURL.Host = ""
	reqURL.Scheme = ""
	reqURL.Path = path.Join("..", lastSegment)
	redirectURL := &url.URL{Path: path.Join(lastSegment, g.subpath), RawQuery: url.Values{"then": {reqURL.String()}, "client_id": {grant.Client.GetId()}, "scope": {grant.Scope}, "redirect_uri": {grant.RedirectURI}}.Encode()}
	w.Header().Set("Location", redirectURL.String())
	w.WriteHeader(http.StatusFound)
	return false, true, nil
}

type perClientGrant struct {
	auto		GrantHandler
	prompt		GrantHandler
	deny		GrantHandler
	defaultMethod	oauthapi.GrantHandlerType
}

func NewPerClientGrant(prompt GrantHandler, defaultMethod oauthapi.GrantHandlerType) GrantHandler {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &perClientGrant{auto: NewAutoGrant(), prompt: prompt, deny: NewEmptyGrant(), defaultMethod: defaultMethod}
}
func (g *perClientGrant) GrantNeeded(user user.Info, grant *api.Grant, w http.ResponseWriter, req *http.Request) (bool, bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	client, ok := grant.Client.GetUserData().(*oauthapi.OAuthClient)
	if !ok {
		return false, false, errors.New("unrecognized OAuth client type")
	}
	method := client.GrantMethod
	if len(method) == 0 {
		method = g.defaultMethod
	}
	switch method {
	case oauthapi.GrantHandlerAuto:
		return g.auto.GrantNeeded(user, grant, w, req)
	case oauthapi.GrantHandlerPrompt:
		return g.prompt.GrantNeeded(user, grant, w, req)
	case oauthapi.GrantHandlerDeny:
		return g.deny.GrantNeeded(user, grant, w, req)
	default:
		return false, false, fmt.Errorf("OAuth client grant method %q unrecognized", method)
	}
}
