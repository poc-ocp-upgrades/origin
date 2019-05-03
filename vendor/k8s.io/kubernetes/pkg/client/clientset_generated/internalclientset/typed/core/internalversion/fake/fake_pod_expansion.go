package fake

import (
 restclient "k8s.io/client-go/rest"
 core "k8s.io/client-go/testing"
 api "k8s.io/kubernetes/pkg/apis/core"
)

func (c *FakePods) Bind(binding *api.Binding) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 action := core.CreateActionImpl{}
 action.Verb = "create"
 action.Resource = podsResource
 action.Subresource = "bindings"
 action.Object = binding
 _, err := c.Fake.Invokes(action, binding)
 return err
}
func (c *FakePods) GetLogs(name string, opts *api.PodLogOptions) *restclient.Request {
 _logClusterCodePath()
 defer _logClusterCodePath()
 action := core.GenericActionImpl{}
 action.Verb = "get"
 action.Namespace = c.ns
 action.Resource = podsResource
 action.Subresource = "logs"
 action.Value = opts
 _, _ = c.Fake.Invokes(action, &api.Pod{})
 return &restclient.Request{}
}
