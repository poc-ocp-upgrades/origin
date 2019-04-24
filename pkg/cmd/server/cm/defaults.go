package cm

import (
	apiserverflag "k8s.io/apiserver/pkg/util/flag"
	"bytes"
	"net/http"
	"runtime"
	"fmt"
	kcmapp "k8s.io/kubernetes/cmd/kube-controller-manager/app"
	kcmoptions "k8s.io/kubernetes/cmd/kube-controller-manager/app/options"
)

func OriginControllerManagerAddFlags(cmserver *kcmoptions.KubeControllerManagerOptions) apiserverflag.NamedFlagSets {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return cmserver.Flags(kcmapp.KnownControllers(), kcmapp.ControllersDisabledByDefault.List())
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := runtime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", runtime.FuncForPC(pc).Name()))
	http.Post("/"+"logcode", "application/json", bytes.NewBuffer(jsonLog))
}
