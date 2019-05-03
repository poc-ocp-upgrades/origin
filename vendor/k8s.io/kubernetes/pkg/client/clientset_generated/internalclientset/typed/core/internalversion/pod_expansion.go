package internalversion

import (
 restclient "k8s.io/client-go/rest"
 "k8s.io/kubernetes/pkg/api/legacyscheme"
 api "k8s.io/kubernetes/pkg/apis/core"
)

type PodExpansion interface {
 Bind(binding *api.Binding) error
 GetLogs(name string, opts *api.PodLogOptions) *restclient.Request
}

func (c *pods) Bind(binding *api.Binding) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return c.client.Post().Namespace(c.ns).Resource("pods").Name(binding.Name).SubResource("binding").Body(binding).Do().Error()
}
func (c *pods) GetLogs(name string, opts *api.PodLogOptions) *restclient.Request {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return c.client.Get().Namespace(c.ns).Name(name).Resource("pods").SubResource("log").VersionedParams(opts, legacyscheme.ParameterCodec)
}
