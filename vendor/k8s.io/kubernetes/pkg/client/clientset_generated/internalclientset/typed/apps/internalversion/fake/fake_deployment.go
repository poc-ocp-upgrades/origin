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

type FakeDeployments struct {
 Fake *FakeApps
 ns   string
}

var deploymentsResource = schema.GroupVersionResource{Group: "apps", Version: "", Resource: "deployments"}
var deploymentsKind = schema.GroupVersionKind{Group: "apps", Version: "", Kind: "Deployment"}

func (c *FakeDeployments) Get(name string, options v1.GetOptions) (result *apps.Deployment, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewGetAction(deploymentsResource, c.ns, name), &apps.Deployment{})
 if obj == nil {
  return nil, err
 }
 return obj.(*apps.Deployment), err
}
func (c *FakeDeployments) List(opts v1.ListOptions) (result *apps.DeploymentList, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewListAction(deploymentsResource, deploymentsKind, c.ns, opts), &apps.DeploymentList{})
 if obj == nil {
  return nil, err
 }
 label, _, _ := testing.ExtractFromListOptions(opts)
 if label == nil {
  label = labels.Everything()
 }
 list := &apps.DeploymentList{ListMeta: obj.(*apps.DeploymentList).ListMeta}
 for _, item := range obj.(*apps.DeploymentList).Items {
  if label.Matches(labels.Set(item.Labels)) {
   list.Items = append(list.Items, item)
  }
 }
 return list, err
}
func (c *FakeDeployments) Watch(opts v1.ListOptions) (watch.Interface, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return c.Fake.InvokesWatch(testing.NewWatchAction(deploymentsResource, c.ns, opts))
}
func (c *FakeDeployments) Create(deployment *apps.Deployment) (result *apps.Deployment, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewCreateAction(deploymentsResource, c.ns, deployment), &apps.Deployment{})
 if obj == nil {
  return nil, err
 }
 return obj.(*apps.Deployment), err
}
func (c *FakeDeployments) Update(deployment *apps.Deployment) (result *apps.Deployment, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewUpdateAction(deploymentsResource, c.ns, deployment), &apps.Deployment{})
 if obj == nil {
  return nil, err
 }
 return obj.(*apps.Deployment), err
}
func (c *FakeDeployments) UpdateStatus(deployment *apps.Deployment) (*apps.Deployment, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewUpdateSubresourceAction(deploymentsResource, "status", c.ns, deployment), &apps.Deployment{})
 if obj == nil {
  return nil, err
 }
 return obj.(*apps.Deployment), err
}
func (c *FakeDeployments) Delete(name string, options *v1.DeleteOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 _, err := c.Fake.Invokes(testing.NewDeleteAction(deploymentsResource, c.ns, name), &apps.Deployment{})
 return err
}
func (c *FakeDeployments) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 action := testing.NewDeleteCollectionAction(deploymentsResource, c.ns, listOptions)
 _, err := c.Fake.Invokes(action, &apps.DeploymentList{})
 return err
}
func (c *FakeDeployments) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *apps.Deployment, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewPatchSubresourceAction(deploymentsResource, c.ns, name, pt, data, subresources...), &apps.Deployment{})
 if obj == nil {
  return nil, err
 }
 return obj.(*apps.Deployment), err
}
