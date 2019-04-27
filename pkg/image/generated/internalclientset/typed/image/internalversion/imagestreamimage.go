package internalversion

import (
	image "github.com/openshift/origin/pkg/image/apis/image"
	scheme "github.com/openshift/origin/pkg/image/generated/internalclientset/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	rest "k8s.io/client-go/rest"
)

type ImageStreamImagesGetter interface {
	ImageStreamImages(namespace string) ImageStreamImageInterface
}
type ImageStreamImageInterface interface {
	Get(name string, options v1.GetOptions) (*image.ImageStreamImage, error)
	ImageStreamImageExpansion
}
type imageStreamImages struct {
	client	rest.Interface
	ns	string
}

func newImageStreamImages(c *ImageClient, namespace string) *imageStreamImages {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &imageStreamImages{client: c.RESTClient(), ns: namespace}
}
func (c *imageStreamImages) Get(name string, options v1.GetOptions) (result *image.ImageStreamImage, err error) {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	result = &image.ImageStreamImage{}
	err = c.client.Get().Namespace(c.ns).Resource("imagestreamimages").Name(name).VersionedParams(&options, scheme.ParameterCodec).Do().Into(result)
	return
}
