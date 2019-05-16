package setdefault

import (
	"fmt"
	goformat "fmt"
	"io"
	storagev1 "k8s.io/api/storage/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apiserver/pkg/admission"
	genericadmissioninitializer "k8s.io/apiserver/pkg/admission/initializer"
	"k8s.io/client-go/informers"
	storagev1listers "k8s.io/client-go/listers/storage/v1"
	"k8s.io/klog"
	api "k8s.io/kubernetes/pkg/apis/core"
	"k8s.io/kubernetes/pkg/apis/core/helper"
	storageutil "k8s.io/kubernetes/pkg/apis/storage/util"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

const (
	PluginName = "DefaultStorageClass"
)

func Register(plugins *admission.Plugins) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	plugins.Register(PluginName, func(config io.Reader) (admission.Interface, error) {
		plugin := newPlugin()
		return plugin, nil
	})
}

type claimDefaulterPlugin struct {
	*admission.Handler
	lister storagev1listers.StorageClassLister
}

var _ admission.Interface = &claimDefaulterPlugin{}
var _ admission.MutationInterface = &claimDefaulterPlugin{}
var _ = genericadmissioninitializer.WantsExternalKubeInformerFactory(&claimDefaulterPlugin{})

func newPlugin() *claimDefaulterPlugin {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &claimDefaulterPlugin{Handler: admission.NewHandler(admission.Create)}
}
func (a *claimDefaulterPlugin) SetExternalKubeInformerFactory(f informers.SharedInformerFactory) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	informer := f.Storage().V1().StorageClasses()
	a.lister = informer.Lister()
	a.SetReadyFunc(informer.Informer().HasSynced)
}
func (a *claimDefaulterPlugin) ValidateInitialization() error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if a.lister == nil {
		return fmt.Errorf("missing lister")
	}
	return nil
}
func (a *claimDefaulterPlugin) Admit(attr admission.Attributes) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if attr.GetResource().GroupResource() != api.Resource("persistentvolumeclaims") {
		return nil
	}
	if len(attr.GetSubresource()) != 0 {
		return nil
	}
	pvc, ok := attr.GetObject().(*api.PersistentVolumeClaim)
	if !ok {
		return nil
	}
	if helper.PersistentVolumeClaimHasClass(pvc) {
		return nil
	}
	klog.V(4).Infof("no storage class for claim %s (generate: %s)", pvc.Name, pvc.GenerateName)
	def, err := getDefaultClass(a.lister)
	if err != nil {
		return admission.NewForbidden(attr, err)
	}
	if def == nil {
		return nil
	}
	klog.V(4).Infof("defaulting storage class for claim %s (generate: %s) to %s", pvc.Name, pvc.GenerateName, def.Name)
	pvc.Spec.StorageClassName = &def.Name
	return nil
}
func getDefaultClass(lister storagev1listers.StorageClassLister) (*storagev1.StorageClass, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	list, err := lister.List(labels.Everything())
	if err != nil {
		return nil, err
	}
	defaultClasses := []*storagev1.StorageClass{}
	for _, class := range list {
		if storageutil.IsDefaultAnnotation(class.ObjectMeta) {
			defaultClasses = append(defaultClasses, class)
			klog.V(4).Infof("getDefaultClass added: %s", class.Name)
		}
	}
	if len(defaultClasses) == 0 {
		return nil, nil
	}
	if len(defaultClasses) > 1 {
		klog.V(4).Infof("getDefaultClass %d defaults found", len(defaultClasses))
		return nil, errors.NewInternalError(fmt.Errorf("%d default StorageClasses were found", len(defaultClasses)))
	}
	return defaultClasses[0], nil
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
