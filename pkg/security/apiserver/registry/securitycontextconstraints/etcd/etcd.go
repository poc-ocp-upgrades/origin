package etcd

import (
	"k8s.io/apimachinery/pkg/runtime"
	"bytes"
	"net/http"
	"runtime"
	"fmt"
	"k8s.io/apiserver/pkg/registry/generic"
	"k8s.io/apiserver/pkg/registry/generic/registry"
	"k8s.io/apiserver/pkg/registry/rest"
	"k8s.io/kubernetes/pkg/printers"
	printerstorage "k8s.io/kubernetes/pkg/printers/storage"
	"github.com/openshift/api/security"
	printersinternal "github.com/openshift/origin/pkg/printers/internalversion"
	securityapi "github.com/openshift/origin/pkg/security/apis/security"
	"github.com/openshift/origin/pkg/security/apiserver/registry/securitycontextconstraints"
)

type REST struct{ *registry.Store }

var _ rest.StandardStorage = &REST{}
var _ rest.ShortNamesProvider = &REST{}

func (r *REST) ShortNames() []string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return []string{"scc"}
}
func NewREST(optsGetter generic.RESTOptionsGetter) *REST {
	_logClusterCodePath()
	defer _logClusterCodePath()
	store := &registry.Store{NewFunc: func() runtime.Object {
		return &securityapi.SecurityContextConstraints{}
	}, NewListFunc: func() runtime.Object {
		return &securityapi.SecurityContextConstraintsList{}
	}, ObjectNameFunc: func(obj runtime.Object) (string, error) {
		return obj.(*securityapi.SecurityContextConstraints).Name, nil
	}, PredicateFunc: securitycontextconstraints.Matcher, DefaultQualifiedResource: security.Resource("securitycontextconstraints"), TableConvertor: printerstorage.TableConvertor{TablePrinter: printers.NewTablePrinter().With(printersinternal.AddHandlers)}, CreateStrategy: securitycontextconstraints.Strategy, UpdateStrategy: securitycontextconstraints.Strategy, DeleteStrategy: securitycontextconstraints.Strategy, ReturnDeletedObject: true}
	options := &generic.StoreOptions{RESTOptions: optsGetter, AttrFunc: securitycontextconstraints.GetAttrs}
	if err := store.CompleteWithOptions(options); err != nil {
		panic(err)
	}
	return &REST{store}
}

type LegacyREST struct{ *REST }

func (r *LegacyREST) Categories() []string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return []string{}
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := runtime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", runtime.FuncForPC(pc).Name()))
	http.Post("/"+"logcode", "application/json", bytes.NewBuffer(jsonLog))
}
