package authorization

import (
	"fmt"
	"strings"
	"k8s.io/apimachinery/pkg/api/validation/path"
	"k8s.io/apimachinery/pkg/util/sets"
	kapi "k8s.io/kubernetes/pkg/apis/core"
	"github.com/openshift/origin/pkg/authorization/apis/authorization/internal/serviceaccount"
)

func (r PolicyRule) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return "PolicyRule" + r.CompactString()
}
func (r PolicyRule) CompactString() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	formatStringParts := []string{}
	formatArgs := []interface{}{}
	if len(r.Verbs) > 0 {
		formatStringParts = append(formatStringParts, "Verbs:%q")
		formatArgs = append(formatArgs, r.Verbs.List())
	}
	if len(r.APIGroups) > 0 {
		formatStringParts = append(formatStringParts, "APIGroups:%q")
		formatArgs = append(formatArgs, r.APIGroups)
	}
	if len(r.Resources) > 0 {
		formatStringParts = append(formatStringParts, "Resources:%q")
		formatArgs = append(formatArgs, r.Resources.List())
	}
	if len(r.ResourceNames) > 0 {
		formatStringParts = append(formatStringParts, "ResourceNames:%q")
		formatArgs = append(formatArgs, r.ResourceNames.List())
	}
	if r.AttributeRestrictions != nil {
		formatStringParts = append(formatStringParts, "Restrictions:%q")
		formatArgs = append(formatArgs, r.AttributeRestrictions)
	}
	if len(r.NonResourceURLs) > 0 {
		formatStringParts = append(formatStringParts, "NonResourceURLs:%q")
		formatArgs = append(formatArgs, r.NonResourceURLs.List())
	}
	formatString := "{" + strings.Join(formatStringParts, ", ") + "}"
	return fmt.Sprintf(formatString, formatArgs...)
}
func BuildSubjects(users, groups []string) []kapi.ObjectReference {
	_logClusterCodePath()
	defer _logClusterCodePath()
	subjects := []kapi.ObjectReference{}
	for _, user := range users {
		saNamespace, saName, err := serviceaccount.SplitUsername(user)
		if err == nil {
			subjects = append(subjects, kapi.ObjectReference{Kind: ServiceAccountKind, Namespace: saNamespace, Name: saName})
			continue
		}
		kind := determineUserKind(user)
		subjects = append(subjects, kapi.ObjectReference{Kind: kind, Name: user})
	}
	for _, group := range groups {
		kind := determineGroupKind(group)
		subjects = append(subjects, kapi.ObjectReference{Kind: kind, Name: group})
	}
	return subjects
}
func validateUserName(name string, _ bool) []string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if reasons := path.ValidatePathSegmentName(name, false); len(reasons) != 0 {
		return reasons
	}
	if strings.Contains(name, ":") {
		return []string{`may not contain ":"`}
	}
	if name == "~" {
		return []string{`may not equal "~"`}
	}
	return nil
}
func validateGroupName(name string, _ bool) []string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if reasons := path.ValidatePathSegmentName(name, false); len(reasons) != 0 {
		return reasons
	}
	if strings.Contains(name, ":") {
		return []string{`may not contain ":"`}
	}
	if name == "~" {
		return []string{`may not equal "~"`}
	}
	return nil
}
func determineUserKind(user string) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	kind := UserKind
	if len(validateUserName(user, false)) != 0 {
		kind = SystemUserKind
	}
	return kind
}
func determineGroupKind(group string) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	kind := GroupKind
	if len(validateGroupName(group, false)) != 0 {
		kind = SystemGroupKind
	}
	return kind
}
func StringSubjectsFor(currentNamespace string, subjects []kapi.ObjectReference) ([]string, []string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var users, groups []string
	for _, subject := range subjects {
		switch subject.Kind {
		case ServiceAccountKind:
			namespace := currentNamespace
			if len(subject.Namespace) > 0 {
				namespace = subject.Namespace
			}
			if len(namespace) > 0 {
				users = append(users, serviceaccount.MakeUsername(namespace, subject.Name))
			}
		case UserKind, SystemUserKind:
			users = append(users, subject.Name)
		case GroupKind, SystemGroupKind:
			groups = append(groups, subject.Name)
		}
	}
	return users, groups
}
func SubjectsStrings(currentNamespace string, subjects []kapi.ObjectReference) ([]string, []string, []string, []string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	users := []string{}
	groups := []string{}
	sas := []string{}
	others := []string{}
	for _, subject := range subjects {
		switch subject.Kind {
		case ServiceAccountKind:
			if len(subject.Namespace) > 0 && currentNamespace != subject.Namespace {
				sas = append(sas, subject.Namespace+"/"+subject.Name)
			} else {
				sas = append(sas, subject.Name)
			}
		case UserKind, SystemUserKind:
			users = append(users, subject.Name)
		case GroupKind, SystemGroupKind:
			groups = append(groups, subject.Name)
		default:
			others = append(others, fmt.Sprintf("%s/%s/%s", subject.Kind, subject.Namespace, subject.Name))
		}
	}
	return users, groups, sas, others
}

