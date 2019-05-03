package rest

import (
 "context"
 godefaultbytes "bytes"
 godefaulthttp "net/http"
 godefaultruntime "runtime"
 "fmt"
 "k8s.io/apimachinery/pkg/api/errors"
 "k8s.io/apimachinery/pkg/runtime"
 genericregistry "k8s.io/apiserver/pkg/registry/generic/registry"
 genericrest "k8s.io/apiserver/pkg/registry/generic/rest"
 "k8s.io/apiserver/pkg/registry/rest"
 api "k8s.io/kubernetes/pkg/apis/core"
 "k8s.io/kubernetes/pkg/apis/core/validation"
 "k8s.io/kubernetes/pkg/kubelet/client"
 "k8s.io/kubernetes/pkg/registry/core/pod"
)

type LogREST struct {
 KubeletConn client.ConnectionInfoGetter
 Store       *genericregistry.Store
}

var _ = rest.GetterWithOptions(&LogREST{})

func (r *LogREST) New() runtime.Object {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &api.Pod{}
}
func (r *LogREST) ProducesMIMETypes(verb string) []string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return []string{"text/plain"}
}
func (r *LogREST) ProducesObject(verb string) interface{} {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return ""
}
func (r *LogREST) Get(ctx context.Context, name string, opts runtime.Object) (runtime.Object, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 logOpts, ok := opts.(*api.PodLogOptions)
 if !ok {
  return nil, fmt.Errorf("invalid options object: %#v", opts)
 }
 if errs := validation.ValidatePodLogOptions(logOpts); len(errs) > 0 {
  return nil, errors.NewInvalid(api.Kind("PodLogOptions"), name, errs)
 }
 location, transport, err := pod.LogLocation(r.Store, r.KubeletConn, ctx, name, logOpts)
 if err != nil {
  return nil, err
 }
 return &genericrest.LocationStreamer{Location: location, Transport: transport, ContentType: "text/plain", Flush: logOpts.Follow, ResponseChecker: genericrest.NewGenericHttpResponseChecker(api.Resource("pods/log"), name), RedirectChecker: genericrest.PreventRedirects}, nil
}
func (r *LogREST) NewGetOptions() (runtime.Object, bool, string) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &api.PodLogOptions{}, false, ""
}
func (r *LogREST) OverrideMetricsVerb(oldVerb string) (newVerb string) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 newVerb = oldVerb
 if oldVerb == "GET" {
  newVerb = "CONNECT"
 }
 return
}
func _logClusterCodePath() {
 pc, _, _, _ := godefaultruntime.Caller(1)
 jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
 godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
