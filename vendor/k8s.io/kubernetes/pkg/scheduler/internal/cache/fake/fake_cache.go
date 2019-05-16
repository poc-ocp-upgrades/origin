package fake

import (
	goformat "fmt"
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/kubernetes/pkg/scheduler/algorithm"
	schedulercache "k8s.io/kubernetes/pkg/scheduler/cache"
	schedulerinternalcache "k8s.io/kubernetes/pkg/scheduler/internal/cache"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

type Cache struct {
	AssumeFunc       func(*v1.Pod)
	ForgetFunc       func(*v1.Pod)
	IsAssumedPodFunc func(*v1.Pod) bool
	GetPodFunc       func(*v1.Pod) *v1.Pod
}

func (c *Cache) AssumePod(pod *v1.Pod) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	c.AssumeFunc(pod)
	return nil
}
func (c *Cache) FinishBinding(pod *v1.Pod) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return nil
}
func (c *Cache) ForgetPod(pod *v1.Pod) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	c.ForgetFunc(pod)
	return nil
}
func (c *Cache) AddPod(pod *v1.Pod) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return nil
}
func (c *Cache) UpdatePod(oldPod, newPod *v1.Pod) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return nil
}
func (c *Cache) RemovePod(pod *v1.Pod) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return nil
}
func (c *Cache) IsAssumedPod(pod *v1.Pod) (bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return c.IsAssumedPodFunc(pod), nil
}
func (c *Cache) GetPod(pod *v1.Pod) (*v1.Pod, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return c.GetPodFunc(pod), nil
}
func (c *Cache) AddNode(node *v1.Node) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return nil
}
func (c *Cache) UpdateNode(oldNode, newNode *v1.Node) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return nil
}
func (c *Cache) RemoveNode(node *v1.Node) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return nil
}
func (c *Cache) UpdateNodeNameToInfoMap(infoMap map[string]*schedulercache.NodeInfo) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return nil
}
func (c *Cache) List(s labels.Selector) ([]*v1.Pod, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return nil, nil
}
func (c *Cache) FilteredList(filter algorithm.PodFilter, selector labels.Selector) ([]*v1.Pod, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return nil, nil
}
func (c *Cache) Snapshot() *schedulerinternalcache.Snapshot {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &schedulerinternalcache.Snapshot{}
}
func (c *Cache) NodeTree() *schedulerinternalcache.NodeTree {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return nil
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
