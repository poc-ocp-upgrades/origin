package scope

import (
	"fmt"
	"strings"
	rbacv1 "k8s.io/api/rbac/v1"
	kapierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime/schema"
	kutilerrors "k8s.io/apimachinery/pkg/util/errors"
	"k8s.io/apimachinery/pkg/util/sets"
	kauthorizer "k8s.io/apiserver/pkg/authorization/authorizer"
	rbaclisters "k8s.io/client-go/listers/rbac/v1"
	kauthorizationapi "k8s.io/kubernetes/pkg/apis/authorization"
	kapi "k8s.io/kubernetes/pkg/apis/core"
	"k8s.io/kubernetes/pkg/apis/rbac"
	rbacv1helpers "k8s.io/kubernetes/pkg/apis/rbac/v1"
	authorizerrbac "k8s.io/kubernetes/plugin/pkg/auth/authorizer/rbac"
	oauthapi "github.com/openshift/api/oauth/v1"
	"github.com/openshift/origin/pkg/api/legacy"
	authorizationapi "github.com/openshift/origin/pkg/authorization/apis/authorization"
	imageapi "github.com/openshift/origin/pkg/image/apis/image"
	projectapi "github.com/openshift/origin/pkg/project/apis/project"
	userapi "github.com/openshift/origin/pkg/user/apis/user"
)

func ScopesToRules(scopes []string, namespace string, clusterRoleGetter rbaclisters.ClusterRoleLister) ([]rbacv1.PolicyRule, error) {
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
	rules := append([]rbacv1.PolicyRule{}, authorizationapi.DiscoveryRule)
	errors := []error{}
	for _, scope := range scopes {
		found := false
		for _, evaluator := range ScopeEvaluators {
			if evaluator.Handles(scope) {
				found = true
				currRules, err := evaluator.ResolveRules(scope, namespace, clusterRoleGetter)
				if err != nil {
					errors = append(errors, err)
					continue
				}
				rules = append(rules, currRules...)
			}
		}
		if !found {
			errors = append(errors, fmt.Errorf("no scope evaluator found for %q", scope))
		}
	}
	return rules, kutilerrors.NewAggregate(errors)
}
func ScopesToVisibleNamespaces(scopes []string, clusterRoleGetter rbaclisters.ClusterRoleLister, ignoreUnhandledScopes bool) (sets.String, error) {
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
	if len(scopes) == 0 {
		return sets.NewString("*"), nil
	}
	visibleNamespaces := sets.String{}
	errors := []error{}
	for _, scope := range scopes {
		found := false
		for _, evaluator := range ScopeEvaluators {
			if evaluator.Handles(scope) {
				found = true
				allowedNamespaces, err := evaluator.ResolveGettableNamespaces(scope, clusterRoleGetter)
				if err != nil {
					errors = append(errors, err)
					continue
				}
				visibleNamespaces.Insert(allowedNamespaces...)
				break
			}
		}
		if !found && !ignoreUnhandledScopes {
			errors = append(errors, fmt.Errorf("no scope evaluator found for %q", scope))
		}
	}
	return visibleNamespaces, kutilerrors.NewAggregate(errors)
}

const (
	UserIndicator		= "user:"
	ClusterRoleIndicator	= "role:"
)

type ScopeEvaluator interface {
	Handles(scope string) bool
	Validate(scope string) error
	Describe(scope string) (description string, warning string, err error)
	ResolveRules(scope, namespace string, clusterRoleGetter rbaclisters.ClusterRoleLister) ([]rbacv1.PolicyRule, error)
	ResolveGettableNamespaces(scope string, clusterRoleGetter rbaclisters.ClusterRoleLister) ([]string, error)
}

var ScopeEvaluators = []ScopeEvaluator{userEvaluator{}, clusterRoleEvaluator{}}

const (
	UserInfo		= UserIndicator + "info"
	UserAccessCheck		= UserIndicator + "check-access"
	UserListScopedProjects	= UserIndicator + "list-scoped-projects"
	UserListAllProjects	= UserIndicator + "list-projects"
	UserFull		= UserIndicator + "full"
)

var defaultSupportedScopesMap = map[string]string{UserInfo: "Read-only access to your user information (including username, identities, and group membership)", UserAccessCheck: `Read-only access to view your privileges (for example, "can I create builds?")`, UserListScopedProjects: `Read-only access to list your projects viewable with this token and view their metadata (display name, description, etc.)`, UserListAllProjects: `Read-only access to list your projects and view their metadata (display name, description, etc.)`, UserFull: `Full read/write access with all of your permissions`}

