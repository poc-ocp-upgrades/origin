package strict

import (
	goformat "fmt"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/klog"
	"k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm/scheme"
	"k8s.io/kubernetes/cmd/kubeadm/app/componentconfigs"
	goos "os"
	godefaultruntime "runtime"
	"sigs.k8s.io/yaml"
	gotime "time"
)

func VerifyUnmarshalStrict(bytes []byte, gvk schema.GroupVersionKind) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var (
		iface interface{}
		err   error
	)
	iface, err = scheme.Scheme.New(gvk)
	if err != nil {
		iface, err = componentconfigs.Scheme.New(gvk)
		if err != nil {
			err := errors.Errorf("unknown configuration %#v for scheme definitions in %q and %q", gvk, scheme.Scheme.Name(), componentconfigs.Scheme.Name())
			klog.Warning(err.Error())
			return err
		}
	}
	if err := yaml.UnmarshalStrict(bytes, iface); err != nil {
		err := errors.Wrapf(err, "error unmarshaling configuration %#v", gvk)
		klog.Warning(err.Error())
		return err
	}
	return nil
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
