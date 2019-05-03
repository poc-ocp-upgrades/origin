package fake

import (
 "k8s.io/apimachinery/pkg/runtime/schema"
 core "k8s.io/client-go/testing"
 policy "k8s.io/kubernetes/pkg/apis/policy"
)

func (c *FakeEvictions) Evict(eviction *policy.Eviction) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 action := core.GetActionImpl{}
 action.Verb = "post"
 action.Namespace = c.ns
 action.Resource = schema.GroupVersionResource{Group: "", Version: "", Resource: "pods"}
 action.Subresource = "eviction"
 _, err := c.Fake.Invokes(action, eviction)
 return err
}
