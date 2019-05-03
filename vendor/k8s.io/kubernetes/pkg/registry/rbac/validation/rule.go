package validation

import (
 "context"
 "errors"
 "fmt"
 "strings"
 "k8s.io/klog"
 rbacv1 "k8s.io/api/rbac/v1"
 utilerrors "k8s.io/apimachinery/pkg/util/errors"
 "k8s.io/apimachinery/pkg/util/sets"
 "k8s.io/apiserver/pkg/authentication/serviceaccount"
 "k8s.io/apiserver/pkg/authentication/user"
 genericapirequest "k8s.io/apiserver/pkg/endpoints/request"
 rbacv1helpers "k8s.io/kubernetes/pkg/apis/rbac/v1"
)

type AuthorizationRuleResolver interface {
 GetRoleReferenceRules(roleRef rbacv1.RoleRef, namespace string) ([]rbacv1.PolicyRule, error)
 RulesFor(user user.Info, namespace string) ([]rbacv1.PolicyRule, error)
 VisitRulesFor(user user.Info, namespace string, visitor func(source fmt.Stringer, rule *rbacv1.PolicyRule, err error) bool)
}

func ConfirmNoEscalation(ctx context.Context, ruleResolver AuthorizationRuleResolver, rules []rbacv1.PolicyRule) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 ruleResolutionErrors := []error{}
 user, ok := genericapirequest.UserFrom(ctx)
 if !ok {
  return fmt.Errorf("no user on context")
 }
 namespace, _ := genericapirequest.NamespaceFrom(ctx)
 ownerRules, err := ruleResolver.RulesFor(user, namespace)
 if err != nil {
  klog.V(1).Infof("non-fatal error getting local rules for %v: %v", user, err)
  ruleResolutionErrors = append(ruleResolutionErrors, err)
 }
 ownerRightsCover, missingRights := Covers(ownerRules, rules)
 if !ownerRightsCover {
  compactMissingRights := missingRights
  if compact, err := CompactRules(missingRights); err == nil {
   compactMissingRights = compact
  }
  missingDescriptions := sets.NewString()
  for _, missing := range compactMissingRights {
   missingDescriptions.Insert(rbacv1helpers.CompactString(missing))
  }
  msg := fmt.Sprintf("user %q (groups=%q) is attempting to grant RBAC permissions not currently held:\n%s", user.GetName(), user.GetGroups(), strings.Join(missingDescriptions.List(), "\n"))
  if len(ruleResolutionErrors) > 0 {
   msg = msg + fmt.Sprintf("; resolution errors: %v", ruleResolutionErrors)
  }
  return errors.New(msg)
 }
 return nil
}

type DefaultRuleResolver struct {
 roleGetter               RoleGetter
 roleBindingLister        RoleBindingLister
 clusterRoleGetter        ClusterRoleGetter
 clusterRoleBindingLister ClusterRoleBindingLister
}

func NewDefaultRuleResolver(roleGetter RoleGetter, roleBindingLister RoleBindingLister, clusterRoleGetter ClusterRoleGetter, clusterRoleBindingLister ClusterRoleBindingLister) *DefaultRuleResolver {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &DefaultRuleResolver{roleGetter, roleBindingLister, clusterRoleGetter, clusterRoleBindingLister}
}

type RoleGetter interface {
 GetRole(namespace, name string) (*rbacv1.Role, error)
}
type RoleBindingLister interface {
 ListRoleBindings(namespace string) ([]*rbacv1.RoleBinding, error)
}
type ClusterRoleGetter interface {
 GetClusterRole(name string) (*rbacv1.ClusterRole, error)
}
type ClusterRoleBindingLister interface {
 ListClusterRoleBindings() ([]*rbacv1.ClusterRoleBinding, error)
}

func (r *DefaultRuleResolver) RulesFor(user user.Info, namespace string) ([]rbacv1.PolicyRule, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 visitor := &ruleAccumulator{}
 r.VisitRulesFor(user, namespace, visitor.visit)
 return visitor.rules, utilerrors.NewAggregate(visitor.errors)
}

type ruleAccumulator struct {
 rules  []rbacv1.PolicyRule
 errors []error
}

func (r *ruleAccumulator) visit(source fmt.Stringer, rule *rbacv1.PolicyRule, err error) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if rule != nil {
  r.rules = append(r.rules, *rule)
 }
 if err != nil {
  r.errors = append(r.errors, err)
 }
 return true
}
func describeSubject(s *rbacv1.Subject, bindingNamespace string) string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 switch s.Kind {
 case rbacv1.ServiceAccountKind:
  if len(s.Namespace) > 0 {
   return fmt.Sprintf("%s %q", s.Kind, s.Name+"/"+s.Namespace)
  }
  return fmt.Sprintf("%s %q", s.Kind, s.Name+"/"+bindingNamespace)
 default:
  return fmt.Sprintf("%s %q", s.Kind, s.Name)
 }
}

type clusterRoleBindingDescriber struct {
 binding *rbacv1.ClusterRoleBinding
 subject *rbacv1.Subject
}

func (d *clusterRoleBindingDescriber) String() string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return fmt.Sprintf("ClusterRoleBinding %q of %s %q to %s", d.binding.Name, d.binding.RoleRef.Kind, d.binding.RoleRef.Name, describeSubject(d.subject, ""))
}

type roleBindingDescriber struct {
 binding *rbacv1.RoleBinding
 subject *rbacv1.Subject
}

