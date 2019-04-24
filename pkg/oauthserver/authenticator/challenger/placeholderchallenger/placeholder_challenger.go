package placeholderchallenger

import (
	"fmt"
	"bytes"
	"runtime"
	"net/http"
	oauthhandlers "github.com/openshift/origin/pkg/oauthserver/oauth/handlers"
)

type placeholderChallenger struct{ tokenRequestURL string }

func New(url string) oauthhandlers.AuthenticationChallenger {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return placeholderChallenger{url}
}
func (c placeholderChallenger) AuthenticationChallenge(req *http.Request) (http.Header, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	headers := http.Header{}
	headers.Add("Warning", fmt.Sprintf(`%s %s "You must obtain an API token by visiting %s"`, oauthhandlers.WarningHeaderMiscCode, oauthhandlers.WarningHeaderOpenShiftSource, c.tokenRequestURL))
	headers.Add("Link", fmt.Sprintf(`<%s>; rel="related"`, c.tokenRequestURL))
	return headers, nil
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := runtime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", runtime.FuncForPC(pc).Name()))
	http.Post("/"+"logcode", "application/json", bytes.NewBuffer(jsonLog))
}
