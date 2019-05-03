package internalversion

import (
 api "k8s.io/kubernetes/pkg/apis/core"
)

type NamespaceExpansion interface {
 Finalize(item *api.Namespace) (*api.Namespace, error)
}

func (c *namespaces) Finalize(namespace *api.Namespace) (result *api.Namespace, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result = &api.Namespace{}
 err = c.client.Put().Resource("namespaces").Name(namespace.Name).SubResource("finalize").Body(namespace).Do().Into(result)
 return
}
