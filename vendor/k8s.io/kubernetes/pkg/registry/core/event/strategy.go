package event

import (
	"context"
	"fmt"
	goformat "fmt"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"k8s.io/apiserver/pkg/registry/generic"
	"k8s.io/apiserver/pkg/registry/rest"
	"k8s.io/apiserver/pkg/storage"
	"k8s.io/apiserver/pkg/storage/names"
	"k8s.io/kubernetes/pkg/api/legacyscheme"
	api "k8s.io/kubernetes/pkg/apis/core"
	"k8s.io/kubernetes/pkg/apis/core/validation"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

type eventStrategy struct {
	runtime.ObjectTyper
	names.NameGenerator
}

var Strategy = eventStrategy{legacyscheme.Scheme, names.SimpleNameGenerator}

func (eventStrategy) DefaultGarbageCollectionPolicy(ctx context.Context) rest.GarbageCollectionPolicy {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return rest.Unsupported
}
func (eventStrategy) NamespaceScoped() bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return true
}
func (eventStrategy) PrepareForCreate(ctx context.Context, obj runtime.Object) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
}
func (eventStrategy) PrepareForUpdate(ctx context.Context, obj, old runtime.Object) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
}
func (eventStrategy) Validate(ctx context.Context, obj runtime.Object) field.ErrorList {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	event := obj.(*api.Event)
	return validation.ValidateEvent(event)
}
func (eventStrategy) Canonicalize(obj runtime.Object) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
}
func (eventStrategy) AllowCreateOnUpdate() bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return true
}
func (eventStrategy) ValidateUpdate(ctx context.Context, obj, old runtime.Object) field.ErrorList {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	event := obj.(*api.Event)
	return validation.ValidateEvent(event)
}
func (eventStrategy) AllowUnconditionalUpdate() bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return true
}
func GetAttrs(obj runtime.Object) (labels.Set, fields.Set, bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	event, ok := obj.(*api.Event)
	if !ok {
		return nil, nil, false, fmt.Errorf("not an event")
	}
	return labels.Set(event.Labels), EventToSelectableFields(event), event.Initializers != nil, nil
}
func MatchEvent(label labels.Selector, field fields.Selector) storage.SelectionPredicate {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return storage.SelectionPredicate{Label: label, Field: field, GetAttrs: GetAttrs}
}
func EventToSelectableFields(event *api.Event) fields.Set {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	objectMetaFieldsSet := generic.ObjectMetaFieldsSet(&event.ObjectMeta, true)
	specificFieldsSet := fields.Set{"involvedObject.kind": event.InvolvedObject.Kind, "involvedObject.namespace": event.InvolvedObject.Namespace, "involvedObject.name": event.InvolvedObject.Name, "involvedObject.uid": string(event.InvolvedObject.UID), "involvedObject.apiVersion": event.InvolvedObject.APIVersion, "involvedObject.resourceVersion": event.InvolvedObject.ResourceVersion, "involvedObject.fieldPath": event.InvolvedObject.FieldPath, "reason": event.Reason, "source": event.Source.Component, "type": event.Type}
	return generic.MergeFieldsSets(objectMetaFieldsSet, specificFieldsSet)
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
