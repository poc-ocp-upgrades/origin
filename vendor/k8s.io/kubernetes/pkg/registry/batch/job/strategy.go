package job

import (
	"context"
	"fmt"
	goformat "fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"k8s.io/apiserver/pkg/registry/generic"
	"k8s.io/apiserver/pkg/registry/rest"
	"k8s.io/apiserver/pkg/storage"
	"k8s.io/apiserver/pkg/storage/names"
	utilfeature "k8s.io/apiserver/pkg/util/feature"
	"k8s.io/kubernetes/pkg/api/legacyscheme"
	"k8s.io/kubernetes/pkg/api/pod"
	"k8s.io/kubernetes/pkg/apis/batch"
	"k8s.io/kubernetes/pkg/apis/batch/validation"
	"k8s.io/kubernetes/pkg/features"
	goos "os"
	godefaultruntime "runtime"
	"strconv"
	gotime "time"
)

type jobStrategy struct {
	runtime.ObjectTyper
	names.NameGenerator
}

var Strategy = jobStrategy{legacyscheme.Scheme, names.SimpleNameGenerator}

func (jobStrategy) DefaultGarbageCollectionPolicy(ctx context.Context) rest.GarbageCollectionPolicy {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return rest.OrphanDependents
}
func (jobStrategy) NamespaceScoped() bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return true
}
func (jobStrategy) PrepareForCreate(ctx context.Context, obj runtime.Object) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	job := obj.(*batch.Job)
	job.Status = batch.JobStatus{}
	if !utilfeature.DefaultFeatureGate.Enabled(features.TTLAfterFinished) {
		job.Spec.TTLSecondsAfterFinished = nil
	}
	pod.DropDisabledAlphaFields(&job.Spec.Template.Spec)
}
func (jobStrategy) PrepareForUpdate(ctx context.Context, obj, old runtime.Object) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	newJob := obj.(*batch.Job)
	oldJob := old.(*batch.Job)
	newJob.Status = oldJob.Status
	if !utilfeature.DefaultFeatureGate.Enabled(features.TTLAfterFinished) {
		newJob.Spec.TTLSecondsAfterFinished = nil
		oldJob.Spec.TTLSecondsAfterFinished = nil
	}
	pod.DropDisabledAlphaFields(&newJob.Spec.Template.Spec)
	pod.DropDisabledAlphaFields(&oldJob.Spec.Template.Spec)
}
func (jobStrategy) Validate(ctx context.Context, obj runtime.Object) field.ErrorList {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	job := obj.(*batch.Job)
	if job.Spec.ManualSelector == nil || *job.Spec.ManualSelector == false {
		generateSelector(job)
	}
	return validation.ValidateJob(job)
}
func generateSelector(obj *batch.Job) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if obj.Spec.Template.Labels == nil {
		obj.Spec.Template.Labels = make(map[string]string)
	}
	_, found := obj.Spec.Template.Labels["job-name"]
	if found {
	} else {
		obj.Spec.Template.Labels["job-name"] = string(obj.ObjectMeta.Name)
	}
	_, found = obj.Spec.Template.Labels["controller-uid"]
	if found {
	} else {
		obj.Spec.Template.Labels["controller-uid"] = string(obj.ObjectMeta.UID)
	}
	if obj.Spec.Selector == nil {
		obj.Spec.Selector = &metav1.LabelSelector{}
	}
	if obj.Spec.Selector.MatchLabels == nil {
		obj.Spec.Selector.MatchLabels = make(map[string]string)
	}
	if _, found := obj.Spec.Selector.MatchLabels["controller-uid"]; !found {
		obj.Spec.Selector.MatchLabels["controller-uid"] = string(obj.ObjectMeta.UID)
	}
}
func (jobStrategy) Canonicalize(obj runtime.Object) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
}
func (jobStrategy) AllowUnconditionalUpdate() bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return true
}
func (jobStrategy) AllowCreateOnUpdate() bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return false
}
func (jobStrategy) ValidateUpdate(ctx context.Context, obj, old runtime.Object) field.ErrorList {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	validationErrorList := validation.ValidateJob(obj.(*batch.Job))
	updateErrorList := validation.ValidateJobUpdate(obj.(*batch.Job), old.(*batch.Job))
	return append(validationErrorList, updateErrorList...)
}

type jobStatusStrategy struct{ jobStrategy }

var StatusStrategy = jobStatusStrategy{Strategy}

func (jobStatusStrategy) PrepareForUpdate(ctx context.Context, obj, old runtime.Object) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	newJob := obj.(*batch.Job)
	oldJob := old.(*batch.Job)
	newJob.Spec = oldJob.Spec
}
func (jobStatusStrategy) ValidateUpdate(ctx context.Context, obj, old runtime.Object) field.ErrorList {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return validation.ValidateJobUpdateStatus(obj.(*batch.Job), old.(*batch.Job))
}
func JobToSelectableFields(job *batch.Job) fields.Set {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	objectMetaFieldsSet := generic.ObjectMetaFieldsSet(&job.ObjectMeta, true)
	specificFieldsSet := fields.Set{"status.successful": strconv.Itoa(int(job.Status.Succeeded))}
	return generic.MergeFieldsSets(objectMetaFieldsSet, specificFieldsSet)
}
func GetAttrs(obj runtime.Object) (labels.Set, fields.Set, bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	job, ok := obj.(*batch.Job)
	if !ok {
		return nil, nil, false, fmt.Errorf("given object is not a job.")
	}
	return labels.Set(job.ObjectMeta.Labels), JobToSelectableFields(job), job.Initializers != nil, nil
}
func MatchJob(label labels.Selector, field fields.Selector) storage.SelectionPredicate {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return storage.SelectionPredicate{Label: label, Field: field, GetAttrs: GetAttrs}
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
