package internalversion

import (
 "k8s.io/apimachinery/pkg/util/net"
 restclient "k8s.io/client-go/rest"
)

type ServiceExpansion interface {
 ProxyGet(scheme, name, port, path string, params map[string]string) restclient.ResponseWrapper
}

func (c *services) ProxyGet(scheme, name, port, path string, params map[string]string) restclient.ResponseWrapper {
 _logClusterCodePath()
 defer _logClusterCodePath()
 request := c.client.Get().Namespace(c.ns).Resource("services").SubResource("proxy").Name(net.JoinSchemeNamePort(scheme, name, port)).Suffix(path)
 for k, v := range params {
  request = request.Param(k, v)
 }
 return request
}
