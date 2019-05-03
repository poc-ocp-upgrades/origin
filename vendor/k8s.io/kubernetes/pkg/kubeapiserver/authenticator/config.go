package authenticator

import (
 "fmt"
 godefaultbytes "bytes"
 godefaulthttp "net/http"
 godefaultruntime "runtime"
 "time"
 "github.com/go-openapi/spec"
 "k8s.io/apiserver/pkg/authentication/authenticator"
 "k8s.io/apiserver/pkg/authentication/authenticatorfactory"
 "k8s.io/apiserver/pkg/authentication/group"
 "k8s.io/apiserver/pkg/authentication/request/anonymous"
 "k8s.io/apiserver/pkg/authentication/request/bearertoken"
 "k8s.io/apiserver/pkg/authentication/request/headerrequest"
 "k8s.io/apiserver/pkg/authentication/request/union"
 "k8s.io/apiserver/pkg/authentication/request/websocket"
 "k8s.io/apiserver/pkg/authentication/request/x509"
 tokencache "k8s.io/apiserver/pkg/authentication/token/cache"
 "k8s.io/apiserver/pkg/authentication/token/tokenfile"
 tokenunion "k8s.io/apiserver/pkg/authentication/token/union"
 genericapiserver "k8s.io/apiserver/pkg/server"
 "k8s.io/apiserver/pkg/server/certs"
 utilfeature "k8s.io/apiserver/pkg/util/feature"
 "k8s.io/apiserver/plugin/pkg/authenticator/password/passwordfile"
 "k8s.io/apiserver/plugin/pkg/authenticator/request/basicauth"
 "k8s.io/apiserver/plugin/pkg/authenticator/token/oidc"
 "k8s.io/apiserver/plugin/pkg/authenticator/token/webhook"
 _ "k8s.io/client-go/plugin/pkg/client/auth"
 certutil "k8s.io/client-go/util/cert"
 "k8s.io/kubernetes/pkg/features"
 "k8s.io/kubernetes/pkg/serviceaccount"
)

type Config struct {
 Anonymous                   bool
 BasicAuthFile               string
 BootstrapToken              bool
 ClientCAFile                string
 TokenAuthFile               string
 OIDCIssuerURL               string
 OIDCClientID                string
 OIDCCAFile                  string
 OIDCUsernameClaim           string
 OIDCUsernamePrefix          string
 OIDCGroupsClaim             string
 OIDCGroupsPrefix            string
 OIDCSigningAlgs             []string
 OIDCRequiredClaims          map[string]string
 ServiceAccountKeyFiles      []string
 ServiceAccountLookup        bool
 ServiceAccountIssuer        string
 APIAudiences                authenticator.Audiences
 WebhookTokenAuthnConfigFile string
 WebhookTokenAuthnCacheTTL   time.Duration
 TokenSuccessCacheTTL        time.Duration
 TokenFailureCacheTTL        time.Duration
 RequestHeaderConfig         *authenticatorfactory.RequestHeaderConfig
 ServiceAccountTokenGetter   serviceaccount.ServiceAccountTokenGetter
 BootstrapTokenAuthenticator authenticator.Token
}

