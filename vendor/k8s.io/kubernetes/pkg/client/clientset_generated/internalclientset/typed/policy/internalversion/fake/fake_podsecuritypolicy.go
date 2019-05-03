package fake

import (
 v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
 labels "k8s.io/apimachinery/pkg/labels"
 schema "k8s.io/apimachinery/pkg/runtime/schema"
 types "k8s.io/apimachinery/pkg/types"
 watch "k8s.io/apimachinery/pkg/watch"
 testing "k8s.io/client-go/testing"
 policy "k8s.io/kubernetes/pkg/apis/policy"
)

type FakePodSecurityPolicies struct{ Fake *FakePolicy }

var podsecuritypoliciesResource = schema.GroupVersionResource{Group: "policy", Version: "", Resource: "podsecuritypolicies"}
var podsecuritypoliciesKind = schema.GroupVersionKind{Group: "policy", Version: "", Kind: "PodSecurityPolicy"}

func (c *FakePodSecurityPolicies) Get(name string, options v1.GetOptions) (result *policy.PodSecurityPolicy, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewRootGetAction(podsecuritypoliciesResource, name), &policy.PodSecurityPolicy{})
 if obj == nil {
  return nil, err
 }
 return obj.(*policy.PodSecurityPolicy), err
}
func (c *FakePodSecurityPolicies) List(opts v1.ListOptions) (result *policy.PodSecurityPolicyList, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewRootListAction(podsecuritypoliciesResource, podsecuritypoliciesKind, opts), &policy.PodSecurityPolicyList{})
 if obj == nil {
  return nil, err
 }
 label, _, _ := testing.ExtractFromListOptions(opts)
 if label == nil {
  label = labels.Everything()
 }
 list := &policy.PodSecurityPolicyList{ListMeta: obj.(*policy.PodSecurityPolicyList).ListMeta}
 for _, item := range obj.(*policy.PodSecurityPolicyList).Items {
  if label.Matches(labels.Set(item.Labels)) {
   list.Items = append(list.Items, item)
  }
 }
 return list, err
}
func (c *FakePodSecurityPolicies) Watch(opts v1.ListOptions) (watch.Interface, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return c.Fake.InvokesWatch(testing.NewRootWatchAction(podsecuritypoliciesResource, opts))
}
func (c *FakePodSecurityPolicies) Create(podSecurityPolicy *policy.PodSecurityPolicy) (result *policy.PodSecurityPolicy, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewRootCreateAction(podsecuritypoliciesResource, podSecurityPolicy), &policy.PodSecurityPolicy{})
 if obj == nil {
  return nil, err
 }
 return obj.(*policy.PodSecurityPolicy), err
}
func (c *FakePodSecurityPolicies) Update(podSecurityPolicy *policy.PodSecurityPolicy) (result *policy.PodSecurityPolicy, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewRootUpdateAction(podsecuritypoliciesResource, podSecurityPolicy), &policy.PodSecurityPolicy{})
 if obj == nil {
  return nil, err
 }
 return obj.(*policy.PodSecurityPolicy), err
}
func (c *FakePodSecurityPolicies) Delete(name string, options *v1.DeleteOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 _, err := c.Fake.Invokes(testing.NewRootDeleteAction(podsecuritypoliciesResource, name), &policy.PodSecurityPolicy{})
 return err
}
func (c *FakePodSecurityPolicies) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 action := testing.NewRootDeleteCollectionAction(podsecuritypoliciesResource, listOptions)
 _, err := c.Fake.Invokes(action, &policy.PodSecurityPolicyList{})
 return err
}
func (c *FakePodSecurityPolicies) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *policy.PodSecurityPolicy, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewRootPatchSubresourceAction(podsecuritypoliciesResource, name, pt, data, subresources...), &policy.PodSecurityPolicy{})
 if obj == nil {
  return nil, err
 }
 return obj.(*policy.PodSecurityPolicy), err
}
