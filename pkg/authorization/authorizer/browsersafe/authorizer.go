package browsersafe

import (
	"fmt"
	goformat "fmt"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/apiserver/pkg/authorization/authorizer"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

const (
	proxyAction = "proxy"
	unsafeProxy = "unsafeproxy"
)

type browserSafeAuthorizer struct {
	delegate            authorizer.Authorizer
	authenticatedGroups sets.String
}

func NewBrowserSafeAuthorizer(delegate authorizer.Authorizer, authenticatedGroups ...string) authorizer.Authorizer {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &browserSafeAuthorizer{delegate: delegate, authenticatedGroups: sets.NewString(authenticatedGroups...)}
}
func (a *browserSafeAuthorizer) Authorize(attributes authorizer.Attributes) (authorizer.Decision, string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	attrs := a.getBrowserSafeAttributes(attributes)
	decision, reason, err := a.delegate.Authorize(attrs)
	safeAttributes, changed := attrs.(*browserSafeAttributes)
	if decision == authorizer.DecisionAllow || !changed {
		return decision, reason, err
	}
	return decision, safeAttributes.reason(reason), err
}
func (a *browserSafeAuthorizer) getBrowserSafeAttributes(attributes authorizer.Attributes) authorizer.Attributes {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
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
	isProxyVerb, isProxySubresource bool
}

func (b *browserSafeAttributes) GetVerb() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if b.isProxyVerb {
		return unsafeProxy
	}
	return b.Attributes.GetVerb()
}
func (b *browserSafeAttributes) GetSubresource() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if b.isProxySubresource {
		return unsafeProxy
	}
	return b.Attributes.GetSubresource()
}
func (b *browserSafeAttributes) reason(reason string) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
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
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
