package headers

import (
	"net/http"
	"bytes"
	"runtime"
	"fmt"
)

var standardHeaders = map[string]string{"Cache-Control": "no-cache, no-store, max-age=0, must-revalidate", "Pragma": "no-cache", "Expires": "0", "Referrer-Policy": "strict-origin-when-cross-origin", "X-Frame-Options": "DENY", "X-Content-Type-Options": "nosniff", "X-DNS-Prefetch-Control": "off", "X-XSS-Protection": "1; mode=block"}

func WithStandardHeaders(handler http.Handler) http.Handler {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h := w.Header()
		for k, v := range standardHeaders {
			h.Set(k, v)
		}
		handler.ServeHTTP(w, r)
	})
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := runtime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", runtime.FuncForPC(pc).Name()))
	http.Post("/"+"logcode", "application/json", bytes.NewBuffer(jsonLog))
}
