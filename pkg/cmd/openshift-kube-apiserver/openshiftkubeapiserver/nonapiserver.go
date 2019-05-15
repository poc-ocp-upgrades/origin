package openshiftkubeapiserver

import (
	kubecontrolplanev1 "github.com/openshift/api/kubecontrolplane/v1"
	osinv1 "github.com/openshift/api/osin/v1"
	oauthutil "github.com/openshift/origin/pkg/oauth/util"
	genericapiserver "k8s.io/apiserver/pkg/server"
	genericmux "k8s.io/apiserver/pkg/server/mux"
	"k8s.io/client-go/informers"
	"k8s.io/klog"
	"net/http"
)

func NewOpenshiftNonAPIConfig(generiConfig *genericapiserver.Config, kubeInformers informers.SharedInformerFactory, oauthConfig *osinv1.OAuthConfig, authConfig kubecontrolplanev1.MasterAuthConfig) (*OpenshiftNonAPIConfig, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var err error
	ret := &OpenshiftNonAPIConfig{GenericConfig: &genericapiserver.RecommendedConfig{Config: *generiConfig, SharedInformerFactory: kubeInformers}}
	ret.ExtraConfig.OAuthMetadata, _, err = oauthutil.PrepOauthMetadata(oauthConfig, authConfig.OAuthMetadataFile)
	if err != nil {
		klog.Errorf("Unable to initialize OAuth authorization server metadata: %v", err)
	}
	return ret, nil
}

type NonAPIExtraConfig struct{ OAuthMetadata []byte }
type OpenshiftNonAPIConfig struct {
	GenericConfig *genericapiserver.RecommendedConfig
	ExtraConfig   NonAPIExtraConfig
}
type OpenshiftNonAPIServer struct {
	GenericAPIServer *genericapiserver.GenericAPIServer
}
type completedOpenshiftNonAPIConfig struct {
	GenericConfig genericapiserver.CompletedConfig
	ExtraConfig   *NonAPIExtraConfig
}
type CompletedOpenshiftNonAPIConfig struct {
	*completedOpenshiftNonAPIConfig
}

func (c *OpenshiftNonAPIConfig) Complete() completedOpenshiftNonAPIConfig {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cfg := completedOpenshiftNonAPIConfig{c.GenericConfig.Complete(), &c.ExtraConfig}
	return cfg
}
func (c completedOpenshiftNonAPIConfig) New(delegationTarget genericapiserver.DelegationTarget) (*OpenshiftNonAPIServer, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	genericServer, err := c.GenericConfig.New("openshift-non-api-routes", delegationTarget)
	if err != nil {
		return nil, err
	}
	s := &OpenshiftNonAPIServer{GenericAPIServer: genericServer}
	if len(c.ExtraConfig.OAuthMetadata) > 0 {
		initOAuthAuthorizationServerMetadataRoute(s.GenericAPIServer.Handler.NonGoRestfulMux, c.ExtraConfig)
	}
	return s, nil
}

const (
	oauthMetadataEndpoint = "/.well-known/oauth-authorization-server"
)

func initOAuthAuthorizationServerMetadataRoute(mux *genericmux.PathRecorderMux, ExtraConfig *NonAPIExtraConfig) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	mux.UnlistedHandleFunc(oauthMetadataEndpoint, func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(ExtraConfig.OAuthMetadata)
	})
}
