package controller

import (
	"github.com/golang/groupcache/lru"
	"hash/fnv"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	hashutil "k8s.io/kubernetes/pkg/util/hash"
	"sync"
)

type objectWithMeta interface{ metav1.Object }

func keyFunc(obj objectWithMeta) uint64 {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
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
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &MatchingCache{cache: lru.New(maxCacheEntries)}
}
func (c *MatchingCache) Add(labelObj objectWithMeta, selectorObj objectWithMeta) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	key := keyFunc(labelObj)
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.cache.Add(key, selectorObj)
}
func (c *MatchingCache) GetMatchingObject(labelObj objectWithMeta) (controller interface{}, exists bool) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	key := keyFunc(labelObj)
	c.mutex.Lock()
	defer c.mutex.Unlock()
	return c.cache.Get(key)
}
func (c *MatchingCache) Update(labelObj objectWithMeta, selectorObj objectWithMeta) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	c.Add(labelObj, selectorObj)
}
func (c *MatchingCache) InvalidateAll() {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.cache = lru.New(c.cache.MaxEntries)
}
