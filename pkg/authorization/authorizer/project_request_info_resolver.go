package authorizer

import (
	"net/http"
	apirequest "k8s.io/apiserver/pkg/endpoints/request"
	"github.com/openshift/origin/pkg/project/apis/project"
)

type projectRequestInfoResolver struct {
	infoFactory apirequest.RequestInfoResolver
}

func NewProjectRequestInfoResolver(infoFactory apirequest.RequestInfoResolver) apirequest.RequestInfoResolver {
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
	return &projectRequestInfoResolver{infoFactory: infoFactory}
}
func (a *projectRequestInfoResolver) NewRequestInfo(req *http.Request) (*apirequest.RequestInfo, error) {
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
	requestInfo, err := a.infoFactory.NewRequestInfo(req)
	if err != nil {
		return requestInfo, err
	}
	if (len(requestInfo.APIGroup) == 0 || requestInfo.APIGroup == project.GroupName) && requestInfo.Resource == "projects" && len(requestInfo.Name) > 0 {
		requestInfo.Namespace = requestInfo.Name
	}
	return requestInfo, nil
}
