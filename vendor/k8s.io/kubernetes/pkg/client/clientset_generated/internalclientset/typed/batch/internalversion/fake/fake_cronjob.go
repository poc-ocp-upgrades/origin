package fake

import (
 v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
 labels "k8s.io/apimachinery/pkg/labels"
 schema "k8s.io/apimachinery/pkg/runtime/schema"
 types "k8s.io/apimachinery/pkg/types"
 watch "k8s.io/apimachinery/pkg/watch"
 testing "k8s.io/client-go/testing"
 batch "k8s.io/kubernetes/pkg/apis/batch"
)

type FakeCronJobs struct {
 Fake *FakeBatch
 ns   string
}

var cronjobsResource = schema.GroupVersionResource{Group: "batch", Version: "", Resource: "cronjobs"}
var cronjobsKind = schema.GroupVersionKind{Group: "batch", Version: "", Kind: "CronJob"}

func (c *FakeCronJobs) Get(name string, options v1.GetOptions) (result *batch.CronJob, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewGetAction(cronjobsResource, c.ns, name), &batch.CronJob{})
 if obj == nil {
  return nil, err
 }
 return obj.(*batch.CronJob), err
}
func (c *FakeCronJobs) List(opts v1.ListOptions) (result *batch.CronJobList, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewListAction(cronjobsResource, cronjobsKind, c.ns, opts), &batch.CronJobList{})
 if obj == nil {
  return nil, err
 }
 label, _, _ := testing.ExtractFromListOptions(opts)
 if label == nil {
  label = labels.Everything()
 }
 list := &batch.CronJobList{ListMeta: obj.(*batch.CronJobList).ListMeta}
 for _, item := range obj.(*batch.CronJobList).Items {
  if label.Matches(labels.Set(item.Labels)) {
   list.Items = append(list.Items, item)
  }
 }
 return list, err
}
func (c *FakeCronJobs) Watch(opts v1.ListOptions) (watch.Interface, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return c.Fake.InvokesWatch(testing.NewWatchAction(cronjobsResource, c.ns, opts))
}
func (c *FakeCronJobs) Create(cronJob *batch.CronJob) (result *batch.CronJob, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewCreateAction(cronjobsResource, c.ns, cronJob), &batch.CronJob{})
 if obj == nil {
  return nil, err
 }
 return obj.(*batch.CronJob), err
}
func (c *FakeCronJobs) Update(cronJob *batch.CronJob) (result *batch.CronJob, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewUpdateAction(cronjobsResource, c.ns, cronJob), &batch.CronJob{})
 if obj == nil {
  return nil, err
 }
 return obj.(*batch.CronJob), err
}
func (c *FakeCronJobs) UpdateStatus(cronJob *batch.CronJob) (*batch.CronJob, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewUpdateSubresourceAction(cronjobsResource, "status", c.ns, cronJob), &batch.CronJob{})
 if obj == nil {
  return nil, err
 }
 return obj.(*batch.CronJob), err
}
func (c *FakeCronJobs) Delete(name string, options *v1.DeleteOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 _, err := c.Fake.Invokes(testing.NewDeleteAction(cronjobsResource, c.ns, name), &batch.CronJob{})
 return err
}
func (c *FakeCronJobs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 action := testing.NewDeleteCollectionAction(cronjobsResource, c.ns, listOptions)
 _, err := c.Fake.Invokes(action, &batch.CronJobList{})
 return err
}
func (c *FakeCronJobs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *batch.CronJob, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewPatchSubresourceAction(cronjobsResource, c.ns, name, pt, data, subresources...), &batch.CronJob{})
 if obj == nil {
  return nil, err
 }
 return obj.(*batch.CronJob), err
}
