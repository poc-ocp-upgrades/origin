package handlers

import (
	"fmt"
	oauthapi "github.com/openshift/api/oauth/v1"
	authapi "github.com/openshift/origin/pkg/oauthserver/api"
	kerrors "k8s.io/apimachinery/pkg/util/errors"
	"k8s.io/apiserver/pkg/endpoints/request"
	"net/http"
	"regexp"
	"strings"
)

type unionAuthenticationHandler struct {
	challengers      map[string]AuthenticationChallenger
	redirectors      *AuthenticationRedirectors
	errorHandler     AuthenticationErrorHandler
	selectionHandler AuthenticationSelectionHandler
}

func NewUnionAuthenticationHandler(passedChallengers map[string]AuthenticationChallenger, passedRedirectors *AuthenticationRedirectors, errorHandler AuthenticationErrorHandler, selectionHandler AuthenticationSelectionHandler) AuthenticationHandler {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	challengers := passedChallengers
	if challengers == nil {
		challengers = make(map[string]AuthenticationChallenger, 1)
	}
	redirectors := passedRedirectors
	if redirectors == nil {
		redirectors = new(AuthenticationRedirectors)
	}
	return &unionAuthenticationHandler{challengers: challengers, redirectors: redirectors, errorHandler: errorHandler, selectionHandler: selectionHandler}
}

const (
	WarningHeaderMiscCode        = "199"
	WarningHeaderOpenShiftSource = "Origin"
	warningHeaderCodeIndex       = 1
	warningHeaderAgentIndex      = 2
	warningHeaderTextIndex       = 3
	warningHeaderDateIndex       = 4
	useRedirectParam             = "idp"
)

var (
	warningRegex = regexp.MustCompile(strings.Join([]string{`^`, `([0-9]{3})`, ` `, `([^ ]+)`, ` `, `"((?:[^"\\]|\\.)*)"`, `(?: "([^"]+)")?`, `$`}, ""))
)

