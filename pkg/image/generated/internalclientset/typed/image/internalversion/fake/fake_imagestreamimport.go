package fake

import (
	image "github.com/openshift/origin/pkg/image/apis/image"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	testing "k8s.io/client-go/testing"
)

type FakeImageStreamImports struct {
	Fake *FakeImage
	ns   string
}

var imagestreamimportsResource = schema.GroupVersionResource{Group: "image.openshift.io", Version: "", Resource: "imagestreamimports"}
var imagestreamimportsKind = schema.GroupVersionKind{Group: "image.openshift.io", Version: "", Kind: "ImageStreamImport"}

func (c *FakeImageStreamImports) Create(imageStreamImport *image.ImageStreamImport) (result *image.ImageStreamImport, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := c.Fake.Invokes(testing.NewCreateAction(imagestreamimportsResource, c.ns, imageStreamImport), &image.ImageStreamImport{})
	if obj == nil {
		return nil, err
	}
	return obj.(*image.ImageStreamImport), err
}
