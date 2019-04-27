package fake

import (
	user "github.com/openshift/origin/pkg/user/apis/user"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	testing "k8s.io/client-go/testing"
)

type FakeUserIdentityMappings struct{ Fake *FakeUser }

var useridentitymappingsResource = schema.GroupVersionResource{Group: "user.openshift.io", Version: "", Resource: "useridentitymappings"}
var useridentitymappingsKind = schema.GroupVersionKind{Group: "user.openshift.io", Version: "", Kind: "UserIdentityMapping"}

func (c *FakeUserIdentityMappings) Get(name string, options v1.GetOptions) (result *user.UserIdentityMapping, err error) {
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
	obj, err := c.Fake.Invokes(testing.NewRootGetAction(useridentitymappingsResource, name), &user.UserIdentityMapping{})
	if obj == nil {
		return nil, err
	}
	return obj.(*user.UserIdentityMapping), err
}
func (c *FakeUserIdentityMappings) Create(userIdentityMapping *user.UserIdentityMapping) (result *user.UserIdentityMapping, err error) {
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
	obj, err := c.Fake.Invokes(testing.NewRootCreateAction(useridentitymappingsResource, userIdentityMapping), &user.UserIdentityMapping{})
	if obj == nil {
		return nil, err
	}
	return obj.(*user.UserIdentityMapping), err
}
func (c *FakeUserIdentityMappings) Update(userIdentityMapping *user.UserIdentityMapping) (result *user.UserIdentityMapping, err error) {
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
	obj, err := c.Fake.Invokes(testing.NewRootUpdateAction(useridentitymappingsResource, userIdentityMapping), &user.UserIdentityMapping{})
	if obj == nil {
		return nil, err
	}
	return obj.(*user.UserIdentityMapping), err
}
func (c *FakeUserIdentityMappings) Delete(name string, options *v1.DeleteOptions) error {
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
	_, err := c.Fake.Invokes(testing.NewRootDeleteAction(useridentitymappingsResource, name), &user.UserIdentityMapping{})
	return err
}
