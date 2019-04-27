package internalversion

import (
	image "github.com/openshift/origin/pkg/image/apis/image"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	rest "k8s.io/client-go/rest"
)

type ImageStreamMappingsGetter interface {
	ImageStreamMappings(namespace string) ImageStreamMappingInterface
}
type ImageStreamMappingInterface interface {
	Create(*image.ImageStreamMapping) (*v1.Status, error)
	ImageStreamMappingExpansion
}
type imageStreamMappings struct {
	client	rest.Interface
	ns	string
}

func newImageStreamMappings(c *ImageClient, namespace string) *imageStreamMappings {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &imageStreamMappings{client: c.RESTClient(), ns: namespace}
}
func (c *imageStreamMappings) Create(imageStreamMapping *image.ImageStreamMapping) (result *v1.Status, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	result = &v1.Status{}
	err = c.client.Post().Namespace(c.ns).Resource("imagestreammappings").Body(imageStreamMapping).Do().Into(result)
	return
}
