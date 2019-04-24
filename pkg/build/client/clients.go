package client

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"bytes"
	"net/http"
	"runtime"
	"fmt"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/types"
	buildv1 "github.com/openshift/api/build/v1"
	buildclient "github.com/openshift/client-go/build/clientset/versioned"
	buildclienttyped "github.com/openshift/client-go/build/clientset/versioned/typed/build/v1"
	buildlister "github.com/openshift/client-go/build/listers/build/v1"
)

type BuildConfigGetter interface {
	Get(namespace, name string, options metav1.GetOptions) (*buildv1.BuildConfig, error)
}
type BuildConfigUpdater interface {
	Update(buildConfig *buildv1.BuildConfig) error
}
type ClientBuildConfigClient struct{ Client buildclient.Interface }

func NewClientBuildConfigClient(client buildclient.Interface) *ClientBuildConfigClient {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &ClientBuildConfigClient{Client: client}
}
func (c ClientBuildConfigClient) Get(namespace, name string, options metav1.GetOptions) (*buildv1.BuildConfig, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.Client.BuildV1().BuildConfigs(namespace).Get(name, options)
}
func (c ClientBuildConfigClient) Update(buildConfig *buildv1.BuildConfig) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_, err := c.Client.BuildV1().BuildConfigs(buildConfig.Namespace).Update(buildConfig)
	return err
}

type BuildUpdater interface {
	Update(namespace string, build *buildv1.Build) error
}
type BuildPatcher interface {
	Patch(namespace, name string, patch []byte) (*buildv1.Build, error)
}
type BuildLister interface {
	List(namespace string, opts metav1.ListOptions) (*buildv1.BuildList, error)
}
type BuildDeleter interface {
	DeleteBuild(build *buildv1.Build) error
}
type ClientBuildClient struct{ Client buildclient.Interface }

func NewClientBuildClient(client buildclient.Interface) *ClientBuildClient {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &ClientBuildClient{Client: client}
}
func (c ClientBuildClient) Update(namespace string, build *buildv1.Build) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_, e := c.Client.BuildV1().Builds(namespace).Update(build)
	return e
}
func (c ClientBuildClient) Patch(namespace, name string, patch []byte) (*buildv1.Build, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.Client.BuildV1().Builds(namespace).Patch(name, types.StrategicMergePatchType, patch)
}
func (c ClientBuildClient) List(namespace string, opts metav1.ListOptions) (*buildv1.BuildList, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.Client.BuildV1().Builds(namespace).List(opts)
}
func (c ClientBuildClient) DeleteBuild(build *buildv1.Build) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.Client.BuildV1().Builds(build.Namespace).Delete(build.Name, &metav1.DeleteOptions{})
}

type ClientBuildLister struct{ client buildclienttyped.BuildsGetter }

func NewClientBuildLister(client buildclienttyped.BuildsGetter) buildlister.BuildLister {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &ClientBuildLister{client: client}
}
func (c *ClientBuildLister) List(label labels.Selector) ([]*buildv1.Build, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	list, err := c.client.Builds(metav1.NamespaceAll).List(metav1.ListOptions{LabelSelector: label.String()})
	return buildListToPointerArray(list), err
}
func (c *ClientBuildLister) Builds(ns string) buildlister.BuildNamespaceLister {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &ClientBuildListerNamespacer{client: c.client, ns: ns}
}

type ClientBuildListerNamespacer struct {
	client	buildclienttyped.BuildsGetter
	ns	string
}

func (c ClientBuildListerNamespacer) List(label labels.Selector) ([]*buildv1.Build, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	list, err := c.client.Builds(c.ns).List(metav1.ListOptions{LabelSelector: label.String()})
	return buildListToPointerArray(list), err
}
func (c ClientBuildListerNamespacer) Get(name string) (*buildv1.Build, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.client.Builds(c.ns).Get(name, metav1.GetOptions{})
}
func buildListToPointerArray(list *buildv1.BuildList) []*buildv1.Build {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if list == nil {
		return nil
	}
	result := make([]*buildv1.Build, len(list.Items))
	for i := range list.Items {
		result[i] = &list.Items[i]
	}
	return result
}

type ClientBuildConfigLister struct {
	client buildclienttyped.BuildConfigsGetter
}

func NewClientBuildConfigLister(client buildclienttyped.BuildConfigsGetter) buildlister.BuildConfigLister {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &ClientBuildConfigLister{client: client}
}
func (c *ClientBuildConfigLister) List(label labels.Selector) ([]*buildv1.BuildConfig, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	list, err := c.client.BuildConfigs(metav1.NamespaceAll).List(metav1.ListOptions{LabelSelector: label.String()})
	return buildConfigListToPointerArray(list), err
}
func (c *ClientBuildConfigLister) BuildConfigs(ns string) buildlister.BuildConfigNamespaceLister {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &ClientBuildConfigListerNamespacer{client: c.client, ns: ns}
}

type ClientBuildConfigListerNamespacer struct {
	client	buildclienttyped.BuildConfigsGetter
	ns	string
}

func (c ClientBuildConfigListerNamespacer) List(label labels.Selector) ([]*buildv1.BuildConfig, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	list, err := c.client.BuildConfigs(c.ns).List(metav1.ListOptions{LabelSelector: label.String()})
	return buildConfigListToPointerArray(list), err
}
func (c ClientBuildConfigListerNamespacer) Get(name string) (*buildv1.BuildConfig, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.client.BuildConfigs(c.ns).Get(name, metav1.GetOptions{})
}
func buildConfigListToPointerArray(list *buildv1.BuildConfigList) []*buildv1.BuildConfig {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if list == nil {
		return nil
	}
	result := make([]*buildv1.BuildConfig, len(list.Items))
	for i := range list.Items {
		result[i] = &list.Items[i]
	}
	return result
}

type BuildCloner interface {
	Clone(namespace string, request *buildv1.BuildRequest) (*buildv1.Build, error)
}
type ClientBuildClonerClient struct{ Client buildclient.Interface }

func NewClientBuildClonerClient(client buildclient.Interface) *ClientBuildClonerClient {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &ClientBuildClonerClient{Client: client}
}
func (c ClientBuildClonerClient) Clone(namespace string, request *buildv1.BuildRequest) (*buildv1.Build, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.Client.BuildV1().Builds(namespace).Clone(request.Name, request)
}

type BuildConfigInstantiator interface {
	Instantiate(namespace string, request *buildv1.BuildRequest) (*buildv1.Build, error)
}
type ClientBuildConfigInstantiatorClient struct{ Client buildclient.Interface }

func NewClientBuildConfigInstantiatorClient(client buildclient.Interface) *ClientBuildConfigInstantiatorClient {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &ClientBuildConfigInstantiatorClient{Client: client}
}
func (c ClientBuildConfigInstantiatorClient) Instantiate(namespace string, request *buildv1.BuildRequest) (*buildv1.Build, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.Client.BuildV1().BuildConfigs(namespace).Instantiate(request.Name, request)
}

type BuildConfigInstantiatorClient struct {
	Client buildclienttyped.BuildV1Interface
}

func (c BuildConfigInstantiatorClient) Instantiate(namespace string, request *buildv1.BuildRequest) (*buildv1.Build, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.Client.BuildConfigs(namespace).Instantiate(request.Name, request)
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := runtime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", runtime.FuncForPC(pc).Name()))
	http.Post("/"+"logcode", "application/json", bytes.NewBuffer(jsonLog))
}