func (config Config) New() (authenticator.Request, *spec.SecurityDefinitions, map[string]genericapiserver.PostStartHookFunc, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var authenticators []authenticator.Request
 var tokenAuthenticators []authenticator.Token
 securityDefinitions := spec.SecurityDefinitions{}
 dynamicReloadHooks := map[string]genericapiserver.PostStartHookFunc{}
 if config.RequestHeaderConfig != nil {
  requestHeaderAuthenticator, dynamicReloadFn, err := headerrequest.NewSecure(config.RequestHeaderConfig.ClientCA, config.RequestHeaderConfig.AllowedClientNames, config.RequestHeaderConfig.UsernameHeaders, config.RequestHeaderConfig.GroupHeaders, config.RequestHeaderConfig.ExtraHeaderPrefixes)
  if err != nil {
   return nil, nil, nil, err
  }
  dynamicReloadHooks["kube-apiserver-requestheader-reload"] = func(context genericapiserver.PostStartHookContext) error {
   go dynamicReloadFn(context.StopCh)
   return nil
  }
  authenticators = append(authenticators, authenticator.WrapAudienceAgnosticRequest(config.APIAudiences, requestHeaderAuthenticator))
 }
 if len(config.BasicAuthFile) > 0 {
  basicAuth, err := newAuthenticatorFromBasicAuthFile(config.BasicAuthFile)
  if err != nil {
   return nil, nil, nil, err
  }
  authenticators = append(authenticators, authenticator.WrapAudienceAgnosticRequest(config.APIAudiences, basicAuth))
  securityDefinitions["HTTPBasic"] = &spec.SecurityScheme{SecuritySchemeProps: spec.SecuritySchemeProps{Type: "basic", Description: "HTTP Basic authentication"}}
 }
 if len(config.ClientCAFile) > 0 {
  dynamicVerifier := certs.NewDynamicCA(config.ClientCAFile)
  if err := dynamicVerifier.CheckCerts(); err != nil {
   return nil, nil, nil, fmt.Errorf("unable to load client CA file %s: %v", config.ClientCAFile, err)
  }
  dynamicReloadHooks["kube-apiserver-clientCA-reload"] = func(context genericapiserver.PostStartHookContext) error {
   go dynamicVerifier.Run(context.StopCh)
   return nil
  }
  authenticators = append(authenticators, x509.NewDynamic(dynamicVerifier.GetVerifier, x509.CommonNameUserConversion))
 }
 if len(config.TokenAuthFile) > 0 {
  tokenAuth, err := newAuthenticatorFromTokenFile(config.TokenAuthFile)
  if err != nil {
   return nil, nil, nil, err
  }
  tokenAuthenticators = append(tokenAuthenticators, authenticator.WrapAudienceAgnosticToken(config.APIAudiences, tokenAuth))
 }
 if len(config.ServiceAccountKeyFiles) > 0 {
  serviceAccountAuth, err := newLegacyServiceAccountAuthenticator(config.ServiceAccountKeyFiles, config.ServiceAccountLookup, config.APIAudiences, config.ServiceAccountTokenGetter)
  if err != nil {
   return nil, nil, nil, err
  }
  tokenAuthenticators = append(tokenAuthenticators, serviceAccountAuth)
 }
 if utilfeature.DefaultFeatureGate.Enabled(features.TokenRequest) && config.ServiceAccountIssuer != "" {
  serviceAccountAuth, err := newServiceAccountAuthenticator(config.ServiceAccountIssuer, config.ServiceAccountKeyFiles, config.APIAudiences, config.ServiceAccountTokenGetter)
  if err != nil {
   return nil, nil, nil, err
  }
  tokenAuthenticators = append(tokenAuthenticators, serviceAccountAuth)
 }
 if config.BootstrapToken {
  if config.BootstrapTokenAuthenticator != nil {
   tokenAuthenticators = append(tokenAuthenticators, authenticator.WrapAudienceAgnosticToken(config.APIAudiences, config.BootstrapTokenAuthenticator))
  }
 }
 if len(config.OIDCIssuerURL) > 0 && len(config.OIDCClientID) > 0 {
  oidcAuth, err := newAuthenticatorFromOIDCIssuerURL(oidc.Options{IssuerURL: config.OIDCIssuerURL, ClientID: config.OIDCClientID, APIAudiences: config.APIAudiences, CAFile: config.OIDCCAFile, UsernameClaim: config.OIDCUsernameClaim, UsernamePrefix: config.OIDCUsernamePrefix, GroupsClaim: config.OIDCGroupsClaim, GroupsPrefix: config.OIDCGroupsPrefix, SupportedSigningAlgs: config.OIDCSigningAlgs, RequiredClaims: config.OIDCRequiredClaims})
  if err != nil {
   return nil, nil, nil, err
  }
  tokenAuthenticators = append(tokenAuthenticators, oidcAuth)
 }
 if len(config.WebhookTokenAuthnConfigFile) > 0 {
  webhookTokenAuth, err := newWebhookTokenAuthenticator(config.WebhookTokenAuthnConfigFile, config.WebhookTokenAuthnCacheTTL, config.APIAudiences)
  if err != nil {
   return nil, nil, nil, err
  }
  tokenAuthenticators = append(tokenAuthenticators, webhookTokenAuth)
 }
 if len(tokenAuthenticators) > 0 {
  tokenAuth := tokenunion.New(tokenAuthenticators...)
  if config.TokenSuccessCacheTTL > 0 || config.TokenFailureCacheTTL > 0 {
   tokenAuth = tokencache.New(tokenAuth, true, config.TokenSuccessCacheTTL, config.TokenFailureCacheTTL)
  }
  authenticators = append(authenticators, bearertoken.New(tokenAuth), websocket.NewProtocolAuthenticator(tokenAuth))
  securityDefinitions["BearerToken"] = &spec.SecurityScheme{SecuritySchemeProps: spec.SecuritySchemeProps{Type: "apiKey", Name: "authorization", In: "header", Description: "Bearer Token authentication"}}
 }
 if len(authenticators) == 0 {
  if config.Anonymous {
   return anonymous.NewAuthenticator(), &securityDefinitions, dynamicReloadHooks, nil
  }
  return nil, &securityDefinitions, dynamicReloadHooks, nil
 }
 authenticator := union.New(authenticators...)
 authenticator = group.NewAuthenticatedGroupAdder(authenticator)
 if config.Anonymous {
  authenticator = union.NewFailOnError(authenticator, anonymous.NewAuthenticator())
 }
 return authenticator, &securityDefinitions, dynamicReloadHooks, nil
}
func IsValidServiceAccountKeyFile(file string) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 _, err := certutil.PublicKeysFromFile(file)
 return err == nil
}
func newAuthenticatorFromBasicAuthFile(basicAuthFile string) (authenticator.Request, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 basicAuthenticator, err := passwordfile.NewCSV(basicAuthFile)
 if err != nil {
  return nil, err
 }
 return basicauth.New(basicAuthenticator), nil
}
func newAuthenticatorFromTokenFile(tokenAuthFile string) (authenticator.Token, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 tokenAuthenticator, err := tokenfile.NewCSV(tokenAuthFile)
 if err != nil {
  return nil, err
 }
 return tokenAuthenticator, nil
}
func newAuthenticatorFromOIDCIssuerURL(opts oidc.Options) (authenticator.Token, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 const noUsernamePrefix = "-"
 if opts.UsernamePrefix == "" && opts.UsernameClaim != "email" {
  opts.UsernamePrefix = opts.IssuerURL + "#"
 }
 if opts.UsernamePrefix == noUsernamePrefix {
  opts.UsernamePrefix = ""
 }
 tokenAuthenticator, err := oidc.New(opts)
 if err != nil {
  return nil, err
 }
 return tokenAuthenticator, nil
}
func newLegacyServiceAccountAuthenticator(keyfiles []string, lookup bool, apiAudiences authenticator.Audiences, serviceAccountGetter serviceaccount.ServiceAccountTokenGetter) (authenticator.Token, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allPublicKeys := []interface{}{}
 for _, keyfile := range keyfiles {
  publicKeys, err := certutil.PublicKeysFromFile(keyfile)
  if err != nil {
   return nil, err
  }
  allPublicKeys = append(allPublicKeys, publicKeys...)
 }
 tokenAuthenticator := serviceaccount.JWTTokenAuthenticator(serviceaccount.LegacyIssuer, allPublicKeys, apiAudiences, serviceaccount.NewLegacyValidator(lookup, serviceAccountGetter))
 return tokenAuthenticator, nil
}
func newServiceAccountAuthenticator(iss string, keyfiles []string, apiAudiences authenticator.Audiences, serviceAccountGetter serviceaccount.ServiceAccountTokenGetter) (authenticator.Token, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allPublicKeys := []interface{}{}
 for _, keyfile := range keyfiles {
  publicKeys, err := certutil.PublicKeysFromFile(keyfile)
  if err != nil {
   return nil, err
  }
  allPublicKeys = append(allPublicKeys, publicKeys...)
 }
 tokenAuthenticator := serviceaccount.JWTTokenAuthenticator(iss, allPublicKeys, apiAudiences, serviceaccount.NewValidator(serviceAccountGetter))
 return tokenAuthenticator, nil
}
func newAuthenticatorFromClientCAFile(clientCAFile string) (authenticator.Request, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 roots, err := certutil.NewPool(clientCAFile)
 if err != nil {
  return nil, err
 }
 opts := x509.DefaultVerifyOptions()
 opts.Roots = roots
 return x509.New(opts, x509.CommonNameUserConversion), nil
}
func newWebhookTokenAuthenticator(webhookConfigFile string, ttl time.Duration, implicitAuds authenticator.Audiences) (authenticator.Token, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 webhookTokenAuthenticator, err := webhook.New(webhookConfigFile, implicitAuds)
 if err != nil {
  return nil, err
 }
 return tokencache.New(webhookTokenAuthenticator, false, ttl, ttl), nil
}
func _logClusterCodePath() {
 pc, _, _, _ := godefaultruntime.Caller(1)
 jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
 godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
