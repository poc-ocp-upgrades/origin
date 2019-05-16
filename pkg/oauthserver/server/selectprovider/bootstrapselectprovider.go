package selectprovider

import (
	goformat "fmt"
	"github.com/openshift/origin/pkg/oauthserver/api"
	"github.com/openshift/origin/pkg/oauthserver/authenticator/password/bootstrap"
	"github.com/openshift/origin/pkg/oauthserver/oauth/handlers"
	"net/http"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

func NewBootstrapSelectProvider(delegate handlers.AuthenticationSelectionHandler, getter bootstrap.BootstrapUserDataGetter) handlers.AuthenticationSelectionHandler {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &bootstrapSelectProvider{delegate: delegate, getter: getter}
}

type bootstrapSelectProvider struct {
	delegate handlers.AuthenticationSelectionHandler
	getter   bootstrap.BootstrapUserDataGetter
}

func (b *bootstrapSelectProvider) SelectAuthentication(providers []api.ProviderInfo, w http.ResponseWriter, req *http.Request) (*api.ProviderInfo, bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if len(providers) <= 1 || providers[0].Name != bootstrap.BootstrapUser {
		return b.delegate.SelectAuthentication(providers, w, req)
	}
	_, ok, err := b.getter.Get()
	if err != nil || !ok {
		return b.delegate.SelectAuthentication(providers[1:], w, req)
	}
	return b.delegate.SelectAuthentication(providers, w, req)
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
