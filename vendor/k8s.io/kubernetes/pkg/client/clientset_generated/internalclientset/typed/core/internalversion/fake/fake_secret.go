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

type FakeSecrets struct {
 Fake *FakeCore
 ns   string
}

var secretsResource = schema.GroupVersionResource{Group: "", Version: "", Resource: "secrets"}
var secretsKind = schema.GroupVersionKind{Group: "", Version: "", Kind: "Secret"}

func (c *FakeSecrets) Get(name string, options v1.GetOptions) (result *core.Secret, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewGetAction(secretsResource, c.ns, name), &core.Secret{})
 if obj == nil {
  return nil, err
 }
 return obj.(*core.Secret), err
}
func (c *FakeSecrets) List(opts v1.ListOptions) (result *core.SecretList, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewListAction(secretsResource, secretsKind, c.ns, opts), &core.SecretList{})
 if obj == nil {
  return nil, err
 }
 label, _, _ := testing.ExtractFromListOptions(opts)
 if label == nil {
  label = labels.Everything()
 }
 list := &core.SecretList{ListMeta: obj.(*core.SecretList).ListMeta}
 for _, item := range obj.(*core.SecretList).Items {
  if label.Matches(labels.Set(item.Labels)) {
   list.Items = append(list.Items, item)
  }
 }
 return list, err
}
func (c *FakeSecrets) Watch(opts v1.ListOptions) (watch.Interface, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return c.Fake.InvokesWatch(testing.NewWatchAction(secretsResource, c.ns, opts))
}
func (c *FakeSecrets) Create(secret *core.Secret) (result *core.Secret, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewCreateAction(secretsResource, c.ns, secret), &core.Secret{})
 if obj == nil {
  return nil, err
 }
 return obj.(*core.Secret), err
}
func (c *FakeSecrets) Update(secret *core.Secret) (result *core.Secret, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewUpdateAction(secretsResource, c.ns, secret), &core.Secret{})
 if obj == nil {
  return nil, err
 }
 return obj.(*core.Secret), err
}
func (c *FakeSecrets) Delete(name string, options *v1.DeleteOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 _, err := c.Fake.Invokes(testing.NewDeleteAction(secretsResource, c.ns, name), &core.Secret{})
 return err
}
func (c *FakeSecrets) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 action := testing.NewDeleteCollectionAction(secretsResource, c.ns, listOptions)
 _, err := c.Fake.Invokes(action, &core.SecretList{})
 return err
}
func (c *FakeSecrets) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *core.Secret, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewPatchSubresourceAction(secretsResource, c.ns, name, pt, data, subresources...), &core.Secret{})
 if obj == nil {
  return nil, err
 }
 return obj.(*core.Secret), err
}
