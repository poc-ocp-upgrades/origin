package openshiftkubeapiserver

import (
	"net/http"
	"strings"
	"k8s.io/klog"
	genericapiserver "k8s.io/apiserver/pkg/server"
	kubecontrolplanev1 "github.com/openshift/api/kubecontrolplane/v1"
	osinv1 "github.com/openshift/api/osin/v1"
	"github.com/openshift/origin/pkg/cmd/openshift-apiserver/openshiftapiserver/configprocessing"
	"github.com/openshift/origin/pkg/oauth/urls"
	"github.com/openshift/origin/pkg/oauth/util"
	"github.com/openshift/origin/pkg/util/httprequest"
)

const (
	openShiftOAuthAPIPrefix		= "/oauth"
	openShiftLoginPrefix		= "/login"
	openShiftLogoutPrefix		= "/logout"
	openShiftOAuthCallbackPrefix	= "/oauth2callback"
)

func BuildHandlerChain(genericConfig *genericapiserver.Config, oauthConfig *osinv1.OAuthConfig, authConfig kubecontrolplanev1.MasterAuthConfig, userAgentMatchingConfig kubecontrolplanev1.UserAgentMatchingConfig, consolePublicURL string) (func(apiHandler http.Handler, kc *genericapiserver.Config) http.Handler, map[string]genericapiserver.PostStartHookFunc, error) {
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
	if oauthMetadataFile := authConfig.OAuthMetadataFile; len(oauthMetadataFile) > 0 {
		if _, _, err := util.LoadOAuthMetadataFile(oauthMetadataFile); err == nil {
			oauthConfig = nil
		}
	}
	extraPostStartHooks := map[string]genericapiserver.PostStartHookFunc{}
	var oauthServerHandler http.Handler
	if oauthConfig != nil {
		var newPostStartHooks map[string]genericapiserver.PostStartHookFunc
		var err error
		oauthServerHandler, newPostStartHooks, err = NewOAuthServerHandler(genericConfig, oauthConfig)
		if err != nil {
			return nil, nil, err
		}
		for name, fn := range newPostStartHooks {
			extraPostStartHooks[name] = fn
		}
	}
	return func(apiHandler http.Handler, genericConfig *genericapiserver.Config) http.Handler {
		handler := versionSkewFilter(apiHandler, userAgentMatchingConfig)
		handler = genericapiserver.DefaultBuildHandlerChain(handler, genericConfig)
		handler = translateLegacyScopeImpersonation(handler)
		handler = configprocessing.WithCacheControl(handler, "no-store")
		handler = withConsoleRedirect(handler, consolePublicURL)
		if oauthConfig != nil {
			handler = withOAuthRedirection(oauthConfig, handler, oauthServerHandler)
		}
		return handler
	}, extraPostStartHooks, nil
}
func withOAuthRedirection(oauthConfig *osinv1.OAuthConfig, handler, oauthServerHandler http.Handler) http.Handler {
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
	if oauthConfig == nil {
		return handler
	}
	klog.Infof("Starting OAuth2 API at %s", urls.OpenShiftOAuthAPIPrefix)
	return WithPatternPrefixHandler(handler, oauthServerHandler, openShiftOAuthAPIPrefix, openShiftLoginPrefix, openShiftLogoutPrefix, openShiftOAuthCallbackPrefix)
}
func WithPatternPrefixHandler(handler http.Handler, patternHandler http.Handler, prefixes ...string) http.Handler {
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
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		for _, p := range prefixes {
			if strings.HasPrefix(req.URL.Path, p) {
				patternHandler.ServeHTTP(w, req)
				return
			}
		}
		handler.ServeHTTP(w, req)
	})
}
func withConsoleRedirect(handler http.Handler, consolePublicURL string) http.Handler {
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
	if len(consolePublicURL) == 0 {
		return handler
	}
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if strings.HasPrefix(req.URL.Path, "/console") || (req.URL.Path == "/" && httprequest.PrefersHTML(req)) {
			http.Redirect(w, req, consolePublicURL, http.StatusFound)
			return
		}
		handler.ServeHTTP(w, req)
	})
}
