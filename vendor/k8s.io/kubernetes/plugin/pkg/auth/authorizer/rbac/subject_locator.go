package rbac

import (
	rbacv1 "k8s.io/api/rbac/v1"
	utilerrors "k8s.io/apimachinery/pkg/util/errors"
	"k8s.io/apiserver/pkg/authentication/user"
	"k8s.io/apiserver/pkg/authorization/authorizer"
	rbacregistryvalidation "k8s.io/kubernetes/pkg/registry/rbac/validation"
)

type RoleToRuleMapper interface {
	GetRoleReferenceRules(roleRef rbacv1.RoleRef, namespace string) ([]rbacv1.PolicyRule, error)
}
type SubjectLocator interface {
	AllowedSubjects(attributes authorizer.Attributes) ([]rbacv1.Subject, error)
}

var _ = SubjectLocator(&SubjectAccessEvaluator{})

type SubjectAccessEvaluator struct {
	superUser                string
	roleBindingLister        rbacregistryvalidation.RoleBindingLister
	clusterRoleBindingLister rbacregistryvalidation.ClusterRoleBindingLister
	roleToRuleMapper         RoleToRuleMapper
}

func NewSubjectAccessEvaluator(roles rbacregistryvalidation.RoleGetter, roleBindings rbacregistryvalidation.RoleBindingLister, clusterRoles rbacregistryvalidation.ClusterRoleGetter, clusterRoleBindings rbacregistryvalidation.ClusterRoleBindingLister, superUser string) *SubjectAccessEvaluator {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	subjectLocator := &SubjectAccessEvaluator{superUser: superUser, roleBindingLister: roleBindings, clusterRoleBindingLister: clusterRoleBindings, roleToRuleMapper: rbacregistryvalidation.NewDefaultRuleResolver(roles, roleBindings, clusterRoles, clusterRoleBindings)}
	return subjectLocator
}
func (r *SubjectAccessEvaluator) AllowedSubjects(requestAttributes authorizer.Attributes) ([]rbacv1.Subject, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	subjects := []rbacv1.Subject{{Kind: rbacv1.GroupKind, APIGroup: rbacv1.GroupName, Name: user.SystemPrivilegedGroup}}
	if len(r.superUser) > 0 {
		subjects = append(subjects, rbacv1.Subject{Kind: rbacv1.UserKind, APIGroup: rbacv1.GroupName, Name: r.superUser})
	}
	errorlist := []error{}
	if clusterRoleBindings, err := r.clusterRoleBindingLister.ListClusterRoleBindings(); err != nil {
		errorlist = append(errorlist, err)
	} else {
		for _, clusterRoleBinding := range clusterRoleBindings {
			rules, err := r.roleToRuleMapper.GetRoleReferenceRules(clusterRoleBinding.RoleRef, "")
			if err != nil {
				errorlist = append(errorlist, err)
			}
			if RulesAllow(requestAttributes, rules...) {
				subjects = append(subjects, clusterRoleBinding.Subjects...)
			}
		}
	}
	if namespace := requestAttributes.GetNamespace(); len(namespace) > 0 {
		if roleBindings, err := r.roleBindingLister.ListRoleBindings(namespace); err != nil {
			errorlist = append(errorlist, err)
		} else {
			for _, roleBinding := range roleBindings {
				rules, err := r.roleToRuleMapper.GetRoleReferenceRules(roleBinding.RoleRef, namespace)
				if err != nil {
					errorlist = append(errorlist, err)
				}
				if RulesAllow(requestAttributes, rules...) {
					subjects = append(subjects, roleBinding.Subjects...)
				}
			}
		}
	}
	dedupedSubjects := []rbacv1.Subject{}
	for _, subject := range subjects {
		found := false
		for _, curr := range dedupedSubjects {
			if curr == subject {
				found = true
				break
			}
		}
		if !found {
			dedupedSubjects = append(dedupedSubjects, subject)
		}
	}
	return subjects, utilerrors.NewAggregate(errorlist)
}
