package impersonatingclient

import (
	goformat "fmt"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apiserver/pkg/authentication/user"
	kclientset "k8s.io/client-go/kubernetes"
	restclient "k8s.io/client-go/rest"
	"k8s.io/client-go/transport"
	"k8s.io/client-go/util/flowcontrol"
	"net/http"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

func NewImpersonatingConfig(user user.Info, config restclient.Config) restclient.Config {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	oldWrapTransport := config.WrapTransport
	if oldWrapTransport == nil {
		oldWrapTransport = func(rt http.RoundTripper) http.RoundTripper {
			return rt
		}
	}
	newConfig := transport.ImpersonationConfig{UserName: user.GetName(), Groups: user.GetGroups(), Extra: user.GetExtra()}
	config.WrapTransport = func(rt http.RoundTripper) http.RoundTripper {
		return transport.NewImpersonatingRoundTripper(newConfig, oldWrapTransport(rt))
	}
	return config
}
func NewImpersonatingKubernetesClientset(user user.Info, config restclient.Config) (kclientset.Interface, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	impersonatingConfig := NewImpersonatingConfig(user, config)
	return kclientset.NewForConfig(&impersonatingConfig)
}

type impersonatingRESTClient struct {
	user     user.Info
	delegate restclient.Interface
}

func NewImpersonatingRESTClient(user user.Info, client restclient.Interface) restclient.Interface {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &impersonatingRESTClient{user: user, delegate: client}
}
func (c *impersonatingRESTClient) impersonate(req *restclient.Request) *restclient.Request {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	req.SetHeader(transport.ImpersonateUserHeader, c.user.GetName())
	req.SetHeader(transport.ImpersonateGroupHeader, c.user.GetGroups()...)
	for k, vv := range c.user.GetExtra() {
		req.SetHeader(transport.ImpersonateUserExtraHeaderPrefix+k, vv...)
	}
	return req
}
func (c *impersonatingRESTClient) Verb(verb string) *restclient.Request {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return c.impersonate(c.delegate.Verb(verb))
}
func (c *impersonatingRESTClient) Post() *restclient.Request {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return c.impersonate(c.delegate.Post())
}
func (c *impersonatingRESTClient) Put() *restclient.Request {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return c.impersonate(c.delegate.Put())
}
func (c *impersonatingRESTClient) Patch(pt types.PatchType) *restclient.Request {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return c.impersonate(c.delegate.Patch(pt))
}
func (c *impersonatingRESTClient) Get() *restclient.Request {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return c.impersonate(c.delegate.Get())
}
func (c *impersonatingRESTClient) Delete() *restclient.Request {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return c.impersonate(c.delegate.Delete())
}
func (c *impersonatingRESTClient) APIVersion() schema.GroupVersion {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return c.delegate.APIVersion()
}
func (c *impersonatingRESTClient) GetRateLimiter() flowcontrol.RateLimiter {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return c.delegate.GetRateLimiter()
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
