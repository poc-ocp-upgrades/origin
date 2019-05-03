package json

import (
 "encoding/json"
 "fmt"
 "reflect"
 "k8s.io/apimachinery/pkg/runtime"
 "k8s.io/apimachinery/pkg/watch"
)

type WatchEvent struct {
 Type   watch.EventType      `json:"type,omitempty" description:"the type of watch event; may be ADDED, MODIFIED, DELETED, or ERROR"`
 Object runtime.RawExtension `json:"object,omitempty" description:"the object being watched; will match the type of the resource endpoint or be a Status object if the type is ERROR"`
}

func Object(encoder runtime.Encoder, event *watch.Event) (interface{}, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 obj, ok := event.Object.(runtime.Object)
 if !ok {
  return nil, fmt.Errorf("the event object cannot be safely converted to JSON: %v", reflect.TypeOf(event.Object).Name())
 }
 data, err := runtime.Encode(encoder, obj)
 if err != nil {
  return nil, err
 }
 return &WatchEvent{event.Type, runtime.RawExtension{Raw: json.RawMessage(data)}}, nil
}
