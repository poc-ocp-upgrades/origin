package fake

import (
 v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
 labels "k8s.io/apimachinery/pkg/labels"
 schema "k8s.io/apimachinery/pkg/runtime/schema"
 types "k8s.io/apimachinery/pkg/types"
 watch "k8s.io/apimachinery/pkg/watch"
 testing "k8s.io/client-go/testing"
 apps "k8s.io/kubernetes/pkg/apis/apps"
)

type FakeControllerRevisions struct {
 Fake *FakeApps
 ns   string
}

var controllerrevisionsResource = schema.GroupVersionResource{Group: "apps", Version: "", Resource: "controllerrevisions"}
var controllerrevisionsKind = schema.GroupVersionKind{Group: "apps", Version: "", Kind: "ControllerRevision"}

func (c *FakeControllerRevisions) Get(name string, options v1.GetOptions) (result *apps.ControllerRevision, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewGetAction(controllerrevisionsResource, c.ns, name), &apps.ControllerRevision{})
 if obj == nil {
  return nil, err
 }
 return obj.(*apps.ControllerRevision), err
}
func (c *FakeControllerRevisions) List(opts v1.ListOptions) (result *apps.ControllerRevisionList, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewListAction(controllerrevisionsResource, controllerrevisionsKind, c.ns, opts), &apps.ControllerRevisionList{})
 if obj == nil {
  return nil, err
 }
 label, _, _ := testing.ExtractFromListOptions(opts)
 if label == nil {
  label = labels.Everything()
 }
 list := &apps.ControllerRevisionList{ListMeta: obj.(*apps.ControllerRevisionList).ListMeta}
 for _, item := range obj.(*apps.ControllerRevisionList).Items {
  if label.Matches(labels.Set(item.Labels)) {
   list.Items = append(list.Items, item)
  }
 }
 return list, err
}
func (c *FakeControllerRevisions) Watch(opts v1.ListOptions) (watch.Interface, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return c.Fake.InvokesWatch(testing.NewWatchAction(controllerrevisionsResource, c.ns, opts))
}
func (c *FakeControllerRevisions) Create(controllerRevision *apps.ControllerRevision) (result *apps.ControllerRevision, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewCreateAction(controllerrevisionsResource, c.ns, controllerRevision), &apps.ControllerRevision{})
 if obj == nil {
  return nil, err
 }
 return obj.(*apps.ControllerRevision), err
}
func (c *FakeControllerRevisions) Update(controllerRevision *apps.ControllerRevision) (result *apps.ControllerRevision, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewUpdateAction(controllerrevisionsResource, c.ns, controllerRevision), &apps.ControllerRevision{})
 if obj == nil {
  return nil, err
 }
 return obj.(*apps.ControllerRevision), err
}
func (c *FakeControllerRevisions) Delete(name string, options *v1.DeleteOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 _, err := c.Fake.Invokes(testing.NewDeleteAction(controllerrevisionsResource, c.ns, name), &apps.ControllerRevision{})
 return err
}
func (c *FakeControllerRevisions) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 action := testing.NewDeleteCollectionAction(controllerrevisionsResource, c.ns, listOptions)
 _, err := c.Fake.Invokes(action, &apps.ControllerRevisionList{})
 return err
}
func (c *FakeControllerRevisions) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *apps.ControllerRevision, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewPatchSubresourceAction(controllerrevisionsResource, c.ns, name, pt, data, subresources...), &apps.ControllerRevision{})
 if obj == nil {
  return nil, err
 }
 return obj.(*apps.ControllerRevision), err
}
