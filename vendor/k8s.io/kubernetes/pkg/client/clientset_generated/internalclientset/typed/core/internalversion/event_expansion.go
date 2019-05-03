package internalversion

import (
 "fmt"
 "k8s.io/api/core/v1"
 metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
 "k8s.io/apimachinery/pkg/fields"
 "k8s.io/apimachinery/pkg/runtime"
 "k8s.io/apimachinery/pkg/types"
 "k8s.io/kubernetes/pkg/api/ref"
 api "k8s.io/kubernetes/pkg/apis/core"
 k8s_api_v1 "k8s.io/kubernetes/pkg/apis/core/v1"
)

type EventExpansion interface {
 CreateWithEventNamespace(event *api.Event) (*api.Event, error)
 UpdateWithEventNamespace(event *api.Event) (*api.Event, error)
 PatchWithEventNamespace(event *api.Event, data []byte) (*api.Event, error)
 Search(scheme *runtime.Scheme, objOrRef runtime.Object) (*api.EventList, error)
 GetFieldSelector(involvedObjectName, involvedObjectNamespace, involvedObjectKind, involvedObjectUID *string) fields.Selector
}

func (e *events) CreateWithEventNamespace(event *api.Event) (*api.Event, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if e.ns != "" && event.Namespace != e.ns {
  return nil, fmt.Errorf("can't create an event with namespace '%v' in namespace '%v'", event.Namespace, e.ns)
 }
 result := &api.Event{}
 err := e.client.Post().NamespaceIfScoped(event.Namespace, len(event.Namespace) > 0).Resource("events").Body(event).Do().Into(result)
 return result, err
}
func (e *events) UpdateWithEventNamespace(event *api.Event) (*api.Event, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result := &api.Event{}
 err := e.client.Put().NamespaceIfScoped(event.Namespace, len(event.Namespace) > 0).Resource("events").Name(event.Name).Body(event).Do().Into(result)
 return result, err
}
func (e *events) PatchWithEventNamespace(incompleteEvent *api.Event, data []byte) (*api.Event, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if e.ns != "" && incompleteEvent.Namespace != e.ns {
  return nil, fmt.Errorf("can't patch an event with namespace '%v' in namespace '%v'", incompleteEvent.Namespace, e.ns)
 }
 result := &api.Event{}
 err := e.client.Patch(types.StrategicMergePatchType).NamespaceIfScoped(incompleteEvent.Namespace, len(incompleteEvent.Namespace) > 0).Resource("events").Name(incompleteEvent.Name).Body(data).Do().Into(result)
 return result, err
}
func (e *events) Search(scheme *runtime.Scheme, objOrRef runtime.Object) (*api.EventList, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 ref, err := ref.GetReference(scheme, objOrRef)
 if err != nil {
  return nil, err
 }
 if e.ns != "" && ref.Namespace != e.ns {
  return nil, fmt.Errorf("won't be able to find any events of namespace '%v' in namespace '%v'", ref.Namespace, e.ns)
 }
 stringRefKind := string(ref.Kind)
 var refKind *string
 if stringRefKind != "" {
  refKind = &stringRefKind
 }
 stringRefUID := string(ref.UID)
 var refUID *string
 if stringRefUID != "" {
  refUID = &stringRefUID
 }
 fieldSelector := e.GetFieldSelector(&ref.Name, &ref.Namespace, refKind, refUID)
 return e.List(metav1.ListOptions{FieldSelector: fieldSelector.String()})
}
func (e *events) GetFieldSelector(involvedObjectName, involvedObjectNamespace, involvedObjectKind, involvedObjectUID *string) fields.Selector {
 _logClusterCodePath()
 defer _logClusterCodePath()
 apiVersion := e.client.APIVersion().String()
 field := fields.Set{}
 if involvedObjectName != nil {
  field[GetInvolvedObjectNameFieldLabel(apiVersion)] = *involvedObjectName
 }
 if involvedObjectNamespace != nil {
  field["involvedObject.namespace"] = *involvedObjectNamespace
 }
 if involvedObjectKind != nil {
  field["involvedObject.kind"] = *involvedObjectKind
 }
 if involvedObjectUID != nil {
  field["involvedObject.uid"] = *involvedObjectUID
 }
 return field.AsSelector()
}
func GetInvolvedObjectNameFieldLabel(version string) string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return "involvedObject.name"
}

type EventSinkImpl struct{ Interface EventInterface }

func (e *EventSinkImpl) Create(event *v1.Event) (*v1.Event, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 internalEvent := &api.Event{}
 err := k8s_api_v1.Convert_v1_Event_To_core_Event(event, internalEvent, nil)
 if err != nil {
  return nil, err
 }
 _, err = e.Interface.CreateWithEventNamespace(internalEvent)
 if err != nil {
  return nil, err
 }
 return event, nil
}
func (e *EventSinkImpl) Update(event *v1.Event) (*v1.Event, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 internalEvent := &api.Event{}
 err := k8s_api_v1.Convert_v1_Event_To_core_Event(event, internalEvent, nil)
 if err != nil {
  return nil, err
 }
 _, err = e.Interface.UpdateWithEventNamespace(internalEvent)
 if err != nil {
  return nil, err
 }
 return event, nil
}
func (e *EventSinkImpl) Patch(event *v1.Event, data []byte) (*v1.Event, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 internalEvent := &api.Event{}
 err := k8s_api_v1.Convert_v1_Event_To_core_Event(event, internalEvent, nil)
 if err != nil {
  return nil, err
 }
 internalEvent, err = e.Interface.PatchWithEventNamespace(internalEvent, data)
 if err != nil {
  return nil, err
 }
 externalEvent := &v1.Event{}
 err = k8s_api_v1.Convert_core_Event_To_v1_Event(internalEvent, externalEvent, nil)
 if err != nil {
  return event, nil
 }
 return externalEvent, nil
}
