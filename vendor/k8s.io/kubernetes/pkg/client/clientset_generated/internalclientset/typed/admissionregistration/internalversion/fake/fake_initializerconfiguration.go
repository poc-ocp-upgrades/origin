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

type FakeInitializerConfigurations struct{ Fake *FakeAdmissionregistration }

var initializerconfigurationsResource = schema.GroupVersionResource{Group: "admissionregistration.k8s.io", Version: "", Resource: "initializerconfigurations"}
var initializerconfigurationsKind = schema.GroupVersionKind{Group: "admissionregistration.k8s.io", Version: "", Kind: "InitializerConfiguration"}

func (c *FakeInitializerConfigurations) Get(name string, options v1.GetOptions) (result *admissionregistration.InitializerConfiguration, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewRootGetAction(initializerconfigurationsResource, name), &admissionregistration.InitializerConfiguration{})
 if obj == nil {
  return nil, err
 }
 return obj.(*admissionregistration.InitializerConfiguration), err
}
func (c *FakeInitializerConfigurations) List(opts v1.ListOptions) (result *admissionregistration.InitializerConfigurationList, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewRootListAction(initializerconfigurationsResource, initializerconfigurationsKind, opts), &admissionregistration.InitializerConfigurationList{})
 if obj == nil {
  return nil, err
 }
 label, _, _ := testing.ExtractFromListOptions(opts)
 if label == nil {
  label = labels.Everything()
 }
 list := &admissionregistration.InitializerConfigurationList{ListMeta: obj.(*admissionregistration.InitializerConfigurationList).ListMeta}
 for _, item := range obj.(*admissionregistration.InitializerConfigurationList).Items {
  if label.Matches(labels.Set(item.Labels)) {
   list.Items = append(list.Items, item)
  }
 }
 return list, err
}
func (c *FakeInitializerConfigurations) Watch(opts v1.ListOptions) (watch.Interface, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return c.Fake.InvokesWatch(testing.NewRootWatchAction(initializerconfigurationsResource, opts))
}
func (c *FakeInitializerConfigurations) Create(initializerConfiguration *admissionregistration.InitializerConfiguration) (result *admissionregistration.InitializerConfiguration, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewRootCreateAction(initializerconfigurationsResource, initializerConfiguration), &admissionregistration.InitializerConfiguration{})
 if obj == nil {
  return nil, err
 }
 return obj.(*admissionregistration.InitializerConfiguration), err
}
func (c *FakeInitializerConfigurations) Update(initializerConfiguration *admissionregistration.InitializerConfiguration) (result *admissionregistration.InitializerConfiguration, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewRootUpdateAction(initializerconfigurationsResource, initializerConfiguration), &admissionregistration.InitializerConfiguration{})
 if obj == nil {
  return nil, err
 }
 return obj.(*admissionregistration.InitializerConfiguration), err
}
func (c *FakeInitializerConfigurations) Delete(name string, options *v1.DeleteOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 _, err := c.Fake.Invokes(testing.NewRootDeleteAction(initializerconfigurationsResource, name), &admissionregistration.InitializerConfiguration{})
 return err
}
func (c *FakeInitializerConfigurations) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 action := testing.NewRootDeleteCollectionAction(initializerconfigurationsResource, listOptions)
 _, err := c.Fake.Invokes(action, &admissionregistration.InitializerConfigurationList{})
 return err
}
func (c *FakeInitializerConfigurations) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *admissionregistration.InitializerConfiguration, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewRootPatchSubresourceAction(initializerconfigurationsResource, name, pt, data, subresources...), &admissionregistration.InitializerConfiguration{})
 if obj == nil {
  return nil, err
 }
 return obj.(*admissionregistration.InitializerConfiguration), err
}
