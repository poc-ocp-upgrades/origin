package fake

import (
 v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
 labels "k8s.io/apimachinery/pkg/labels"
 schema "k8s.io/apimachinery/pkg/runtime/schema"
 types "k8s.io/apimachinery/pkg/types"
 watch "k8s.io/apimachinery/pkg/watch"
 testing "k8s.io/client-go/testing"
 core "k8s.io/kubernetes/pkg/apis/core"
)

type FakeNodes struct{ Fake *FakeCore }

var nodesResource = schema.GroupVersionResource{Group: "", Version: "", Resource: "nodes"}
var nodesKind = schema.GroupVersionKind{Group: "", Version: "", Kind: "Node"}

func (c *FakeNodes) Get(name string, options v1.GetOptions) (result *core.Node, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewRootGetAction(nodesResource, name), &core.Node{})
 if obj == nil {
  return nil, err
 }
 return obj.(*core.Node), err
}
func (c *FakeNodes) List(opts v1.ListOptions) (result *core.NodeList, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewRootListAction(nodesResource, nodesKind, opts), &core.NodeList{})
 if obj == nil {
  return nil, err
 }
 label, _, _ := testing.ExtractFromListOptions(opts)
 if label == nil {
  label = labels.Everything()
 }
 list := &core.NodeList{ListMeta: obj.(*core.NodeList).ListMeta}
 for _, item := range obj.(*core.NodeList).Items {
  if label.Matches(labels.Set(item.Labels)) {
   list.Items = append(list.Items, item)
  }
 }
 return list, err
}
func (c *FakeNodes) Watch(opts v1.ListOptions) (watch.Interface, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return c.Fake.InvokesWatch(testing.NewRootWatchAction(nodesResource, opts))
}
func (c *FakeNodes) Create(node *core.Node) (result *core.Node, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewRootCreateAction(nodesResource, node), &core.Node{})
 if obj == nil {
  return nil, err
 }
 return obj.(*core.Node), err
}
func (c *FakeNodes) Update(node *core.Node) (result *core.Node, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewRootUpdateAction(nodesResource, node), &core.Node{})
 if obj == nil {
  return nil, err
 }
 return obj.(*core.Node), err
}
func (c *FakeNodes) UpdateStatus(node *core.Node) (*core.Node, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewRootUpdateSubresourceAction(nodesResource, "status", node), &core.Node{})
 if obj == nil {
  return nil, err
 }
 return obj.(*core.Node), err
}
func (c *FakeNodes) Delete(name string, options *v1.DeleteOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 _, err := c.Fake.Invokes(testing.NewRootDeleteAction(nodesResource, name), &core.Node{})
 return err
}
func (c *FakeNodes) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 action := testing.NewRootDeleteCollectionAction(nodesResource, listOptions)
 _, err := c.Fake.Invokes(action, &core.NodeList{})
 return err
}
func (c *FakeNodes) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *core.Node, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewRootPatchSubresourceAction(nodesResource, name, pt, data, subresources...), &core.Node{})
 if obj == nil {
  return nil, err
 }
 return obj.(*core.Node), err
}
