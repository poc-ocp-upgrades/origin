package v1

import (
	v1 "github.com/openshift/api/authorization/v1"
	authorization "github.com/openshift/origin/pkg/authorization/apis/authorization"
	apicorev1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	core "k8s.io/kubernetes/pkg/apis/core"
	corev1 "k8s.io/kubernetes/pkg/apis/core/v1"
	rbac "k8s.io/kubernetes/pkg/apis/rbac"
	unsafe "unsafe"
)

func init() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	localSchemeBuilder.Register(RegisterConversions)
}
func RegisterConversions(s *runtime.Scheme) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := s.AddGeneratedConversionFunc((*v1.Action)(nil), (*authorization.Action)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_Action_To_authorization_Action(a.(*v1.Action), b.(*authorization.Action), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*authorization.Action)(nil), (*v1.Action)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_authorization_Action_To_v1_Action(a.(*authorization.Action), b.(*v1.Action), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.ClusterRole)(nil), (*authorization.ClusterRole)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_ClusterRole_To_authorization_ClusterRole(a.(*v1.ClusterRole), b.(*authorization.ClusterRole), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*authorization.ClusterRole)(nil), (*v1.ClusterRole)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_authorization_ClusterRole_To_v1_ClusterRole(a.(*authorization.ClusterRole), b.(*v1.ClusterRole), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.ClusterRoleBinding)(nil), (*authorization.ClusterRoleBinding)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_ClusterRoleBinding_To_authorization_ClusterRoleBinding(a.(*v1.ClusterRoleBinding), b.(*authorization.ClusterRoleBinding), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*authorization.ClusterRoleBinding)(nil), (*v1.ClusterRoleBinding)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_authorization_ClusterRoleBinding_To_v1_ClusterRoleBinding(a.(*authorization.ClusterRoleBinding), b.(*v1.ClusterRoleBinding), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.ClusterRoleBindingList)(nil), (*authorization.ClusterRoleBindingList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_ClusterRoleBindingList_To_authorization_ClusterRoleBindingList(a.(*v1.ClusterRoleBindingList), b.(*authorization.ClusterRoleBindingList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*authorization.ClusterRoleBindingList)(nil), (*v1.ClusterRoleBindingList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_authorization_ClusterRoleBindingList_To_v1_ClusterRoleBindingList(a.(*authorization.ClusterRoleBindingList), b.(*v1.ClusterRoleBindingList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.ClusterRoleList)(nil), (*authorization.ClusterRoleList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_ClusterRoleList_To_authorization_ClusterRoleList(a.(*v1.ClusterRoleList), b.(*authorization.ClusterRoleList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*authorization.ClusterRoleList)(nil), (*v1.ClusterRoleList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_authorization_ClusterRoleList_To_v1_ClusterRoleList(a.(*authorization.ClusterRoleList), b.(*v1.ClusterRoleList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.GroupRestriction)(nil), (*authorization.GroupRestriction)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_GroupRestriction_To_authorization_GroupRestriction(a.(*v1.GroupRestriction), b.(*authorization.GroupRestriction), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*authorization.GroupRestriction)(nil), (*v1.GroupRestriction)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_authorization_GroupRestriction_To_v1_GroupRestriction(a.(*authorization.GroupRestriction), b.(*v1.GroupRestriction), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.IsPersonalSubjectAccessReview)(nil), (*authorization.IsPersonalSubjectAccessReview)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_IsPersonalSubjectAccessReview_To_authorization_IsPersonalSubjectAccessReview(a.(*v1.IsPersonalSubjectAccessReview), b.(*authorization.IsPersonalSubjectAccessReview), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*authorization.IsPersonalSubjectAccessReview)(nil), (*v1.IsPersonalSubjectAccessReview)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_authorization_IsPersonalSubjectAccessReview_To_v1_IsPersonalSubjectAccessReview(a.(*authorization.IsPersonalSubjectAccessReview), b.(*v1.IsPersonalSubjectAccessReview), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.LocalResourceAccessReview)(nil), (*authorization.LocalResourceAccessReview)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_LocalResourceAccessReview_To_authorization_LocalResourceAccessReview(a.(*v1.LocalResourceAccessReview), b.(*authorization.LocalResourceAccessReview), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*authorization.LocalResourceAccessReview)(nil), (*v1.LocalResourceAccessReview)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_authorization_LocalResourceAccessReview_To_v1_LocalResourceAccessReview(a.(*authorization.LocalResourceAccessReview), b.(*v1.LocalResourceAccessReview), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.LocalSubjectAccessReview)(nil), (*authorization.LocalSubjectAccessReview)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_LocalSubjectAccessReview_To_authorization_LocalSubjectAccessReview(a.(*v1.LocalSubjectAccessReview), b.(*authorization.LocalSubjectAccessReview), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*authorization.LocalSubjectAccessReview)(nil), (*v1.LocalSubjectAccessReview)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_authorization_LocalSubjectAccessReview_To_v1_LocalSubjectAccessReview(a.(*authorization.LocalSubjectAccessReview), b.(*v1.LocalSubjectAccessReview), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.PolicyRule)(nil), (*authorization.PolicyRule)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_PolicyRule_To_authorization_PolicyRule(a.(*v1.PolicyRule), b.(*authorization.PolicyRule), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*authorization.PolicyRule)(nil), (*v1.PolicyRule)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_authorization_PolicyRule_To_v1_PolicyRule(a.(*authorization.PolicyRule), b.(*v1.PolicyRule), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.ResourceAccessReview)(nil), (*authorization.ResourceAccessReview)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_ResourceAccessReview_To_authorization_ResourceAccessReview(a.(*v1.ResourceAccessReview), b.(*authorization.ResourceAccessReview), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*authorization.ResourceAccessReview)(nil), (*v1.ResourceAccessReview)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_authorization_ResourceAccessReview_To_v1_ResourceAccessReview(a.(*authorization.ResourceAccessReview), b.(*v1.ResourceAccessReview), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.ResourceAccessReviewResponse)(nil), (*authorization.ResourceAccessReviewResponse)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_ResourceAccessReviewResponse_To_authorization_ResourceAccessReviewResponse(a.(*v1.ResourceAccessReviewResponse), b.(*authorization.ResourceAccessReviewResponse), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*authorization.ResourceAccessReviewResponse)(nil), (*v1.ResourceAccessReviewResponse)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_authorization_ResourceAccessReviewResponse_To_v1_ResourceAccessReviewResponse(a.(*authorization.ResourceAccessReviewResponse), b.(*v1.ResourceAccessReviewResponse), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.Role)(nil), (*authorization.Role)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_Role_To_authorization_Role(a.(*v1.Role), b.(*authorization.Role), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*authorization.Role)(nil), (*v1.Role)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_authorization_Role_To_v1_Role(a.(*authorization.Role), b.(*v1.Role), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.RoleBinding)(nil), (*authorization.RoleBinding)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_RoleBinding_To_authorization_RoleBinding(a.(*v1.RoleBinding), b.(*authorization.RoleBinding), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*authorization.RoleBinding)(nil), (*v1.RoleBinding)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_authorization_RoleBinding_To_v1_RoleBinding(a.(*authorization.RoleBinding), b.(*v1.RoleBinding), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.RoleBindingList)(nil), (*authorization.RoleBindingList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_RoleBindingList_To_authorization_RoleBindingList(a.(*v1.RoleBindingList), b.(*authorization.RoleBindingList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*authorization.RoleBindingList)(nil), (*v1.RoleBindingList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_authorization_RoleBindingList_To_v1_RoleBindingList(a.(*authorization.RoleBindingList), b.(*v1.RoleBindingList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.RoleBindingRestriction)(nil), (*authorization.RoleBindingRestriction)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_RoleBindingRestriction_To_authorization_RoleBindingRestriction(a.(*v1.RoleBindingRestriction), b.(*authorization.RoleBindingRestriction), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*authorization.RoleBindingRestriction)(nil), (*v1.RoleBindingRestriction)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_authorization_RoleBindingRestriction_To_v1_RoleBindingRestriction(a.(*authorization.RoleBindingRestriction), b.(*v1.RoleBindingRestriction), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.RoleBindingRestrictionList)(nil), (*authorization.RoleBindingRestrictionList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_RoleBindingRestrictionList_To_authorization_RoleBindingRestrictionList(a.(*v1.RoleBindingRestrictionList), b.(*authorization.RoleBindingRestrictionList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*authorization.RoleBindingRestrictionList)(nil), (*v1.RoleBindingRestrictionList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_authorization_RoleBindingRestrictionList_To_v1_RoleBindingRestrictionList(a.(*authorization.RoleBindingRestrictionList), b.(*v1.RoleBindingRestrictionList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.RoleBindingRestrictionSpec)(nil), (*authorization.RoleBindingRestrictionSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_RoleBindingRestrictionSpec_To_authorization_RoleBindingRestrictionSpec(a.(*v1.RoleBindingRestrictionSpec), b.(*authorization.RoleBindingRestrictionSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*authorization.RoleBindingRestrictionSpec)(nil), (*v1.RoleBindingRestrictionSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_authorization_RoleBindingRestrictionSpec_To_v1_RoleBindingRestrictionSpec(a.(*authorization.RoleBindingRestrictionSpec), b.(*v1.RoleBindingRestrictionSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.RoleList)(nil), (*authorization.RoleList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_RoleList_To_authorization_RoleList(a.(*v1.RoleList), b.(*authorization.RoleList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*authorization.RoleList)(nil), (*v1.RoleList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_authorization_RoleList_To_v1_RoleList(a.(*authorization.RoleList), b.(*v1.RoleList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.SelfSubjectRulesReview)(nil), (*authorization.SelfSubjectRulesReview)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_SelfSubjectRulesReview_To_authorization_SelfSubjectRulesReview(a.(*v1.SelfSubjectRulesReview), b.(*authorization.SelfSubjectRulesReview), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*authorization.SelfSubjectRulesReview)(nil), (*v1.SelfSubjectRulesReview)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_authorization_SelfSubjectRulesReview_To_v1_SelfSubjectRulesReview(a.(*authorization.SelfSubjectRulesReview), b.(*v1.SelfSubjectRulesReview), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.SelfSubjectRulesReviewSpec)(nil), (*authorization.SelfSubjectRulesReviewSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_SelfSubjectRulesReviewSpec_To_authorization_SelfSubjectRulesReviewSpec(a.(*v1.SelfSubjectRulesReviewSpec), b.(*authorization.SelfSubjectRulesReviewSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*authorization.SelfSubjectRulesReviewSpec)(nil), (*v1.SelfSubjectRulesReviewSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_authorization_SelfSubjectRulesReviewSpec_To_v1_SelfSubjectRulesReviewSpec(a.(*authorization.SelfSubjectRulesReviewSpec), b.(*v1.SelfSubjectRulesReviewSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.ServiceAccountReference)(nil), (*authorization.ServiceAccountReference)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_ServiceAccountReference_To_authorization_ServiceAccountReference(a.(*v1.ServiceAccountReference), b.(*authorization.ServiceAccountReference), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*authorization.ServiceAccountReference)(nil), (*v1.ServiceAccountReference)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_authorization_ServiceAccountReference_To_v1_ServiceAccountReference(a.(*authorization.ServiceAccountReference), b.(*v1.ServiceAccountReference), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.ServiceAccountRestriction)(nil), (*authorization.ServiceAccountRestriction)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_ServiceAccountRestriction_To_authorization_ServiceAccountRestriction(a.(*v1.ServiceAccountRestriction), b.(*authorization.ServiceAccountRestriction), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*authorization.ServiceAccountRestriction)(nil), (*v1.ServiceAccountRestriction)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_authorization_ServiceAccountRestriction_To_v1_ServiceAccountRestriction(a.(*authorization.ServiceAccountRestriction), b.(*v1.ServiceAccountRestriction), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.SubjectAccessReview)(nil), (*authorization.SubjectAccessReview)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_SubjectAccessReview_To_authorization_SubjectAccessReview(a.(*v1.SubjectAccessReview), b.(*authorization.SubjectAccessReview), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*authorization.SubjectAccessReview)(nil), (*v1.SubjectAccessReview)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_authorization_SubjectAccessReview_To_v1_SubjectAccessReview(a.(*authorization.SubjectAccessReview), b.(*v1.SubjectAccessReview), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.SubjectAccessReviewResponse)(nil), (*authorization.SubjectAccessReviewResponse)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_SubjectAccessReviewResponse_To_authorization_SubjectAccessReviewResponse(a.(*v1.SubjectAccessReviewResponse), b.(*authorization.SubjectAccessReviewResponse), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*authorization.SubjectAccessReviewResponse)(nil), (*v1.SubjectAccessReviewResponse)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_authorization_SubjectAccessReviewResponse_To_v1_SubjectAccessReviewResponse(a.(*authorization.SubjectAccessReviewResponse), b.(*v1.SubjectAccessReviewResponse), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.SubjectRulesReview)(nil), (*authorization.SubjectRulesReview)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_SubjectRulesReview_To_authorization_SubjectRulesReview(a.(*v1.SubjectRulesReview), b.(*authorization.SubjectRulesReview), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*authorization.SubjectRulesReview)(nil), (*v1.SubjectRulesReview)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_authorization_SubjectRulesReview_To_v1_SubjectRulesReview(a.(*authorization.SubjectRulesReview), b.(*v1.SubjectRulesReview), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.SubjectRulesReviewSpec)(nil), (*authorization.SubjectRulesReviewSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_SubjectRulesReviewSpec_To_authorization_SubjectRulesReviewSpec(a.(*v1.SubjectRulesReviewSpec), b.(*authorization.SubjectRulesReviewSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*authorization.SubjectRulesReviewSpec)(nil), (*v1.SubjectRulesReviewSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_authorization_SubjectRulesReviewSpec_To_v1_SubjectRulesReviewSpec(a.(*authorization.SubjectRulesReviewSpec), b.(*v1.SubjectRulesReviewSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.SubjectRulesReviewStatus)(nil), (*authorization.SubjectRulesReviewStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_SubjectRulesReviewStatus_To_authorization_SubjectRulesReviewStatus(a.(*v1.SubjectRulesReviewStatus), b.(*authorization.SubjectRulesReviewStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*authorization.SubjectRulesReviewStatus)(nil), (*v1.SubjectRulesReviewStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_authorization_SubjectRulesReviewStatus_To_v1_SubjectRulesReviewStatus(a.(*authorization.SubjectRulesReviewStatus), b.(*v1.SubjectRulesReviewStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.UserRestriction)(nil), (*authorization.UserRestriction)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_UserRestriction_To_authorization_UserRestriction(a.(*v1.UserRestriction), b.(*authorization.UserRestriction), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*authorization.UserRestriction)(nil), (*v1.UserRestriction)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_authorization_UserRestriction_To_v1_UserRestriction(a.(*authorization.UserRestriction), b.(*v1.UserRestriction), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*authorization.ClusterRoleBinding)(nil), (*v1.ClusterRoleBinding)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_authorization_ClusterRoleBinding_To_v1_ClusterRoleBinding(a.(*authorization.ClusterRoleBinding), b.(*v1.ClusterRoleBinding), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*authorization.LocalSubjectAccessReview)(nil), (*v1.LocalSubjectAccessReview)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_authorization_LocalSubjectAccessReview_To_v1_LocalSubjectAccessReview(a.(*authorization.LocalSubjectAccessReview), b.(*v1.LocalSubjectAccessReview), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*authorization.PolicyRule)(nil), (*v1.PolicyRule)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_authorization_PolicyRule_To_v1_PolicyRule(a.(*authorization.PolicyRule), b.(*v1.PolicyRule), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*authorization.ResourceAccessReviewResponse)(nil), (*v1.ResourceAccessReviewResponse)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_authorization_ResourceAccessReviewResponse_To_v1_ResourceAccessReviewResponse(a.(*authorization.ResourceAccessReviewResponse), b.(*v1.ResourceAccessReviewResponse), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*authorization.RoleBinding)(nil), (*v1.RoleBinding)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_authorization_RoleBinding_To_v1_RoleBinding(a.(*authorization.RoleBinding), b.(*v1.RoleBinding), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*authorization.SelfSubjectRulesReviewSpec)(nil), (*v1.SelfSubjectRulesReviewSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_authorization_SelfSubjectRulesReviewSpec_To_v1_SelfSubjectRulesReviewSpec(a.(*authorization.SelfSubjectRulesReviewSpec), b.(*v1.SelfSubjectRulesReviewSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*authorization.SubjectAccessReview)(nil), (*v1.SubjectAccessReview)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_authorization_SubjectAccessReview_To_v1_SubjectAccessReview(a.(*authorization.SubjectAccessReview), b.(*v1.SubjectAccessReview), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*v1.ClusterRoleBinding)(nil), (*authorization.ClusterRoleBinding)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_ClusterRoleBinding_To_authorization_ClusterRoleBinding(a.(*v1.ClusterRoleBinding), b.(*authorization.ClusterRoleBinding), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*v1.LocalSubjectAccessReview)(nil), (*authorization.LocalSubjectAccessReview)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_LocalSubjectAccessReview_To_authorization_LocalSubjectAccessReview(a.(*v1.LocalSubjectAccessReview), b.(*authorization.LocalSubjectAccessReview), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*v1.PolicyRule)(nil), (*authorization.PolicyRule)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_PolicyRule_To_authorization_PolicyRule(a.(*v1.PolicyRule), b.(*authorization.PolicyRule), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*v1.ResourceAccessReviewResponse)(nil), (*authorization.ResourceAccessReviewResponse)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_ResourceAccessReviewResponse_To_authorization_ResourceAccessReviewResponse(a.(*v1.ResourceAccessReviewResponse), b.(*authorization.ResourceAccessReviewResponse), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*v1.RoleBinding)(nil), (*authorization.RoleBinding)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_RoleBinding_To_authorization_RoleBinding(a.(*v1.RoleBinding), b.(*authorization.RoleBinding), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*v1.SelfSubjectRulesReviewSpec)(nil), (*authorization.SelfSubjectRulesReviewSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_SelfSubjectRulesReviewSpec_To_authorization_SelfSubjectRulesReviewSpec(a.(*v1.SelfSubjectRulesReviewSpec), b.(*authorization.SelfSubjectRulesReviewSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*v1.SubjectAccessReview)(nil), (*authorization.SubjectAccessReview)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_SubjectAccessReview_To_authorization_SubjectAccessReview(a.(*v1.SubjectAccessReview), b.(*authorization.SubjectAccessReview), scope)
	}); err != nil {
		return err
	}
	return nil
}
func autoConvert_v1_Action_To_authorization_Action(in *v1.Action, out *authorization.Action, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Namespace = in.Namespace
	out.Verb = in.Verb
	out.Group = in.Group
	out.Version = in.Version
	out.Resource = in.Resource
	out.ResourceName = in.ResourceName
	out.Path = in.Path
	out.IsNonResourceURL = in.IsNonResourceURL
	if err := runtime.Convert_runtime_RawExtension_To_runtime_Object(&in.Content, &out.Content, s); err != nil {
		return err
	}
	return nil
}
func Convert_v1_Action_To_authorization_Action(in *v1.Action, out *authorization.Action, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_Action_To_authorization_Action(in, out, s)
}
func autoConvert_authorization_Action_To_v1_Action(in *authorization.Action, out *v1.Action, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Namespace = in.Namespace
	out.Verb = in.Verb
	out.Group = in.Group
	out.Version = in.Version
	out.Resource = in.Resource
	out.ResourceName = in.ResourceName
	out.Path = in.Path
	out.IsNonResourceURL = in.IsNonResourceURL
	if err := runtime.Convert_runtime_Object_To_runtime_RawExtension(&in.Content, &out.Content, s); err != nil {
		return err
	}
	return nil
}
func Convert_authorization_Action_To_v1_Action(in *authorization.Action, out *v1.Action, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_authorization_Action_To_v1_Action(in, out, s)
}
func autoConvert_v1_ClusterRole_To_authorization_ClusterRole(in *v1.ClusterRole, out *authorization.ClusterRole, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ObjectMeta = in.ObjectMeta
	if in.Rules != nil {
		in, out := &in.Rules, &out.Rules
		*out = make([]authorization.PolicyRule, len(*in))
		for i := range *in {
			if err := Convert_v1_PolicyRule_To_authorization_PolicyRule(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Rules = nil
	}
	out.AggregationRule = (*rbac.AggregationRule)(unsafe.Pointer(in.AggregationRule))
	return nil
}
func Convert_v1_ClusterRole_To_authorization_ClusterRole(in *v1.ClusterRole, out *authorization.ClusterRole, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_ClusterRole_To_authorization_ClusterRole(in, out, s)
}
func autoConvert_authorization_ClusterRole_To_v1_ClusterRole(in *authorization.ClusterRole, out *v1.ClusterRole, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ObjectMeta = in.ObjectMeta
	if in.Rules != nil {
		in, out := &in.Rules, &out.Rules
		*out = make([]v1.PolicyRule, len(*in))
		for i := range *in {
			if err := Convert_authorization_PolicyRule_To_v1_PolicyRule(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Rules = nil
	}
	out.AggregationRule = (*rbacv1.AggregationRule)(unsafe.Pointer(in.AggregationRule))
	return nil
}
func Convert_authorization_ClusterRole_To_v1_ClusterRole(in *authorization.ClusterRole, out *v1.ClusterRole, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_authorization_ClusterRole_To_v1_ClusterRole(in, out, s)
}
func autoConvert_v1_ClusterRoleBinding_To_authorization_ClusterRoleBinding(in *v1.ClusterRoleBinding, out *authorization.ClusterRoleBinding, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ObjectMeta = in.ObjectMeta
	if in.Subjects != nil {
		in, out := &in.Subjects, &out.Subjects
		*out = make([]core.ObjectReference, len(*in))
		for i := range *in {
			if err := corev1.Convert_v1_ObjectReference_To_core_ObjectReference(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Subjects = nil
	}
	if err := corev1.Convert_v1_ObjectReference_To_core_ObjectReference(&in.RoleRef, &out.RoleRef, s); err != nil {
		return err
	}
	return nil
}
func autoConvert_authorization_ClusterRoleBinding_To_v1_ClusterRoleBinding(in *authorization.ClusterRoleBinding, out *v1.ClusterRoleBinding, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ObjectMeta = in.ObjectMeta
	if in.Subjects != nil {
		in, out := &in.Subjects, &out.Subjects
		*out = make([]apicorev1.ObjectReference, len(*in))
		for i := range *in {
			if err := corev1.Convert_core_ObjectReference_To_v1_ObjectReference(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Subjects = nil
	}
	if err := corev1.Convert_core_ObjectReference_To_v1_ObjectReference(&in.RoleRef, &out.RoleRef, s); err != nil {
		return err
	}
	return nil
}
func autoConvert_v1_ClusterRoleBindingList_To_authorization_ClusterRoleBindingList(in *v1.ClusterRoleBindingList, out *authorization.ClusterRoleBindingList, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]authorization.ClusterRoleBinding, len(*in))
		for i := range *in {
			if err := Convert_v1_ClusterRoleBinding_To_authorization_ClusterRoleBinding(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}
func Convert_v1_ClusterRoleBindingList_To_authorization_ClusterRoleBindingList(in *v1.ClusterRoleBindingList, out *authorization.ClusterRoleBindingList, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_ClusterRoleBindingList_To_authorization_ClusterRoleBindingList(in, out, s)
}
func autoConvert_authorization_ClusterRoleBindingList_To_v1_ClusterRoleBindingList(in *authorization.ClusterRoleBindingList, out *v1.ClusterRoleBindingList, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]v1.ClusterRoleBinding, len(*in))
		for i := range *in {
			if err := Convert_authorization_ClusterRoleBinding_To_v1_ClusterRoleBinding(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}
func Convert_authorization_ClusterRoleBindingList_To_v1_ClusterRoleBindingList(in *authorization.ClusterRoleBindingList, out *v1.ClusterRoleBindingList, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_authorization_ClusterRoleBindingList_To_v1_ClusterRoleBindingList(in, out, s)
}
func autoConvert_v1_ClusterRoleList_To_authorization_ClusterRoleList(in *v1.ClusterRoleList, out *authorization.ClusterRoleList, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]authorization.ClusterRole, len(*in))
		for i := range *in {
			if err := Convert_v1_ClusterRole_To_authorization_ClusterRole(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}
func Convert_v1_ClusterRoleList_To_authorization_ClusterRoleList(in *v1.ClusterRoleList, out *authorization.ClusterRoleList, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_ClusterRoleList_To_authorization_ClusterRoleList(in, out, s)
}
func autoConvert_authorization_ClusterRoleList_To_v1_ClusterRoleList(in *authorization.ClusterRoleList, out *v1.ClusterRoleList, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]v1.ClusterRole, len(*in))
		for i := range *in {
			if err := Convert_authorization_ClusterRole_To_v1_ClusterRole(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}
func Convert_authorization_ClusterRoleList_To_v1_ClusterRoleList(in *authorization.ClusterRoleList, out *v1.ClusterRoleList, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_authorization_ClusterRoleList_To_v1_ClusterRoleList(in, out, s)
}
func autoConvert_v1_GroupRestriction_To_authorization_GroupRestriction(in *v1.GroupRestriction, out *authorization.GroupRestriction, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Groups = *(*[]string)(unsafe.Pointer(&in.Groups))
	out.Selectors = *(*[]metav1.LabelSelector)(unsafe.Pointer(&in.Selectors))
	return nil
}
func Convert_v1_GroupRestriction_To_authorization_GroupRestriction(in *v1.GroupRestriction, out *authorization.GroupRestriction, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_GroupRestriction_To_authorization_GroupRestriction(in, out, s)
}
func autoConvert_authorization_GroupRestriction_To_v1_GroupRestriction(in *authorization.GroupRestriction, out *v1.GroupRestriction, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Groups = *(*[]string)(unsafe.Pointer(&in.Groups))
	out.Selectors = *(*[]metav1.LabelSelector)(unsafe.Pointer(&in.Selectors))
	return nil
}
func Convert_authorization_GroupRestriction_To_v1_GroupRestriction(in *authorization.GroupRestriction, out *v1.GroupRestriction, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_authorization_GroupRestriction_To_v1_GroupRestriction(in, out, s)
}
func autoConvert_v1_IsPersonalSubjectAccessReview_To_authorization_IsPersonalSubjectAccessReview(in *v1.IsPersonalSubjectAccessReview, out *authorization.IsPersonalSubjectAccessReview, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return nil
}
func Convert_v1_IsPersonalSubjectAccessReview_To_authorization_IsPersonalSubjectAccessReview(in *v1.IsPersonalSubjectAccessReview, out *authorization.IsPersonalSubjectAccessReview, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_IsPersonalSubjectAccessReview_To_authorization_IsPersonalSubjectAccessReview(in, out, s)
}
func autoConvert_authorization_IsPersonalSubjectAccessReview_To_v1_IsPersonalSubjectAccessReview(in *authorization.IsPersonalSubjectAccessReview, out *v1.IsPersonalSubjectAccessReview, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return nil
}
func Convert_authorization_IsPersonalSubjectAccessReview_To_v1_IsPersonalSubjectAccessReview(in *authorization.IsPersonalSubjectAccessReview, out *v1.IsPersonalSubjectAccessReview, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_authorization_IsPersonalSubjectAccessReview_To_v1_IsPersonalSubjectAccessReview(in, out, s)
}
func autoConvert_v1_LocalResourceAccessReview_To_authorization_LocalResourceAccessReview(in *v1.LocalResourceAccessReview, out *authorization.LocalResourceAccessReview, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := Convert_v1_Action_To_authorization_Action(&in.Action, &out.Action, s); err != nil {
		return err
	}
	return nil
}
func Convert_v1_LocalResourceAccessReview_To_authorization_LocalResourceAccessReview(in *v1.LocalResourceAccessReview, out *authorization.LocalResourceAccessReview, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_LocalResourceAccessReview_To_authorization_LocalResourceAccessReview(in, out, s)
}
func autoConvert_authorization_LocalResourceAccessReview_To_v1_LocalResourceAccessReview(in *authorization.LocalResourceAccessReview, out *v1.LocalResourceAccessReview, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := Convert_authorization_Action_To_v1_Action(&in.Action, &out.Action, s); err != nil {
		return err
	}
	return nil
}
func Convert_authorization_LocalResourceAccessReview_To_v1_LocalResourceAccessReview(in *authorization.LocalResourceAccessReview, out *v1.LocalResourceAccessReview, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_authorization_LocalResourceAccessReview_To_v1_LocalResourceAccessReview(in, out, s)
}
func autoConvert_v1_LocalSubjectAccessReview_To_authorization_LocalSubjectAccessReview(in *v1.LocalSubjectAccessReview, out *authorization.LocalSubjectAccessReview, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := Convert_v1_Action_To_authorization_Action(&in.Action, &out.Action, s); err != nil {
		return err
	}
	out.User = in.User
	return nil
}
func autoConvert_authorization_LocalSubjectAccessReview_To_v1_LocalSubjectAccessReview(in *authorization.LocalSubjectAccessReview, out *v1.LocalSubjectAccessReview, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := Convert_authorization_Action_To_v1_Action(&in.Action, &out.Action, s); err != nil {
		return err
	}
	out.User = in.User
	return nil
}
func autoConvert_v1_PolicyRule_To_authorization_PolicyRule(in *v1.PolicyRule, out *authorization.PolicyRule, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := runtime.Convert_runtime_RawExtension_To_runtime_Object(&in.AttributeRestrictions, &out.AttributeRestrictions, s); err != nil {
		return err
	}
	out.APIGroups = *(*[]string)(unsafe.Pointer(&in.APIGroups))
	return nil
}
func autoConvert_authorization_PolicyRule_To_v1_PolicyRule(in *authorization.PolicyRule, out *v1.PolicyRule, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := runtime.Convert_runtime_Object_To_runtime_RawExtension(&in.AttributeRestrictions, &out.AttributeRestrictions, s); err != nil {
		return err
	}
	out.APIGroups = *(*[]string)(unsafe.Pointer(&in.APIGroups))
	return nil
}
func autoConvert_v1_ResourceAccessReview_To_authorization_ResourceAccessReview(in *v1.ResourceAccessReview, out *authorization.ResourceAccessReview, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := Convert_v1_Action_To_authorization_Action(&in.Action, &out.Action, s); err != nil {
		return err
	}
	return nil
}
func Convert_v1_ResourceAccessReview_To_authorization_ResourceAccessReview(in *v1.ResourceAccessReview, out *authorization.ResourceAccessReview, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_ResourceAccessReview_To_authorization_ResourceAccessReview(in, out, s)
}
func autoConvert_authorization_ResourceAccessReview_To_v1_ResourceAccessReview(in *authorization.ResourceAccessReview, out *v1.ResourceAccessReview, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := Convert_authorization_Action_To_v1_Action(&in.Action, &out.Action, s); err != nil {
		return err
	}
	return nil
}
func Convert_authorization_ResourceAccessReview_To_v1_ResourceAccessReview(in *authorization.ResourceAccessReview, out *v1.ResourceAccessReview, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_authorization_ResourceAccessReview_To_v1_ResourceAccessReview(in, out, s)
}
func autoConvert_v1_ResourceAccessReviewResponse_To_authorization_ResourceAccessReviewResponse(in *v1.ResourceAccessReviewResponse, out *authorization.ResourceAccessReviewResponse, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Namespace = in.Namespace
	out.EvaluationError = in.EvaluationError
	return nil
}
func autoConvert_authorization_ResourceAccessReviewResponse_To_v1_ResourceAccessReviewResponse(in *authorization.ResourceAccessReviewResponse, out *v1.ResourceAccessReviewResponse, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Namespace = in.Namespace
	out.EvaluationError = in.EvaluationError
	return nil
}
func autoConvert_v1_Role_To_authorization_Role(in *v1.Role, out *authorization.Role, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ObjectMeta = in.ObjectMeta
	if in.Rules != nil {
		in, out := &in.Rules, &out.Rules
		*out = make([]authorization.PolicyRule, len(*in))
		for i := range *in {
			if err := Convert_v1_PolicyRule_To_authorization_PolicyRule(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Rules = nil
	}
	return nil
}
func Convert_v1_Role_To_authorization_Role(in *v1.Role, out *authorization.Role, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_Role_To_authorization_Role(in, out, s)
}
func autoConvert_authorization_Role_To_v1_Role(in *authorization.Role, out *v1.Role, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ObjectMeta = in.ObjectMeta
	if in.Rules != nil {
		in, out := &in.Rules, &out.Rules
		*out = make([]v1.PolicyRule, len(*in))
		for i := range *in {
			if err := Convert_authorization_PolicyRule_To_v1_PolicyRule(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Rules = nil
	}
	return nil
}
func Convert_authorization_Role_To_v1_Role(in *authorization.Role, out *v1.Role, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_authorization_Role_To_v1_Role(in, out, s)
}
func autoConvert_v1_RoleBinding_To_authorization_RoleBinding(in *v1.RoleBinding, out *authorization.RoleBinding, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ObjectMeta = in.ObjectMeta
	if in.Subjects != nil {
		in, out := &in.Subjects, &out.Subjects
		*out = make([]core.ObjectReference, len(*in))
		for i := range *in {
			if err := corev1.Convert_v1_ObjectReference_To_core_ObjectReference(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Subjects = nil
	}
	if err := corev1.Convert_v1_ObjectReference_To_core_ObjectReference(&in.RoleRef, &out.RoleRef, s); err != nil {
		return err
	}
	return nil
}
func autoConvert_authorization_RoleBinding_To_v1_RoleBinding(in *authorization.RoleBinding, out *v1.RoleBinding, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ObjectMeta = in.ObjectMeta
	if in.Subjects != nil {
		in, out := &in.Subjects, &out.Subjects
		*out = make([]apicorev1.ObjectReference, len(*in))
		for i := range *in {
			if err := corev1.Convert_core_ObjectReference_To_v1_ObjectReference(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Subjects = nil
	}
	if err := corev1.Convert_core_ObjectReference_To_v1_ObjectReference(&in.RoleRef, &out.RoleRef, s); err != nil {
		return err
	}
	return nil
}
func autoConvert_v1_RoleBindingList_To_authorization_RoleBindingList(in *v1.RoleBindingList, out *authorization.RoleBindingList, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]authorization.RoleBinding, len(*in))
		for i := range *in {
			if err := Convert_v1_RoleBinding_To_authorization_RoleBinding(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}
func Convert_v1_RoleBindingList_To_authorization_RoleBindingList(in *v1.RoleBindingList, out *authorization.RoleBindingList, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_RoleBindingList_To_authorization_RoleBindingList(in, out, s)
}
func autoConvert_authorization_RoleBindingList_To_v1_RoleBindingList(in *authorization.RoleBindingList, out *v1.RoleBindingList, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]v1.RoleBinding, len(*in))
		for i := range *in {
			if err := Convert_authorization_RoleBinding_To_v1_RoleBinding(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}
func Convert_authorization_RoleBindingList_To_v1_RoleBindingList(in *authorization.RoleBindingList, out *v1.RoleBindingList, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_authorization_RoleBindingList_To_v1_RoleBindingList(in, out, s)
}
func autoConvert_v1_RoleBindingRestriction_To_authorization_RoleBindingRestriction(in *v1.RoleBindingRestriction, out *authorization.RoleBindingRestriction, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1_RoleBindingRestrictionSpec_To_authorization_RoleBindingRestrictionSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	return nil
}
func Convert_v1_RoleBindingRestriction_To_authorization_RoleBindingRestriction(in *v1.RoleBindingRestriction, out *authorization.RoleBindingRestriction, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_RoleBindingRestriction_To_authorization_RoleBindingRestriction(in, out, s)
}
func autoConvert_authorization_RoleBindingRestriction_To_v1_RoleBindingRestriction(in *authorization.RoleBindingRestriction, out *v1.RoleBindingRestriction, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_authorization_RoleBindingRestrictionSpec_To_v1_RoleBindingRestrictionSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	return nil
}
func Convert_authorization_RoleBindingRestriction_To_v1_RoleBindingRestriction(in *authorization.RoleBindingRestriction, out *v1.RoleBindingRestriction, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_authorization_RoleBindingRestriction_To_v1_RoleBindingRestriction(in, out, s)
}
func autoConvert_v1_RoleBindingRestrictionList_To_authorization_RoleBindingRestrictionList(in *v1.RoleBindingRestrictionList, out *authorization.RoleBindingRestrictionList, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ListMeta = in.ListMeta
	out.Items = *(*[]authorization.RoleBindingRestriction)(unsafe.Pointer(&in.Items))
	return nil
}
func Convert_v1_RoleBindingRestrictionList_To_authorization_RoleBindingRestrictionList(in *v1.RoleBindingRestrictionList, out *authorization.RoleBindingRestrictionList, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_RoleBindingRestrictionList_To_authorization_RoleBindingRestrictionList(in, out, s)
}
func autoConvert_authorization_RoleBindingRestrictionList_To_v1_RoleBindingRestrictionList(in *authorization.RoleBindingRestrictionList, out *v1.RoleBindingRestrictionList, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ListMeta = in.ListMeta
	out.Items = *(*[]v1.RoleBindingRestriction)(unsafe.Pointer(&in.Items))
	return nil
}
func Convert_authorization_RoleBindingRestrictionList_To_v1_RoleBindingRestrictionList(in *authorization.RoleBindingRestrictionList, out *v1.RoleBindingRestrictionList, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_authorization_RoleBindingRestrictionList_To_v1_RoleBindingRestrictionList(in, out, s)
}
func autoConvert_v1_RoleBindingRestrictionSpec_To_authorization_RoleBindingRestrictionSpec(in *v1.RoleBindingRestrictionSpec, out *authorization.RoleBindingRestrictionSpec, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.UserRestriction = (*authorization.UserRestriction)(unsafe.Pointer(in.UserRestriction))
	out.GroupRestriction = (*authorization.GroupRestriction)(unsafe.Pointer(in.GroupRestriction))
	out.ServiceAccountRestriction = (*authorization.ServiceAccountRestriction)(unsafe.Pointer(in.ServiceAccountRestriction))
	return nil
}
func Convert_v1_RoleBindingRestrictionSpec_To_authorization_RoleBindingRestrictionSpec(in *v1.RoleBindingRestrictionSpec, out *authorization.RoleBindingRestrictionSpec, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_RoleBindingRestrictionSpec_To_authorization_RoleBindingRestrictionSpec(in, out, s)
}
func autoConvert_authorization_RoleBindingRestrictionSpec_To_v1_RoleBindingRestrictionSpec(in *authorization.RoleBindingRestrictionSpec, out *v1.RoleBindingRestrictionSpec, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.UserRestriction = (*v1.UserRestriction)(unsafe.Pointer(in.UserRestriction))
	out.GroupRestriction = (*v1.GroupRestriction)(unsafe.Pointer(in.GroupRestriction))
	out.ServiceAccountRestriction = (*v1.ServiceAccountRestriction)(unsafe.Pointer(in.ServiceAccountRestriction))
	return nil
}
func Convert_authorization_RoleBindingRestrictionSpec_To_v1_RoleBindingRestrictionSpec(in *authorization.RoleBindingRestrictionSpec, out *v1.RoleBindingRestrictionSpec, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_authorization_RoleBindingRestrictionSpec_To_v1_RoleBindingRestrictionSpec(in, out, s)
}
func autoConvert_v1_RoleList_To_authorization_RoleList(in *v1.RoleList, out *authorization.RoleList, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]authorization.Role, len(*in))
		for i := range *in {
			if err := Convert_v1_Role_To_authorization_Role(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}
func Convert_v1_RoleList_To_authorization_RoleList(in *v1.RoleList, out *authorization.RoleList, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_RoleList_To_authorization_RoleList(in, out, s)
}
func autoConvert_authorization_RoleList_To_v1_RoleList(in *authorization.RoleList, out *v1.RoleList, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]v1.Role, len(*in))
		for i := range *in {
			if err := Convert_authorization_Role_To_v1_Role(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}
func Convert_authorization_RoleList_To_v1_RoleList(in *authorization.RoleList, out *v1.RoleList, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_authorization_RoleList_To_v1_RoleList(in, out, s)
}
func autoConvert_v1_SelfSubjectRulesReview_To_authorization_SelfSubjectRulesReview(in *v1.SelfSubjectRulesReview, out *authorization.SelfSubjectRulesReview, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := Convert_v1_SelfSubjectRulesReviewSpec_To_authorization_SelfSubjectRulesReviewSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1_SubjectRulesReviewStatus_To_authorization_SubjectRulesReviewStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}
func Convert_v1_SelfSubjectRulesReview_To_authorization_SelfSubjectRulesReview(in *v1.SelfSubjectRulesReview, out *authorization.SelfSubjectRulesReview, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_SelfSubjectRulesReview_To_authorization_SelfSubjectRulesReview(in, out, s)
}
func autoConvert_authorization_SelfSubjectRulesReview_To_v1_SelfSubjectRulesReview(in *authorization.SelfSubjectRulesReview, out *v1.SelfSubjectRulesReview, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := Convert_authorization_SelfSubjectRulesReviewSpec_To_v1_SelfSubjectRulesReviewSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_authorization_SubjectRulesReviewStatus_To_v1_SubjectRulesReviewStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}
func Convert_authorization_SelfSubjectRulesReview_To_v1_SelfSubjectRulesReview(in *authorization.SelfSubjectRulesReview, out *v1.SelfSubjectRulesReview, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_authorization_SelfSubjectRulesReview_To_v1_SelfSubjectRulesReview(in, out, s)
}
func autoConvert_v1_SelfSubjectRulesReviewSpec_To_authorization_SelfSubjectRulesReviewSpec(in *v1.SelfSubjectRulesReviewSpec, out *authorization.SelfSubjectRulesReviewSpec, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return nil
}
func autoConvert_authorization_SelfSubjectRulesReviewSpec_To_v1_SelfSubjectRulesReviewSpec(in *authorization.SelfSubjectRulesReviewSpec, out *v1.SelfSubjectRulesReviewSpec, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return nil
}
func autoConvert_v1_ServiceAccountReference_To_authorization_ServiceAccountReference(in *v1.ServiceAccountReference, out *authorization.ServiceAccountReference, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Name = in.Name
	out.Namespace = in.Namespace
	return nil
}
func Convert_v1_ServiceAccountReference_To_authorization_ServiceAccountReference(in *v1.ServiceAccountReference, out *authorization.ServiceAccountReference, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_ServiceAccountReference_To_authorization_ServiceAccountReference(in, out, s)
}
func autoConvert_authorization_ServiceAccountReference_To_v1_ServiceAccountReference(in *authorization.ServiceAccountReference, out *v1.ServiceAccountReference, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Name = in.Name
	out.Namespace = in.Namespace
	return nil
}
func Convert_authorization_ServiceAccountReference_To_v1_ServiceAccountReference(in *authorization.ServiceAccountReference, out *v1.ServiceAccountReference, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_authorization_ServiceAccountReference_To_v1_ServiceAccountReference(in, out, s)
}
func autoConvert_v1_ServiceAccountRestriction_To_authorization_ServiceAccountRestriction(in *v1.ServiceAccountRestriction, out *authorization.ServiceAccountRestriction, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ServiceAccounts = *(*[]authorization.ServiceAccountReference)(unsafe.Pointer(&in.ServiceAccounts))
	out.Namespaces = *(*[]string)(unsafe.Pointer(&in.Namespaces))
	return nil
}
func Convert_v1_ServiceAccountRestriction_To_authorization_ServiceAccountRestriction(in *v1.ServiceAccountRestriction, out *authorization.ServiceAccountRestriction, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_ServiceAccountRestriction_To_authorization_ServiceAccountRestriction(in, out, s)
}
func autoConvert_authorization_ServiceAccountRestriction_To_v1_ServiceAccountRestriction(in *authorization.ServiceAccountRestriction, out *v1.ServiceAccountRestriction, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ServiceAccounts = *(*[]v1.ServiceAccountReference)(unsafe.Pointer(&in.ServiceAccounts))
	out.Namespaces = *(*[]string)(unsafe.Pointer(&in.Namespaces))
	return nil
}
func Convert_authorization_ServiceAccountRestriction_To_v1_ServiceAccountRestriction(in *authorization.ServiceAccountRestriction, out *v1.ServiceAccountRestriction, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_authorization_ServiceAccountRestriction_To_v1_ServiceAccountRestriction(in, out, s)
}
func autoConvert_v1_SubjectAccessReview_To_authorization_SubjectAccessReview(in *v1.SubjectAccessReview, out *authorization.SubjectAccessReview, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := Convert_v1_Action_To_authorization_Action(&in.Action, &out.Action, s); err != nil {
		return err
	}
	out.User = in.User
	return nil
}
func autoConvert_authorization_SubjectAccessReview_To_v1_SubjectAccessReview(in *authorization.SubjectAccessReview, out *v1.SubjectAccessReview, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := Convert_authorization_Action_To_v1_Action(&in.Action, &out.Action, s); err != nil {
		return err
	}
	out.User = in.User
	return nil
}
func autoConvert_v1_SubjectAccessReviewResponse_To_authorization_SubjectAccessReviewResponse(in *v1.SubjectAccessReviewResponse, out *authorization.SubjectAccessReviewResponse, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Namespace = in.Namespace
	out.Allowed = in.Allowed
	out.Reason = in.Reason
	out.EvaluationError = in.EvaluationError
	return nil
}
func Convert_v1_SubjectAccessReviewResponse_To_authorization_SubjectAccessReviewResponse(in *v1.SubjectAccessReviewResponse, out *authorization.SubjectAccessReviewResponse, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_SubjectAccessReviewResponse_To_authorization_SubjectAccessReviewResponse(in, out, s)
}
func autoConvert_authorization_SubjectAccessReviewResponse_To_v1_SubjectAccessReviewResponse(in *authorization.SubjectAccessReviewResponse, out *v1.SubjectAccessReviewResponse, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Namespace = in.Namespace
	out.Allowed = in.Allowed
	out.Reason = in.Reason
	out.EvaluationError = in.EvaluationError
	return nil
}
func Convert_authorization_SubjectAccessReviewResponse_To_v1_SubjectAccessReviewResponse(in *authorization.SubjectAccessReviewResponse, out *v1.SubjectAccessReviewResponse, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_authorization_SubjectAccessReviewResponse_To_v1_SubjectAccessReviewResponse(in, out, s)
}
func autoConvert_v1_SubjectRulesReview_To_authorization_SubjectRulesReview(in *v1.SubjectRulesReview, out *authorization.SubjectRulesReview, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := Convert_v1_SubjectRulesReviewSpec_To_authorization_SubjectRulesReviewSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1_SubjectRulesReviewStatus_To_authorization_SubjectRulesReviewStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}
func Convert_v1_SubjectRulesReview_To_authorization_SubjectRulesReview(in *v1.SubjectRulesReview, out *authorization.SubjectRulesReview, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_SubjectRulesReview_To_authorization_SubjectRulesReview(in, out, s)
}
func autoConvert_authorization_SubjectRulesReview_To_v1_SubjectRulesReview(in *authorization.SubjectRulesReview, out *v1.SubjectRulesReview, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := Convert_authorization_SubjectRulesReviewSpec_To_v1_SubjectRulesReviewSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_authorization_SubjectRulesReviewStatus_To_v1_SubjectRulesReviewStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}
func Convert_authorization_SubjectRulesReview_To_v1_SubjectRulesReview(in *authorization.SubjectRulesReview, out *v1.SubjectRulesReview, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_authorization_SubjectRulesReview_To_v1_SubjectRulesReview(in, out, s)
}
func autoConvert_v1_SubjectRulesReviewSpec_To_authorization_SubjectRulesReviewSpec(in *v1.SubjectRulesReviewSpec, out *authorization.SubjectRulesReviewSpec, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.User = in.User
	out.Groups = *(*[]string)(unsafe.Pointer(&in.Groups))
	out.Scopes = *(*[]string)(unsafe.Pointer(&in.Scopes))
	return nil
}
func Convert_v1_SubjectRulesReviewSpec_To_authorization_SubjectRulesReviewSpec(in *v1.SubjectRulesReviewSpec, out *authorization.SubjectRulesReviewSpec, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_SubjectRulesReviewSpec_To_authorization_SubjectRulesReviewSpec(in, out, s)
}
func autoConvert_authorization_SubjectRulesReviewSpec_To_v1_SubjectRulesReviewSpec(in *authorization.SubjectRulesReviewSpec, out *v1.SubjectRulesReviewSpec, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.User = in.User
	out.Groups = *(*[]string)(unsafe.Pointer(&in.Groups))
	out.Scopes = *(*v1.OptionalScopes)(unsafe.Pointer(&in.Scopes))
	return nil
}
func Convert_authorization_SubjectRulesReviewSpec_To_v1_SubjectRulesReviewSpec(in *authorization.SubjectRulesReviewSpec, out *v1.SubjectRulesReviewSpec, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_authorization_SubjectRulesReviewSpec_To_v1_SubjectRulesReviewSpec(in, out, s)
}
func autoConvert_v1_SubjectRulesReviewStatus_To_authorization_SubjectRulesReviewStatus(in *v1.SubjectRulesReviewStatus, out *authorization.SubjectRulesReviewStatus, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in.Rules != nil {
		in, out := &in.Rules, &out.Rules
		*out = make([]authorization.PolicyRule, len(*in))
		for i := range *in {
			if err := Convert_v1_PolicyRule_To_authorization_PolicyRule(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Rules = nil
	}
	out.EvaluationError = in.EvaluationError
	return nil
}
func Convert_v1_SubjectRulesReviewStatus_To_authorization_SubjectRulesReviewStatus(in *v1.SubjectRulesReviewStatus, out *authorization.SubjectRulesReviewStatus, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_SubjectRulesReviewStatus_To_authorization_SubjectRulesReviewStatus(in, out, s)
}
func autoConvert_authorization_SubjectRulesReviewStatus_To_v1_SubjectRulesReviewStatus(in *authorization.SubjectRulesReviewStatus, out *v1.SubjectRulesReviewStatus, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in.Rules != nil {
		in, out := &in.Rules, &out.Rules
		*out = make([]v1.PolicyRule, len(*in))
		for i := range *in {
			if err := Convert_authorization_PolicyRule_To_v1_PolicyRule(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Rules = nil
	}
	out.EvaluationError = in.EvaluationError
	return nil
}
func Convert_authorization_SubjectRulesReviewStatus_To_v1_SubjectRulesReviewStatus(in *authorization.SubjectRulesReviewStatus, out *v1.SubjectRulesReviewStatus, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_authorization_SubjectRulesReviewStatus_To_v1_SubjectRulesReviewStatus(in, out, s)
}
func autoConvert_v1_UserRestriction_To_authorization_UserRestriction(in *v1.UserRestriction, out *authorization.UserRestriction, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Users = *(*[]string)(unsafe.Pointer(&in.Users))
	out.Groups = *(*[]string)(unsafe.Pointer(&in.Groups))
	out.Selectors = *(*[]metav1.LabelSelector)(unsafe.Pointer(&in.Selectors))
	return nil
}
func Convert_v1_UserRestriction_To_authorization_UserRestriction(in *v1.UserRestriction, out *authorization.UserRestriction, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1_UserRestriction_To_authorization_UserRestriction(in, out, s)
}
func autoConvert_authorization_UserRestriction_To_v1_UserRestriction(in *authorization.UserRestriction, out *v1.UserRestriction, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Users = *(*[]string)(unsafe.Pointer(&in.Users))
	out.Groups = *(*[]string)(unsafe.Pointer(&in.Groups))
	out.Selectors = *(*[]metav1.LabelSelector)(unsafe.Pointer(&in.Selectors))
	return nil
}
func Convert_authorization_UserRestriction_To_v1_UserRestriction(in *authorization.UserRestriction, out *v1.UserRestriction, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_authorization_UserRestriction_To_v1_UserRestriction(in, out, s)
}
