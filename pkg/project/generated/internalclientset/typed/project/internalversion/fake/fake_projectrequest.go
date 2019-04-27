package fake

import (
	project "github.com/openshift/origin/pkg/project/apis/project"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	testing "k8s.io/client-go/testing"
)

type FakeProjectRequests struct{ Fake *FakeProject }

var projectrequestsResource = schema.GroupVersionResource{Group: "project.openshift.io", Version: "", Resource: "projectrequests"}
var projectrequestsKind = schema.GroupVersionKind{Group: "project.openshift.io", Version: "", Kind: "ProjectRequest"}

func (c *FakeProjectRequests) Create(projectRequest *project.ProjectRequest) (result *project.Project, err error) {
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
	obj, err := c.Fake.Invokes(testing.NewRootCreateAction(projectrequestsResource, projectRequest), &project.Project{})
	if obj == nil {
		return nil, err
	}
	return obj.(*project.Project), err
}
