package fuzzer

import (
	"fmt"
	goformat "fmt"
	"github.com/google/gofuzz"
	runtimeserializer "k8s.io/apimachinery/pkg/runtime/serializer"
	kubectrlmgrconfig "k8s.io/kubernetes/pkg/controller/apis/config"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

func Funcs(codecs runtimeserializer.CodecFactory) []interface{} {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
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
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
