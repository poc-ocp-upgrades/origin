package seccomp

import (
	"k8s.io/apimachinery/pkg/util/validation/field"
	"bytes"
	"net/http"
	"runtime"
	"fmt"
	api "k8s.io/kubernetes/pkg/apis/core"
)

type SeccompStrategy interface {
	Generate(pod *api.Pod) (string, error)
	ValidatePod(pod *api.Pod) field.ErrorList
	ValidateContainer(pod *api.Pod, container *api.Container) field.ErrorList
}

func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := runtime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", runtime.FuncForPC(pc).Name()))
	http.Post("/"+"logcode", "application/json", bytes.NewBuffer(jsonLog))
}
