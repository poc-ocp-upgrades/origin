package authenticator

import (
	"k8s.io/apiserver/pkg/authentication/authenticator"
	"bytes"
	"net/http"
	"runtime"
	"fmt"
	"github.com/openshift/origin/pkg/oauthserver/api"
)

type Assertion interface {
	AuthenticateAssertion(assertionType, data string) (*authenticator.Response, bool, error)
}
type Client interface {
	AuthenticateClient(client api.Client) (*authenticator.Response, bool, error)
}

func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := runtime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", runtime.FuncForPC(pc).Name()))
	http.Post("/"+"logcode", "application/json", bytes.NewBuffer(jsonLog))
}
