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

type FakeNamespaces struct{ Fake *FakeCore }

var namespacesResource = schema.GroupVersionResource{Group: "", Version: "", Resource: "namespaces"}
var namespacesKind = schema.GroupVersionKind{Group: "", Version: "", Kind: "Namespace"}

func (c *FakeNamespaces) Get(name string, options v1.GetOptions) (result *core.Namespace, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewRootGetAction(namespacesResource, name), &core.Namespace{})
 if obj == nil {
  return nil, err
 }
 return obj.(*core.Namespace), err
}
func (c *FakeNamespaces) List(opts v1.ListOptions) (result *core.NamespaceList, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewRootListAction(namespacesResource, namespacesKind, opts), &core.NamespaceList{})
 if obj == nil {
  return nil, err
 }
 label, _, _ := testing.ExtractFromListOptions(opts)
 if label == nil {
  label = labels.Everything()
 }
 list := &core.NamespaceList{ListMeta: obj.(*core.NamespaceList).ListMeta}
 for _, item := range obj.(*core.NamespaceList).Items {
  if label.Matches(labels.Set(item.Labels)) {
   list.Items = append(list.Items, item)
  }
 }
 return list, err
}
func (c *FakeNamespaces) Watch(opts v1.ListOptions) (watch.Interface, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return c.Fake.InvokesWatch(testing.NewRootWatchAction(namespacesResource, opts))
}
func (c *FakeNamespaces) Create(namespace *core.Namespace) (result *core.Namespace, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewRootCreateAction(namespacesResource, namespace), &core.Namespace{})
 if obj == nil {
  return nil, err
 }
 return obj.(*core.Namespace), err
}
func (c *FakeNamespaces) Update(namespace *core.Namespace) (result *core.Namespace, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewRootUpdateAction(namespacesResource, namespace), &core.Namespace{})
 if obj == nil {
  return nil, err
 }
 return obj.(*core.Namespace), err
}
func (c *FakeNamespaces) UpdateStatus(namespace *core.Namespace) (*core.Namespace, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewRootUpdateSubresourceAction(namespacesResource, "status", namespace), &core.Namespace{})
 if obj == nil {
  return nil, err
 }
 return obj.(*core.Namespace), err
}
func (c *FakeNamespaces) Delete(name string, options *v1.DeleteOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 _, err := c.Fake.Invokes(testing.NewRootDeleteAction(namespacesResource, name), &core.Namespace{})
 return err
}
func (c *FakeNamespaces) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 action := testing.NewRootDeleteCollectionAction(namespacesResource, listOptions)
 _, err := c.Fake.Invokes(action, &core.NamespaceList{})
 return err
}
func (c *FakeNamespaces) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *core.Namespace, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewRootPatchSubresourceAction(namespacesResource, name, pt, data, subresources...), &core.Namespace{})
 if obj == nil {
  return nil, err
 }
 return obj.(*core.Namespace), err
}
