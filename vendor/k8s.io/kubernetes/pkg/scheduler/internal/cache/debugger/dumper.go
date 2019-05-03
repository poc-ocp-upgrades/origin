package debugger

import (
 "fmt"
 "strings"
 "k8s.io/klog"
 "k8s.io/api/core/v1"
 "k8s.io/kubernetes/pkg/scheduler/cache"
 internalcache "k8s.io/kubernetes/pkg/scheduler/internal/cache"
 "k8s.io/kubernetes/pkg/scheduler/internal/queue"
)

type CacheDumper struct {
 cache    internalcache.Cache
 podQueue queue.SchedulingQueue
}

func (d *CacheDumper) DumpAll() {
 _logClusterCodePath()
 defer _logClusterCodePath()
 d.dumpNodes()
 d.dumpSchedulingQueue()
}
func (d *CacheDumper) dumpNodes() {
 _logClusterCodePath()
 defer _logClusterCodePath()
 snapshot := d.cache.Snapshot()
 klog.Info("Dump of cached NodeInfo")
 for _, nodeInfo := range snapshot.Nodes {
  klog.Info(printNodeInfo(nodeInfo))
 }
}
func (d *CacheDumper) dumpSchedulingQueue() {
 _logClusterCodePath()
 defer _logClusterCodePath()
 waitingPods := d.podQueue.WaitingPods()
 var podData strings.Builder
 for _, p := range waitingPods {
  podData.WriteString(printPod(p))
 }
 klog.Infof("Dump of scheduling queue:\n%s", podData.String())
}
func printNodeInfo(n *cache.NodeInfo) string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var nodeData strings.Builder
 nodeData.WriteString(fmt.Sprintf("\nNode name: %+v\nRequested Resources: %+v\nAllocatable Resources:%+v\nNumber of Pods: %v\nPods:\n", n.Node().Name, n.RequestedResource(), n.AllocatableResource(), len(n.Pods())))
 for _, p := range n.Pods() {
  nodeData.WriteString(printPod(p))
 }
 return nodeData.String()
}
func printPod(p *v1.Pod) string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return fmt.Sprintf("name: %v, namespace: %v, uid: %v, phase: %v, nominated node: %v\n", p.Name, p.Namespace, p.UID, p.Status.Phase, p.Status.NominatedNodeName)
}
