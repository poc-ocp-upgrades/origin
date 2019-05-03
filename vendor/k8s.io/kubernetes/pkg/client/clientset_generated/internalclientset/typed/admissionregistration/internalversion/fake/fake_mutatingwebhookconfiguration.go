package fake

import (
 v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
 labels "k8s.io/apimachinery/pkg/labels"
 schema "k8s.io/apimachinery/pkg/runtime/schema"
 types "k8s.io/apimachinery/pkg/types"
 watch "k8s.io/apimachinery/pkg/watch"
 testing "k8s.io/client-go/testing"
 admissionregistration "k8s.io/kubernetes/pkg/apis/admissionregistration"
)

type FakeMutatingWebhookConfigurations struct{ Fake *FakeAdmissionregistration }

var mutatingwebhookconfigurationsResource = schema.GroupVersionResource{Group: "admissionregistration.k8s.io", Version: "", Resource: "mutatingwebhookconfigurations"}
var mutatingwebhookconfigurationsKind = schema.GroupVersionKind{Group: "admissionregistration.k8s.io", Version: "", Kind: "MutatingWebhookConfiguration"}

func (c *FakeMutatingWebhookConfigurations) Get(name string, options v1.GetOptions) (result *admissionregistration.MutatingWebhookConfiguration, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewRootGetAction(mutatingwebhookconfigurationsResource, name), &admissionregistration.MutatingWebhookConfiguration{})
 if obj == nil {
  return nil, err
 }
 return obj.(*admissionregistration.MutatingWebhookConfiguration), err
}
func (c *FakeMutatingWebhookConfigurations) List(opts v1.ListOptions) (result *admissionregistration.MutatingWebhookConfigurationList, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewRootListAction(mutatingwebhookconfigurationsResource, mutatingwebhookconfigurationsKind, opts), &admissionregistration.MutatingWebhookConfigurationList{})
 if obj == nil {
  return nil, err
 }
 label, _, _ := testing.ExtractFromListOptions(opts)
 if label == nil {
  label = labels.Everything()
 }
 list := &admissionregistration.MutatingWebhookConfigurationList{ListMeta: obj.(*admissionregistration.MutatingWebhookConfigurationList).ListMeta}
 for _, item := range obj.(*admissionregistration.MutatingWebhookConfigurationList).Items {
  if label.Matches(labels.Set(item.Labels)) {
   list.Items = append(list.Items, item)
  }
 }
 return list, err
}
func (c *FakeMutatingWebhookConfigurations) Watch(opts v1.ListOptions) (watch.Interface, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return c.Fake.InvokesWatch(testing.NewRootWatchAction(mutatingwebhookconfigurationsResource, opts))
}
func (c *FakeMutatingWebhookConfigurations) Create(mutatingWebhookConfiguration *admissionregistration.MutatingWebhookConfiguration) (result *admissionregistration.MutatingWebhookConfiguration, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewRootCreateAction(mutatingwebhookconfigurationsResource, mutatingWebhookConfiguration), &admissionregistration.MutatingWebhookConfiguration{})
 if obj == nil {
  return nil, err
 }
 return obj.(*admissionregistration.MutatingWebhookConfiguration), err
}
func (c *FakeMutatingWebhookConfigurations) Update(mutatingWebhookConfiguration *admissionregistration.MutatingWebhookConfiguration) (result *admissionregistration.MutatingWebhookConfiguration, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewRootUpdateAction(mutatingwebhookconfigurationsResource, mutatingWebhookConfiguration), &admissionregistration.MutatingWebhookConfiguration{})
 if obj == nil {
  return nil, err
 }
 return obj.(*admissionregistration.MutatingWebhookConfiguration), err
}
func (c *FakeMutatingWebhookConfigurations) Delete(name string, options *v1.DeleteOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 _, err := c.Fake.Invokes(testing.NewRootDeleteAction(mutatingwebhookconfigurationsResource, name), &admissionregistration.MutatingWebhookConfiguration{})
 return err
}
func (c *FakeMutatingWebhookConfigurations) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 action := testing.NewRootDeleteCollectionAction(mutatingwebhookconfigurationsResource, listOptions)
 _, err := c.Fake.Invokes(action, &admissionregistration.MutatingWebhookConfigurationList{})
 return err
}
func (c *FakeMutatingWebhookConfigurations) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *admissionregistration.MutatingWebhookConfiguration, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewRootPatchSubresourceAction(mutatingwebhookconfigurationsResource, name, pt, data, subresources...), &admissionregistration.MutatingWebhookConfiguration{})
 if obj == nil {
  return nil, err
 }
 return obj.(*admissionregistration.MutatingWebhookConfiguration), err
}
