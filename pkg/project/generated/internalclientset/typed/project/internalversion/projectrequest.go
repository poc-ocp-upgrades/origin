package internalversion

import (
	project "github.com/openshift/origin/pkg/project/apis/project"
	rest "k8s.io/client-go/rest"
)

type ProjectRequestsGetter interface {
	ProjectRequests() ProjectRequestInterface
}
type ProjectRequestInterface interface {
	Create(*project.ProjectRequest) (*project.Project, error)
	ProjectRequestExpansion
}
type projectRequests struct{ client rest.Interface }

func newProjectRequests(c *ProjectClient) *projectRequests {
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
	return &projectRequests{client: c.RESTClient()}
}
func (c *projectRequests) Create(projectRequest *project.ProjectRequest) (result *project.Project, err error) {
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
	result = &project.Project{}
	err = c.client.Post().Resource("projectrequests").Body(projectRequest).Do().Into(result)
	return
}
