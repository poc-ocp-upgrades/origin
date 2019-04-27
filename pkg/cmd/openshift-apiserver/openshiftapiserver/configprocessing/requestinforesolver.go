package configprocessing

import (
	oauthorizer "github.com/openshift/origin/pkg/authorization/authorizer"
	"k8s.io/apimachinery/pkg/util/sets"
	apirequest "k8s.io/apiserver/pkg/endpoints/request"
)

func OpenshiftRequestInfoResolver() apirequest.RequestInfoResolver {
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
	requestInfoFactory := &apirequest.RequestInfoFactory{APIPrefixes: sets.NewString("api", "apis"), GrouplessAPIPrefixes: sets.NewString("api")}
	personalSARRequestInfoResolver := oauthorizer.NewPersonalSARRequestInfoResolver(requestInfoFactory)
	projectRequestInfoResolver := oauthorizer.NewProjectRequestInfoResolver(personalSARRequestInfoResolver)
	return projectRequestInfoResolver
}
