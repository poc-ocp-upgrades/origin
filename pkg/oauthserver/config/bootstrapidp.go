package config

import (
	goformat "fmt"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

type BootstrapIdentityProvider struct{ v1.TypeMeta }

func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
