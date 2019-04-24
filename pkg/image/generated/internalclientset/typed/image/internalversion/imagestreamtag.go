package internalversion

import (
	image "github.com/openshift/origin/pkg/image/apis/image"
	scheme "github.com/openshift/origin/pkg/image/generated/internalclientset/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	rest "k8s.io/client-go/rest"
)

type ImageStreamTagsGetter interface {
	ImageStreamTags(namespace string) ImageStreamTagInterface
}
type ImageStreamTagInterface interface {
	Create(*image.ImageStreamTag) (*image.ImageStreamTag, error)
	Update(*image.ImageStreamTag) (*image.ImageStreamTag, error)
	Delete(name string, options *v1.DeleteOptions) error
	Get(name string, options v1.GetOptions) (*image.ImageStreamTag, error)
	ImageStreamTagExpansion
}
type imageStreamTags struct {
	client	rest.Interface
	ns	string
}

func newImageStreamTags(c *ImageClient, namespace string) *imageStreamTags {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &imageStreamTags{client: c.RESTClient(), ns: namespace}
}
func (c *imageStreamTags) Get(name string, options v1.GetOptions) (result *image.ImageStreamTag, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	result = &image.ImageStreamTag{}
	err = c.client.Get().Namespace(c.ns).Resource("imagestreamtags").Name(name).VersionedParams(&options, scheme.ParameterCodec).Do().Into(result)
	return
}
func (c *imageStreamTags) Create(imageStreamTag *image.ImageStreamTag) (result *image.ImageStreamTag, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	result = &image.ImageStreamTag{}
	err = c.client.Post().Namespace(c.ns).Resource("imagestreamtags").Body(imageStreamTag).Do().Into(result)
	return
}
func (c *imageStreamTags) Update(imageStreamTag *image.ImageStreamTag) (result *image.ImageStreamTag, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	result = &image.ImageStreamTag{}
	err = c.client.Put().Namespace(c.ns).Resource("imagestreamtags").Name(imageStreamTag.Name).Body(imageStreamTag).Do().Into(result)
	return
}
func (c *imageStreamTags) Delete(name string, options *v1.DeleteOptions) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.client.Delete().Namespace(c.ns).Resource("imagestreamtags").Name(name).Body(options).Do().Error()
}
