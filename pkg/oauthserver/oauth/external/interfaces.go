package external

import (
	"github.com/RangelReale/osincli"
	authapi "github.com/openshift/origin/pkg/oauthserver/api"
	"net/http"
)

type Provider interface {
	NewConfig() (*osincli.ClientConfig, error)
	GetTransport() (http.RoundTripper, error)
	AddCustomParameters(*osincli.AuthorizeRequest)
	GetUserIdentity(*osincli.AccessData) (authapi.UserIdentityInfo, bool, error)
}
type State interface {
	Generate(w http.ResponseWriter, req *http.Request) (string, error)
	Check(state string, req *http.Request) (bool, error)
}
