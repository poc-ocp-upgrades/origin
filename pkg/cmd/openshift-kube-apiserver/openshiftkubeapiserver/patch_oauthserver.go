package openshiftkubeapiserver

import (
	"net/http"
	osinv1 "github.com/openshift/api/osin/v1"
	"github.com/openshift/origin/pkg/oauthserver/oauthserver"
	genericapiserver "k8s.io/apiserver/pkg/server"
)

func NewOAuthServerConfigFromMasterConfig(genericConfig *genericapiserver.Config, oauthConfig *osinv1.OAuthConfig) (*oauthserver.OAuthServerConfig, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	oauthServerConfig, err := oauthserver.NewOAuthServerConfig(*oauthConfig, genericConfig.LoopbackClientConfig)
	if err != nil {
		return nil, err
	}
	oauthServerConfig.GenericConfig.CorsAllowedOriginList = genericConfig.CorsAllowedOriginList
	oauthServerConfig.GenericConfig.SecureServing = genericConfig.SecureServing
	oauthServerConfig.GenericConfig.AuditBackend = genericConfig.AuditBackend
	oauthServerConfig.GenericConfig.AuditPolicyChecker = genericConfig.AuditPolicyChecker
	return oauthServerConfig, nil
}
func NewOAuthServerHandler(genericConfig *genericapiserver.Config, oauthConfig *osinv1.OAuthConfig) (http.Handler, map[string]genericapiserver.PostStartHookFunc, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if oauthConfig == nil {
		return http.NotFoundHandler(), nil, nil
	}
	config, err := NewOAuthServerConfigFromMasterConfig(genericConfig, oauthConfig)
	if err != nil {
		return nil, nil, err
	}
	oauthServer, err := config.Complete().New(genericapiserver.NewEmptyDelegate())
	if err != nil {
		return nil, nil, err
	}
	return oauthServer.GenericAPIServer.PrepareRun().GenericAPIServer.Handler.FullHandlerChain, map[string]genericapiserver.PostStartHookFunc{"oauth.openshift.io-startoauthclientsbootstrapping": config.StartOAuthClientsBootstrapping}, nil
}
