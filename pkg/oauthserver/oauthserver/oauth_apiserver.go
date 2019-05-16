package oauthserver

import (
	"crypto/sha256"
	"fmt"
	oauthv1 "github.com/openshift/api/oauth/v1"
	osinv1 "github.com/openshift/api/osin/v1"
	oauthclient "github.com/openshift/client-go/oauth/clientset/versioned/typed/oauth/v1"
	routeclient "github.com/openshift/client-go/route/clientset/versioned/typed/route/v1"
	userclient "github.com/openshift/client-go/user/clientset/versioned/typed/user/v1"
	"github.com/openshift/origin/pkg/cmd/server/apis/config/latest"
	"github.com/openshift/origin/pkg/oauth/urls"
	"github.com/openshift/origin/pkg/oauthserver/authenticator/password/bootstrap"
	"github.com/openshift/origin/pkg/oauthserver/config"
	"github.com/openshift/origin/pkg/oauthserver/server/crypto"
	"github.com/openshift/origin/pkg/oauthserver/server/headers"
	"github.com/openshift/origin/pkg/oauthserver/server/session"
	"github.com/openshift/origin/pkg/oauthserver/userregistry/identitymapper"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	genericapiserver "k8s.io/apiserver/pkg/server"
	kclientset "k8s.io/client-go/kubernetes"
	corev1 "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/rest"
	"net/http"
	"net/url"
	"time"
)

var (
	scheme = runtime.NewScheme()
	codecs = serializer.NewCodecFactory(scheme)
)

func init() {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	utilruntime.Must(osinv1.Install(scheme))
}
func NewOAuthServerConfig(oauthConfig osinv1.OAuthConfig, userClientConfig *rest.Config, genericConfig *genericapiserver.RecommendedConfig) (*OAuthServerConfig, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	decoder := codecs.UniversalDecoder(osinv1.GroupVersion)
	for i, idp := range oauthConfig.IdentityProviders {
		if idp.Provider.Object != nil {
			break
		}
		idpObject, err := runtime.Decode(decoder, idp.Provider.Raw)
		if err != nil {
			return nil, err
		}
		oauthConfig.IdentityProviders[i].Provider.Object = idpObject
	}
	if genericConfig == nil {
		genericConfig = genericapiserver.NewRecommendedConfig(codecs)
	}
	genericConfig.LoopbackClientConfig = userClientConfig
	userClient, err := userclient.NewForConfig(userClientConfig)
	if err != nil {
		return nil, err
	}
	oauthClient, err := oauthclient.NewForConfig(userClientConfig)
	if err != nil {
		return nil, err
	}
	eventsClient, err := corev1.NewForConfig(userClientConfig)
	if err != nil {
		return nil, err
	}
	routeClient, err := routeclient.NewForConfig(userClientConfig)
	if err != nil {
		return nil, err
	}
	kubeClient, err := kclientset.NewForConfig(userClientConfig)
	if err != nil {
		return nil, err
	}
	bootstrapUserDataGetter := bootstrap.NewBootstrapUserDataGetter(kubeClient.CoreV1(), kubeClient.CoreV1())
	var sessionAuth session.SessionAuthenticator
	if oauthConfig.SessionConfig != nil {
		secure := isHTTPS(oauthConfig.MasterPublicURL)
		auth, err := buildSessionAuth(secure, oauthConfig.SessionConfig, bootstrapUserDataGetter)
		if err != nil {
			return nil, err
		}
		sessionAuth = auth
		oauthConfig.IdentityProviders = append([]osinv1.IdentityProvider{{Name: bootstrap.BootstrapUser, UseAsChallenger: true, UseAsLogin: true, MappingMethod: string(identitymapper.MappingMethodClaim), Provider: runtime.RawExtension{Object: &config.BootstrapIdentityProvider{}}}}, oauthConfig.IdentityProviders...)
	}
	if len(oauthConfig.IdentityProviders) == 0 {
		oauthConfig.IdentityProviders = []osinv1.IdentityProvider{{Name: "defaultDenyAll", UseAsChallenger: true, UseAsLogin: true, MappingMethod: string(identitymapper.MappingMethodClaim), Provider: runtime.RawExtension{Object: &osinv1.DenyAllPasswordIdentityProvider{}}}}
	}
	ret := &OAuthServerConfig{GenericConfig: genericConfig, ExtraOAuthConfig: ExtraOAuthConfig{Options: oauthConfig, KubeClient: kubeClient, EventsClient: eventsClient.Events(""), RouteClient: routeClient, UserClient: userClient.Users(), IdentityClient: userClient.Identities(), UserIdentityMappingClient: userClient.UserIdentityMappings(), OAuthAccessTokenClient: oauthClient.OAuthAccessTokens(), OAuthAuthorizeTokenClient: oauthClient.OAuthAuthorizeTokens(), OAuthClientClient: oauthClient.OAuthClients(), OAuthClientAuthorizationClient: oauthClient.OAuthClientAuthorizations(), SessionAuth: sessionAuth, BootstrapUserDataGetter: bootstrapUserDataGetter}}
	genericConfig.BuildHandlerChainFunc = ret.buildHandlerChainForOAuth
	return ret, nil
}
func buildSessionAuth(secure bool, config *osinv1.SessionConfig, getter bootstrap.BootstrapUserDataGetter) (session.SessionAuthenticator, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	secrets, err := getSessionSecrets(config.SessionSecretsFile)
	if err != nil {
		return nil, err
	}
	sessionStore := session.NewStore(config.SessionName, secure, secrets...)
	sessionAuthenticator := session.NewAuthenticator(sessionStore, time.Duration(config.SessionMaxAgeSeconds)*time.Second)
	return session.NewBootstrapAuthenticator(sessionAuthenticator, getter, sessionStore), nil
}
func getSessionSecrets(filename string) ([][]byte, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var secrets [][]byte
	if len(filename) != 0 {
		sessionSecrets, err := latest.ReadSessionSecrets(filename)
		if err != nil {
			return nil, fmt.Errorf("error reading sessionSecretsFile %s: %v", filename, err)
		}
		if len(sessionSecrets.Secrets) == 0 {
			return nil, fmt.Errorf("sessionSecretsFile %s contained no secrets", filename)
		}
		for _, s := range sessionSecrets.Secrets {
			secrets = append(secrets, []byte(s.Authentication))
			secrets = append(secrets, []byte(s.Encryption))
		}
	} else {
		const (
			sha256KeyLenBits = sha256.BlockSize * 8
			aes256KeyLenBits = 256
		)
		secrets = append(secrets, crypto.RandomBits(sha256KeyLenBits))
		secrets = append(secrets, crypto.RandomBits(aes256KeyLenBits))
	}
	return secrets, nil
}
func isHTTPS(u string) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	parsedURL, err := url.Parse(u)
	return err == nil && parsedURL.Scheme == "https"
}

