package v1

import (
	"fmt"
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sort"
)

type PolicyRuleBuilder struct {
	PolicyRule rbacv1.PolicyRule `protobuf:"bytes,1,opt,name=policyRule"`
}

func NewRule(verbs ...string) *PolicyRuleBuilder {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &PolicyRuleBuilder{PolicyRule: rbacv1.PolicyRule{Verbs: verbs}}
}
func (r *PolicyRuleBuilder) Groups(groups ...string) *PolicyRuleBuilder {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	r.PolicyRule.APIGroups = append(r.PolicyRule.APIGroups, groups...)
	return r
}
func (r *PolicyRuleBuilder) Resources(resources ...string) *PolicyRuleBuilder {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	r.PolicyRule.Resources = append(r.PolicyRule.Resources, resources...)
	return r
}
func (r *PolicyRuleBuilder) Names(names ...string) *PolicyRuleBuilder {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	r.PolicyRule.ResourceNames = append(r.PolicyRule.ResourceNames, names...)
	return r
}
func (r *PolicyRuleBuilder) URLs(urls ...string) *PolicyRuleBuilder {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	r.PolicyRule.NonResourceURLs = append(r.PolicyRule.NonResourceURLs, urls...)
	return r
}
func (r *PolicyRuleBuilder) RuleOrDie() rbacv1.PolicyRule {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	ret, err := r.Rule()
	if err != nil {
		panic(err)
	}
	return ret
}
func (r *PolicyRuleBuilder) Rule() (rbacv1.PolicyRule, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if len(r.PolicyRule.Verbs) == 0 {
		return rbacv1.PolicyRule{}, fmt.Errorf("verbs are required: %#v", r.PolicyRule)
	}
	switch {
	case len(r.PolicyRule.NonResourceURLs) > 0:
		if len(r.PolicyRule.APIGroups) != 0 || len(r.PolicyRule.Resources) != 0 || len(r.PolicyRule.ResourceNames) != 0 {
			return rbacv1.PolicyRule{}, fmt.Errorf("non-resource rule may not have apiGroups, resources, or resourceNames: %#v", r.PolicyRule)
		}
	case len(r.PolicyRule.Resources) > 0:
		if len(r.PolicyRule.NonResourceURLs) != 0 {
			return rbacv1.PolicyRule{}, fmt.Errorf("resource rule may not have nonResourceURLs: %#v", r.PolicyRule)
		}
		if len(r.PolicyRule.APIGroups) == 0 {
			return rbacv1.PolicyRule{}, fmt.Errorf("resource rule must have apiGroups: %#v", r.PolicyRule)
		}
	default:
		return rbacv1.PolicyRule{}, fmt.Errorf("a rule must have either nonResourceURLs or resources: %#v", r.PolicyRule)
	}
	sort.Strings(r.PolicyRule.Resources)
	sort.Strings(r.PolicyRule.ResourceNames)
	sort.Strings(r.PolicyRule.APIGroups)
	sort.Strings(r.PolicyRule.NonResourceURLs)
	sort.Strings(r.PolicyRule.Verbs)
	return r.PolicyRule, nil
}

type ClusterRoleBindingBuilder struct {
	ClusterRoleBinding rbacv1.ClusterRoleBinding `protobuf:"bytes,1,opt,name=clusterRoleBinding"`
}

