package headers

import "net/http"

const (
	authzHeader     = "Authorization"
	copyAuthzHeader = "oauth.openshift.io:" + authzHeader
)

func WithPreserveAuthorizationHeader(handler http.Handler) http.Handler {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if vv, ok := r.Header[authzHeader]; ok {
			r.Header[copyAuthzHeader] = vv
		}
		handler.ServeHTTP(w, r)
	})
}
func WithRestoreAuthorizationHeader(handler http.Handler) http.Handler {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if vv, ok := r.Header[copyAuthzHeader]; ok {
			r.Header[authzHeader] = vv
			delete(r.Header, copyAuthzHeader)
		}
		handler.ServeHTTP(w, r)
	})
}
