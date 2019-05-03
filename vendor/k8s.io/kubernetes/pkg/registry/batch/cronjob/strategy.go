package cronjob

import (
 "context"
 "k8s.io/apimachinery/pkg/runtime"
 "k8s.io/apimachinery/pkg/util/validation/field"
 "k8s.io/apiserver/pkg/registry/rest"
 "k8s.io/apiserver/pkg/storage/names"
 "k8s.io/kubernetes/pkg/api/legacyscheme"
 "k8s.io/kubernetes/pkg/api/pod"
 "k8s.io/kubernetes/pkg/apis/batch"
 "k8s.io/kubernetes/pkg/apis/batch/validation"
)

type cronJobStrategy struct {
 runtime.ObjectTyper
 names.NameGenerator
}

var Strategy = cronJobStrategy{legacyscheme.Scheme, names.SimpleNameGenerator}

func (cronJobStrategy) DefaultGarbageCollectionPolicy(ctx context.Context) rest.GarbageCollectionPolicy {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return rest.OrphanDependents
}
func (cronJobStrategy) NamespaceScoped() bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return true
}
func (cronJobStrategy) PrepareForCreate(ctx context.Context, obj runtime.Object) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 cronJob := obj.(*batch.CronJob)
 cronJob.Status = batch.CronJobStatus{}
 pod.DropDisabledAlphaFields(&cronJob.Spec.JobTemplate.Spec.Template.Spec)
}
func (cronJobStrategy) PrepareForUpdate(ctx context.Context, obj, old runtime.Object) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 newCronJob := obj.(*batch.CronJob)
 oldCronJob := old.(*batch.CronJob)
 newCronJob.Status = oldCronJob.Status
 pod.DropDisabledAlphaFields(&newCronJob.Spec.JobTemplate.Spec.Template.Spec)
 pod.DropDisabledAlphaFields(&oldCronJob.Spec.JobTemplate.Spec.Template.Spec)
}
func (cronJobStrategy) Validate(ctx context.Context, obj runtime.Object) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 cronJob := obj.(*batch.CronJob)
 return validation.ValidateCronJob(cronJob)
}
func (cronJobStrategy) Canonicalize(obj runtime.Object) {
 _logClusterCodePath()
 defer _logClusterCodePath()
}
func (cronJobStrategy) AllowUnconditionalUpdate() bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return true
}
func (cronJobStrategy) AllowCreateOnUpdate() bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return false
}
func (cronJobStrategy) ValidateUpdate(ctx context.Context, obj, old runtime.Object) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return validation.ValidateCronJobUpdate(obj.(*batch.CronJob), old.(*batch.CronJob))
}

type cronJobStatusStrategy struct{ cronJobStrategy }

var StatusStrategy = cronJobStatusStrategy{Strategy}

func (cronJobStatusStrategy) PrepareForUpdate(ctx context.Context, obj, old runtime.Object) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 newJob := obj.(*batch.CronJob)
 oldJob := old.(*batch.CronJob)
 newJob.Spec = oldJob.Spec
}
func (cronJobStatusStrategy) ValidateUpdate(ctx context.Context, obj, old runtime.Object) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return field.ErrorList{}
}
