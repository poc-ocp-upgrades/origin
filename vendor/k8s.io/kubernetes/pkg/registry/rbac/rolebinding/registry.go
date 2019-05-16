package rolebinding

import (
	"context"
	goformat "fmt"
	rbacv1 "k8s.io/api/rbac/v1"
	metainternalversion "k8s.io/apimachinery/pkg/apis/meta/internalversion"
	genericapirequest "k8s.io/apiserver/pkg/endpoints/request"
	"k8s.io/apiserver/pkg/registry/rest"
	"k8s.io/kubernetes/pkg/apis/rbac"
	rbacv1helpers "k8s.io/kubernetes/pkg/apis/rbac/v1"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

type Registry interface {
	ListRoleBindings(ctx context.Context, options *metainternalversion.ListOptions) (*rbacv1.RoleBindingList, error)
}
type storage struct{ rest.Lister }

func NewRegistry(s rest.StandardStorage) Registry {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &storage{s}
}
func (s *storage) ListRoleBindings(ctx context.Context, options *metainternalversion.ListOptions) (*rbacv1.RoleBindingList, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
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
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
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
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