func (d *roleBindingDescriber) String() string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return fmt.Sprintf("RoleBinding %q of %s %q to %s", d.binding.Name+"/"+d.binding.Namespace, d.binding.RoleRef.Kind, d.binding.RoleRef.Name, describeSubject(d.subject, d.binding.Namespace))
}
func (r *DefaultRuleResolver) VisitRulesFor(user user.Info, namespace string, visitor func(source fmt.Stringer, rule *rbacv1.PolicyRule, err error) bool) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if clusterRoleBindings, err := r.clusterRoleBindingLister.ListClusterRoleBindings(); err != nil {
  if !visitor(nil, nil, err) {
   return
  }
 } else {
  sourceDescriber := &clusterRoleBindingDescriber{}
  for _, clusterRoleBinding := range clusterRoleBindings {
   subjectIndex, applies := appliesTo(user, clusterRoleBinding.Subjects, "")
   if !applies {
    continue
   }
   rules, err := r.GetRoleReferenceRules(clusterRoleBinding.RoleRef, "")
   if err != nil {
    if !visitor(nil, nil, err) {
     return
    }
    continue
   }
   sourceDescriber.binding = clusterRoleBinding
   sourceDescriber.subject = &clusterRoleBinding.Subjects[subjectIndex]
   for i := range rules {
    if !visitor(sourceDescriber, &rules[i], nil) {
     return
    }
   }
  }
 }
 if len(namespace) > 0 {
  if roleBindings, err := r.roleBindingLister.ListRoleBindings(namespace); err != nil {
   if !visitor(nil, nil, err) {
    return
   }
  } else {
   sourceDescriber := &roleBindingDescriber{}
   for _, roleBinding := range roleBindings {
    subjectIndex, applies := appliesTo(user, roleBinding.Subjects, namespace)
    if !applies {
     continue
    }
    rules, err := r.GetRoleReferenceRules(roleBinding.RoleRef, namespace)
    if err != nil {
     if !visitor(nil, nil, err) {
      return
     }
     continue
    }
    sourceDescriber.binding = roleBinding
    sourceDescriber.subject = &roleBinding.Subjects[subjectIndex]
    for i := range rules {
     if !visitor(sourceDescriber, &rules[i], nil) {
      return
     }
    }
   }
  }
 }
}
func (r *DefaultRuleResolver) GetRoleReferenceRules(roleRef rbacv1.RoleRef, bindingNamespace string) ([]rbacv1.PolicyRule, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 switch roleRef.Kind {
 case "Role":
  role, err := r.roleGetter.GetRole(bindingNamespace, roleRef.Name)
  if err != nil {
   return nil, err
  }
  return role.Rules, nil
 case "ClusterRole":
  clusterRole, err := r.clusterRoleGetter.GetClusterRole(roleRef.Name)
  if err != nil {
   return nil, err
  }
  return clusterRole.Rules, nil
 default:
  return nil, fmt.Errorf("unsupported role reference kind: %q", roleRef.Kind)
 }
}
func appliesTo(user user.Info, bindingSubjects []rbacv1.Subject, namespace string) (int, bool) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 for i, bindingSubject := range bindingSubjects {
  if appliesToUser(user, bindingSubject, namespace) {
   return i, true
  }
 }
 return 0, false
}
func appliesToUser(user user.Info, subject rbacv1.Subject, namespace string) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 switch subject.Kind {
 case rbacv1.UserKind:
  return user.GetName() == subject.Name
 case rbacv1.GroupKind:
  return has(user.GetGroups(), subject.Name)
 case rbacv1.ServiceAccountKind:
  saNamespace := namespace
  if len(subject.Namespace) > 0 {
   saNamespace = subject.Namespace
  }
  if len(saNamespace) == 0 {
   return false
  }
  return serviceaccount.MatchesUsername(saNamespace, subject.Name, user.GetName())
 default:
  return false
 }
}
func NewTestRuleResolver(roles []*rbacv1.Role, roleBindings []*rbacv1.RoleBinding, clusterRoles []*rbacv1.ClusterRole, clusterRoleBindings []*rbacv1.ClusterRoleBinding) (AuthorizationRuleResolver, *StaticRoles) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 r := StaticRoles{roles: roles, roleBindings: roleBindings, clusterRoles: clusterRoles, clusterRoleBindings: clusterRoleBindings}
 return newMockRuleResolver(&r), &r
}
func newMockRuleResolver(r *StaticRoles) AuthorizationRuleResolver {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return NewDefaultRuleResolver(r, r, r, r)
}

type StaticRoles struct {
 roles               []*rbacv1.Role
 roleBindings        []*rbacv1.RoleBinding
 clusterRoles        []*rbacv1.ClusterRole
 clusterRoleBindings []*rbacv1.ClusterRoleBinding
}

func (r *StaticRoles) GetRole(namespace, name string) (*rbacv1.Role, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if len(namespace) == 0 {
  return nil, errors.New("must provide namespace when getting role")
 }
 for _, role := range r.roles {
  if role.Namespace == namespace && role.Name == name {
   return role, nil
  }
 }
 return nil, errors.New("role not found")
}
func (r *StaticRoles) GetClusterRole(name string) (*rbacv1.ClusterRole, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 for _, clusterRole := range r.clusterRoles {
  if clusterRole.Name == name {
   return clusterRole, nil
  }
 }
 return nil, errors.New("clusterrole not found")
}
func (r *StaticRoles) ListRoleBindings(namespace string) ([]*rbacv1.RoleBinding, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if len(namespace) == 0 {
  return nil, errors.New("must provide namespace when listing role bindings")
 }
 roleBindingList := []*rbacv1.RoleBinding{}
 for _, roleBinding := range r.roleBindings {
  if roleBinding.Namespace != namespace {
   continue
  }
  roleBindingList = append(roleBindingList, roleBinding)
 }
 return roleBindingList, nil
}
func (r *StaticRoles) ListClusterRoleBindings() ([]*rbacv1.ClusterRoleBinding, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return r.clusterRoleBindings, nil
}
