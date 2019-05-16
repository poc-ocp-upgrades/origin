package storage

import (
	goformat "fmt"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apiserver/pkg/authentication/authenticator"
	"k8s.io/apiserver/pkg/registry/generic"
	genericregistry "k8s.io/apiserver/pkg/registry/generic/registry"
	"k8s.io/apiserver/pkg/registry/rest"
	api "k8s.io/kubernetes/pkg/apis/core"
	"k8s.io/kubernetes/pkg/printers"
	printersinternal "k8s.io/kubernetes/pkg/printers/internalversion"
	printerstorage "k8s.io/kubernetes/pkg/printers/storage"
	"k8s.io/kubernetes/pkg/registry/core/serviceaccount"
	token "k8s.io/kubernetes/pkg/serviceaccount"
	goos "os"
	godefaultruntime "runtime"
	"time"
	gotime "time"
)

type REST struct {
	*genericregistry.Store
	Token *TokenREST
}

func NewREST(optsGetter generic.RESTOptionsGetter, issuer token.TokenGenerator, auds authenticator.Audiences, max time.Duration, podStorage, secretStorage *genericregistry.Store) *REST {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	store := &genericregistry.Store{NewFunc: func() runtime.Object {
		return &api.ServiceAccount{}
	}, NewListFunc: func() runtime.Object {
		return &api.ServiceAccountList{}
	}, DefaultQualifiedResource: api.Resource("serviceaccounts"), CreateStrategy: serviceaccount.Strategy, UpdateStrategy: serviceaccount.Strategy, DeleteStrategy: serviceaccount.Strategy, ReturnDeletedObject: true, TableConvertor: printerstorage.TableConvertor{TablePrinter: printers.NewTablePrinter().With(printersinternal.AddHandlers)}}
	options := &generic.StoreOptions{RESTOptions: optsGetter}
	if err := store.CompleteWithOptions(options); err != nil {
		panic(err)
	}
	var trest *TokenREST
	if issuer != nil && podStorage != nil && secretStorage != nil {
		trest = &TokenREST{svcaccts: store, pods: podStorage, secrets: secretStorage, issuer: issuer, auds: auds, maxExpirationSeconds: int64(max.Seconds())}
	}
	return &REST{Store: store, Token: trest}
}

var _ rest.ShortNamesProvider = &REST{}

func (r *REST) ShortNames() []string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return []string{"sa"}
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
