package controllers

import (
	lru "github.com/hashicorp/golang-lru"
	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apiserver/pkg/storage/etcd"
	"k8s.io/client-go/tools/cache"
	"k8s.io/kubernetes/pkg/controller"
)

type MutationCache interface {
	GetByKey(key string) (interface{}, bool, error)
	Mutation(interface{})
}
type ResourceVersionComparator interface {
	CompareResourceVersion(lhs, rhs runtime.Object) int
}

func NewEtcdMutationCache(backingCache cache.Store) MutationCache {
	_logClusterCodePath()
	defer _logClusterCodePath()
	lru, err := lru.New(100)
	if err != nil {
		panic(err)
	}
	return &mutationCache{backingCache: backingCache, mutationCache: lru, comparator: etcd.APIObjectVersioner{}}
}

type mutationCache struct {
	backingCache  cache.Store
	mutationCache *lru.Cache
	comparator    ResourceVersionComparator
}

func (c *mutationCache) GetByKey(key string) (interface{}, bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, exists, err := c.backingCache.GetByKey(key)
	if err != nil {
		return nil, false, err
	}
	if !exists {
		return nil, false, nil
	}
	objRuntime, ok := obj.(runtime.Object)
	if !ok {
		return obj, true, nil
	}
	mutatedObj, exists := c.mutationCache.Get(key)
	if !exists {
		return obj, true, nil
	}
	mutatedObjRuntime, ok := mutatedObj.(runtime.Object)
	if !ok {
		return obj, true, nil
	}
	if c.comparator.CompareResourceVersion(objRuntime, mutatedObjRuntime) >= 0 {
		c.mutationCache.Remove(key)
		return obj, true, nil
	}
	return mutatedObj, true, nil
}
func (c *mutationCache) Mutation(obj interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	key, err := controller.KeyFunc(obj)
	if err != nil {
		utilruntime.HandleError(err)
		return
	}
	c.mutationCache.Add(key, obj)
}
