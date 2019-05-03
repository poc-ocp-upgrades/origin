package internalversion

import (
	image "github.com/openshift/origin/pkg/image/apis/image"
	scheme "github.com/openshift/origin/pkg/image/generated/internalclientset/scheme"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	"time"
)

type ImageStreamsGetter interface {
	ImageStreams(namespace string) ImageStreamInterface
}
type ImageStreamInterface interface {
	Create(*image.ImageStream) (*image.ImageStream, error)
	Update(*image.ImageStream) (*image.ImageStream, error)
	UpdateStatus(*image.ImageStream) (*image.ImageStream, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*image.ImageStream, error)
	List(opts v1.ListOptions) (*image.ImageStreamList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *image.ImageStream, err error)
	Secrets(imageStreamName string, options v1.GetOptions) (*corev1.SecretList, error)
	ImageStreamExpansion
}
type imageStreams struct {
	client rest.Interface
	ns     string
}

func newImageStreams(c *ImageClient, namespace string) *imageStreams {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &imageStreams{client: c.RESTClient(), ns: namespace}
}
func (c *imageStreams) Get(name string, options v1.GetOptions) (result *image.ImageStream, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	result = &image.ImageStream{}
	err = c.client.Get().Namespace(c.ns).Resource("imagestreams").Name(name).VersionedParams(&options, scheme.ParameterCodec).Do().Into(result)
	return
}
func (c *imageStreams) List(opts v1.ListOptions) (result *image.ImageStreamList, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &image.ImageStreamList{}
	err = c.client.Get().Namespace(c.ns).Resource("imagestreams").VersionedParams(&opts, scheme.ParameterCodec).Timeout(timeout).Do().Into(result)
	return
}
func (c *imageStreams) Watch(opts v1.ListOptions) (watch.Interface, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().Namespace(c.ns).Resource("imagestreams").VersionedParams(&opts, scheme.ParameterCodec).Timeout(timeout).Watch()
}
func (c *imageStreams) Create(imageStream *image.ImageStream) (result *image.ImageStream, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	result = &image.ImageStream{}
	err = c.client.Post().Namespace(c.ns).Resource("imagestreams").Body(imageStream).Do().Into(result)
	return
}
func (c *imageStreams) Update(imageStream *image.ImageStream) (result *image.ImageStream, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	result = &image.ImageStream{}
	err = c.client.Put().Namespace(c.ns).Resource("imagestreams").Name(imageStream.Name).Body(imageStream).Do().Into(result)
	return
}
func (c *imageStreams) UpdateStatus(imageStream *image.ImageStream) (result *image.ImageStream, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	result = &image.ImageStream{}
	err = c.client.Put().Namespace(c.ns).Resource("imagestreams").Name(imageStream.Name).SubResource("status").Body(imageStream).Do().Into(result)
	return
}
func (c *imageStreams) Delete(name string, options *v1.DeleteOptions) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.client.Delete().Namespace(c.ns).Resource("imagestreams").Name(name).Body(options).Do().Error()
}
func (c *imageStreams) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().Namespace(c.ns).Resource("imagestreams").VersionedParams(&listOptions, scheme.ParameterCodec).Timeout(timeout).Body(options).Do().Error()
}
func (c *imageStreams) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *image.ImageStream, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	result = &image.ImageStream{}
	err = c.client.Patch(pt).Namespace(c.ns).Resource("imagestreams").SubResource(subresources...).Name(name).Body(data).Do().Into(result)
	return
}
func (c *imageStreams) Secrets(imageStreamName string, options v1.GetOptions) (result *corev1.SecretList, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	result = &corev1.SecretList{}
	err = c.client.Get().Namespace(c.ns).Resource("imagestreams").Name(imageStreamName).SubResource("secrets").VersionedParams(&options, scheme.ParameterCodec).Do().Into(result)
	return
}
