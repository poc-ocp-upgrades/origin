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

type FakePodDisruptionBudgets struct {
 Fake *FakePolicy
 ns   string
}

var poddisruptionbudgetsResource = schema.GroupVersionResource{Group: "policy", Version: "", Resource: "poddisruptionbudgets"}
var poddisruptionbudgetsKind = schema.GroupVersionKind{Group: "policy", Version: "", Kind: "PodDisruptionBudget"}

func (c *FakePodDisruptionBudgets) Get(name string, options v1.GetOptions) (result *policy.PodDisruptionBudget, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewGetAction(poddisruptionbudgetsResource, c.ns, name), &policy.PodDisruptionBudget{})
 if obj == nil {
  return nil, err
 }
 return obj.(*policy.PodDisruptionBudget), err
}
func (c *FakePodDisruptionBudgets) List(opts v1.ListOptions) (result *policy.PodDisruptionBudgetList, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewListAction(poddisruptionbudgetsResource, poddisruptionbudgetsKind, c.ns, opts), &policy.PodDisruptionBudgetList{})
 if obj == nil {
  return nil, err
 }
 label, _, _ := testing.ExtractFromListOptions(opts)
 if label == nil {
  label = labels.Everything()
 }
 list := &policy.PodDisruptionBudgetList{ListMeta: obj.(*policy.PodDisruptionBudgetList).ListMeta}
 for _, item := range obj.(*policy.PodDisruptionBudgetList).Items {
  if label.Matches(labels.Set(item.Labels)) {
   list.Items = append(list.Items, item)
  }
 }
 return list, err
}
func (c *FakePodDisruptionBudgets) Watch(opts v1.ListOptions) (watch.Interface, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return c.Fake.InvokesWatch(testing.NewWatchAction(poddisruptionbudgetsResource, c.ns, opts))
}
func (c *FakePodDisruptionBudgets) Create(podDisruptionBudget *policy.PodDisruptionBudget) (result *policy.PodDisruptionBudget, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewCreateAction(poddisruptionbudgetsResource, c.ns, podDisruptionBudget), &policy.PodDisruptionBudget{})
 if obj == nil {
  return nil, err
 }
 return obj.(*policy.PodDisruptionBudget), err
}
func (c *FakePodDisruptionBudgets) Update(podDisruptionBudget *policy.PodDisruptionBudget) (result *policy.PodDisruptionBudget, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewUpdateAction(poddisruptionbudgetsResource, c.ns, podDisruptionBudget), &policy.PodDisruptionBudget{})
 if obj == nil {
  return nil, err
 }
 return obj.(*policy.PodDisruptionBudget), err
}
func (c *FakePodDisruptionBudgets) UpdateStatus(podDisruptionBudget *policy.PodDisruptionBudget) (*policy.PodDisruptionBudget, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewUpdateSubresourceAction(poddisruptionbudgetsResource, "status", c.ns, podDisruptionBudget), &policy.PodDisruptionBudget{})
 if obj == nil {
  return nil, err
 }
 return obj.(*policy.PodDisruptionBudget), err
}
func (c *FakePodDisruptionBudgets) Delete(name string, options *v1.DeleteOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 _, err := c.Fake.Invokes(testing.NewDeleteAction(poddisruptionbudgetsResource, c.ns, name), &policy.PodDisruptionBudget{})
 return err
}
func (c *FakePodDisruptionBudgets) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 action := testing.NewDeleteCollectionAction(poddisruptionbudgetsResource, c.ns, listOptions)
 _, err := c.Fake.Invokes(action, &policy.PodDisruptionBudgetList{})
 return err
}
func (c *FakePodDisruptionBudgets) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *policy.PodDisruptionBudget, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewPatchSubresourceAction(poddisruptionbudgetsResource, c.ns, name, pt, data, subresources...), &policy.PodDisruptionBudget{})
 if obj == nil {
  return nil, err
 }
 return obj.(*policy.PodDisruptionBudget), err
}
