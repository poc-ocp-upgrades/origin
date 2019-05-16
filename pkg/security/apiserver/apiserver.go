package apiserver

import (
	goformat "fmt"
	securityapiv1 "github.com/openshift/api/security/v1"
	securityv1informer "github.com/openshift/client-go/security/informers/externalversions"
	"github.com/openshift/origin/pkg/security/apiserver/registry/podsecuritypolicyreview"
	"github.com/openshift/origin/pkg/security/apiserver/registry/podsecuritypolicyselfsubjectreview"
	"github.com/openshift/origin/pkg/security/apiserver/registry/podsecuritypolicysubjectreview"
	"github.com/openshift/origin/pkg/security/apiserver/registry/rangeallocations"
	sccstorage "github.com/openshift/origin/pkg/security/apiserver/registry/securitycontextconstraints/etcd"
	oscc "github.com/openshift/origin/pkg/security/apiserver/securitycontextconstraints"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/apiserver/pkg/authorization/authorizer"
	"k8s.io/apiserver/pkg/registry/rest"
	genericapiserver "k8s.io/apiserver/pkg/server"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	restclient "k8s.io/client-go/rest"
	goos "os"
	godefaultruntime "runtime"
	"sync"
	gotime "time"
)

type ExtraConfig struct {
	KubeAPIServerClientConfig *restclient.Config
	SecurityInformers         securityv1informer.SharedInformerFactory
	KubeInformers             informers.SharedInformerFactory
	Authorizer                authorizer.Authorizer
	Scheme                    *runtime.Scheme
	Codecs                    serializer.CodecFactory
	makeV1Storage             sync.Once
	v1Storage                 map[string]rest.Storage
	v1StorageErr              error
}
type SecurityAPIServerConfig struct {
	GenericConfig *genericapiserver.RecommendedConfig
	ExtraConfig   ExtraConfig
}
type SecurityAPIServer struct {
	GenericAPIServer *genericapiserver.GenericAPIServer
}
type completedConfig struct {
	GenericConfig genericapiserver.CompletedConfig
	ExtraConfig   *ExtraConfig
}
type CompletedConfig struct{ *completedConfig }

func (c *SecurityAPIServerConfig) Complete() completedConfig {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	cfg := completedConfig{c.GenericConfig.Complete(), &c.ExtraConfig}
	return cfg
}
func (c completedConfig) New(delegationTarget genericapiserver.DelegationTarget) (*SecurityAPIServer, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	genericServer, err := c.GenericConfig.New("security.openshift.io-apiserver", delegationTarget)
	if err != nil {
		return nil, err
	}
	s := &SecurityAPIServer{GenericAPIServer: genericServer}
	v1Storage, err := c.V1RESTStorage()
	if err != nil {
		return nil, err
	}
	apiGroupInfo := genericapiserver.NewDefaultAPIGroupInfo(securityapiv1.GroupName, c.ExtraConfig.Scheme, metav1.ParameterCodec, c.ExtraConfig.Codecs)
	apiGroupInfo.VersionedResourcesStorageMap[securityapiv1.SchemeGroupVersion.Version] = v1Storage
	if err := s.GenericAPIServer.InstallAPIGroup(&apiGroupInfo); err != nil {
		return nil, err
	}
	return s, nil
}
func (c *completedConfig) V1RESTStorage() (map[string]rest.Storage, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	c.ExtraConfig.makeV1Storage.Do(func() {
		c.ExtraConfig.v1Storage, c.ExtraConfig.v1StorageErr = c.newV1RESTStorage()
	})
	return c.ExtraConfig.v1Storage, c.ExtraConfig.v1StorageErr
}
func (c *completedConfig) newV1RESTStorage() (map[string]rest.Storage, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	kubeClient, err := kubernetes.NewForConfig(c.ExtraConfig.KubeAPIServerClientConfig)
	if err != nil {
		return nil, err
	}
	sccStorage := sccstorage.NewREST()
	sccMatcher := oscc.NewDefaultSCCMatcher(c.ExtraConfig.SecurityInformers.Security().V1().SecurityContextConstraints().Lister(), c.ExtraConfig.Authorizer)
	podSecurityPolicyReviewStorage := podsecuritypolicyreview.NewREST(sccMatcher, c.ExtraConfig.KubeInformers.Core().V1().ServiceAccounts().Lister(), kubeClient)
	podSecurityPolicySubjectStorage := podsecuritypolicysubjectreview.NewREST(sccMatcher, kubeClient)
	podSecurityPolicySelfSubjectReviewStorage := podsecuritypolicyselfsubjectreview.NewREST(sccMatcher, kubeClient)
	uidRangeStorage := rangeallocations.NewREST(c.GenericConfig.RESTOptionsGetter)
	v1Storage := map[string]rest.Storage{}
	v1Storage["securityContextConstraints"] = sccStorage
	v1Storage["podSecurityPolicyReviews"] = podSecurityPolicyReviewStorage
	v1Storage["podSecurityPolicySubjectReviews"] = podSecurityPolicySubjectStorage
	v1Storage["podSecurityPolicySelfSubjectReviews"] = podSecurityPolicySelfSubjectReviewStorage
	v1Storage["rangeallocations"] = uidRangeStorage
	return v1Storage, nil
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
