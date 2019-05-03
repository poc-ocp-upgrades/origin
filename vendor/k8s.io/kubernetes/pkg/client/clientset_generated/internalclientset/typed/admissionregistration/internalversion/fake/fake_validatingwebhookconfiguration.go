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

type FakeValidatingWebhookConfigurations struct{ Fake *FakeAdmissionregistration }

var validatingwebhookconfigurationsResource = schema.GroupVersionResource{Group: "admissionregistration.k8s.io", Version: "", Resource: "validatingwebhookconfigurations"}
var validatingwebhookconfigurationsKind = schema.GroupVersionKind{Group: "admissionregistration.k8s.io", Version: "", Kind: "ValidatingWebhookConfiguration"}

func (c *FakeValidatingWebhookConfigurations) Get(name string, options v1.GetOptions) (result *admissionregistration.ValidatingWebhookConfiguration, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewRootGetAction(validatingwebhookconfigurationsResource, name), &admissionregistration.ValidatingWebhookConfiguration{})
 if obj == nil {
  return nil, err
 }
 return obj.(*admissionregistration.ValidatingWebhookConfiguration), err
}
func (c *FakeValidatingWebhookConfigurations) List(opts v1.ListOptions) (result *admissionregistration.ValidatingWebhookConfigurationList, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewRootListAction(validatingwebhookconfigurationsResource, validatingwebhookconfigurationsKind, opts), &admissionregistration.ValidatingWebhookConfigurationList{})
 if obj == nil {
  return nil, err
 }
 label, _, _ := testing.ExtractFromListOptions(opts)
 if label == nil {
  label = labels.Everything()
 }
 list := &admissionregistration.ValidatingWebhookConfigurationList{ListMeta: obj.(*admissionregistration.ValidatingWebhookConfigurationList).ListMeta}
 for _, item := range obj.(*admissionregistration.ValidatingWebhookConfigurationList).Items {
  if label.Matches(labels.Set(item.Labels)) {
   list.Items = append(list.Items, item)
  }
 }
 return list, err
}
func (c *FakeValidatingWebhookConfigurations) Watch(opts v1.ListOptions) (watch.Interface, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return c.Fake.InvokesWatch(testing.NewRootWatchAction(validatingwebhookconfigurationsResource, opts))
}
func (c *FakeValidatingWebhookConfigurations) Create(validatingWebhookConfiguration *admissionregistration.ValidatingWebhookConfiguration) (result *admissionregistration.ValidatingWebhookConfiguration, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewRootCreateAction(validatingwebhookconfigurationsResource, validatingWebhookConfiguration), &admissionregistration.ValidatingWebhookConfiguration{})
 if obj == nil {
  return nil, err
 }
 return obj.(*admissionregistration.ValidatingWebhookConfiguration), err
}
func (c *FakeValidatingWebhookConfigurations) Update(validatingWebhookConfiguration *admissionregistration.ValidatingWebhookConfiguration) (result *admissionregistration.ValidatingWebhookConfiguration, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewRootUpdateAction(validatingwebhookconfigurationsResource, validatingWebhookConfiguration), &admissionregistration.ValidatingWebhookConfiguration{})
 if obj == nil {
  return nil, err
 }
 return obj.(*admissionregistration.ValidatingWebhookConfiguration), err
}
func (c *FakeValidatingWebhookConfigurations) Delete(name string, options *v1.DeleteOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 _, err := c.Fake.Invokes(testing.NewRootDeleteAction(validatingwebhookconfigurationsResource, name), &admissionregistration.ValidatingWebhookConfiguration{})
 return err
}
func (c *FakeValidatingWebhookConfigurations) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 action := testing.NewRootDeleteCollectionAction(validatingwebhookconfigurationsResource, listOptions)
 _, err := c.Fake.Invokes(action, &admissionregistration.ValidatingWebhookConfigurationList{})
 return err
}
func (c *FakeValidatingWebhookConfigurations) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *admissionregistration.ValidatingWebhookConfiguration, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewRootPatchSubresourceAction(validatingwebhookconfigurationsResource, name, pt, data, subresources...), &admissionregistration.ValidatingWebhookConfiguration{})
 if obj == nil {
  return nil, err
 }
 return obj.(*admissionregistration.ValidatingWebhookConfiguration), err
}
