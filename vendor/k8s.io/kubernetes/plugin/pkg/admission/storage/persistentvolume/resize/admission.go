package resize

import (
	"fmt"
	goformat "fmt"
	"io"
	"k8s.io/apiserver/pkg/admission"
	genericadmissioninitializer "k8s.io/apiserver/pkg/admission/initializer"
	"k8s.io/client-go/informers"
	storagev1listers "k8s.io/client-go/listers/storage/v1"
	api "k8s.io/kubernetes/pkg/apis/core"
	apihelper "k8s.io/kubernetes/pkg/apis/core/helper"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

const (
	PluginName = "PersistentVolumeClaimResize"
)

func Register(plugins *admission.Plugins) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	plugins.Register(PluginName, func(config io.Reader) (admission.Interface, error) {
		plugin := newPlugin()
		return plugin, nil
	})
}

var _ admission.Interface = &persistentVolumeClaimResize{}
var _ admission.ValidationInterface = &persistentVolumeClaimResize{}
var _ = genericadmissioninitializer.WantsExternalKubeInformerFactory(&persistentVolumeClaimResize{})

type persistentVolumeClaimResize struct {
	*admission.Handler
	scLister storagev1listers.StorageClassLister
}

func newPlugin() *persistentVolumeClaimResize {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &persistentVolumeClaimResize{Handler: admission.NewHandler(admission.Update)}
}
func (pvcr *persistentVolumeClaimResize) SetExternalKubeInformerFactory(f informers.SharedInformerFactory) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	scInformer := f.Storage().V1().StorageClasses()
	pvcr.scLister = scInformer.Lister()
	pvcr.SetReadyFunc(scInformer.Informer().HasSynced)
}
func (pvcr *persistentVolumeClaimResize) ValidateInitialization() error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if pvcr.scLister == nil {
		return fmt.Errorf("missing storageclass lister")
	}
	return nil
}
func (pvcr *persistentVolumeClaimResize) Validate(a admission.Attributes) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if a.GetResource().GroupResource() != api.Resource("persistentvolumeclaims") {
		return nil
	}
	if len(a.GetSubresource()) != 0 {
		return nil
	}
	pvc, ok := a.GetObject().(*api.PersistentVolumeClaim)
	if !ok {
		return nil
	}
	oldPvc, ok := a.GetOldObject().(*api.PersistentVolumeClaim)
	if !ok {
		return nil
	}
	oldSize := oldPvc.Spec.Resources.Requests[api.ResourceStorage]
	newSize := pvc.Spec.Resources.Requests[api.ResourceStorage]
	if newSize.Cmp(oldSize) <= 0 {
		return nil
	}
	if oldPvc.Status.Phase != api.ClaimBound {
		return admission.NewForbidden(a, fmt.Errorf("Only bound persistent volume claims can be expanded"))
	}
	if !pvcr.allowResize(pvc, oldPvc) {
		return admission.NewForbidden(a, fmt.Errorf("only dynamically provisioned pvc can be resized and "+"the storageclass that provisions the pvc must support resize"))
	}
	return nil
}
func (pvcr *persistentVolumeClaimResize) allowResize(pvc, oldPvc *api.PersistentVolumeClaim) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	pvcStorageClass := apihelper.GetPersistentVolumeClaimClass(pvc)
	oldPvcStorageClass := apihelper.GetPersistentVolumeClaimClass(oldPvc)
	if pvcStorageClass == "" || oldPvcStorageClass == "" || pvcStorageClass != oldPvcStorageClass {
		return false
	}
	sc, err := pvcr.scLister.Get(pvcStorageClass)
	if err != nil {
		return false
	}
	if sc.AllowVolumeExpansion != nil {
		return *sc.AllowVolumeExpansion
	}
	return false
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
