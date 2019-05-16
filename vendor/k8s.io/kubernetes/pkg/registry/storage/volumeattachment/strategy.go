package volumeattachment

import (
	"context"
	goformat "fmt"
	storageapiv1beta1 "k8s.io/api/storage/v1beta1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/validation/field"
	genericapirequest "k8s.io/apiserver/pkg/endpoints/request"
	"k8s.io/apiserver/pkg/storage/names"
	"k8s.io/kubernetes/pkg/api/legacyscheme"
	"k8s.io/kubernetes/pkg/apis/storage"
	"k8s.io/kubernetes/pkg/apis/storage/validation"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

type volumeAttachmentStrategy struct {
	runtime.ObjectTyper
	names.NameGenerator
}

var Strategy = volumeAttachmentStrategy{legacyscheme.Scheme, names.SimpleNameGenerator}

func (volumeAttachmentStrategy) NamespaceScoped() bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return false
}
func (volumeAttachmentStrategy) PrepareForCreate(ctx context.Context, obj runtime.Object) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var groupVersion schema.GroupVersion
	if requestInfo, found := genericapirequest.RequestInfoFrom(ctx); found {
		groupVersion = schema.GroupVersion{Group: requestInfo.APIGroup, Version: requestInfo.APIVersion}
	}
	switch groupVersion {
	case storageapiv1beta1.SchemeGroupVersion:
	default:
		volumeAttachment := obj.(*storage.VolumeAttachment)
		volumeAttachment.Status = storage.VolumeAttachmentStatus{}
	}
}
func (volumeAttachmentStrategy) Validate(ctx context.Context, obj runtime.Object) field.ErrorList {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	volumeAttachment := obj.(*storage.VolumeAttachment)
	errs := validation.ValidateVolumeAttachment(volumeAttachment)
	var groupVersion schema.GroupVersion
	if requestInfo, found := genericapirequest.RequestInfoFrom(ctx); found {
		groupVersion = schema.GroupVersion{Group: requestInfo.APIGroup, Version: requestInfo.APIVersion}
	}
	switch groupVersion {
	case storageapiv1beta1.SchemeGroupVersion:
	default:
		errs = append(errs, validation.ValidateVolumeAttachmentV1(volumeAttachment)...)
	}
	return errs
}
func (volumeAttachmentStrategy) Canonicalize(obj runtime.Object) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
}
func (volumeAttachmentStrategy) AllowCreateOnUpdate() bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return false
}
func (volumeAttachmentStrategy) PrepareForUpdate(ctx context.Context, obj, old runtime.Object) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var groupVersion schema.GroupVersion
	if requestInfo, found := genericapirequest.RequestInfoFrom(ctx); found {
		groupVersion = schema.GroupVersion{Group: requestInfo.APIGroup, Version: requestInfo.APIVersion}
	}
	switch groupVersion {
	case storageapiv1beta1.SchemeGroupVersion:
	default:
		newVolumeAttachment := obj.(*storage.VolumeAttachment)
		oldVolumeAttachment := old.(*storage.VolumeAttachment)
		newVolumeAttachment.Status = oldVolumeAttachment.Status
	}
}
func (volumeAttachmentStrategy) ValidateUpdate(ctx context.Context, obj, old runtime.Object) field.ErrorList {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	newVolumeAttachmentObj := obj.(*storage.VolumeAttachment)
	oldVolumeAttachmentObj := old.(*storage.VolumeAttachment)
	errorList := validation.ValidateVolumeAttachment(newVolumeAttachmentObj)
	return append(errorList, validation.ValidateVolumeAttachmentUpdate(newVolumeAttachmentObj, oldVolumeAttachmentObj)...)
}
func (volumeAttachmentStrategy) AllowUnconditionalUpdate() bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return false
}

type volumeAttachmentStatusStrategy struct{ volumeAttachmentStrategy }

var StatusStrategy = volumeAttachmentStatusStrategy{Strategy}

func (volumeAttachmentStatusStrategy) PrepareForUpdate(ctx context.Context, obj, old runtime.Object) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	newVolumeAttachment := obj.(*storage.VolumeAttachment)
	oldVolumeAttachment := old.(*storage.VolumeAttachment)
	newVolumeAttachment.Spec = oldVolumeAttachment.Spec
	oldMeta := oldVolumeAttachment.ObjectMeta
	newMeta := &newVolumeAttachment.ObjectMeta
	newMeta.SetDeletionTimestamp(oldMeta.GetDeletionTimestamp())
	newMeta.SetGeneration(oldMeta.GetGeneration())
	newMeta.SetSelfLink(oldMeta.GetSelfLink())
	newMeta.SetLabels(oldMeta.GetLabels())
	newMeta.SetAnnotations(oldMeta.GetAnnotations())
	newMeta.SetFinalizers(oldMeta.GetFinalizers())
	newMeta.SetOwnerReferences(oldMeta.GetOwnerReferences())
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
