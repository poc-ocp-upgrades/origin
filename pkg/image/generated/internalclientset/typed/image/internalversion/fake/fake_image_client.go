package fake

import (
	internalversion "github.com/openshift/origin/pkg/image/generated/internalclientset/typed/image/internalversion"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeImage struct{ *testing.Fake }

func (c *FakeImage) Images() internalversion.ImageResourceInterface {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &FakeImages{c}
}
func (c *FakeImage) ImageSignatures() internalversion.ImageSignatureInterface {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &FakeImageSignatures{c}
}
func (c *FakeImage) ImageStreams(namespace string) internalversion.ImageStreamInterface {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &FakeImageStreams{c, namespace}
}
func (c *FakeImage) ImageStreamImages(namespace string) internalversion.ImageStreamImageInterface {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &FakeImageStreamImages{c, namespace}
}
func (c *FakeImage) ImageStreamImports(namespace string) internalversion.ImageStreamImportInterface {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &FakeImageStreamImports{c, namespace}
}
func (c *FakeImage) ImageStreamMappings(namespace string) internalversion.ImageStreamMappingInterface {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &FakeImageStreamMappings{c, namespace}
}
func (c *FakeImage) ImageStreamTags(namespace string) internalversion.ImageStreamTagInterface {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &FakeImageStreamTags{c, namespace}
}
func (c *FakeImage) RESTClient() rest.Interface {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var ret *rest.RESTClient
	return ret
}
