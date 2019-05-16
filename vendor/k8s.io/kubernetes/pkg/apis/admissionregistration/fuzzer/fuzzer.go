package fuzzer

import (
	goformat "fmt"
	fuzz "github.com/google/gofuzz"
	runtimeserializer "k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/kubernetes/pkg/apis/admissionregistration"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

var Funcs = func(codecs runtimeserializer.CodecFactory) []interface{} {
	return []interface{}{func(obj *admissionregistration.Webhook, c fuzz.Continue) {
		c.FuzzNoCustom(obj)
		p := admissionregistration.FailurePolicyType("Fail")
		obj.FailurePolicy = &p
		s := admissionregistration.SideEffectClassUnknown
		obj.SideEffects = &s
	}}
}

func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
