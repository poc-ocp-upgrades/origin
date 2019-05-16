package latest

import (
	goformat "fmt"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/runtime/serializer/json"
	"k8s.io/apimachinery/pkg/runtime/serializer/versioning"
	schedulerapi "k8s.io/kubernetes/pkg/scheduler/api"
	_ "k8s.io/kubernetes/pkg/scheduler/api/v1"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

const Version = "v1"
const OldestVersion = "v1"

var Versions = []string{"v1"}
var Codec runtime.Codec

func init() {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	jsonSerializer := json.NewSerializer(json.DefaultMetaFactory, schedulerapi.Scheme, schedulerapi.Scheme, true)
	Codec = versioning.NewDefaultingCodecForScheme(schedulerapi.Scheme, jsonSerializer, jsonSerializer, schema.GroupVersion{Version: Version}, runtime.InternalGroupVersioner)
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
