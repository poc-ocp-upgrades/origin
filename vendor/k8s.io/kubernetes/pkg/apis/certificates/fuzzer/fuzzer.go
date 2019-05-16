package fuzzer

import (
	goformat "fmt"
	fuzz "github.com/google/gofuzz"
	runtimeserializer "k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/kubernetes/pkg/apis/certificates"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

var Funcs = func(codecs runtimeserializer.CodecFactory) []interface{} {
	return []interface{}{func(obj *certificates.CertificateSigningRequestSpec, c fuzz.Continue) {
		c.FuzzNoCustom(obj)
		obj.Usages = []certificates.KeyUsage{certificates.UsageKeyEncipherment}
	}}
}

func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
