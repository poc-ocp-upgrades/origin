package etcd

import (
	goformat "fmt"
	"github.com/openshift/api/network"
	networkapi "github.com/openshift/origin/pkg/network/apis/network"
	"github.com/openshift/origin/pkg/network/apiserver/registry/egressnetworkpolicy"
	printersinternal "github.com/openshift/origin/pkg/printers/internalversion"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apiserver/pkg/registry/generic"
	"k8s.io/apiserver/pkg/registry/generic/registry"
	"k8s.io/apiserver/pkg/registry/rest"
	"k8s.io/kubernetes/pkg/printers"
	printerstorage "k8s.io/kubernetes/pkg/printers/storage"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

type REST struct{ *registry.Store }

var _ rest.StandardStorage = &REST{}

func NewREST(optsGetter generic.RESTOptionsGetter) (*REST, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	store := &registry.Store{NewFunc: func() runtime.Object {
		return &networkapi.EgressNetworkPolicy{}
	}, NewListFunc: func() runtime.Object {
		return &networkapi.EgressNetworkPolicyList{}
	}, DefaultQualifiedResource: network.Resource("egressnetworkpolicies"), TableConvertor: printerstorage.TableConvertor{TablePrinter: printers.NewTablePrinter().With(printersinternal.AddHandlers)}, CreateStrategy: egressnetworkpolicy.Strategy, UpdateStrategy: egressnetworkpolicy.Strategy, DeleteStrategy: egressnetworkpolicy.Strategy}
	options := &generic.StoreOptions{RESTOptions: optsGetter}
	if err := store.CompleteWithOptions(options); err != nil {
		return nil, err
	}
	return &REST{store}, nil
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
