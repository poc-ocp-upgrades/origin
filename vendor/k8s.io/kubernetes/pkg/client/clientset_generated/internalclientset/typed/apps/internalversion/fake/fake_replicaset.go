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

type FakeReplicaSets struct {
 Fake *FakeApps
 ns   string
}

var replicasetsResource = schema.GroupVersionResource{Group: "apps", Version: "", Resource: "replicasets"}
var replicasetsKind = schema.GroupVersionKind{Group: "apps", Version: "", Kind: "ReplicaSet"}

func (c *FakeReplicaSets) Get(name string, options v1.GetOptions) (result *apps.ReplicaSet, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewGetAction(replicasetsResource, c.ns, name), &apps.ReplicaSet{})
 if obj == nil {
  return nil, err
 }
 return obj.(*apps.ReplicaSet), err
}
func (c *FakeReplicaSets) List(opts v1.ListOptions) (result *apps.ReplicaSetList, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewListAction(replicasetsResource, replicasetsKind, c.ns, opts), &apps.ReplicaSetList{})
 if obj == nil {
  return nil, err
 }
 label, _, _ := testing.ExtractFromListOptions(opts)
 if label == nil {
  label = labels.Everything()
 }
 list := &apps.ReplicaSetList{ListMeta: obj.(*apps.ReplicaSetList).ListMeta}
 for _, item := range obj.(*apps.ReplicaSetList).Items {
  if label.Matches(labels.Set(item.Labels)) {
   list.Items = append(list.Items, item)
  }
 }
 return list, err
}
func (c *FakeReplicaSets) Watch(opts v1.ListOptions) (watch.Interface, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return c.Fake.InvokesWatch(testing.NewWatchAction(replicasetsResource, c.ns, opts))
}
func (c *FakeReplicaSets) Create(replicaSet *apps.ReplicaSet) (result *apps.ReplicaSet, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewCreateAction(replicasetsResource, c.ns, replicaSet), &apps.ReplicaSet{})
 if obj == nil {
  return nil, err
 }
 return obj.(*apps.ReplicaSet), err
}
func (c *FakeReplicaSets) Update(replicaSet *apps.ReplicaSet) (result *apps.ReplicaSet, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewUpdateAction(replicasetsResource, c.ns, replicaSet), &apps.ReplicaSet{})
 if obj == nil {
  return nil, err
 }
 return obj.(*apps.ReplicaSet), err
}
func (c *FakeReplicaSets) UpdateStatus(replicaSet *apps.ReplicaSet) (*apps.ReplicaSet, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewUpdateSubresourceAction(replicasetsResource, "status", c.ns, replicaSet), &apps.ReplicaSet{})
 if obj == nil {
  return nil, err
 }
 return obj.(*apps.ReplicaSet), err
}
func (c *FakeReplicaSets) Delete(name string, options *v1.DeleteOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 _, err := c.Fake.Invokes(testing.NewDeleteAction(replicasetsResource, c.ns, name), &apps.ReplicaSet{})
 return err
}
func (c *FakeReplicaSets) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 action := testing.NewDeleteCollectionAction(replicasetsResource, c.ns, listOptions)
 _, err := c.Fake.Invokes(action, &apps.ReplicaSetList{})
 return err
}
func (c *FakeReplicaSets) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *apps.ReplicaSet, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewPatchSubresourceAction(replicasetsResource, c.ns, name, pt, data, subresources...), &apps.ReplicaSet{})
 if obj == nil {
  return nil, err
 }
 return obj.(*apps.ReplicaSet), err
}