func DefaultSupportedScopes() []string {
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
	return sets.StringKeySet(defaultSupportedScopesMap).List()
}
func DefaultSupportedScopesMap() map[string]string {
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
	return defaultSupportedScopesMap
}
func DescribeScopes(scopes []string) map[string]string {
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
	ret := map[string]string{}
	for _, s := range scopes {
		val, ok := defaultSupportedScopesMap[s]
		if ok {
			ret[s] = val
		} else {
			ret[s] = ""
		}
	}
	return ret
}

type userEvaluator struct{}

func (userEvaluator) Handles(scope string) bool {
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
	switch scope {
	case UserFull, UserInfo, UserAccessCheck, UserListScopedProjects, UserListAllProjects:
		return true
	}
	return false
}
func (e userEvaluator) Validate(scope string) error {
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
	if e.Handles(scope) {
		return nil
	}
	return fmt.Errorf("unrecognized scope: %v", scope)
}
func (userEvaluator) Describe(scope string) (string, string, error) {
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
	switch scope {
	case UserInfo, UserAccessCheck, UserListScopedProjects, UserListAllProjects:
		return defaultSupportedScopesMap[scope], "", nil
	case UserFull:
		return defaultSupportedScopesMap[scope], `Includes any access you have to escalating resources like secrets`, nil
	default:
		return "", "", fmt.Errorf("unrecognized scope: %v", scope)
	}
}
func (userEvaluator) ResolveRules(scope, namespace string, _ rbaclisters.ClusterRoleLister) ([]rbacv1.PolicyRule, error) {
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
	switch scope {
	case UserInfo:
		return []rbacv1.PolicyRule{rbacv1helpers.NewRule("get").Groups(userapi.GroupName, legacy.GroupName).Resources("users").Names("~").RuleOrDie()}, nil
	case UserAccessCheck:
		return []rbacv1.PolicyRule{rbacv1helpers.NewRule("create").Groups(kauthorizationapi.GroupName).Resources("selfsubjectaccessreviews").RuleOrDie(), rbacv1helpers.NewRule("create").Groups(authorizationapi.GroupName, legacy.GroupName).Resources("selfsubjectrulesreviews").RuleOrDie()}, nil
	case UserListScopedProjects:
		return []rbacv1.PolicyRule{rbacv1helpers.NewRule("list", "watch").Groups(projectapi.GroupName, legacy.GroupName).Resources("projects").RuleOrDie()}, nil
	case UserListAllProjects:
		return []rbacv1.PolicyRule{rbacv1helpers.NewRule("list", "watch").Groups(projectapi.GroupName, legacy.GroupName).Resources("projects").RuleOrDie(), rbacv1helpers.NewRule("get").Groups(kapi.GroupName).Resources("namespaces").RuleOrDie()}, nil
	case UserFull:
		return []rbacv1.PolicyRule{rbacv1helpers.NewRule(rbac.VerbAll).Groups(rbac.APIGroupAll).Resources(rbac.ResourceAll).RuleOrDie(), rbacv1helpers.NewRule(rbac.VerbAll).URLs(rbac.NonResourceAll).RuleOrDie()}, nil
	default:
		return nil, fmt.Errorf("unrecognized scope: %v", scope)
	}
}
func (userEvaluator) ResolveGettableNamespaces(scope string, _ rbaclisters.ClusterRoleLister) ([]string, error) {
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
	switch scope {
	case UserFull, UserListAllProjects:
		return []string{"*"}, nil
	default:
		return []string{}, nil
	}
}

var escalatingScopeResources = []schema.GroupResource{{Group: kapi.GroupName, Resource: "secrets"}, {Group: imageapi.GroupName, Resource: "imagestreams/secrets"}, {Group: legacy.GroupName, Resource: "imagestreams/secrets"}, {Group: oauthapi.GroupName, Resource: "oauthauthorizetokens"}, {Group: legacy.GroupName, Resource: "oauthauthorizetokens"}, {Group: oauthapi.GroupName, Resource: "oauthaccesstokens"}, {Group: legacy.GroupName, Resource: "oauthaccesstokens"}, {Group: authorizationapi.GroupName, Resource: "roles"}, {Group: legacy.GroupName, Resource: "roles"}, {Group: authorizationapi.GroupName, Resource: "rolebindings"}, {Group: legacy.GroupName, Resource: "rolebindings"}, {Group: authorizationapi.GroupName, Resource: "clusterroles"}, {Group: legacy.GroupName, Resource: "clusterroles"}, {Group: authorizationapi.GroupName, Resource: "clusterrolebindings"}, {Group: legacy.GroupName, Resource: "clusterrolebindings"}}

