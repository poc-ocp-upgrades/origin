package persistentvolume

import (
	"context"
	"fmt"
	goformat "fmt"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"k8s.io/apiserver/pkg/registry/generic"
	"k8s.io/apiserver/pkg/storage"
	"k8s.io/apiserver/pkg/storage/names"
	"k8s.io/kubernetes/pkg/api/legacyscheme"
	pvutil "k8s.io/kubernetes/pkg/api/persistentvolume"
	api "k8s.io/kubernetes/pkg/apis/core"
	"k8s.io/kubernetes/pkg/apis/core/validation"
	volumevalidation "k8s.io/kubernetes/pkg/volume/validation"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

type persistentvolumeStrategy struct {
	runtime.ObjectTyper
	names.NameGenerator
}

var Strategy = persistentvolumeStrategy{legacyscheme.Scheme, names.SimpleNameGenerator}

func (persistentvolumeStrategy) NamespaceScoped() bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return false
}
func (persistentvolumeStrategy) PrepareForCreate(ctx context.Context, obj runtime.Object) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	pv := obj.(*api.PersistentVolume)
	pv.Status = api.PersistentVolumeStatus{}
	pvutil.DropDisabledFields(&pv.Spec, nil)
}
func (persistentvolumeStrategy) Validate(ctx context.Context, obj runtime.Object) field.ErrorList {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	persistentvolume := obj.(*api.PersistentVolume)
	errorList := validation.ValidatePersistentVolume(persistentvolume)
	return append(errorList, volumevalidation.ValidatePersistentVolume(persistentvolume)...)
}
func (persistentvolumeStrategy) Canonicalize(obj runtime.Object) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
}
func (persistentvolumeStrategy) AllowCreateOnUpdate() bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return false
}
func (persistentvolumeStrategy) PrepareForUpdate(ctx context.Context, obj, old runtime.Object) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	newPv := obj.(*api.PersistentVolume)
	oldPv := old.(*api.PersistentVolume)
	newPv.Status = oldPv.Status
	pvutil.DropDisabledFields(&newPv.Spec, &oldPv.Spec)
}
func (persistentvolumeStrategy) ValidateUpdate(ctx context.Context, obj, old runtime.Object) field.ErrorList {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	newPv := obj.(*api.PersistentVolume)
	errorList := validation.ValidatePersistentVolume(newPv)
	errorList = append(errorList, volumevalidation.ValidatePersistentVolume(newPv)...)
	return append(errorList, validation.ValidatePersistentVolumeUpdate(newPv, old.(*api.PersistentVolume))...)
}
func (persistentvolumeStrategy) AllowUnconditionalUpdate() bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return true
}

type persistentvolumeStatusStrategy struct{ persistentvolumeStrategy }

var StatusStrategy = persistentvolumeStatusStrategy{Strategy}

func (persistentvolumeStatusStrategy) PrepareForUpdate(ctx context.Context, obj, old runtime.Object) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	newPv := obj.(*api.PersistentVolume)
	oldPv := old.(*api.PersistentVolume)
	newPv.Spec = oldPv.Spec
}
func (persistentvolumeStatusStrategy) ValidateUpdate(ctx context.Context, obj, old runtime.Object) field.ErrorList {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return validation.ValidatePersistentVolumeStatusUpdate(obj.(*api.PersistentVolume), old.(*api.PersistentVolume))
}
func GetAttrs(obj runtime.Object) (labels.Set, fields.Set, bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	persistentvolumeObj, ok := obj.(*api.PersistentVolume)
	if !ok {
		return nil, nil, false, fmt.Errorf("not a persistentvolume")
	}
	return labels.Set(persistentvolumeObj.Labels), PersistentVolumeToSelectableFields(persistentvolumeObj), persistentvolumeObj.Initializers != nil, nil
}
func MatchPersistentVolumes(label labels.Selector, field fields.Selector) storage.SelectionPredicate {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return storage.SelectionPredicate{Label: label, Field: field, GetAttrs: GetAttrs}
}
func PersistentVolumeToSelectableFields(persistentvolume *api.PersistentVolume) fields.Set {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	objectMetaFieldsSet := generic.ObjectMetaFieldsSet(&persistentvolume.ObjectMeta, false)
	specificFieldsSet := fields.Set{"name": persistentvolume.Name}
	return generic.MergeFieldsSets(objectMetaFieldsSet, specificFieldsSet)
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
