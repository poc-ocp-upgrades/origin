package internalversion

import (
	image "github.com/openshift/origin/pkg/image/apis/image"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	rest "k8s.io/client-go/rest"
)

type ImageSignaturesGetter interface {
	ImageSignatures() ImageSignatureInterface
}
type ImageSignatureInterface interface {
	Create(*image.ImageSignature) (*image.ImageSignature, error)
	Delete(name string, options *v1.DeleteOptions) error
	ImageSignatureExpansion
}
type imageSignatures struct{ client rest.Interface }

func newImageSignatures(c *ImageClient) *imageSignatures {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &imageSignatures{client: c.RESTClient()}
}
func (c *imageSignatures) Create(imageSignature *image.ImageSignature) (result *image.ImageSignature, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	result = &image.ImageSignature{}
	err = c.client.Post().Resource("imagesignatures").Body(imageSignature).Do().Into(result)
	return
}
func (c *imageSignatures) Delete(name string, options *v1.DeleteOptions) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.client.Delete().Resource("imagesignatures").Name(name).Body(options).Do().Error()
}
