package clusterrolebinding

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
 ListClusterRoleBindings(ctx context.Context, options *metainternalversion.ListOptions) (*rbacv1.ClusterRoleBindingList, error)
}
type storage struct{ rest.Lister }

func NewRegistry(s rest.StandardStorage) Registry {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &storage{s}
}
func (s *storage) ListClusterRoleBindings(ctx context.Context, options *metainternalversion.ListOptions) (*rbacv1.ClusterRoleBindingList, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := s.List(ctx, options)
 if err != nil {
  return nil, err
 }
 ret := &rbacv1.ClusterRoleBindingList{}
 if err := rbacv1helpers.Convert_rbac_ClusterRoleBindingList_To_v1_ClusterRoleBindingList(obj.(*rbac.ClusterRoleBindingList), ret, nil); err != nil {
  return nil, err
 }
 return ret, nil
}

type AuthorizerAdapter struct{ Registry Registry }

func (a AuthorizerAdapter) ListClusterRoleBindings() ([]*rbacv1.ClusterRoleBinding, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 list, err := a.Registry.ListClusterRoleBindings(genericapirequest.NewContext(), &metainternalversion.ListOptions{})
 if err != nil {
  return nil, err
 }
 ret := []*rbacv1.ClusterRoleBinding{}
 for i := range list.Items {
  ret = append(ret, &list.Items[i])
 }
 return ret, nil
}
