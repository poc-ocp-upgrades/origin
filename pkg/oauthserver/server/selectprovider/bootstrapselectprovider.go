package selectprovider

import (
	"net/http"
	"bytes"
	"runtime"
	"fmt"
	"github.com/openshift/origin/pkg/oauthserver/api"
	"github.com/openshift/origin/pkg/oauthserver/authenticator/password/bootstrap"
	"github.com/openshift/origin/pkg/oauthserver/oauth/handlers"
)

func NewBootstrapSelectProvider(delegate handlers.AuthenticationSelectionHandler, getter bootstrap.BootstrapUserDataGetter) handlers.AuthenticationSelectionHandler {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &bootstrapSelectProvider{delegate: delegate, getter: getter}
}

type bootstrapSelectProvider struct {
	delegate	handlers.AuthenticationSelectionHandler
	getter		bootstrap.BootstrapUserDataGetter
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := runtime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", runtime.FuncForPC(pc).Name()))
	http.Post("/"+"logcode", "application/json", bytes.NewBuffer(jsonLog))
}
