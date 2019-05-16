package headers

import (
	goformat "fmt"
	"net/http"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

var standardHeaders = map[string]string{"Cache-Control": "no-cache, no-store, max-age=0, must-revalidate", "Pragma": "no-cache", "Expires": "0", "Referrer-Policy": "strict-origin-when-cross-origin", "X-Frame-Options": "DENY", "X-Content-Type-Options": "nosniff", "X-DNS-Prefetch-Control": "off", "X-XSS-Protection": "1; mode=block"}

func WithStandardHeaders(handler http.Handler) http.Handler {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h := w.Header()
		for k, v := range standardHeaders {
			h.Set(k, v)
		}
		handler.ServeHTTP(w, r)
	})
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
