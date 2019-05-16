package storageobjectinuseprotection

import (
	goformat "fmt"
	"io"
	"k8s.io/apiserver/pkg/admission"
	"k8s.io/apiserver/pkg/util/feature"
	"k8s.io/klog"
	api "k8s.io/kubernetes/pkg/apis/core"
	"k8s.io/kubernetes/pkg/features"
	volumeutil "k8s.io/kubernetes/pkg/volume/util"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

const (
	PluginName = "StorageObjectInUseProtection"
)

func Register(plugins *admission.Plugins) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	plugins.Register(PluginName, func(config io.Reader) (admission.Interface, error) {
		plugin := newPlugin()
		return plugin, nil
	})
}

type storageProtectionPlugin struct{ *admission.Handler }

var _ admission.Interface = &storageProtectionPlugin{}

func newPlugin() *storageProtectionPlugin {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &storageProtectionPlugin{Handler: admission.NewHandler(admission.Create)}
}

var (
	pvResource  = api.Resource("persistentvolumes")
	pvcResource = api.Resource("persistentvolumeclaims")
)

func (c *storageProtectionPlugin) Admit(a admission.Attributes) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if !feature.DefaultFeatureGate.Enabled(features.StorageObjectInUseProtection) {
		return nil
	}
	switch a.GetResource().GroupResource() {
	case pvResource:
		return c.admitPV(a)
	case pvcResource:
		return c.admitPVC(a)
	default:
		return nil
	}
}
func (c *storageProtectionPlugin) admitPV(a admission.Attributes) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if len(a.GetSubresource()) != 0 {
		return nil
	}
	pv, ok := a.GetObject().(*api.PersistentVolume)
	if !ok {
		return nil
	}
	for _, f := range pv.Finalizers {
		if f == volumeutil.PVProtectionFinalizer {
			return nil
		}
	}
	klog.V(4).Infof("adding PV protection finalizer to %s", pv.Name)
	pv.Finalizers = append(pv.Finalizers, volumeutil.PVProtectionFinalizer)
	return nil
}
func (c *storageProtectionPlugin) admitPVC(a admission.Attributes) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if len(a.GetSubresource()) != 0 {
		return nil
	}
	pvc, ok := a.GetObject().(*api.PersistentVolumeClaim)
	if !ok {
		return nil
	}
	for _, f := range pvc.Finalizers {
		if f == volumeutil.PVCProtectionFinalizer {
			return nil
		}
	}
	klog.V(4).Infof("adding PVC protection finalizer to %s/%s", pvc.Namespace, pvc.Name)
	pvc.Finalizers = append(pvc.Finalizers, volumeutil.PVCProtectionFinalizer)
	return nil
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
