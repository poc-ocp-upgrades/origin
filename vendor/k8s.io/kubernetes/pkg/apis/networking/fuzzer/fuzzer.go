package fuzzer

import (
	goformat "fmt"
	fuzz "github.com/google/gofuzz"
	runtimeserializer "k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/kubernetes/pkg/apis/networking"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

var Funcs = func(codecs runtimeserializer.CodecFactory) []interface{} {
	return []interface{}{func(np *networking.NetworkPolicyPeer, c fuzz.Continue) {
		c.FuzzNoCustom(np)
		if np.IPBlock != nil {
			np.IPBlock = &networking.IPBlock{CIDR: "192.168.1.0/24", Except: []string{"192.168.1.1/24", "192.168.1.2/24"}}
		}
	}, func(np *networking.NetworkPolicy, c fuzz.Continue) {
		c.FuzzNoCustom(np)
		if len(np.Spec.PolicyTypes) == 0 {
			np.Spec.PolicyTypes = []networking.PolicyType{networking.PolicyTypeIngress}
		}
	}}
}

func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
