package restrictusers

import (
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	kerrors "k8s.io/apimachinery/pkg/util/errors"
	"k8s.io/client-go/kubernetes"
	"k8s.io/kubernetes/pkg/apis/rbac"
	authorizationapi "github.com/openshift/api/authorization/v1"
	userapi "github.com/openshift/api/user/v1"
	userclient "github.com/openshift/client-go/user/clientset/versioned/typed/user/v1"
)

type SubjectChecker interface {
	Allowed(rbac.Subject, *RoleBindingRestrictionContext) (bool, error)
}
type UnionSubjectChecker []SubjectChecker

func NewUnionSubjectChecker(checkers []SubjectChecker) UnionSubjectChecker {
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
	return UnionSubjectChecker(checkers)
}
func (checkers UnionSubjectChecker) Allowed(subject rbac.Subject, ctx *RoleBindingRestrictionContext) (bool, error) {
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
	for _, checker := range []SubjectChecker(checkers) {
		allowed, err := checker.Allowed(subject, ctx)
		if err != nil {
			errs = append(errs, err)
		} else if allowed {
			return true, nil
		}
	}
	return false, kerrors.NewAggregate(errs)
}

type RoleBindingRestrictionContext struct {
	userClient	userclient.UserV1Interface
	kclient		kubernetes.Interface
	groupCache	GroupCache
	userToLabelSet	map[string]labels.Set
	groupToLabelSet	map[string]labels.Set
	namespace	string
}

func newRoleBindingRestrictionContext(ns string, kc kubernetes.Interface, userClient userclient.UserV1Interface, groupCache GroupCache) (*RoleBindingRestrictionContext, error) {
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
	return &RoleBindingRestrictionContext{namespace: ns, kclient: kc, userClient: userClient, groupCache: groupCache, userToLabelSet: map[string]labels.Set{}, groupToLabelSet: map[string]labels.Set{}}, nil
}
func (ctx *RoleBindingRestrictionContext) labelSetForUser(subject rbac.Subject) (labels.Set, error) {
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
	if subject.Kind != rbac.UserKind {
		return labels.Set{}, fmt.Errorf("not a user: %q", subject.Name)
	}
	labelSet, ok := ctx.userToLabelSet[subject.Name]
	if ok {
		return labelSet, nil
	}
	user, err := ctx.userClient.Users().Get(subject.Name, metav1.GetOptions{})
	if err != nil {
		return labels.Set{}, err
	}
	ctx.userToLabelSet[subject.Name] = labels.Set(user.Labels)
	return ctx.userToLabelSet[subject.Name], nil
}
func (ctx *RoleBindingRestrictionContext) groupsForUser(subject rbac.Subject) ([]*userapi.Group, error) {
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
	if subject.Kind != rbac.UserKind {
		return []*userapi.Group{}, fmt.Errorf("not a user: %q", subject.Name)
	}
	return ctx.groupCache.GroupsFor(subject.Name)
}
func (ctx *RoleBindingRestrictionContext) labelSetForGroup(subject rbac.Subject) (labels.Set, error) {
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
	if subject.Kind != rbac.GroupKind {
		return labels.Set{}, fmt.Errorf("not a group: %q", subject.Name)
	}
	labelSet, ok := ctx.groupToLabelSet[subject.Name]
	if ok {
		return labelSet, nil
	}
	group, err := ctx.userClient.Groups().Get(subject.Name, metav1.GetOptions{})
	if err != nil {
		return labels.Set{}, err
	}
	ctx.groupToLabelSet[subject.Name] = labels.Set(group.Labels)
	return ctx.groupToLabelSet[subject.Name], nil
}

type UserSubjectChecker struct {
	userRestriction *authorizationapi.UserRestriction
}