func NewClusterBinding(clusterRoleName string) *ClusterRoleBindingBuilder {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &ClusterRoleBindingBuilder{ClusterRoleBinding: rbacv1.ClusterRoleBinding{ObjectMeta: metav1.ObjectMeta{Name: clusterRoleName}, RoleRef: rbacv1.RoleRef{APIGroup: GroupName, Kind: "ClusterRole", Name: clusterRoleName}}}
}
func (r *ClusterRoleBindingBuilder) Groups(groups ...string) *ClusterRoleBindingBuilder {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	for _, group := range groups {
		r.ClusterRoleBinding.Subjects = append(r.ClusterRoleBinding.Subjects, rbacv1.Subject{APIGroup: rbacv1.GroupName, Kind: rbacv1.GroupKind, Name: group})
	}
	return r
}
func (r *ClusterRoleBindingBuilder) Users(users ...string) *ClusterRoleBindingBuilder {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	for _, user := range users {
		r.ClusterRoleBinding.Subjects = append(r.ClusterRoleBinding.Subjects, rbacv1.Subject{APIGroup: rbacv1.GroupName, Kind: rbacv1.UserKind, Name: user})
	}
	return r
}
func (r *ClusterRoleBindingBuilder) SAs(namespace string, serviceAccountNames ...string) *ClusterRoleBindingBuilder {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	for _, saName := range serviceAccountNames {
		r.ClusterRoleBinding.Subjects = append(r.ClusterRoleBinding.Subjects, rbacv1.Subject{Kind: rbacv1.ServiceAccountKind, Namespace: namespace, Name: saName})
	}
	return r
}
func (r *ClusterRoleBindingBuilder) BindingOrDie() rbacv1.ClusterRoleBinding {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	ret, err := r.Binding()
	if err != nil {
		panic(err)
	}
	return ret
}
func (r *ClusterRoleBindingBuilder) Binding() (rbacv1.ClusterRoleBinding, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if len(r.ClusterRoleBinding.Subjects) == 0 {
		return rbacv1.ClusterRoleBinding{}, fmt.Errorf("subjects are required: %#v", r.ClusterRoleBinding)
	}
	return r.ClusterRoleBinding, nil
}

type RoleBindingBuilder struct{ RoleBinding rbacv1.RoleBinding }

func NewRoleBinding(roleName, namespace string) *RoleBindingBuilder {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &RoleBindingBuilder{RoleBinding: rbacv1.RoleBinding{ObjectMeta: metav1.ObjectMeta{Name: roleName, Namespace: namespace}, RoleRef: rbacv1.RoleRef{APIGroup: GroupName, Kind: "Role", Name: roleName}}}
}
func NewRoleBindingForClusterRole(roleName, namespace string) *RoleBindingBuilder {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &RoleBindingBuilder{RoleBinding: rbacv1.RoleBinding{ObjectMeta: metav1.ObjectMeta{Name: roleName, Namespace: namespace}, RoleRef: rbacv1.RoleRef{APIGroup: GroupName, Kind: "ClusterRole", Name: roleName}}}
}
func (r *RoleBindingBuilder) Groups(groups ...string) *RoleBindingBuilder {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	for _, group := range groups {
		r.RoleBinding.Subjects = append(r.RoleBinding.Subjects, rbacv1.Subject{Kind: rbacv1.GroupKind, APIGroup: GroupName, Name: group})
	}
	return r
}
func (r *RoleBindingBuilder) Users(users ...string) *RoleBindingBuilder {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	for _, user := range users {
		r.RoleBinding.Subjects = append(r.RoleBinding.Subjects, rbacv1.Subject{Kind: rbacv1.UserKind, APIGroup: GroupName, Name: user})
	}
	return r
}
func (r *RoleBindingBuilder) SAs(namespace string, serviceAccountNames ...string) *RoleBindingBuilder {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	for _, saName := range serviceAccountNames {
		r.RoleBinding.Subjects = append(r.RoleBinding.Subjects, rbacv1.Subject{Kind: rbacv1.ServiceAccountKind, Namespace: namespace, Name: saName})
	}
	return r
}
func (r *RoleBindingBuilder) BindingOrDie() rbacv1.RoleBinding {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	ret, err := r.Binding()
	if err != nil {
		panic(err)
	}
	return ret
}
func (r *RoleBindingBuilder) Binding() (rbacv1.RoleBinding, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if len(r.RoleBinding.Subjects) == 0 {
		return rbacv1.RoleBinding{}, fmt.Errorf("subjects are required: %#v", r.RoleBinding)
	}
	return r.RoleBinding, nil
}
