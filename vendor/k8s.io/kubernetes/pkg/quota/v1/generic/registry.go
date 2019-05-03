package generic

import (
 "sync"
 "k8s.io/apimachinery/pkg/runtime/schema"
 quota "k8s.io/kubernetes/pkg/quota/v1"
)

type simpleRegistry struct {
 lock       sync.RWMutex
 evaluators map[schema.GroupResource]quota.Evaluator
}

func NewRegistry(evaluators []quota.Evaluator) quota.Registry {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &simpleRegistry{evaluators: evaluatorsByGroupResource(evaluators)}
}
func (r *simpleRegistry) Add(e quota.Evaluator) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 r.lock.Lock()
 defer r.lock.Unlock()
 r.evaluators[e.GroupResource()] = e
}
func (r *simpleRegistry) Remove(e quota.Evaluator) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 r.lock.Lock()
 defer r.lock.Unlock()
 delete(r.evaluators, e.GroupResource())
}
func (r *simpleRegistry) Get(gr schema.GroupResource) quota.Evaluator {
 _logClusterCodePath()
 defer _logClusterCodePath()
 r.lock.RLock()
 defer r.lock.RUnlock()
 return r.evaluators[gr]
}
func (r *simpleRegistry) List() []quota.Evaluator {
 _logClusterCodePath()
 defer _logClusterCodePath()
 r.lock.RLock()
 defer r.lock.RUnlock()
 return evaluatorsList(r.evaluators)
}
func evaluatorsByGroupResource(items []quota.Evaluator) map[schema.GroupResource]quota.Evaluator {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result := map[schema.GroupResource]quota.Evaluator{}
 for _, item := range items {
  result[item.GroupResource()] = item
 }
 return result
}
func evaluatorsList(input map[schema.GroupResource]quota.Evaluator) []quota.Evaluator {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var result []quota.Evaluator
 for _, item := range input {
  result = append(result, item)
 }
 return result
}
