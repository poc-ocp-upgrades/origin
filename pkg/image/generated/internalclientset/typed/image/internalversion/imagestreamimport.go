package internalversion

import (
	image "github.com/openshift/origin/pkg/image/apis/image"
	rest "k8s.io/client-go/rest"
)

type ImageStreamImportsGetter interface {
	ImageStreamImports(namespace string) ImageStreamImportInterface
}
type ImageStreamImportInterface interface {
	Create(*image.ImageStreamImport) (*image.ImageStreamImport, error)
	ImageStreamImportExpansion
}
type imageStreamImports struct {
	client	rest.Interface
	ns	string
}

func newImageStreamImports(c *ImageClient, namespace string) *imageStreamImports {
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
	return &imageStreamImports{client: c.RESTClient(), ns: namespace}
}
func (c *imageStreamImports) Create(imageStreamImport *image.ImageStreamImport) (result *image.ImageStreamImport, err error) {
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
	result = &image.ImageStreamImport{}
	err = c.client.Post().Namespace(c.ns).Resource("imagestreamimports").Body(imageStreamImport).Do().Into(result)
	return
}
