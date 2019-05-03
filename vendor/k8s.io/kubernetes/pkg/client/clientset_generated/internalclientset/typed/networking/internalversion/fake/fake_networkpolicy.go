package fake

import (
 v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
 labels "k8s.io/apimachinery/pkg/labels"
 schema "k8s.io/apimachinery/pkg/runtime/schema"
 types "k8s.io/apimachinery/pkg/types"
 watch "k8s.io/apimachinery/pkg/watch"
 testing "k8s.io/client-go/testing"
 networking "k8s.io/kubernetes/pkg/apis/networking"
)

type FakeNetworkPolicies struct {
 Fake *FakeNetworking
 ns   string
}

var networkpoliciesResource = schema.GroupVersionResource{Group: "networking.k8s.io", Version: "", Resource: "networkpolicies"}
var networkpoliciesKind = schema.GroupVersionKind{Group: "networking.k8s.io", Version: "", Kind: "NetworkPolicy"}

func (c *FakeNetworkPolicies) Get(name string, options v1.GetOptions) (result *networking.NetworkPolicy, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewGetAction(networkpoliciesResource, c.ns, name), &networking.NetworkPolicy{})
 if obj == nil {
  return nil, err
 }
 return obj.(*networking.NetworkPolicy), err
}
func (c *FakeNetworkPolicies) List(opts v1.ListOptions) (result *networking.NetworkPolicyList, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewListAction(networkpoliciesResource, networkpoliciesKind, c.ns, opts), &networking.NetworkPolicyList{})
 if obj == nil {
  return nil, err
 }
 label, _, _ := testing.ExtractFromListOptions(opts)
 if label == nil {
  label = labels.Everything()
 }
 list := &networking.NetworkPolicyList{ListMeta: obj.(*networking.NetworkPolicyList).ListMeta}
 for _, item := range obj.(*networking.NetworkPolicyList).Items {
  if label.Matches(labels.Set(item.Labels)) {
   list.Items = append(list.Items, item)
  }
 }
 return list, err
}
func (c *FakeNetworkPolicies) Watch(opts v1.ListOptions) (watch.Interface, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return c.Fake.InvokesWatch(testing.NewWatchAction(networkpoliciesResource, c.ns, opts))
}
func (c *FakeNetworkPolicies) Create(networkPolicy *networking.NetworkPolicy) (result *networking.NetworkPolicy, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewCreateAction(networkpoliciesResource, c.ns, networkPolicy), &networking.NetworkPolicy{})
 if obj == nil {
  return nil, err
 }
 return obj.(*networking.NetworkPolicy), err
}
func (c *FakeNetworkPolicies) Update(networkPolicy *networking.NetworkPolicy) (result *networking.NetworkPolicy, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewUpdateAction(networkpoliciesResource, c.ns, networkPolicy), &networking.NetworkPolicy{})
 if obj == nil {
  return nil, err
 }
 return obj.(*networking.NetworkPolicy), err
}
func (c *FakeNetworkPolicies) Delete(name string, options *v1.DeleteOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 _, err := c.Fake.Invokes(testing.NewDeleteAction(networkpoliciesResource, c.ns, name), &networking.NetworkPolicy{})
 return err
}
func (c *FakeNetworkPolicies) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 action := testing.NewDeleteCollectionAction(networkpoliciesResource, c.ns, listOptions)
 _, err := c.Fake.Invokes(action, &networking.NetworkPolicyList{})
 return err
}
func (c *FakeNetworkPolicies) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *networking.NetworkPolicy, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewPatchSubresourceAction(networkpoliciesResource, c.ns, name, pt, data, subresources...), &networking.NetworkPolicy{})
 if obj == nil {
  return nil, err
 }
 return obj.(*networking.NetworkPolicy), err
}
