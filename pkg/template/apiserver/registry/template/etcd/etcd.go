package etcd

import (
	goformat "fmt"
	templategroup "github.com/openshift/api/template"
	printersinternal "github.com/openshift/origin/pkg/printers/internalversion"
	templateapi "github.com/openshift/origin/pkg/template/apis/template"
	"github.com/openshift/origin/pkg/template/apiserver/registry/template"
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
		return &templateapi.Template{}
	}, NewListFunc: func() runtime.Object {
		return &templateapi.TemplateList{}
	}, DefaultQualifiedResource: templategroup.Resource("templates"), TableConvertor: printerstorage.TableConvertor{TablePrinter: printers.NewTablePrinter().With(printersinternal.AddHandlers)}, CreateStrategy: template.Strategy, UpdateStrategy: template.Strategy, DeleteStrategy: template.Strategy, ReturnDeletedObject: true}
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
