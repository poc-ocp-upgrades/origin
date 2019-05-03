package azure

import (
 "fmt"
 "sync"
 "time"
 "k8s.io/client-go/tools/cache"
)

type getFunc func(key string) (interface{}, error)
type cacheEntry struct {
 key  string
 data interface{}
 lock sync.Mutex
}

func cacheKeyFunc(obj interface{}) (string, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return obj.(*cacheEntry).key, nil
}

type timedCache struct {
 store  cache.Store
 lock   sync.Mutex
 getter getFunc
}

func newTimedcache(ttl time.Duration, getter getFunc) (*timedCache, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if getter == nil {
  return nil, fmt.Errorf("getter is not provided")
 }
 return &timedCache{getter: getter, store: cache.NewTTLStore(cacheKeyFunc, ttl)}, nil
}
func (t *timedCache) getInternal(key string) (*cacheEntry, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 entry, exists, err := t.store.GetByKey(key)
 if err != nil {
  return nil, err
 }
 if exists {
  return entry.(*cacheEntry), nil
 }
 t.lock.Lock()
 defer t.lock.Unlock()
 entry, exists, err = t.store.GetByKey(key)
 if err != nil {
  return nil, err
 }
 if exists {
  return entry.(*cacheEntry), nil
 }
 newEntry := &cacheEntry{key: key, data: nil}
 t.store.Add(newEntry)
 return newEntry, nil
}
func (t *timedCache) Get(key string) (interface{}, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 entry, err := t.getInternal(key)
 if err != nil {
  return nil, err
 }
 if entry.data == nil {
  entry.lock.Lock()
  defer entry.lock.Unlock()
  if entry.data == nil {
   data, err := t.getter(key)
   if err != nil {
    return nil, err
   }
   entry.data = data
  }
 }
 return entry.data, nil
}
func (t *timedCache) Delete(key string) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return t.store.Delete(&cacheEntry{key: key})
}
