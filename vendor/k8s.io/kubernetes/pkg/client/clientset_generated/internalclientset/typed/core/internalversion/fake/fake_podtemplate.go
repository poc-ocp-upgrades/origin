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

type FakePodTemplates struct {
 Fake *FakeCore
 ns   string
}

var podtemplatesResource = schema.GroupVersionResource{Group: "", Version: "", Resource: "podtemplates"}
var podtemplatesKind = schema.GroupVersionKind{Group: "", Version: "", Kind: "PodTemplate"}

func (c *FakePodTemplates) Get(name string, options v1.GetOptions) (result *core.PodTemplate, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewGetAction(podtemplatesResource, c.ns, name), &core.PodTemplate{})
 if obj == nil {
  return nil, err
 }
 return obj.(*core.PodTemplate), err
}
func (c *FakePodTemplates) List(opts v1.ListOptions) (result *core.PodTemplateList, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewListAction(podtemplatesResource, podtemplatesKind, c.ns, opts), &core.PodTemplateList{})
 if obj == nil {
  return nil, err
 }
 label, _, _ := testing.ExtractFromListOptions(opts)
 if label == nil {
  label = labels.Everything()
 }
 list := &core.PodTemplateList{ListMeta: obj.(*core.PodTemplateList).ListMeta}
 for _, item := range obj.(*core.PodTemplateList).Items {
  if label.Matches(labels.Set(item.Labels)) {
   list.Items = append(list.Items, item)
  }
 }
 return list, err
}
func (c *FakePodTemplates) Watch(opts v1.ListOptions) (watch.Interface, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return c.Fake.InvokesWatch(testing.NewWatchAction(podtemplatesResource, c.ns, opts))
}
func (c *FakePodTemplates) Create(podTemplate *core.PodTemplate) (result *core.PodTemplate, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewCreateAction(podtemplatesResource, c.ns, podTemplate), &core.PodTemplate{})
 if obj == nil {
  return nil, err
 }
 return obj.(*core.PodTemplate), err
}
func (c *FakePodTemplates) Update(podTemplate *core.PodTemplate) (result *core.PodTemplate, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewUpdateAction(podtemplatesResource, c.ns, podTemplate), &core.PodTemplate{})
 if obj == nil {
  return nil, err
 }
 return obj.(*core.PodTemplate), err
}
func (c *FakePodTemplates) Delete(name string, options *v1.DeleteOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 _, err := c.Fake.Invokes(testing.NewDeleteAction(podtemplatesResource, c.ns, name), &core.PodTemplate{})
 return err
}
func (c *FakePodTemplates) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 action := testing.NewDeleteCollectionAction(podtemplatesResource, c.ns, listOptions)
 _, err := c.Fake.Invokes(action, &core.PodTemplateList{})
 return err
}
func (c *FakePodTemplates) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *core.PodTemplate, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewPatchSubresourceAction(podtemplatesResource, c.ns, name, pt, data, subresources...), &core.PodTemplate{})
 if obj == nil {
  return nil, err
 }
 return obj.(*core.PodTemplate), err
}