type ExtraOAuthConfig struct {
	Options                        osinv1.OAuthConfig
	KubeClient                     kclientset.Interface
	EventsClient                   corev1.EventInterface
	RouteClient                    routeclient.RouteV1Interface
	UserClient                     userclient.UserInterface
	IdentityClient                 userclient.IdentityInterface
	UserIdentityMappingClient      userclient.UserIdentityMappingInterface
	OAuthAccessTokenClient         oauthclient.OAuthAccessTokenInterface
	OAuthAuthorizeTokenClient      oauthclient.OAuthAuthorizeTokenInterface
	OAuthClientClient              oauthclient.OAuthClientInterface
	OAuthClientAuthorizationClient oauthclient.OAuthClientAuthorizationInterface
	SessionAuth                    session.SessionAuthenticator
	BootstrapUserDataGetter        bootstrap.BootstrapUserDataGetter
}
type OAuthServerConfig struct {
	GenericConfig    *genericapiserver.RecommendedConfig
	ExtraOAuthConfig ExtraOAuthConfig
}
type OAuthServer struct {
	GenericAPIServer *genericapiserver.GenericAPIServer
	PublicURL        url.URL
}
type completedOAuthConfig struct {
	GenericConfig    genericapiserver.CompletedConfig
	ExtraOAuthConfig *ExtraOAuthConfig
}
type CompletedOAuthConfig struct{ *completedOAuthConfig }

func (c *OAuthServerConfig) Complete() completedOAuthConfig {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	cfg := completedOAuthConfig{c.GenericConfig.Complete(), &c.ExtraOAuthConfig}
	return cfg
}
func (c completedOAuthConfig) New(delegationTarget genericapiserver.DelegationTarget) (*OAuthServer, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	genericServer, err := c.GenericConfig.New("openshift-oauth", delegationTarget)
	if err != nil {
		return nil, err
	}
	s := &OAuthServer{GenericAPIServer: genericServer}
	return s, nil
}
func (c *OAuthServerConfig) buildHandlerChainForOAuth(startingHandler http.Handler, genericConfig *genericapiserver.Config) http.Handler {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	handler, err := c.WithOAuth(startingHandler)
	if err != nil {
		panic(err)
	}
	handler = headers.WithRestoreAuthorizationHeader(handler)
	handler = genericapiserver.DefaultBuildHandlerChain(handler, genericConfig)
	handler = headers.WithPreserveAuthorizationHeader(handler)
	handler = headers.WithStandardHeaders(handler)
	return handler
}
func (c *OAuthServerConfig) StartOAuthClientsBootstrapping(context genericapiserver.PostStartHookContext) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	go func() {
		_ = wait.PollUntil(1*time.Second, func() (done bool, err error) {
			browserClient := oauthv1.OAuthClient{ObjectMeta: metav1.ObjectMeta{Name: openShiftBrowserClientID}, Secret: crypto.Random256BitsString(), RespondWithChallenges: false, RedirectURIs: []string{urls.OpenShiftOAuthTokenDisplayURL(c.ExtraOAuthConfig.Options.MasterPublicURL)}, GrantMethod: oauthv1.GrantHandlerAuto}
			if err := ensureOAuthClient(browserClient, c.ExtraOAuthConfig.OAuthClientClient, true, true); err != nil {
				utilruntime.HandleError(err)
				return false, nil
			}
			cliClient := oauthv1.OAuthClient{ObjectMeta: metav1.ObjectMeta{Name: openShiftCLIClientID}, Secret: "", RespondWithChallenges: true, RedirectURIs: []string{urls.OpenShiftOAuthTokenImplicitURL(c.ExtraOAuthConfig.Options.MasterPublicURL)}, GrantMethod: oauthv1.GrantHandlerAuto}
			if err := ensureOAuthClient(cliClient, c.ExtraOAuthConfig.OAuthClientClient, false, false); err != nil {
				utilruntime.HandleError(err)
				return false, nil
			}
			return true, nil
		}, context.StopCh)
	}()
	return nil
}
