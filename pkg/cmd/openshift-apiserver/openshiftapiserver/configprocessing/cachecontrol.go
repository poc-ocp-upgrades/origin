package configprocessing

import (
	goformat "fmt"
	"net/http"
	goos "os"
	godefaultruntime "runtime"
	"strings"
	gotime "time"
)

var cacheExcludedPathPrefixes = []string{"/swagger-2.0.0.json", "/swagger-2.0.0.pb-v1", "/swagger-2.0.0.pb-v1.gz", "/swagger.json", "/swaggerapi", "/openapi/"}

func WithCacheControl(handler http.Handler, value string) http.Handler {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if _, ok := w.Header()["Cache-Control"]; ok {
			handler.ServeHTTP(w, req)
			return
		}
		for _, prefix := range cacheExcludedPathPrefixes {
			if strings.HasPrefix(req.URL.Path, prefix) {
				handler.ServeHTTP(w, req)
				return
			}
		}
		w.Header().Set("Cache-Control", value)
		handler.ServeHTTP(w, req)
	})
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
