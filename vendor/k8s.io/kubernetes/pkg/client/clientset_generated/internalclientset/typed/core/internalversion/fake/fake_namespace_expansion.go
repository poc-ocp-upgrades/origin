package fake

import (
 core "k8s.io/client-go/testing"
 api "k8s.io/kubernetes/pkg/apis/core"
)

func (c *FakeNamespaces) Finalize(namespace *api.Namespace) (*api.Namespace, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 action := core.CreateActionImpl{}
 action.Verb = "create"
 action.Resource = namespacesResource
 action.Subresource = "finalize"
 action.Object = namespace
 obj, err := c.Fake.Invokes(action, namespace)
 if obj == nil {
  return nil, err
 }
 return obj.(*api.Namespace), err
}
