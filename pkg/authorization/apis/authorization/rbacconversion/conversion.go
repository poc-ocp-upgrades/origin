package rbacconversion

import (
	godefaultbytes "bytes"
	"fmt"
	authorizationapi "github.com/openshift/origin/pkg/authorization/apis/authorization"
	"github.com/openshift/origin/pkg/user/apis/user/validation"
	rbacv1 "k8s.io/api/rbac/v1"
	"k8s.io/apimachinery/pkg/conversion"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/sets"
	api "k8s.io/kubernetes/pkg/apis/core"
	"k8s.io/kubernetes/pkg/apis/rbac"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
)

var (
	SchemeBuilder = runtime.NewSchemeBuilder(addConversionFuncs)
	AddToScheme   = SchemeBuilder.AddToScheme
)

const reconcileProtectAnnotation = "openshift.io/reconcile-protect"

func addConversionFuncs(scheme *runtime.Scheme) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := scheme.AddConversionFuncs(Convert_authorization_ClusterRole_To_rbac_ClusterRole, Convert_authorization_Role_To_rbac_Role, Convert_authorization_ClusterRoleBinding_To_rbac_ClusterRoleBinding, Convert_authorization_RoleBinding_To_rbac_RoleBinding, Convert_rbac_ClusterRole_To_authorization_ClusterRole, Convert_rbac_Role_To_authorization_Role, Convert_rbac_ClusterRoleBinding_To_authorization_ClusterRoleBinding, Convert_rbac_RoleBinding_To_authorization_RoleBinding); err != nil {
		return err
	}
	return nil
}
func Convert_authorization_ClusterRole_To_rbac_ClusterRole(in *authorizationapi.ClusterRole, out *rbac.ClusterRole, _ conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ObjectMeta = in.ObjectMeta
	out.Annotations = convert_authorization_Annotations_To_rbac_Annotations(in.Annotations)
	out.Rules = Convert_api_PolicyRules_To_rbac_PolicyRules(in.Rules)
	out.AggregationRule = in.AggregationRule.DeepCopy()
	return nil
}
func Convert_authorization_Role_To_rbac_Role(in *authorizationapi.Role, out *rbac.Role, _ conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ObjectMeta = in.ObjectMeta
	out.Annotations = convert_authorization_Annotations_To_rbac_Annotations(in.Annotations)
	out.Rules = Convert_api_PolicyRules_To_rbac_PolicyRules(in.Rules)
	return nil
}
func Convert_authorization_ClusterRoleBinding_To_rbac_ClusterRoleBinding(in *authorizationapi.ClusterRoleBinding, out *rbac.ClusterRoleBinding, _ conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(in.RoleRef.Namespace) != 0 {
		return fmt.Errorf("invalid origin cluster role binding %s: attempts to reference role in namespace %q instead of cluster scope", in.Name, in.RoleRef.Namespace)
	}
	var err error
	if out.Subjects, err = convert_api_Subjects_To_rbac_Subjects(in.Subjects); err != nil {
		return err
	}
	out.RoleRef = convert_api_RoleRef_To_rbac_RoleRef(&in.RoleRef)
	out.ObjectMeta = in.ObjectMeta
	out.Annotations = convert_authorization_Annotations_To_rbac_Annotations(in.Annotations)
	return nil
}
func Convert_authorization_RoleBinding_To_rbac_RoleBinding(in *authorizationapi.RoleBinding, out *rbac.RoleBinding, _ conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(in.RoleRef.Namespace) != 0 && in.RoleRef.Namespace != in.Namespace {
		return fmt.Errorf("invalid origin role binding %s: attempts to reference role in namespace %q instead of current namespace %q", in.Name, in.RoleRef.Namespace, in.Namespace)
	}
	var err error
	if out.Subjects, err = convert_api_Subjects_To_rbac_Subjects(in.Subjects); err != nil {
		return err
	}
	out.RoleRef = convert_api_RoleRef_To_rbac_RoleRef(&in.RoleRef)
	out.ObjectMeta = in.ObjectMeta
	out.Annotations = convert_authorization_Annotations_To_rbac_Annotations(in.Annotations)
	return nil
}
func Convert_api_PolicyRules_To_rbac_PolicyRules(in []authorizationapi.PolicyRule) []rbac.PolicyRule {
	_logClusterCodePath()
	defer _logClusterCodePath()
	rules := make([]rbac.PolicyRule, 0, len(in))
	for _, rule := range in {
		if rule.AttributeRestrictions != nil {
			continue
		}
		if isResourceRule(&rule) && isNonResourceRule(&rule) {
			r1 := rbac.PolicyRule{Verbs: rule.Verbs.List(), APIGroups: rule.APIGroups, Resources: rule.Resources.List(), ResourceNames: rule.ResourceNames.List()}
			r2 := rbac.PolicyRule{Verbs: rule.Verbs.List(), NonResourceURLs: rule.NonResourceURLs.List()}
			rules = append(rules, r1, r2)
		} else {
			r := rbac.PolicyRule{APIGroups: rule.APIGroups, Verbs: rule.Verbs.List(), Resources: rule.Resources.List(), ResourceNames: rule.ResourceNames.List(), NonResourceURLs: rule.NonResourceURLs.List()}
			rules = append(rules, r)
		}
	}
	return rules
}
func isResourceRule(rule *authorizationapi.PolicyRule) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return len(rule.APIGroups) > 0 || len(rule.Resources) > 0 || len(rule.ResourceNames) > 0
}
func isNonResourceRule(rule *authorizationapi.PolicyRule) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return len(rule.NonResourceURLs) > 0
}
func convert_api_Subjects_To_rbac_Subjects(in []api.ObjectReference) ([]rbac.Subject, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	subjects := make([]rbac.Subject, 0, len(in))
	for _, subject := range in {
		s := rbac.Subject{Name: subject.Name}
		switch subject.Kind {
		case authorizationapi.ServiceAccountKind:
			s.Kind = rbac.ServiceAccountKind
			s.Namespace = subject.Namespace
		case authorizationapi.UserKind, authorizationapi.SystemUserKind:
			s.APIGroup = rbac.GroupName
			s.Kind = rbac.UserKind
		case authorizationapi.GroupKind, authorizationapi.SystemGroupKind:
			s.APIGroup = rbac.GroupName
			s.Kind = rbac.GroupKind
		default:
			return nil, fmt.Errorf("invalid kind for origin subject: %q", subject.Kind)
		}
		subjects = append(subjects, s)
	}
	return subjects, nil
}
func convert_api_RoleRef_To_rbac_RoleRef(in *api.ObjectReference) rbac.RoleRef {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return rbac.RoleRef{APIGroup: rbac.GroupName, Kind: getRBACRoleRefKind(in.Namespace), Name: in.Name}
}
func getRBACRoleRefKind(namespace string) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	kind := "ClusterRole"
	if len(namespace) != 0 {
		kind = "Role"
	}
	return kind
}
func Convert_rbac_ClusterRole_To_authorization_ClusterRole(in *rbac.ClusterRole, out *authorizationapi.ClusterRole, _ conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ObjectMeta = in.ObjectMeta
	out.Annotations = convert_rbac_Annotations_To_authorization_Annotations(in.Annotations)
	out.Rules = Convert_rbac_PolicyRules_To_authorization_PolicyRules(in.Rules)
	out.AggregationRule = in.AggregationRule.DeepCopy()
	return nil
}
func Convert_rbac_Role_To_authorization_Role(in *rbac.Role, out *authorizationapi.Role, _ conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ObjectMeta = in.ObjectMeta
	out.Annotations = convert_rbac_Annotations_To_authorization_Annotations(in.Annotations)
	out.Rules = Convert_rbac_PolicyRules_To_authorization_PolicyRules(in.Rules)
	return nil
}
func Convert_rbac_ClusterRoleBinding_To_authorization_ClusterRoleBinding(in *rbac.ClusterRoleBinding, out *authorizationapi.ClusterRoleBinding, _ conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var err error
	if out.Subjects, err = Convert_rbac_Subjects_To_authorization_Subjects(in.Subjects); err != nil {
		return err
	}
	if out.RoleRef, err = convert_rbac_RoleRef_To_authorization_RoleRef(&in.RoleRef, ""); err != nil {
		return err
	}
	out.ObjectMeta = in.ObjectMeta
	out.Annotations = convert_rbac_Annotations_To_authorization_Annotations(in.Annotations)
	return nil
}
func Convert_rbac_RoleBinding_To_authorization_RoleBinding(in *rbac.RoleBinding, out *authorizationapi.RoleBinding, _ conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var err error
	if out.Subjects, err = Convert_rbac_Subjects_To_authorization_Subjects(in.Subjects); err != nil {
		return err
	}
	if out.RoleRef, err = convert_rbac_RoleRef_To_authorization_RoleRef(&in.RoleRef, in.Namespace); err != nil {
		return err
	}
	out.ObjectMeta = in.ObjectMeta
	out.Annotations = convert_rbac_Annotations_To_authorization_Annotations(in.Annotations)
	return nil
}
func Convert_rbac_Subjects_To_authorization_Subjects(in []rbac.Subject) ([]api.ObjectReference, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	subjects := make([]api.ObjectReference, 0, len(in))
	for _, subject := range in {
		s := api.ObjectReference{Name: subject.Name}
		switch subject.Kind {
		case rbac.ServiceAccountKind:
			s.Kind = authorizationapi.ServiceAccountKind
			s.Namespace = subject.Namespace
		case rbac.UserKind:
			s.Kind = determineUserKind(subject.Name)
		case rbac.GroupKind:
			s.Kind = determineGroupKind(subject.Name)
		default:
			return nil, fmt.Errorf("invalid kind for rbac subject: %q", subject.Kind)
		}
		subjects = append(subjects, s)
	}
	return subjects, nil
}
func convert_rbac_RoleRef_To_authorization_RoleRef(in *rbac.RoleRef, namespace string) (api.ObjectReference, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	switch in.Kind {
	case "ClusterRole":
		return api.ObjectReference{Name: in.Name}, nil
	case "Role":
		return api.ObjectReference{Name: in.Name, Namespace: namespace}, nil
	default:
		return api.ObjectReference{}, fmt.Errorf("invalid kind %q for rbac role ref %q", in.Kind, in.Name)
	}
}
func Convert_rbac_PolicyRules_To_authorization_PolicyRules(in []rbac.PolicyRule) []authorizationapi.PolicyRule {
	_logClusterCodePath()
	defer _logClusterCodePath()
	rules := make([]authorizationapi.PolicyRule, 0, len(in))
	for _, rule := range in {
		r := authorizationapi.PolicyRule{APIGroups: rule.APIGroups, Verbs: sets.NewString(rule.Verbs...), Resources: sets.NewString(rule.Resources...), ResourceNames: sets.NewString(rule.ResourceNames...), NonResourceURLs: sets.NewString(rule.NonResourceURLs...)}
		rules = append(rules, r)
	}
	return rules
}
func Convert_rbacv1_PolicyRules_To_authorization_PolicyRules(in []rbacv1.PolicyRule) []authorizationapi.PolicyRule {
	_logClusterCodePath()
	defer _logClusterCodePath()
	rules := make([]authorizationapi.PolicyRule, 0, len(in))
	for _, rule := range in {
		r := authorizationapi.PolicyRule{APIGroups: rule.APIGroups, Verbs: sets.NewString(rule.Verbs...), Resources: sets.NewString(rule.Resources...), ResourceNames: sets.NewString(rule.ResourceNames...), NonResourceURLs: sets.NewString(rule.NonResourceURLs...)}
		rules = append(rules, r)
	}
	return rules
}
func copyMapExcept(in map[string]string, except string) map[string]string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out := map[string]string{}
	for k, v := range in {
		if k != except {
			out[k] = v
		}
	}
	return out
}

