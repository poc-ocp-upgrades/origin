package util

import (
	rbacv1 "k8s.io/api/rbac/v1"
	"k8s.io/kubernetes/pkg/apis/rbac"
	rbacv1helpers "k8s.io/kubernetes/pkg/apis/rbac/v1"
	authorizationapi "github.com/openshift/origin/pkg/authorization/apis/authorization"
	"github.com/openshift/origin/pkg/authorization/apis/authorization/rbacconversion"
)

func ClusterRoleToRBAC(obj *authorizationapi.ClusterRole) (*rbacv1.ClusterRole, error) {
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
	objCopy := obj.DeepCopy()
	convertedObjInternal := &rbac.ClusterRole{}
	if err := rbacconversion.Convert_authorization_ClusterRole_To_rbac_ClusterRole(objCopy, convertedObjInternal, nil); err != nil {
		return nil, err
	}
	convertedObj := &rbacv1.ClusterRole{}
	if err := rbacv1helpers.Convert_rbac_ClusterRole_To_v1_ClusterRole(convertedObjInternal, convertedObj, nil); err != nil {
		return nil, err
	}
	return convertedObj, nil
}
func ClusterRoleBindingToRBAC(obj *authorizationapi.ClusterRoleBinding) (*rbacv1.ClusterRoleBinding, error) {
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
	objCopy := obj.DeepCopy()
	convertedObjInternal := &rbac.ClusterRoleBinding{}
	if err := rbacconversion.Convert_authorization_ClusterRoleBinding_To_rbac_ClusterRoleBinding(objCopy, convertedObjInternal, nil); err != nil {
		return nil, err
	}
	convertedObj := &rbacv1.ClusterRoleBinding{}
	if err := rbacv1helpers.Convert_rbac_ClusterRoleBinding_To_v1_ClusterRoleBinding(convertedObjInternal, convertedObj, nil); err != nil {
		return nil, err
	}
	return convertedObj, nil
}
func RoleToRBAC(obj *authorizationapi.Role) (*rbacv1.Role, error) {
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
	objCopy := obj.DeepCopy()
	convertedObjInternal := &rbac.Role{}
	if err := rbacconversion.Convert_authorization_Role_To_rbac_Role(objCopy, convertedObjInternal, nil); err != nil {
		return nil, err
	}
	convertedObj := &rbacv1.Role{}
	if err := rbacv1helpers.Convert_rbac_Role_To_v1_Role(convertedObjInternal, convertedObj, nil); err != nil {
		return nil, err
	}
	return convertedObj, nil
}
func RoleBindingToRBAC(obj *authorizationapi.RoleBinding) (*rbacv1.RoleBinding, error) {
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
	objCopy := obj.DeepCopy()
	convertedObjInternal := &rbac.RoleBinding{}
	if err := rbacconversion.Convert_authorization_RoleBinding_To_rbac_RoleBinding(objCopy, convertedObjInternal, nil); err != nil {
		return nil, err
	}
	convertedObj := &rbacv1.RoleBinding{}
	if err := rbacv1helpers.Convert_rbac_RoleBinding_To_v1_RoleBinding(convertedObjInternal, convertedObj, nil); err != nil {
		return nil, err
	}
	return convertedObj, nil
}
func ClusterRoleFromRBAC(obj *rbacv1.ClusterRole) (*authorizationapi.ClusterRole, error) {
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
	objCopy := obj.DeepCopy()
	convertedObjInternal := &rbac.ClusterRole{}
	if err := rbacv1helpers.Convert_v1_ClusterRole_To_rbac_ClusterRole(objCopy, convertedObjInternal, nil); err != nil {
		return nil, err
	}
	convertedObj := &authorizationapi.ClusterRole{}
	if err := rbacconversion.Convert_rbac_ClusterRole_To_authorization_ClusterRole(convertedObjInternal, convertedObj, nil); err != nil {
		return nil, err
	}
	return convertedObj, nil
}
func ClusterRoleBindingFromRBAC(obj *rbacv1.ClusterRoleBinding) (*authorizationapi.ClusterRoleBinding, error) {
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
	objCopy := obj.DeepCopy()
	convertedObjInternal := &rbac.ClusterRoleBinding{}
	if err := rbacv1helpers.Convert_v1_ClusterRoleBinding_To_rbac_ClusterRoleBinding(objCopy, convertedObjInternal, nil); err != nil {
		return nil, err
	}
	convertedObj := &authorizationapi.ClusterRoleBinding{}
	if err := rbacconversion.Convert_rbac_ClusterRoleBinding_To_authorization_ClusterRoleBinding(convertedObjInternal, convertedObj, nil); err != nil {
		return nil, err
	}
	return convertedObj, nil
}
func RoleFromRBAC(obj *rbacv1.Role) (*authorizationapi.Role, error) {
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
	objCopy := obj.DeepCopy()
	convertedObjInternal := &rbac.Role{}
	if err := rbacv1helpers.Convert_v1_Role_To_rbac_Role(objCopy, convertedObjInternal, nil); err != nil {
		return nil, err
	}
	convertedObj := &authorizationapi.Role{}
	if err := rbacconversion.Convert_rbac_Role_To_authorization_Role(convertedObjInternal, convertedObj, nil); err != nil {
		return nil, err
	}
	return convertedObj, nil
}
func RoleBindingFromRBAC(obj *rbacv1.RoleBinding) (*authorizationapi.RoleBinding, error) {
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
	objCopy := obj.DeepCopy()
	convertedObjInternal := &rbac.RoleBinding{}
	if err := rbacv1helpers.Convert_v1_RoleBinding_To_rbac_RoleBinding(objCopy, convertedObjInternal, nil); err != nil {
		return nil, err
	}
	convertedObj := &authorizationapi.RoleBinding{}
	if err := rbacconversion.Convert_rbac_RoleBinding_To_authorization_RoleBinding(convertedObjInternal, convertedObj, nil); err != nil {
		return nil, err
	}
	return convertedObj, nil
}
