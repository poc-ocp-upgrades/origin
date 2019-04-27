package controller

import (
	"k8s.io/apimachinery/pkg/util/sets"
	kauthorizer "k8s.io/apiserver/pkg/authorization/authorizer"
)

type bypassAuthorizer struct {
	paths		sets.String
	authorizer	kauthorizer.Authorizer
}

func newBypassAuthorizer(auth kauthorizer.Authorizer, paths ...string) kauthorizer.Authorizer {
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
	return bypassAuthorizer{paths: sets.NewString(paths...), authorizer: auth}
}
func (a bypassAuthorizer) Authorize(attributes kauthorizer.Attributes) (allowed kauthorizer.Decision, reason string, err error) {
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
	if !attributes.IsResourceRequest() && a.paths.Has(attributes.GetPath()) {
		return kauthorizer.DecisionAllow, "always allowed", nil
	}
	return a.authorizer.Authorize(attributes)
}