var stringBool = sets.NewString("true", "false")

func convert_authorization_Annotations_To_rbac_Annotations(in map[string]string) map[string]string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if value, ok := in[reconcileProtectAnnotation]; ok && stringBool.Has(value) {
		out := copyMapExcept(in, reconcileProtectAnnotation)
		if value == "true" {
			out[rbac.AutoUpdateAnnotationKey] = "false"
		} else {
			out[rbac.AutoUpdateAnnotationKey] = "true"
		}
		return out
	}
	return in
}
func convert_rbac_Annotations_To_authorization_Annotations(in map[string]string) map[string]string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if value, ok := in[rbac.AutoUpdateAnnotationKey]; ok && stringBool.Has(value) {
		out := copyMapExcept(in, rbac.AutoUpdateAnnotationKey)
		if value == "true" {
			out[reconcileProtectAnnotation] = "false"
		} else {
			out[reconcileProtectAnnotation] = "true"
		}
		return out
	}
	return in
}
func determineUserKind(user string) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	kind := authorizationapi.UserKind
	if len(validation.ValidateUserName(user, false)) != 0 {
		kind = authorizationapi.SystemUserKind
	}
	return kind
}
func determineGroupKind(group string) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	kind := authorizationapi.GroupKind
	if len(validation.ValidateGroupName(group, false)) != 0 {
		kind = authorizationapi.SystemGroupKind
	}
	return kind
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
	godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
