package cronjob

import (
 "fmt"
 "sync"
 batchv1 "k8s.io/api/batch/v1"
 batchv1beta1 "k8s.io/api/batch/v1beta1"
 "k8s.io/api/core/v1"
 metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
 "k8s.io/apimachinery/pkg/labels"
 "k8s.io/apimachinery/pkg/types"
 clientset "k8s.io/client-go/kubernetes"
 "k8s.io/client-go/tools/record"
)

type sjControlInterface interface {
 UpdateStatus(sj *batchv1beta1.CronJob) (*batchv1beta1.CronJob, error)
}
type realSJControl struct{ KubeClient clientset.Interface }

var _ sjControlInterface = &realSJControl{}

func (c *realSJControl) UpdateStatus(sj *batchv1beta1.CronJob) (*batchv1beta1.CronJob, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return c.KubeClient.BatchV1beta1().CronJobs(sj.Namespace).UpdateStatus(sj)
}

type fakeSJControl struct{ Updates []batchv1beta1.CronJob }

var _ sjControlInterface = &fakeSJControl{}

func (c *fakeSJControl) UpdateStatus(sj *batchv1beta1.CronJob) (*batchv1beta1.CronJob, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 c.Updates = append(c.Updates, *sj)
 return sj, nil
}

type jobControlInterface interface {
 GetJob(namespace, name string) (*batchv1.Job, error)
 CreateJob(namespace string, job *batchv1.Job) (*batchv1.Job, error)
 UpdateJob(namespace string, job *batchv1.Job) (*batchv1.Job, error)
 PatchJob(namespace string, name string, pt types.PatchType, data []byte, subresources ...string) (*batchv1.Job, error)
 DeleteJob(namespace string, name string) error
}
type realJobControl struct {
 KubeClient clientset.Interface
 Recorder   record.EventRecorder
}

var _ jobControlInterface = &realJobControl{}

func copyLabels(template *batchv1beta1.JobTemplateSpec) labels.Set {
 _logClusterCodePath()
 defer _logClusterCodePath()
 l := make(labels.Set)
 for k, v := range template.Labels {
  l[k] = v
 }
 return l
}
func copyAnnotations(template *batchv1beta1.JobTemplateSpec) labels.Set {
 _logClusterCodePath()
 defer _logClusterCodePath()
 a := make(labels.Set)
 for k, v := range template.Annotations {
  a[k] = v
 }
 return a
}
func (r realJobControl) GetJob(namespace, name string) (*batchv1.Job, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return r.KubeClient.BatchV1().Jobs(namespace).Get(name, metav1.GetOptions{})
}
func (r realJobControl) UpdateJob(namespace string, job *batchv1.Job) (*batchv1.Job, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return r.KubeClient.BatchV1().Jobs(namespace).Update(job)
}
func (r realJobControl) PatchJob(namespace string, name string, pt types.PatchType, data []byte, subresources ...string) (*batchv1.Job, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return r.KubeClient.BatchV1().Jobs(namespace).Patch(name, pt, data, subresources...)
}
func (r realJobControl) CreateJob(namespace string, job *batchv1.Job) (*batchv1.Job, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return r.KubeClient.BatchV1().Jobs(namespace).Create(job)
}
func (r realJobControl) DeleteJob(namespace string, name string) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 background := metav1.DeletePropagationBackground
 return r.KubeClient.BatchV1().Jobs(namespace).Delete(name, &metav1.DeleteOptions{PropagationPolicy: &background})
}

type fakeJobControl struct {
 sync.Mutex
 Job           *batchv1.Job
 Jobs          []batchv1.Job
 DeleteJobName []string
 Err           error
 UpdateJobName []string
 PatchJobName  []string
 Patches       [][]byte
}

var _ jobControlInterface = &fakeJobControl{}

func (f *fakeJobControl) CreateJob(namespace string, job *batchv1.Job) (*batchv1.Job, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 f.Lock()
 defer f.Unlock()
 if f.Err != nil {
  return nil, f.Err
 }
 job.SelfLink = fmt.Sprintf("/api/batch/v1/namespaces/%s/jobs/%s", namespace, job.Name)
 f.Jobs = append(f.Jobs, *job)
 job.UID = "test-uid"
 return job, nil
}
func (f *fakeJobControl) GetJob(namespace, name string) (*batchv1.Job, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 f.Lock()
 defer f.Unlock()
 if f.Err != nil {
  return nil, f.Err
 }
 return f.Job, nil
}
func (f *fakeJobControl) UpdateJob(namespace string, job *batchv1.Job) (*batchv1.Job, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 f.Lock()
 defer f.Unlock()
 if f.Err != nil {
  return nil, f.Err
 }
 f.UpdateJobName = append(f.UpdateJobName, job.Name)
 return job, nil
}
func (f *fakeJobControl) PatchJob(namespace string, name string, pt types.PatchType, data []byte, subresources ...string) (*batchv1.Job, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 f.Lock()
 defer f.Unlock()
 if f.Err != nil {
  return nil, f.Err
 }
 f.PatchJobName = append(f.PatchJobName, name)
 f.Patches = append(f.Patches, data)
 return &batchv1.Job{}, nil
}
func (f *fakeJobControl) DeleteJob(namespace string, name string) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 f.Lock()
 defer f.Unlock()
 if f.Err != nil {
  return f.Err
 }
 f.DeleteJobName = append(f.DeleteJobName, name)
 return nil
}
func (f *fakeJobControl) Clear() {
 _logClusterCodePath()
 defer _logClusterCodePath()
 f.Lock()
 defer f.Unlock()
 f.DeleteJobName = []string{}
 f.Jobs = []batchv1.Job{}
 f.Err = nil
}

type podControlInterface interface {
 ListPods(namespace string, opts metav1.ListOptions) (*v1.PodList, error)
 DeletePod(namespace string, name string) error
}
type realPodControl struct {
 KubeClient clientset.Interface
 Recorder   record.EventRecorder
}

var _ podControlInterface = &realPodControl{}

func (r realPodControl) ListPods(namespace string, opts metav1.ListOptions) (*v1.PodList, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return r.KubeClient.CoreV1().Pods(namespace).List(opts)
}
func (r realPodControl) DeletePod(namespace string, name string) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return r.KubeClient.CoreV1().Pods(namespace).Delete(name, nil)
}

type fakePodControl struct {
 sync.Mutex
 Pods          []v1.Pod
 DeletePodName []string
 Err           error
}

var _ podControlInterface = &fakePodControl{}

func (f *fakePodControl) ListPods(namespace string, opts metav1.ListOptions) (*v1.PodList, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 f.Lock()
 defer f.Unlock()
 if f.Err != nil {
  return nil, f.Err
 }
 return &v1.PodList{Items: f.Pods}, nil
}
func (f *fakePodControl) DeletePod(namespace string, name string) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 f.Lock()
 defer f.Unlock()
 if f.Err != nil {
  return f.Err
 }
 f.DeletePodName = append(f.DeletePodName, name)
 return nil
}
