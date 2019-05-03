package debugger

import (
 corelisters "k8s.io/client-go/listers/core/v1"
 internalcache "k8s.io/kubernetes/pkg/scheduler/internal/cache"
 internalqueue "k8s.io/kubernetes/pkg/scheduler/internal/queue"
)

type CacheDebugger struct {
 Comparer CacheComparer
 Dumper   CacheDumper
}

func New(nodeLister corelisters.NodeLister, podLister corelisters.PodLister, cache internalcache.Cache, podQueue internalqueue.SchedulingQueue) *CacheDebugger {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &CacheDebugger{Comparer: CacheComparer{NodeLister: nodeLister, PodLister: podLister, Cache: cache, PodQueue: podQueue}, Dumper: CacheDumper{cache: cache, podQueue: podQueue}}
}
