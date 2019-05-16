package authorizer

import (
	"github.com/openshift/origin/pkg/project/apis/project"
	apirequest "k8s.io/apiserver/pkg/endpoints/request"
	"net/http"
)

type projectRequestInfoResolver struct {
	infoFactory apirequest.RequestInfoResolver
}

func NewProjectRequestInfoResolver(infoFactory apirequest.RequestInfoResolver) apirequest.RequestInfoResolver {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &projectRequestInfoResolver{infoFactory: infoFactory}
}
func (a *projectRequestInfoResolver) NewRequestInfo(req *http.Request) (*apirequest.RequestInfo, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	requestInfo, err := a.infoFactory.NewRequestInfo(req)
	if err != nil {
		return requestInfo, err
	}
	if (len(requestInfo.APIGroup) == 0 || requestInfo.APIGroup == project.GroupName) && requestInfo.Resource == "projects" && len(requestInfo.Name) > 0 {
		requestInfo.Namespace = requestInfo.Name
	}
	return requestInfo, nil
}
