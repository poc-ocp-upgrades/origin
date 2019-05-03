package internalversion

import (
	"github.com/openshift/origin/pkg/user/generated/internalclientset/scheme"
	rest "k8s.io/client-go/rest"
)

type UserInterface interface {
	RESTClient() rest.Interface
	GroupsGetter
	IdentitiesGetter
	UsersGetter
	UserIdentityMappingsGetter
}
type UserClient struct{ restClient rest.Interface }

func (c *UserClient) Groups() GroupInterface {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return newGroups(c)
}
func (c *UserClient) Identities() IdentityInterface {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return newIdentities(c)
}
func (c *UserClient) Users() UserResourceInterface {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return newUsers(c)
}
func (c *UserClient) UserIdentityMappings() UserIdentityMappingInterface {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return newUserIdentityMappings(c)
}
func NewForConfig(c *rest.Config) (*UserClient, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &UserClient{client}, nil
}
func NewForConfigOrDie(c *rest.Config) *UserClient {
	_logClusterCodePath()
	defer _logClusterCodePath()
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}
func New(c rest.Interface) *UserClient {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &UserClient{c}
}
func setConfigDefaults(config *rest.Config) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	config.APIPath = "/apis"
	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}
	if config.GroupVersion == nil || config.GroupVersion.Group != scheme.Scheme.PrioritizedVersionsForGroup("user.openshift.io")[0].Group {
		gv := scheme.Scheme.PrioritizedVersionsForGroup("user.openshift.io")[0]
		config.GroupVersion = &gv
	}
	config.NegotiatedSerializer = scheme.Codecs
	if config.QPS == 0 {
		config.QPS = 5
	}
	if config.Burst == 0 {
		config.Burst = 10
	}
	return nil
}
func (c *UserClient) RESTClient() rest.Interface {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if c == nil {
		return nil
	}
	return c.restClient
}
