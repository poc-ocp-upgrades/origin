package fake

import (
	image "github.com/openshift/origin/pkg/image/apis/image"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	testing "k8s.io/client-go/testing"
)

type FakeImageStreamMappings struct {
	Fake	*FakeImage
	ns	string
}

var imagestreammappingsResource = schema.GroupVersionResource{Group: "image.openshift.io", Version: "", Resource: "imagestreammappings"}
var imagestreammappingsKind = schema.GroupVersionKind{Group: "image.openshift.io", Version: "", Kind: "ImageStreamMapping"}

func (c *FakeImageStreamMappings) Create(imageStreamMapping *image.ImageStreamMapping) (result *v1.Status, err error) {
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
	obj, err := c.Fake.Invokes(testing.NewCreateAction(imagestreammappingsResource, c.ns, imageStreamMapping), &v1.Status{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Status), err
}
