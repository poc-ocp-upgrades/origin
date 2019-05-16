package clusterrole

import (
	"context"
	goformat "fmt"
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	genericapirequest "k8s.io/apiserver/pkg/endpoints/request"
	"k8s.io/apiserver/pkg/registry/rest"
	"k8s.io/kubernetes/pkg/apis/rbac"
	rbacv1helpers "k8s.io/kubernetes/pkg/apis/rbac/v1"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

type Registry interface {
	GetClusterRole(ctx context.Context, name string, options *metav1.GetOptions) (*rbacv1.ClusterRole, error)
}
type storage struct{ rest.Getter }

func NewRegistry(s rest.StandardStorage) Registry {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &storage{s}
}
func (s *storage) GetClusterRole(ctx context.Context, name string, options *metav1.GetOptions) (*rbacv1.ClusterRole, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	obj, err := s.Get(ctx, name, options)
	if err != nil {
		return nil, err
	}
	ret := &rbacv1.ClusterRole{}
	if err := rbacv1helpers.Convert_rbac_ClusterRole_To_v1_ClusterRole(obj.(*rbac.ClusterRole), ret, nil); err != nil {
		return nil, err
	}
	return ret, nil
}

type AuthorizerAdapter struct{ Registry Registry }

func (a AuthorizerAdapter) GetClusterRole(name string) (*rbacv1.ClusterRole, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return a.Registry.GetClusterRole(genericapirequest.NewContext(), name, &metav1.GetOptions{})
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
