package login

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/openshift/origin/pkg/oauthserver"
	"github.com/openshift/origin/pkg/oauthserver/oauth/handlers"
	"github.com/openshift/origin/pkg/oauthserver/prometheus"
	"github.com/openshift/origin/pkg/oauthserver/server/csrf"
	"github.com/openshift/origin/pkg/oauthserver/server/errorpage"
	"github.com/openshift/origin/pkg/oauthserver/server/redirect"
	"html/template"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apiserver/pkg/authentication/authenticator"
	"k8s.io/klog"
	"net/http"
)

const (
	thenParam             = "then"
	csrfParam             = "csrf"
	usernameParam         = "username"
	passwordParam         = "password"
	reasonParam           = "reason"
	errorCodeUserRequired = "user_required"
	errorCodeTokenExpired = "token_expired"
	errorCodeAccessDenied = "access_denied"
)

var errorMessages = map[string]string{errorCodeUserRequired: "Login is required. Please try again.", errorCodeTokenExpired: "Could not check CSRF token. Please try again.", errorCodeAccessDenied: "Invalid login or password. Please try again."}

type PasswordAuthenticator interface {
	authenticator.Password
	handlers.AuthenticationSuccessHandler
}
type LoginFormRenderer interface {
	Render(form LoginForm, w http.ResponseWriter, req *http.Request)
}
type LoginForm struct {
	ProviderName string
	Action       string
	Error        string
	ErrorCode    string
	Names        LoginFormFields
	Values       LoginFormFields
}
type LoginFormFields struct {
	Then     string
	CSRF     string
	Username string
	Password string
}
type Login struct {
	provider string
	csrf     csrf.CSRF
	auth     PasswordAuthenticator
	render   LoginFormRenderer
}

func NewLogin(provider string, csrf csrf.CSRF, auth PasswordAuthenticator, render LoginFormRenderer) *Login {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &Login{provider: provider, csrf: csrf, auth: auth, render: render}
}
func (l *Login) Install(mux oauthserver.Mux, prefix string) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	mux.Handle(prefix, l)
}
func (l *Login) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	switch req.Method {
	case http.MethodGet:
		l.handleLoginForm(w, req)
	case http.MethodPost:
		l.handleLogin(w, req)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
func (l *Login) handleLoginForm(w http.ResponseWriter, req *http.Request) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	uri, err := getBaseURL(req)
	if err != nil {
		utilruntime.HandleError(fmt.Errorf("Unable to generate base URL: %v", err))
		http.Error(w, "Unable to determine URL", http.StatusInternalServerError)
		return
	}
	form := LoginForm{ProviderName: l.provider, Action: uri.String(), Names: LoginFormFields{Then: thenParam, CSRF: csrfParam, Username: usernameParam, Password: passwordParam}}
	if then := req.URL.Query().Get(thenParam); redirect.IsServerRelativeURL(then) {
		form.Values.Then = then
	} else {
		http.Redirect(w, req, "/", http.StatusFound)
		return
	}
	form.ErrorCode = req.URL.Query().Get(reasonParam)
	if len(form.ErrorCode) > 0 {
		if msg, hasMsg := errorMessages[form.ErrorCode]; hasMsg {
			form.Error = msg
		} else {
			form.Error = errorpage.AuthenticationErrorMessage(form.ErrorCode)
		}
	}
	form.Values.CSRF = l.csrf.Generate(w, req)
	l.render.Render(form, w, req)
}
func (l *Login) handleLogin(w http.ResponseWriter, req *http.Request) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if ok := l.csrf.Check(req, req.FormValue(csrfParam)); !ok {
		klog.V(4).Infof("Invalid CSRF token: %s", req.FormValue(csrfParam))
		failed(errorCodeTokenExpired, w, req)
		return
	}
	then := req.FormValue(thenParam)
	if !redirect.IsServerRelativeURL(then) {
		http.Redirect(w, req, "/", http.StatusFound)
		return
	}
	username, password := req.FormValue(usernameParam), req.FormValue(passwordParam)
	if len(username) == 0 {
		failed(errorCodeUserRequired, w, req)
		return
	}
	if len(password) == 0 {
		failed(errorCodeAccessDenied, w, req)
		return
	}
	result := metrics.SuccessResult
	defer func() {
		metrics.RecordFormPasswordAuth(result)
	}()
	authResponse, ok, err := l.auth.AuthenticatePassword(context.TODO(), username, password)
	if err != nil {
		utilruntime.HandleError(fmt.Errorf(`Error authenticating %q with provider %q: %v`, username, l.provider, err))
		failed(errorpage.AuthenticationErrorCode(err), w, req)
		result = metrics.ErrorResult
		return
	}
	if !ok {
		klog.V(4).Infof(`Login with provider %q failed for %q`, l.provider, username)
		failed(errorCodeAccessDenied, w, req)
		result = metrics.FailResult
		return
	}
	klog.V(4).Infof(`Login with provider %q succeeded for %q: %#v`, l.provider, username, authResponse.User)
	l.auth.AuthenticationSucceeded(authResponse.User, then, w, req)
}
func NewLoginFormRenderer(customLoginTemplateFile string) (*loginTemplateRenderer, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	r := &loginTemplateRenderer{}
	if len(customLoginTemplateFile) > 0 {
		customTemplate, err := template.ParseFiles(customLoginTemplateFile)
		if err != nil {
			return nil, err
		}
		r.loginTemplate = customTemplate
	} else {
		r.loginTemplate = defaultLoginTemplate
	}
	return r, nil
}
func ValidateLoginTemplate(templateContent []byte) []error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var allErrs []error
	template, err := template.New("loginTemplateTest").Parse(string(templateContent))
	if err != nil {
		return append(allErrs, err)
	}
	form := LoginForm{Action: "MyAction", Error: "MyError", Names: LoginFormFields{Then: "MyThenName", CSRF: "MyCSRFName", Username: "MyUsernameName", Password: "MyPasswordName"}, Values: LoginFormFields{Then: "MyThenValue", CSRF: "MyCSRFValue", Username: "MyUsernameValue"}}
	var buffer bytes.Buffer
	err = template.Execute(&buffer, form)
	if err != nil {
		return append(allErrs, err)
	}
	output := buffer.Bytes()
	var testFields = map[string]string{"Action": form.Action, "Error": form.Error, "Names.Then": form.Names.Then, "Names.CSRF": form.Values.CSRF, "Names.Username": form.Names.Username, "Names.Password": form.Names.Password, "Values.Then": form.Values.Then, "Values.CSRF": form.Values.CSRF, "Values.Username": form.Values.Username}
	for field, value := range testFields {
		if !bytes.Contains(output, []byte(value)) {
			allErrs = append(allErrs, errors.New(fmt.Sprintf("template is missing parameter {{ .%s }}", field)))
		}
	}
	return allErrs
}

type loginTemplateRenderer struct{ loginTemplate *template.Template }

func (r loginTemplateRenderer) Render(form LoginForm, w http.ResponseWriter, req *http.Request) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	w.Header().Add("Content-Type", "text/html; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := r.loginTemplate.Execute(w, form); err != nil {
		utilruntime.HandleError(fmt.Errorf("unable to render login template: %v", err))
	}
}
