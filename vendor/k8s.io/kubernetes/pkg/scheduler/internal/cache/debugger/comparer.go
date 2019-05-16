package debugger

import (
	goformat "fmt"
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/labels"
	corelisters "k8s.io/client-go/listers/core/v1"
	"k8s.io/klog"
	schedulercache "k8s.io/kubernetes/pkg/scheduler/cache"
	schedulerinternalcache "k8s.io/kubernetes/pkg/scheduler/internal/cache"
	internalqueue "k8s.io/kubernetes/pkg/scheduler/internal/queue"
	goos "os"
	godefaultruntime "runtime"
	"sort"
	"strings"
	gotime "time"
)

type CacheComparer struct {
	NodeLister corelisters.NodeLister
	PodLister  corelisters.PodLister
	Cache      schedulerinternalcache.Cache
	PodQueue   internalqueue.SchedulingQueue
}

func (c *CacheComparer) Compare() error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	klog.V(3).Info("cache comparer started")
	defer klog.V(3).Info("cache comparer finished")
	nodes, err := c.NodeLister.List(labels.Everything())
	if err != nil {
		return err
	}
	pods, err := c.PodLister.List(labels.Everything())
	if err != nil {
		return err
	}
	snapshot := c.Cache.Snapshot()
	waitingPods := c.PodQueue.WaitingPods()
	if missed, redundant := c.CompareNodes(nodes, snapshot.Nodes); len(missed)+len(redundant) != 0 {
		klog.Warningf("cache mismatch: missed nodes: %s; redundant nodes: %s", missed, redundant)
	}
	if missed, redundant := c.ComparePods(pods, waitingPods, snapshot.Nodes); len(missed)+len(redundant) != 0 {
		klog.Warningf("cache mismatch: missed pods: %s; redundant pods: %s", missed, redundant)
	}
	return nil
}
func (c *CacheComparer) CompareNodes(nodes []*v1.Node, nodeinfos map[string]*schedulercache.NodeInfo) (missed, redundant []string) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	actual := []string{}
	for _, node := range nodes {
		actual = append(actual, node.Name)
	}
	cached := []string{}
	for nodeName := range nodeinfos {
		cached = append(cached, nodeName)
	}
	return compareStrings(actual, cached)
}
func (c *CacheComparer) ComparePods(pods, waitingPods []*v1.Pod, nodeinfos map[string]*schedulercache.NodeInfo) (missed, redundant []string) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	actual := []string{}
	for _, pod := range pods {
		actual = append(actual, string(pod.UID))
	}
	cached := []string{}
	for _, nodeinfo := range nodeinfos {
		for _, pod := range nodeinfo.Pods() {
			cached = append(cached, string(pod.UID))
		}
	}
	for _, pod := range waitingPods {
		cached = append(cached, string(pod.UID))
	}
	return compareStrings(actual, cached)
}
func compareStrings(actual, cached []string) (missed, redundant []string) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	missed, redundant = []string{}, []string{}
	sort.Strings(actual)
	sort.Strings(cached)
	compare := func(i, j int) int {
		if i == len(actual) {
			return 1
		} else if j == len(cached) {
			return -1
		}
		return strings.Compare(actual[i], cached[j])
	}
	for i, j := 0, 0; i < len(actual) || j < len(cached); {
		switch compare(i, j) {
		case 0:
			i++
			j++
		case -1:
			missed = append(missed, actual[i])
			i++
		case 1:
			redundant = append(redundant, cached[j])
			j++
		}
	}
	return
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
