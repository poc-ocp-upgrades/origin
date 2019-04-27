package fake

import (
	template "github.com/openshift/origin/pkg/template/apis/template"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

type FakeBrokerTemplateInstances struct{ Fake *FakeTemplate }

var brokertemplateinstancesResource = schema.GroupVersionResource{Group: "template.openshift.io", Version: "", Resource: "brokertemplateinstances"}
var brokertemplateinstancesKind = schema.GroupVersionKind{Group: "template.openshift.io", Version: "", Kind: "BrokerTemplateInstance"}

func (c *FakeBrokerTemplateInstances) Get(name string, options v1.GetOptions) (result *template.BrokerTemplateInstance, err error) {
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
	obj, err := c.Fake.Invokes(testing.NewRootGetAction(brokertemplateinstancesResource, name), &template.BrokerTemplateInstance{})
	if obj == nil {
		return nil, err
	}
	return obj.(*template.BrokerTemplateInstance), err
}
func (c *FakeBrokerTemplateInstances) List(opts v1.ListOptions) (result *template.BrokerTemplateInstanceList, err error) {
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
	obj, err := c.Fake.Invokes(testing.NewRootListAction(brokertemplateinstancesResource, brokertemplateinstancesKind, opts), &template.BrokerTemplateInstanceList{})
	if obj == nil {
		return nil, err
	}
	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &template.BrokerTemplateInstanceList{ListMeta: obj.(*template.BrokerTemplateInstanceList).ListMeta}
	for _, item := range obj.(*template.BrokerTemplateInstanceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}
func (c *FakeBrokerTemplateInstances) Watch(opts v1.ListOptions) (watch.Interface, error) {
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
	return c.Fake.InvokesWatch(testing.NewRootWatchAction(brokertemplateinstancesResource, opts))
}
func (c *FakeBrokerTemplateInstances) Create(brokerTemplateInstance *template.BrokerTemplateInstance) (result *template.BrokerTemplateInstance, err error) {
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
	obj, err := c.Fake.Invokes(testing.NewRootCreateAction(brokertemplateinstancesResource, brokerTemplateInstance), &template.BrokerTemplateInstance{})
	if obj == nil {
		return nil, err
	}
	return obj.(*template.BrokerTemplateInstance), err
}
func (c *FakeBrokerTemplateInstances) Update(brokerTemplateInstance *template.BrokerTemplateInstance) (result *template.BrokerTemplateInstance, err error) {
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
	obj, err := c.Fake.Invokes(testing.NewRootUpdateAction(brokertemplateinstancesResource, brokerTemplateInstance), &template.BrokerTemplateInstance{})
	if obj == nil {
		return nil, err
	}
	return obj.(*template.BrokerTemplateInstance), err
}
func (c *FakeBrokerTemplateInstances) Delete(name string, options *v1.DeleteOptions) error {
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
	_, err := c.Fake.Invokes(testing.NewRootDeleteAction(brokertemplateinstancesResource, name), &template.BrokerTemplateInstance{})
	return err
}
func (c *FakeBrokerTemplateInstances) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
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
	action := testing.NewRootDeleteCollectionAction(brokertemplateinstancesResource, listOptions)
	_, err := c.Fake.Invokes(action, &template.BrokerTemplateInstanceList{})
	return err
}
func (c *FakeBrokerTemplateInstances) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *template.BrokerTemplateInstance, err error) {
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
	obj, err := c.Fake.Invokes(testing.NewRootPatchSubresourceAction(brokertemplateinstancesResource, name, pt, data, subresources...), &template.BrokerTemplateInstance{})
	if obj == nil {
		return nil, err
	}
	return obj.(*template.BrokerTemplateInstance), err
}
