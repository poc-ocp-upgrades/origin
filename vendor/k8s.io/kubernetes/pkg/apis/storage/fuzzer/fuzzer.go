package fuzzer

import (
	goformat "fmt"
	fuzz "github.com/google/gofuzz"
	runtimeserializer "k8s.io/apimachinery/pkg/runtime/serializer"
	api "k8s.io/kubernetes/pkg/apis/core"
	"k8s.io/kubernetes/pkg/apis/storage"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

var Funcs = func(codecs runtimeserializer.CodecFactory) []interface{} {
	return []interface{}{func(obj *storage.StorageClass, c fuzz.Continue) {
		c.FuzzNoCustom(obj)
		reclamationPolicies := []api.PersistentVolumeReclaimPolicy{api.PersistentVolumeReclaimDelete, api.PersistentVolumeReclaimRetain}
		obj.ReclaimPolicy = &reclamationPolicies[c.Rand.Intn(len(reclamationPolicies))]
		bindingModes := []storage.VolumeBindingMode{storage.VolumeBindingImmediate, storage.VolumeBindingWaitForFirstConsumer}
		obj.VolumeBindingMode = &bindingModes[c.Rand.Intn(len(bindingModes))]
	}}
}

func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
