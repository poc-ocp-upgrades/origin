package controller

import (
 "hash/fnv"
 "sync"
 "github.com/golang/groupcache/lru"
 metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
 hashutil "k8s.io/kubernetes/pkg/util/hash"
)

type objectWithMeta interface{ metav1.Object }

func keyFunc(obj objectWithMeta) uint64 {
 _logClusterCodePath()
 defer _logClusterCodePath()
 hash := fnv.New32a()
 hashutil.DeepHashObject(hash, &equivalenceLabelObj{namespace: obj.GetNamespace(), labels: obj.GetLabels()})
 return uint64(hash.Sum32())
}

type equivalenceLabelObj struct {
 namespace string
 labels    map[string]string
}
type MatchingCache struct {
 mutex sync.RWMutex
 cache *lru.Cache
}

func NewMatchingCache(maxCacheEntries int) *MatchingCache {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &MatchingCache{cache: lru.New(maxCacheEntries)}
}
func (c *MatchingCache) Add(labelObj objectWithMeta, selectorObj objectWithMeta) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 key := keyFunc(labelObj)
 c.mutex.Lock()
 defer c.mutex.Unlock()
 c.cache.Add(key, selectorObj)
}
func (c *MatchingCache) GetMatchingObject(labelObj objectWithMeta) (controller interface{}, exists bool) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 key := keyFunc(labelObj)
 c.mutex.Lock()
 defer c.mutex.Unlock()
 return c.cache.Get(key)
}
func (c *MatchingCache) Update(labelObj objectWithMeta, selectorObj objectWithMeta) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 c.Add(labelObj, selectorObj)
}
func (c *MatchingCache) InvalidateAll() {
 _logClusterCodePath()
 defer _logClusterCodePath()
 c.mutex.Lock()
 defer c.mutex.Unlock()
 c.cache = lru.New(c.cache.MaxEntries)
}
