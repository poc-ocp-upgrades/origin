package fake

import (
 v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
 labels "k8s.io/apimachinery/pkg/labels"
 schema "k8s.io/apimachinery/pkg/runtime/schema"
 types "k8s.io/apimachinery/pkg/types"
 watch "k8s.io/apimachinery/pkg/watch"
 testing "k8s.io/client-go/testing"
 autoscaling "k8s.io/kubernetes/pkg/apis/autoscaling"
 core "k8s.io/kubernetes/pkg/apis/core"
)

type FakeReplicationControllers struct {
 Fake *FakeCore
 ns   string
}

var replicationcontrollersResource = schema.GroupVersionResource{Group: "", Version: "", Resource: "replicationcontrollers"}
var replicationcontrollersKind = schema.GroupVersionKind{Group: "", Version: "", Kind: "ReplicationController"}

func (c *FakeReplicationControllers) Get(name string, options v1.GetOptions) (result *core.ReplicationController, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewGetAction(replicationcontrollersResource, c.ns, name), &core.ReplicationController{})
 if obj == nil {
  return nil, err
 }
 return obj.(*core.ReplicationController), err
}
func (c *FakeReplicationControllers) List(opts v1.ListOptions) (result *core.ReplicationControllerList, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewListAction(replicationcontrollersResource, replicationcontrollersKind, c.ns, opts), &core.ReplicationControllerList{})
 if obj == nil {
  return nil, err
 }
 label, _, _ := testing.ExtractFromListOptions(opts)
 if label == nil {
  label = labels.Everything()
 }
 list := &core.ReplicationControllerList{ListMeta: obj.(*core.ReplicationControllerList).ListMeta}
 for _, item := range obj.(*core.ReplicationControllerList).Items {
  if label.Matches(labels.Set(item.Labels)) {
   list.Items = append(list.Items, item)
  }
 }
 return list, err
}
func (c *FakeReplicationControllers) Watch(opts v1.ListOptions) (watch.Interface, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return c.Fake.InvokesWatch(testing.NewWatchAction(replicationcontrollersResource, c.ns, opts))
}
func (c *FakeReplicationControllers) Create(replicationController *core.ReplicationController) (result *core.ReplicationController, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewCreateAction(replicationcontrollersResource, c.ns, replicationController), &core.ReplicationController{})
 if obj == nil {
  return nil, err
 }
 return obj.(*core.ReplicationController), err
}
func (c *FakeReplicationControllers) Update(replicationController *core.ReplicationController) (result *core.ReplicationController, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewUpdateAction(replicationcontrollersResource, c.ns, replicationController), &core.ReplicationController{})
 if obj == nil {
  return nil, err
 }
 return obj.(*core.ReplicationController), err
}
func (c *FakeReplicationControllers) UpdateStatus(replicationController *core.ReplicationController) (*core.ReplicationController, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewUpdateSubresourceAction(replicationcontrollersResource, "status", c.ns, replicationController), &core.ReplicationController{})
 if obj == nil {
  return nil, err
 }
 return obj.(*core.ReplicationController), err
}
func (c *FakeReplicationControllers) Delete(name string, options *v1.DeleteOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 _, err := c.Fake.Invokes(testing.NewDeleteAction(replicationcontrollersResource, c.ns, name), &core.ReplicationController{})
 return err
}
func (c *FakeReplicationControllers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 action := testing.NewDeleteCollectionAction(replicationcontrollersResource, c.ns, listOptions)
 _, err := c.Fake.Invokes(action, &core.ReplicationControllerList{})
 return err
}
func (c *FakeReplicationControllers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *core.ReplicationController, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewPatchSubresourceAction(replicationcontrollersResource, c.ns, name, pt, data, subresources...), &core.ReplicationController{})
 if obj == nil {
  return nil, err
 }
 return obj.(*core.ReplicationController), err
}
func (c *FakeReplicationControllers) GetScale(replicationControllerName string, options v1.GetOptions) (result *autoscaling.Scale, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewGetSubresourceAction(replicationcontrollersResource, c.ns, "scale", replicationControllerName), &autoscaling.Scale{})
 if obj == nil {
  return nil, err
 }
 return obj.(*autoscaling.Scale), err
}
func (c *FakeReplicationControllers) UpdateScale(replicationControllerName string, scale *autoscaling.Scale) (result *autoscaling.Scale, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewUpdateSubresourceAction(replicationcontrollersResource, "scale", c.ns, scale), &autoscaling.Scale{})
 if obj == nil {
  return nil, err
 }
 return obj.(*autoscaling.Scale), err
}
