package fake

import (
	image "github.com/openshift/origin/pkg/image/apis/image"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	testing "k8s.io/client-go/testing"
)

type FakeImageSignatures struct{ Fake *FakeImage }

var imagesignaturesResource = schema.GroupVersionResource{Group: "image.openshift.io", Version: "", Resource: "imagesignatures"}
var imagesignaturesKind = schema.GroupVersionKind{Group: "image.openshift.io", Version: "", Kind: "ImageSignature"}

func (c *FakeImageSignatures) Create(imageSignature *image.ImageSignature) (result *image.ImageSignature, err error) {
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
	obj, err := c.Fake.Invokes(testing.NewRootCreateAction(imagesignaturesResource, imageSignature), &image.ImageSignature{})
	if obj == nil {
		return nil, err
	}
	return obj.(*image.ImageSignature), err
}
func (c *FakeImageSignatures) Delete(name string, options *v1.DeleteOptions) error {
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
	_, err := c.Fake.Invokes(testing.NewRootDeleteAction(imagesignaturesResource, name), &image.ImageSignature{})
	return err
}
