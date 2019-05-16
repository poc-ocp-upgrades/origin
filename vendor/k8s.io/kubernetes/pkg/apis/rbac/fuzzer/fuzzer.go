package fuzzer

import (
	goformat "fmt"
	fuzz "github.com/google/gofuzz"
	runtimeserializer "k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/kubernetes/pkg/apis/rbac"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

var Funcs = func(codecs runtimeserializer.CodecFactory) []interface{} {
	return []interface{}{func(r *rbac.RoleRef, c fuzz.Continue) {
		c.FuzzNoCustom(r)
		if len(r.APIGroup) == 0 {
			r.APIGroup = rbac.GroupName
		}
	}, func(r *rbac.Subject, c fuzz.Continue) {
		switch c.Int31n(3) {
		case 0:
			r.Kind = rbac.ServiceAccountKind
			r.APIGroup = ""
			c.FuzzNoCustom(&r.Name)
			c.FuzzNoCustom(&r.Namespace)
		case 1:
			r.Kind = rbac.UserKind
			r.APIGroup = rbac.GroupName
			c.FuzzNoCustom(&r.Name)
			for r.Name == "*" {
				c.FuzzNoCustom(&r.Name)
			}
		case 2:
			r.Kind = rbac.GroupKind
			r.APIGroup = rbac.GroupName
			c.FuzzNoCustom(&r.Name)
		}
	}}
}

func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
