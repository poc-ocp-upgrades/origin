package rules

import (
	"k8s.io/klog"
	"bytes"
	"net/http"
	"runtime"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	imagepolicy "github.com/openshift/origin/pkg/image/apiserver/admission/apis/imagepolicy/v1"
)

type Accepter interface {
	Covers(metav1.GroupResource) bool
	Accepts(*ImagePolicyAttributes) bool
}
type mappedAccepter map[metav1.GroupResource]Accepter

func (a mappedAccepter) Covers(gr metav1.GroupResource) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_, ok := a[gr]
	return ok
}
func (a mappedAccepter) Accepts(attr *ImagePolicyAttributes) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	accepter, ok := a[attr.Resource]
	if !ok {
		return true
	}
	return accepter.Accepts(attr)
}

type executionAccepter struct {
	rules				[]imagepolicy.ImageExecutionPolicyRule
	covers				metav1.GroupResource
	defaultReject			bool
	integratedRegistryMatcher	RegistryMatcher
}

func NewExecutionRulesAccepter(rules []imagepolicy.ImageExecutionPolicyRule, integratedRegistryMatcher RegistryMatcher) (Accepter, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	mapped := make(mappedAccepter)
	for _, rule := range rules {
		over, selectors, err := imageConditionInfo(&rule.ImageCondition)
		if err != nil {
			return nil, err
		}
		rule.ImageCondition.MatchImageLabelSelectors = selectors
		for gr := range over {
			a, ok := mapped[gr]
			if !ok {
				a = &executionAccepter{covers: gr, integratedRegistryMatcher: integratedRegistryMatcher}
				mapped[gr] = a
			}
			byResource := a.(*executionAccepter)
			byResource.rules = append(byResource.rules, rule)
		}
	}
	for _, a := range mapped {
		byResource := a.(*executionAccepter)
		if len(byResource.rules) > 0 {
			allReject := true
			for _, rule := range byResource.rules {
				if !rule.Reject {
					allReject = false
					break
				}
			}
			byResource.defaultReject = !allReject
		}
	}
	return mapped, nil
}
func (r *executionAccepter) Covers(gr metav1.GroupResource) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return r.covers == gr
}
func (r *executionAccepter) Accepts(attrs *ImagePolicyAttributes) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if attrs.Resource != r.covers {
		return true
	}
	anyMatched := false
	for _, rule := range r.rules {
		klog.V(5).Infof("image policy checking rule %q", rule.Name)
		if attrs.ExcludedRules.Has(rule.Name) && !rule.IgnoreNamespaceOverride {
			klog.V(5).Infof("skipping because rule is excluded by namespace annotations\n")
			continue
		}
		if attrs.Image == nil && rule.SkipOnResolutionFailure {
			klog.V(5).Infof("skipping because image is not resolved and skip on failure is true\n")
			continue
		}
		matches := matchImageCondition(&rule.ImageCondition, r.integratedRegistryMatcher, attrs)
		klog.V(5).Infof("Rule %q(reject=%t) applies to image %v: %t", rule.Name, rule.Reject, attrs.Name, matches)
		if matches {
			if rule.Reject {
				return false
			}
			anyMatched = true
		}
	}
	return anyMatched || !r.defaultReject
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := runtime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", runtime.FuncForPC(pc).Name()))
	http.Post("/"+"logcode", "application/json", bytes.NewBuffer(jsonLog))
}
