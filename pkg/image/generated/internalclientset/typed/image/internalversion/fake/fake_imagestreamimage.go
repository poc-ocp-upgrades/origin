package fake

import (
	image "github.com/openshift/origin/pkg/image/apis/image"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	testing "k8s.io/client-go/testing"
)

type FakeImageStreamImages struct {
	Fake	*FakeImage
	ns	string
}

var imagestreamimagesResource = schema.GroupVersionResource{Group: "image.openshift.io", Version: "", Resource: "imagestreamimages"}
var imagestreamimagesKind = schema.GroupVersionKind{Group: "image.openshift.io", Version: "", Kind: "ImageStreamImage"}

func (c *FakeImageStreamImages) Get(name string, options v1.GetOptions) (result *image.ImageStreamImage, err error) {
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
	obj, err := c.Fake.Invokes(testing.NewGetAction(imagestreamimagesResource, c.ns, name), &image.ImageStreamImage{})
	if obj == nil {
		return nil, err
	}
	return obj.(*image.ImageStreamImage), err
}
