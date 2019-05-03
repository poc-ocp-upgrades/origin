package fake

import (
 rest "k8s.io/client-go/rest"
 testing "k8s.io/client-go/testing"
 internalversion "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset/typed/rbac/internalversion"
)

type FakeRbac struct{ *testing.Fake }

func (c *FakeRbac) ClusterRoles() internalversion.ClusterRoleInterface {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &FakeClusterRoles{c}
}
func (c *FakeRbac) ClusterRoleBindings() internalversion.ClusterRoleBindingInterface {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &FakeClusterRoleBindings{c}
}
func (c *FakeRbac) Roles(namespace string) internalversion.RoleInterface {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &FakeRoles{c, namespace}
}
func (c *FakeRbac) RoleBindings(namespace string) internalversion.RoleBindingInterface {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &FakeRoleBindings{c, namespace}
}
func (c *FakeRbac) RESTClient() rest.Interface {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var ret *rest.RESTClient
 return ret
}