func NewUserSubjectChecker(userRestriction *authorizationapi.UserRestriction) UserSubjectChecker {
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
	return UserSubjectChecker{userRestriction: userRestriction}
}
func (checker UserSubjectChecker) Allowed(subject rbac.Subject, ctx *RoleBindingRestrictionContext) (bool, error) {
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
	if subject.Kind != rbac.UserKind {
		return false, nil
	}
	for _, userName := range checker.userRestriction.Users {
		if subject.Name == userName {
			return true, nil
		}
	}
	if len(checker.userRestriction.Groups) != 0 {
		subjectGroups, err := ctx.groupsForUser(subject)
		if err != nil {
			return false, err
		}
		for _, groupName := range checker.userRestriction.Groups {
			for _, group := range subjectGroups {
				if group.Name == groupName {
					return true, nil
				}
			}
		}
	}
	if len(checker.userRestriction.Selectors) != 0 {
		labelSet, err := ctx.labelSetForUser(subject)
		if err != nil {
			return false, err
		}
		for _, labelSelector := range checker.userRestriction.Selectors {
			selector, err := metav1.LabelSelectorAsSelector(&labelSelector)
			if err != nil {
				return false, err
			}
			if selector.Matches(labelSet) {
				return true, nil
			}
		}
	}
	return false, nil
}

type GroupSubjectChecker struct {
	groupRestriction *authorizationapi.GroupRestriction
}

func NewGroupSubjectChecker(groupRestriction *authorizationapi.GroupRestriction) GroupSubjectChecker {
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
	return GroupSubjectChecker{groupRestriction: groupRestriction}
}
func (checker GroupSubjectChecker) Allowed(subject rbac.Subject, ctx *RoleBindingRestrictionContext) (bool, error) {
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
	if subject.Kind != rbac.GroupKind {
		return false, nil
	}
	for _, groupName := range checker.groupRestriction.Groups {
		if subject.Name == groupName {
			return true, nil
		}
	}
	if len(checker.groupRestriction.Selectors) != 0 {
		labelSet, err := ctx.labelSetForGroup(subject)
		if err != nil {
			return false, err
		}
		for _, labelSelector := range checker.groupRestriction.Selectors {
			selector, err := metav1.LabelSelectorAsSelector(&labelSelector)
			if err != nil {
				return false, err
			}
			if selector.Matches(labelSet) {
				return true, nil
			}
		}
	}
	return false, nil
}

type ServiceAccountSubjectChecker struct {
	serviceAccountRestriction *authorizationapi.ServiceAccountRestriction
}

func NewServiceAccountSubjectChecker(serviceAccountRestriction *authorizationapi.ServiceAccountRestriction) ServiceAccountSubjectChecker {
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
	return ServiceAccountSubjectChecker{serviceAccountRestriction: serviceAccountRestriction}
}
func (checker ServiceAccountSubjectChecker) Allowed(subject rbac.Subject, ctx *RoleBindingRestrictionContext) (bool, error) {
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
	if subject.Kind != rbac.ServiceAccountKind {
		return false, nil
	}
	subjectNamespace := subject.Namespace
	if len(subjectNamespace) == 0 {
		subjectNamespace = ctx.namespace
	}
	for _, namespace := range checker.serviceAccountRestriction.Namespaces {
		if subjectNamespace == namespace {
			return true, nil
		}
	}
	for _, serviceAccountRef := range checker.serviceAccountRestriction.ServiceAccounts {
		serviceAccountNamespace := serviceAccountRef.Namespace
		if len(serviceAccountNamespace) == 0 {
			serviceAccountNamespace = ctx.namespace
		}
		if subject.Name == serviceAccountRef.Name && subjectNamespace == serviceAccountNamespace {
			return true, nil
		}
	}
	return false, nil
}
func NewSubjectChecker(spec *authorizationapi.RoleBindingRestrictionSpec) (SubjectChecker, error) {
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
	switch {
	case spec.UserRestriction != nil:
		return NewUserSubjectChecker(spec.UserRestriction), nil
	case spec.GroupRestriction != nil:
		return NewGroupSubjectChecker(spec.GroupRestriction), nil
	case spec.ServiceAccountRestriction != nil:
		return NewServiceAccountSubjectChecker(spec.ServiceAccountRestriction), nil
	}
	return nil, fmt.Errorf("invalid RoleBindingRestrictionSpec: %v", spec)
}
