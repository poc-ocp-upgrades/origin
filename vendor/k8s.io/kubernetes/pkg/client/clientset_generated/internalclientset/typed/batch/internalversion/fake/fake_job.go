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

type FakeJobs struct {
 Fake *FakeBatch
 ns   string
}

var jobsResource = schema.GroupVersionResource{Group: "batch", Version: "", Resource: "jobs"}
var jobsKind = schema.GroupVersionKind{Group: "batch", Version: "", Kind: "Job"}

func (c *FakeJobs) Get(name string, options v1.GetOptions) (result *batch.Job, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewGetAction(jobsResource, c.ns, name), &batch.Job{})
 if obj == nil {
  return nil, err
 }
 return obj.(*batch.Job), err
}
func (c *FakeJobs) List(opts v1.ListOptions) (result *batch.JobList, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewListAction(jobsResource, jobsKind, c.ns, opts), &batch.JobList{})
 if obj == nil {
  return nil, err
 }
 label, _, _ := testing.ExtractFromListOptions(opts)
 if label == nil {
  label = labels.Everything()
 }
 list := &batch.JobList{ListMeta: obj.(*batch.JobList).ListMeta}
 for _, item := range obj.(*batch.JobList).Items {
  if label.Matches(labels.Set(item.Labels)) {
   list.Items = append(list.Items, item)
  }
 }
 return list, err
}
func (c *FakeJobs) Watch(opts v1.ListOptions) (watch.Interface, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return c.Fake.InvokesWatch(testing.NewWatchAction(jobsResource, c.ns, opts))
}
func (c *FakeJobs) Create(job *batch.Job) (result *batch.Job, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewCreateAction(jobsResource, c.ns, job), &batch.Job{})
 if obj == nil {
  return nil, err
 }
 return obj.(*batch.Job), err
}
func (c *FakeJobs) Update(job *batch.Job) (result *batch.Job, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewUpdateAction(jobsResource, c.ns, job), &batch.Job{})
 if obj == nil {
  return nil, err
 }
 return obj.(*batch.Job), err
}
func (c *FakeJobs) UpdateStatus(job *batch.Job) (*batch.Job, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewUpdateSubresourceAction(jobsResource, "status", c.ns, job), &batch.Job{})
 if obj == nil {
  return nil, err
 }
 return obj.(*batch.Job), err
}
func (c *FakeJobs) Delete(name string, options *v1.DeleteOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 _, err := c.Fake.Invokes(testing.NewDeleteAction(jobsResource, c.ns, name), &batch.Job{})
 return err
}
func (c *FakeJobs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 action := testing.NewDeleteCollectionAction(jobsResource, c.ns, listOptions)
 _, err := c.Fake.Invokes(action, &batch.JobList{})
 return err
}
func (c *FakeJobs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *batch.Job, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := c.Fake.Invokes(testing.NewPatchSubresourceAction(jobsResource, c.ns, name, pt, data, subresources...), &batch.Job{})
 if obj == nil {
  return nil, err
 }
 return obj.(*batch.Job), err
}
