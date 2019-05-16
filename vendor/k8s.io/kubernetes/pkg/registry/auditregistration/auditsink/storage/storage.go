package storage

import (
	goformat "fmt"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apiserver/pkg/registry/generic"
	genericregistry "k8s.io/apiserver/pkg/registry/generic/registry"
	"k8s.io/kubernetes/pkg/apis/auditregistration"
	auditstrategy "k8s.io/kubernetes/pkg/registry/auditregistration/auditsink"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

type REST struct{ *genericregistry.Store }

func NewREST(optsGetter generic.RESTOptionsGetter) *REST {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	store := &genericregistry.Store{NewFunc: func() runtime.Object {
		return &auditregistration.AuditSink{}
	}, NewListFunc: func() runtime.Object {
		return &auditregistration.AuditSinkList{}
	}, ObjectNameFunc: func(obj runtime.Object) (string, error) {
		return obj.(*auditregistration.AuditSink).Name, nil
	}, DefaultQualifiedResource: auditregistration.Resource("auditsinks"), CreateStrategy: auditstrategy.Strategy, UpdateStrategy: auditstrategy.Strategy, DeleteStrategy: auditstrategy.Strategy}
	options := &generic.StoreOptions{RESTOptions: optsGetter}
	if err := store.CompleteWithOptions(options); err != nil {
		panic(err)
	}
	return &REST{store}
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
