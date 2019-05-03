package fake

import (
 restclient "k8s.io/client-go/rest"
 core "k8s.io/client-go/testing"
)

func (c *FakeServices) ProxyGet(scheme, name, port, path string, params map[string]string) restclient.ResponseWrapper {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return c.Fake.InvokesProxy(core.NewProxyGetAction(servicesResource, c.ns, scheme, name, port, path, params))
}
