package apiserver

import (
	"fmt"
	goformat "fmt"
	networkapiv1 "github.com/openshift/api/network/v1"
	clusternetworketcd "github.com/openshift/origin/pkg/network/apiserver/registry/clusternetwork/etcd"
	egressnetworkpolicyetcd "github.com/openshift/origin/pkg/network/apiserver/registry/egressnetworkpolicy/etcd"
	hostsubnetetcd "github.com/openshift/origin/pkg/network/apiserver/registry/hostsubnet/etcd"
	netnamespaceetcd "github.com/openshift/origin/pkg/network/apiserver/registry/netnamespace/etcd"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/apiserver/pkg/registry/rest"
	genericapiserver "k8s.io/apiserver/pkg/server"
	goos "os"
	godefaultruntime "runtime"
	"sync"
	gotime "time"
)

type ExtraConfig struct {
	Scheme        *runtime.Scheme
	Codecs        serializer.CodecFactory
	makeV1Storage sync.Once
	v1Storage     map[string]rest.Storage
	v1StorageErr  error
}
type NetworkAPIServerConfig struct {
	GenericConfig *genericapiserver.RecommendedConfig
	ExtraConfig   ExtraConfig
}
type NetworkAPIServer struct {
	GenericAPIServer *genericapiserver.GenericAPIServer
}
type completedConfig struct {
	GenericConfig genericapiserver.CompletedConfig
	ExtraConfig   *ExtraConfig
}
type CompletedConfig struct{ *completedConfig }

func (c *NetworkAPIServerConfig) Complete() completedConfig {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	cfg := completedConfig{c.GenericConfig.Complete(), &c.ExtraConfig}
	return cfg
}
func (c completedConfig) New(delegationTarget genericapiserver.DelegationTarget) (*NetworkAPIServer, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	genericServer, err := c.GenericConfig.New("network.openshift.io-apiserver", delegationTarget)
	if err != nil {
		return nil, err
	}
	s := &NetworkAPIServer{GenericAPIServer: genericServer}
	v1Storage, err := c.V1RESTStorage()
	if err != nil {
		return nil, err
	}
	apiGroupInfo := genericapiserver.NewDefaultAPIGroupInfo(networkapiv1.GroupName, c.ExtraConfig.Scheme, metav1.ParameterCodec, c.ExtraConfig.Codecs)
	apiGroupInfo.VersionedResourcesStorageMap[networkapiv1.SchemeGroupVersion.Version] = v1Storage
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
	hostSubnetStorage, err := hostsubnetetcd.NewREST(c.GenericConfig.RESTOptionsGetter)
	if err != nil {
		return nil, fmt.Errorf("error building REST storage: %v", err)
	}
	netNamespaceStorage, err := netnamespaceetcd.NewREST(c.GenericConfig.RESTOptionsGetter)
	if err != nil {
		return nil, fmt.Errorf("error building REST storage: %v", err)
	}
	clusterNetworkStorage, err := clusternetworketcd.NewREST(c.GenericConfig.RESTOptionsGetter)
	if err != nil {
		return nil, fmt.Errorf("error building REST storage: %v", err)
	}
	egressNetworkPolicyStorage, err := egressnetworkpolicyetcd.NewREST(c.GenericConfig.RESTOptionsGetter)
	if err != nil {
		return nil, fmt.Errorf("error building REST storage: %v", err)
	}
	v1Storage := map[string]rest.Storage{}
	v1Storage["hostSubnets"] = hostSubnetStorage
	v1Storage["netNamespaces"] = netNamespaceStorage
	v1Storage["clusterNetworks"] = clusterNetworkStorage
	v1Storage["egressNetworkPolicies"] = egressNetworkPolicyStorage
	return v1Storage, nil
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
