package role

import (
 "context"
 rbacv1 "k8s.io/api/rbac/v1"
 metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
 genericapirequest "k8s.io/apiserver/pkg/endpoints/request"
 "k8s.io/apiserver/pkg/registry/rest"
 "k8s.io/kubernetes/pkg/apis/rbac"
 rbacv1helpers "k8s.io/kubernetes/pkg/apis/rbac/v1"
)

type Registry interface {
 GetRole(ctx context.Context, name string, options *metav1.GetOptions) (*rbacv1.Role, error)
}
type storage struct{ rest.Getter }

func NewRegistry(s rest.StandardStorage) Registry {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &storage{s}
}
func (s *storage) GetRole(ctx context.Context, name string, options *metav1.GetOptions) (*rbacv1.Role, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, err := s.Get(ctx, name, options)
 if err != nil {
  return nil, err
 }
 ret := &rbacv1.Role{}
 if err := rbacv1helpers.Convert_rbac_Role_To_v1_Role(obj.(*rbac.Role), ret, nil); err != nil {
  return nil, err
 }
 return ret, nil
}

type AuthorizerAdapter struct{ Registry Registry }

func (a AuthorizerAdapter) GetRole(namespace, name string) (*rbacv1.Role, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return a.Registry.GetRole(genericapirequest.WithNamespace(genericapirequest.NewContext(), namespace), name, &metav1.GetOptions{})
}
