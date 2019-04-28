package fake

import (
	image "github.com/openshift/origin/pkg/image/apis/image"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	testing "k8s.io/client-go/testing"
)

type FakeImageStreamTags struct {
	Fake	*FakeImage
	ns	string
}

var imagestreamtagsResource = schema.GroupVersionResource{Group: "image.openshift.io", Version: "", Resource: "imagestreamtags"}
var imagestreamtagsKind = schema.GroupVersionKind{Group: "image.openshift.io", Version: "", Kind: "ImageStreamTag"}

func (c *FakeImageStreamTags) Get(name string, options v1.GetOptions) (result *image.ImageStreamTag, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := c.Fake.Invokes(testing.NewGetAction(imagestreamtagsResource, c.ns, name), &image.ImageStreamTag{})
	if obj == nil {
		return nil, err
	}
	return obj.(*image.ImageStreamTag), err
}
func (c *FakeImageStreamTags) Create(imageStreamTag *image.ImageStreamTag) (result *image.ImageStreamTag, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := c.Fake.Invokes(testing.NewCreateAction(imagestreamtagsResource, c.ns, imageStreamTag), &image.ImageStreamTag{})
	if obj == nil {
		return nil, err
	}
	return obj.(*image.ImageStreamTag), err
}
func (c *FakeImageStreamTags) Update(imageStreamTag *image.ImageStreamTag) (result *image.ImageStreamTag, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := c.Fake.Invokes(testing.NewUpdateAction(imagestreamtagsResource, c.ns, imageStreamTag), &image.ImageStreamTag{})
	if obj == nil {
		return nil, err
	}
	return obj.(*image.ImageStreamTag), err
}
func (c *FakeImageStreamTags) Delete(name string, options *v1.DeleteOptions) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_, err := c.Fake.Invokes(testing.NewDeleteAction(imagestreamtagsResource, c.ns, name), &image.ImageStreamTag{})
	return err
}
