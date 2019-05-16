package placeholderchallenger

import (
	"fmt"
	goformat "fmt"
	oauthhandlers "github.com/openshift/origin/pkg/oauthserver/oauth/handlers"
	"net/http"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

type placeholderChallenger struct{ tokenRequestURL string }

func New(url string) oauthhandlers.AuthenticationChallenger {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return placeholderChallenger{url}
}
func (c placeholderChallenger) AuthenticationChallenge(req *http.Request) (http.Header, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	headers := http.Header{}
	headers.Add("Warning", fmt.Sprintf(`%s %s "You must obtain an API token by visiting %s"`, oauthhandlers.WarningHeaderMiscCode, oauthhandlers.WarningHeaderOpenShiftSource, c.tokenRequestURL))
	headers.Add("Link", fmt.Sprintf(`<%s>; rel="related"`, c.tokenRequestURL))
	return headers, nil
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