func (authHandler *unionAuthenticationHandler) AuthenticationNeeded(apiClient authapi.Client, w http.ResponseWriter, req *http.Request) (bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	client, ok := apiClient.GetUserData().(*oauthapi.OAuthClient)
	if !ok {
		return false, fmt.Errorf("apiClient data was not an oauthapi.OAuthClient")
	}
	if client.RespondWithChallenges {
		errors := []error{}
		headers := http.Header(make(map[string][]string))
		for _, challengingHandler := range authHandler.challengers {
			currHeaders, err := challengingHandler.AuthenticationChallenge(req)
			if err != nil {
				errors = append(errors, err)
				continue
			}
			mergeHeaders(headers, currHeaders)
		}
		if len(headers) > 0 {
			mergeHeaders(w.Header(), headers)
			redirectHeader := w.Header().Get("Location")
			redirectHeaders := w.Header()[http.CanonicalHeaderKey("Location")]
			challengeHeader := w.Header().Get("WWW-Authenticate")
			switch {
			case len(redirectHeader) > 0 && len(challengeHeader) > 0:
				errors = append(errors, fmt.Errorf("redirect header (Location: %s) and challenge header (WWW-Authenticate: %s) cannot both be set", redirectHeader, challengeHeader))
				return false, kerrors.NewAggregate(errors)
			case len(redirectHeaders) > 1:
				errors = append(errors, fmt.Errorf("cannot set multiple redirect headers: %s", strings.Join(redirectHeaders, ", ")))
				return false, kerrors.NewAggregate(errors)
			case len(redirectHeader) > 0:
				w.WriteHeader(http.StatusFound)
			default:
				w.WriteHeader(http.StatusUnauthorized)
				ctx := req.Context()
				if !ok {
					return false, fmt.Errorf("no context found for request to audit")
				}
				ev := request.AuditEventFrom(ctx)
				if ev != nil {
					ev.ResponseStatus.Message = getAuthMethods(req)
				}
			}
			if warnings, hasWarnings := w.Header()[http.CanonicalHeaderKey("Warning")]; hasWarnings {
				for _, warning := range warnings {
					warningParts := warningRegex.FindStringSubmatch(warning)
					if len(warningParts) != 0 && warningParts[warningHeaderCodeIndex] == WarningHeaderMiscCode {
						fmt.Fprintln(w, warningParts[warningHeaderTextIndex])
					}
				}
			}
			return true, nil
		}
		return false, kerrors.NewAggregate(errors)
	}
	redirectHandlerName := req.URL.Query().Get(useRedirectParam)
	if len(redirectHandlerName) > 0 {
		redirectHandler, ok := authHandler.redirectors.Get(redirectHandlerName)
		if !ok {
			return false, fmt.Errorf("Unable to locate redirect handler: %v", redirectHandlerName)
		}
		err := redirectHandler.AuthenticationRedirect(w, req)
		if err != nil {
			return authHandler.errorHandler.AuthenticationError(err, w, req)
		}
		return true, nil
	}
	if authHandler.selectionHandler != nil {
		providers := []authapi.ProviderInfo{}
		for _, name := range authHandler.redirectors.GetNames() {
			u := *req.URL
			q := u.Query()
			q.Set(useRedirectParam, name)
			u.RawQuery = q.Encode()
			providerInfo := authapi.ProviderInfo{Name: name, URL: u.String()}
			providers = append(providers, providerInfo)
		}
		selectedProvider, handled, err := authHandler.selectionHandler.SelectAuthentication(providers, w, req)
		if err != nil {
			return authHandler.errorHandler.AuthenticationError(err, w, req)
		}
		if handled {
			return handled, nil
		}
		if selectedProvider != nil {
			redirectHandler, ok := authHandler.redirectors.Get(selectedProvider.Name)
			if !ok {
				return false, fmt.Errorf("Unable to locate redirect handler: %v", selectedProvider.Name)
			}
			err := redirectHandler.AuthenticationRedirect(w, req)
			if err != nil {
				return authHandler.errorHandler.AuthenticationError(err, w, req)
			}
			return true, nil
		}
	}
	if authHandler.redirectors.Count() == 1 {
		redirectHandler, ok := authHandler.redirectors.Get(authHandler.redirectors.GetNames()[0])
		if !ok {
			return authHandler.errorHandler.AuthenticationError(fmt.Errorf("No valid redirectors"), w, req)
		}
		err := redirectHandler.AuthenticationRedirect(w, req)
		if err != nil {
			return authHandler.errorHandler.AuthenticationError(err, w, req)
		}
		return true, nil
	} else if authHandler.redirectors.Count() > 1 {
		return false, fmt.Errorf("Too many potential redirect handlers: %v", authHandler.redirectors)
	}
	return false, nil
}
func mergeHeaders(dest http.Header, toAdd http.Header) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	for key, values := range toAdd {
		for _, value := range values {
			dest.Add(key, value)
		}
	}
}
func getAuthMethods(req *http.Request) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	authMethods := []string{}
	if _, _, ok := req.BasicAuth(); ok {
		authMethods = append(authMethods, "basic")
	}
	auth := strings.TrimSpace(req.Header.Get("Authorization"))
	parts := strings.Split(auth, " ")
	if len(parts) > 1 && strings.ToLower(parts[0]) == "bearer" {
		authMethods = append(authMethods, "bearer")
	}
	token := strings.TrimSpace(req.URL.Query().Get("access_token"))
	if len(token) > 0 {
		authMethods = append(authMethods, "access_token")
	}
	if req.TLS != nil && len(req.TLS.PeerCertificates) > 0 {
		authMethods = append(authMethods, "x509")
	}
	if len(authMethods) > 0 {
		return fmt.Sprintf("Authentication failed, attempted: %s", strings.Join(authMethods, ", "))
	}
	return "Authentication failed, no credentials provided"
}
