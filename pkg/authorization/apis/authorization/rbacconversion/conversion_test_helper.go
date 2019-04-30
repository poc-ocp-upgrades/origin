package rbacconversion

import (
	"strings"
	"k8s.io/apimachinery/pkg/util/sets"
	authorizationapi "github.com/openshift/origin/pkg/authorization/apis/authorization"
)

func Covers(ownerRules, servantRules []authorizationapi.PolicyRule) (bool, []authorizationapi.PolicyRule) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	subrules := []authorizationapi.PolicyRule{}
	for _, servantRule := range servantRules {
		subrules = append(subrules, BreakdownRule(servantRule)...)
	}
	uncoveredRules := []authorizationapi.PolicyRule{}
	for _, subrule := range subrules {
		covered := false
		for _, ownerRule := range ownerRules {
			if ruleCovers(ownerRule, subrule) {
				covered = true
				break
			}
		}
		if !covered {
			uncoveredRules = append(uncoveredRules, subrule)
		}
	}
	return (len(uncoveredRules) == 0), uncoveredRules
}
func BreakdownRule(rule authorizationapi.PolicyRule) []authorizationapi.PolicyRule {
	_logClusterCodePath()
	defer _logClusterCodePath()
	subrules := []authorizationapi.PolicyRule{}
	if rule.AttributeRestrictions != nil {
		return subrules
	}
	for _, group := range rule.APIGroups {
		subrules = append(subrules, breakdownRuleForGroup(group, rule)...)
	}
	if len(rule.APIGroups) == 0 {
		for _, subrule := range breakdownRuleForGroup("", rule) {
			subrule.APIGroups = nil
			subrules = append(subrules, subrule)
		}
	}
	for nonResourceURL := range rule.NonResourceURLs {
		for verb := range rule.Verbs {
			subrules = append(subrules, authorizationapi.PolicyRule{Verbs: sets.NewString(verb), NonResourceURLs: sets.NewString(nonResourceURL)})
		}
	}
	return subrules
}
func breakdownRuleForGroup(group string, rule authorizationapi.PolicyRule) []authorizationapi.PolicyRule {
	_logClusterCodePath()
	defer _logClusterCodePath()
	subrules := []authorizationapi.PolicyRule{}
	for resource := range rule.Resources {
		for verb := range rule.Verbs {
			if len(rule.ResourceNames) > 0 {
				for _, resourceName := range rule.ResourceNames.List() {
					subrules = append(subrules, authorizationapi.PolicyRule{APIGroups: []string{group}, Resources: sets.NewString(resource), Verbs: sets.NewString(verb), ResourceNames: sets.NewString(resourceName)})
				}
			} else {
				subrules = append(subrules, authorizationapi.PolicyRule{APIGroups: []string{group}, Resources: sets.NewString(resource), Verbs: sets.NewString(verb)})
			}
		}
	}
	return subrules
}
func ruleCovers(ownerRule, subrule authorizationapi.PolicyRule) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if ownerRule.AttributeRestrictions != nil {
		return false
	}
	allResources := ownerRule.Resources
	ownerGroups := sets.NewString(ownerRule.APIGroups...)
	groupMatches := ownerGroups.Has(authorizationapi.APIGroupAll) || ownerGroups.HasAll(subrule.APIGroups...) || (len(ownerRule.APIGroups) == 0 && len(subrule.APIGroups) == 0)
	verbMatches := ownerRule.Verbs.Has(authorizationapi.VerbAll) || ownerRule.Verbs.HasAll(subrule.Verbs.List()...)
	resourceMatches := ownerRule.Resources.Has(authorizationapi.ResourceAll) || allResources.HasAll(subrule.Resources.List()...)
	resourceNameMatches := false
	if len(subrule.ResourceNames) == 0 {
		resourceNameMatches = (len(ownerRule.ResourceNames) == 0)
	} else {
		resourceNameMatches = (len(ownerRule.ResourceNames) == 0) || ownerRule.ResourceNames.HasAll(subrule.ResourceNames.List()...)
	}
	nonResourceCovers := nonResourceRuleCovers(ownerRule.NonResourceURLs, subrule.NonResourceURLs)
	return verbMatches && resourceMatches && resourceNameMatches && groupMatches && nonResourceCovers
}
func nonResourceRuleCovers(allowedPaths sets.String, requestedPaths sets.String) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if allowedPaths.Has(authorizationapi.NonResourceAll) {
		return true
	}
	for requestedPath := range requestedPaths {
		if allowedPaths.Has(requestedPath) {
			continue
		}
		prefixMatch := false
		for allowedPath := range allowedPaths {
			if strings.HasSuffix(allowedPath, "*") {
				if strings.HasPrefix(requestedPath, allowedPath[0:len(allowedPath)-1]) {
					return true
				}
			}
		}
		if !prefixMatch {
			return false
		}
	}
	return true
}
