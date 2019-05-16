package garbagecollector

import (
	"github.com/golang/groupcache/lru"
	"k8s.io/apimachinery/pkg/types"
	"sync"
)

type UIDCache struct {
	mutex sync.Mutex
	cache *lru.Cache
}

func NewUIDCache(maxCacheEntries int) *UIDCache {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &UIDCache{cache: lru.New(maxCacheEntries)}
}
func (c *UIDCache) Add(uid types.UID) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.cache.Add(uid, nil)
}
func (c *UIDCache) Has(uid types.UID) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	c.mutex.Lock()
	defer c.mutex.Unlock()
	_, found := c.cache.Get(uid)
	return found
}
