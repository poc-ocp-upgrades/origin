package fake

import (
 v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
 labels "k8s.io/apimachinery/pkg/labels"
 schema "k8s.io/apimachinery/pkg/runtime/schema"
 types "k8s.io/apimachinery/pkg/types"
 watch "k8s.io/apimachinery/pkg/watch"
 testing "k8s.io/client-go/testing"
 auditregistration "k8s.io/kubernetes/pkg/apis/auditregistration"
)

type FakeAuditSinks struct{ Fake *FakeAuditregistration }

var auditsinksResource = schema.GroupVersionResource{Group: "auditregistration.k8s.io", Version: "", Resource: "auditsinks"}
var auditsinksKind = schema.GroupVersionKind{Group: "auditregistration.k8s.io", Version: "", Kind: "AuditSink"}

func (c *FakeAuditSinks) Get(name string, options v1.GetOptions) (result *auditregistration.AuditSink, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewRootGetAction(auditsinksResource, name), &auditregistration.AuditSink{})
 if obj == nil {
  return nil, err
 }
 return obj.(*auditregistration.AuditSink), err
}
func (c *FakeAuditSinks) List(opts v1.ListOptions) (result *auditregistration.AuditSinkList, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewRootListAction(auditsinksResource, auditsinksKind, opts), &auditregistration.AuditSinkList{})
 if obj == nil {
  return nil, err
 }
 label, _, _ := testing.ExtractFromListOptions(opts)
 if label == nil {
  label = labels.Everything()
 }
 list := &auditregistration.AuditSinkList{ListMeta: obj.(*auditregistration.AuditSinkList).ListMeta}
 for _, item := range obj.(*auditregistration.AuditSinkList).Items {
  if label.Matches(labels.Set(item.Labels)) {
   list.Items = append(list.Items, item)
  }
 }
 return list, err
}
func (c *FakeAuditSinks) Watch(opts v1.ListOptions) (watch.Interface, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return c.Fake.InvokesWatch(testing.NewRootWatchAction(auditsinksResource, opts))
}
func (c *FakeAuditSinks) Create(auditSink *auditregistration.AuditSink) (result *auditregistration.AuditSink, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewRootCreateAction(auditsinksResource, auditSink), &auditregistration.AuditSink{})
 if obj == nil {
  return nil, err
 }
 return obj.(*auditregistration.AuditSink), err
}
func (c *FakeAuditSinks) Update(auditSink *auditregistration.AuditSink) (result *auditregistration.AuditSink, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewRootUpdateAction(auditsinksResource, auditSink), &auditregistration.AuditSink{})
 if obj == nil {
  return nil, err
 }
 return obj.(*auditregistration.AuditSink), err
}
func (c *FakeAuditSinks) Delete(name string, options *v1.DeleteOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 _, err := c.Fake.Invokes(testing.NewRootDeleteAction(auditsinksResource, name), &auditregistration.AuditSink{})
 return err
}
func (c *FakeAuditSinks) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 action := testing.NewRootDeleteCollectionAction(auditsinksResource, listOptions)
 _, err := c.Fake.Invokes(action, &auditregistration.AuditSinkList{})
 return err
}
func (c *FakeAuditSinks) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *auditregistration.AuditSink, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewRootPatchSubresourceAction(auditsinksResource, name, pt, data, subresources...), &auditregistration.AuditSink{})
 if obj == nil {
  return nil, err
 }
 return obj.(*auditregistration.AuditSink), err
}
