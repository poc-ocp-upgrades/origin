package fake

import (
	authorization "github.com/openshift/origin/pkg/authorization/apis/authorization"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

type FakeClusterRoleBindings struct{ Fake *FakeAuthorization }

var clusterrolebindingsResource = schema.GroupVersionResource{Group: "authorization.openshift.io", Version: "", Resource: "clusterrolebindings"}
var clusterrolebindingsKind = schema.GroupVersionKind{Group: "authorization.openshift.io", Version: "", Kind: "ClusterRoleBinding"}

func (c *FakeClusterRoleBindings) Get(name string, options v1.GetOptions) (result *authorization.ClusterRoleBinding, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := c.Fake.Invokes(testing.NewRootGetAction(clusterrolebindingsResource, name), &authorization.ClusterRoleBinding{})
	if obj == nil {
		return nil, err
	}
	return obj.(*authorization.ClusterRoleBinding), err
}
func (c *FakeClusterRoleBindings) List(opts v1.ListOptions) (result *authorization.ClusterRoleBindingList, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := c.Fake.Invokes(testing.NewRootListAction(clusterrolebindingsResource, clusterrolebindingsKind, opts), &authorization.ClusterRoleBindingList{})
	if obj == nil {
		return nil, err
	}
	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &authorization.ClusterRoleBindingList{ListMeta: obj.(*authorization.ClusterRoleBindingList).ListMeta}
	for _, item := range obj.(*authorization.ClusterRoleBindingList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}
func (c *FakeClusterRoleBindings) Watch(opts v1.ListOptions) (watch.Interface, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.Fake.InvokesWatch(testing.NewRootWatchAction(clusterrolebindingsResource, opts))
}
func (c *FakeClusterRoleBindings) Create(clusterRoleBinding *authorization.ClusterRoleBinding) (result *authorization.ClusterRoleBinding, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := c.Fake.Invokes(testing.NewRootCreateAction(clusterrolebindingsResource, clusterRoleBinding), &authorization.ClusterRoleBinding{})
	if obj == nil {
		return nil, err
	}
	return obj.(*authorization.ClusterRoleBinding), err
}
func (c *FakeClusterRoleBindings) Update(clusterRoleBinding *authorization.ClusterRoleBinding) (result *authorization.ClusterRoleBinding, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := c.Fake.Invokes(testing.NewRootUpdateAction(clusterrolebindingsResource, clusterRoleBinding), &authorization.ClusterRoleBinding{})
	if obj == nil {
		return nil, err
	}
	return obj.(*authorization.ClusterRoleBinding), err
}
func (c *FakeClusterRoleBindings) Delete(name string, options *v1.DeleteOptions) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_, err := c.Fake.Invokes(testing.NewRootDeleteAction(clusterrolebindingsResource, name), &authorization.ClusterRoleBinding{})
	return err
}
func (c *FakeClusterRoleBindings) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	action := testing.NewRootDeleteCollectionAction(clusterrolebindingsResource, listOptions)
	_, err := c.Fake.Invokes(action, &authorization.ClusterRoleBindingList{})
	return err
}
func (c *FakeClusterRoleBindings) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *authorization.ClusterRoleBinding, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := c.Fake.Invokes(testing.NewRootPatchSubresourceAction(clusterrolebindingsResource, name, pt, data, subresources...), &authorization.ClusterRoleBinding{})
	if obj == nil {
		return nil, err
	}
	return obj.(*authorization.ClusterRoleBinding), err
}
