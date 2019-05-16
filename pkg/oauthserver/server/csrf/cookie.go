package csrf

import (
	goformat "fmt"
	"github.com/openshift/origin/pkg/oauthserver/server/crypto"
	"net/http"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

type cookieCsrf struct {
	name   string
	path   string
	domain string
	secure bool
}

func NewCookieCSRF(name, path, domain string, secure bool) CSRF {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &cookieCsrf{name: name, path: path, domain: domain, secure: secure}
}
func (c *cookieCsrf) Generate(w http.ResponseWriter, req *http.Request) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	cookie, err := req.Cookie(c.name)
	if err == nil && len(cookie.Value) > 0 {
		return cookie.Value
	}
	cookie = &http.Cookie{Name: c.name, Value: crypto.Random256BitsString(), Path: c.path, Domain: c.domain, Secure: c.secure, HttpOnly: true}
	http.SetCookie(w, cookie)
	return cookie.Value
}
func (c *cookieCsrf) Check(req *http.Request, value string) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if len(value) == 0 {
		return false
	}
	cookie, err := req.Cookie(c.name)
	if err != nil {
		return false
	}
	return crypto.IsEqualConstantTime(cookie.Value, value)
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
