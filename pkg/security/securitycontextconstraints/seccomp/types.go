package seccomp

import (
	goformat "fmt"
	"k8s.io/apimachinery/pkg/util/validation/field"
	api "k8s.io/kubernetes/pkg/apis/core"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

type SeccompStrategy interface {
	Generate(pod *api.Pod) (string, error)
	ValidatePod(pod *api.Pod) field.ErrorList
	ValidateContainer(pod *api.Pod, container *api.Container) field.ErrorList
}

func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
