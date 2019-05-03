package internalversion

import (
 policy "k8s.io/kubernetes/pkg/apis/policy"
)

type EvictionExpansion interface {
 Evict(eviction *policy.Eviction) error
}

func (c *evictions) Evict(eviction *policy.Eviction) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return c.client.Post().AbsPath("/api/v1").Namespace(eviction.Namespace).Resource("pods").Name(eviction.Name).SubResource("eviction").Body(eviction).Do().Error()
}
