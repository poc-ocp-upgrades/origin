package rolebinding

import (
 "context"
 rbacv1 "k8s.io/api/rbac/v1"
 metainternalversion "k8s.io/apimachinery/pkg/apis/meta/internalversion"
 genericapirequest "k8s.io/apiserver/pkg/endpoints/request"
 "k8s.io/apiserver/pkg/registry/rest"
 "k8s.io/kubernetes/pkg/apis/rbac"
 rbacv1helpers "k8s.io/kubernetes/pkg/apis/rbac/v1"
)

type Registry interface {
 ListRoleBindings(ctx context.Context, options *metainternalversion.ListOptions) (*rbacv1.RoleBindingList, error)
}
type storage struct{ rest.Lister }

func NewRegistry(s rest.StandardStorage) Registry {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &storage{s}
}
func (s *storage) ListRoleBindings(ctx context.Context, options *metainternalversion.ListOptions) (*rbacv1.RoleBindingList, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := s.List(ctx, options)
 if err != nil {
  return nil, err
 }
 ret := &rbacv1.RoleBindingList{}
 if err := rbacv1helpers.Convert_rbac_RoleBindingList_To_v1_RoleBindingList(obj.(*rbac.RoleBindingList), ret, nil); err != nil {
  return nil, err
 }
 return ret, nil
}

type AuthorizerAdapter struct{ Registry Registry }

func (a AuthorizerAdapter) ListRoleBindings(namespace string) ([]*rbacv1.RoleBinding, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 list, err := a.Registry.ListRoleBindings(genericapirequest.WithNamespace(genericapirequest.NewContext(), namespace), &metainternalversion.ListOptions{})
 if err != nil {
  return nil, err
 }
 ret := []*rbacv1.RoleBinding{}
 for i := range list.Items {
  ret = append(ret, &list.Items[i])
 }
 return ret, nil
}
