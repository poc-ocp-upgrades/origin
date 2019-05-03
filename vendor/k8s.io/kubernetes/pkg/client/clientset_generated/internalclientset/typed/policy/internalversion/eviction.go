package internalversion

import (
 rest "k8s.io/client-go/rest"
)

type EvictionsGetter interface {
 Evictions(namespace string) EvictionInterface
}
type EvictionInterface interface{ EvictionExpansion }
type evictions struct {
 client rest.Interface
 ns     string
}

func newEvictions(c *PolicyClient, namespace string) *evictions {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &evictions{client: c.RESTClient(), ns: namespace}
}