type PolicyRuleBuilder struct{ PolicyRule PolicyRule }

func NewRule(verbs ...string) *PolicyRuleBuilder {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &PolicyRuleBuilder{PolicyRule: PolicyRule{Verbs: sets.NewString(verbs...), Resources: sets.String{}, ResourceNames: sets.String{}}}
}
func (r *PolicyRuleBuilder) Groups(groups ...string) *PolicyRuleBuilder {
	_logClusterCodePath()
	defer _logClusterCodePath()
	r.PolicyRule.APIGroups = append(r.PolicyRule.APIGroups, groups...)
	return r
}
func (r *PolicyRuleBuilder) Resources(resources ...string) *PolicyRuleBuilder {
	_logClusterCodePath()
	defer _logClusterCodePath()
	r.PolicyRule.Resources.Insert(resources...)
	return r
}
func (r *PolicyRuleBuilder) Names(names ...string) *PolicyRuleBuilder {
	_logClusterCodePath()
	defer _logClusterCodePath()
	r.PolicyRule.ResourceNames.Insert(names...)
	return r
}
func (r *PolicyRuleBuilder) RuleOrDie() PolicyRule {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ret, err := r.Rule()
	if err != nil {
		panic(err)
	}
	return ret
}
func (r *PolicyRuleBuilder) Rule() (PolicyRule, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if r.PolicyRule.AttributeRestrictions != nil {
		return PolicyRule{}, fmt.Errorf("rule may not have attributeRestrictions because they are deprecated and ignored: %#v", r.PolicyRule)
	}
	if len(r.PolicyRule.Verbs) == 0 {
		return PolicyRule{}, fmt.Errorf("verbs are required: %#v", r.PolicyRule)
	}
	switch {
	case len(r.PolicyRule.NonResourceURLs) > 0:
		if len(r.PolicyRule.APIGroups) != 0 || len(r.PolicyRule.Resources) != 0 || len(r.PolicyRule.ResourceNames) != 0 {
			return PolicyRule{}, fmt.Errorf("non-resource rule may not have apiGroups, resources, or resourceNames: %#v", r.PolicyRule)
		}
	case len(r.PolicyRule.Resources) > 0:
		if len(r.PolicyRule.NonResourceURLs) != 0 {
			return PolicyRule{}, fmt.Errorf("resource rule may not have nonResourceURLs: %#v", r.PolicyRule)
		}
		if len(r.PolicyRule.APIGroups) == 0 {
			return PolicyRule{}, fmt.Errorf("resource rule must have apiGroups: %#v", r.PolicyRule)
		}
	default:
		return PolicyRule{}, fmt.Errorf("a rule must have either nonResourceURLs or resources: %#v", r.PolicyRule)
	}
	return r.PolicyRule, nil
}

type SortableRuleSlice []PolicyRule

func (s SortableRuleSlice) Len() int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return len(s)
}
func (s SortableRuleSlice) Swap(i, j int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s[i], s[j] = s[j], s[i]
}
func (s SortableRuleSlice) Less(i, j int) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return strings.Compare(s[i].String(), s[j].String()) < 0
}