type clusterRoleEvaluator struct{}

var clusterRoleEvaluatorInstance = clusterRoleEvaluator{}

func (clusterRoleEvaluator) Handles(scope string) bool {
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
	return strings.HasPrefix(scope, ClusterRoleIndicator)
}
func (e clusterRoleEvaluator) Validate(scope string) error {
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
	_, _, _, err := e.parseScope(scope)
	return err
}
func (e clusterRoleEvaluator) parseScope(scope string) (string, string, bool, error) {
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
	if !e.Handles(scope) {
		return "", "", false, fmt.Errorf("bad format for scope %v", scope)
	}
	return ParseClusterRoleScope(scope)
}
func ParseClusterRoleScope(scope string) (string, string, bool, error) {
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
	if !strings.HasPrefix(scope, ClusterRoleIndicator) {
		return "", "", false, fmt.Errorf("bad format for scope %v", scope)
	}
	escalating := false
	if strings.HasSuffix(scope, ":!") {
		escalating = true
		scope = scope[:strings.LastIndex(scope, ":")]
	}
	tokens := strings.SplitN(scope, ":", 2)
	if len(tokens) != 2 {
		return "", "", false, fmt.Errorf("bad format for scope %v", scope)
	}
	lastColonIndex := strings.LastIndex(tokens[1], ":")
	if lastColonIndex <= 0 || lastColonIndex == (len(tokens[1])-1) {
		return "", "", false, fmt.Errorf("bad format for scope %v", scope)
	}
	return tokens[1][0:lastColonIndex], tokens[1][lastColonIndex+1:], escalating, nil
}
func (e clusterRoleEvaluator) Describe(scope string) (string, string, error) {
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
	roleName, scopeNamespace, escalating, err := e.parseScope(scope)
	if err != nil {
		return "", "", err
	}
	scopePhrase := ""
	if scopeNamespace == authorizationapi.ScopesAllNamespaces {
		scopePhrase = "server-wide"
	} else {
		scopePhrase = fmt.Sprintf("in project %q", scopeNamespace)
	}
	warning := ""
	escalatingPhrase := ""
	if escalating {
		warning = fmt.Sprintf("Includes access to escalating resources like secrets")
	} else {
		escalatingPhrase = ", except access escalating resources like secrets"
	}
	description := fmt.Sprintf("Anything you can do %s that is also allowed by the %q role%s", scopePhrase, roleName, escalatingPhrase)
	return description, warning, nil
}
func (e clusterRoleEvaluator) ResolveRules(scope, namespace string, clusterRoleGetter rbaclisters.ClusterRoleLister) ([]rbacv1.PolicyRule, error) {
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
	_, scopeNamespace, _, err := e.parseScope(scope)
	if err != nil {
		return nil, err
	}
	if !(scopeNamespace == authorizationapi.ScopesAllNamespaces || scopeNamespace == namespace) {
		return []rbacv1.PolicyRule{}, nil
	}
	return e.resolveRules(scope, clusterRoleGetter)
}
func has(set []string, value string) bool {
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
	for _, element := range set {
		if value == element {
			return true
		}
	}
	return false
}
func (e clusterRoleEvaluator) resolveRules(scope string, clusterRoleGetter rbaclisters.ClusterRoleLister) ([]rbacv1.PolicyRule, error) {
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
	roleName, _, escalating, err := e.parseScope(scope)
	if err != nil {
		return nil, err
	}
	role, err := clusterRoleGetter.Get(roleName)
	if err != nil {
		if kapierrors.IsNotFound(err) {
			return []rbacv1.PolicyRule{}, nil
		}
		return nil, err
	}
	rules := []rbacv1.PolicyRule{}
	for _, rule := range role.Rules {
		if escalating {
			rules = append(rules, rule)
			continue
		}
		if has(rule.Verbs, rbacv1.VerbAll) || has(rule.Resources, rbacv1.ResourceAll) || has(rule.APIGroups, rbacv1.APIGroupAll) {
			continue
		}
		safeRule := removeEscalatingResources(rule)
		rules = append(rules, safeRule)
	}
	return rules, nil
}
func (e clusterRoleEvaluator) ResolveGettableNamespaces(scope string, clusterRoleGetter rbaclisters.ClusterRoleLister) ([]string, error) {
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
	_, scopeNamespace, _, err := e.parseScope(scope)
	if err != nil {
		return nil, err
	}
	rules, err := e.resolveRules(scope, clusterRoleGetter)
	if err != nil {
		return nil, err
	}
	attributes := kauthorizer.AttributesRecord{APIGroup: kapi.GroupName, Verb: "get", Resource: "namespaces", ResourceRequest: true}
	if authorizerrbac.RulesAllow(attributes, rules...) {
		return []string{scopeNamespace}, nil
	}
	return []string{}, nil
}
func remove(array []string, item string) []string {
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
	newar := array[:0]
	for _, element := range array {
		if element != item {
			newar = append(newar, element)
		}
	}
	return newar
}
func removeEscalatingResources(in rbacv1.PolicyRule) rbacv1.PolicyRule {
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
	var ruleCopy *rbacv1.PolicyRule
	for _, resource := range escalatingScopeResources {
		if !(has(in.APIGroups, resource.Group) && has(in.Resources, resource.Resource)) {
			continue
		}
		if ruleCopy == nil {
			ruleCopy = in.DeepCopy()
		}
		ruleCopy.Resources = remove(ruleCopy.Resources, resource.Resource)
	}
	if ruleCopy != nil {
		return *ruleCopy
	}
	return in
}
func ValidateScopeRestrictions(client *oauthapi.OAuthClient, scopes ...string) error {
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
	if len(client.ScopeRestrictions) == 0 {
		return nil
	}
	if len(scopes) == 0 {
		return fmt.Errorf("%v may not request unscoped tokens", client.Name)
	}
	errs := []error{}
	for _, scope := range scopes {
		if err := validateScopeRestrictions(client, scope); err != nil {
			errs = append(errs, err)
		}
	}
	return kutilerrors.NewAggregate(errs)
}
func validateScopeRestrictions(client *oauthapi.OAuthClient, scope string) error {
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
	errs := []error{}
	if len(client.ScopeRestrictions) == 0 {
		return nil
	}
	for _, restriction := range client.ScopeRestrictions {
		if len(restriction.ExactValues) > 0 {
			if err := validateLiteralScopeRestrictions(scope, restriction.ExactValues); err != nil {
				errs = append(errs, err)
				continue
			}
			return nil
		}
		if restriction.ClusterRole != nil {
			if !clusterRoleEvaluatorInstance.Handles(scope) {
				continue
			}
			if err := validateClusterRoleScopeRestrictions(scope, *restriction.ClusterRole); err != nil {
				errs = append(errs, err)
				continue
			}
			return nil
		}
	}
	if len(errs) == 0 {
		errs = append(errs, fmt.Errorf("%v did not match any scope restriction", scope))
	}
	return kutilerrors.NewAggregate(errs)
}
func validateLiteralScopeRestrictions(scope string, literals []string) error {
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
	for _, literal := range literals {
		if literal == scope {
			return nil
		}
	}
	return fmt.Errorf("%v not found in %v", scope, literals)
}
func validateClusterRoleScopeRestrictions(scope string, restriction oauthapi.ClusterRoleScopeRestriction) error {
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
	role, namespace, escalating, err := clusterRoleEvaluatorInstance.parseScope(scope)
	if err != nil {
		return err
	}
	foundName := false
	for _, restrictedRoleName := range restriction.RoleNames {
		if restrictedRoleName == "*" || restrictedRoleName == role {
			foundName = true
			break
		}
	}
	if !foundName {
		return fmt.Errorf("%v does not use an approved name", scope)
	}
	foundNamespace := false
	for _, restrictedNamespace := range restriction.Namespaces {
		if restrictedNamespace == "*" || restrictedNamespace == namespace {
			foundNamespace = true
			break
		}
	}
	if !foundNamespace {
		return fmt.Errorf("%v does not use an approved namespace", scope)
	}
	if escalating && !restriction.AllowEscalation {
		return fmt.Errorf("%v is not allowed to escalate", scope)
	}
	return nil
}
