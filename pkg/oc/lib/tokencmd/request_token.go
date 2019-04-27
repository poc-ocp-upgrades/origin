package tokencmd

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"strings"
	"github.com/RangelReale/osincli"
	"k8s.io/klog"
	apierrs "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/sets"
	restclient "k8s.io/client-go/rest"
	"github.com/openshift/origin/pkg/oauth/urls"
	"github.com/openshift/origin/pkg/oauth/util"
)

const (
	csrfTokenHeader						= "X-CSRF-Token"
	oauthMetadataEndpoint					= "/.well-known/oauth-authorization-server"
	openShiftCLIClientID					= "openshift-challenging-client"
	pkce_s256						= "S256"
	token			osincli.AuthorizeRequestType	= "token"
)

type ChallengeHandler interface {
	CanHandle(headers http.Header) bool
	HandleChallenge(requestURL string, headers http.Header) (http.Header, bool, error)
	CompleteChallenge(requestURL string, headers http.Header) error
	Release() error
}
type RequestTokenOptions struct {
	ClientConfig	*restclient.Config
	Handler		ChallengeHandler
	OsinConfig	*osincli.ClientConfig
	Issuer		string
	TokenFlow	bool
}

func RequestToken(clientCfg *restclient.Config, reader io.Reader, defaultUsername string, defaultPassword string) (string, error) {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	return NewRequestTokenOptions(clientCfg, reader, defaultUsername, defaultPassword, false).RequestToken()
}
func NewRequestTokenOptions(clientCfg *restclient.Config, reader io.Reader, defaultUsername string, defaultPassword string, tokenFlow bool) *RequestTokenOptions {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	var handlers []ChallengeHandler
	if GSSAPIEnabled() {
		klog.V(6).Info("GSSAPI Enabled")
		handlers = append(handlers, NewNegotiateChallengeHandler(NewGSSAPINegotiator(defaultUsername)))
	}
	if SSPIEnabled() {
		klog.V(6).Info("SSPI Enabled")
		handlers = append(handlers, NewNegotiateChallengeHandler(NewSSPINegotiator(defaultUsername, defaultPassword, clientCfg.Host, reader)))
	}
	handlers = append(handlers, &BasicChallengeHandler{Host: clientCfg.Host, Reader: reader, Username: defaultUsername, Password: defaultPassword})
	var handler ChallengeHandler
	if len(handlers) == 1 {
		handler = handlers[0]
	} else {
		handler = NewMultiHandler(handlers...)
	}
	return &RequestTokenOptions{ClientConfig: clientCfg, Handler: handler, TokenFlow: tokenFlow}
}
func (o *RequestTokenOptions) SetDefaultOsinConfig() error {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	if o.OsinConfig != nil {
		return fmt.Errorf("osin config is already set to: %#v", *o.OsinConfig)
	}
	rt, err := restclient.TransportFor(o.ClientConfig)
	if err != nil {
		return err
	}
	requestURL := strings.TrimRight(o.ClientConfig.Host, "/") + oauthMetadataEndpoint
	resp, err := request(rt, requestURL, nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("couldn't get %v: unexpected response status %v", requestURL, resp.StatusCode)
	}
	metadata := &util.OauthAuthorizationServerMetadata{}
	if err := json.NewDecoder(resp.Body).Decode(metadata); err != nil {
		return err
	}
	config := &osincli.ClientConfig{ClientId: openShiftCLIClientID, AuthorizeUrl: metadata.AuthorizationEndpoint, TokenUrl: metadata.TokenEndpoint, RedirectUrl: urls.OpenShiftOAuthTokenImplicitURL(metadata.Issuer)}
	if !o.TokenFlow && sets.NewString(metadata.CodeChallengeMethodsSupported...).Has(pkce_s256) {
		if err := osincli.PopulatePKCE(config); err != nil {
			return err
		}
	}
	o.OsinConfig = config
	o.Issuer = metadata.Issuer
	return nil
}
func (o *RequestTokenOptions) RequestToken() (string, error) {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	defer func() {
		if err := o.Handler.Release(); err != nil {
			klog.V(4).Infof("error releasing handler: %v", err)
		}
	}()
	if o.OsinConfig == nil {
		if err := o.SetDefaultOsinConfig(); err != nil {
			return "", err
		}
	}
	rt, err := transportWithSystemRoots(o.Issuer, o.ClientConfig)
	if err != nil {
		return "", err
	}
	client, err := osincli.NewClient(o.OsinConfig)
	if err != nil {
		return "", err
	}
	client.Transport = rt
	authorizeRequest := client.NewAuthorizeRequest(osincli.CODE)
	var oauthTokenFunc func(redirectURL string) (accessToken string, oauthError error)
	if o.TokenFlow {
		authorizeRequest.Type = token
		oauthTokenFunc = oauthTokenFlow
	} else {
		oauthTokenFunc = func(redirectURL string) (accessToken string, oauthError error) {
			return oauthCodeFlow(client, authorizeRequest, redirectURL)
		}
	}
	requestURL := authorizeRequest.GetAuthorizeUrl().String()
	requestHeaders := http.Header{}
	requestedURLSet := sets.NewString()
	requestedURLList := []string{}
	handledChallenge := false
	for {
		resp, err := request(rt, requestURL, requestHeaders)
		if err != nil {
			return "", err
		}
		defer resp.Body.Close()
		if resp.StatusCode == http.StatusUnauthorized {
			if resp.Header.Get("WWW-Authenticate") != "" {
				if !o.Handler.CanHandle(resp.Header) {
					return "", apierrs.NewUnauthorized("unhandled challenge")
				}
				newRequestHeaders, shouldRetry, err := o.Handler.HandleChallenge(requestURL, resp.Header)
				if err != nil {
					return "", err
				}
				if !shouldRetry {
					return "", apierrs.NewUnauthorized("challenger chose not to retry the request")
				}
				handledChallenge = true
				requestedURLSet = sets.NewString()
				requestedURLList = []string{}
				requestHeaders = newRequestHeaders
				continue
			}
			unauthorizedError := apierrs.NewUnauthorized("")
			if details, err := ioutil.ReadAll(resp.Body); err == nil && len(details) > 0 {
				unauthorizedError.ErrStatus.Details = &metav1.StatusDetails{Causes: []metav1.StatusCause{{Message: string(details)}}}
			}
			return "", unauthorizedError
		}
		if handledChallenge {
			if err := o.Handler.CompleteChallenge(requestURL, resp.Header); err != nil {
				return "", err
			}
		}
		if resp.StatusCode == http.StatusFound {
			redirectURL := resp.Header.Get("Location")
			accessToken, err := oauthTokenFunc(redirectURL)
			if err != nil {
				return "", err
			}
			if len(accessToken) > 0 {
				return accessToken, nil
			}
			requestedURLList = append(requestedURLList, redirectURL)
			if !requestedURLSet.Has(redirectURL) {
				requestedURLSet.Insert(redirectURL)
				requestURL = redirectURL
				continue
			}
			return "", apierrs.NewInternalError(fmt.Errorf("redirect loop: %s", strings.Join(requestedURLList, " -> ")))
		}
		return "", apierrs.NewInternalError(fmt.Errorf("unexpected response: %d", resp.StatusCode))
	}
}
func oauthTokenFlow(location string) (string, error) {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	u, err := url.Parse(location)
	if err != nil {
		return "", err
	}
	if oauthErr := oauthErrFromValues(u.Query()); oauthErr != nil {
		return "", oauthErr
	}
	fragment := ""
	if parts := strings.SplitN(location, "#", 2); len(parts) == 2 {
		fragment = parts[1]
	}
	fragmentValues, err := url.ParseQuery(fragment)
	if err != nil {
		return "", err
	}
	return fragmentValues.Get("access_token"), nil
}
func oauthCodeFlow(client *osincli.Client, authorizeRequest *osincli.AuthorizeRequest, location string) (string, error) {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	req, err := http.NewRequest(http.MethodGet, location, nil)
	if err != nil {
		return "", err
	}
	req.ParseForm()
	if oauthErr := oauthErrFromValues(req.Form); oauthErr != nil {
		return "", oauthErr
	}
	if len(req.Form.Get("code")) == 0 {
		return "", nil
	}
	authorizeData, err := authorizeRequest.HandleRequest(req)
	if err != nil {
		return "", osinToOAuthError(err)
	}
	accessRequest := client.NewAccessRequest(osincli.AUTHORIZATION_CODE, authorizeData)
	accessData, err := accessRequest.GetToken()
	if err != nil {
		return "", osinToOAuthError(err)
	}
	return accessData.AccessToken, nil
}
func osinToOAuthError(err error) error {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	if osinErr, ok := err.(*osincli.Error); ok {
		return createOAuthError(osinErr.Id, osinErr.Description)
	}
	return err
}
func oauthErrFromValues(values url.Values) error {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	if errorCode := values.Get("error"); len(errorCode) > 0 {
		errorDescription := values.Get("error_description")
		return createOAuthError(errorCode, errorDescription)
	}
	return nil
}
func createOAuthError(errorCode, errorDescription string) error {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fmt.Errorf("%s %s", errorCode, errorDescription)
}
func request(rt http.RoundTripper, requestURL string, requestHeaders http.Header) (*http.Response, error) {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	req, err := http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		return nil, err
	}
	for k, v := range requestHeaders {
		req.Header[k] = v
	}
	req.Header.Set(csrfTokenHeader, "1")
	return rt.RoundTrip(req)
}
func transportWithSystemRoots(issuer string, clientConfig *restclient.Config) (http.RoundTripper, error) {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	configWithSystemRoots := restclient.CopyConfig(clientConfig)
	configWithSystemRoots.CAFile = ""
	configWithSystemRoots.CAData = nil
	systemRootsRT, err := restclient.TransportFor(configWithSystemRoots)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodHead, issuer, nil)
	if err != nil {
		return nil, err
	}
	_, err = systemRootsRT.RoundTrip(req)
	switch err.(type) {
	case nil:
		klog.V(4).Info("using system roots as no error was encountered")
		return systemRootsRT, nil
	case x509.UnknownAuthorityError, x509.HostnameError, x509.CertificateInvalidError, x509.SystemRootsError, tls.RecordHeaderError, *net.OpError:
		klog.V(4).Infof("falling back to kubeconfig CA due to possible x509 error: %v", err)
		return restclient.TransportFor(clientConfig)
	default:
		switch err {
		case io.EOF, io.ErrUnexpectedEOF, io.ErrNoProgress:
			klog.V(4).Infof("falling back to kubeconfig CA due to possible IO error: %v", err)
			return restclient.TransportFor(clientConfig)
		}
		klog.V(4).Infof("unexpected error during system roots probe: %v", err)
		return nil, err
	}
}
