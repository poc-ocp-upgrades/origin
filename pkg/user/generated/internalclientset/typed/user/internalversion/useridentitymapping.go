package internalversion

import (
	user "github.com/openshift/origin/pkg/user/apis/user"
	scheme "github.com/openshift/origin/pkg/user/generated/internalclientset/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	rest "k8s.io/client-go/rest"
)

type UserIdentityMappingsGetter interface {
	UserIdentityMappings() UserIdentityMappingInterface
}
type UserIdentityMappingInterface interface {
	Create(*user.UserIdentityMapping) (*user.UserIdentityMapping, error)
	Update(*user.UserIdentityMapping) (*user.UserIdentityMapping, error)
	Delete(name string, options *v1.DeleteOptions) error
	Get(name string, options v1.GetOptions) (*user.UserIdentityMapping, error)
	UserIdentityMappingExpansion
}
type userIdentityMappings struct{ client rest.Interface }

func newUserIdentityMappings(c *UserClient) *userIdentityMappings {
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
	return &userIdentityMappings{client: c.RESTClient()}
}
func (c *userIdentityMappings) Get(name string, options v1.GetOptions) (result *user.UserIdentityMapping, err error) {
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
	result = &user.UserIdentityMapping{}
	err = c.client.Get().Resource("useridentitymappings").Name(name).VersionedParams(&options, scheme.ParameterCodec).Do().Into(result)
	return
}
func (c *userIdentityMappings) Create(userIdentityMapping *user.UserIdentityMapping) (result *user.UserIdentityMapping, err error) {
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
	result = &user.UserIdentityMapping{}
	err = c.client.Post().Resource("useridentitymappings").Body(userIdentityMapping).Do().Into(result)
	return
}
func (c *userIdentityMappings) Update(userIdentityMapping *user.UserIdentityMapping) (result *user.UserIdentityMapping, err error) {
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
	result = &user.UserIdentityMapping{}
	err = c.client.Put().Resource("useridentitymappings").Name(userIdentityMapping.Name).Body(userIdentityMapping).Do().Into(result)
	return
}
func (c *userIdentityMappings) Delete(name string, options *v1.DeleteOptions) error {
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
	return c.client.Delete().Resource("useridentitymappings").Name(name).Body(options).Do().Error()
}
