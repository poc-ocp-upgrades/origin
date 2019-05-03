package garbagecollector

import (
 "sync"
 "github.com/golang/groupcache/lru"
 "k8s.io/apimachinery/pkg/types"
)

type UIDCache struct {
 mutex sync.Mutex
 cache *lru.Cache
}

func NewUIDCache(maxCacheEntries int) *UIDCache {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &UIDCache{cache: lru.New(maxCacheEntries)}
}
func (c *UIDCache) Add(uid types.UID) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 c.mutex.Lock()
 defer c.mutex.Unlock()
 c.cache.Add(uid, nil)
}
func (c *UIDCache) Has(uid types.UID) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 c.mutex.Lock()
 defer c.mutex.Unlock()
 _, found := c.cache.Get(uid)
 return found
}
