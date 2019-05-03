package internalversion

import (
 rest "k8s.io/client-go/rest"
 "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset/scheme"
)

type RbacInterface interface {
 RESTClient() rest.Interface
 ClusterRolesGetter
 ClusterRoleBindingsGetter
 RolesGetter
 RoleBindingsGetter
}
type RbacClient struct{ restClient rest.Interface }

func (c *RbacClient) ClusterRoles() ClusterRoleInterface {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return newClusterRoles(c)
}
func (c *RbacClient) ClusterRoleBindings() ClusterRoleBindingInterface {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return newClusterRoleBindings(c)
}
func (c *RbacClient) Roles(namespace string) RoleInterface {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return newRoles(c, namespace)
}
func (c *RbacClient) RoleBindings(namespace string) RoleBindingInterface {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return newRoleBindings(c, namespace)
}
func NewForConfig(c *rest.Config) (*RbacClient, error) {
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
 return &RbacClient{client}, nil
}
func NewForConfigOrDie(c *rest.Config) *RbacClient {
 _logClusterCodePath()
 defer _logClusterCodePath()
 client, err := NewForConfig(c)
 if err != nil {
  panic(err)
 }
 return client
}
func New(c rest.Interface) *RbacClient {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &RbacClient{c}
}
func setConfigDefaults(config *rest.Config) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 config.APIPath = "/apis"
 if config.UserAgent == "" {
  config.UserAgent = rest.DefaultKubernetesUserAgent()
 }
 if config.GroupVersion == nil || config.GroupVersion.Group != scheme.Scheme.PrioritizedVersionsForGroup("rbac.authorization.k8s.io")[0].Group {
  gv := scheme.Scheme.PrioritizedVersionsForGroup("rbac.authorization.k8s.io")[0]
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
func (c *RbacClient) RESTClient() rest.Interface {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if c == nil {
  return nil
 }
 return c.restClient
}
