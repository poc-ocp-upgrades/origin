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
	"k8s.io/apiserver/pkg/storage"
	"k8s.io/kubernetes/pkg/printers"
	printerstorage "k8s.io/kubernetes/pkg/printers/storage"
	"github.com/openshift/api/oauth"
	oauthapi "github.com/openshift/origin/pkg/oauth/apis/oauth"
	"github.com/openshift/origin/pkg/oauth/apiserver/registry/oauthclient"
	"github.com/openshift/origin/pkg/oauth/apiserver/registry/oauthclientauthorization"
	printersinternal "github.com/openshift/origin/pkg/printers/internalversion"
)

type REST struct{ *registry.Store }

var _ rest.StandardStorage = &REST{}

func NewREST(optsGetter generic.RESTOptionsGetter, clientGetter oauthclient.Getter) (*REST, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	strategy := oauthclientauthorization.NewStrategy(clientGetter)
	store := &registry.Store{NewFunc: func() runtime.Object {
		return &oauthapi.OAuthClientAuthorization{}
	}, NewListFunc: func() runtime.Object {
		return &oauthapi.OAuthClientAuthorizationList{}
	}, DefaultQualifiedResource: oauth.Resource("oauthclientauthorizations"), TableConvertor: printerstorage.TableConvertor{TablePrinter: printers.NewTablePrinter().With(printersinternal.AddHandlers)}, CreateStrategy: strategy, UpdateStrategy: strategy, DeleteStrategy: strategy}
	options := &generic.StoreOptions{RESTOptions: optsGetter, AttrFunc: storage.AttrFunc(storage.DefaultNamespaceScopedAttr).WithFieldMutation(oauthapi.OAuthClientAuthorizationFieldSelector)}
	if err := store.CompleteWithOptions(options); err != nil {
		return nil, err
	}
	return &REST{store}, nil
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := runtime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", runtime.FuncForPC(pc).Name()))
	http.Post("/"+"logcode", "application/json", bytes.NewBuffer(jsonLog))
}
